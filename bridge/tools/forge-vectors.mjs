import assert from 'node:assert/strict';
import { createPrivateKey, createPublicKey, sign } from 'node:crypto';
import { mkdir, writeFile } from 'node:fs/promises';
import {
  PRECEDENCE,
  applyOperation,
  canonicalOpenBytes,
  canonicalOperationBytes,
  canonicalWithdrawalBytes,
  newChain,
  openHash,
  PROTOCOL_VERSION,
  stable,
  VERSION,
  withdrawalHash
} from '../src/continuity.mjs';
import {
  canonicalFreshAdmissionBytes,
  canonicalProposalBytes,
  freshAdmissionHash,
  freshSurfaceHash,
  handoffHash,
  sha256
} from '../src/ik002.mjs';

const H = label => sha256(Buffer.from(label));
const clone = value => structuredClone(value);

function seedKey(label) {
  const seed = Buffer.from(H(label), 'hex');
  const prefix = Buffer.from('302e020100300506032b657004220420', 'hex');
  const privateKey = createPrivateKey({
    key: Buffer.concat([prefix, seed]),
    format: 'der',
    type: 'pkcs8'
  });
  const publicKey = createPublicKey(privateKey);
  return {
    privateKey,
    publicDer: publicKey.export({ format: 'der', type: 'spki' }).toString('base64')
  };
}

const A = seedKey('CONTINUITY-001-A');
const B = seedKey('CONTINUITY-001-B');
const C = seedKey('CONTINUITY-001-C');
const D = seedKey('CONTINUITY-001-D');
const PROFILE = H('CONTINUITY-BASE-PROFILE-001');
const DOMAIN = [H('carry-alpha'), H('carry-beta'), H('carry-gamma')].sort();

function signOpen(value, keyA = A.privateKey, keyB = B.privateKey) {
  const open = { ...clone(value), signature_a: '', signature_b: '' };
  const payload = canonicalOpenBytes(open);
  open.signature_a = sign(null, payload, keyA).toString('base64');
  open.signature_b = sign(null, payload, keyB).toString('base64');
  return open;
}

function originOpen() {
  return signOpen({
    version: VERSION,
    relation_id: H('continuity-origin-relation'),
    profile_commitment: PROFILE,
    invocation_id: H('continuity-origin-invocation'),
    input_commitment: H('continuity-origin-input'),
    bounds_commitment: H('continuity-origin-bounds'),
    head_a: H('continuity-origin-head-a'),
    verifier_a: A.publicDer,
    head_b: H('continuity-origin-head-b'),
    verifier_b: B.publicDer,
    handoff_domain: DOMAIN,
    signature_a: '',
    signature_b: ''
  });
}

function signProposal(value, keyA = A.privateKey, keyB = B.privateKey) {
  const proposal = { ...clone(value), signature_a: '', signature_b: '' };
  const payload = canonicalProposalBytes(proposal);
  proposal.signature_a = sign(null, payload, keyA).toString('base64');
  proposal.signature_b = sign(null, payload, keyB).toString('base64');
  return proposal;
}

function crossing(relation, resolution, manifest = null, nextSuffix = '1') {
  return signProposal({
    version: VERSION,
    relation_id: relation.relation_id,
    profile_commitment: relation.profile_commitment,
    epoch: relation.epoch,
    prev_head_a: relation.head_a,
    prev_head_b: relation.head_b,
    resolution,
    next_head_a: resolution === '1' ? H(`continuity-head-a-${nextSuffix}`) : null,
    next_head_b: resolution === '1' ? H(`continuity-head-b-${nextSuffix}`) : null,
    nonce_a: relation.nonce_a + 1,
    nonce_b: relation.nonce_b + 1,
    handoff_manifest: resolution === '↻' ? manifest : null,
    signature_a: '',
    signature_b: ''
  });
}

function successorSurface() {
  return {
    version: VERSION,
    relation_id: H('continuity-successor-relation'),
    profile_commitment: PROFILE,
    invocation_id: H('continuity-successor-invocation'),
    input_commitment: H('continuity-successor-input'),
    bounds_commitment: H('continuity-successor-bounds'),
    head_a: H('continuity-successor-head-a'),
    verifier_a: C.publicDer,
    head_b: H('continuity-successor-head-b'),
    verifier_b: D.publicDer
  };
}

function successorManifest(surface = successorSurface()) {
  return {
    successor_input_commitment: surface.input_commitment,
    successor_admission_commitment: freshSurfaceHash(surface),
    preserved: [DOMAIN[0]],
    transformed: [{ from: DOMAIN[1], to: H('carry-beta-prime') }],
    ended: [DOMAIN[2]],
    refused: [],
    unresolved: []
  };
}

function signFresh(value, keyA = C.privateKey, keyB = D.privateKey) {
  const fresh = { ...clone(value), signature_a: '', signature_b: '' };
  const payload = canonicalFreshAdmissionBytes(fresh);
  fresh.signature_a = sign(null, payload, keyA).toString('base64');
  fresh.signature_b = sign(null, payload, keyB).toString('base64');
  return fresh;
}

function successorOperation(renewResult, manifest, surface = successorSurface()) {
  const fresh = signFresh({
    ...surface,
    predecessor_commitment: renewResult.closure_hash,
    handoff_commitment: handoffHash(manifest),
    signature_a: '',
    signature_b: ''
  });
  return {
    predecessor_closure_hash: renewResult.closure_hash,
    handoff_manifest: clone(manifest),
    fresh_admission: fresh
  };
}

function signWithdrawal(value, privateKey) {
  const withdrawal = { ...clone(value), signature: '' };
  withdrawal.signature = sign(null, canonicalWithdrawalBytes(withdrawal), privateKey).toString('base64');
  return withdrawal;
}

function withdrawal(relation, side) {
  return signWithdrawal({
    version: VERSION,
    relation_id: relation.relation_id,
    profile_commitment: relation.profile_commitment,
    epoch: relation.epoch,
    head_a: relation.head_a,
    head_b: relation.head_b,
    side,
    nonce: side === 'A' ? relation.nonce_a + 1 : relation.nonce_b + 1,
    signature: ''
  }, side === 'A' ? C.privateKey : D.privateKey);
}

function operation(kind, value) {
  return { kind, value };
}

const vectors = [];

function vector(id, name, initialState, op) {
  const before = stable(initialState);
  const expected = applyOperation(initialState, op);
  if (!expected.accepted) assert.equal(stable(expected.state), before, `${id}: refusal mutated state`);
  let canonicalUnsignedHex = null;
  try { canonicalUnsignedHex = canonicalOperationBytes(op).toString('hex'); } catch {}
  vectors.push({
    id,
    name,
    initial_state: initialState,
    operation: op,
    canonical_unsigned_hex: canonicalUnsignedHex,
    expected
  });
  return expected;
}

const empty = newChain(PROFILE);
const open = originOpen();
const opened = vector('O001', 'fresh relation opens without predecessor', empty, operation('OPEN_RELATION', open));
assert.equal(opened.accepted, true);
assert.equal(opened.operation_hash, openHash(open));

vector('O002', 'relation id cannot be reused', opened.state, operation('OPEN_RELATION', open));

{
  const wrongProfile = signOpen({ ...originOpen(), relation_id: H('wrong-profile-relation'), profile_commitment: H('wrong-profile') });
  vector('O003', 'open must use genesis-pinned profile', empty, operation('OPEN_RELATION', wrongProfile));
}

{
  const sameKey = signOpen({
    ...originOpen(), relation_id: H('same-verifier-relation'),
    verifier_b: A.publicDer, signature_a: '', signature_b: ''
  }, A.privateKey, A.privateKey);
  vector('O004', 'two chamber slots require distinct technical verifiers', empty, operation('OPEN_RELATION', sameKey));
}

{
  const bad = clone(open);
  bad.signature_a = Buffer.alloc(64, 7).toString('base64');
  vector('O005', 'open rejects invalid A signature', empty, operation('OPEN_RELATION', bad));
}

const relation0 = opened.state.relations.find(value => value.relation_id === open.relation_id);
const completeProposal = crossing(relation0, '1', null, 'complete-1');
const completed = vector('C001', 'complete changes both heads atomically', opened.state, operation('APPLY_CROSSING', completeProposal));
assert.equal(completed.accepted, true);

{
  const partial = {
    ...crossing(relation0, '1', null, 'partial'),
    next_head_b: null,
    signature_a: '',
    signature_b: ''
  };
  const signed = signProposal(partial);
  vector('C002', 'partial successor head set is refused', opened.state, operation('APPLY_CROSSING', signed));
}

vector('C003', 'accepted prior state loses a concurrent crossing race', completed.state, operation('APPLY_CROSSING', crossing(relation0, '1', null, 'race')));

{
  const mismatched = crossing(relation0, '1', null, 'mismatch');
  const other = crossing(relation0, '1', null, 'other');
  mismatched.signature_b = other.signature_b;
  vector('C004', 'both signatures must bind identical bytes', opened.state, operation('APPLY_CROSSING', mismatched));
}

const surface = successorSurface();
const manifest = successorManifest(surface);
const relation1 = completed.state.relations.find(value => value.relation_id === open.relation_id);
const renewProposal = crossing(relation1, '↻', manifest);
const renewed = vector('C005', 'renew closes predecessor and stores handoff closure', completed.state, operation('APPLY_CROSSING', renewProposal));
assert.equal(renewed.accepted, true);
assert.equal(renewed.state.relations.find(value => value.relation_id === open.relation_id).status, 'ENDED');

vector('C006', 'renewed predecessor cannot move again', renewed.state, operation('APPLY_CROSSING', crossing(relation1, '1', null, 'after-renew')));

const admit = successorOperation(renewed, manifest, surface);
const succeeded = vector('S001', 'stored renew closure admits one distinct successor', renewed.state, operation('ADMIT_SUCCESSOR', admit));
assert.equal(succeeded.accepted, true);
assert.equal(succeeded.operation_hash, freshAdmissionHash(admit.fresh_admission));

vector('S002', 'renew closure is consumed exactly once', succeeded.state, operation('ADMIT_SUCCESSOR', admit));

{
  const phantom = clone(admit);
  phantom.predecessor_closure_hash = H('phantom-renew-closure');
  phantom.fresh_admission.predecessor_commitment = phantom.predecessor_closure_hash;
  phantom.fresh_admission = signFresh(phantom.fresh_admission);
  vector('S003', 'caller-supplied phantom closure is refused', renewed.state, operation('ADMIT_SUCCESSOR', phantom));
}

{
  const mismatched = clone(admit);
  mismatched.handoff_manifest.preserved = [];
  vector('S004', 'successor must bind exact stored handoff commitment', renewed.state, operation('ADMIT_SUCCESSOR', mismatched));
}

const successorRelation = succeeded.state.relations.find(value => value.relation_id === surface.relation_id);
const withdrawB = withdrawal(successorRelation, 'B');
const withdrawn = vector('W001', 'B unilaterally ends future joint effect', succeeded.state, operation('WITHDRAW_CHAMBER', withdrawB));
assert.equal(withdrawn.accepted, true);
const withdrawnRelation = withdrawn.state.relations.find(value => value.relation_id === surface.relation_id);
assert.equal(withdrawnRelation.head_a, successorRelation.head_a);
assert.equal(withdrawnRelation.head_b, successorRelation.head_b);
assert.equal(withdrawnRelation.nonce_a, successorRelation.nonce_a);
assert.equal(withdrawnRelation.nonce_b, successorRelation.nonce_b + 1);

{
  const bad = withdrawal(successorRelation, 'A');
  bad.signature = Buffer.alloc(64, 9).toString('base64');
  vector('W002', 'withdrawal requires only the selected chamber signature', succeeded.state, operation('WITHDRAW_CHAMBER', bad));
}

{
  const stale = withdrawal(successorRelation, 'A');
  stale.nonce = successorRelation.nonce_a;
  stale.signature = sign(null, canonicalWithdrawalBytes(stale), C.privateKey).toString('base64');
  vector('W003', 'withdrawal nonce is replay-safe', succeeded.state, operation('WITHDRAW_CHAMBER', stale));
}

vector('W004', 'ended relation refuses a second withdrawal', withdrawn.state, operation('WITHDRAW_CHAMBER', withdrawB));

const sequenceOperations = [
  operation('OPEN_RELATION', open),
  operation('APPLY_CROSSING', completeProposal),
  operation('APPLY_CROSSING', renewProposal),
  operation('ADMIT_SUCCESSOR', admit),
  operation('WITHDRAW_CHAMBER', withdrawB)
];
let sequenceState = newChain(PROFILE);
const sequenceSteps = [];
for (const op of sequenceOperations) {
  const result = applyOperation(sequenceState, op);
  assert.equal(result.accepted, true);
  sequenceSteps.push({
    operation: op,
    canonical_unsigned_hex: canonicalOperationBytes(op).toString('hex'),
    expected: result
  });
  sequenceState = result.state;
}

const body = {
  suite_id: 'CONTINUITY_L1_VECTORS_001',
  protocol_version: PROTOCOL_VERSION,
  kernel_version: VERSION,
  profile_commitment: PROFILE,
  domains_hex: {
    open_relation: Buffer.from('IK:OPEN_RELATION:001\0').toString('hex'),
    crossing: Buffer.from('IK:CROSSING:002\0').toString('hex'),
    relation_state: Buffer.from('IK:RELATION_STATE:002\0').toString('hex'),
    handoff: Buffer.from('IK:HANDOFF:002\0').toString('hex'),
    closure: Buffer.from('IK:CLOSURE:002\0').toString('hex'),
    fresh_admission: Buffer.from('IK:FRESH_ADMISSION:002\0').toString('hex'),
    fresh_surface: Buffer.from('IK:FRESH_SURFACE:002\0').toString('hex'),
    withdraw_chamber: Buffer.from('IK:WITHDRAW_CHAMBER:001\0').toString('hex'),
    withdraw_closure: Buffer.from('IK:WITHDRAW_CLOSURE:001\0').toString('hex')
  },
  precedence: PRECEDENCE,
  operations: vectors,
  lifecycle_sequence: {
    id: 'L001',
    name: 'fresh open, complete, terminal renew, one successor, unilateral withdrawal',
    initial_state: newChain(PROFILE),
    steps: sequenceSteps,
    final_state: sequenceState
  }
};

const suite = { ...body, suite_hash: sha256(Buffer.from(stable(body))) };
const target = new URL('../vectors/continuity_vectors_001.json', import.meta.url);
await mkdir(new URL('../vectors/', import.meta.url), { recursive: true });
await writeFile(target, `${JSON.stringify(suite, null, 2)}\n`);

assert.equal(withdrawalHash(withdrawB), withdrawn.operation_hash);
console.log(JSON.stringify({
  suite_id: suite.suite_id,
  suite_hash: suite.suite_hash,
  vectors: suite.operations.length,
  lifecycle_steps: suite.lifecycle_sequence.steps.length
}));
