import { createPublicKey, verify } from 'node:crypto';
import {
  VERSION,
  canonicalFreshAdmissionBytes,
  canonicalProposalBytes,
  closureHash,
  freshAdmissionHash,
  freshSurfaceHash,
  handoffHash,
  proposalHash,
  sha256,
  stateHash,
  stable
} from './ik002.mjs';
import {
  exactKeys,
  isNextCounter,
  manifestRule,
  publicKey,
  samePublicKey,
  signature,
  sortedUniqueHashes,
  successorDomain,
  validFreshAdmission,
  validManifestShape,
  validProposalShape,
  validState
} from './ik002.mjs';

export { VERSION, stable };

export const PROTOCOL_VERSION = 'CONTINUITY_L1_001';
const HEX32 = /^[0-9a-f]{64}$/;

const OPEN_DOMAIN = Buffer.from('IK:OPEN_RELATION:001\0', 'utf8');
const WITHDRAW_DOMAIN = Buffer.from('IK:WITHDRAW_CHAMBER:001\0', 'utf8');
const WITHDRAW_CLOSURE_DOMAIN = Buffer.from('IK:WITHDRAW_CLOSURE:001\0', 'utf8');

const CHAIN_KEYS = [
  'protocol_version', 'kernel_version', 'profile_commitment',
  'relations', 'closures', 'successors'
].sort();
const OPEN_KEYS = [
  'version', 'relation_id', 'profile_commitment', 'invocation_id',
  'input_commitment', 'bounds_commitment', 'head_a', 'verifier_a',
  'head_b', 'verifier_b', 'handoff_domain', 'signature_a', 'signature_b'
].sort();
const WITHDRAW_KEYS = [
  'version', 'relation_id', 'profile_commitment', 'epoch', 'head_a',
  'head_b', 'side', 'nonce', 'signature'
].sort();
const SUCCESSOR_KEYS = [
  'predecessor_closure_hash', 'handoff_manifest', 'fresh_admission'
].sort();
const CLOSURE_RECORD_KEYS = [
  'closure_hash', 'relation_id', 'epoch', 'prior_state_hash',
  'proposal_hash', 'resolution', 'next_state_hash', 'handoff_commitment'
].sort();
const SUCCESSOR_RECORD_KEYS = [
  'predecessor_closure_hash', 'admission_hash', 'relation_id'
].sort();

export function parseStrictJson(source) {
  if (typeof source !== 'string') throw new TypeError('json source');
  let cursor = 0;
  const fail = reason => { throw new SyntaxError(`${reason} at byte ${cursor}`); };
  const whitespace = () => {
    while (cursor < source.length && /[\t\n\r ]/.test(source[cursor])) cursor += 1;
  };
  const string = () => {
    if (source[cursor] !== '"') fail('expected string');
    const start = cursor;
    cursor += 1;
    while (cursor < source.length) {
      const character = source[cursor];
      cursor += 1;
      if (character === '"') return JSON.parse(source.slice(start, cursor));
      if (character === '\\') {
        if (cursor >= source.length) fail('unterminated escape');
        const escaped = source[cursor];
        cursor += 1;
        if (escaped === 'u') {
          if (!/^[0-9a-fA-F]{4}$/.test(source.slice(cursor, cursor + 4))) fail('invalid unicode escape');
          cursor += 4;
        } else if (!['"', '\\', '/', 'b', 'f', 'n', 'r', 't'].includes(escaped)) {
          fail('invalid escape');
        }
      } else if (character.charCodeAt(0) < 0x20) {
        fail('unescaped control character');
      }
    }
    fail('unterminated string');
  };
  const value = () => {
    whitespace();
    const character = source[cursor];
    if (character === '{') {
      cursor += 1;
      whitespace();
      const keys = new Set();
      if (source[cursor] === '}') { cursor += 1; return; }
      while (cursor < source.length) {
        const key = string();
        if (keys.has(key)) fail(`duplicate key ${JSON.stringify(key)}`);
        keys.add(key);
        whitespace();
        if (source[cursor] !== ':') fail('expected colon');
        cursor += 1;
        value();
        whitespace();
        if (source[cursor] === '}') { cursor += 1; return; }
        if (source[cursor] !== ',') fail('expected comma');
        cursor += 1;
        whitespace();
      }
      fail('unterminated object');
    }
    if (character === '[') {
      cursor += 1;
      whitespace();
      if (source[cursor] === ']') { cursor += 1; return; }
      while (cursor < source.length) {
        value();
        whitespace();
        if (source[cursor] === ']') { cursor += 1; return; }
        if (source[cursor] !== ',') fail('expected comma');
        cursor += 1;
      }
      fail('unterminated array');
    }
    if (character === '"') { string(); return; }
    for (const literal of ['true', 'false', 'null']) {
      if (source.startsWith(literal, cursor)) { cursor += literal.length; return; }
    }
    const number = source.slice(cursor).match(/^-?(?:0|[1-9]\d*)(?:\.\d+)?(?:[eE][+-]?\d+)?/);
    if (number) { cursor += number[0].length; return; }
    fail('invalid value');
  };
  whitespace();
  value();
  whitespace();
  if (cursor !== source.length) fail('trailing data');
  return JSON.parse(source);
}

export const Codes = Object.freeze({
  ACCEPTED: 'ACCEPTED',
  OPEN_ACCEPTED: 'OPEN_ACCEPTED',
  SUCCESSOR_ACCEPTED: 'SUCCESSOR_ACCEPTED',
  WITHDRAW_ACCEPTED: 'WITHDRAW_ACCEPTED',
  MALFORMED_CHAIN_STATE: 'MALFORMED_CHAIN_STATE',
  MALFORMED_OPEN: 'MALFORMED_OPEN',
  MALFORMED_PROPOSAL: 'MALFORMED_PROPOSAL',
  MALFORMED_SUCCESSOR: 'MALFORMED_SUCCESSOR',
  MALFORMED_WITHDRAWAL: 'MALFORMED_WITHDRAWAL',
  UNSUPPORTED_VERSION: 'UNSUPPORTED_VERSION',
  PROFILE_MISMATCH: 'PROFILE_MISMATCH',
  RELATION_NOT_FOUND: 'RELATION_NOT_FOUND',
  RELATION_EXISTS: 'RELATION_EXISTS',
  VERIFIERS_NOT_DISTINCT: 'VERIFIERS_NOT_DISTINCT',
  RELATION_ENDED: 'RELATION_ENDED',
  STALE_EPOCH: 'STALE_EPOCH',
  STALE_HEAD_A: 'STALE_HEAD_A',
  STALE_HEAD_B: 'STALE_HEAD_B',
  NONCE_A_MISMATCH: 'NONCE_A_MISMATCH',
  NONCE_B_MISMATCH: 'NONCE_B_MISMATCH',
  RESOLUTION_INVALID: 'RESOLUTION_INVALID',
  OUTPUT_RULE_VIOLATION: 'OUTPUT_RULE_VIOLATION',
  HANDOFF_RULE_VIOLATION: 'HANDOFF_RULE_VIOLATION',
  SIGNATURE_A_INVALID: 'SIGNATURE_A_INVALID',
  SIGNATURE_B_INVALID: 'SIGNATURE_B_INVALID',
  OPEN_SIGNATURE_A_INVALID: 'OPEN_SIGNATURE_A_INVALID',
  OPEN_SIGNATURE_B_INVALID: 'OPEN_SIGNATURE_B_INVALID',
  PREDECESSOR_CLOSURE_NOT_FOUND: 'PREDECESSOR_CLOSURE_NOT_FOUND',
  PREDECESSOR_NOT_RENEW: 'PREDECESSOR_NOT_RENEW',
  PREDECESSOR_ALREADY_CONSUMED: 'PREDECESSOR_ALREADY_CONSUMED',
  PREDECESSOR_COMMITMENT_MISMATCH: 'PREDECESSOR_COMMITMENT_MISMATCH',
  HANDOFF_COMMITMENT_MISMATCH: 'HANDOFF_COMMITMENT_MISMATCH',
  FRESH_INPUT_MISMATCH: 'FRESH_INPUT_MISMATCH',
  SUCCESSOR_ADMISSION_COMMITMENT_MISMATCH: 'SUCCESSOR_ADMISSION_COMMITMENT_MISMATCH',
  FRESH_SIGNATURE_A_INVALID: 'FRESH_SIGNATURE_A_INVALID',
  FRESH_SIGNATURE_B_INVALID: 'FRESH_SIGNATURE_B_INVALID',
  SIDE_INVALID: 'SIDE_INVALID',
  WITHDRAW_NONCE_MISMATCH: 'WITHDRAW_NONCE_MISMATCH',
  WITHDRAW_SIGNATURE_INVALID: 'WITHDRAW_SIGNATURE_INVALID'
});

export const PRECEDENCE = Object.freeze({
  open: Object.freeze([
    Codes.MALFORMED_CHAIN_STATE, Codes.MALFORMED_OPEN,
    Codes.UNSUPPORTED_VERSION, Codes.PROFILE_MISMATCH,
    Codes.RELATION_EXISTS, Codes.VERIFIERS_NOT_DISTINCT,
    Codes.OPEN_SIGNATURE_A_INVALID, Codes.OPEN_SIGNATURE_B_INVALID
  ]),
  crossing: Object.freeze([
    Codes.MALFORMED_CHAIN_STATE, Codes.MALFORMED_PROPOSAL,
    Codes.UNSUPPORTED_VERSION, Codes.RELATION_NOT_FOUND,
    Codes.PROFILE_MISMATCH, Codes.RELATION_ENDED, Codes.STALE_EPOCH,
    Codes.STALE_HEAD_A, Codes.STALE_HEAD_B, Codes.NONCE_A_MISMATCH,
    Codes.NONCE_B_MISMATCH, Codes.RESOLUTION_INVALID,
    Codes.OUTPUT_RULE_VIOLATION, Codes.HANDOFF_RULE_VIOLATION,
    Codes.SIGNATURE_A_INVALID, Codes.SIGNATURE_B_INVALID
  ]),
  successor: Object.freeze([
    Codes.MALFORMED_CHAIN_STATE, Codes.MALFORMED_SUCCESSOR,
    Codes.UNSUPPORTED_VERSION, Codes.PREDECESSOR_CLOSURE_NOT_FOUND,
    Codes.PREDECESSOR_NOT_RENEW, Codes.PREDECESSOR_ALREADY_CONSUMED,
    Codes.RELATION_EXISTS, Codes.PROFILE_MISMATCH,
    Codes.PREDECESSOR_COMMITMENT_MISMATCH,
    Codes.HANDOFF_COMMITMENT_MISMATCH, Codes.FRESH_INPUT_MISMATCH,
    Codes.SUCCESSOR_ADMISSION_COMMITMENT_MISMATCH,
    Codes.VERIFIERS_NOT_DISTINCT, Codes.FRESH_SIGNATURE_A_INVALID,
    Codes.FRESH_SIGNATURE_B_INVALID
  ]),
  withdraw: Object.freeze([
    Codes.MALFORMED_CHAIN_STATE, Codes.MALFORMED_WITHDRAWAL,
    Codes.UNSUPPORTED_VERSION, Codes.RELATION_NOT_FOUND,
    Codes.PROFILE_MISMATCH, Codes.RELATION_ENDED, Codes.STALE_EPOCH,
    Codes.STALE_HEAD_A, Codes.STALE_HEAD_B, Codes.SIDE_INVALID,
    Codes.WITHDRAW_NONCE_MISMATCH, Codes.WITHDRAW_SIGNATURE_INVALID
  ])
});

const clone = value => structuredClone(value);
const integer = value => Number.isSafeInteger(value) && value >= 0;

function u32(value) {
  if (!Number.isInteger(value) || value < 0 || value > 0xffffffff) throw new RangeError('u32');
  const out = Buffer.alloc(4);
  out.writeUInt32BE(value);
  return out;
}

function u64(value) {
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
  if (typeof value !== 'string' || !HEX32.test(value)) throw new TypeError('hash32');
  return Buffer.from(value, 'hex');
}

const hashList = values => Buffer.concat([u32(values.length), ...values.map(hash32)]);

function sideByte(side) {
  if (side === 'A') return Buffer.from([0]);
  if (side === 'B') return Buffer.from([1]);
  throw new TypeError('side');
}

function verifySignature(publicDerB64, payload, signatureB64) {
  try {
    const key = createPublicKey({
      key: Buffer.from(publicDerB64, 'base64'),
      format: 'der',
      type: 'spki'
    });
    return verify(null, payload, key, Buffer.from(signatureB64, 'base64'));
  } catch {
    return false;
  }
}

export function canonicalOpenBytes(open) {
  return Buffer.concat([
    OPEN_DOMAIN,
    text(open.version),
    hash32(open.relation_id),
    hash32(open.profile_commitment),
    hash32(open.invocation_id),
    hash32(open.input_commitment),
    hash32(open.bounds_commitment),
    hash32(open.head_a),
    bytes(Buffer.from(open.verifier_a, 'base64')),
    hash32(open.head_b),
    bytes(Buffer.from(open.verifier_b, 'base64')),
    hashList(open.handoff_domain)
  ]);
}

export const openHash = open => sha256(canonicalOpenBytes(open));

export function canonicalWithdrawalBytes(withdrawal) {
  return Buffer.concat([
    WITHDRAW_DOMAIN,
    text(withdrawal.version),
    hash32(withdrawal.relation_id),
    hash32(withdrawal.profile_commitment),
    u64(withdrawal.epoch),
    hash32(withdrawal.head_a),
    hash32(withdrawal.head_b),
    sideByte(withdrawal.side),
    u64(withdrawal.nonce)
  ]);
}

export const withdrawalHash = withdrawal => sha256(canonicalWithdrawalBytes(withdrawal));

export function withdrawalClosureHash(fields) {
  return sha256(Buffer.concat([
    WITHDRAW_CLOSURE_DOMAIN,
    hash32(fields.prior_state_hash),
    hash32(fields.withdrawal_hash),
    sideByte(fields.side),
    hash32(fields.next_state_hash)
  ]));
}

export function newChain(profileCommitment) {
  if (!HEX32.test(profileCommitment)) throw new TypeError('profile_commitment');
  return {
    protocol_version: PROTOCOL_VERSION,
    kernel_version: VERSION,
    profile_commitment: profileCommitment,
    relations: [],
    closures: [],
    successors: []
  };
}

function validClosureRecord(value) {
  return exactKeys(value, CLOSURE_RECORD_KEYS)
    && HEX32.test(value.closure_hash)
    && HEX32.test(value.relation_id)
    && integer(value.epoch)
    && HEX32.test(value.prior_state_hash)
    && HEX32.test(value.proposal_hash)
    && ['1', '0', '↻', 'WITHDRAW_A', 'WITHDRAW_B'].includes(value.resolution)
    && HEX32.test(value.next_state_hash)
    && (value.handoff_commitment === null || HEX32.test(value.handoff_commitment));
}

function validSuccessorRecord(value) {
  return exactKeys(value, SUCCESSOR_RECORD_KEYS)
    && HEX32.test(value.predecessor_closure_hash)
    && HEX32.test(value.admission_hash)
    && HEX32.test(value.relation_id);
}

function expectedClosureHash(closure) {
  if (closure.resolution === 'WITHDRAW_A' || closure.resolution === 'WITHDRAW_B') {
    return withdrawalClosureHash({
      prior_state_hash: closure.prior_state_hash,
      withdrawal_hash: closure.proposal_hash,
      side: closure.resolution === 'WITHDRAW_A' ? 'A' : 'B',
      next_state_hash: closure.next_state_hash
    });
  }
  return closureHash(closure);
}

function validRelationHistory(relation, closures) {
  if (relation.version !== VERSION
      || samePublicKey(relation.verifier_a, relation.verifier_b)
      || closures.length !== relation.epoch) return false;
  if (closures.length === 0) {
    return relation.status === 'ACTIVE'
      && relation.nonce_a === 0
      && relation.nonce_b === 0;
  }

  const ordered = [...closures].sort((a, b) => a.epoch - b.epoch);
  let expectedNonceA = 0;
  let expectedNonceB = 0;
  for (let index = 0; index < ordered.length; index += 1) {
    const closure = ordered[index];
    if (closure.epoch !== index) return false;
    if (index > 0 && ordered[index - 1].next_state_hash !== closure.prior_state_hash) return false;

    const terminal = closure.resolution !== '1';
    if (index < ordered.length - 1 && terminal) return false;
    if ((closure.resolution === '↻') !== (closure.handoff_commitment !== null)) return false;

    if (closure.resolution === 'WITHDRAW_A') expectedNonceA += 1;
    else if (closure.resolution === 'WITHDRAW_B') expectedNonceB += 1;
    else {
      expectedNonceA += 1;
      expectedNonceB += 1;
    }
  }

  const finalClosure = ordered.at(-1);
  const expectedStatus = finalClosure.resolution === '1' ? 'ACTIVE' : 'ENDED';
  return relation.status === expectedStatus
    && relation.nonce_a === expectedNonceA
    && relation.nonce_b === expectedNonceB
    && finalClosure.next_state_hash === stateHash(relation);
}

function acyclicLineage(edges) {
  const nextByRelation = new Map();
  for (const [from, to] of edges) {
    if (from === to || nextByRelation.has(from)) return false;
    nextByRelation.set(from, to);
  }
  for (const start of nextByRelation.keys()) {
    const visited = new Set();
    let relationID = start;
    while (nextByRelation.has(relationID)) {
      if (visited.has(relationID)) return false;
      visited.add(relationID);
      relationID = nextByRelation.get(relationID);
    }
  }
  return true;
}

export function validChain(chain) {
  if (!exactKeys(chain, CHAIN_KEYS)
      || chain.protocol_version !== PROTOCOL_VERSION
      || chain.kernel_version !== VERSION
      || !HEX32.test(chain.profile_commitment)
      || !Array.isArray(chain.relations)
      || !Array.isArray(chain.closures)
      || !Array.isArray(chain.successors)) return false;
  if (!chain.relations.every(validState)
      || !chain.relations.every(value => value.profile_commitment === chain.profile_commitment)
      || !chain.closures.every(validClosureRecord)
      || !chain.successors.every(validSuccessorRecord)) return false;
  const orderedUnique = (values, key) => values.every((value, index) => index === 0 || values[index - 1][key] < value[key]);
  if (!orderedUnique(chain.relations, 'relation_id')
      || !orderedUnique(chain.closures, 'closure_hash')
      || !orderedUnique(chain.successors, 'predecessor_closure_hash')) return false;

  try {
    const relationByID = new Map(chain.relations.map(value => [value.relation_id, value]));
    const closureByHash = new Map();
    const closuresByRelation = new Map(chain.relations.map(value => [value.relation_id, []]));
    const relationEpochs = new Set();
    for (const closure of chain.closures) {
      if (!relationByID.has(closure.relation_id)
          || closure.closure_hash !== expectedClosureHash(closure)) return false;
      const epochKey = `${closure.relation_id}:${closure.epoch}`;
      if (relationEpochs.has(epochKey)) return false;
      relationEpochs.add(epochKey);
      closureByHash.set(closure.closure_hash, closure);
      closuresByRelation.get(closure.relation_id).push(closure);
    }
    for (const relation of chain.relations) {
      if (!validRelationHistory(relation, closuresByRelation.get(relation.relation_id))) return false;
    }

    const ownedRelations = new Set();
    const lineage = [];
    for (const successor of chain.successors) {
      const predecessor = closureByHash.get(successor.predecessor_closure_hash);
      if (!predecessor
          || predecessor.resolution !== '↻'
          || !relationByID.has(successor.relation_id)
          || ownedRelations.has(successor.relation_id)) return false;
      ownedRelations.add(successor.relation_id);
      lineage.push([predecessor.relation_id, successor.relation_id]);
    }
    return acyclicLineage(lineage);
  } catch {
    return false;
  }
}

function validOpen(open) {
  return exactKeys(open, OPEN_KEYS)
    && typeof open.version === 'string'
    && HEX32.test(open.relation_id)
    && HEX32.test(open.profile_commitment)
    && HEX32.test(open.invocation_id)
    && HEX32.test(open.input_commitment)
    && HEX32.test(open.bounds_commitment)
    && HEX32.test(open.head_a)
    && publicKey(open.verifier_a)
    && HEX32.test(open.head_b)
    && publicKey(open.verifier_b)
    && sortedUniqueHashes(open.handoff_domain)
    && signature(open.signature_a)
    && signature(open.signature_b);
}

function validWithdrawal(withdrawal) {
  return exactKeys(withdrawal, WITHDRAW_KEYS)
    && typeof withdrawal.version === 'string'
    && HEX32.test(withdrawal.relation_id)
    && HEX32.test(withdrawal.profile_commitment)
    && integer(withdrawal.epoch)
    && HEX32.test(withdrawal.head_a)
    && HEX32.test(withdrawal.head_b)
    && typeof withdrawal.side === 'string'
    && integer(withdrawal.nonce)
    && signature(withdrawal.signature);
}

function refusal(code, chain, operationHash = null, relationID = null) {
  return {
    accepted: false,
    code,
    operation_hash: operationHash,
    relation_id: relationID,
    relation_state_hash: null,
    closure_hash: null,
    state: clone(chain)
  };
}

function acceptance(code, chain, operationHash, relation, closure = null) {
  return {
    accepted: true,
    code,
    operation_hash: operationHash,
    relation_id: relation.relation_id,
    relation_state_hash: stateHash(relation),
    closure_hash: closure?.closure_hash ?? null,
    state: chain
  };
}

const relationIndex = (chain, relationID) => chain.relations.findIndex(value => value.relation_id === relationID);
const closureIndex = (chain, closureID) => chain.closures.findIndex(value => value.closure_hash === closureID);
const successorIndex = (chain, closureID) => chain.successors.findIndex(value => value.predecessor_closure_hash === closureID);

function orderedInsert(values, value, key) {
  values.push(value);
  values.sort((a, b) => a[key].localeCompare(b[key]));
}

export function applyOpen(chain, open) {
  if (!validChain(chain)) return refusal(Codes.MALFORMED_CHAIN_STATE, chain);
  if (!validOpen(open)) return refusal(Codes.MALFORMED_OPEN, chain);
  if (open.version !== VERSION) return refusal(Codes.UNSUPPORTED_VERSION, chain);
  if (open.profile_commitment !== chain.profile_commitment) return refusal(Codes.PROFILE_MISMATCH, chain);
  if (relationIndex(chain, open.relation_id) !== -1) return refusal(Codes.RELATION_EXISTS, chain, null, open.relation_id);
  if (samePublicKey(open.verifier_a, open.verifier_b)) return refusal(Codes.VERIFIERS_NOT_DISTINCT, chain, null, open.relation_id);

  const payload = canonicalOpenBytes(open);
  const digest = sha256(payload);
  if (!verifySignature(open.verifier_a, payload, open.signature_a)) {
    return refusal(Codes.OPEN_SIGNATURE_A_INVALID, chain, digest, open.relation_id);
  }
  if (!verifySignature(open.verifier_b, payload, open.signature_b)) {
    return refusal(Codes.OPEN_SIGNATURE_B_INVALID, chain, digest, open.relation_id);
  }

  const relation = {
    version: VERSION,
    relation_id: open.relation_id,
    profile_commitment: open.profile_commitment,
    epoch: 0,
    status: 'ACTIVE',
    head_a: open.head_a,
    verifier_a: open.verifier_a,
    nonce_a: 0,
    head_b: open.head_b,
    verifier_b: open.verifier_b,
    nonce_b: 0,
    handoff_domain: clone(open.handoff_domain)
  };
  const next = clone(chain);
  orderedInsert(next.relations, relation, 'relation_id');
  return acceptance(Codes.OPEN_ACCEPTED, next, digest, relation);
}

function crossingRefusal(code, chain, proposal, digest = null) {
  return refusal(code, chain, digest, proposal?.relation_id ?? null);
}

export function applyCrossing(chain, proposal) {
  if (!validChain(chain)) return refusal(Codes.MALFORMED_CHAIN_STATE, chain);
  if (!validProposalShape(proposal)) return crossingRefusal(Codes.MALFORMED_PROPOSAL, chain, proposal);
  if (proposal.version !== VERSION) return crossingRefusal(Codes.UNSUPPORTED_VERSION, chain, proposal);
  const index = relationIndex(chain, proposal.relation_id);
  if (index === -1) return crossingRefusal(Codes.RELATION_NOT_FOUND, chain, proposal);
  const relation = chain.relations[index];
  if (proposal.profile_commitment !== chain.profile_commitment
      || proposal.profile_commitment !== relation.profile_commitment) return crossingRefusal(Codes.PROFILE_MISMATCH, chain, proposal);
  if (relation.status === 'ENDED') return crossingRefusal(Codes.RELATION_ENDED, chain, proposal);
  if (proposal.epoch !== relation.epoch) return crossingRefusal(Codes.STALE_EPOCH, chain, proposal);
  if (proposal.prev_head_a !== relation.head_a) return crossingRefusal(Codes.STALE_HEAD_A, chain, proposal);
  if (proposal.prev_head_b !== relation.head_b) return crossingRefusal(Codes.STALE_HEAD_B, chain, proposal);
  if (!isNextCounter(relation.nonce_a, proposal.nonce_a)) return crossingRefusal(Codes.NONCE_A_MISMATCH, chain, proposal);
  if (!isNextCounter(relation.nonce_b, proposal.nonce_b)) return crossingRefusal(Codes.NONCE_B_MISMATCH, chain, proposal);
  if (!['1', '0', '↻'].includes(proposal.resolution)) return crossingRefusal(Codes.RESOLUTION_INVALID, chain, proposal);

  let payload;
  let digest;
  try {
    payload = canonicalProposalBytes(proposal);
    digest = proposalHash(proposal);
  } catch {
    return crossingRefusal(Codes.MALFORMED_PROPOSAL, chain, proposal);
  }
  const hasA = proposal.next_head_a !== null;
  const hasB = proposal.next_head_b !== null;
  if (proposal.resolution === '1' ? !(hasA && hasB) : (hasA || hasB)) {
    return crossingRefusal(Codes.OUTPUT_RULE_VIOLATION, chain, proposal, digest);
  }
  if (proposal.resolution === '↻') {
    if (proposal.handoff_manifest === null || !manifestRule(relation.handoff_domain, proposal.handoff_manifest)) {
      return crossingRefusal(Codes.HANDOFF_RULE_VIOLATION, chain, proposal, digest);
    }
  } else if (proposal.handoff_manifest !== null) {
    return crossingRefusal(Codes.HANDOFF_RULE_VIOLATION, chain, proposal, digest);
  }
  if (!verifySignature(relation.verifier_a, payload, proposal.signature_a)) {
    return crossingRefusal(Codes.SIGNATURE_A_INVALID, chain, proposal, digest);
  }
  if (!verifySignature(relation.verifier_b, payload, proposal.signature_b)) {
    return crossingRefusal(Codes.SIGNATURE_B_INVALID, chain, proposal, digest);
  }

  const nextRelation = {
    ...clone(relation),
    epoch: relation.epoch + 1,
    status: proposal.resolution === '1' ? 'ACTIVE' : 'ENDED',
    head_a: proposal.resolution === '1' ? proposal.next_head_a : relation.head_a,
    head_b: proposal.resolution === '1' ? proposal.next_head_b : relation.head_b,
    nonce_a: proposal.nonce_a,
    nonce_b: proposal.nonce_b
  };
  const priorDigest = stateHash(relation);
  const nextDigest = stateHash(nextRelation);
  const handoffCommitment = proposal.resolution === '↻' ? handoffHash(proposal.handoff_manifest) : null;
  const closureFields = {
    prior_state_hash: priorDigest,
    proposal_hash: digest,
    resolution: proposal.resolution,
    next_state_hash: nextDigest,
    handoff_commitment: handoffCommitment
  };
  const closure = {
    closure_hash: closureHash(closureFields),
    relation_id: relation.relation_id,
    epoch: proposal.epoch,
    ...closureFields
  };
  const next = clone(chain);
  next.relations[index] = nextRelation;
  orderedInsert(next.closures, closure, 'closure_hash');
  return acceptance(Codes.ACCEPTED, next, digest, nextRelation, closure);
}

function validSuccessor(value) {
  return exactKeys(value, SUCCESSOR_KEYS)
    && HEX32.test(value.predecessor_closure_hash)
    && validManifestShape(value.handoff_manifest)
    && validFreshAdmission(value.fresh_admission);
}

export function applySuccessor(chain, operation) {
  if (!validChain(chain)) return refusal(Codes.MALFORMED_CHAIN_STATE, chain);
  if (!validSuccessor(operation)) return refusal(Codes.MALFORMED_SUCCESSOR, chain);
  const fresh = operation.fresh_admission;
  if (fresh.version !== VERSION) return refusal(Codes.UNSUPPORTED_VERSION, chain, null, fresh.relation_id);
  const closureAt = closureIndex(chain, operation.predecessor_closure_hash);
  if (closureAt === -1) return refusal(Codes.PREDECESSOR_CLOSURE_NOT_FOUND, chain, null, fresh.relation_id);
  const predecessor = chain.closures[closureAt];
  if (predecessor.resolution !== '↻') return refusal(Codes.PREDECESSOR_NOT_RENEW, chain, null, fresh.relation_id);
  if (successorIndex(chain, predecessor.closure_hash) !== -1) return refusal(Codes.PREDECESSOR_ALREADY_CONSUMED, chain, null, fresh.relation_id);
  if (relationIndex(chain, fresh.relation_id) !== -1) return refusal(Codes.RELATION_EXISTS, chain, null, fresh.relation_id);
  if (fresh.profile_commitment !== chain.profile_commitment) return refusal(Codes.PROFILE_MISMATCH, chain, null, fresh.relation_id);
  if (fresh.predecessor_commitment !== predecessor.closure_hash) return refusal(Codes.PREDECESSOR_COMMITMENT_MISMATCH, chain, null, fresh.relation_id);

  const manifestCommitment = handoffHash(operation.handoff_manifest);
  if (fresh.handoff_commitment !== manifestCommitment
      || predecessor.handoff_commitment !== manifestCommitment) {
    return refusal(Codes.HANDOFF_COMMITMENT_MISMATCH, chain, null, fresh.relation_id);
  }
  const predecessorRelation = chain.relations[relationIndex(chain, predecessor.relation_id)];
  if (!predecessorRelation
      || !manifestRule(predecessorRelation.handoff_domain, operation.handoff_manifest)) {
    return refusal(Codes.HANDOFF_COMMITMENT_MISMATCH, chain, null, fresh.relation_id);
  }
  if (fresh.input_commitment !== operation.handoff_manifest.successor_input_commitment) {
    return refusal(Codes.FRESH_INPUT_MISMATCH, chain, null, fresh.relation_id);
  }
  if (freshSurfaceHash(fresh) !== operation.handoff_manifest.successor_admission_commitment) {
    return refusal(Codes.SUCCESSOR_ADMISSION_COMMITMENT_MISMATCH, chain, null, fresh.relation_id);
  }
  if (samePublicKey(fresh.verifier_a, fresh.verifier_b)) return refusal(Codes.VERIFIERS_NOT_DISTINCT, chain, null, fresh.relation_id);

  const payload = canonicalFreshAdmissionBytes(fresh);
  const digest = freshAdmissionHash(fresh);
  if (!verifySignature(fresh.verifier_a, payload, fresh.signature_a)) {
    return refusal(Codes.FRESH_SIGNATURE_A_INVALID, chain, digest, fresh.relation_id);
  }
  if (!verifySignature(fresh.verifier_b, payload, fresh.signature_b)) {
    return refusal(Codes.FRESH_SIGNATURE_B_INVALID, chain, digest, fresh.relation_id);
  }

  const relation = {
    version: VERSION,
    relation_id: fresh.relation_id,
    profile_commitment: fresh.profile_commitment,
    epoch: 0,
    status: 'ACTIVE',
    head_a: fresh.head_a,
    verifier_a: fresh.verifier_a,
    nonce_a: 0,
    head_b: fresh.head_b,
    verifier_b: fresh.verifier_b,
    nonce_b: 0,
    handoff_domain: successorDomain(operation.handoff_manifest)
  };
  const next = clone(chain);
  orderedInsert(next.relations, relation, 'relation_id');
  orderedInsert(next.successors, {
    predecessor_closure_hash: predecessor.closure_hash,
    admission_hash: digest,
    relation_id: relation.relation_id
  }, 'predecessor_closure_hash');
  return acceptance(Codes.SUCCESSOR_ACCEPTED, next, digest, relation);
}

export function applyWithdrawal(chain, withdrawal) {
  if (!validChain(chain)) return refusal(Codes.MALFORMED_CHAIN_STATE, chain);
  if (!validWithdrawal(withdrawal)) return refusal(Codes.MALFORMED_WITHDRAWAL, chain);
  if (withdrawal.version !== VERSION) return refusal(Codes.UNSUPPORTED_VERSION, chain, null, withdrawal.relation_id);
  const index = relationIndex(chain, withdrawal.relation_id);
  if (index === -1) return refusal(Codes.RELATION_NOT_FOUND, chain, null, withdrawal.relation_id);
  const relation = chain.relations[index];
  if (withdrawal.profile_commitment !== chain.profile_commitment
      || withdrawal.profile_commitment !== relation.profile_commitment) return refusal(Codes.PROFILE_MISMATCH, chain, null, withdrawal.relation_id);
  if (relation.status === 'ENDED') return refusal(Codes.RELATION_ENDED, chain, null, withdrawal.relation_id);
  if (withdrawal.epoch !== relation.epoch) return refusal(Codes.STALE_EPOCH, chain, null, withdrawal.relation_id);
  if (withdrawal.head_a !== relation.head_a) return refusal(Codes.STALE_HEAD_A, chain, null, withdrawal.relation_id);
  if (withdrawal.head_b !== relation.head_b) return refusal(Codes.STALE_HEAD_B, chain, null, withdrawal.relation_id);
  if (!['A', 'B'].includes(withdrawal.side)) return refusal(Codes.SIDE_INVALID, chain, null, withdrawal.relation_id);
  const priorNonce = withdrawal.side === 'A' ? relation.nonce_a : relation.nonce_b;
  if (!isNextCounter(priorNonce, withdrawal.nonce)) {
    return refusal(Codes.WITHDRAW_NONCE_MISMATCH, chain, null, withdrawal.relation_id);
  }

  const payload = canonicalWithdrawalBytes(withdrawal);
  const digest = sha256(payload);
  const verifier = withdrawal.side === 'A' ? relation.verifier_a : relation.verifier_b;
  if (!verifySignature(verifier, payload, withdrawal.signature)) {
    return refusal(Codes.WITHDRAW_SIGNATURE_INVALID, chain, digest, withdrawal.relation_id);
  }

  const nextRelation = {
    ...clone(relation),
    epoch: relation.epoch + 1,
    status: 'ENDED',
    nonce_a: withdrawal.side === 'A' ? withdrawal.nonce : relation.nonce_a,
    nonce_b: withdrawal.side === 'B' ? withdrawal.nonce : relation.nonce_b
  };
  const priorDigest = stateHash(relation);
  const nextDigest = stateHash(nextRelation);
  const resolution = withdrawal.side === 'A' ? 'WITHDRAW_A' : 'WITHDRAW_B';
  const closureDigest = withdrawalClosureHash({
    prior_state_hash: priorDigest,
    withdrawal_hash: digest,
    side: withdrawal.side,
    next_state_hash: nextDigest
  });
  const closure = {
    closure_hash: closureDigest,
    relation_id: relation.relation_id,
    epoch: withdrawal.epoch,
    prior_state_hash: priorDigest,
    proposal_hash: digest,
    resolution,
    next_state_hash: nextDigest,
    handoff_commitment: null
  };
  const next = clone(chain);
  next.relations[index] = nextRelation;
  orderedInsert(next.closures, closure, 'closure_hash');
  return acceptance(Codes.WITHDRAW_ACCEPTED, next, digest, nextRelation, closure);
}

export function applyOperation(chain, operation) {
  if (!operation || typeof operation !== 'object' || Array.isArray(operation)) {
    return refusal(Codes.MALFORMED_CHAIN_STATE, chain);
  }
  switch (operation.kind) {
    case 'OPEN_RELATION': return applyOpen(chain, operation.value);
    case 'APPLY_CROSSING': return applyCrossing(chain, operation.value);
    case 'ADMIT_SUCCESSOR': return applySuccessor(chain, operation.value);
    case 'WITHDRAW_CHAMBER': return applyWithdrawal(chain, operation.value);
    default: return refusal(Codes.MALFORMED_CHAIN_STATE, chain);
  }
}

export function canonicalOperationBytes(operation) {
  switch (operation.kind) {
    case 'OPEN_RELATION': return canonicalOpenBytes(operation.value);
    case 'APPLY_CROSSING': return canonicalProposalBytes(operation.value);
    case 'ADMIT_SUCCESSOR': return canonicalFreshAdmissionBytes(operation.value.fresh_admission);
    case 'WITHDRAW_CHAMBER': return canonicalWithdrawalBytes(operation.value);
    default: throw new TypeError('operation kind');
  }
}
