import assert from 'node:assert/strict';
import { createPublicKey } from 'node:crypto';
import { readFile } from 'node:fs/promises';
import test from 'node:test';
import {
  Codes,
  PRECEDENCE,
  applyOperation,
  canonicalOperationBytes,
  parseStrictJson,
  stable,
  validChain
} from '../src/continuity.mjs';
import {
  closureHash,
  handoffHash,
  isNextCounter,
  publicKey,
  sha256,
  stateHash,
  u64
} from '../src/ik002.mjs';

const suite = JSON.parse(await readFile(new URL('../vectors/continuity_vectors_001.json', import.meta.url)));

test('suite hash seals every executable vector', () => {
  const { suite_hash: expected, ...body } = suite;
  assert.equal(sha256(Buffer.from(stable(body))), expected);
});

for (const vector of suite.operations) {
  test(`${vector.id}: ${vector.name}`, () => {
    assert.equal(validChain(vector.initial_state), true);
    const before = stable(vector.initial_state);
    const result = applyOperation(vector.initial_state, vector.operation);
    assert.deepEqual(result, vector.expected);
    if (vector.canonical_unsigned_hex !== null) {
      assert.equal(canonicalOperationBytes(vector.operation).toString('hex'), vector.canonical_unsigned_hex);
    }
    if (!result.accepted) assert.equal(stable(result.state), before);
    assert.equal(validChain(result.state), true);
  });
}

test('full lifecycle sequence is deterministic and terminal', () => {
  let state = suite.lifecycle_sequence.initial_state;
  for (const step of suite.lifecycle_sequence.steps) {
    assert.equal(canonicalOperationBytes(step.operation).toString('hex'), step.canonical_unsigned_hex);
    const first = applyOperation(state, step.operation);
    const second = applyOperation(state, step.operation);
    assert.deepEqual(first, step.expected);
    assert.deepEqual(second, first);
    assert.equal(first.accepted, true);
    state = first.state;
  }
  assert.deepEqual(state, suite.lifecycle_sequence.final_state);
  assert.equal(state.relations.length, 2);
  assert.equal(state.relations.every(value => value.status === 'ENDED'), true);
  assert.equal(state.successors.length, 1);
});

test('refusal result owns a deep state copy', () => {
  const source = suite.operations.find(value => value.id === 'S002');
  const before = stable(source.initial_state);
  const result = applyOperation(source.initial_state, source.operation);
  assert.equal(result.accepted, false);
  assert.notEqual(result.state, source.initial_state);

  const changedDomain = sha256(Buffer.from('changed-domain'));
  const changedCommitment = sha256(Buffer.from('changed-handoff-commitment'));
  const relation = result.state.relations.find(value => value.handoff_domain.length > 0);
  const closure = result.state.closures.find(value => value.handoff_commitment !== null);
  assert.ok(relation);
  assert.ok(closure);
  relation.handoff_domain[0] = changedDomain;
  closure.handoff_commitment = changedCommitment;

  assert.equal(stable(source.initial_state), before);
  assert.notEqual(stable(result.state), before);
});

test('declared precedence is deeply frozen', () => {
  assert.equal(Object.isFrozen(PRECEDENCE), true);
  for (const values of Object.values(PRECEDENCE)) assert.equal(Object.isFrozen(values), true);
  assert.throws(() => { PRECEDENCE.crossing[0] = Codes.ACCEPTED; }, TypeError);
  assert.equal(PRECEDENCE.crossing[0], Codes.MALFORMED_CHAIN_STATE);
});

test('persisted closure hashes and relation references are enforced', () => {
  const renewed = structuredClone(suite.operations.find(value => value.id === 'C005').expected.state);
  const hashCorrupt = structuredClone(renewed);
  hashCorrupt.closures[0].proposal_hash = sha256(Buffer.from('hash-corrupt'));
  assert.equal(validChain(hashCorrupt), false);

  const dangling = structuredClone(renewed);
  dangling.closures[0].relation_id = sha256(Buffer.from('missing-relation'));
  assert.equal(validChain(dangling), false);
});

test('persisted closure epochs form one contiguous state-hash chain', () => {
  const renewed = structuredClone(suite.operations.find(value => value.id === 'C005').expected.state);
  const ordered = [...renewed.closures].sort((a, b) => a.epoch - b.epoch);
  const later = ordered[1];
  later.prior_state_hash = sha256(Buffer.from('detached-prior-state'));
  later.closure_hash = closureHash(later);
  renewed.closures.sort((a, b) => a.closure_hash.localeCompare(b.closure_hash));
  assert.equal(validChain(renewed), false);
});

test('a successor relation has one accepted renew-lineage owner', () => {
  const succeeded = structuredClone(suite.operations.find(value => value.id === 'S001').expected.state);
  const target = succeeded.relations.find(value => value.relation_id === succeeded.successors[0].relation_id);
  const predecessor = {
    ...structuredClone(target),
    relation_id: sha256(Buffer.from('second-predecessor-relation'))
  };
  const ended = {
    ...structuredClone(predecessor),
    epoch: 1,
    status: 'ENDED',
    nonce_a: 1,
    nonce_b: 1
  };
  const fabricatedRenew = {
    relation_id: ended.relation_id,
    epoch: 0,
    prior_state_hash: stateHash(predecessor),
    proposal_hash: sha256(Buffer.from('second-renew-proposal')),
    resolution: '↻',
    next_state_hash: stateHash(ended),
    handoff_commitment: sha256(Buffer.from('second-handoff'))
  };
  fabricatedRenew.closure_hash = closureHash(fabricatedRenew);
  succeeded.relations.push(ended);
  succeeded.relations.sort((a, b) => a.relation_id.localeCompare(b.relation_id));
  succeeded.closures.push(fabricatedRenew);
  succeeded.closures.sort((a, b) => a.closure_hash.localeCompare(b.closure_hash));
  succeeded.successors.push({
    predecessor_closure_hash: fabricatedRenew.closure_hash,
    admission_hash: sha256(Buffer.from('second-successor-admission')),
    relation_id: target.relation_id
  });
  succeeded.successors.sort((a, b) => a.predecessor_closure_hash.localeCompare(b.predecessor_closure_hash));
  assert.equal(validChain(succeeded), false);
});

test('Ed25519 SPKI rejects trailing DER accepted by the platform parser', () => {
  const source = suite.operations.find(value => value.id === 'O001');
  const operation = structuredClone(source.operation);
  const canonical = Buffer.from(operation.value.verifier_a, 'base64');
  const trailing = Buffer.concat([canonical, Buffer.from([0])]);
  assert.equal(createPublicKey({ key: trailing, format: 'der', type: 'spki' }).asymmetricKeyType, 'ed25519');
  assert.equal(publicKey(trailing.toString('base64')), false);
  operation.value.verifier_a = trailing.toString('base64');
  assert.equal(applyOperation(source.initial_state, operation).code, Codes.MALFORMED_OPEN);
});

test('non-minimal SPKI encoding of the same Ed25519 material is rejected', () => {
  const source = suite.operations.find(value => value.id === 'O001');
  const operation = structuredClone(source.operation);
  const canonical = Buffer.from(operation.value.verifier_a, 'base64');
  assert.equal(canonical[0], 0x30);
  assert.equal(canonical[1], canonical.length - 2);
  const alternate = Buffer.concat([
    Buffer.from([0x30, 0x81, canonical[1]]),
    canonical.subarray(2)
  ]);
  const reparsed = createPublicKey({ key: alternate, format: 'der', type: 'spki' })
    .export({ format: 'der', type: 'spki' });
  assert.equal(reparsed.equals(canonical), true);
  assert.equal(alternate.equals(canonical), false);
  assert.equal(publicKey(alternate.toString('base64')), false);
  operation.value.verifier_b = alternate.toString('base64');
  assert.equal(applyOperation(source.initial_state, operation).code, Codes.MALFORMED_OPEN);
});

test('nested unknown and duplicate JSON keys are refused', () => {
  const source = suite.operations.find(value => value.id === 'C005');
  const unknown = structuredClone(source.operation);
  unknown.value.handoff_manifest.transformed[0].unknown = sha256(Buffer.from('unknown'));
  assert.equal(applyOperation(source.initial_state, unknown).code, Codes.MALFORMED_PROPOSAL);
  assert.throws(
    () => parseStrictJson('{"kind":"APPLY_CROSSING","value":{"resolution":"1","resolution":"0"}}'),
    /duplicate key "resolution"/
  );
});

test('nullable fields distinguish explicit null from omission', () => {
  const source = suite.operations.find(value => value.id === 'C005');
  assert.equal(source.operation.value.next_head_a, null);
  assert.equal(applyOperation(source.initial_state, source.operation).accepted, true);
  const omitted = structuredClone(source.operation);
  delete omitted.value.next_head_a;
  assert.equal(applyOperation(source.initial_state, omitted).code, Codes.MALFORMED_PROPOSAL);
});

test('successor revalidates the committed manifest against predecessor domain', () => {
  const source = suite.operations.find(value => value.id === 'S001');
  const chain = structuredClone(source.initial_state);
  const operation = structuredClone(source.operation);
  operation.value.handoff_manifest.preserved = [];
  const commitment = handoffHash(operation.value.handoff_manifest);
  const closure = chain.closures.find(value => value.resolution === '↻');
  closure.handoff_commitment = commitment;
  closure.closure_hash = closureHash(closure);
  chain.closures.sort((a, b) => a.closure_hash.localeCompare(b.closure_hash));
  operation.value.predecessor_closure_hash = closure.closure_hash;
  operation.value.fresh_admission.predecessor_commitment = closure.closure_hash;
  operation.value.fresh_admission.handoff_commitment = commitment;
  assert.equal(validChain(chain), true);
  const before = stable(chain);
  const result = applyOperation(chain, operation);
  assert.equal(result.code, Codes.HANDOFF_COMMITMENT_MISMATCH);
  assert.equal(stable(result.state), before);
});

test('counter domain accepts MAX_SAFE-1 to MAX_SAFE and refuses exhaustion', () => {
  const maximum = Number.MAX_SAFE_INTEGER;
  assert.equal(isNextCounter(maximum - 1, maximum), true);
  assert.equal(isNextCounter(maximum, maximum), false);
  assert.equal(isNextCounter(maximum, maximum + 1), false);
  assert.equal(u64(maximum - 1).toString('hex'), '001ffffffffffffe');
  assert.equal(u64(maximum).toString('hex'), '001fffffffffffff');

  const source = suite.operations.find(value => value.id === 'C001');
  const operation = structuredClone(source.operation);
  operation.value.nonce_a = maximum + 1;
  assert.equal(applyOperation(source.initial_state, operation).code, Codes.MALFORMED_PROPOSAL);
});
