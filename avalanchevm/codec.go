package avalanchevm

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

// canonicalJSON is deliberately restricted to typed structs and slices. The
// portable kernel has its own binary canonical encodings; this codec is only
// the Avalanche operation/block/storage envelope around those semantics.
func canonicalJSON(value any) ([]byte, error) {
	var out bytes.Buffer
	encoder := json.NewEncoder(&out)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(value); err != nil {
		return nil, err
	}
	data := out.Bytes()
	if len(data) == 0 || data[len(data)-1] != '\n' {
		return nil, errors.New("canonical JSON encoder omitted terminal newline")
	}
	return bytes.Clone(data[:len(data)-1]), nil
}

// inspectJSON rejects alternate numeric spellings and duplicate keys before
// typed decoding. JSON null remains legal: it is part of the IK crossing and
// closure surface and is checked by the concrete kernel type validators.
func inspectJSON(data []byte) error {
	if len(data) == 0 || !utf8.Valid(data) || bytes.HasPrefix(data, []byte{0xef, 0xbb, 0xbf}) {
		return errors.New("JSON must be non-empty UTF-8 without BOM")
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	var readValue func() error
	readValue = func() error {
		token, err := decoder.Token()
		if err != nil {
			return err
		}
		if number, ok := token.(json.Number); ok {
			spelling := number.String()
			if strings.HasPrefix(spelling, "-") || strings.ContainsAny(spelling, ".eE") || (len(spelling) > 1 && spelling[0] == '0') {
				return fmt.Errorf("non-canonical unsigned integer %q", spelling)
			}
			return nil
		}
		delim, ok := token.(json.Delim)
		if !ok {
			return nil
		}
		switch delim {
		case '{':
			seen := make(map[string]struct{})
			for decoder.More() {
				keyToken, err := decoder.Token()
				if err != nil {
					return err
				}
				key, ok := keyToken.(string)
				if !ok {
					return errors.New("JSON object key is not a string")
				}
				if _, duplicate := seen[key]; duplicate {
					return fmt.Errorf("duplicate JSON key %q", key)
				}
				seen[key] = struct{}{}
				if err := readValue(); err != nil {
					return err
				}
			}
			end, err := decoder.Token()
			if err != nil || end != json.Delim('}') {
				return errors.New("unterminated JSON object")
			}
		case '[':
			for decoder.More() {
				if err := readValue(); err != nil {
					return err
				}
			}
			end, err := decoder.Token()
			if err != nil || end != json.Delim(']') {
				return errors.New("unterminated JSON array")
			}
		default:
			return fmt.Errorf("unexpected JSON delimiter %q", delim)
		}
		return nil
	}
	if err := readValue(); err != nil {
		return err
	}
	if _, err := decoder.Token(); !errors.Is(err, io.EOF) {
		if err == nil {
			return errors.New("unexpected trailing JSON value")
		}
		return err
	}
	return nil
}

func decodeStrict(data []byte, destination any) error {
	if err := inspectJSON(data); err != nil {
		return err
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	decoder.UseNumber()
	if err := decoder.Decode(destination); err != nil {
		return err
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		if err == nil {
			return errors.New("unexpected trailing JSON value")
		}
		return err
	}
	return nil
}

func decodeCanonical(data []byte, destination any) error {
	if err := decodeStrict(data, destination); err != nil {
		return err
	}
	encoded, err := canonicalJSON(destination)
	if err != nil {
		return err
	}
	if !bytes.Equal(data, encoded) {
		return errors.New("JSON is not CONTINUITY_CANONICAL_JSON_001")
	}
	return nil
}

func domainHash(domain string, data []byte) string {
	hash := sha256.New()
	hash.Write([]byte(domain))
	hash.Write([]byte{0})
	var size [8]byte
	binary.BigEndian.PutUint64(size[:], uint64(len(data)))
	hash.Write(size[:])
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}

func requireHash(name, value string) error {
	if len(value) != sha256.Size*2 {
		return fmt.Errorf("%s must be 64 lowercase hexadecimal characters", name)
	}
	decoded, err := hex.DecodeString(value)
	if err != nil || hex.EncodeToString(decoded) != value {
		return fmt.Errorf("%s must be 64 lowercase hexadecimal characters", name)
	}
	return nil
}
