import { createHash, createPublicKey } from 'node:crypto';

export const VERSION = '002';
export const HEX32 = /^[0-9a-f]{64}$/;

const PROPOSAL_DOMAIN = Buffer.from('IK:CROSSING:002\0', 'utf8');
const STATE_DOMAIN = Buffer.from('IK:RELATION_STATE:002\0', 'utf8');
const HANDOFF_DOMAIN = Buffer.from('IK:HANDOFF:002\0', 'utf8');
const CLOSURE_DOMAIN = Buffer.from('IK:CLOSURE:002\0', 'utf8');
const FRESH_DOMAIN = Buffer.from('IK:FRESH_ADMISSION:002\0', 'utf8');
const FRESH_SURFACE_DOMAIN = Buffer.from('IK:FRESH_SURFACE:002\0', 'utf8');
const ED25519_SPKI_PREFIX = Buffer.from('302a300506032b6570032100', 'hex');

const STATE_KEYS = [
  'version', 'relation_id', 'profile_commitment', 'epoch', 'status',
  'head_a', 'verifier_a', 'nonce_a', 'head_b', 'verifier_b', 'nonce_b',
  'handoff_domain'
].sort();
const PROPOSAL_KEYS = [
  'version', 'relation_id', 'profile_commitment', 'epoch', 'prev_head_a',
  'prev_head_b', 'resolution', 'next_head_a', 'next_head_b', 'nonce_a',
  'nonce_b', 'handoff_manifest', 'signature_a', 'signature_b'
].sort();
const MANIFEST_KEYS = [
  'successor_input_commitment', 'successor_admission_commitment',
  'preserved', 'transformed', 'ended', 'refused', 'unresolved'
].sort();
const TRANSFORM_KEYS = ['from', 'to'].sort();
const FRESH_KEYS = [
  'version', 'relation_id', 'profile_commitment', 'invocation_id',
  'predecessor_commitment', 'handoff_commitment', 'input_commitment',
  'bounds_commitment', 'head_a', 'verifier_a', 'head_b', 'verifier_b',
  'signature_a', 'signature_b'
].sort();

function u32(value) {
  if (!Number.isInteger(value) || value < 0 || value > 0xffffffff) throw new RangeError('u32');
  const out = Buffer.alloc(4);
  out.writeUInt32BE(value);
  return out;
}

export function u64(value) {
  const n = typeof value === 'bigint' ? value : BigInt(value);
  if (n < 0n || n > 0xffffffffffffffffn) throw new RangeError('u64');
  const out = Buffer.alloc(8);
  out.writeBigUInt64BE(n);
  return out;
}

function bytes(value) {
  const raw = Buffer.isBuffer(value) ? value : Buffer.from(value);
  return Buffer.concat([u32(raw.length), raw]);
}

const text = value => bytes(Buffer.from(value, 'utf8'));

function hash32(value) {
  if (!HEX32.test(value)) throw new TypeError('hash32');
  return Buffer.from(value, 'hex');
}

function optionalHash(value) {
  return value === null
    ? Buffer.from([0])
    : Buffer.concat([Buffer.from([1]), hash32(value)]);
}

function resolutionByte(value) {
  if (value === '1') return Buffer.from([1]);
  if (value === '0') return Buffer.from([0]);
  if (value === '↻') return Buffer.from([2]);
  throw new TypeError('resolution');
}

const hashList = values => Buffer.concat([u32(values.length), ...values.map(hash32)]);

function transformedList(values) {
  return Buffer.concat([
    u32(values.length),
    ...values.flatMap(({ from, to }) => [hash32(from), hash32(to)])
  ]);
}

function optionalManifest(value) {
  return value === null
    ? Buffer.from([0])
    : Buffer.concat([Buffer.from([1]), canonicalHandoffBytes(value)]);
}

export function canonicalHandoffBytes(handoff) {
  return Buffer.concat([
    HANDOFF_DOMAIN,
    hash32(handoff.successor_input_commitment),
    hash32(handoff.successor_admission_commitment),
    hashList(handoff.preserved),
    transformedList(handoff.transformed),
    hashList(handoff.ended),
    hashList(handoff.refused),
    hashList(handoff.unresolved)
  ]);
}

export function canonicalFreshSurfaceBytes(admission) {
  return Buffer.concat([
    FRESH_SURFACE_DOMAIN,
    text(admission.version),
    hash32(admission.relation_id),
    hash32(admission.profile_commitment),
    hash32(admission.invocation_id),
    hash32(admission.input_commitment),
    hash32(admission.bounds_commitment),
    hash32(admission.head_a),
    bytes(Buffer.from(admission.verifier_a, 'base64')),
    hash32(admission.head_b),
    bytes(Buffer.from(admission.verifier_b, 'base64'))
  ]);
}

export function canonicalProposalBytes(proposal) {
  return Buffer.concat([
    PROPOSAL_DOMAIN,
    text(proposal.version),
    hash32(proposal.relation_id),
    hash32(proposal.profile_commitment),
    u64(proposal.epoch),
    hash32(proposal.prev_head_a),
    hash32(proposal.prev_head_b),
    resolutionByte(proposal.resolution),
    optionalHash(proposal.next_head_a),
    optionalHash(proposal.next_head_b),
    u64(proposal.nonce_a),
    u64(proposal.nonce_b),
    optionalManifest(proposal.handoff_manifest)
  ]);
}

export function canonicalStateBytes(state) {
  const status = state.status === 'ACTIVE'
    ? 1
    : state.status === 'ENDED'
      ? 0
      : (() => { throw new TypeError('status'); })();
  return Buffer.concat([
    STATE_DOMAIN,
    text(state.version),
    hash32(state.relation_id),
    hash32(state.profile_commitment),
    u64(state.epoch),
    Buffer.from([status]),
    hash32(state.head_a),
    bytes(Buffer.from(state.verifier_a, 'base64')),
    u64(state.nonce_a),
    hash32(state.head_b),
    bytes(Buffer.from(state.verifier_b, 'base64')),
    u64(state.nonce_b),
    hashList(state.handoff_domain)
  ]);
}

export function canonicalFreshAdmissionBytes(admission) {
  return Buffer.concat([
    FRESH_DOMAIN,
    text(admission.version),
    hash32(admission.relation_id),
    hash32(admission.profile_commitment),
    hash32(admission.invocation_id),
    hash32(admission.predecessor_commitment),
    hash32(admission.handoff_commitment),
    hash32(admission.input_commitment),
    hash32(admission.bounds_commitment),
    hash32(admission.head_a),
    bytes(Buffer.from(admission.verifier_a, 'base64')),
    hash32(admission.head_b),
    bytes(Buffer.from(admission.verifier_b, 'base64'))
  ]);
}

export const sha256 = data => createHash('sha256').update(data).digest('hex');
export const stateHash = state => sha256(canonicalStateBytes(state));
export const proposalHash = proposal => sha256(canonicalProposalBytes(proposal));
export const handoffHash = handoff => sha256(canonicalHandoffBytes(handoff));
export const freshAdmissionHash = admission => sha256(canonicalFreshAdmissionBytes(admission));
export const freshSurfaceHash = admission => sha256(canonicalFreshSurfaceBytes(admission));

export function closureHash(fields) {
  return sha256(Buffer.concat([
    CLOSURE_DOMAIN,
    hash32(fields.prior_state_hash),
    hash32(fields.proposal_hash),
    resolutionByte(fields.resolution),
    hash32(fields.next_state_hash),
    optionalHash(fields.handoff_commitment)
  ]));
}

export function stable(value) {
  if (value === null || typeof value !== 'object') return JSON.stringify(value);
  if (Array.isArray(value)) return `[${value.map(stable).join(',')}]`;
  return `{${Object.keys(value).sort().map(key => `${JSON.stringify(key)}:${stable(value[key])}`).join(',')}}`;
}

export const exactKeys = (value, keys) => value
  && typeof value === 'object'
  && !Array.isArray(value)
  && JSON.stringify(Object.keys(value).sort()) === JSON.stringify(keys);
export const integer = value => Number.isSafeInteger(value) && value >= 0;
export const isNextCounter = (current, candidate) => integer(current)
  && integer(candidate)
  && current < Number.MAX_SAFE_INTEGER
  && candidate === current + 1;
export const sortedUniqueHashes = values => Array.isArray(values)
  && values.every(value => typeof value === 'string' && HEX32.test(value))
  && values.every((value, index) => index === 0 || values[index - 1] < value);

const canonicalBase64 = value => {
  if (typeof value !== 'string' || value.length === 0) return false;
  try { return Buffer.from(value, 'base64').toString('base64') === value; } catch { return false; }
};

export const publicKeyMaterial = value => {
  if (!canonicalBase64(value)) return null;
  try {
    const submitted = Buffer.from(value, 'base64');
    const key = createPublicKey({
      key: submitted,
      format: 'der',
      type: 'spki'
    });
    if (key.asymmetricKeyType !== 'ed25519') return null;
    const canonical = key.export({ format: 'der', type: 'spki' });
    if (!Buffer.isBuffer(canonical)
        || !canonical.equals(submitted)
        || canonical.length !== ED25519_SPKI_PREFIX.length + 32
        || !canonical.subarray(0, ED25519_SPKI_PREFIX.length).equals(ED25519_SPKI_PREFIX)) return null;
    return Buffer.from(canonical.subarray(ED25519_SPKI_PREFIX.length));
  } catch {
    return null;
  }
};

export const publicKey = value => publicKeyMaterial(value) !== null;

export function samePublicKey(left, right) {
  const leftMaterial = publicKeyMaterial(left);
  const rightMaterial = publicKeyMaterial(right);
  return leftMaterial !== null
    && rightMaterial !== null
    && leftMaterial.equals(rightMaterial);
}

export const signature = value => canonicalBase64(value) && Buffer.from(value, 'base64').length === 64;

export function validManifestShape(handoff) {
  return exactKeys(handoff, MANIFEST_KEYS)
    && HEX32.test(handoff.successor_input_commitment)
    && HEX32.test(handoff.successor_admission_commitment)
    && Array.isArray(handoff.preserved) && handoff.preserved.every(value => typeof value === 'string' && HEX32.test(value))
    && Array.isArray(handoff.transformed) && handoff.transformed.every(value => exactKeys(value, TRANSFORM_KEYS) && HEX32.test(value.from) && HEX32.test(value.to))
    && Array.isArray(handoff.ended) && handoff.ended.every(value => typeof value === 'string' && HEX32.test(value))
    && Array.isArray(handoff.refused) && handoff.refused.every(value => typeof value === 'string' && HEX32.test(value))
    && Array.isArray(handoff.unresolved) && handoff.unresolved.every(value => typeof value === 'string' && HEX32.test(value));
}

export function validState(state) {
  return exactKeys(state, STATE_KEYS)
    && typeof state.version === 'string'
    && HEX32.test(state.relation_id)
    && HEX32.test(state.profile_commitment)
    && integer(state.epoch)
    && ['ACTIVE', 'ENDED'].includes(state.status)
    && HEX32.test(state.head_a)
    && publicKey(state.verifier_a)
    && integer(state.nonce_a)
    && HEX32.test(state.head_b)
    && publicKey(state.verifier_b)
    && integer(state.nonce_b)
    && sortedUniqueHashes(state.handoff_domain);
}

export function validProposalShape(proposal) {
  return exactKeys(proposal, PROPOSAL_KEYS)
    && typeof proposal.version === 'string'
    && HEX32.test(proposal.relation_id)
    && HEX32.test(proposal.profile_commitment)
    && integer(proposal.epoch)
    && HEX32.test(proposal.prev_head_a)
    && HEX32.test(proposal.prev_head_b)
    && typeof proposal.resolution === 'string'
    && (proposal.next_head_a === null || HEX32.test(proposal.next_head_a))
    && (proposal.next_head_b === null || HEX32.test(proposal.next_head_b))
    && integer(proposal.nonce_a)
    && integer(proposal.nonce_b)
    && (proposal.handoff_manifest === null || validManifestShape(proposal.handoff_manifest))
    && signature(proposal.signature_a)
    && signature(proposal.signature_b);
}

export function validFreshAdmission(admission) {
  return exactKeys(admission, FRESH_KEYS)
    && typeof admission.version === 'string'
    && HEX32.test(admission.relation_id)
    && HEX32.test(admission.profile_commitment)
    && HEX32.test(admission.invocation_id)
    && HEX32.test(admission.predecessor_commitment)
    && HEX32.test(admission.handoff_commitment)
    && HEX32.test(admission.input_commitment)
    && HEX32.test(admission.bounds_commitment)
    && HEX32.test(admission.head_a)
    && publicKey(admission.verifier_a)
    && HEX32.test(admission.head_b)
    && publicKey(admission.verifier_b)
    && signature(admission.signature_a)
    && signature(admission.signature_b);
}

export function manifestRule(domain, handoff) {
  const canonicalArrays = sortedUniqueHashes(handoff.preserved)
    && sortedUniqueHashes(handoff.ended)
    && sortedUniqueHashes(handoff.refused)
    && sortedUniqueHashes(handoff.unresolved)
    && handoff.transformed.every((value, index, all) => index === 0
      || all[index - 1].from < value.from
      || (all[index - 1].from === value.from && all[index - 1].to < value.to));
  if (!canonicalArrays) return false;

  const sources = [
    ...handoff.preserved,
    ...handoff.transformed.map(value => value.from),
    ...handoff.ended,
    ...handoff.refused,
    ...handoff.unresolved
  ].sort();
  if (JSON.stringify(sources) !== JSON.stringify(domain)) return false;
  if (new Set(sources).size !== sources.length) return false;

  const carried = [
    ...handoff.preserved,
    ...handoff.transformed.map(value => value.to),
    ...handoff.unresolved
  ].sort();
  if (new Set(carried).size !== carried.length) return false;
  const prohibited = new Set([...handoff.ended, ...handoff.refused]);
  return carried.every(value => !prohibited.has(value));
}

export function successorDomain(handoff) {
  return [
    ...handoff.preserved,
    ...handoff.transformed.map(value => value.to),
    ...handoff.unresolved
  ].sort();
}
