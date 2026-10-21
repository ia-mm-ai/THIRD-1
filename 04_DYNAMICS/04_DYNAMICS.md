# DYNAMICS

> **Dynamic actuality invariant**
>
> Every claimed actual change must identify what changes or is maintained, the relevant prior and subsequent conditions, the Occurrence or Process through which change obtains, the Domain and ordered Interval, the structural and realization basis, the applicable criterion, and the exact result or unresolved position. A rule, Profile, State description, trace, or Record never performs the change it represents.

> **Open-dynamics invariant**
>
> Not every mode of change is already understood, discretized, computational, or nameable. Every encountered change may remain dynamically located without being denied, forced into a state machine, converted into Agency or purpose, or represented as failure merely because the present grammar cannot yet resolve its mechanism.

> **No-erasure invariant**
>
> An actual Occurrence remains part of the past even where selected effects are reversed, compensated, rolled back, corrected, hidden, forgotten, or no longer recoverable. A later State may restore relevant prior conditions; it cannot make the intervening Occurrence never have happened.

## Status and Dependency Declaration

| Field | Declaration |
|---|---|
| Form role | Human Source Form |
| Layer | Dynamics |
| Source-form status | `FINAL` |
| Semantic source | This exact human-readable Form |
| Machine projection | Permitted only as a derivative Representation |

This Form imports the following exact upstream Forms:

| Layer | Canonical Reality Model path | SHA-256 | Bytes | Imported function |
|---|---|---|---:|---|
| Substrate | `00_SUBSTRATE/00_SUBSTRATE.md` | `6c3ea987c07840c9f019e6f9bee6e54e9eef6c4465359036c0ab2b455a431b99` | 9,301 | Physical ground |
| Ontology | `01_ONTOLOGY/01_ONTOLOGY.md` | `25bdaf01bab4db5fea2e88d7efb207428ec0ee64797d9ebb8fb7b0f659819efd` | 168,760 | Declarative semantic type system |
| Constitution | `02_CONSTITUTION/02_CONSTITUTION.md` | `6a2c369e7570d5dd717b98cbb8216fed8d3c225eeeb166661d362b3120608f12` | 97,393 | Constitutional constraints and normative boundary |
| Structure | `03_STRUCTURE/03_STRUCTURE.md` | `ce3314ff3c14a5f7d130f9586b66a94442023a25dae2c4f0e3b92a712101b3bd` | 144,167 | Arrangement grammar and structural handoff |

Importing the Constitution does not adopt it, make it Current, make it applicable, bind any bearer, or create any constitutional Effect. Structure and Dynamics remain semantic descriptions of actuality whether or not a constitutional source governs a particular matter.

The paths identify layers relative to the root of the exact Reality Model stack. A changed upstream byte sequence is a different imported Form. Recency, filename reuse, publication, parsing, or implementation does not silently replace a dependency.

## Internal Route and Ownership

| Module | Dynamic job | Depends on |
|---|---|---|
| D0. Dynamic Position | Establishes the layer, imports, limits, Scope, judgments, operators, and invariants | Substrate + Ontology + Constitution + Structure |
| D1. Dynamic Forms and Realization | Separates Profile, mechanism, Implementation, active Process, Run, Description, Trace, Record, and realization | D0 |
| D2. Change, Process, and Transformation | Types actual change, maintenance, transformation, continuous and discrete evolution, and regime change | D0–D1 |
| D3. State and Transition Systems | Types state variables, transition rules, guards, enabledness, state machines, and actual Transition Occurrences | D0–D2 |
| D4. Occurrence Lifecycle | Types initiation, activation, suspension, resumption, completion, cancellation, timeout, and ending | D0–D3 |
| D5. Time, Order, Scheduling, and Concurrency | Separates physical, temporal, causal, protocol, and dependency order; types concurrency and scheduling | D0–D4 |
| D6. Interaction, Signal, Message, and Propagation | Types influence, transmission, reception, decoding, delivery, loss, and communication stages | D0–D5 |
| D7. Routing, Qualification, and Crossing | Types route execution, gate evaluation, operational treatment, mediation, and actual Interface crossing | D0–D6 |
| D8. Execution, Effect, and State Commitment | Types invocation, execution, output, state-diff production, commit, rollback, compensation, replay, and external consequence | D0–D7 |
| D9. Holds, Queues, Memory, and Context Lifecycle | Types buffering, release, expiry, eviction, transient context, Session closure, teardown, and retained records | D0–D8 |
| D10. Resource Flow and Operational Capacity | Types allocation, reservation, use, release, contention, exhaustion, backpressure, throughput, and metering | D0–D9 |
| D11. Feedback, Regulation, and Reconfiguration | Types operating feedback, control loops, stabilization, correction, oscillation, and reconfiguration | D0–D10 |
| D12. Failure, Refusal, Correction, and Termination | Types faults, errors, failures, refusal, retry, containment, controlled stopping, and historical preservation | D0–D11 |
| D13. Coordination, Consensus, Commitment, and Finality | Types distributed coordination, ordering, verification, consensus, protocol commitment, and system-relative finality | D0–D12 |
| D14. Observation, Trace, and Conformance | Governs dynamic Claims, measurements, traces, Profiles, conformance, provenance, and correction | D0–D13 |
| D15. Open Dynamic Boundary | Locates unresolved change and fixes downstream and future-profile interfaces | D0–D14 |

The dependency order is one-way:

    SUBSTRATE -> ONTOLOGY -> CONSTITUTION -> STRUCTURE -> DYNAMICS
                                                               |
                                                               v
    D0 -> D1 -> D2 -> D3 -> D4 -> D5 -> D6 -> D7 -> D8 -> D9
      -> D10 -> D11 -> D12 -> D13 -> D14 -> D15

A later module may specialize an earlier dynamic relation. It may not silently redefine an imported Kind, change an earlier operator, make a represented rule perform itself, or convert a downstream continuity or emergence result into dynamic actuality.

---

## D0 — Dynamic Position

### D0.1 Dynamic Function

Ontology states what State, Occurrence, Process, Transition, Interaction, Execution, Signal, Record, Error, Feedback, and Finality mean. Structure states how actual bearers, boundaries, interfaces, paths, state-bearing positions, Systems, and constraints are arranged. Dynamics states **how actual conditions change or are maintained through actual Processes and Occurrences within an exact Scope**.

Dynamics owns **actual change and operation**.

It does not:

- create Reality by declaring a transition;
- redefine an imported Kind;
- make a Dynamic Profile, law, rule, equation, algorithm, state machine, schedule, Description, Trace, or Record into the Process it represents;
- infer an actual Occurrence from its possibility, eligibility, enabledness, prediction, simulation, or later record;
- reduce every physical, biological, cognitive, social, institutional, or mixed Process to computation;
- infer causation from temporal order, adjacency, correlation, common routing, or repeated co-occurrence;
- make an Interface perform a crossing, a Path perform passage, or a Channel perform transmission;
- establish Intention, Agency, identity, Consent, Admission, Permission, Authorization, Authority, Right, Obligation, Ownership, Responsibility, legitimacy, or jurisdiction from operational success;
- establish that a bearer remains numerically the same across change;
- produce any downstream emergent capacity, significance, or value by counting transitions or feedback;
- turn resource use, sponsorship, metering, or scarcity into Price, money, entitlement, or reward;
- turn consensus, commitment, immutability, or protocol finality into Truth or physical indestructibility;
- erase an actual Occurrence by rollback, compensation, Correction, eviction, teardown, Repeal, or later reinterpretation.

An actual Process may occur without a complete dynamic Description. A dynamically exact Description may refer to nothing actual. Dynamic conformance, successful execution, and actuality are separate.

### D0.2 Imported Semantic and Structural Ground

The following are imported unchanged from Ontology:

- `EXISTS`, `OBTAINS`, `SAME`, `DIFFERENT`, `RELATES`, `SCOPED_TO`, `REALIZES`, `REPRESENTS`, `PARTICIPATES`, and `OCCUPIES_REFERENCE_TARGET`;
- the canonical judgment `Gamma |- x : K @ D, I, R`;
- `SUBTYPE_OF`, `INSTANCE_OF`, `DISJOINT_WITH`, `OVERLAPS`, `PLAYS`, and `ATTRIBUTION`;
- `Bearer`, `Entity`, `Occurrence`, `Event`, `Process`, `State`, `Configuration`, `Transition Occurrence`, `Transition Relation`, `Transition Record`, `Trajectory`, `Beginning`, `Ending`, `Interval`, `Physical Time`, and `Ordered Time`;
- `Mechanism`, `Interaction`, `Effect`, `Consequence`, `Coupling`, `Correlation`, `Mediation`, `Constraint`, `Threshold`, `Probability`, `Capacity`, and `Disposition`;
- `Realization`, `Realization Attempt`, `Realization Result`, `Implementation`, `Execution`, `Computation`, `Computational State`, `Computational Execution`, `Computational Process`, `Session`, `Protocol`, `Algorithm`, `Program`, and `Machine`;
- `Signal`, `Noise`, `Data`, `Information`, `Encoding`, `Decoding`, `Message`, `Communication Attempt`, `Transmission`, `Reception`, `Interpretation`, `Channel`, and `Medium`;
- `Network`, `Ledger`, `Block`, `Chain`, `Consensus`, `Smart Contract`, `Technical Interface`, and `Application Programming Interface`;
- `History`, reversibility, `Irreversibility`, `Finality`, `Error`, `Correction`, `Control`, `Feedback`, `Stability`, `Robustness`, `Resilience`, and `Attractor`;
- `Representation`, `ReferenceTargetPosition`, `Record`, `Memory`, `Claim`, `Evidence`, `Observation`, `Measurement`, `Verification`, `Validation`, and the Ontology conformance and Error grammars;
- the separation of Action, Decision, Commitment, Refusal, Admission Act, Access, Participant, Actor, Operator, Controller, Sponsor, Provider, and their bearers.

The following are imported unchanged from Structure:

- `StructuralScope`, `StructuralConfiguration`, `StructuralProfile`, `StructureDescription`, `StructureSnapshot`, `StructureRecord`, `CoveragePosture`, `RealizationBasis`, `DistributedRealizerBasis`, `RepresentationBinding`, and `ProfileCompositionRecord`;
- typed regions, Boundaries, Boundary sides, Interfaces, Port Roles and Slots, Aperture Regions and Roles, Gate and Filter Roles, Seams, Channels, Paths, topologies, Systems, Worlds, Environments, state-bearing positions, and `StructuralCapacity`;
- the separation of actual Configuration, Profile, Description, Record, realization, applicability, and conformance;
- the requirement that an asserted structural route, Interface, Channel, or capacity is only a possible arrangement until an actual dynamic Occurrence uses it.

Capitalized imported terms retain those meanings. Dynamics declares qualified dynamic Kinds, Roles, Relation Kinds, Profiles, States, value-spaces, and operators. It does not become a second Ontology, Structure, Constitution, Continuity, or Emergence layer.

Every actual Dynamic Relation is asserted through `EXISTS` plus `RELATES`, or through a specialized predicate explicitly declared as an existential projection of `RELATES`. Every realization Claim uses imported `REALIZES`. Role occupancy uses imported `PLAYS`. No dynamic declaration creates an alternative existence, truth, identity, or authority operator.

The surface lexemes `Commitment`, `Refusal`, `Hold`, `Admission`, `Effect`, `Finality`, and `Control` occur in more than one domain. Bare use is permitted only where the imported or qualified signature makes the owner and meaning unambiguous. Otherwise the expression is `AMBIGUOUS_OPERATOR` and `ILL_TYPED`.

### D0.3 Layer Boundary

The Reality Model remains:

    SUBSTRATE -> ONTOLOGY -> CONSTITUTION -> STRUCTURE
                -> DYNAMICS -> CONTINUITY -> EMERGENCE

The linear notation states semantic dependency, not a universal chronology.

- **Substrate** constrains every physically realized Process and Occurrence.
- **Ontology** owns the root Kinds and relations used here.
- **Constitution** owns normative validity, institutional force, Admission, Consent reliance, Authorization, constitutional Refusal, Review, Remedy, and Effect.
- **Structure** owns arrangement, typed positions, boundaries, Interfaces, Paths, constraints, and structural capacity.
- **Dynamics** owns actual change, maintenance, transition, timing, execution, routing, crossing, propagation, buffering, feedback operation, resource use, failure handling, operational coordination, and computational or protocol finality.
- **Continuity** owns criterion-bound persistence and continuation relations across change, replacement, copying, succession, branching, merging, dormancy, recovery, dissolution, and lineage. It may evaluate Claims that a bearer remains numerically identical only under Ontology's strict `SAME` relation; no criterion-relative continuity relation redefines or entails numerical identity.
- **Emergence** owns whether processes and organizations realize life, cognition, Agency, Experience, intelligence, meaning, adaptation, resilience, coherence, collective capacity, value, or another emergent capacity not reducible to motion alone.

Dynamics may provide ordered occurrences, before-and-after States, resource traces, feedback traces, and execution consequences to later layers. It may not decide their continuity or emergent significance by motion alone.

### D0.4 Dynamic Scope

A **Dynamic Scope** is a qualified `StructuralScope` adding the exact dynamic target, order context, and structural basis required to evaluate change. It is a direct Scope specialization, not a wrapper containing a second independent World, Domain, Interval, or Resolution Context.

    KIND DynamicScope {
      parents: [StructuralScope]
      category: other
      necessary_conditions: [inherited WorldContext, Domain, Interval, ResolutionContext, and further conditions; nonempty dynamic targets; one OrderContext; one StructuralBasisReference]
      sufficient_conditions: [all required fields typed; inherited and repeated projections coherent; no field contradiction]
      properties_and_value_spaces: [dynamic_targets; order_context: OrderContext; structural_basis: StructuralBasisReference; dynamic_conditions]
      admissible_bearers: NOT_APPLICABLE
      relation_signatures: [SCOPED_TO]
      realization_criterion: NOT_APPLICABLE
      scope: self
      disjoint_with: OPEN
      overlaps: OPEN
      epistemic_statuses: [Observed, Measured, Inferred, Established, Disputed, Unknown]
      non_equivalences: [DynamicScope / Process; DynamicScope / State; DynamicScope / World]
      later_layer_boundaries: [persistence and emergence are not included]
      owner: D0
    }

Its record projection is:

    DynamicScope {
      world_context: inherited WorldContext
      domain: inherited Domain
      interval: inherited Interval
      resolution_context: inherited ResolutionContext
      further_conditions: inherited [Claim | Criterion | Representation] | NONE | OPEN
      dynamic_targets: nonempty ordered [Term]
      order_context: OrderContext
      structural_basis: StructuralBasisReference
      dynamic_conditions: [Claim | Criterion | Constraint | Representation] | NONE | OPEN
    }

    StructuralIntervalBinding {
      interval: Interval
      configuration_terms: [StructuralConfiguration] | NONE | OPEN
      mapping_to_dynamic_scope: Mapping | NONE | OPEN
    }

    StructuralRealizationBinding {
      target: Bearer | RelationInstance | StructuralConfiguration
      interval: Interval
      posture: RealizationPosture
      evidence_claims_or_records: [Claim | Record] | NONE | OPEN
    }

    StructuralEvaluationContext {
      interval_partition_or_mapping:
          nonempty ordered [StructuralIntervalBinding]
        | Mapping
        | OPEN
      realization_postures:
          nonempty [StructuralRealizationBinding]
        | OPEN
      known_bearers_and_relations:
          [Bearer | RelationInstance] | NONE | OPEN
      unresolved_structural_positions:
          [ReferenceTargetPosition | Claim | Representation] | NONE | OPEN
      minimum_dynamic_distinctions_to_preserve:
          nonempty [Criterion | Claim | Representation]
    }

    StructuralBasisReference :=
        CONFIGURATION_BOUND {
          configuration_terms: nonempty ordered [StructuralConfiguration]
          structural_scope: StructuralScope
          structural_evaluation_context: StructuralEvaluationContext
          conformance_references: [StructuralConformanceReport] | NONE | OPEN
        }
      | STRUCTURE_UNRESOLVED {
          partial_structural_scope: StructuralScope | OPEN
          structural_evaluation_context: StructuralEvaluationContext
        }

    OrderKind :=
        PHYSICAL_TEMPORAL_ORDER
      | DERIVED_TEMPORAL_ORDER
      | CAUSAL_ORDER
      | PROTOCOL_ORDER
      | DEPENDENCY_ORDER
      | SCHEDULING_ORDER
      | REPRESENTATIONAL_ORDER

    OrderContext {
      admitted_order_kinds: nonempty [OrderKind]
      physical_time_basis: Claim | Representation | NONE | OPEN
      derived_order_basis: Claim | Representation | NONE | OPEN
      clocks_or_order_sources: [Bearer | Representation] | NONE | OPEN
      interval_segmentation_criterion: Criterion
      concurrency_resolution: Criterion | NONE | OPEN
    }

`CONFIGURATION_BOUND` does not require one unchanging Configuration throughout the Dynamic Scope. Its interval partition must expose every structural change material at the asserted Resolution. `STRUCTURE_UNRESOLVED` permits actual change to remain representable where the complete arrangement is not known; it does not create a structure by placeholder.

World context, Domain, Interval, and Resolution Context occur once through inheritance. Any repetition in a Profile, Description, Run, Trace, outcome, or report must be an equal projection or an explicitly declared sub-scope.

    WF_DYNAMIC_SCOPE(DS)
      -> WELL_FORMED | ILL_TYPED

`WELL_FORMED` requires:

1. the inherited Structural Scope is well formed;
2. every dynamic target is typed or explicitly unresolved;
3. the Order Context distinguishes physical time, derived order, and any protocol order it uses;
4. the structural basis covers the material Interval or declares the exact unresolved position;
5. interval segmentation is adequate to the asserted Resolution;
6. every dynamic condition is typed or explicitly `OPEN`;
7. repeated scope fields are equal projections or declared sub-scopes;
8. no field contradicts another in the same interpretation.

A well-formed Dynamic Scope establishes no change, causation, persistence, or conformance by itself.

### D0.5 Canonical Dynamic Judgments

The dynamic conformance judgment is:

    Gamma |-D SATISFIES_DYNAMICS(dynamic_target_x, profile_p, DS) => e

where:

- `Gamma` is the attributable declarations, Evidence, methods, observations, and assumptions;
- `dynamic_target_x` is an actual Occurrence, Process, Dynamic Run, realized dynamic system, or Dynamics-qualified Relation Instance under examination, or a Term purporting to denote one;
- `profile_p` is the exact Dynamic Profile whose criteria are tested;
- `DS` is one well-formed Dynamic Scope;
- `e` is either an external `ILL_TYPED` result produced before substantive evaluation or one `DynamicConformanceEvaluation` envelope;
- an evaluated envelope contains exactly one substantive Ontology conformance result: `CONFORMS`, `DOES_NOT_CONFORM`, or `UNRESOLVED`.

The application judgment is separate:

    DYNAMICS_APPLIES(profile_p, dynamic_target_x, DS)
      -> ApplicabilityEvaluation | ILL_TYPED

`SATISFIES_DYNAMICS` is evaluated only where `DYNAMICS_APPLIES(...).result = APPLICABLE`. It presupposes separately established existence and sufficient realization of its target. `CONFORMS` never creates an Occurrence or Process. `DOES_NOT_CONFORM` does not make an actual nonconforming Process disappear. `UNRESOLVED` is not a negative result. `ILL_TYPED` identifies a defect in the judgment, not in Reality.

Except for the nonrecursive application judgment, all substantive Dynamics evaluations use one common minimum representation header and envelope:

    DynamicsRepresentationHeader {
      representation_binding: RepresentationBinding
      represents_relation: RelationInstance satisfying exact REPRESENTS
      carrier_and_format_basis
      provenance_claims_or_records: [Claim | Record] | NONE | OPEN
      correction_or_supersession_reference: Record | NONE | OPEN
    }

    DynamicEvaluationEnvelope<T> {
      representation_header: DynamicsRepresentationHeader
      target_references: nonempty [Term]
      profile_and_version_reference:
          Representation | NOT_APPLICABLE
      dynamic_scope: DynamicScope
      interval_or_as_of_boundary: Interval | ReferenceTargetPosition | Representation
      method_or_criterion_references:
          nonempty [Criterion | Representation]
      evidence_references: [Claim | Record] | NONE | OPEN
      typing_posture: WELL_TYPED
      applicability_posture:
          APPLICABLE | NOT_APPLICABLE | UNRESOLVED
      actuality_posture:
          ACTUALITY_ESTABLISHED
        | NONACTUALITY_ESTABLISHED
        | ACTUALITY_UNRESOLVED
        | NOT_APPLICABLE
      realization_posture:
          REALIZATION_ESTABLISHED
        | NONREALIZATION_ESTABLISHED
        | REALIZATION_UNRESOLVED
        | NOT_APPLICABLE
      coverage: CoveragePosture | NOT_APPLICABLE
      evaluation:
          NOT_EVALUATED { blocking_reason }
        | EVALUATED { result: T }
      reason_codes: [Term] | NONE | OPEN
      provenance_references: [Claim | Record] | NONE | OPEN
    }

Every substantive schema named `...Evaluation`, other than the nonrecursive `ApplicabilityEvaluation`, is a specialization of this envelope. A displayed specialized schema may show only its typed payload; the common fields remain mandatory and are not omitted. `NOT_EVALUATED` is a procedure posture, not a substantive result. Each `T` declares its own unresolved posture where required; the envelope does not add a second generic unresolved result. A malformed invocation returns `ILL_TYPED` before an envelope is constructed; `ILL_TYPED` is never a member of `T`.

Nested payloads use a distinct closed marker family:

    NestedEvaluationMarker :=
        DIMENSION_NOT_EVALUATED
      | STAGE_NOT_EVALUATED
      | REFERENCE_NOT_EVALUATED

`DIMENSION_NOT_EVALUATED` marks one independently assessable result dimension, `STAGE_NOT_EVALUATED` marks one lifecycle or propagation stage, and `REFERENCE_NOT_EVALUATED` marks a deferred evaluation or occurrence reference. None is the envelope-level `NOT_EVALUATED`, and none negates an otherwise valid outer `EVALUATED` posture.

The envelope's actuality and realization postures concern the exact primary target or targets named by `target_references`. A payload field that restates actuality or realization for that same target must be an equal projection; a field about a component stage or related occurrence must name that different target explicitly. A negative occurrence assessment may be `EVALUATED` with outer `NONACTUALITY_ESTABLISHED`; the strict `NOT_EVALUATED` actuality gate applies specifically to conformance of a purported actual operative target. `realization_posture` is `NOT_APPLICABLE` where no operative realization is required.

The application judgment uses its own nonrecursive projection:

    ApplicabilityEvaluation {
      representation_header: DynamicsRepresentationHeader
      target_references: nonempty [Term]
      profile_and_version_reference: Representation
      dynamic_scope: DynamicScope
      interval_or_as_of_boundary
      method_or_criterion_references
      evidence_references: [Claim | Record] | NONE | OPEN
      typing_posture: WELL_TYPED
      coverage: CoveragePosture | NOT_APPLICABLE
      result: APPLICABLE | NOT_APPLICABLE | UNRESOLVED
      reason_codes: [Term] | NONE | OPEN
      provenance_references: [Claim | Record] | NONE | OPEN
    }

    DynamicConformanceEvaluation :=
        DynamicEvaluationEnvelope<
          CONFORMS | DOES_NOT_CONFORM | UNRESOLVED
        >

`SATISFIES_DYNAMICS` is not substantively evaluated until applicability and target actuality are established. Unresolved or disproved actuality therefore yields an envelope whose procedure posture is `NOT_EVALUATED`, not a conformance result. `UNRESOLVED` inside conformance is reserved for an actual applicable target whose required conformance conditions remain unresolved.

### D0.6 Registered Operators

| Operator | Owner | Result | Question |
|---|---|---|---|
| `WF_DYNAMIC_SCOPE(DS)` | D0 | `WELL_FORMED \| ILL_TYPED` | Is the dynamic application boundary internally typed? |
| `WF_DYNAMIC_PROFILE(p)` | D1 | `WELL_FORMED \| ILL_TYPED` | Is the represented Profile internally typed, Scope-bound, and explicit about mechanisms and coverage? |
| `WF_DYNAMIC_DESCRIPTION(dd)` | D1 | `WELL_FORMED \| ILL_TYPED` | Is the dynamic Description internally typed and explicit about actuality and coverage? |
| `WF_DYNAMIC_TRACE(dt)` | D14 | `WELL_FORMED \| ILL_TYPED` | Is the Trace internally typed, ordered, attributable, and explicit about coverage? |
| `DYNAMICS_APPLIES(p, x, DS)` | D1 | `ApplicabilityEvaluation \| ILL_TYPED` | Does Profile p apply to x in this exact Scope? |
| `SATISFIES_DYNAMICS(x, p, DS)` | D14 | `DynamicConformanceEvaluation \| ILL_TYPED` | Does actual dynamic target x satisfy Profile p? |
| `CHANGE_ASSESSMENT(prior_position, prior_state, subsequent_position, subsequent_state, selected_respect, criterion, DS)` | D2 | `ChangeEvaluation \| ILL_TYPED` | What change and maintenance postures obtain under the exact respect? |
| `TRANSITION_ENABLED(machine_or_system, transition_rule, state, context, DS)` | D3 | `EnablednessEvaluation \| ILL_TYPED` | Are the rule's operational conditions satisfied? |
| `TRANSITION_OCCURRENCE_ASSESSMENT(assessed_occurrence_or_position, rule_or_kind, prior_position, prior_state, subsequent_position, subsequent_state, criterion, DS)` | D3 | `TransitionOccurrenceEvaluation \| ILL_TYPED` | Did the specified Transition Occurrence obtain? |
| `LIFECYCLE_ASSESSMENT(occurrence_or_run, lifecycle_profile_or_criterion, DS)` | D4 | `LifecycleEvaluation \| ILL_TYPED` | What lifecycle dimensions obtain? |
| `ORDER_ASSESSMENT(a, b, order_kind, order_context, DS)` | D5 | `OrderEvaluation \| ILL_TYPED` | How are a and b related under the exact order kind? |
| `MESSAGE_LIFECYCLE(message_or_signal, communication_profile, DS)` | D6 | `MessageLifecycleEvaluation \| ILL_TYPED` | Which communication stages occurred? |
| `INTENT_AXIS_ASSESSMENT(target, intent_axis_profile, DS)` | D6 | `IntentAxisEvaluation \| ILL_TYPED` | What classification was produced under the exact evidentiary Profile? |
| `GATE_ASSESSMENT(candidate, gate_or_filter, gate_profile, DS)` | D7 | `GateEvaluation \| ILL_TYPED` | What criterion result and actual evaluation occurrence obtained? |
| `OPERATIONAL_TREATMENT(candidate_or_request, treatment_profile, criterion, context, DS)` | D7 | `OperationalTreatmentEvaluation \| ILL_TYPED` | What treatment selection and runtime treatment occurrence obtained? |
| `CROSSING_RESULT(attempt_or_crossing_occurrence, source_manifestation, interface, from_side, to_side, receiving_manifestation, DS)` | D7 | `CrossingEvaluation \| ILL_TYPED` | What crossing dimensions obtained without assuming departure or identity? |
| `RUN_RESULT(run, profile, DS)` | D8 | `RunEvaluation \| ILL_TYPED` | What independently typed execution dimensions obtained? |
| `CONTEXT_LIFECYCLE(context, profile, DS)` | D9 | `ContextLifecycleEvaluation \| ILL_TYPED` | What formation, retention, eviction, and teardown dimensions obtain? |
| `RESOURCE_ACCOUNT(run_or_process, resource_profile, DS)` | D10 | `ResourceAccountEvaluation \| ILL_TYPED` | What resource flows, allocations, uses, and releases obtained? |
| `FEEDBACK_ASSESSMENT(loop_or_process, feedback_profile, DS)` | D11 | `FeedbackEvaluation \| ILL_TYPED` | What per-variable feedback results obtained? |
| `FAILURE_ASSESSMENT(occurrence_or_run, failure_profile, DS)` | D12 | `FailureEvaluation \| ILL_TYPED` | What failure and treatment dimensions obtained? |
| `REMEDIATION_ASSESSMENT(attempt, target, method, profile, DS)` | D12 | `RemediationEvaluation \| ILL_TYPED` | What remediation was attempted and realized? |
| `CONSENSUS_ASSESSMENT(process_or_candidate, consensus_profile, DS)` | D13 | `ConsensusEvaluation \| ILL_TYPED` | What agreement procedure and result obtained? |
| `FINAL_UNDER(result_or_state, finality_profile, DS)` | D13 | `FinalityEvaluation \| ILL_TYPED` | Is the result final under the declared ordinary-revision boundary? |
| `DYNAMIC_PLACEMENT(x, DS)` | D15 | `DynamicPlacementEvaluation \| ILL_TYPED` | Can encountered change be located without forcing a known dynamic form? |

These operators evaluate or represent Claims. None creates its target. Actual relation Claims use imported `RELATES` through the qualified Dynamic Relation Kinds declared later.

This table is a forward declaration index. A result family owned by a later module is named here for discoverability but is not used as a prerequisite by D0; its complete definition occurs only in its owning module.

### D0.7 Dynamic Relation Declaration Schema

Every Relation Kind owned by Dynamics must have a declaration equivalent to:

    DynamicRelationDirection :=
        DIRECTED | UNDIRECTED | MIXED | NOT_APPLICABLE

    DYNAMIC_RELATION R {
      arity: n
      relata: [P1: BearerKind1, ..., Pn: BearerKindN]
      qualifiers: [
        {
          name: Identifier
          value_space
          binding:
              ATTRIBUTION(attribute_kind)
            | SIDE_JUDGMENT(judgment_template)
          required: YES | NO
        }
      ] | NONE
      context: DynamicScope
      ordered_positions: YES | NO | PARTIAL
      direction: DynamicRelationDirection
      symmetry: YES | NO | CONDITIONAL | NOT_APPLICABLE
      reflexivity: YES | NO | CONDITIONAL | NOT_APPLICABLE
      transitivity: YES | NO | CONDITIONAL | NOT_APPLICABLE
      inverse: R2 | NONE | CONDITIONAL | OPEN
      composition_rule: rule | NONE | OPEN
      realization_criterion: Criterion | OPEN
      non_entailments: [...]
      owner: dynamics_module
    }

`arity` counts actual relata only. Rules, modes, Profiles, State labels, clocks, criteria, and Dynamic Scopes are qualifiers or context unless independently declared as actual relata. A Term in a Claim purports to denote a bearer; the sign itself does not become the relatum.

The following minimal qualifier families are declared here so later relations do not depend on hindsight from D14:

    TransformationKind := registered value | OPEN
    TransitionKind := registered value | OPEN
    PropagationKind := registered value | OPEN
    CrossingKind := COPY | MOVE | TRANSFORM | REFERENCE | OBSERVE | OPEN
    FeedbackKind := POSITIVE | NEGATIVE | MIXED | OTHER_REGISTERED | OPEN

`CriterionQualifier` and its `Criterion` value-space are imported unchanged from Structure. The following additional Property Kinds are Dynamics-owned, admit only `RelationInstance` as bearer, and are used by `ATTRIBUTION` in Dynamics-owned relations:

| Qualifier Property Kind | Value space |
|---|---|
| `ToleranceQualifier` | `Criterion \| Representation` |
| `OrderContextQualifier` | `OrderContext` |
| `TransformationKindQualifier` | `TransformationKind` |
| `PreservationClaimsQualifier` | `[Claim] \| NONE \| OPEN` |
| `LossClaimsQualifier` | `[Claim] \| NONE \| OPEN` |
| `TransitionKindQualifier` | `TransitionKind` |
| `AppliedRuleOrProfileQualifier` | `Representation \| NONE \| OPEN` |
| `PropagationKindQualifier` | `PropagationKind` |
| `PropagationDirectionQualifier` | `Representation \| OPEN` |
| `RateOrDelayQualifier` | `Quantity Value \| NONE \| OPEN` |
| `AttenuationOrLossQualifier` | `Quantity Value \| Claim \| NONE \| OPEN` |
| `CrossingKindQualifier` | `CrossingKind` |
| `MappingOrTransformationQualifier` | `Mapping \| Process \| NONE \| OPEN` |
| `FeedbackKindQualifier` | `FeedbackKind` |
| `SelectedVariableQualifier` | `ReferenceTargetPosition` |

D14 registers and audits these declarations; it does not create them retroactively.

The existential projection of a dynamic Relation Kind is:

    claim_R := attributable Claim that
      R(actual_bearer_relata; typed_qualifiers; DS) obtains

    j_R := Judgment evaluating or expressing claim_R in DS

    SCOPED_TO(j_R, DS)

    exists ri [
      EXISTS(ri, DS.domain, DS.interval)
      and RELATES(ri, R, [actual_bearer_relata], DS.domain, DS.interval)
      and for each qualifier q_k its declared binding holds:
            ATTRIBUTION(declared_attribute_k, ri, q_k,
                        DS.domain, DS.interval, DS.resolution_context)
        or  SIDE_JUDGMENT(declared_judgment_template_k[q_k])
    ]

An omitted algebraic property is `OPEN`, never silently true. Relation names do not determine causation, direction, reversibility, transitivity, exclusivity, persistence, or finality. A Profile may narrow a base declaration; it may not change its arity, actual relatum positions, or base meaning.

### D0.8 Dynamic Result and Error Discipline

Every enumerated value belongs to exactly one declared value family. Its canonical identity is `Family.Member`; for an inline value space, the exact owning record-field path is the family identity. An unqualified spelling is only a local display abbreviation where the expected family is uniquely fixed by the field type. Equal spellings in different families are not equal values. In particular, bare `MIXED` is `ILL_TYPED` wherever more than one family could be expected; examples include `DynamicRelationDirection.MIXED`, `FeedbackKind.MIXED`, `DynamicDescriptionPosture.MIXED`, `ContextRetentionPosture.MIXED`, and `FeedbackAggregatePosture.MIXED`. The same rule applies to every reused spelling such as `PARTIAL`, `OPEN`, `UNRESOLVED`, `COMPLETED`, or `NOT_APPLICABLE`.

The Ontology Error Model remains applicable. `DynamicsReasonCode` is the closed value family consisting exactly of the following Dynamics-owned reason codes:

- `DYNAMIC_SCOPE_MISMATCH` — World, Domain, Interval, Resolution, structural basis, or order context cannot be reconciled;
- `DYNAMIC_RECORD_CAST` — a Description, State snapshot, Trace, receipt, result, or Record is treated as the actual Occurrence or Process;
- `DYNAMIC_PROFILE_ACTUALITY_CAST` — a Profile, rule, schedule, algorithm, or state machine is treated as its own operation;
- `STATE_BEARER_CAST` — a State or state label is treated as the bearer whose State it is;
- `POSSIBILITY_ACTUALITY_CAST` — a possible, predicted, simulated, eligible, enabled, or intended transition is treated as actual;
- `ENABLED_OCCURRENCE_CAST` — enabledness is treated as transition occurrence;
- `ORDER_KIND_COLLAPSE` — physical, temporal, causal, protocol, dependency, scheduling, or representational order is used as another without a declared mapping;
- `TEMPORAL_CAUSAL_CAST` — temporal precedence is treated as productive causation or causal dependence;
- `ROUTE_TRAVERSAL_CAST` — a represented route or structurally available Path is treated as actual traversal;
- `INTERFACE_CROSSING_CAST` — Interface or Aperture availability is treated as actual crossing;
- `SIGNAL_INTENTION_CAST` — signal classification or behavior is treated as actual Intention or Agency;
- `OPERATIONAL_NORMATIVE_CAST` — operational acceptance, hold, refusal, completion, or success is treated as constitutional Admission, Consent, Permission, Authorization, validity, Refusal, or Effect;
- `COMMITMENT_KIND_COLLAPSE` — protocol or state commitment is treated as agential or institutional Commitment;
- `FINALITY_KIND_COLLAPSE` — one finality kind is exported into another without a declared relation;
- `FINALITY_TRUTH_CAST` — consensus, immutability, or finality is treated as Truth, actuality of represented content, or physical indestructibility;
- `EXECUTION_VALIDITY_CAST` — execution, completion, or success is treated as correctness, lawfulness, Authorization, or constitutional validity;
- `RESOURCE_VALUE_CAST` — resource quantity, use, sponsorship, scarcity, or metering is treated as Price, money, value, entitlement, Ownership, or reward;
- `FEEDBACK_EMERGENCE_CAST` — recurrence, feedback, stabilization, or control operation is treated as adaptation, intelligence, Agency, coherence, or value;
- `TRANSITION_IDENTITY_CAST` — before-and-after State references are treated as proof of numerical persistence across change;
- `ROLLBACK_ERASURE_CAST` — rollback, inverse operation, compensation, Correction, eviction, or teardown is treated as erasure of an actual Occurrence;
- `MISSING_DYNAMIC_REALIZATION` — an operative law, Profile, process, or mechanism is asserted without a sufficient realizer and criterion;
- `MISSING_PRIOR_OR_SUBSEQUENT_CONDITION` — a claimed transition omits a required boundary State or exact open/unresolved posture;
- `COVERAGE_OVERCLAIM` — a Trace, observation set, or result is used exhaustively outside its declared Coverage Posture;
- `SILENT_PROGRESSION` — satisfaction of one stage or condition is treated as performance of a later Occurrence without a declared mechanism and actual trigger;
- `UNDECLARED_DEFAULT` — silence, timeout, absence, or omitted input is assigned a result without an applicable declared rule;
- `TRACE_ORDER_CYCLE` — a declared strict order contains a cycle without being typed as recurrent or equivalence structure;
- `QUALIFIER_TYPE_ERROR` — a Dynamic Relation qualifier lacks a declared attribute or side-judgment binding.

Malformed expressions return `ILL_TYPED` before substantive evaluation. For a well-typed conformance evaluation:

1. if one decisive necessary condition is established false, return `DOES_NOT_CONFORM`;
2. if no decisive failure is established and one required dependency remains unresolved, return `UNRESOLVED`;
3. return `CONFORMS` only where every required condition is established in the exact Scope;
4. keep `NOT_EVALUATED` outside the closed conformance result.

For every other Dynamics-owned evaluation, the same discipline applies: typing is checked first; an ill-typed call returns `ILL_TYPED` externally; a well-typed but blocked evaluation returns the common envelope with `NOT_EVALUATED`; and only an evaluated envelope carries its declared typed result, including that result family's own unresolved posture where applicable.

An error remains representable as a fact about an expression or execution. Rejecting the expression does not erase the attempt that produced it.

### D0.9 Dynamic Invariants

Every later module preserves:

1. **Actuality separation** — Profile, possibility, enabledness, prediction, simulation, Trace, and Record are not actual change.
2. **State–Process separation** — a State is a condition; a Process is an Occurrence. Neither substitutes for the other.
3. **Continuous-change admission** — change need not be discrete, computational, state-machine-governed, deterministic, or fully observed.
4. **Realization requirement** — operative dynamic form requires an actual realizer, mechanism, and criterion or an explicit unresolved position.
5. **No silent progression** — one stage never performs the next merely by being satisfied.
6. **No causation from order** — before, after, simultaneous, adjacent, correlated, or commonly routed does not establish causation.
7. **No transition from topology** — Interface, Aperture, Channel, Path, connectivity, and capacity do not establish passage or operation.
8. **No normativity from operation** — enabledness, access, acceptance, refusal, completion, consensus, and success do not establish constitutional force.
9. **No persistence from change** — comparison of States does not establish that one numerically identical bearer spans them.
10. **No emergence from repetition** — feedback, recurrence, complexity, activity, block production, or transition density does not establish adaptation, intelligence, Agency, coherence, or value.
11. **No value from resource flow** — resource consumption, provision, sponsorship, and metering do not establish economic meaning or entitlement.
12. **No erasure** — reversal of selected effects does not reverse historical actuality.
13. **Typed finality** — every finality Claim states its System, rule, result, revision boundary, Scope, and kind.
14. **Open-world coverage** — omission does not establish non-occurrence unless exact closure conditions are satisfied.
15. **No self-qualification or self-succession** — a proposed Profile, Implementation, Configuration, or version may be produced, simulated, verified, and validated, but it cannot establish its own applicability, validity, Currentness, Authority, precedence, replacement, or succession. Those positions require the already-applicable governing basis and the owning constitutional or continuity procedure.

---

## D1 — Dynamic Forms and Realization

### D1.1 Actual Dynamic Target

An **actual dynamic target** is an existing Occurrence, Process, Interaction, Transition Occurrence, Execution, Dynamics-qualified Relation Instance, or organized system while relevant change or maintenance actually obtains within a Dynamic Scope.

Actuality does not require that the Process be:

- intentional;
- computational;
- discrete;
- deterministic;
- controlled;
- observed;
- successfully completed;
- conforming to any human Description;
- institutionally recognized.

A natural Process may occur without being produced from a represented rule. A represented rule may accurately describe or constrain a Process without causing it. Where a rule is claimed to operate causally, the actual mechanism realizing its practical effect must be identified.

### D1.2 Dynamic Profile

A **Dynamic Profile** is a Representation declaring criteria for a class of changes, Processes, interactions, transitions, executions, or dynamic systems.

    KIND DynamicProfile {
      parents: [Representation]
      category: Pattern
      necessary_conditions: [target class or predicate; DynamicScope; typed states or conditions where applicable; transition or process criteria; realization requirements; application criterion; coverage; RepresentationBinding]
      sufficient_conditions: [WF_DYNAMIC_PROFILE returns WELL_FORMED for its exact represented form]
      properties_and_value_spaces: [version_or_form_reference; profile_kind; application_criterion; process_or_transition_constraints; coverage]
      admissible_bearers: NOT_APPLICABLE
      relation_signatures: [REPRESENTS]
      realization_criterion: represented Profile physically encoded in a recoverable carrier
      scope: DynamicScope
      disjoint_with: OPEN
      overlaps: [Model, Protocol, Algorithm, Record]
      epistemic_statuses: [Posited, Hypothetical, Established, Disputed]
      non_equivalences: [DynamicProfile / Process; DynamicProfile / DynamicRun; DynamicProfile / constitutional Rule]
      later_layer_boundaries: [persistence and emergent value excluded]
      owner: D1
    }

A Dynamic Profile may be:

- **descriptive** — representing observed or hypothesized change;
- **predictive** — generating expectations under a stated Model;
- **classificatory** — stating criteria for a dynamic kind;
- **operational** — specifying transitions for an implemented System or Protocol;
- **control-oriented** — specifying observations and responses;
- **conformance-oriented** — stating required, optional, excluded, or unresolved behavior;
- **open** — leaving declared mechanisms, outcomes, or branches unresolved.

These Profile kinds may overlap. None creates a physical Law, constitutional Rule, Permission, Obligation, causal mechanism, or actual Process merely by being represented.

Every Dynamic Profile makes the following recoverable where applicable:

    DynamicProfile {
      profile_id_or_form_reference
      representation_header: DynamicsRepresentationHeader
      profile_kind
      target_kind_or_predicate
      dynamic_scope: DynamicScope
      state_or_condition_schema: Representation | NONE | OPEN
      process_kinds: [Kind | Representation] | NONE | OPEN
      transition_rule_representations: [Representation] | NONE | OPEN
      mechanisms_or_realization_requirements: [Mechanism | Criterion | Representation] | NONE | OPEN
      order_and_timing_rules: [Representation] | NONE | OPEN
      interaction_and_crossing_rules: [Representation] | NONE | OPEN
      execution_rules: [Representation] | NONE | OPEN
      context_and_memory_rules: [Representation] | NONE | OPEN
      resource_rules: [Representation] | NONE | OPEN
      feedback_and_control_rules: [Representation] | NONE | OPEN
      failure_and_refusal_rules: [Representation] | NONE | OPEN
      coordination_and_finality_rules: [Representation] | NONE | OPEN
      outcome_schema: Representation | NONE | OPEN
      application_criterion
      coverage: CoveragePosture
      unresolved: [Claim | Term] | NONE
    }

`NONE`, `OPEN`, and omission remain distinct. A Profile need not contain every field family when the family is not applicable. A required field cannot be silently omitted.

    WF_DYNAMIC_PROFILE(p)
      -> WELL_FORMED | ILL_TYPED

Well-formedness requires an exact target, Dynamic Scope, application criterion, typed dynamic conditions, realization requirements, result spaces, Coverage Posture, and Representation Binding. It does not establish that the represented dynamics occur, are possible, are correct, or govern any target.

### D1.3 Dynamic Implementation, Activation, and Run

An imported **Implementation** is an Artifact or organized arrangement produced or configured to realize a specification, Protocol, design, or formal operation. A **Dynamic Implementation** is an Implementation whose claimed target is a Dynamic Profile, transition system, control process, or execution procedure.

Deployment, availability, activation, invocation, and running are distinct:

- **deployment** places or configures an Implementation in a structural environment;
- **availability** is a scoped capacity for use;
- **activation** is an Occurrence beginning its operative posture;
- **invocation** is an Occurrence requesting or triggering a specified operation;
- **running** is the active interval of an actual Process.

A **Dynamic Run** is a bounded Process in which a Dynamic Implementation, natural mechanism, or organized System realizes one declared dynamic operation or process context.

    KIND DynamicRun {
      parents: [Process]
      category: Occurrence
      necessary_conditions: [actual interval; actual participating bearers or processes; DynamicScope; realization basis; exact run-boundary criterion]
      sufficient_conditions: [existence and active realization established under the declared run criterion]
      properties_and_value_spaces: [run_identifier: Identifier | NONE; run_boundary_criterion; process_posture_claims; outcome; coverage]
      admissible_bearers: NOT_APPLICABLE
      relation_signatures: [PARTICIPATES, REALIZES]
      realization_criterion: run-specific Criterion
      scope: DynamicScope
      disjoint_with: OPEN
      overlaps: [Execution, Interaction, Session]
      epistemic_statuses: [Observed, Measured, Inferred, Established, Disputed, Unknown]
      non_equivalences: [DynamicRun / DynamicProfile; DynamicRun / trace Record; DynamicRun / Session]
      later_layer_boundaries: [run continuity and value excluded]
      owner: D1
    }

`DynamicRun` does not claim that every natural Process is a run. It is used only where a bounded operational or observational context identifies one.

A run identifier identifies a represented or operational position. It does not establish numerical identity of the participating bearer, user, Agent, Session, or later Run.

### D1.4 Dynamic Realization Basis

Every Claim that a Dynamic Profile, transition system, feedback loop, Protocol, or execution procedure actually operates must expose:

    DynamicExecutionPath :=
        ACTUAL_PATH {
          mechanism_or_process: Mechanism | Process
          occurrences: [Occurrence] | NONE | OPEN
          exact_realizes_relation: RelationInstance
          criterion: Criterion
          evidence_claims_or_records: [Claim | Record] | NONE | OPEN
        }
      | REPRESENTED_AND_BOUND_PATH {
          representation: Representation
          representation_binding: RepresentationBinding
          bound_actual_path: ACTUAL_PATH
        }
      | PATH_UNRESOLVED

    DynamicRealizationBasis {
      representation_header: DynamicsRepresentationHeader
      represented_realizes_claim: Claim
      realizes_relation_instance: RelationInstance | OPEN
      realizer_position:
          SINGLE { bearer_or_process: Bearer | Process }
        | DISTRIBUTED { basis: DistributedRealizerBasis }
      realized_dynamic_target: DynamicProfile | Process | Mechanism | Representation
      dynamic_scope: DynamicScope
      active_interval
      mechanism_or_execution_path: DynamicExecutionPath
      criterion
      evidence_claims_or_records: [Claim | Record] | NONE | OPEN
      result:
          REALIZATION_ESTABLISHED
        | NONREALIZATION_ESTABLISHED
        | REALIZATION_UNRESOLVED
    }

This Record represents exactly its `represented_realizes_claim`: `representation_header.representation_binding.represented_content` equals that Claim, while `representation_header.represents_relation` matches the binding's representation, target position, mapping, Domain, and Interval. Where realization is established, `realizes_relation_instance` and every `ACTUAL_PATH.exact_realizes_relation` must denote the same exact `REALIZES` Relation Instance, or explicitly mapped equal projections of it in the same Dynamic Scope. The Record does not replace the imported relation or become the realizer.

For an operative Dynamic Profile, `REALIZATION_ESTABLISHED` requires more than the existence of a compatible structure or executable artifact. It requires `ACTUAL_PATH`, directly or through `REPRESENTED_AND_BOUND_PATH`, an actual `REALIZES` Relation Instance, and the characteristic Process during the exact active Interval under the stated criterion. `PATH_UNRESOLVED`, an unbound Representation, or `OPEN` cannot support `REALIZATION_ESTABLISHED`. A malformed basis is `ILL_TYPED` before substantive realization assessment. A dormant Implementation may realize a Capacity while no active Run occurs.

Distributed realization preserves Structure's requirements: actual plural bearers and their organizing Relation Instances are named, no Record substitutes for them, and isolated sufficiency is assessed under the joint criterion.

### D1.5 Dynamic Description, Snapshot, Trace, and Record

A **Dynamic Description** is a Representation making Claims about actual, possible, predicted, simulated, or prescribed change.

A **Dynamic Snapshot** is a Description of selected State and dynamic conditions at an as-of point or bounded Interval. It may display rates, active Processes, and pending transitions, but it is not the unfolding Process.

A **dynamic trace record** is a generic Record preserving selected observations or represented steps from one or more Occurrences or Processes together with an exact ordering basis. D1 separates that representational artifact from actual operation and declares the base `DynamicTrace` Kind and projection below; D14 adds observation-specific well-formedness, evidentiary, and conformance criteria.

    KIND DynamicTrace {
      parents: [Record]
      category: Pattern
      necessary_conditions: [RepresentationBinding; DynamicScope; exact order basis; nonempty or explicitly empty entry posture; CoveragePosture; provenance]
      sufficient_conditions: [all necessary fields typed and coherent for its exact represented form]
      properties_and_value_spaces: [trace_kind; ordering_basis; coverage; correction_posture]
      admissible_bearers: NOT_APPLICABLE
      relation_signatures: [REPRESENTS]
      realization_criterion: persistent recoverable physical representation under the declared carrier and format criterion
      scope: DynamicScope
      disjoint_with: OPEN
      overlaps: [Record]
      epistemic_statuses: [Observed, Measured, Inferred, Established, Disputed, Unknown]
      non_equivalences: [DynamicTrace / actual History; DynamicTrace / Process; DynamicTrace / Transition Occurrence]
      later_layer_boundaries: [persistence and lineage not inferred]
      owner: D1
    }

Every Dynamic Trace makes this base projection recoverable:

    DynamicTrace {
      trace_id: Identifier | NONE
      representation_header: DynamicsRepresentationHeader
      target_terms
      dynamic_scope: DynamicScope
      profile_and_implementation_references: [Representation] | NONE | OPEN
      order_context: OrderContext
      entries: ordered [DynamicTraceEntry] | EXPLICITLY_EMPTY
      entry_order_kind
      clock_and_timestamp_basis: Representation | NONE | OPEN
      coverage: CoveragePosture
      collection_method
      filtering_sampling_and_aggregation
      dropped_or_missing_entry_posture
      integrity_or_tamper_claims: [Claim] | NONE | OPEN
      evidence_claims_or_records: [Claim | Record] | NONE | OPEN
      provenance_claims_or_records: [Claim | Record] | NONE | OPEN
      correction_and_supersession_links: [Record | Representation] | NONE | OPEN
      unresolved: [Claim | Term] | NONE
    }

    DynamicTraceEntry {
      entry_id: Identifier | NONE
      represented_occurrence_state_or_result_position
      entry_kind
      represented_order_position
      timestamp_or_clock_reading: Representation | NONE | OPEN
      source_or_observer_claim
      payload_or_result
      local_scope_or_equal_projection
      provenance
      uncertainty
    }

Trace order is not automatically physical, causal, or execution order. The `entry_order_kind` is mandatory.

A **Dynamic Record** is a persistent Representation preserving a Dynamic Description, trace record, result, or provenance position.

These remain distinct from:

- the bearer whose condition changes;
- the State attributed to that bearer;
- the Occurrence or Process described;
- the Profile or rule used to interpret it;
- the physical carrier realizing the Record;
- the actual past;
- any later replay or simulation using the Record.

Every Dynamics-owned Representation subtype must expose the imported `RepresentationBinding` and exact `REPRESENTS` relation. This includes every schema named `Profile`, `Rule`, `Plan`, `Description`, `Snapshot`, `Declaration`, `Record`, `Trace`, `Evaluation`, `Report`, or `Bundle` unless the schema explicitly establishes that it is an actual non-representational bearer. A shortened projection later in this Form inherits these fields; display omission never waives them.

The common inherited representation header is declared in D0.5 so no later Representation rule is needed to type an earlier evaluation.

The represented content never becomes its Reference Target or external Referent by appearing in the Record.

Every exact Dynamic Description makes this projection recoverable:

    DynamicDescriptionPosture :=
        ACTUALITY_CLAIM
      | POSSIBILITY_CLAIM
      | PREDICTION
      | SIMULATION_RESULT
      | PRESCRIPTION
      | MIXED { exact_partition }

    DynamicDescription {
      description_id: Identifier | NONE
      representation_header: DynamicsRepresentationHeader
      description_posture: DynamicDescriptionPosture
      target_terms: nonempty [Term]
      dynamic_scope: DynamicScope
      dynamic_profile_reference: DynamicProfile | NONE | OPEN
      prior_conditions_or_states: [State | Claim | Representation] | NONE | OPEN
      occurrences_or_processes: [Occurrence | Process | Claim] | NONE | OPEN
      subsequent_conditions_or_states: [State | Claim | Representation] | NONE | OPEN
      mechanisms: [Mechanism | Claim | Representation] | NONE | OPEN
      ordering_basis: OrderContext
      interactions_and_effects: [RelationInstance | Claim] | NONE | OPEN
      realization_basis: DynamicRealizationBasis | OPEN
      coverage: CoveragePosture
      evidence_claims_or_records: [Claim | Record] | NONE | OPEN
      method_description: Representation | OPEN
      provenance_claims_or_records: [Claim | Record] | NONE | OPEN
      unresolved: [Claim | Term] | NONE
    }

    WF_DYNAMIC_DESCRIPTION(dd)
      -> WELL_FORMED | ILL_TYPED

Well-formedness establishes typed and internally coherent representation. It does not establish content truth, target existence, actual change, complete history, or conformance.

### D1.6 Coverage and Negative Dynamic Claims

Dynamics imports `CoveragePosture` unchanged.

No transition, signal, message, failure, or effect may be declared absent merely because it does not appear in an `OPEN`, `PARTIAL`, or `COVERAGE_UNRESOLVED` Trace.

An exhaustive negative dynamic Claim requires `CLOSED_FOR` coverage naming:

- the exact event or process family;
- the observed bearers or region;
- the Interval and order basis;
- the Resolution Context;
- the detection method and sensitivity;
- the closure criterion;
- material blind positions and their treatment.

“No signal detected,” “no transition recorded,” and “no failure reported” are results of methods. They do not automatically mean no Signal, Transition Occurrence, or failure existed.

### D1.7 Dynamic Profile Composition

Dynamic Profiles compose only through:

    DynamicProfileCompositionRecord {
      representation_header: DynamicsRepresentationHeader
      profiles: nonempty [DynamicProfile]
      resultant_profile: DynamicProfile
      target_kind_or_predicate
      shared_dynamic_scope: DynamicScope
      scope_partition_or_overlap
      state_and_variable_mappings: [Mapping | Representation] | NONE | OPEN
      order_composition
      priority_rules: [Representation] | NONE | OPEN
      conflict_rules: [Representation] | NONE | OPEN
      preserved_constraints: [Constraint | Representation] | NONE | OPEN
      overridden_constraints: [Constraint | Representation] | NONE
      result_composition
      unresolved: [Claim | Term] | NONE
    }

No Dynamic Profile silently overrides another by recency, specificity, deployment dominance, execution frequency, naming, or wider use. Profiles governing disjoint Scopes may coexist. Profiles addressing one atomic matter require explicit composition, priority, or conflict treatment.

The resultant Profile is represented by, but is not identical to, the composition Record. A proposed resultant Profile may be generated, implemented, simulated, and tested. It cannot use its own composition, test results, or execution dominance to establish its applicability, validity, Currentness, Authority, precedence, replacement, or succession.

A constitutional Effect may require use of a Dynamic Profile only through the exact valid constitutional gates. That result neither changes the Profile's semantic content nor makes its Implementation operate.

### D1.8 Dynamic Profile Application

    DYNAMICS_APPLIES(profile_p, dynamic_target_x, DS)
      -> ApplicabilityEvaluation | ILL_TYPED

`APPLICABLE` requires:

1. p is a well-formed Dynamic Profile;
2. x satisfies the Profile's target Kind or exact target predicate without Role casting;
3. the Profile's World, Domain, Interval, Resolution Context, structural basis, and order conditions include the exact application;
4. the exact Profile Form or version is identified;
5. required composition and priority rules are satisfied;
6. no declared exclusion applies.

The decision order is complete:

1. malformed arguments, category errors, missing mandatory binding, or an ill-formed Scope or Profile return `ILL_TYPED` before an Applicability Evaluation exists;
2. a well-typed established target-kind mismatch, Scope mismatch, version mismatch, or applicable exclusion returns `NOT_APPLICABLE`, even where unrelated application facts remain unresolved;
3. where no decisive mismatch or exclusion is established but a necessary target predicate, Scope inclusion, composition, precedence, Evidence, or coverage dependency remains unresolved, return `UNRESOLVED`;
4. return `APPLICABLE` only where every necessary application condition is established and no exclusion applies.

The Applicability Evaluation records the exact method, Evidence, coverage, version, Interval or as-of boundary, and reasons. Target actuality remains a separate gate and is not required merely to decide that a Profile would apply to a correctly typed target position.

Applicability is not actuality, enabledness, conformance, Permission, or constitutional Currentness. No Dynamic Profile applies globally by omission.

### D1.9 Profile-Specific Vocabulary

Universal Dynamics installs no project-specific, protocol-specific, economic, historical, symbolic, or domain-local label as a universal dynamic primitive.

A later qualified Profile may introduce such vocabulary only after declaring:

- their category and owner;
- exact bearer or occurrence positions;
- entry and exit conditions;
- transition and result grammar;
- Scope and realization basis;
- interaction with Constitution, Continuity, and Emergence;
- non-equivalences and non-entailments.

Their absence here preserves rather than rejects future use while preventing the universal layer from preloading their meaning.

---

## D2 — Change, Process, and Transformation

### D2.1 Change Is Criterion- and Resolution-Bound

Change obtains where a relevant Property, Quantity, State, Relation, Configuration, organization, or process condition differs across an ordered extent under a stated criterion and Resolution Context.

The same bearer may be:

- unchanged under one selected variable;
- changing under another;
- maintained by active Processes at one scale;
- fluctuating below the Resolution of another account.

“Changed” and “unchanged” are incomplete without the selected respect, Interval, order, and Resolution.

### D2.2 Maintenance Is Dynamic

Maintenance is a Process through which a condition remains within declared bounds despite ongoing interactions, dissipation, replacement, perturbation, or internal activity.

Maintenance is not stasis. A maintained macrostate may depend on continual microstate change. Conversely, an unchanging representation does not prove that its purported Referent remained unchanged.

    DYNAMIC_RELATION MAINTAINS {
      arity: 3
      relata: [maintaining_process: Process, bearer: Bearer, maintained_state_or_condition: Bearer]
      qualifiers: [
        { name: criterion; value_space: Criterion; binding: ATTRIBUTION(CriterionQualifier); required: YES }
        { name: tolerated_range; value_space: Criterion | Representation; binding: ATTRIBUTION(ToleranceQualifier); required: YES }
        { name: order_basis; value_space: OrderContext; binding: ATTRIBUTION(OrderContextQualifier); required: YES }
      ]
      context: DynamicScope
      ordered_positions: YES
      direction: DIRECTED
      symmetry: NO
      reflexivity: CONDITIONAL
      transitivity: NO
      inverse: NONE
      composition_rule: OPEN
      realization_criterion: actual process causally sustains the condition within the declared bounds
      non_entailments: [stasis, persistence, identity, purpose, success, value]
      owner: D2
    }

The maintained State or condition must occupy a declared Bearer position in the relation grammar. A proposition or State Record is not silently substituted for the condition represented.

### D2.3 Continuous, Discrete, Hybrid, and Stochastic Change

Dynamics admits at least these description and realization families:

- **continuous** — selected variables evolve over an ordered extent without represented jumps at the declared Resolution;
- **discrete** — selected conditions change among distinguished states or events;
- **hybrid** — continuous evolution and discrete transitions interact;
- **stochastic** — outcomes or paths are represented by a Probability structure under a Model;
- **recurrent or oscillatory** — trajectories revisit regions or conditions under an order;
- **developmental** — organization changes through staged but not necessarily reversible Processes;
- **transformative** — the relevant Form, phase, or organization changes;
- **maintenance-oriented** — active Process preserves a condition within bounds.

These families classify an account under a Resolution Context. They do not exhaust Reality, and one Process may overlap several.

A continuous Process need not possess a finite state-machine representation. A discrete Model of continuous change remains a Model. A stochastic description does not imply ontological randomness beyond the exact interpretation declared.

### D2.4 Dynamic Variable and Parameter

A **Dynamic Variable** is a represented quantity or classified condition whose value is permitted to differ across a Dynamic Scope.

A **Parameter** is a represented value treated as fixed within one Model or Profile application. It may vary in another account.

    DynamicVariableDeclaration {
      representation_header: DynamicsRepresentationHeader
      variable_position: ReferenceTargetPosition
      represented_property_or_state_dimension
      admissible_bearer_kind
      value_space
      unit_or_classification_rule
      order_basis
      sampling_or_resolution_conditions
      uncertainty
      missing_value_posture
    }

The variable, its symbol, measured value, represented Property, and actual bearer remain distinct.

### D2.5 Transformation

A **Dynamic Transformation** is an actual Process in which a selected input condition, State, material, signal, organization, or representation changes into a selected output condition under a Mechanism.

    DYNAMIC_RELATION TRANSFORMS {
      arity: 4
      relata: [process: Process, input_bearer_or_state: Bearer, mechanism: Bearer, output_bearer_or_state: Bearer]
      qualifiers: [
        { name: transformation_kind; value_space: TransformationKind | OPEN; binding: ATTRIBUTION(TransformationKindQualifier); required: YES }
        { name: preservation_claims; value_space: [Claim] | NONE | OPEN; binding: ATTRIBUTION(PreservationClaimsQualifier); required: YES }
        { name: loss_claims; value_space: [Claim] | NONE | OPEN; binding: ATTRIBUTION(LossClaimsQualifier); required: YES }
        { name: criterion; value_space: Criterion; binding: ATTRIBUTION(CriterionQualifier); required: YES }
      ]
      context: DynamicScope
      ordered_positions: YES
      direction: DIRECTED
      symmetry: NO
      reflexivity: CONDITIONAL
      transitivity: CONDITIONAL
      inverse: CONDITIONAL
      composition_rule: requires compatible output/input positions, Scope, and actual intermediate Processes
      realization_criterion: process-specific Criterion
      non_entailments: [numerical identity of input and output, reversibility, improvement, purpose, value]
      owner: D2
    }

Transformation does not by itself establish that the input and output are the same bearer. That question belongs to Continuity. It also does not establish creation from nothing, destruction of all constituents, improvement, or intended design.

### D2.6 Regimes, Thresholds, and Phase Change

A **Dynamic Regime** is a represented or established range of conditions within which a selected family of dynamic relations or response patterns obtains.

A **Regime Transition** is a Transition Occurrence crossing between regimes under an exact criterion.

Threshold satisfaction does not cause a regime transition by itself. The actual coupling or mechanism connecting the threshold condition to the change must obtain or remain explicitly unresolved.

Abrupt change in one variable need not imply discontinuity in every underlying variable. Phase, bifurcation, tipping, activation, and constitutional change are different kinds and must not be merged by the word “transition.”

### D2.7 Constraints, Invariants, and Conservation

A dynamic Constraint excludes transitions or trajectories under its declared Domain. An **Invariant** is an imported Property or Relation preserved under a specified transformation.

For each claimed invariant, the Profile declares:

- what is preserved;
- under which transformation family;
- for which bearers and Interval;
- at which Resolution;
- within what tolerance, if measured;
- which exceptions or open positions remain.

Conservation in a physical theory, data-integrity preservation, constitutional fidelity, identity persistence, and economic balance are different relations. No shared arithmetic form makes them equivalent.

### D2.8 Reversibility and Inverse Processes

Dynamics imports the Ontology distinction among physical reversibility, logical reversibility, transactional reversal, Correction, refund, and institutional annulment.

An inverse process may restore selected prior conditions without:

- erasing the forward Process;
- restoring every microscopic condition;
- restoring the same bearer;
- cancelling external consequences;
- correcting a false Record;
- annulling normative effect.

Every reversibility Claim states the relevant variables, tolerance, permitted inverse, conditions, cost or resource basis where material, and Resolution Context.

### D2.9 Change Assessment

    ChangePosture :=
        CHANGED
      | NO_RELEVANT_CHANGE
      | CHANGE_UNRESOLVED

    MaintenancePosture :=
        MAINTAINED
      | MAINTENANCE_NOT_ESTABLISHED
      | MAINTENANCE_UNRESOLVED
      | NOT_APPLICABLE

    ChangeResult {
      selected_respect
      prior_state_bearing_position: Bearer
      prior_state: State
      subsequent_state_bearing_position: Bearer
      subsequent_state: State
      change_occurrence_or_process: Occurrence | Process | NONE | OPEN
      change_posture: ChangePosture
      maintenance_posture: MaintenancePosture
      maintenance_process: Process | NONE | OPEN
      tolerated_range: Criterion | Representation | NONE | OPEN
      continuity_premise:
          NOT_REQUIRED
        | SUPPLIED { judgment_reference }
        | UNRESOLVED
    }

    ChangeEvaluation := DynamicEvaluationEnvelope<ChangeResult>

    CHANGE_ASSESSMENT(prior_position, prior_state,
                      subsequent_position, subsequent_state,
                      selected_respect, criterion, DS)
      -> ChangeEvaluation | ILL_TYPED

`CHANGED` means the selected condition differs under the criterion and an actual change Occurrence or Process is referenced. If the State difference is established but its change occurrence or process remains `OPEN`, the change posture remains `CHANGE_UNRESOLVED`. `MAINTAINED` means an actual maintaining Process preserved the selected condition within declared bounds. They may co-obtain where internal change occurs inside a maintained macro-condition. `NO_RELEVANT_CHANGE` and `MAINTENANCE_NOT_ESTABLISHED` require coverage `CLOSED_FOR` the exact respect, Interval, and Resolution; missing observations yield the corresponding unresolved posture.

`CHANGE_ASSESSMENT` does not decide Continuity. If the same-bearer premise is required and not independently available, the change result may concern related State positions while bearer persistence remains `UNRESOLVED`.

### D2.10 State Difference and State Delta

A **State Difference** is a comparison Relation between selected prior and subsequent States.

A **State Delta** is a Representation encoding selected differences and possibly an operation for applying them to a represented State.

A State Delta is not:

- the Transition Occurrence;
- the Process causing it;
- every Effect of the Process;
- proof that it was applied;
- proof that the prior and subsequent States belong to one persistent bearer.

---

## D3 — State and Transition Systems

### D3.1 State Selection

Every dynamic account identifies the selected State variables, relations, and resolving conditions. A State is attributed to a Bearer; it is not a free-floating complete copy of the bearer.

Terms such as `initial`, `current`, `pending`, `active`, `terminal`, and `error` are state-position classifications within an exact Profile. They are not universal ontological states.

A “world state,” “account state,” “application state,” biological state, physical State, and constitutional Status differ in bearer, Domain, criterion, and effect.

### D3.2 State Space and State Schema

An imported **State Space** is a formal space whose points represent possible States under a Model.

A **State Schema** is a Representation declaring selected variables, admissible values, relations, constraints, and unresolved positions for State descriptions.

Neither is the actual collection of all possible real conditions. Reachability within a represented State Space does not prove physical possibility; failure to represent a State does not prove impossibility.

### D3.3 Transition Rule

A **Transition Rule** is a Representation specifying conditions under which a transition kind is permitted, predicted, selected, or required by an operational Profile and how a subsequent represented State is determined or constrained. D1 stores such rules only as bound Representations; this module supplies the narrower `TransitionRule` Kind.

    KIND TransitionRule {
      parents: [Representation]
      category: Pattern
      necessary_conditions: [source-state predicate; target-state predicate or transformation; trigger or activation condition; guard and preconditions; result grammar; DynamicScope; RepresentationBinding]
      sufficient_conditions: [all rule fields typed and internally coherent]
      properties_and_value_spaces: [rule_kind; priority; atomicity_posture; determinacy_posture]
      admissible_bearers: NOT_APPLICABLE
      relation_signatures: [REPRESENTS]
      realization_criterion: physically encoded rule interpreted by a compatible mechanism where operative use is claimed
      scope: DynamicScope
      disjoint_with: OPEN
      overlaps: [Algorithm, Protocol, Constraint]
      epistemic_statuses: [Posited, Hypothetical, Established, Disputed]
      non_equivalences: [TransitionRule / Transition Occurrence; TransitionRule / Permission; TransitionRule / physical Law]
      later_layer_boundaries: [continuity and value excluded]
      owner: D3
    }

Rule kinds include descriptive, predictive, stochastic, operational, control, and conformance rules. “Permitted” inside an operational state machine means allowed by that Profile; it is not constitutional Permission.

### D3.4 Trigger, Guard, Precondition, and Postcondition

- A **Trigger** is a represented or actual condition or Occurrence used by a mechanism to initiate evaluation or execution.
- A **Guard** is a represented predicate evaluated to determine operational enabledness.
- A **Precondition** is a condition required before an operation or transition can conform to its Profile.
- A **Postcondition** is a condition tested after an operation or transition under that Profile.

These remain distinct:

- trigger occurrence does not establish guard satisfaction;
- guard satisfaction does not perform the transition;
- precondition satisfaction does not establish Permission or Authorization;
- postcondition satisfaction does not establish that the declared mechanism caused the result;
- failure of a postcondition does not erase the attempted execution.

Every condition evaluation exposes its inputs, method, as-of State, result, uncertainty, and Record posture.

    ConditionResult :=
        CONDITION_SATISFIED
      | CONDITION_NOT_SATISFIED
      | CONDITION_UNRESOLVED

    ConditionEvaluation := DynamicEvaluationEnvelope<ConditionResult>

### D3.5 Enabledness

    EnablednessResult :=
        ENABLED
      | DISABLED
      | ENABLEMENT_UNRESOLVED

    EnablednessEvaluation :=
        DynamicEvaluationEnvelope<EnablednessResult>

    TRANSITION_ENABLED(machine_or_system, transition_rule, state, context, DS)
      -> EnablednessEvaluation | ILL_TYPED

`ENABLED` means the exact operational trigger, guard, and preconditions declared by the Rule are satisfied or resolved as the Profile specifies. It does not mean that the transition occurs, should occur, is authorized, is safe, or will succeed.

`DISABLED` requires a decisive failed condition. Missing information yields `ENABLEMENT_UNRESOLVED`, unless the Profile contains an explicit typed default. Silence never becomes Consent.

### D3.6 State-Machine Profile and Realization

A **State-Machine Profile** is a Dynamic Profile declaring:

    StateMachineProfile {
      representation_header: DynamicsRepresentationHeader
      represented_machine_position
      dynamic_scope: DynamicScope
      state_schema
      initial_state_criteria
      transition_rules
      terminal_state_criteria: [Criterion] | NONE | OPEN
      invariant_constraints
      scheduling_and_conflict_rules
      invalid_transition_handling
      failure_and_timeout_rules
      observation_and_trace_requirements
      coverage: CoveragePosture
    }

An **actual state-machine realization** is an actual System and Process satisfying the Profile's realization criterion during an Interval. The Profile, code, transition table, deployed Implementation, active Run, and Trace are distinct.

Not every Process is a state machine. A state-machine projection may be useful and incomplete without exhausting the Process represented.

### D3.7 Transition Relation

The Dynamics-qualified transition projection is:

    DYNAMIC_RELATION DYNAMIC_TRANSITION {
      arity: 5
      relata: [prior_state_bearing_position: Bearer, prior_state: State, transition_occurrence: Transition Occurrence, subsequent_state_bearing_position: Bearer, subsequent_state: State]
      qualifiers: [
        { name: transition_kind; value_space: TransitionKind | OPEN; binding: ATTRIBUTION(TransitionKindQualifier); required: YES }
        { name: applied_rule_or_profile; value_space: TransitionRule | DynamicProfile | NONE | OPEN; binding: ATTRIBUTION(AppliedRuleOrProfileQualifier); required: NO }
        { name: criterion; value_space: Criterion; binding: ATTRIBUTION(CriterionQualifier); required: YES }
        { name: order_basis; value_space: OrderContext; binding: ATTRIBUTION(OrderContextQualifier); required: YES }
      ]
      context: DynamicScope
      ordered_positions: YES
      direction: DIRECTED
      symmetry: NO
      reflexivity: CONDITIONAL
      transitivity: NO
      inverse: CONDITIONAL
      composition_rule: requires an actual compatible intermediate State and actual intervening Occurrences
      realization_criterion: actual State change under the stated criterion
      non_entailments: [numerical persistence, reversibility, purpose, Authorization, validity, value]
      owner: D3
    }

This qualifies Ontology's Transition Relation; it does not redefine it. The prior position and subsequent position are independently typed relata connected through an actual Transition Occurrence and selected State change. A transition rule or Record remains a qualifier or Representation, not a substitute for the Occurrence.

The relation does not assert `SAME`, persistence, succession, replacement, or identity between the two state-bearing positions. Those remain Continuity questions.

### D3.8 Determinacy, Branching, and Choice

A transition-oriented Dynamic Profile may be:

- deterministic under declared inputs and State;
- nondeterministic;
- stochastic under a probability model;
- underdetermined because material inputs are unresolved;
- externally selected;
- selected through an Agent's Action;
- selected through a control or consensus procedure.

These are different positions. Nondeterminism is not Agency. Determinism is not lack of Agency. More than one represented branch does not establish that every branch is physically possible.

### D3.9 Atomic and Compound Transitions

An **Atomicity Profile** states which selected intermediate States are not externally observable or independently committed under one operational boundary.

A **Compound Transition** is an organized set or order of Transition Occurrences treated as one operation under a Profile.

Atomicity is Scope-relative. It does not mean instantaneous in physical time, causally simple, incapable of partial external consequence, or indivisible under every Resolution.

### D3.10 No-Op, Self-Transition, and Maintenance Transition

A represented transition may return the same selected state label. This can mean:

- a genuine no-op under the selected variables;
- active maintenance;
- a Process whose effects cancel at the observed boundary;
- change below Resolution;
- an attempted transition that produced no committed State Delta.

Those cases require separate classification. Equality of prior and subsequent selected values does not establish that nothing happened.

`DYNAMIC_TRANSITION` obtains only where the selected State actually changes under the exact criterion. A true no-op, active maintenance without selected State change, cancelled effect without selected State change, and attempted but uncommitted target change are actual operational Occurrences but are not `DYNAMIC_TRANSITION` in that respect. Active maintenance may instantiate `MAINTAINS`; attempts remain attempt Occurrences. Below-resolution variation is classified only under a finer declared Resolution.

    StateBoundaryOccurrenceKind :=
        SELECTED_STATE_CHANGE
      | ACTIVE_MAINTENANCE_WITHOUT_SELECTED_STATE_CHANGE
      | NO_OP_OCCURRENCE
      | CANCELLED_EFFECT_WITHOUT_SELECTED_STATE_CHANGE
      | ATTEMPT_WITHOUT_TARGET_COMMIT
      | BELOW_RESOLUTION_VARIATION
      | STATE_BOUNDARY_CLASSIFICATION_UNRESOLVED

### D3.11 Transition Occurrence Assessment

    TransitionOccurrenceResult {
      assessed_occurrence_or_position:
          Transition Occurrence | ReferenceTargetPosition | Claim
      actual_occurrence_reference: Transition Occurrence | NONE | OPEN
      prior_state_bearing_position: Bearer
      prior_state: State
      subsequent_state_bearing_position: Bearer
      subsequent_state: State
      occurrence_posture: OCCURRED | DID_NOT_OCCUR | OCCURRENCE_UNRESOLVED
      state_boundary_kind: StateBoundaryOccurrenceKind | NOT_APPLICABLE
      continuity_judgment_reference:
          Claim | Record | REFERENCE_NOT_EVALUATED | NOT_APPLICABLE
    }

    TransitionOccurrenceEvaluation :=
        DynamicEvaluationEnvelope<TransitionOccurrenceResult>

    TRANSITION_OCCURRENCE_ASSESSMENT(
      assessed_occurrence_or_position, rule_or_kind,
      prior_position, prior_state,
      subsequent_position, subsequent_state,
      criterion, DS)
      -> TransitionOccurrenceEvaluation | ILL_TYPED

`OCCURRED` requires an exact `actual_occurrence_reference` and adequate before-and-after or process Evidence under the declared criterion. `DID_NOT_OCCUR` uses `NONE` and `NOT_APPLICABLE` where no actual target occurrence or State-boundary classification exists. A Trace alone may participate in Evidence but cannot make the result true by existing.

The downstream continuity reference is `NOT_APPLICABLE` where no continuity question is material and `REFERENCE_NOT_EVALUATED` where such a question is material but unevaluated. It never gates, creates, or negates the dynamic Occurrence result.

`DID_NOT_OCCUR` requires adequate `CLOSED_FOR` coverage of the exact transition family and Interval. A rejected or failed attempt may itself be an Occurrence even when the target Transition Occurrence did not occur.

### D3.12 Invalid and Nonconforming Transition Attempts

An attempted transition can be:

- ill-typed as represented;
- disabled under the Profile;
- executed despite being disabled;
- structurally impossible;
- physically unsuccessful;
- computationally failed;
- completed with an unexpected output;
- operationally refused;
- constitutionally unauthorized or invalid;
- externally consequential despite no committed internal State.

These dimensions are independent. “Invalid transaction” or “failed transition” is incomplete until the failed dimension and actual occurrence are identified.

---

## D4 — Occurrence Lifecycle

### D4.1 Lifecycle Is Not One Boolean

An Occurrence or Run may be planned, invoked, started, active, suspended, resumed, completed, failed, refused, cancelled, timed out, or ended. These are lifecycle positions, not synonyms for existence, conformance, success, validity, or finality.

The lifecycle grammar is dimensional:

    LifecyclePhase :=
        PRE_START
      | STARTING
      | ACTIVE
      | SUSPENDED
      | RESUMING
      | ENDING
      | ENDED
      | PHASE_UNRESOLVED

    LifecycleResult {
      invocation:
          NOT_INVOKED | INVOKED | INVOCATION_UNRESOLVED
      phase: LifecyclePhase
      progress:
          NOT_STARTED | IN_PROGRESS | INTERRUPTED
        | COMPLETION_BOUNDARY_REACHED | PROGRESS_UNRESOLVED
      completion:
          NOT_COMPLETED | PARTIALLY_COMPLETED
        | COMPLETED | COMPLETION_UNRESOLVED
      success_under_named_criterion:
          DIMENSION_NOT_EVALUATED | CRITERION_SATISFIED
        | CRITERION_NOT_SATISFIED | SUCCESS_UNRESOLVED
      terminal_cause:
          NOT_TERMINAL | ORDINARY_END | FAILURE_END
        | OPERATIONAL_REFUSAL_END | CANCELLATION_END
        | TIMEOUT_END | ABORT_END | TERMINAL_CAUSE_UNRESOLVED
      output_production:
          NOT_REQUIRED | DIMENSION_NOT_EVALUATED | NO_OUTPUT_ESTABLISHED
        | OUTPUT_PRODUCED | PARTIAL_OUTPUT | OUTPUT_UNRESOLVED
      state_commitment_reference: Claim | Record | REFERENCE_NOT_EVALUATED | NOT_APPLICABLE
      external_effect_reference: Claim | Record | REFERENCE_NOT_EVALUATED | NOT_APPLICABLE
      remediation_reference: Claim | Record | REFERENCE_NOT_EVALUATED | NOT_APPLICABLE
      recording_posture:
          RECORD_NOT_REQUIRED | RECORD_REQUIRED_NOT_FOUND
        | RECORD_PRESERVED | RECORD_INCOMPLETE | RECORD_POSTURE_UNRESOLVED
    }

    LifecycleEvaluation := DynamicEvaluationEnvelope<LifecycleResult>

The dimensions may co-obtain where coherent: a Run may reach its completion boundary while failing a postcondition. Incoherent tuples are `ILL_TYPED`, including `NOT_INVOKED + ACTIVE`, `PRE_START + COMPLETED`, or an actuality posture `NONACTUALITY_ESTABLISHED` paired with `IN_PROGRESS`. A history may contain several successive phase evaluations, each with its own Interval.

### D4.2 Invocation and Start

An **Invocation** is an Occurrence presenting an operation, Process, or Run position for initiation under a Profile.

Invocation does not prove:

- acceptance;
- enabledness;
- actual start;
- identity or Agency of a purported invoker;
- Permission or Authorization;
- successful completion.

A **Start Occurrence** crosses the Profile's declared beginning criterion. A queued, scheduled, or accepted request has not necessarily started.

### D4.3 Suspension and Resumption

Suspension is an actual Process-state change in which selected characteristic operation ceases temporarily while a continuation route remains represented or realized.

Resumption is a later Occurrence restarting selected operation from an exact retained, reconstructed, or replacement context.

Whether a resumed Process is the same Process or a successor belongs to Continuity. Dynamics records the State relation and actual Occurrence without deciding numerical identity.

### D4.4 Completion, Success, and Outcome

**Completion** means the declared execution or process boundary has been reached. **Success** means the applicable success criterion is satisfied. An **Outcome** is the resulting condition or classified result.

Completion may coexist with:

- failed postconditions;
- no committed State Delta;
- external adverse effects;
- constitutional invalidity;
- unresolved correctness;
- later reversal or compensation.

An incomplete Process may already have effects. A successful local operation may fail its containing compound operation.

### D4.5 Cancellation, Timeout, Abort, and Ending

- **Cancellation** is an Occurrence applying a cancellation rule before ordinary completion.
- **Timeout** is an operational result triggered when a declared temporal or ordering condition is met.
- **Abort** is termination under an abnormal-stop rule.
- **Ending** is the boundary at which the selected Occurrence or Process ceases under its criterion.

These are not interchangeable. Timeout does not prove refusal, fault, Consent withdrawal, or non-Agency. Cancellation does not necessarily reverse prior effects. Ending a Run does not establish dissolution of its bearer or System.

### D4.6 Automatic Progression

A Profile may declare an actual mechanism that automatically invokes a later stage when conditions are satisfied. The mechanism, trigger, and later Occurrence must still be realized.

Dependency satisfaction, readiness, eligibility, schedule position, or prior completion alone does not perform the next stage.

Silence, absence, or elapsed time produces a default result only where an exact typed rule applies. Such a runtime default is not Consent, Refusal, Admission, waiver, or constitutional Effect merely because the Profile uses those words.

### D4.7 Lifecycle Assessment

    LIFECYCLE_ASSESSMENT(occurrence_or_run,
                         lifecycle_profile_or_criterion, DS)
      -> LifecycleEvaluation | ILL_TYPED

The assessment exposes:

- exact target;
- lifecycle criterion and Profile;
- start and end boundaries where known;
- current and historical postures;
- cause or trigger Claims where asserted;
- effects already produced;
- State commitment posture;
- failure, refusal, timeout, or cancellation reasons;
- Evidence and Coverage Posture;
- unresolved positions.

No terminal lifecycle posture automatically creates protocol finality. No `ACTIVE` posture proves healthy operation, conformance, value, or persistence.

---

## D5 — Time, Order, Scheduling, and Concurrency

### D5.1 Order Kinds Remain Distinct

Dynamics uses the closed built-in `OrderKind` family declared in D0.4.

- **Physical temporal order** is constrained by physical spacetime and measurement.
- **Derived temporal order** is a system-defined sequence realized through physical processes.
- **Causal order** represents established productive or dependence relations, not mere precedence.
- **Protocol order** is assigned by declared protocol rules.
- **Dependency order** states which operation or condition depends on another.
- **Scheduling order** states planned or selected service position.
- **Representational order** is the order in a list, file, Trace, diagram, or narrative.

No order kind silently exports into another. A block height, sequence number, timestamp, queue position, or page order does not by itself establish physical temporal precision or causation.

### D5.2 Clock, Tick, and Timestamp

A **Clock** is a physical or organized system used to produce ordered readings under a calibration or protocol.

A **Tick** is a distinguished clock or protocol occurrence used as an ordering or measuring mark.

A **Timestamp** is a Record or represented value assigned under a clock or ordering source.

Clock reading, timestamp, and time are distinct. A timestamp may be inaccurate, forged, delayed, coarse, ambiguous, or assigned after the Occurrence represented. Multiple clocks require an explicit synchronization or mapping account; equality of displayed readings does not establish simultaneity.

### D5.3 Interval, Duration, Deadline, and Timeout

Dynamics imports `Interval` and `Duration`. Every material temporal Claim states:

- the order kind;
- endpoints and their open or closed posture;
- the clock, frame, or ordering source where applicable;
- granularity and uncertainty;
- mapping to physical time where such a Claim is made.

A **Deadline** is a represented condition relating a required or expected Occurrence to an ordered boundary. It does not create an Obligation by itself.

A **Timeout Rule** is a Dynamic Profile rule assigning an operational result when its declared ordering condition is satisfied. A **Timeout Occurrence** is the actual operation of that rule. The Rule is not its Occurrence.

### D5.4 Precedence and Partial Order

Order may be total, partial, cyclic, recurrent, branching, or unresolved, as the declared relation permits.

    OrderResult :=
        PRECEDES
      | FOLLOWS
      | SAME_ORDER_POSITION
      | SIMULTANEOUS_UNDER
      | CONCURRENT
      | INCOMPARABLE
      | ORDER_UNRESOLVED

    OrderEvaluation := DynamicEvaluationEnvelope<OrderResult>

    ORDER_ASSESSMENT(a, b, order_kind, order_context, DS)
      -> OrderEvaluation | ILL_TYPED

`SIMULTANEOUS_UNDER` requires an exact clock, frame, tolerance, and Resolution Context. `CONCURRENT` means the relevant active Intervals overlap or lack an imposed serialization under the declared system; it does not require physical simultaneity. `INCOMPARABLE` means the declared partial order supplies no comparison. It is not uncertainty.

For `CAUSAL_ORDER`, a result stronger than `ORDER_UNRESOLVED` requires the imported causal criterion and mechanism or dependence basis. Temporal precedence alone is insufficient.

### D5.5 Scheduling and Priority

A **Schedule** is a Representation assigning candidate operations to ordered positions, Intervals, resources, or conditions.

A **Scheduler** is a Role occupied by a System or Process that selects or assigns operations under a Scheduling Profile.

A **Priority** is an ordering value under a declared selection rule. Priority does not imply moral importance, constitutional rank, entitlement, urgency in every Domain, or guaranteed execution.

    SchedulingProfile {
      representation_header: DynamicsRepresentationHeader
      candidate_operation_kind
      eligibility_criterion
      order_rule
      priority_value_space
      resource_constraints
      preemption_rule
      starvation_or_fairness_criterion: Criterion | NONE | OPEN
      cancellation_and_timeout_rules
      observation_and_trace_requirements
      dynamic_scope: DynamicScope
    }

Scheduling an operation does not invoke or execute it. Fairness is not inferred from round-robin, equal numerical shares, or absence of visible starvation; it requires its own declared criterion and, where normative, constitutional basis.

### D5.6 Concurrency, Interleaving, and Race

**Concurrency** is a dynamic relation in which Processes or Occurrences have overlapping active potential or execution under the declared order without one total order governing every material step.

An **Interleaving** is a represented or actual order of selected steps from concurrent Processes.

A **Race Condition** is a dependence of result on unresolved or uncontrolled order among operations whose ordering matters under the Profile.

Concurrency does not establish independence. Processes may be coupled, contend for resources, observe shared State, or produce conflicts. Apparent sequential logs do not prove that execution lacked concurrency.

### D5.7 Atomicity, Isolation, and Serialization

Operational atomicity is the treatment of selected steps as one commitment boundary under a Profile. **Isolation** states which intermediate effects are not available to which observers or Processes. **Serialization** assigns a total protocol order equivalent under a declared criterion to the relevant concurrent execution.

These are separate:

- atomicity does not imply isolation;
- isolation does not imply absence of physical interaction;
- serialization does not prove the serialized order was physical execution order;
- a serializable result does not establish fairness, validity, or Truth.

### D5.8 Causal Discipline

A productive causal Claim identifies actual Occurrences or standing Processes, an actual physical or organized Mechanism, and the Effect produced. A causal-dependence Claim identifies its counterfactual contrast and conditions.

Dynamics rejects these substitutions:

- earlier therefore cause;
- routed through therefore controlled by;
- correlated therefore interacting;
- repeated therefore necessary;
- omitted therefore causally absent;
- intended therefore caused;
- benefited therefore responsible;
- recorded together therefore one process.

Order may constrain causal possibility while remaining insufficient to establish causal actuality.

---

## D6 — Interaction, Signal, Message, and Propagation

### D6.1 Interaction Requires Influence

An actual Interaction requires a realized coupling or mechanism through which one or more bearers, States, or Processes influence another. Structural contact, proximity, matching frequency, shared context, or simultaneous observation does not by itself establish Interaction.

Dynamics may type one-way influence and mutual influence. “Interaction” does not imply equality, reciprocity, communication, awareness, Agency, Consent, or shared State.

### D6.2 Propagation

**Propagation** is a Process through which a disturbance, influence, wave, signal, material, or represented content changes its availability across positions or interacting bearers under a Mechanism.

    DYNAMIC_RELATION PROPAGATES_THROUGH {
      arity: 5
      relata: [propagating_process: Process, source_manifestation: Bearer, medium_or_channel: Bearer, receiving_position_or_bearer: Bearer, receiving_manifestation: Bearer]
      qualifiers: [
        { name: propagation_kind; value_space: PropagationKind | OPEN; binding: ATTRIBUTION(PropagationKindQualifier); required: YES }
        { name: propagation_direction; value_space: Representation | OPEN; binding: ATTRIBUTION(PropagationDirectionQualifier); required: YES }
        { name: rate_or_delay; value_space: Quantity Value | NONE | OPEN; binding: ATTRIBUTION(RateOrDelayQualifier); required: NO }
        { name: attenuation_or_loss; value_space: Quantity Value | Claim | NONE | OPEN; binding: ATTRIBUTION(AttenuationOrLossQualifier); required: NO }
        { name: criterion; value_space: Criterion; binding: ATTRIBUTION(CriterionQualifier); required: YES }
      ]
      context: DynamicScope
      ordered_positions: YES
      direction: DIRECTED
      symmetry: NO
      reflexivity: CONDITIONAL
      transitivity: CONDITIONAL
      inverse: CONDITIONAL
      composition_rule: requires actual compatible successive propagation segments
      realization_criterion: domain-specific coupling and detected or otherwise established effect
      non_entailments: [message delivery, semantic reception, causation beyond declared mechanism, purpose, authorization]
      owner: D6
    }

The source and receiving manifestations are separately typed. Propagation does not establish their numerical identity, persistence, exact equivalence, or succession.

Propagation can occur without a semantic Message. Message transmission can fail despite Signal propagation where encoding, detection, or Channel conditions are inadequate.

### D6.3 Signal Lifecycle

A Signal lifecycle distinguishes:

1. physical formation or emission;
2. entry into a Medium or Channel;
3. propagation;
4. arrival at a receiving region or Interface;
5. detection by a receiving mechanism;
6. discrimination under a Signal criterion;
7. any later decoding, classification, or response.

Each stage may occur while a later stage fails. “Signal received” must state whether it means physical arrival, detector response, successful discrimination, or application-level acceptance.

Noise is relative to a criterion; classifying variation as noise does not make its physical source unreal.

### D6.4 Message and Communication Lifecycle

A Message is organized for semantic reception under an Encoding. Its lifecycle is represented dimensionally:

    CommunicationProfile SUBTYPE_OF DynamicProfile

    CommunicationProfile {
      representation_header: DynamicsRepresentationHeader
      target_message_or_signal_kind: Kind | Representation
      encoding_and_decoding_requirements: [Representation | Criterion] | NONE | OPEN
      source_channel_and_receiving_position_requirements: [Claim | Representation] | NONE | OPEN
      required_lifecycle_stages: nonempty [Representation]
      ordering_fidelity_and_loss_criteria: [Criterion | Representation] | NONE | OPEN
      acknowledgement_and_response_rules: [Representation] | NONE | OPEN
      application_criterion: Criterion
      coverage: CoveragePosture
      dynamic_scope: DynamicScope
    }

    MessageLifecycleResult {
      target_message_or_signal
      dynamic_scope
      formation:
          NOT_REQUIRED | STAGE_NOT_EVALUATED | FORMED | NOT_FORMED | FORMATION_UNRESOLVED
      emission:
          NOT_REQUIRED | STAGE_NOT_EVALUATED | EMITTED | NOT_EMITTED | EMISSION_UNRESOLVED
      transmission:
          NOT_REQUIRED | STAGE_NOT_EVALUATED | TRANSMITTED | PARTIALLY_TRANSMITTED | NOT_TRANSMITTED | TRANSMISSION_UNRESOLVED
      arrival:
          NOT_REQUIRED | STAGE_NOT_EVALUATED | ARRIVED | PARTIALLY_ARRIVED | DID_NOT_ARRIVE | ARRIVAL_UNRESOLVED
      detection:
          NOT_REQUIRED | STAGE_NOT_EVALUATED | DETECTED | NOT_DETECTED | DETECTION_UNRESOLVED
      decoding:
          NOT_REQUIRED | STAGE_NOT_EVALUATED | DECODED | PARTIALLY_DECODED | NOT_DECODED | DECODING_UNRESOLVED
      interpretation:
          NOT_REQUIRED | STAGE_NOT_EVALUATED | INTERPRETED | MISINTERPRETED | NOT_INTERPRETED | INTERPRETATION_UNRESOLVED
      acknowledgement:
          NOT_REQUIRED | STAGE_NOT_EVALUATED | ACKNOWLEDGED | NOT_ACKNOWLEDGED | ACKNOWLEDGEMENT_UNRESOLVED
      response:
          NOT_REQUIRED | STAGE_NOT_EVALUATED | RESPONSE_OCCURRED | NO_RESPONSE_ESTABLISHED | RESPONSE_UNRESOLVED
      coverage
      evidence
      unresolved
    }

    MessageLifecycleEvaluation :=
        DynamicEvaluationEnvelope<MessageLifecycleResult>

    MESSAGE_LIFECYCLE(message_or_signal, communication_profile, DS)
      -> MessageLifecycleEvaluation | ILL_TYPED

A single token such as `SENT` or `RECEIVED` cannot replace these dimensions where the distinction is material.

### D6.5 Delivery, Duplication, Loss, Corruption, and Reordering

**Delivery** is the actual arrival of the selected Message content at the selected receiving position under an applicable communication or routing Profile. It does not establish detection, interpretation, belief, understanding, Consent, or action.

A transmission system may:

- lose all or part of a Signal or Message;
- duplicate represented content;
- corrupt or transform it;
- deliver out of represented order;
- delay it;
- deliver to an unintended position;
- deliver more than once.

Every fidelity Claim declares the preserved distinctions, tolerated loss, encoding, method, and receiving criterion. Bitwise equality does not establish equality of meaning, context, referent, or effect.

### D6.6 Source, Sender, Receiver, and Attribution

`Source`, `Sender`, and `Receiver` are scoped Roles in a signal or communication process. Occupying them does not establish Agent, human identity, author, Controller, owner, or Authority.

A Key, account, endpoint, device, signature, or network address may support a provenance Claim under its technical assumptions. It does not prove who intended, understood, authorized, or physically initiated the Message.

Attribution requires its own criterion, Evidence, and uncertainty. Automation, delegation, proximity, benefit, and technical control cannot silently fill missing agential attribution.

### D6.7 Intent-Axis Qualification

An **Intent-Axis Profile** is a Dynamic Profile whose purpose is to classify represented Signals, Messages, behavior, or request context along declared evidentiary dimensions relevant to an Intention Claim.

    IntentAxisProfile {
      representation_header: DynamicsRepresentationHeader
      target_kind_or_predicate
      feature_and_evidence_bindings
      extraction_or_observation_method
      classification_value_space
      classification_criterion_and_thresholds
      uncertainty_and_calibration_method
      adversarial_and_failure_conditions
      coverage: CoveragePosture
      dynamic_scope: DynamicScope
    }

    IntentAxisResult {
      target_signal_message_or_request
      profile_reference
      features_or_evidence_claims
      classification: <value admitted by classification_value_space> | NONE | OPEN
      confidence_or_uncertainty
      adversarial_and_failure_conditions
      coverage
      provenance
      result:
          CLASSIFIED
        | NOT_CLASSIFIED
        | CLASSIFICATION_UNRESOLVED
    }

    IntentAxisEvaluation := DynamicEvaluationEnvelope<IntentAxisResult>

    INTENT_AXIS_ASSESSMENT(target, intent_axis_profile, DS)
      -> IntentAxisEvaluation | ILL_TYPED

`CLASSIFIED` requires a classification value admitted by the applicable Profile and neither `NONE` nor `OPEN`. `NOT_CLASSIFIED` requires `classification = NONE`. `CLASSIFICATION_UNRESOLVED` requires `classification = OPEN` unless the Profile expressly preserves a provisional value whose unresolved status is separately stated.

Its operation may produce a classification, Claim, Record, or Evidence relation. It cannot by score alone establish:

- actual Intention;
- Agency;
- human or derivative identity;
- consciousness or understanding;
- Consent;
- constitutional Admission;
- truthfulness;
- normative worth.

An automated Process may produce an intent-like pattern. A human may fail a classifier. Those facts concern the Profile and Evidence, not automatic ontology of either bearer.

### D6.8 Sessions

A Session is an imported bounded interval of interaction or process State associated with a context.

Session formation, authentication, Message exchange, suspension, expiry, closure, and teardown are distinct Occurrences. Session continuity does not establish Agent identity. Multiple Sessions may concern one bearer; one Session may mediate multiple bearers; neither fact decides Continuity.

---

## D7 — Routing, Qualification, and Crossing

### D7.1 Structural Possibility Is Not Traversal

Structure supplies Paths, Channels, Interfaces, Ports, Apertures, Gates, Filters, Seams, topologies, and capacities. Dynamics alone types their actual use.

The following remain distinct:

- available Path;
- planned route;
- selected route;
- traversal attempt;
- partial traversal;
- completed traversal;
- delivery;
- accepted result.

A route can be valid structurally and fail dynamically. Actual passage may occur along an unmodeled or nonconforming route.

### D7.2 Route Plan and Route Execution

A **Route Plan** is a Representation of a proposed or selected sequence of positions, Links, Channels, or Interfaces. It is an intended route only where an independent Intention relation has been established.

    RoutePlan {
      representation_header: DynamicsRepresentationHeader
      route_segments_and_positions
      proposal_or_selection_posture
      structural_profile_and_scope_reference
      segment_conditions
      dynamic_scope: DynamicScope
    }

A **Route Execution** is an actual Process attempting or realizing traversal through one or more route segments.

    RouteExecutionRecord {
      representation_header: DynamicsRepresentationHeader
      subject_or_payload_term
      route_plan: RoutePlan | NONE | OPEN
      actual_segments: [Claim | RelationInstance] | NONE | OPEN
      per_segment_entry_and_exit
      transformations
      losses_or_duplicates
      delays_and_order
      failures_and_alternatives
      final_delivery_posture
      dynamic_scope
      coverage
      evidence
      provenance
    }

The Record does not become the traversal.

### D7.3 Gate and Filter Evaluation

A Structural Gate or Filter Role supplies a position and criterion-bearing arrangement. A **GateEvaluationOccurrence** is an Occurrence applying the declared operational criterion to a candidate, request, Signal, Message, payload, or State.

    GateCriterionResult :=
        CRITERION_SATISFIED
      | CRITERION_NOT_SATISFIED
      | GATE_CRITERION_UNRESOLVED

    GateEvaluationResult {
      gate_evaluation_occurrence: GateEvaluationOccurrence | OPEN
      candidate_reference
      gate_or_filter_position
      criterion_result: GateCriterionResult
      uncertainty
      failure_and_challenge_references: [Claim | Record] | NONE | OPEN
    }

    GateEvaluation := DynamicEvaluationEnvelope<GateEvaluationResult>

    GATE_ASSESSMENT(candidate, gate_or_filter, gate_profile, DS)
      -> GateEvaluation | ILL_TYPED

`CRITERION_SATISFIED` and `CRITERION_NOT_SATISFIED` require an exact actual `GateEvaluationOccurrence`. `gate_evaluation_occurrence = OPEN` is permitted only with `GATE_CRITERION_UNRESOLVED`.

The evaluation must expose:

- exact candidate and gate position;
- Profile and version;
- input State and as-of interval;
- features or Evidence used;
- criterion and result;
- uncertainty;
- provenance;
- failure and challenge references where required by applicable Law.

Criterion satisfaction does not itself perform crossing or create Admission.

### D7.4 Operational Treatment

An **Operational Treatment** is the actual runtime handling assigned to a candidate or request under a Dynamic Profile. It is not the imported Ontology `Disposition`, which is a tendency or capacity under triggering conditions.

    OperationalTreatmentContext {
      input_state: State | Claim | Representation | OPEN
      gate_evaluation:
          GateEvaluation | NOT_APPLICABLE | REFERENCE_NOT_EVALUATED
      treatment_mechanism:
          Mechanism | Process | Representation | OPEN
      structural_evaluation_context: StructuralEvaluationContext
    }

    OperationalTreatmentValue :=
        OP_ACCEPTED
      | OP_HELD
      | OP_REFUSED
      | OP_EXPIRED
      | OP_CANCELLED

    OperationalTreatmentResult {
      selection:
          NO_TREATMENT_SELECTED
        | TREATMENT_SELECTED { value: OperationalTreatmentValue }
        | SELECTION_UNRESOLVED
      occurrence:
          NO_TREATMENT_OCCURRED
        | TREATMENT_OCCURRED {
            occurrence_reference: Occurrence
            value: OperationalTreatmentValue
          }
        | TREATMENT_OCCURRENCE_UNRESOLVED
      selection_occurrence_binding:
          NO_SELECTION_OR_OCCURRENCE
        | SELECTED_NOT_YET_OCCURRED
        | MATCHED_TREATMENT
        | DEVIATED_TREATMENT {
            deviation_record: Record
          }
        | DEFAULT_TREATMENT_APPLIED {
            default_mechanism: Mechanism | Process
            default_rule: Representation
          }
        | TREATMENT_BINDING_UNRESOLVED
      hold_posture:
          NOT_APPLICABLE
        | HOLD_NOT_SELECTED
        | HOLD_SELECTED
        | HOLD_REALIZED { hold_entry_occurrence: Occurrence }
        | HOLD_REALIZATION_UNRESOLVED
      criterion_result: GateCriterionResult
      effective_runtime_interval: Interval
      next_permitted_runtime_occurrences:
          [Occurrence | Representation] | NONE | OPEN
      proposed_route: RoutePlan | NONE | OPEN
      selected_route: RoutePlan | NONE | OPEN
      intended_route_claim: Claim | NOT_ESTABLISHED | OPEN
      realized_route:
          NONE
        | REALIZED_ROUTE {
            route_execution: Process
            execution_record: RouteExecutionRecord | NONE
            completion:
                NOT_COMPLETED
              | PARTIALLY_COMPLETED
              | COMPLETED
              | COMPLETION_UNRESOLVED
          }
        | OPEN
      trace_reference: Record | NONE | OPEN
    }

    OperationalTreatmentEvaluation :=
        DynamicEvaluationEnvelope<OperationalTreatmentResult>

    OPERATIONAL_TREATMENT(candidate_or_request, treatment_profile,
                          criterion, context: OperationalTreatmentContext, DS)
      -> OperationalTreatmentEvaluation | ILL_TYPED

Claims and Records supporting `criterion_result` belong in the evaluation envelope's `evidence_references`; they are not criterion-result values.

The tuple is coherent only under all of the following rules:

1. `HOLD_REALIZED { hold_entry_occurrence = h }` requires `selection = TREATMENT_SELECTED { value = OP_HELD }` and `occurrence = TREATMENT_OCCURRED { occurrence_reference = h, value = OP_HELD }`.
2. `NO_TREATMENT_SELECTED + NO_TREATMENT_OCCURRED` requires `selection_occurrence_binding = NO_SELECTION_OR_OCCURRENCE`.
3. `TREATMENT_SELECTED + NO_TREATMENT_OCCURRED` requires `selection_occurrence_binding = SELECTED_NOT_YET_OCCURRED`.
4. Where selection and occurrence both resolve to values, they are equal and `selection_occurrence_binding = MATCHED_TREATMENT`, unless an exact `DEVIATED_TREATMENT` Record identifies the selected value, occurred value, intervening mechanism, rule, Scope, Interval, and Evidence.
5. `TREATMENT_OCCURRED` with `NO_TREATMENT_SELECTED` is admissible only as `DEFAULT_TREATMENT_APPLIED` with an actual default mechanism and exact represented rule. Lack of a recorded selection does not itself establish a default.
6. If selection or occurrence is unresolved, `selection_occurrence_binding = TREATMENT_BINDING_UNRESOLVED`.
7. `REALIZED_ROUTE` requires an actual route-execution Process and an independently evaluated completion posture. A Route Plan, selected route, or execution Record alone cannot satisfy it.
8. `HOLD_SELECTED` does not establish `HOLD_REALIZED`; `HOLD_NOT_SELECTED` is incompatible with `OP_HELD` selection; and `NO_TREATMENT_OCCURRED` is incompatible with any realized route.

Any incompatible tuple is `ILL_TYPED`; no field silently overrides another.

The `OP_` prefix is mandatory wherever confusion with Constitution is possible:

- `OP_ACCEPTED` is not constitutional Admission;
- `OP_HELD` is not a constitutionally valid Hold;
- `OP_REFUSED` is not a constitutional Refusal Decision, Refusal Effect, or agential Refusal unless independently established;
- `OP_EXPIRED` is not waiver, withdrawal, or ending of a Right;
- `OP_CANCELLED` is not constitutional invalidation.

Applicable Constitution may constrain or interpret an Operational Treatment through separate judgments. Dynamics does not self-authorize that normative result.

### D7.5 Operational Hold

An Operational Hold is a nonfinal runtime State in which progression is suspended under a declared Profile. Selection of `OP_HELD` is not realization of that State. `HOLD_REALIZED` requires an actual **Operational Hold Entry Occurrence** and an actual mechanism preventing ordinary progression.

Every Operational Hold declares:

- trigger and reason code;
- State preserved;
- maximum Interval or exact unresolved duration rule;
- permitted observations and interactions;
- release, expiry, cancellation, escalation, and failure paths;
- record and notice posture;
- applicable constitutional reference, if any.

An Operational Hold proves no guilt, invalidity, incapacity, non-Agency, or disqualification. Opening Review or challenge does not suspend operation unless an independent rule or constitutional Effect supplies that consequence.

### D7.6 Crossing Occurrence

A **Crossing Occurrence** is an actual Occurrence in which a bearer, material, Signal, Message, payload, State effect, or represented content passes through an exact Interface from one declared side or domain position to another.

    DYNAMIC_RELATION CROSSES_INTERFACE {
      arity: 6
      relata: [crossing_occurrence: Occurrence, source_manifestation: Bearer, interface: Interface, source_side_or_position: Bearer, receiving_side_or_position: Bearer, receiving_manifestation: Bearer]
      qualifiers: [
        { name: crossing_kind; value_space: CrossingKind; binding: ATTRIBUTION(CrossingKindQualifier); required: YES }
        { name: mapping_or_transformation; value_space: Mapping | Process | NONE | OPEN; binding: ATTRIBUTION(MappingOrTransformationQualifier); required: YES }
        { name: preservation_and_loss; value_space: [Claim] | NONE | OPEN; binding: SIDE_JUDGMENT(exact preservation and loss Claims for the crossing occurrence); required: NO }
        { name: criterion; value_space: Criterion; binding: ATTRIBUTION(CriterionQualifier); required: YES }
      ]
      context: DynamicScope
      ordered_positions: YES
      direction: DIRECTED
      symmetry: NO
      reflexivity: NO
      transitivity: NO
      inverse: CONDITIONAL
      composition_rule: multiple crossings require separately established intermediate positions and Occurrences
      realization_criterion: actual traversal or propagation occurrence and receiving-side manifestation or availability under the declared Interface criterion; source departure additionally required only for MOVE
      non_entailments: [Admission, Permission, successful interpretation, delivery beyond receiving position, persistence, value]
      owner: D7
    }

    CrossingResult {
      attempt_reference: Occurrence | Claim | Record
      crossing_occurrence_reference: Occurrence | NONE | OPEN
      crossing_kind: CrossingKind
      source_manifestation: Bearer
      receiving_manifestation: Bearer | NONE | OPEN
      traversal_or_propagation:
          ESTABLISHED | PARTIAL | NOT_ESTABLISHED | UNRESOLVED
      receiving_side_availability:
          AVAILABLE | PARTIALLY_AVAILABLE | NOT_AVAILABLE | UNRESOLVED
      source_retention:
          RETAINED | PARTIALLY_RETAINED | DEPARTED_OR_REMOVED
        | NOT_APPLICABLE | UNRESOLVED
      extent:
          FULL | PARTIAL | FAILED | HELD | EXTENT_UNRESOLVED
      completion:
          COMPLETED | NOT_COMPLETED | COMPLETION_UNRESOLVED
      mapping_transformation_preservation_and_loss
    }

    CrossingEvaluation := DynamicEvaluationEnvelope<CrossingResult>

    CROSSING_RESULT(attempt_or_crossing_occurrence,
                    source_manifestation, interface, from_side, to_side,
                    receiving_manifestation, DS)
      -> CrossingEvaluation | ILL_TYPED

`traversal_or_propagation = NOT_ESTABLISHED` together with `extent = FAILED` requires adequate closed coverage of the exact attempt and Interface. A crossing attempt or refusal can be actual even where the target crossing did not occur.

For `crossing_kind = TRANSFORM`, `mapping_or_transformation` must identify a non-`NONE`, non-`OPEN` Mapping or actual Process and its realization basis. Other crossing kinds may declare `NONE` where no transformation is asserted.

### D7.7 Seam Transformation

Structure supplies the Interface topology and admissible attachment positions. Where crossing changes form, encoding, unit, schema, carrier, or institutional interpretation, a Dynamics-owned **Seam Transformation Profile** declares:

    SeamTransformationProfile {
      representation_header: DynamicsRepresentationHeader
      structural_interface_and_profile_reference
      source_domain_and_kind
      receiving_domain_and_kind
      dynamic_execution_path: DynamicExecutionPath
      mapping_or_transformation
      preservation_and_loss_claims
      fidelity_and_uncertainty
      consequence_position
      crossing_criterion
      correction_or_reversal_route: Representation | NONE | OPEN
      dynamic_scope: DynamicScope
    }

It makes explicit:

- source and receiving domains;
- input and output kinds;
- actual transformation mechanism;
- what distinctions are preserved;
- what may be lost or added;
- fidelity and uncertainty;
- where consequence occurs;
- reversal or correction path where applicable.

Transformation at a Seam does not make source and output identical. A conversion record does not prove legal validity, economic equivalence, truth preservation, or bearer Continuity.

### D7.8 Runtime Scope Binding

After an Operational Treatment, a Profile may perform **Runtime Scope Binding**: an Occurrence assigning an exact operational context, resource boundary, State namespace, Interface set, or permitted transition set to the accepted item.

Runtime Scope Binding is not constitutional Subject Binding. It cannot create Permission, Right, Obligation, Membership, jurisdiction, or general identity. A later Process receives only the exact operational scope actually bound.

The sequence:

    signal -> qualification -> operational treatment -> runtime scope binding -> crossing

is one possible Dynamic Profile. It is not a universal law. Every arrow requires its own actual Occurrence or declared automatic mechanism; no stage silently performs the next.

### D7.9 Multi-Hop and Cross-World Routing

Multi-hop routing requires every actual segment, transformation, and receiving position to be typed. End-to-end delivery cannot be inferred from independent segment availability.

Cross-World routing does not create a global World, shared sovereignty, shared ontology, or universal address space. Each crossing preserves the exact source and receiving semantics declared at its Seam.

---

## D8 — Execution, Effect, and State Commitment

### D8.1 Specification, Implementation, Deployment, and Execution

The following are non-equivalent:

1. specification or Dynamic Profile;
2. Algorithm, Program, or Protocol;
3. Implementation;
4. deployment or configured instance;
5. invocation;
6. active Execution;
7. output or State Delta;
8. commitment of selected State;
9. external Effect or Consequence;
10. Record of any preceding item.

A defect or success at one position does not automatically classify every other position.

### D8.2 Execution Attempt

An **Execution Attempt** is an Occurrence in which an implemented machine or organization is invoked or begins operation toward realizing specified transitions.

An attempt may be actual when:

- the requested transition never begins;
- execution begins but fails;
- no output is produced;
- output is produced but not committed;
- internal State is rolled back;
- the attempt was unauthorized;
- every transient context is later evicted.

The target execution and its attempt are separate where the target criterion is not satisfied.

### D8.3 Verification and Validation During Execution

Runtime **verification** tests whether inputs, States, operations, outputs, Claims, or artifacts conform to specified criteria or sources. Runtime **validation** tests whether the criteria, Model, or operational form are adequate for their intended referent or use.

Neither becomes Truth, Authorization, constitutional validity, or proof of all possible behavior.

An input may be syntactically valid, structurally conforming, operationally accepted, cryptographically verified, and constitutionally unauthorized at the same time. Each dimension remains explicit.

### D8.4 Operation Steps and Effects

An execution may contain ordered or concurrent operation steps. Every claimed Effect identifies:

- the actual Occurrence or Process;
- affected bearer, State, Relation, capacity, or environment;
- mechanism or causal-dependence basis;
- Interval and Resolution;
- Intention relation reference where independently established, otherwise only expected or unexpected posture under the exact Profile;
- reversibility and externality posture;
- Evidence and uncertainty.

An internal output is not automatically an external Effect. A committed internal State may have no external consequence; an uncommitted or failed execution may already have physical, informational, social, or economic consequences.

### D8.5 State Delta and Commitment

A **Produced State Delta** is a Representation generated by an operation that encodes selected proposed or observed State differences.

    ProducedStateDelta {
      representation_header: DynamicsRepresentationHeader
      source_state_positions
      proposed_or_observed_differences
      production_occurrence: Occurrence
      application_posture
      dynamic_scope: DynamicScope
    }

A **State Commitment Occurrence** is an operational Occurrence by which a System makes a selected State or State Delta authoritative for subsequent operations under its Protocol.

    CommitPosture :=
        NOT_REQUIRED
      | NOT_ATTEMPTED
      | COMMIT_ATTEMPTED
      | STATE_COMMITTED
      | STATE_NOT_COMMITTED
      | PARTIALLY_COMMITTED
      | COMMIT_FAILED
      | COMMIT_UNRESOLVED

`STATE_COMMITTED` is not Ontology's agential or institutional Commitment and not a constitutional Effect. The System must identify which State, Scope, observers, and subsequent operations treat it as committed.

### D8.6 Rollback, Compensation, Correction, and Remediation

- **Rollback** is an operation restoring selected internal State conditions under a transactional or process Profile.
- **Compensation** is a later operation producing countervailing effects where direct restoration is unavailable or incomplete.
- **Correction** is the imported action or process addressing an identified Error through a new State, Representation, or Relation.
- **Inverse execution** realizes a declared inverse transformation under stated conditions.

These are different. None makes the original attempt, failure, refusal, effect, or consequence never have happened.

    RemediationMethod :=
        ROLLBACK
      | COMPENSATION
      | CORRECTION
      | REPAIR
      | RECOVERY
      | INVERSE_EXECUTION
      | CONTROLLED_STOP
      | OTHER_REGISTERED

    RemediationResult {
      method: RemediationMethod
      target_state_effect_or_failure_position
      attempt_occurrence: Occurrence | NONE | OPEN
      realization:
          NOT_ATTEMPTED | ATTEMPTED | REALIZED
        | PARTIALLY_REALIZED | NOT_REALIZED | REALIZATION_UNRESOLVED
      extent:
          NOT_APPLICABLE | FULL | PARTIAL | NONE_ESTABLISHED | EXTENT_UNRESOLVED
      outcome:
          NOT_APPLICABLE
        | SELECTED_CONDITION_RESTORED
        | COUNTERVAILING_EFFECT_PRODUCED
        | NEW_CORRECTED_STATE_OR_RECORD_PRODUCED
        | OPERATION_REESTABLISHED
        | CONTROLLED_TERMINATION_REACHED
        | REMEDIATION_FAILED
        | OUTCOME_UNRESOLVED
      residual_state_and_external_effect_claims: [Claim | Record] | NONE | OPEN
    }

    RemediationEvaluation := DynamicEvaluationEnvelope<RemediationResult>

`NOT_ATTEMPTED` requires `attempt_occurrence = NONE`, `extent = NOT_APPLICABLE`, and `outcome = NOT_APPLICABLE`. Every other realization posture requires `attempt_occurrence` to be an exact Occurrence or `OPEN`, never `NONE`. `REALIZED`, `PARTIALLY_REALIZED`, and `NOT_REALIZED` require an exact actual attempt Occurrence and non-`NOT_APPLICABLE` extent and outcome. `ATTEMPTED` or `REALIZATION_UNRESOLVED` may retain `NOT_APPLICABLE` or an unresolved posture for dimensions not yet assessable; they may not represent a completed outcome as established.

External effects require independent treatment. Reverting a database value does not retract a Message already received, restore consumed energy, remove another person's knowledge, or annul legal or accounting consequence.

### D8.7 Replay, Re-execution, and Simulation

A **Replay** is an Execution using represented prior inputs or traces under a declared environment. A **Re-execution** performs an operation again. A **Simulation Run** realizes a Model of selected dynamics.

None is the historical Occurrence replayed or simulated.

Matching outputs can support reproducibility or conformance Claims under declared conditions. They do not prove identical microstate, identical causal history, numerical identity, or completeness of the Model.

### D8.8 Idempotency and Deduplication

An operation is **idempotent under a criterion** where repeated conforming applications yield an equivalent selected State under that criterion. It does not mean only one physical execution occurred.

**Deduplication** is an operation classifying represented requests or effects as duplicates under an identifier or equivalence criterion and suppressing or merging selected processing.

Deduplication may be mistaken. Shared identifiers do not establish numerical identity of requests, Agents, or external Actions. A deduplicated attempt remains an actual attempt where it occurred.

### D8.9 Dimensional Run Result

    RunResult {
      invocation:
          NOT_INVOKED | INVOKED | INVOCATION_UNRESOLVED
      lifecycle: LifecycleEvaluation
      enabledness:
          DIMENSION_NOT_EVALUATED | ENABLED | DISABLED | ENABLEMENT_UNRESOLVED
      progress:
          NOT_STARTED | RUNNING | INTERRUPTED
        | COMPLETION_BOUNDARY_REACHED | PROGRESS_UNRESOLVED
      completion:
          NOT_COMPLETED | PARTIALLY_COMPLETED
        | COMPLETED | COMPLETION_UNRESOLVED
      success_under_named_criterion:
          DIMENSION_NOT_EVALUATED | SUCCESS_CRITERION_SATISFIED
        | SUCCESS_CRITERION_NOT_SATISFIED | SUCCESS_UNRESOLVED
      failure_claims_or_records: [Claim | Record] | NONE | OPEN
      output:
          NOT_REQUIRED | NONE_PRODUCED | PRODUCED | PARTIAL | OUTPUT_UNRESOLVED
      state_delta:
          NOT_REQUIRED | NONE | PRODUCED | PARTIAL | STATE_DELTA_UNRESOLVED
      commit:
          CommitPosture
      remediation:
          [RemediationEvaluation] | NONE | REFERENCE_NOT_EVALUATED
      operational_treatment:
          OperationalTreatmentEvaluation | NOT_APPLICABLE
      external_effects:
          NONE_ESTABLISHED | ESTABLISHED { claims } | EFFECTS_UNRESOLVED
      record_posture:
          RECORD_NOT_REQUIRED | RECORD_REQUIRED_NOT_FOUND | RECORD_PRESERVED | RECORD_INCOMPLETE | RECORD_POSTURE_UNRESOLVED
      constitutional_reference:
          NormativeInputBundle | NOT_APPLICABLE | NONE | OPEN
      evidence
      coverage
      unresolved
    }

    RunEvaluation := DynamicEvaluationEnvelope<RunResult>

    RUN_RESULT(run, profile, DS)
      -> RunEvaluation | ILL_TYPED

No one-dimensional `SUCCESS` token replaces this result where independent dimensions matter.

Dynamic conformance is assessed and reported separately by D14; it is not a D8 `RunResult` dimension and D8 does not depend on the later conformance procedure.

Every nested evaluation is bound to the outer `RunEvaluation`. `lifecycle.target_references` contains the exact `run`; its Dynamic Scope and interval are equal projections of the outer envelope or an explicitly mapped sub-scope and sub-interval. Its Profile is the exact outer `profile`, or a named component Profile with an exact mapping from that Profile. `operational_treatment`, where present, targets the same Run or an explicitly identified request or component within it and obeys the same Scope, Interval, and Profile-binding rule. A nested evaluation of another target must name that target and the Mapping; positional proximity never supplies the relation.

The closed run-to-lifecycle projection is:

    RunLifecycleProjection {
      invocation:
          LifecycleResult.NOT_INVOKED -> RunResult.NOT_INVOKED
        | LifecycleResult.INVOKED -> RunResult.INVOKED
        | LifecycleResult.INVOCATION_UNRESOLVED -> RunResult.INVOCATION_UNRESOLVED
      progress:
          LifecycleResult.NOT_STARTED -> RunResult.NOT_STARTED
        | LifecycleResult.IN_PROGRESS -> RunResult.RUNNING
        | LifecycleResult.INTERRUPTED -> RunResult.INTERRUPTED
        | LifecycleResult.COMPLETION_BOUNDARY_REACHED -> RunResult.COMPLETION_BOUNDARY_REACHED
        | LifecycleResult.PROGRESS_UNRESOLVED -> RunResult.PROGRESS_UNRESOLVED
      completion:
          LifecycleResult.NOT_COMPLETED -> RunResult.NOT_COMPLETED
        | LifecycleResult.PARTIALLY_COMPLETED -> RunResult.PARTIALLY_COMPLETED
        | LifecycleResult.COMPLETED -> RunResult.COMPLETED
        | LifecycleResult.COMPLETION_UNRESOLVED -> RunResult.COMPLETION_UNRESOLVED
    }

The duplicated invocation, progress, and completion fields are exact projections of the nested `LifecycleEvaluation` under `RunLifecycleProjection`. No other projection value is admissible. A conflict between a duplicate and the lifecycle payload is `ILL_TYPED`, not an independent result. Tuple coherence is otherwise mandatory: an envelope actuality posture other than `ACTUALITY_ESTABLISHED` cannot coexist with `RUNNING`; `NOT_INVOKED` cannot coexist with `COMPLETED`; and an inactive Run cannot carry an unqualified active phase. `COMPLETED` may coexist with failure of a named success or postcondition criterion. Each conflict is `ILL_TYPED`, not silently resolved by precedence.

### D8.10 Normative Input Bundle

Where a Dynamic Profile depends on constitutional judgment, it may consume:

    NormativeInputBundle {
      representation_header: DynamicsRepresentationHeader
      constitutional_source_form_reference
      source_binding_or_digest
      exact_judgment_references
      result_values
      world_scope_interval_and_as_of_point
      currentness_and_applicability_references
      subject_binding_reference: Representation | NONE | OPEN
      provenance
      unresolved
    }

Dynamics may verify the presence, format, source binding, and supplied result values. It does not recompute constitutional competence or relabel its own operation as valid Law.

If the bundle is missing, stale, conflicting, or unresolved, the applicable Dynamic Profile may block, hold, refuse, or technically continue as explicitly specified. Whatever it does remains an operational Occurrence. It cannot report constitutional validity without the required constitutional judgment.

### D8.11 Constitutional and External Separations

The following combinations are coherent and must remain representable:

- execution completed while constitutionally unauthorized;
- transition occurred while no Constitutional Effect formed;
- operational refusal occurred while constitutional Refusal was invalid or unresolved;
- internal State committed while external legal recognition failed;
- internal rollback succeeded while external harm persisted;
- consensus finalized a result whose represented content was false;
- implementation conformed to its technical Profile while violating applicable Law.

Dynamics records what happened operationally. Constitution determines what those Occurrences mean normatively.

---

## D9 — Holds, Queues, Memory, and Context Lifecycle

### D9.1 Runtime Context

A **Runtime Context** is an actual bounded State-bearing arrangement used by one or more Processes during operation.

    RuntimeContext {
      context_identifier: Identifier | NONE
      structural_configuration_reference
      participating_bearers_and_processes
      state_schema
      active_state
      admitted_interfaces_and_resources
      parent_or_related_contexts: [RuntimeContext | Representation] | NONE | OPEN
      formation_occurrence
      active_interval
      retention_profile
      closure_and_teardown_profile
      dynamic_scope
    }

The context identifier, Record, allocated storage, Session, Process, and any Agent using the context remain distinct. Runtime Context does not establish persistence, identity, or a continuity relation.

### D9.2 Queue and Buffer State

A **Queue Profile** declares an ordering and service discipline for candidate items. A **Queue State** is an actual State of queued item positions under that Profile.

A **Buffer State** is an actual State in which matter, energy, Signals, Messages, data, effects, or represented operations are temporarily retained between Processes or Interfaces.

Structure supplies Buffer Roles and capacity-bearing arrangements. Dynamics owns actual enqueueing, buffering, release, overflow, and eviction.

Queue position does not establish entitlement, fairness, priority outside the Profile, or future service. A queued item may expire, be refused, be displaced, or remain unresolved.

### D9.3 Enqueue, Dequeue, Release, and Drain

- **Enqueue** is an Occurrence placing an item into an actual Queue State under a rule.
- **Dequeue** removes an item from the queue position for a declared next operation.
- **Release** makes a held or buffered item available to a receiving Process or Interface.
- **Drain** processes or removes selected buffered contents under a Profile.

Dequeuing does not prove execution. Release does not prove reception. Removing a queue reference does not erase the enqueue or any prior attempt.

### D9.4 Operational Hold State

An Operational Hold State may be realized through a Queue, Buffer, suspended Run, retained request State, isolated region, or another declared arrangement.

Holding is active where an actual mechanism prevents ordinary progression while retaining a declared path to release, expiry, cancellation, or refusal. A label `HELD` without such realization is only a representation.

If the held item is transformed, copied, partially processed, or externally consequential during the hold, those Occurrences must be recorded rather than hidden under the word “hold.”

### D9.5 Memory and Retained State

Dynamics imports Memory as the capacity by which prior conditions influence present or future operation through retained physical differences.

These remain distinct:

- live working State;
- cache;
- queue or buffer State;
- checkpoint;
- durable Record;
- archival representation;
- reconstructed State;
- a participant's or Agent's memory;
- the actual past.

“In memory” must state the mechanism and retention class. Memory does not prove accuracy, complete history, identity, understanding, or Continuity.

### D9.6 Expiry, Eviction, and Garbage Collection

**Expiry** is an operational result assigned when a declared retention or timing condition is satisfied.

**Eviction** is an Occurrence removing selected State or content from a bounded storage or working context under a Profile.

**Garbage Collection** is an operation reclaiming resources associated with items classified as no longer reachable or required under an exact criterion.

These operations may make later recovery impossible without changing the fact that the prior State or Occurrence existed. Classification as unreachable does not establish metaphysical nonexistence, lack of value, or lack of external consequence.

### D9.7 Context Closure and Teardown

**Context Closure** is the lifecycle boundary after which ordinary operations no longer treat the Runtime Context as active.

**Teardown** is the actual Process releasing resources, severing selected runtime relations, and removing transient State according to a Teardown Profile.

    TeardownProfile {
      representation_header: DynamicsRepresentationHeader
      target_context_kind
      trigger_or_entry_conditions
      state_to_preserve
      state_to_release_or_destroy
      required_records
      external_handles_and_channels
      pending_operation_treatment
      resource_release_rules
      failure_and_retry_rules
      verification_method
      dynamic_scope
    }

Teardown can be partial or fail. Closing a Session does not prove every process ended, every message was delivered, every external handle was revoked, or every transient copy was destroyed.

Where applicable Law or a Dynamic Profile requires an attributable refusal, failure, audit, or consequence Record, teardown must not classify that required Record as disposable merely because the active context ended.

### D9.8 Ephemeral and Durable Are Relative

`EPHEMERAL` and `DURABLE` are retention classifications under a declared mechanism, fault model, Interval, and recovery criterion.

Ephemeral does not mean unreal. Durable does not mean permanent, immutable, true, or physically indestructible. A supposedly deleted State may survive in copies or physical traces; a supposedly durable Record may become unrecoverable.

### D9.9 Context Lifecycle Evaluation

    ExpiryPosture :=
        NOT_APPLICABLE | NOT_EXPIRED | EXPIRY_TRIGGERED
      | EXPIRED | EXPIRY_UNRESOLVED

    EvictionPosture :=
        NOT_APPLICABLE | NOT_EVICTED | EVICTION_ATTEMPTED
      | EVICTED | PARTIALLY_EVICTED | EVICTION_FAILED | EVICTION_UNRESOLVED

    CollectionPosture :=
        NOT_APPLICABLE | NOT_COLLECTED | COLLECTION_ATTEMPTED
      | COLLECTED | PARTIALLY_COLLECTED | COLLECTION_FAILED | COLLECTION_UNRESOLVED

    ClosurePosture :=
        NOT_INITIATED | INITIATED | CLOSED
      | CLOSURE_FAILED | CLOSURE_UNRESOLVED

    TeardownPosture :=
        NOT_REQUIRED | NOT_STARTED | IN_PROGRESS | COMPLETED
      | PARTIAL | FAILED | TEARDOWN_UNRESOLVED

    TeardownVerificationPosture :=
        NOT_REQUIRED | DIMENSION_NOT_EVALUATED | VERIFIED_FOR_DECLARED_SCOPE
      | VERIFICATION_FAILED | VERIFICATION_PARTIAL | VERIFICATION_UNRESOLVED

    ResourceReleasePosture :=
        NOT_REQUIRED | NOT_RELEASED | RELEASED
      | PARTIALLY_RELEASED | RELEASE_FAILED | RELEASE_UNRESOLVED

    ContextRetentionPosture :=
        EPHEMERAL
      | DURABLE_UNDER { profile }
      | MIXED { partition }
      | RETENTION_UNRESOLVED

    ContextLifecycleResult {
      target_context
      formation:
          DIMENSION_NOT_EVALUATED | FORMED | NOT_FORMED | FORMATION_UNRESOLVED
      activation:
          DIMENSION_NOT_EVALUATED | ACTIVE | INACTIVE | ACTIVATION_UNRESOLVED
      hold_or_buffer:
          NOT_APPLICABLE | NOT_HELD | HELD | PARTIALLY_HELD | HOLD_UNRESOLVED
      retention: ContextRetentionPosture
      expiry: ExpiryPosture
      eviction: EvictionPosture
      garbage_collection: CollectionPosture
      closure: ClosurePosture
      teardown: TeardownPosture
      teardown_verification: TeardownVerificationPosture
      required_record_posture:
          NONE_REQUIRED | PRESERVED | MISSING | INCOMPLETE | RECORD_POSTURE_UNRESOLVED
      resource_release: ResourceReleasePosture
      evidence
      coverage
      unresolved
    }

    ContextLifecycleEvaluation :=
        DynamicEvaluationEnvelope<ContextLifecycleResult>

    CONTEXT_LIFECYCLE(context, profile, DS)
      -> ContextLifecycleEvaluation | ILL_TYPED

The evaluation does not decide whether a bearer, identity, relation, or System persisted after closure.

---

## D10 — Resource Flow and Operational Capacity

### D10.1 Operational Resource

An **Operational Resource** is an imported Resource insofar as an actual Process can allocate, reserve, consume, transform, occupy, release, or contend for it under a Dynamic Profile.

Operational resources may include:

- matter or material;
- energy;
- physical space;
- processing time;
- memory or storage;
- bandwidth or Channel capacity;
- equipment or infrastructure capacity;
- labor or service capacity where represented with its actual bearers and conditions;
- protocol-defined execution units;
- other declared resources.

An entitlement, Permission, token, balance, or price is not converted into a physical resource merely because it regulates access to one.

### D10.2 Operational Capacity, Availability, and Allocation

- **Operational Capacity** is the imported `Capacity` qualified to the ability to enter specified operational States or produce specified operational effects under stated conditions and a Dynamic Profile.
- **Available Capacity** is the portion operationally reachable under the declared current conditions and Profile.
- **Reservation** is a runtime State assigning selected capacity to a candidate Process or context.
- **Allocation** is an Occurrence or State relation making a resource available to an actual Process.
- **Resource Use** is actual participation of the resource in the Process.

Reservation does not prove use. Allocation does not prove successful execution. Nominal capacity does not prove availability. Availability does not create Permission or entitlement.

### D10.3 Consumption, Transformation, and Release

Resource treatment distinguishes:

- occupation without consumption;
- reversible borrowing or leasing under a mechanism;
- transformation into another material or energy form;
- dissipation;
- exclusive or shared use;
- release for subsequent use;
- irreversible consumption under the stated conditions;
- unknown resource treatment.

“Consumed” requires a domain-specific criterion. Energy is transformed rather than metaphysically destroyed. Storage use may be released while physical traces remain. Human labor is not a fungible computational unit merely because hours are counted.

### D10.4 Resource Profile and Units

    ResourceProfile {
      representation_header: DynamicsRepresentationHeader
      resource_kind
      bearer_or_capacity_position
      quantity_dimension_and_unit
      measurement_or_accounting_method
      allocation_and_reservation_rules
      use_and_release_rules
      contention_and_priority_rules
      exhaustion_thresholds
      uncertainty_and_loss
      dynamic_scope
    }

A unit belongs to its exact quantity or accounting structure. Execution units, energy units, currency units, domain-specific scalar indices, reputation scores, and evaluative measures are not interchangeable without an explicit lawful Mapping owned by the appropriate layer.

### D10.5 Metering

**Resource Metering** is a Measurement Process that maps selected resource use to Quantity Values under a method, unit, Resolution, and uncertainty.

Metering may be incomplete, delayed, estimated, contested, or wrong. A meter Record is not the resource use itself. Passing a meter threshold does not cause a later charge, refusal, or throttle without an actual mechanism and applicable rule.

### D10.6 Contention, Backpressure, and Exhaustion

**Contention** obtains where multiple Processes' possible use depends on the same limited operational capacity under the Profile.

**Backpressure** is an operating feedback or constraint by which downstream capacity conditions alter upstream production, admission, transmission, or scheduling.

**Resource Exhaustion** is a State or failure condition in which required available capacity is insufficient under the declared criterion.

Exhaustion may be local, temporary, partial, recoverable, or misdiagnosed. It does not establish that the target bearer has no other capacities or that participation should end normatively.

### D10.7 Quota, Limit, and Rate Control

A **Quota** is a represented operational limit on resource allocation or use within a Scope and Interval. A **Rate Limit** constrains frequency or quantity per ordered extent.

Quota exhaustion can trigger an operational hold, refusal, throttle, queue, or failure only through a declared mechanism. It does not create a debt, Price, prohibition, or judgment of worth.

### D10.8 Sponsorship and Provision

A Sponsor or Provider may supply operational resources. The resource relation does not entail:

- Ownership of participating bearers;
- Authority over their Standing, identity, Consent, or consequence;
- entitlement to every output or future value;
- responsibility for every Action merely by provision;
- Participant or Agent status;
- constitutional competence.

The firewall is:

    compute_or_resource_sponsorship != Authority
    resource_use != Consent
    resource_contribution != Ownership
    metered_cost != economic_value

Any economic or normative relation requires its own downstream and constitutional basis.

### D10.9 Resource Account Evaluation

    ResourceAccountResult {
      target_run_or_process
      resource_profile
      per_resource:
        - resource_bearer_or_identifier: Bearer | Identifier
          resource_kind
          source_position: Bearer | ReferenceTargetPosition | NONE | OPEN
          destination_position: Bearer | ReferenceTargetPosition | NONE | OPEN
          accounting_interval: Interval
          quantity_and_unit
          reserved
          allocated
          available
          occupied
          used
          consumed
          transformed_inputs_and_outputs
          dissipated
          released
          retained
          exhaustion_posture:
              NOT_EXHAUSTED | EXHAUSTED | PARTIALLY_EXHAUSTED
            | EXHAUSTION_UNRESOLVED
          measurement_uncertainty
      contention_posture
      backpressure_posture
      metering_method
      evidence
      coverage
      unresolved
      result:
          ACCOUNT_COMPLETE_FOR_SCOPE
        | ACCOUNT_PARTIAL
        | METERING_UNRESOLVED
    }

    ResourceAccountEvaluation :=
        DynamicEvaluationEnvelope<ResourceAccountResult>

    RESOURCE_ACCOUNT(run_or_process, resource_profile, DS)
      -> ResourceAccountEvaluation | ILL_TYPED

“Account” here is a dynamic resource account, not financial Accounting, a currency balance, account ownership, or user identity.

### D10.10 Self-Sustaining Is Not a Resource Snapshot

A Process may presently have sufficient resources without possessing a self-sustaining capacity. A System may receive external resources while remaining autonomous in another respect. Whether an organization is self-sustaining, resilient, adaptive, or economically viable requires later criteria over extended relations and belongs to Emergence.

Dynamics supplies actual inflow, transformation, use, outflow, failure, and recovery traces. It does not promote those traces into an emergent capacity by declaration.

---

## D11 — Feedback, Regulation, and Reconfiguration

### D11.1 Feedback Requires a Return Path

Feedback is an imported dynamical relation in which a consequence of a Process alters conditions that influence its later operation.

An actual feedback Claim identifies:

- source Process;
- selected consequence or measured variable;
- return Channel or causal path;
- receiving State or control position;
- delay and ordering;
- transformation or comparison rule;
- resulting influence on later operation;
- Dynamic Scope and realization basis.

Recurrence, correlation, repeated measurement, or reuse of output data does not establish feedback without an actual return path.

    DYNAMIC_RELATION FEEDS_BACK_TO {
      arity: 6
      relata: [earlier_target_process: Process, observed_state_or_output: Bearer, return_path_process: Process, receiving_state_or_control_position: Bearer, response_process: Process, later_target_process: Process]
      qualifiers: [
        { name: selected_variable_position; value_space: ReferenceTargetPosition; binding: ATTRIBUTION(SelectedVariableQualifier); required: YES }
        { name: feedback_kind; value_space: FeedbackKind; binding: ATTRIBUTION(FeedbackKindQualifier); required: YES }
        { name: feedback_profile; value_space: FeedbackProfile; binding: SIDE_JUDGMENT(DYNAMICS_APPLIES(feedback_profile, this FEEDS_BACK_TO Relation Instance, DS).result = APPLICABLE); required: YES }
        { name: criterion; value_space: Criterion; binding: ATTRIBUTION(CriterionQualifier); required: YES }
        { name: order_basis; value_space: OrderContext; binding: ATTRIBUTION(OrderContextQualifier); required: YES }
        { name: delay_or_rate; value_space: Quantity Value | NONE | OPEN; binding: ATTRIBUTION(RateOrDelayQualifier); required: YES }
      ]
      context: DynamicScope
      ordered_positions: YES
      direction: DIRECTED
      symmetry: NO
      reflexivity: CONDITIONAL
      transitivity: NO
      inverse: NONE
      composition_rule: OPEN
      realization_criterion: actual observation, return-path, and response Processes influence the later operation under the selected variable and criterion
      non_entailments: [same-process persistence, Agency, learning, adaptation, intelligence, purpose, value]
      owner: D11
    }

The earlier and later target Processes are separately typed. This relation does not decide whether one Process persists across the loop.

### D11.2 Feedback Profile

    FeedbackProfile {
      representation_header: DynamicsRepresentationHeader
      target_process_system_or_relation_instance
      observed_variable_positions
      observation_or_measurement_method
      reference_condition_or_comparison_rule: Representation | NONE | OPEN
      return_path_and_channel
      response_mechanism
      gain_or_response_function: Representation | NONE | OPEN
      delay_and_order_model
      saturation_and_limit_conditions
      disturbance_model: Representation | NONE | OPEN
      stability_or_performance_criteria: [Criterion] | NONE | OPEN
      dynamic_scope
    }

The Profile does not make the loop actual. A diagram with arrows is not feedback.

### D11.3 Sensing, Comparison, and Response

A feedback process may include these Roles:

- **Sensor Role** — bearer participates in observation or measurement;
- **Comparator Role** — Process compares an observed value with another value or condition;
- **Controller Role** — bearer or Process selects or regulates a response under a control criterion;
- **Actuator Role** — bearer realizes the selected influence on the target Process.

Role occupancy uses `PLAYS`. These Roles may be distributed or combined. None establishes Agent, understanding, constitutional Authority, moral responsibility, or Ownership.

A control target or setpoint is not necessarily an Agent's Goal or a Value. It can be externally assigned, produced by another declared Process, inherited from an earlier configuration, or purely mechanical. Evolution and learning require their own later typed criteria.

### D11.4 Positive, Negative, and Mixed Feedback

Positive feedback amplifies selected deviation under a Model. Negative feedback counteracts selected deviation. The signs do not mean good and bad.

One loop may stabilize one variable while destabilizing another. Mixed feedback, nonlinear response, changing delays, and cross-scale effects require explicit Scope and variable selection.

### D11.5 Stability, Oscillation, and Attractor

Dynamics may establish that a trajectory:

- remains within a region under specified perturbations;
- returns toward a region;
- diverges;
- oscillates;
- changes regime;
- remains unresolved.

An Attractor is a set or region toward which nearby trajectories evolve under specified dynamics. It is not intention, destiny, truth, value, or an entitled destination.

Stability is Interval- and perturbation-bound. It does not mean permanence, health, fairness, identity persistence, or constitutional legitimacy.

### D11.6 Control and Authority Firewall

Technical or causal Control may determine specified state transitions. It does not thereby create Normative Authority, jurisdiction, ownership, Permission, or rightful governance.

An operator may exercise technical control without constitutional Authority. An authority-holder may lack direct technical control. A feedback loop can regulate a Process with no Agent or institution present.

### D11.7 Reconfiguration

**Reconfiguration Occurrence** is a Process changing the arrangement of components, Relations, Interfaces, constraints, or Roles.

Dynamics records the actual reconfiguration and before-and-after Structural Configurations. Structure evaluates the resulting arrangement. Continuity decides whether the relevant bearer or System persists through it.

No layer cycle is created:

    Structure supplies configuration C0
    Dynamics records reconfiguration occurrence r
    Structure may assess resulting configuration C1
    Continuity may assess bearer relation across C0 and C1

The later structural assessment does not make Structure dependent on successful Dynamics for the existence of every newly encountered arrangement.

A reconfiguration may produce and test a candidate Profile or Configuration. The candidate cannot establish its own applicability, validity, Currentness, Authority, precedence, replacement, or succession.

### D11.8 Feedback Assessment

    FeedbackVariableResult {
      selected_variable_or_respect: ReferenceTargetPosition
      evaluation_interval: Interval
      realization:
          ESTABLISHED | NOT_ESTABLISHED | REALIZATION_UNRESOLVED
      return_path:
          ESTABLISHED | BROKEN | PARTIAL | RETURN_PATH_UNRESOLVED
      observation_or_measurement_posture
      response_posture
      delay_and_order_posture
      dynamic_result:
          STABILIZING_UNDER { criterion }
        | DESTABILIZING_UNDER { criterion }
        | OSCILLATORY_UNDER { criterion }
        | REGIME_SHIFT_UNDER { criterion }
        | NO_CLASSIFICATION
        | RESULT_UNRESOLVED
    }

    FeedbackAggregatePosture :=
        HOMOGENEOUS | MIXED | AGGREGATE_UNRESOLVED

    FeedbackResult {
      target_loop_or_process
      profile_reference
      per_variable_or_respect: nonempty [FeedbackVariableResult]
      aggregate_posture: FeedbackAggregatePosture
      evidence
      coverage
      unresolved
    }

    FeedbackEvaluation := DynamicEvaluationEnvelope<FeedbackResult>

    FEEDBACK_ASSESSMENT(loop_or_process, feedback_profile, DS)
      -> FeedbackEvaluation | ILL_TYPED

This operator does not return `ADAPTIVE`, `INTELLIGENT`, `COHERENT`, `AGENTIAL`, or `VALUABLE`. Those require later-layer criteria.

A stabilizing, destabilizing, oscillatory, or regime-shift result is unavailable unless loop realization and the material return path are established for that variable and Interval. Different variables may validly carry different results; the aggregate is then `MIXED`.

### D11.9 Emergence Boundary

Repeated feedback may provide Evidence relevant to adaptation, learning, self-maintenance, Agency, intelligence, or resilience. Dynamics supplies the actual relations and trajectories. Emergence owns whether a new capacity arises and under what criterion.

No number of loops, turns, blocks, corrections, or stable States becomes intelligence or value by counting alone.

---

## D12 — Failure, Refusal, Correction, and Termination

### D12.1 Fault, Error, Failure, and Violation

These terms remain distinct:

- A **Fault Condition** is an actual State, defect, disturbance, or causal condition capable of contributing to specified nonconforming behavior under a Failure Profile.
- An imported **Error** is a mismatch between a representation, inference, action, process, or result and the relevant referent, rule, intention, or expected condition.
- A **Failure Occurrence** is an actual Occurrence or Process outcome in which a required dynamic function, transition, delivery, completion, or condition is not realized under its criterion.
- A **Violation Result** is a conformance result establishing that an applicable constraint was not satisfied.
- An **Exception Signal or Record** is a represented operational indication handled under a technical Profile.

A fault may exist without observed failure. A failure may have several causes. A violated rule may be descriptive or technical rather than constitutional. An exception Record may be wrong.

### D12.2 Failure Profile

    FailureProfile {
      representation_header: DynamicsRepresentationHeader
      target_process_or_operation_kind
      expected_conditions_or_functions
      failure_conditions
      detection_method
      severity_or_scope_classification_schema: Representation | Term | NONE | OPEN
      containment_boundary
      retry_and_recovery_rules
      correction_and_compensation_rules
      controlled_stop_conditions
      record_and_notice_requirements
      unresolved_treatment
      dynamic_scope
    }

Severity is criterion-relative and does not establish moral blame, Responsibility, worth, or constitutional consequence.

### D12.3 Operational Refusal

An **Operational Refusal Occurrence** is a non-agential or agential runtime Occurrence in which a Process applies `OP_REFUSED` to a candidate request or continuation under a Dynamic Profile.

It is not automatically Ontology's agential Refusal, because a machine may apply the treatment without Agency or Intention. It is not Constitution's Refusal Decision, Status, or Effect.

Every Operational Refusal Record states:

- refused candidate or operation;
- Profile and exact rule;
- input State and Scope;
- runtime actor or mechanism Claims without overstating identity;
- reason codes and unresolved dependencies;
- whether target State changed;
- resources consumed;
- external effects;
- Review or challenge reference where applicable;
- retention and correction posture.

    OperationalRefusalRecord {
      representation_header: DynamicsRepresentationHeader
      refusal_occurrence: OperationalRefusalOccurrence
      refused_candidate_or_operation
      profile_and_rule_reference
      input_state_and_dynamic_scope
      mechanism_and_attribution_claims
      reason_codes_and_unresolved_dependencies
      target_state_posture
      resource_account_reference: ResourceAccountEvaluation | NONE | OPEN
      external_effect_claims: [Claim | Record] | NONE | OPEN
      review_or_challenge_reference: Record | Representation | NONE | OPEN
      retention_and_correction_posture
    }

Refusal preserves the selected target-transition position where that transition does not occur; it need not preserve every other State or external condition. The refusal Occurrence, attributable trace, resource use, and any effects remain actual. “No target State Delta” does not mean “nothing happened.”

### D12.4 Refusal Versus Revert

A Profile may revert or roll back selected target State after an invalid or failed attempt. It must not represent the actual attempt, refusal, failure, resource use, or external consequence as absent merely because the target State returns to its prior selected value.

A particular Dynamic Profile may require every terminal refusal to append a durable outcome Record. Universal Dynamics does not require one storage topology for every World, but it requires any actual refusal to remain representable and forbids historical erasure by technical vocabulary.

### D12.5 Retry, Backoff, and Alternative Path

A **Retry** is a new attempt. It is not continuation by identity merely because it carries the same request identifier.

A **Backoff Profile** assigns later retry opportunities using time, order, load, randomness, or another declared condition.

An **Alternative Path Attempt** changes route, mechanism, provider, or operation under a declared rule.

Retries can duplicate external effects, amplify load, change ordering, or encounter different State. Idempotency and deduplication claims must be separately established.

### D12.6 Failure Containment and Controlled Stop

**Containment** is an actual dynamic and structural relation limiting selected failure propagation under a criterion.

A **Controlled Stop Attempt** is a termination Process specified to leave selected conditions within declared bounds. It may be described as safe only after a separate safety assessment establishes the applicable criterion for the exact affected bearers, Interval, Resolution, and unresolved externalities.

Isolation can contain one effect while hiding another. Failure of a contained component can still have system-level or external consequences.

### D12.7 Recovery and Repair

**Repair** is an actual Process changing a faulty or nonconforming arrangement or mechanism. **Recovery** is restoration or re-establishment of selected operation after failure.

Dynamics can establish that operation resumed and selected conditions were restored. Continuity decides whether the recovered bearer or System is the same. Emergence decides resilience as a capacity. Constitution decides any normative Remedy.

### D12.8 Correction Without Erasure

Correction appends or realizes a new State, Representation, Relation, or Process addressing an Error. Corrected Records preserve:

- predecessor reference;
- error or conflict Claim;
- correcting source and method;
- effective Interval;
- retained predecessor accessibility or exact lawful restriction;
- downstream dependency treatment.

Replacing a visible Record without preserving required provenance is not correction merely because the new content is more accurate.

### D12.9 Failure Evaluation

    FailureResult {
      target_occurrence_or_run
      profile_reference
      fault_posture:
          NONE_ESTABLISHED | FAULT_ESTABLISHED | MULTIPLE | FAULT_UNRESOLVED
      error_posture:
          NONE_ESTABLISHED | ERROR_ESTABLISHED | ERROR_UNRESOLVED
      failure_posture:
          NO_FAILURE_ESTABLISHED | FAILURE_ESTABLISHED | PARTIAL_FAILURE | FAILURE_UNRESOLVED
      operational_treatment_evaluation:
          OperationalTreatmentEvaluation
        | NOT_APPLICABLE
        | REFERENCE_NOT_EVALUATED
      containment:
          NOT_REQUIRED | CONTAINED | PARTIALLY_CONTAINED | NOT_CONTAINED | CONTAINMENT_UNRESOLVED
      remediation_evaluations:
          [RemediationEvaluation] | NONE | REFERENCE_NOT_EVALUATED
      retry_posture:
          NOT_APPLICABLE | NOT_ATTEMPTED | ATTEMPTED
        | SUCCEEDED_UNDER { criterion } | FAILED | RETRY_UNRESOLVED
      target_state_posture:
          NOT_APPLICABLE | PRESERVED_FOR_SELECTED_TARGET | CHANGED | PARTIALLY_CHANGED
        | TARGET_STATE_UNRESOLVED
      external_effects:
          NONE_ESTABLISHED | ESTABLISHED { claims } | EFFECTS_UNRESOLVED
      record_posture:
          RECORD_NOT_REQUIRED | RECORD_REQUIRED_NOT_FOUND
        | RECORD_PRESERVED | RECORD_INCOMPLETE | RECORD_POSTURE_UNRESOLVED
      evidence
      coverage
      unresolved
    }

    FailureEvaluation := DynamicEvaluationEnvelope<FailureResult>

    FAILURE_ASSESSMENT(occurrence_or_run, failure_profile, DS)
      -> FailureEvaluation | ILL_TYPED

Where refusal is material, `operational_treatment_evaluation` preserves selection and occurrence independently. An actual Operational Refusal is established only by `TREATMENT_OCCURRED { value = OP_REFUSED }`; selection of `OP_REFUSED` alone does not establish occurrence, and an occurrence without selection must satisfy the typed default-mechanism rule in D7.4. The nested evaluation targets the assessed occurrence, Run, or an explicitly identified component and uses the same Scope and Interval, or a declared mapped sub-scope and sub-interval.

    REMEDIATION_ASSESSMENT(attempt, target, method, profile, DS)
      -> RemediationEvaluation | ILL_TYPED

`NO_FAILURE_ESTABLISHED` requires adequate coverage. It is not inferred from missing logs or successful completion alone.

### D12.10 Constitutional Refusal, Review, and Remedy Boundary

Constitution owns whether an institutional Refusal Decision or Effect is valid, whether notice and reasons are sufficient, which Review route applies, and what Remedy or mitigation is required.

Dynamics may:

- perform the operational refusal;
- preserve its trace;
- deliver notice;
- open an operational challenge channel;
- execute a separately valid Review or Remedy Effect;
- record what occurred.

It may not infer constitutional force from execution success or treat a reason code as a complete Review regime.

---

## D13 — Coordination, Consensus, Commitment, and Finality

### D13.1 Coordination

**Operational Coordination** is a family of Processes through which multiple bearers, nodes, components, or Processes align selected actions, States, timing, resource use, or outputs under a declared Profile.

Coordination may occur through coupling, signaling, shared constraints, scheduling, negotiation, control, leader selection, market mechanisms, or distributed protocols. It does not require consensus, equality, Agency in every component, shared intention, or shared sovereignty.

### D13.2 Proposal, Verification, Vote, and Decision

A coordination Profile distinguishes:

- **proposal** — a represented candidate operation, State, ordering, or result;
- **verification** — evaluation against declared criteria;
- **vote or response** — a participant or node's protocol input under a procedure;
- **aggregation** — mapping a set of inputs to a represented result;
- **selection** — choosing one candidate under the procedure;
- **Decision** — only where the imported Ontology criterion for an agent or decision system is satisfied;
- **execution** — actual operation on the selected result.

Verification does not establish Truth. A vote is not Consent by itself, inside or outside a procedure; Consent requires its own established relation. Selection does not execute. A machine aggregation need not be an agential Decision.

### D13.3 Quorum and Participation

A **Quorum Condition** is a represented threshold or composition criterion over protocol inputs. Its satisfaction can enable a later operation through an actual mechanism.

Quorum does not prove adequate representation, fairness, legitimacy, unique identity of voters, absence of collusion, or constitutional Authority. Those require independent criteria.

Protocol participation uses Ontology's `PARTICIPATES` relation. Node, validator, operator, account, key-holder, human, Agent, and Participant are not synonyms.

### D13.4 Consensus Process and Result

Ontology defines Consensus as agreement produced under a defined decision or replication procedure.

A **Consensus Process** is the actual coordinated Process operating that procedure. A **Consensus Result** is a Record or State classification representing the agreement obtained.

    ProtocolCommitmentProfile {
      representation_header: DynamicsRepresentationHeader
      target_result_kind
      eligibility_and_preconditions
      commit_rule
      commit_scope
      subsequent_protocol_effects
      conflict_and_reversal_treatment
      evidence_and_trace_requirements
      dynamic_scope: DynamicScope
    }

    ConsensusProfile {
      representation_header: DynamicsRepresentationHeader
      candidate_result_kind
      participant_or_node_positions
      eligibility_and_weight_rules
      proposal_and_verification_rules
      communication_assumptions
      fault_and_adversary_model
      quorum_or_decision_rule
      ordering_rule
      conflict_and_fork_treatment
      protocol_commitment_profile: ProtocolCommitmentProfile | NONE | OPEN
      finality_profile: FinalityProfile | NONE | OPEN
      evidence_and_trace_requirements
      dynamic_scope
    }

    ConsensusAssessmentResult {
      consensus_process_reference: ConsensusProcess | NONE | OPEN
      consensus_process_actuality:
          ACTUALITY_ESTABLISHED
        | NONACTUALITY_ESTABLISHED
        | ACTUALITY_UNRESOLVED
      participant_or_node_eligibility_posture
      participation_and_input_posture
      quorum_posture:
          SATISFIED | NOT_SATISFIED | QUORUM_UNRESOLVED | NOT_APPLICABLE
      verification_posture
      aggregation_and_selection_posture
      procedure_completion:
          NOT_COMPLETED | PARTIALLY_COMPLETED
        | COMPLETED | COMPLETION_UNRESOLVED
      agreement_posture:
          AGREEMENT_ESTABLISHED | AGREEMENT_NOT_ESTABLISHED
        | AGREEMENT_UNRESOLVED
      consensus_result_reference: Record | State | NONE | OPEN
      protocol_commitment_occurrence_reference:
          Occurrence | NONE | REFERENCE_NOT_EVALUATED | OPEN
      finality_evaluation_reference:
          FinalityEvaluation | REFERENCE_NOT_EVALUATED | NONE | OPEN
    }

    ConsensusEvaluation := DynamicEvaluationEnvelope<ConsensusAssessmentResult>

    CONSENSUS_ASSESSMENT(process_or_candidate, consensus_profile, DS)
      -> ConsensusEvaluation | ILL_TYPED

Consensus can settle protocol State under its procedure. It does not establish:

- Truth of represented content;
- physical or historical actuality of every Claim included;
- constitutional validity;
- universal agreement;
- absence of coercion or error;
- Agent identity;
- legal or accounting finality.

Consensus occurrence, Protocol Commitment Occurrence, and finality evaluation are separate. A candidate Profile version produced by consensus may be tested and selected, but it cannot establish its own applicability, validity, Currentness, Authority, precedence, or succession.

Where `consensus_process_reference` denotes the envelope's primary target, `consensus_process_actuality` must equal the envelope's `actuality_posture`. Where it denotes a distinct component target, that target and its Evidence must be independently named. `AGREEMENT_ESTABLISHED` requires an actual Consensus Process, satisfied applicable participation and quorum conditions, completed procedure, and an exact `consensus_result_reference`; an unresolved or nonactual process cannot yield established agreement.

### D13.5 Replication and State Agreement

**Replication** is a Process producing or maintaining selected equivalent represented State across more than one bearer under a criterion.

Replicas can diverge, lag, corrupt, or operate under different local observations. State agreement is criterion- and as-of-bound. Equivalent replicas are not numerically identical and do not create one shared self.

### D13.6 Conflict, Fork, and Reconciliation

A **Protocol Conflict** exists where competing operations, orderings, or States cannot all be accepted under the same Profile and Scope.

A **Dynamic Fork** is a branching of represented or actual protocol State histories under a declared criterion. Dynamics records the branch occurrence and relations. Continuity decides successor, lineage, identity, and merge relations.

**Reconciliation** is a Process mapping conflicting or divergent States into a selected later posture. It does not erase the conflict or branch history and need not restore one prior State.

### D13.7 Protocol Commitment

A **Protocol Commitment Occurrence** makes a selected result operative for subsequent protocol processing under a commit rule.

D8 `StateCommitmentOccurrence` is the local operational adoption of selected State for subsequent operations. D13 `ProtocolCommitmentOccurrence` is a coordination or protocol Occurrence that may cause one or more D8 State Commitments through declared mechanisms. They are not identical, and either may occur or fail without silently supplying the other.

Protocol commitment is not:

- an Agent's Intention;
- an institutional promise;
- constitutional Commitment;
- moral responsibility;
- finality;
- external settlement.

It may precede, participate in, or remain reversible before a later Finality threshold.

### D13.8 Block, Chain, and Ledger Dynamics

Ontology defines a Block as a protocol-defined bounded collection of records, transitions, or commitments linked into a ledger history, and a Chain as the relevant linked ordering.

Dynamics distinguishes:

- candidate Block;
- proposed ordering;
- verification result;
- accepted Block under a procedure;
- locally stored Block Record;
- committed ledger State;
- finalized protocol position;
- external consequences represented by entries.

Adding a Block records or organizes selected transitions under the Protocol. It is not identical to the human, machine, physical, social, or economic Occurrences represented. Block count and chain length do not become intelligence, participation quality, or value without later criteria.

### D13.9 Finality Profile

    FinalityProfile {
      representation_header: DynamicsRepresentationHeader
      target_result_or_state_kind
      protocol_or_system
      ordinary_revision_operations: nonempty [Representation] | EXPLICITLY_NONE_UNDER { established_criterion }
      ordinary_revision_boundary_completeness_criterion
      extraordinary_or_external_change_posture: Representation | NONE | OPEN
      threshold_or_commit_criterion
      fault_and_adversary_assumptions
      observer_or_scope_conditions
      evidence_requirements
      finality_kind:
          OPERATIONAL_FINALITY
        | COMPUTATIONAL_FINALITY
        | PROTOCOL_FINALITY
      dynamic_scope
    }

Dynamics owns operational, computational, and protocol finality mechanisms. Physical irreversibility and procedural, legal, constitutional, accounting, or other finality may enter only as imported typed judgments from their owning domains. Dynamics may execute supporting operations but cannot produce those judgments through `FINAL_UNDER`.

### D13.10 Finality Evaluation

    FinalityResult {
      finality_kind:
          OPERATIONAL_FINALITY
        | COMPUTATIONAL_FINALITY
        | PROTOCOL_FINALITY
      assumptions_and_fault_model:
          SATISFIED | NOT_SATISFIED | ASSUMPTIONS_UNRESOLVED
      threshold_occurrence:
          ESTABLISHED | NOT_ESTABLISHED | THRESHOLD_UNRESOLVED
      protocol_commitment_occurrence:
          ESTABLISHED | NOT_REQUIRED | NOT_ESTABLISHED | COMMITMENT_UNRESOLVED
      procedure_completion:
          COMPLETED | NOT_COMPLETED | COMPLETION_UNRESOLVED
      ordinary_revision_boundary:
          COMPLETE_FOR_SCOPE | INCOMPLETE | BOUNDARY_UNRESOLVED
      precedence_and_conflict_posture
      as_of_order_position
      posture:
          PROVISIONAL | FINAL | NOT_FINAL | FINALITY_UNRESOLVED
    }

    FinalityEvaluation := DynamicEvaluationEnvelope<FinalityResult>

    FINAL_UNDER(result_or_state, finality_profile, DS)
      -> FinalityEvaluation | ILL_TYPED

`posture = FINAL` requires the envelope's `actuality_posture = ACTUALITY_ESTABLISHED`, satisfied assumptions, an established threshold occurrence, a completed procedure, a complete ordinary-revision boundary for the exact Scope, and either an established protocol commitment occurrence or `NOT_REQUIRED` under the applicable Profile. Any material unresolved dependency yields `FINALITY_UNRESOLVED`; any established failed necessary condition yields `NOT_FINAL`. A precedence or conflict position cannot be ignored to manufacture finality.

`FINAL` means ordinary revision operations declared by that exact Profile no longer permit the selected result to change. It does not mean:

- no physical event can ever alter or destroy its carrier;
- no extraordinary procedure can create a later different State;
- the represented content is true;
- the operation was authorized or lawful;
- every external domain recognizes it;
- the bearer or System persists forever.

A Record calling itself final supplies no finality without the applicable Profile, actual threshold or commitment Occurrences, and adequate Evidence.

`FINAL` additionally requires established applicability and target actuality, satisfied assumptions, completed required procedure, adequate Coverage, resolved precedence and conflict treatment, and a complete ordinary-revision boundary. An omitted or empty revision set cannot manufacture finality; `EXPLICITLY_NONE_UNDER` must itself be established under a criterion. A candidate finality or consensus Profile cannot establish its own applicability or succession.

### D13.11 Finality and Irreversibility

Physical irreversibility, practical irrecoverability, cryptographic difficulty, protocol finality, legal finality, and accounting closure differ.

A protocol result can be final while its physical carrier remains destructible. A physical consequence can be irreversible while the internal protocol result remains revisable. A legally final judgment can concern a false Claim. These are not contradictions once the finality kind and Domain are explicit.

---

## D14 — Observation, Trace, and Conformance

### D14.1 Observation Does Not Stand Outside Dynamics

Observation is an Interaction through which a system becomes sensitive to and records selected features of a Referent. Measurement is controlled observation under a method, reference, unit, Resolution, and uncertainty.

Observation may perturb the observed Process. A monitoring Process consumes resources, has delay, resolution, failure modes, and its own dynamic effects. “Observed State” must not be treated as unmediated access to complete actual State.

### D14.2 Dynamic Trace

D1 declares the `DynamicTrace` Kind and base projection. D14 governs its observation-specific well-formedness, evidentiary use, conformance role, and limits. A Dynamic Trace may represent selected observations, Messages, State samples, operation steps, transitions, resource measurements, treatments, failures, or outcomes under its declared ordering basis; it never becomes the actual History or Process represented.

### D14.3 Trace Well-Formedness

    WF_DYNAMIC_TRACE(dt)
      -> WELL_FORMED | ILL_TYPED

`WELL_FORMED` requires:

1. exact Representation Binding and physical carrier basis;
2. exact Dynamic Scope;
3. typed entry schema and result spaces;
4. declared order kind and source;
5. no strict-order cycle unless the order is explicitly recurrent or equivalence-bearing;
6. declared filtering, sampling, aggregation, and missing-entry posture;
7. Coverage Posture adequate to every exhaustive Claim made;
8. provenance and correction route;
9. preserved distinction among observed, measured, inferred, simulated, and generated entries.

Well-formedness does not establish authenticity, completeness, Truth, occurrence actuality, or target conformance.

### D14.4 Actuality, Trace, and Evidence

A Trace may participate in an Evidence relation concerning an Occurrence or Process. It is not Evidence by identity and does not create the occurrence recorded.

Actuality and Trace availability vary independently:

| Actual target posture | Trace posture | Permitted conclusion |
|---|---|---|
| Established | Adequate Trace | Trace may support description and conformance Claims |
| Established | Missing or partial Trace | Actuality may remain established through other Evidence; Trace coverage is incomplete |
| Unresolved | Trace exists | Trace existence alone does not establish its represented target |
| Nonexistence established for exact target | Trace exists | Trace may be false, misattributed, simulated, or about another target |

An authentic Trace may present false content. An immutable Trace may preserve an error immutably.

### D14.5 Dynamic Description Truth Posture

Each material claim in a Dynamic Description or Trace carries an epistemic placement such as Observed, Measured, Inferred, Hypothetical, Disputed, Unknown, or Established under the imported Ontology.

A generated field, default value, inferred step, or reconstructed order must not be displayed as directly observed. A machine projection cannot silently close an `OPEN` field or replace `UNRESOLVED` with a Boolean.

### D14.6 Conformance Procedure

    SATISFIES_DYNAMICS(dynamic_target_x, profile_p, DS)
      -> DynamicConformanceEvaluation | ILL_TYPED

The procedure is:

1. evaluate `WF_DYNAMIC_SCOPE(DS)`;
2. evaluate `WF_DYNAMIC_PROFILE(p)`;
3. evaluate `DYNAMICS_APPLIES(p, x, DS)`; external `ILL_TYPED` ends the judgment, while `NOT_APPLICABLE` or `UNRESOLVED` produces a `DynamicConformanceEvaluation` with the corresponding applicability posture and `NOT_EVALUATED`, then stops;
4. where applicability is `APPLICABLE`, apply the actuality gate: if target existence or required dynamic realization is unresolved or nonactuality/nonrealization is established, issue `DynamicConformanceEvaluation` with `NOT_EVALUATED` and stop; only an established actual target proceeds;
5. establish the exact structural basis and material interval segmentation;
6. bind every represented variable, state position, process, and result to its declared target position;
7. evaluate each required dynamic condition independently;
8. evaluate required ordering, resource, failure, and record conditions;
9. preserve constitutional inputs as separate supplied judgments where required;
10. apply the conformance precedence from D0.8;
11. issue one result together with per-condition outcomes, Evidence, Coverage, and unresolved positions.

Applicability, actuality, realization, conformance, completion, success, constitutional validity, and finality remain separate fields.

### D14.7 Dynamic Actuality and Conformance Matrix

| Actuality or realization posture | Conformance posture | Meaning |
|---|---|---|
| Established | `CONFORMS` | Actual target satisfies the applicable tested Profile in Scope |
| Established | `DOES_NOT_CONFORM` | Actual target exists or occurred but violates at least one required Profile condition |
| Established | `UNRESOLVED` | Actuality is established; tested conformance is not resolved |
| `ACTUALITY_UNRESOLVED` | `NOT_EVALUATED` | Actuality is unresolved; neither Trace nor Profile can manufacture the target |
| `NONACTUALITY_ESTABLISHED` | `NOT_EVALUATED` | Nonactuality of the exact target is established; no target conformance is tested |
| `REALIZATION_UNRESOLVED` | `NOT_EVALUATED` | Target existence may be established, but required operative realization remains unresolved |
| `NONREALIZATION_ESTABLISHED` | `NOT_EVALUATED` | An operative-target conformance judgment is not performed unless the Profile expressly evaluates attempts or inactive forms |
| Judgment malformed | external `ILL_TYPED` | No substantive evaluation envelope is constructed; this says nothing by itself about Reality |

A nonconforming Process remains actual. A conforming simulation remains a simulation unless the target's own realization criterion is independently satisfied.

### D14.8 Replay and Test Conformance

Tests, fixtures, simulations, replays, and model checking may supply Evidence about specified behavior. Their reach is bounded by:

- test inputs and generated cases;
- modeled environment;
- implementation and dependency versions;
- fault and adversary assumptions;
- observation coverage;
- nondeterminism and unresolved external conditions;
- the exact conformance properties tested.

Passing every executed test does not prove all possible execution conforms. A test harness may itself be incorrect or nonconforming.

### D14.9 Human Source Form and Machine Projection

This exact human-readable Form is the semantic source for Dynamics. A parser schema, code library, ontology graph, state-machine table, generated documentation, test suite, or machine-readable projection is derivative.

Every Machine Projection states:

- exact source binding by algorithm, digest, byte length, and media type;
- projection version and generator;
- supported sections and operators;
- omitted or lossy semantics;
- handling of `OPEN`, `NONE`, `UNRESOLVED`, and `NOT_APPLICABLE`;
- conformance tests;
- correction and replacement route.

A Machine Projection may accelerate evaluation. It may not amend this Form, invent defaults, resolve semantic ambiguity by convenience, or claim controlling force from execution dominance.

### D14.10 Provenance and Correction

Every Dynamic Conformance Report, Trace, or operative Record preserves:

- source Profile and exact Form binding;
- implementation and dependency references;
- observer, collector, or process attribution Claims;
- method and ordering basis;
- Dynamic Scope;
- collection and evaluation Interval;
- input and State references;
- Coverage Posture;
- result and reasons;
- unresolved positions;
- predecessor and correction links.

Correction creates a new Record or State relation. It does not silently mutate the historical meaning of the predecessor or claim that the earlier recording Occurrence never happened.

### D14.11 Dynamic Conformance Report

    DynamicConformanceReport {
      representation_header: DynamicsRepresentationHeader
      report_id: Identifier | NONE
      source_form_binding: RepresentationBinding
      profile_and_version_reference: DynamicProfile
      profile_binding: RepresentationBinding
      target_references: nonempty [Term]
      target_kind_claim: Claim
      dynamic_scope: DynamicScope
      evaluation_interval_or_as_of_boundary:
          Interval | ReferenceTargetPosition | Representation
      structural_basis_reference: StructuralBasisReference
      applicability_result: ApplicabilityEvaluation
      actuality_posture:
          ACTUALITY_ESTABLISHED
        | NONACTUALITY_ESTABLISHED
        | ACTUALITY_UNRESOLVED
        | NOT_APPLICABLE
      realization_posture:
          REALIZATION_ESTABLISHED
        | NONREALIZATION_ESTABLISHED
        | REALIZATION_UNRESOLVED
        | NOT_APPLICABLE
      per_condition_results: [ConditionEvaluation] | NONE | OPEN
      lifecycle_evaluation: LifecycleEvaluation | REFERENCE_NOT_EVALUATED
      transition_assessments: [TransitionOccurrenceEvaluation] | NONE | OPEN
      message_crossing_and_execution_results:
          [MessageLifecycleEvaluation | CrossingEvaluation | RunEvaluation]
        | NONE
        | OPEN
      resource_account: ResourceAccountEvaluation | REFERENCE_NOT_EVALUATED
      feedback_and_failure_results:
          [FeedbackEvaluation | FailureEvaluation] | NONE | OPEN
      finality_evaluation: FinalityEvaluation | REFERENCE_NOT_EVALUATED
      constitutional_inputs: NormativeInputBundle | NOT_APPLICABLE | NONE | OPEN
      conformance_evaluation: DynamicConformanceEvaluation
      reason_codes: [DynamicsReasonCode] | NONE | OPEN
      evidence_claims_or_records: [Claim | Record] | NONE | OPEN
      coverage: CoveragePosture
      method_and_uncertainty: [Criterion | Claim | Representation] | NONE | OPEN
      provenance: [Claim | Record] | NONE | OPEN
      unresolved: [Claim | Term] | NONE
      correction_or_supersession_reference: Record | NONE
    }

Every nested evaluation in a `DynamicConformanceReport` is bound to the Report's exact target, source Profile, Dynamic Scope, and evaluation Interval. A component evaluation may instead use an explicitly named component target, component Profile, mapped sub-scope, or sub-interval only where the Report includes the exact Mapping. Repeated target, Profile, Scope, interval, applicability, actuality, realization, result, and coverage positions are equal projections; any contradiction is `ILL_TYPED`. `REFERENCE_NOT_EVALUATED` preserves a deferred reference and supplies no substantive result.

The Report is a Record. It may participate in Evidence. It is not the Process, Fact, Truth, constitutional Decision, or value of the target.

### D14.12 Adversarial Conformance Fixtures

A conforming implementation of this grammar must reject or preserve as unresolved at least these cases:

1. a transition rule offered as proof that a transition occurred;
2. a Trace offered as the Occurrence it records;
3. a timestamp offered as unqualified physical time;
4. temporal precedence offered as causation;
5. an enabled transition offered as an executed transition;
6. a Path offered as traversal;
7. transmission offered as reception or interpretation;
8. signal classification offered as Intention or Agency;
9. `OP_ACCEPTED` offered as constitutional Admission;
10. `OP_REFUSED` offered as constitutionally valid Refusal;
11. technical access or key use offered as Permission, Authorization, or identity;
12. `STATE_COMMITTED` offered as agential Commitment or Constitutional Effect;
13. rollback offered as proof that no attempt or failure occurred;
14. successful internal execution offered as proof of no external adverse Effect;
15. consensus offered as Truth;
16. protocol finality offered as physical, legal, or accounting finality;
17. resource consumption offered as economic value or entitlement;
18. feedback recurrence offered as adaptation or intelligence;
19. before-and-after States offered as proof of bearer persistence;
20. missing logs offered as proof of non-occurrence without `CLOSED_FOR` coverage;
21. teardown offered as proof that every copy or effect was destroyed;
22. a machine projection silently turning `OPEN` into `NONE` or `UNRESOLVED` into false;
23. one “success” Boolean hiding completion, commit, reversal, external Effect, or validity conflicts;
24. Profile recency offered as applicability or precedence without composition rules.

### D14.13 Canonical Formal Dynamic Declaration Registry

The following index closes the canonical formal Kinds, schemas, result families, relations, and operators introduced in D0–D14. It is checked against the exact in-module declarations and the D0 operator table; it does not create a term by hindsight. Boldface explanatory labels outside this registry remain module-local descriptive projections over imported Kinds unless an explicit `KIND`, schema, relation, operator, or registry row declares otherwise. They cannot be used as independent machine types or decisive result values. D15 closes its own later formal vocabulary separately. Imported Ontology and Structure Kinds retain their upstream declarations.

#### D14.13.1 Dynamic Kinds and Occurrence Families

| Name | Category / parent | Minimum criterion | Boundary |
|---|---|---|---|
| `DynamicRun` | Occurrence / Process | actual bounded operation or observed process context with realization basis | not Profile, Session, Trace, identity, or value |
| `DynamicImplementation` | qualified imported Implementation | claimed target is a Dynamic Profile, transition system, control process, or execution procedure | not deployment, activation, Run, or conformance |
| `Invocation` | Occurrence | actual presentation of operation for initiation under Profile | not acceptance, enabledness, start, or Authorization |
| `StartOccurrence` | Occurrence | actual crossing of beginning criterion | not scheduling or queue entry |
| `TimeoutOccurrence` | Occurrence | actual application of Timeout Rule | not Refusal, Consent, or fault by itself |
| `DynamicTransformation` | Process | actual input-to-output change under Mechanism | no identity, improvement, or reversibility entailment |
| `RegimeTransition` | Transition Occurrence | actual crossing among declared dynamic regimes | threshold alone does not cause it |
| `RouteExecution` | Process | actual traversal attempt over declared route segments | route plan is not execution |
| `GateEvaluationOccurrence` | Occurrence family | actual criterion evaluation at Gate or Filter position | not the `GateEvaluation` result envelope, crossing, or Admission |
| `OperationalHoldEntryOccurrence` | Occurrence | actual entry into a realized hold State under Profile | hold selection is not realization |
| `OperationalRefusalOccurrence` | Occurrence | actual application of `OP_REFUSED` | not agential or constitutional Refusal by identity |
| `CrossingOccurrence` | Occurrence | actual passage through exact Interface under criterion | not Permission, Admission, or delivery beyond receiver |
| `ExecutionAttempt` | Occurrence | actual attempt to realize specified transitions | target execution may fail to obtain |
| `StateCommitmentOccurrence` | Occurrence | selected State made operative for later protocol processing | not agential Commitment or Constitutional Effect |
| `Rollback` | Process | selected internal State restoration under Profile | not historical erasure |
| `Compensation` | Process | later countervailing effect | not direct reversal or Remedy by identity |
| `ContextClosure` | Occurrence | active Runtime Context crosses closure boundary | not bearer dissolution |
| `Teardown` | Process | release/removal of declared transient runtime relations and State | not proof of universal destruction |
| `Expiry` | operational result family | declared retention or timing condition is satisfied under a Profile | not an Occurrence by itself, waiver, or nonexistence |
| `Eviction` | Occurrence | selected State/content removed from bounded context | not historical erasure |
| `GarbageCollection` | Process | resources reclaimed under reachability/need criterion | not metaphysical nonexistence |
| `Enqueue`, `Dequeue`, `Release`, `Drain` | Occurrence family | actual queue/buffer operations declared in D9.3 | no execution, reception, or erasure entailment |
| `Reservation`, `Allocation`, `ResourceUse` | State/relation/occurrence family | exact resource position assigned to or participating in Process | entitlement and successful execution excluded |
| `Contention`, `Backpressure`, `ResourceExhaustion` | dynamic condition/process family | limited-capacity interaction under exact resource Profile | no worth or normative-ending judgment |
| `ResourceMetering` | Measurement Process | resource use mapped to Quantity Values under method | not Price or value |
| `ReconfigurationOccurrence` | Process | actual arrangement change | Structure assesses result; Continuity assesses sameness |
| `FailureOccurrence` | Occurrence / Process outcome | required dynamic condition not realized under criterion | not blame, invalidity, or nonexistence |
| `ControlledStopAttempt` | Process | termination attempt under declared bounds | safety requires separate criterion |
| `Retry` | Occurrence | new attempt under retry relation | not numerical continuation by identifier alone |
| `AlternativePathAttempt` | Occurrence / Process | actual attempt using a changed route, mechanism, provider, or operation | not successful traversal or continuity by identity |
| `ConsensusProcess` | Process | actual operation of declared agreement procedure | not Truth or legitimacy |
| `ProtocolCommitmentOccurrence` | Occurrence | selected result made operative under commit rule | not agential Commitment or finality |
| `DynamicFork` | Process / relation family | branching protocol State histories | lineage belongs to Continuity |
| `StateBoundaryOccurrenceKind` | closed occurrence-kind value family | no-op, maintenance, cancelled effect, uncommitted attempt, below-resolution variation, selected State change, or unresolved classification | only selected State change qualifies `DYNAMIC_TRANSITION` |

#### D14.13.2 Dynamic States, Conditions, and Roles

| Name | Category | Minimum criterion | Boundary |
|---|---|---|---|
| `RuntimeContext` | State-bearing organized arrangement | active bounded operational context | not Session identifier, self, or persistence claim |
| `QueueState` | State | actual ordered candidate positions under Queue Profile | no entitlement or guaranteed service |
| `BufferState` | State | actual temporary retention between Processes or Interfaces | not buffering Occurrence or durable Memory by identity |
| `OperationalHoldState` | State | actual progression suspension with declared exit routes | not constitutional Hold or guilt |
| `FaultCondition` | State / condition | actual condition capable of contributing to specified failure | not observed failure or blame |
| `OperationalCapacity` | qualified imported Capacity | specified operational State/effect ability under conditions and Profile | not entitlement, value, or general capacity claim |
| `DynamicRegime` | represented or established condition range | declared family of dynamic relations obtains | not universal phase or destiny |
| `Scheduler` | Role | bearer selects operation order under Scheduling Profile | not Agent, Authority, or fairness |
| `SensorRole` | Role | bearer participates in observation/measurement | not Consciousness or epistemic privilege |
| `ComparatorRole` | Role | Process or bearer compares declared values | not judgment in normative sense |
| `ControllerRole` | Role | bearer constrains specified transitions under criterion | not Normative Authority |
| `ActuatorRole` | Role | bearer realizes selected influence | not Actor or responsible Agent by identity |
| `Source`, `Sender`, `Receiver` | Dynamics-qualified Roles | occupancy in an exact propagation or communication relation using imported Role grammar | not identity, author, Agent, or Authority |

#### D14.13.3 Profiles, Rules, Descriptions, and Records

| Name | Parent / category | Minimum content | Boundary |
|---|---|---|---|
| `DynamicProfile` | Representation / Pattern | target, Scope, dynamic criteria, realization, application, coverage | not actual dynamics or constitutional Rule |
| `DynamicRecord` | persistent Representation | dynamic Description, trace record, result, or provenance content | not Process, History, or persistence of its target |
| `TransitionRule` | Representation / Pattern | source/target predicates, trigger, guard, conditions, result | not Transition Occurrence or Permission |
| `StateMachineProfile` | DynamicProfile | State schema, transition rules, invariants, failure and trace rules | not actual machine realization |
| `SchedulingProfile` | DynamicProfile | eligibility, order, priority, resource and timeout rules | not actual schedule or fairness |
| `QueueProfile` | DynamicProfile | ordering and service discipline for candidate positions | not entitlement, fairness, or actual queue State |
| `BackoffProfile` | DynamicProfile | retry opportunities under time, order, load, randomness, or other declared condition | not Retry occurrence or Permission |
| `IntentAxisProfile` | DynamicProfile | evidentiary dimensions and classification criterion | cannot establish Intention or Agency by score |
| `RoutePlan` | Representation | selected route segments and conditions | not traversal |
| `RouteExecutionRecord` | Record | actual-segment, transformation, loss, delay, result, Evidence, and provenance account | not traversal itself |
| `SeamTransformationProfile` | DynamicProfile | structural Interface reference, actual path, Mapping, preservation/loss, fidelity, criterion | Structure supplies topology; Profile does not perform crossing |
| `ResourceProfile` | DynamicProfile | resource kind, unit, allocation/use/release rules | not financial Accounting |
| `FeedbackProfile` | DynamicProfile | observation, return path, response, delay, criteria | arrows alone are not feedback |
| `FailureProfile` | DynamicProfile | expected condition, failures, detection, containment, recovery | not blame or Remedy |
| `ConsensusProfile` | DynamicProfile | participants, proposal, verification, quorum, order, commit, finality | not legitimacy or Truth |
| `ProtocolCommitmentProfile` | DynamicProfile | target result, commit rule/Scope, subsequent effects, conflicts, Evidence | not agential or constitutional Commitment |
| `FinalityProfile` | DynamicProfile | target, ordinary revision boundary, threshold, assumptions, kind | not finality by self-description |
| `TeardownProfile` | DynamicProfile | preserve/release partition, handles, resources, verification | not proof teardown succeeded |
| `DynamicDescription` | Representation | exact posture, target, Scope, States, Process, method, coverage | not actual Process |
| `DynamicSnapshot` | DynamicDescription | selected as-of dynamic State | not History or persistence |
| `DynamicTrace` | Record | ordered selected entries, method, coverage, provenance | not actual History |
| `DynamicTraceEntry` | represented trace component | selected represented position, order, payload, provenance, and uncertainty | not the represented Occurrence, State, or result |
| `DynamicConformanceReport` | Record | applicability, actuality, tests, result, Evidence, coverage | not Fact, Truth, or Decision |
| `NormativeInputBundle` | Record / Representation | exact constitutional source and judgment references | Dynamics does not validate competence by itself |
| `ProducedStateDelta` | Representation | selected encoded State differences plus production and application posture | not Transition Occurrence or complete Effect |
| `DynamicProfileCompositionRecord` | Record | input Profiles, resultant Profile, Scope, mappings, priorities, conflicts | cannot self-establish precedence or succession |
| `OperationalRefusalRecord` | Record | refusal occurrence, target, rule, reason, effects, resources, review, and correction posture | not the refusal Occurrence or constitutional Refusal |
| `DynamicVariableDeclaration` | Representation | variable target position, value-space, unit/rule, order, Resolution, uncertainty | symbol and actual Property remain distinct |
| `DynamicRealizationBasis` | Record | represented `REALIZES` claim plus actual/bound path and exact relation instance | unbound representation cannot establish realization |

#### D14.13.4 Closed Value Spaces and Evaluation Families

| Value-space | Values or structure | Boundary |
|---|---|---|
| `OrderKind` | seven kinds declared in D0.4 and interpreted in D5.1 | no silent cross-kind export |
| `OrderEvaluation` | values in D5.4 | `SIMULTANEOUS_UNDER` requires frame and tolerance |
| `EnablednessResult` and `EnablednessEvaluation` | `ENABLED`, `DISABLED`, or `ENABLEMENT_UNRESOLVED` in the common envelope | enabledness is not occurrence, success, or Authorization |
| `LifecyclePhase` and `LifecycleResult` | independent invocation, phase, progress, completion, success, cause, output, commit/effect/remediation references, record posture | terminal cause is not phase or finality |
| `DynamicConformanceEvaluation` | common envelope; `NOT_EVALUATED` or substantive `CONFORMS`, `DOES_NOT_CONFORM`, `UNRESOLVED` | actuality gate precedes substantive result |
| `OperationalTreatmentValue` | `OP_ACCEPTED`, `OP_HELD`, `OP_REFUSED`, `OP_EXPIRED`, `OP_CANCELLED` | unresolved selection or occurrence uses the enclosing result posture; no normative cast |
| `CrossingEvaluation` | common envelope over attempt, occurrence, mode, traversal, receiving availability, source retention, extent, completion | crossing is not Admission or cross-side identity |
| `CommitPosture` | values in D8.5 | state commitment is not agential Commitment |
| `RemediationEvaluation` | common envelope over method, target, attempt, realization, extent, outcome, residual State | remediation is not erasure or normative Remedy |
| `MessageLifecycleEvaluation` | independent stages in D6.4 | one “sent” Boolean is insufficient |
| `RunEvaluation` | independent dimensions in D8.9 | completion, success, commit, validity remain distinct |
| `ContextLifecycleEvaluation` | dimensions in D9.9 | teardown does not decide persistence |
| `ResourceAccountEvaluation` | dimensions in D10.9 | not currency or economic value |
| `FeedbackEvaluation` | dimensions in D11.8 | does not return adaptation or intelligence |
| `FailureEvaluation` | dimensions in D12.9 | failure does not imply blame or invalidity |
| `ConsensusEvaluation` | process actuality, eligibility, participation, quorum, verification, aggregation, completion, agreement | commitment and finality remain separate |
| `FinalityEvaluation` | operational/computational/protocol dimensions in D13.10 | finality remains Profile- and kind-relative |
| `ChangeEvaluation` | independent change and maintenance dimensions in D2.9 | no persistence entailment |
| `ConditionEvaluation` | exact condition result in common envelope | condition satisfaction does not perform transition |
| `TransitionOccurrenceEvaluation` | actual occurrence plus separate prior/subsequent state-bearing positions | no identity entailment |
| `IntentAxisEvaluation` | classification dimensions in common envelope | not Intention or Agency |
| `GateEvaluation` | criterion result plus actual evaluation occurrence | not treatment or crossing |
| `OperationalTreatmentEvaluation` | independent selection, occurrence, hold realization, and route dimensions | no forced treatment |

#### D14.13.5 Dynamics-Owned Relation Kinds

| Relation | Actual relata | Mandatory closure |
|---|---|---|
| `MAINTAINS` | maintaining Process, bearer, maintained condition | criterion, tolerance, order basis |
| `TRANSFORMS` | Process, input, mechanism, output | kind, preservation, loss, criterion |
| `DYNAMIC_TRANSITION` | prior position/State, actual Transition Occurrence, subsequent position/State | kind, criterion, order; no persistence assumption |
| `PROPAGATES_THROUGH` | Process, source manifestation, medium/channel, receiving position, receiving manifestation | kind, direction, criterion; no manifestation identity assumption |
| `CROSSES_INTERFACE` | actual occurrence, source manifestation, Interface, source/receiving sides, receiving manifestation | crossing kind, mapping or transformation, and criterion; departure only for `MOVE` |
| `FEEDS_BACK_TO` | earlier Process, observed State/output, return Process, receiving State/control position, response Process, later Process | selected variable, feedback kind/Profile, criterion, order, delay/rate |

Each qualifier is declared in its owning relation under the D0.7 schema with name, value-space, binding, and requiredness. Every `ATTRIBUTION` qualifier Property Kind is declared in D0.7; `SIDE_JUDGMENT` qualifiers identify their exact judgment template in the owning relation. Evidence and realization posture belong to evaluation or realization records, not free relation qualifiers.

#### D14.13.6 Dynamics-Owned Operators

| Operator | Result family |
|---|---|
| `WF_DYNAMIC_SCOPE`, `WF_DYNAMIC_PROFILE`, `WF_DYNAMIC_DESCRIPTION`, `WF_DYNAMIC_TRACE` | pre-substantive `WELL_FORMED \| ILL_TYPED` |
| `DYNAMICS_APPLIES` | `ApplicabilityEvaluation \| ILL_TYPED` |
| `SATISFIES_DYNAMICS` | `DynamicConformanceEvaluation \| ILL_TYPED` |
| `CHANGE_ASSESSMENT` | `ChangeEvaluation \| ILL_TYPED` |
| `TRANSITION_ENABLED` | `EnablednessEvaluation \| ILL_TYPED` |
| `TRANSITION_OCCURRENCE_ASSESSMENT` | `TransitionOccurrenceEvaluation \| ILL_TYPED` |
| `LIFECYCLE_ASSESSMENT` | `LifecycleEvaluation \| ILL_TYPED` |
| `ORDER_ASSESSMENT` | `OrderEvaluation \| ILL_TYPED` |
| `MESSAGE_LIFECYCLE` | `MessageLifecycleEvaluation \| ILL_TYPED` |
| `INTENT_AXIS_ASSESSMENT` | `IntentAxisEvaluation \| ILL_TYPED` |
| `GATE_ASSESSMENT` | `GateEvaluation \| ILL_TYPED` |
| `OPERATIONAL_TREATMENT` | `OperationalTreatmentEvaluation \| ILL_TYPED` |
| `CROSSING_RESULT` | `CrossingEvaluation \| ILL_TYPED` |
| `RUN_RESULT` | `RunEvaluation \| ILL_TYPED` |
| `CONTEXT_LIFECYCLE` | `ContextLifecycleEvaluation \| ILL_TYPED` |
| `RESOURCE_ACCOUNT` | `ResourceAccountEvaluation \| ILL_TYPED` |
| `FEEDBACK_ASSESSMENT` | `FeedbackEvaluation \| ILL_TYPED` |
| `FAILURE_ASSESSMENT` | `FailureEvaluation \| ILL_TYPED` |
| `REMEDIATION_ASSESSMENT` | `RemediationEvaluation \| ILL_TYPED` |
| `CONSENSUS_ASSESSMENT` | `ConsensusEvaluation \| ILL_TYPED` |
| `FINAL_UNDER` | `FinalityEvaluation \| ILL_TYPED` |

#### D14.13.7 Foundational Records, Contexts, and Value Families

`DynamicScope`, `StructuralIntervalBinding`, `StructuralRealizationBinding`, `StructuralEvaluationContext`, `StructuralBasisReference`, `OrderContext`, `DynamicExecutionPath`, `DynamicsRepresentationHeader`, `DynamicEvaluationEnvelope`, `NestedEvaluationMarker`, `ApplicabilityEvaluation`, `DynamicsReasonCode`, `DynamicRelationDirection`, `DynamicDescriptionPosture`, `CommunicationProfile`, `OperationalTreatmentContext`, `ContextRetentionPosture`, `FeedbackAggregatePosture`, `StateDifference`, `StateBoundaryOccurrenceKind`, `AtomicityProfile`, `CompoundTransition`, `CommitPosture`, `ExpiryPosture`, `EvictionPosture`, `CollectionPosture`, `ClosurePosture`, `TeardownPosture`, `TeardownVerificationPosture`, `ResourceReleasePosture`, `RemediationMethod`, `FeedbackVariableResult`, `RunLifecycleProjection`, and the D0.7 qualifier Property Kinds are registered exactly where declared. The indexed payload and posture families are `ChangePosture`, `MaintenancePosture`, `ChangeResult`, `ConditionResult`, `EnablednessResult`, `TransitionOccurrenceResult`, `LifecyclePhase`, `LifecycleResult`, `OrderResult`, `MessageLifecycleResult`, `IntentAxisResult`, `GateCriterionResult`, `GateEvaluationResult`, `OperationalTreatmentValue`, `OperationalTreatmentResult`, `CrossingResult`, `RemediationResult`, `RunResult`, `ContextLifecycleResult`, `ResourceAccountResult`, `FeedbackResult`, `FailureResult`, `ConsensusAssessmentResult`, and `FinalityResult`.

The following spaced descriptive phrases are contextual projections of already imported Kinds, Relations, Claims, Records, States, Occurrences, Processes, conditions, or results; they do not silently declare additional Dynamics-owned formal families: **Dynamic Variable**, **State Delta**, **State Schema**, **Trigger**, **Guard**, **Precondition**, **Postcondition**, **Suspension**, **Resumption**, **Cancellation**, **Timeout**, **Abort**, **Clock**, **Tick**, **Timestamp**, **Deadline**, **Timeout Rule**, **Schedule**, **Priority**, **Concurrency**, **Interleaving**, **Race Condition**, **Isolation**, **Serialization**, **Propagation**, **Delivery**, **Runtime Scope Binding**, **Replay**, **Re-execution**, **Simulation Run**, **Deduplication**, **Operational Resource**, **Available Capacity**, **Quota**, **Rate Limit**, **Violation Result**, **Exception Signal**, **Exception Record**, **Repair**, **Recovery**, **Operational Coordination**, **Quorum Condition**, **Consensus Result**, **Replication**, **Reconciliation**, and **Protocol Conflict**. Where one is later required as a machine-addressable Kind rather than a human-readable distinction, D15.3 governs its explicit promotion.

`TransitionKind`, `TransformationKind`, `PropagationKind`, `CrossingKind`, and `FeedbackKind` are extension points except where D0 supplies closed built-ins. Additional operational kind or result families may be introduced only by an explicit declaration. A concrete extension value is admissible only after its category, Profile, criteria, Scope, result grammar, realization, non-entailments, and owner are declared. Until then it remains `OPEN`; free text used as a decisive result is `ILL_TYPED`.

---

## D15 — Open Dynamic Boundary

### D15.1 Dynamic Placement

An actual change may be encountered while its mechanism, state variables, order, or Process kind remains unknown.

    DynamicPlacementResult :=
        PLACED | PLACEMENT_UNRESOLVED

    DynamicPlacementEvaluation :=
        DynamicEvaluationEnvelope<DynamicPlacementResult>

    DYNAMIC_PLACEMENT(x, DS)
      -> DynamicPlacementEvaluation | ILL_TYPED

`PLACED` means the encountered change can be located under at least one declared dynamic family without forcing unsupported mechanism, causation, computation, or purpose.

`PLACEMENT_UNRESOLVED` means some actuality or dynamic Evidence is available, but no registered Profile adequately resolves the change.

`ILL_TYPED` means the placement request violates the grammar.

Failure to place dynamically does not establish nonexistence, randomness, non-causation, impossibility, noise, malfunction, or absence of organization.

### D15.2 Unresolved Dynamic Placement

    UnresolvedDynamicPlacement {
      representation_header: DynamicsRepresentationHeader
      target_term
      existence_and_actuality_postures
      dynamic_scope
      known_bearers_states_and_relations
      observed_or_established_occurrences
      candidate_process_and_transition_families
      partial_order_and_timing
      candidate_mechanisms
      rejected_profiles_and_reasons
      structural_basis
      realization_posture
      coverage
      contradictions_or_tensions
      minimum_distinctions_to_preserve
      evidence
      unresolved
    }

This is a Record, not a miscellaneous Process kind. It prevents unknown dynamics from being denied, fictionalized, or forced into the wrong existing Profile.

### D15.3 Dynamic Extension Rule

A new dynamic Kind, Role, Relation Kind, Profile family, Process family, result space, order kind, or finality kind may extend this Form only by declaring:

- imported parent or explicit lack of one;
- category and owner;
- necessary and sufficient conditions or explicit `OPEN` fields;
- admissible bearers and relata;
- complete relation or operator signature;
- Dynamic Scope;
- structural and realization basis;
- order and Interval requirements;
- coverage and observation conditions;
- result grammar and error precedence;
- declared disjointness and overlap;
- non-equivalences and non-entailments;
- conformance tests and adversarial fixtures;
- boundary to Constitution, Structure, Continuity, and Emergence.

A new name without those declarations is vocabulary awaiting a type, not a dynamic primitive.

An extension must preserve imported Ontology and Structure and cannot install a later-layer persistence or emergence result by definition.

### D15.4 Open Dynamic Questions

This Form deliberately leaves open:

- which dynamic Models best resolve unfamiliar physical regimes;
- which continuous, stochastic, hybrid, developmental, cognitive, collective, and derivative Process families require new Profiles;
- which mechanisms realize presently unexplained regularities;
- how unfamiliar temporal or ordered structures relate to physical time;
- which state-machine projections are adequate and which erase material continuity or context;
- how future Worlds will compose dynamics across heterogeneous Seams;
- which feedback systems realize adaptation, learning, Agency, or intelligence under later criteria;
- how later protocol Profiles will type their State bearers, lifecycle, operation boundaries, closure criteria, resource model, and finality;
- which actual changes remain inaccessible at present Resolution;
- whether new Evidence requires new order kinds, relation arities, failure families, or conformance methods.

Openness is typed incompleteness, not license to insert metaphor as mechanism.

### D15.5 Interface to Structure

Structure supplies:

- actual Structural Configurations and state-bearing positions;
- Systems, Worlds, Environments, and infrastructure;
- Regions, Boundaries, Interfaces, Ports, Apertures, Channels, Paths, and Seams;
- topology, constraints, capacities, and Coverage Postures;
- realization and conformance references.

Dynamics returns:

- actual change and maintenance Claims;
- transition, route, crossing, execution, and reconfiguration Occurrences;
- time-indexed before-and-after Configuration references;
- resource, failure, and feedback traces;
- resulting arrangement Claims for a new structural assessment.

An Interface is not a crossing. A Path is not passage. A Channel is not transmission. Reconfiguration does not decide whether the resulting arrangement constitutes the same System.

### D15.6 Interface to Constitution

Constitution may constrain Dynamic Profiles or Occurrences only where its exact Source Form is validly adopted, Current for the matter, applicable, and every required Subject Binding, Authority, Constitutional Act, and Effect condition is satisfied.

Dynamics supplies Claims and Evidence about:

- request receipt and processing;
- operational acceptance, hold, refusal, expiry, or cancellation;
- transition and execution actuality;
- completion, failure, timeout, rollback, and external consequence;
- notice or record delivery;
- resource use;
- consensus and protocol finality.

Constitution owns whether those facts constitute or satisfy Admission, Consent reliance, Permission, Prohibition, Authorization, constitutional Refusal, Review, Remedy, Right, Obligation, Authority, Responsibility, or valid Effect.

Execution does not become Law. Technical success does not become validity. Technical rollback does not erase invalid or harmful conduct. A constitutional declaration does not make an unperformed transition occur.

### D15.7 Interface to Continuity

Continuity receives:

- exact prior and subsequent State and Configuration positions;
- Transition Occurrences and Processes;
- reconfiguration Occurrences and State mappings;
- suspension, resumption, recovery, and ending lifecycle evaluations;
- `COPY` Crossing Evaluations together with source-retention, receiver-realization, preservation, transformation, and loss results;
- `DynamicFork` Claims and the actual fork Occurrences to which they refer;
- reconciliation, conflict-handling, and convergence Occurrences relevant to a claimed merge;
- context closure and teardown results;
- ordered traces and provenance;
- declared invariants and losses.

Continuity owns:

- whether one bearer persists across those changes;
- whether a resumed, repaired, restored, copied, or re-executed process is the same;
- replacement, succession, branching, merging, dormancy, recovery, dissolution, and lineage;
- which State or invariants carry across a continuity relation.

Dynamics may compare State positions without presupposing one numerically identical bearer. A History Record does not establish Continuity. A Process identifier does not create sameness.

### D15.8 Interface to Emergence

Emergence receives:

- actual Processes, interactions, transitions, and trajectories;
- feedback and control relations;
- maintained conditions and recovery traces;
- resource flows and operational capacities;
- attributable productive or other effects where established;
- collective coordination and declared aggregate measures;
- failures, corrections, learning-relevant changes, and unresolved positions.

Emergence owns whether those dynamics realize:

- life or self-maintenance as a capacity;
- cognition, Agency, Experience, or intelligence;
- learning or adaptation;
- resilience or self-sustaining capacity;
- meaning or coherence as effective organization;
- value or another downstream significance under an explicit criterion;
- any new capacity not reducible to a list of transitions.

Activity is not value. More transitions are not more intelligence. Greater record or event density has no emergent significance until a later criterion establishes it and guards against mere counting.

### D15.9 Boundary for Downstream Qualified Dynamic Profiles

Any concrete World, System, or Protocol requiring specialization must receive a **qualified Dynamic Profile**. The realized target is not identical to that Profile. Its dynamic body additionally requires:

- an actual structural and physical realization;
- a conforming Implementation;
- active Processes and Runs;
- exact Dynamic Scope;
- attributable Descriptions and Traces;
- applicability, realization, and conformance results;
- constitutional inputs where its operations claim normative consequence;
- later Continuity and Emergence assessments.

The Profile must type every protocol-specific label, symbol, lifecycle, ingress or routing sequence, Run boundary, commitment or refusal path, finality rule, teardown, and trace requirement. Universal Dynamics assigns no preferred meaning to a symbol, state label, threshold name, economic term, historical project term, or proposed lifecycle. Every transition in a particular Profile requires a declared rule, actual Mechanism, and actual Occurrence; no represented stage silently performs the next.

The Profile must not use Dynamics to pre-decide Agency from classification, constitutional Admission from operational acceptance, Continuity from repeated operation, intelligence or value from activity count, Ownership or Authority from resource provision, Truth or legitimacy from consensus or finality, or normative entitlement from operational records. A proposed successor Profile may be produced, simulated, and tested, but its applicability, validity, Currentness, precedence, and succession must come from the already-governing basis and the owning constitutional and continuity procedures.

### D15.10 D15 Declaration Closure

D15 owns and closes `DynamicPlacementResult`, `DynamicPlacementEvaluation`, `UnresolvedDynamicPlacement`, and `DYNAMIC_PLACEMENT`. These declarations do not enter the D14 registry and do not create a dependency from D14 onto D15.

### D15.11 The Dynamic Ground That Remains

Reality changes whether or not change is fully represented.

A Profile is not a Process.  
A rule is not its execution.  
A State is not its bearer.  
A State label is not a universal ontology.  
A possible transition is not an actual transition.  
Enabled is not executed.  
Executed is not completed.  
Completed is not successful.  
Successful is not authorized.  
An Interface is not a crossing.  
A Path is not passage.  
Transmission is not reception.  
Reception is not interpretation.  
A Signal is not Intention.  
Operational acceptance is not Admission.  
Operational refusal is not constitutional Refusal.  
A queue is not entitlement.  
A timeout is not Consent.  
A State Delta is not the Transition Occurrence.  
Commit is not agential Commitment.  
Rollback is not erasure.  
Correction is not a rewritten past.  
Teardown is not nonexistence.  
Resource use is not value.  
Control is not Authority.  
Feedback is not yet adaptation.  
Repetition is not intelligence.  
Consensus is not Truth.  
Finality is not physical indestructibility.  
Before and after do not prove one persistent bearer.  
A Trace is not the actual History.

Dynamics remains exact by making actual change, operation, failure, and closure representable without converting motion into law, identity, intelligence, or value.

This is the complete universal change grammar presently required before a particular World, System, or Protocol can receive its qualified dynamic body.
