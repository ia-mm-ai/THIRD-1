# HANDOFF_PROTOCOL_001 — the GitHub seam between two assistants

> Status: adopted by Marko Markota on 2026-10-21. The seam is `ia-mm-ai/THIRD-1`, this repository. It applies to any assistant working on this lineage, in any sandbox, on any machine.

## 0. Lineage position

```text
ia-mm-ai/THIRD-1      THE SEAM. Reality Model 00–06, [integrity_kernel], derived runtime forms, VM_IK.
ia-mm-ai/PRESENCE     ARCHIVED PREDECESSOR. VM005, the Avalanche L1 WJcV97GN…, its founding run. Read-only. Never reopened.
```

PRESENCE is referenced, never continued. Its closing evidence is carried as one bound artifact:

| Artifact | SHA-256 | Bytes | Contains |
|---|---|---|---|
| `claude-founding-run-001.bundle` (git bundle, ref `claude/founding-run-001` at `471da1c7e101d0f1a1cb976be87f1a7448673959`) | `ea4ea5b3cd33cc307fcc4d78129b20a9a6aa72c9caa354515829d7d2d8b850ca` | 48,656,260 | the PRESENCE_L1_RELEASE_CANDIDATE history through the VM005 founding run closed STAND on 2026-10-21 |

Nothing in that bundle is current law here. It proves a predecessor completed; it confers no identity, authority, continuity, or value on anything in THIRD-1.

## 1. Positions

```text
Chamber A   ChatGPT (cloud sandbox)   drafts, derives, seals candidate releases
Chamber B   Claude (Marko's machine)  verifies against disk, runs, embodies on node and chain
Seam        Marko                     the only party who admits a crossing
```

Neither chamber can address the other. Nothing an assistant says about the other's work is evidence; only files are. The seam has exactly one surface: a branch of this repository on GitHub.

## 2. Branch naming

```text
chatgpt/<topic>-<NNN>     written only by Chamber A
claude/<topic>-<NNN>      written only by Chamber B
main                      written only by Marko, by merge
```

One writer per branch, always. A chamber never commits to the other's branch and never force-pushes. A branch that has crossed is frozen at the commit that crossed; corrections go on a new numbered branch.

## 3. Hash binding against files on disk

A release, manifest, or document on a branch may bind other files only by `path + byte_length + sha256`, and every bound path MUST be present in the same branch at the same commit. A hash that does not resolve to a file in the branch is a phantom and the whole release is refused, whatever its status field says.

The root bindings of this lineage are the exact human source forms in `00_SUBSTRATE/` through `06_EMERGENCE/` and `[integrity_kernel]/`, as they exist on `main`. A derived form that binds a different byte sequence for any of them is not a successor; it is a different document.

Before Chamber B acts on anything from a `chatgpt/*` branch it:

1. fetches the exact commit, never a summary or screenshot;
2. recomputes every declared `sha256` and `byte_length` against the checked-out files;
3. recomputes the upstream bindings against the root forms on `main`;
4. records the result as a `VERIFICATION_<NNN>.json` on a `claude/*` branch, with the verified commit hash;
5. stops on the first mismatch and reports it to the seam.

Chamber A does the symmetric check on anything it consumes from a `claude/*` branch.

Precedence is unchanged: exact source bytes on disk > machine projection > derived form > implementation > execution record. A newer file is not a successor; recency, filename reuse and sandbox state carry nothing.

## 4. What each chamber may do

**Chamber A (ChatGPT) may:** read any branch; write `chatgpt/*` branches; derive forms, runtime code, vectors and sealed candidate releases; run its own tests in its own sandbox; propose irreversible steps as parameterized inputs with every value explicit.

**Chamber A may not:** reach Marko's disk, any validator, SSH keys, Ledger artifacts, the P-Chain, or any chain API; claim a step is executed, deployed, funded or accepted; bind a file that is not in the branch; write to `claude/*` or `main`; declare its own candidate adopted or current; continue anything from PRESENCE by implication.

**Chamber B (Claude) may:** read any branch; write `claude/*` branches; verify bindings; run read-only checks on infrastructure; build and run locally; submit mainnet operations ONLY after a `VERIFICATION_<NNN>.json` PASS, a LOCALITY_INTERLOCK `external-write` ALLOWED, and Marko's explicit per-step authorization in the current turn; record every exchange as evidence files.

**Chamber B may not:** edit a `chatgpt/*` branch; "fix" a Chamber A artifact in place (it opens a `claude/*` branch and names the defect); submit anything whose bytes were not verified in the same session; push, merge or delete without Marko's word; enter credentials; treat a Chamber A seal as authority.

**The seam (Marko) alone:** merges to `main`; authorizes each irreversible step; decides what crosses from a branch to a node or chain; decides what crosses from a node back to a branch; pushes, since neither chamber holds his credentials. Silence is not admission. A screenshot is not admission. "Proceed" from an earlier turn is not admission.

## 5. Crossing record

Every crossing of an artifact into a node or chain produces, on a `claude/*` branch, in a folder named for the stage:

```text
<STAGE>_<NNN>/
  VERIFICATION_<NNN>.json      what was verified, against which commit, result
  <operation bytes>            the exact bytes submitted
  <raw exchange>.txt           request, response, receipt, record, state readback
  <STAGE>_FINAL_STATE_<NNN>.json
```

Chamber A consumes those files to continue. It does not consume Chamber B's prose.

## 6. Non-collapses

```text
sealed                 ≠ verified
verified               ≠ authorized
authorized             ≠ executed
executed               ≠ accepted
accepted               ≠ adopted or current
sandbox file           ≠ file on disk
summary                ≠ bytes
screenshot             ≠ evidence
predecessor            ≠ current
one assistant's claim  ≠ the other's precondition
```

## 7. Closure

A chamber proposes; the seam admits; a chamber embodies; the record returns. No side becomes the other, no side becomes the seam, and nothing crosses on the strength of having been said.
