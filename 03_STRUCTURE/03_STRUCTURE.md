# STRUCTURE

> **Structural completeness invariant**
>
> Every claimed actual arrangement must identify its bearers, typed Relation Instances, World context, Domain, Interval, Resolution Context, realization basis, and coverage posture. Where a Boundary is claimed, its criterion is mandatory; otherwise the boundary position is explicitly `NONE`, `OPEN`, or `NOT_APPLICABLE` as the exact grammar permits. No structural position creates identity, Agency, Authority, Ownership, persistence, or value by itself.

> **Open-structure invariant**
>
> Not every arrangement is already understood or named. Every encountered arrangement can remain structurally located without being denied, converted into a diagram, forced into a tree or graph, or promoted into a System that has not been shown to exist.

## Status and Dependency Declaration

| Field | Declaration |
|---|---|
| Form role | Human Source Form |
| Layer | Structure |
| Source-form status | `FINAL` |
| Semantic source | This exact human-readable Form |
| Machine projection | Permitted only as a derivative Representation |

This Form imports the following exact upstream Forms:

| Layer | Canonical Reality Model path | SHA-256 | Bytes | Imported function |
|---|---|---|---:|---|
| Substrate | `00_SUBSTRATE/00_SUBSTRATE.md` | `6c3ea987c07840c9f019e6f9bee6e54e9eef6c4465359036c0ab2b455a431b99` | 9,301 | Physical ground |
| Ontology | `01_ONTOLOGY/01_ONTOLOGY.md` | `25bdaf01bab4db5fea2e88d7efb207428ec0ee64797d9ebb8fb7b0f659819efd` | 168,760 | Declarative semantic type system |
| Constitution | `02_CONSTITUTION/02_CONSTITUTION.md` | `6a2c369e7570d5dd717b98cbb8216fed8d3c225eeeb166661d362b3120608f12` | 97,393 | Constitutional constraints and layer boundary |

Importing the Constitution does not adopt it, make it Current, make it applicable, bind any bearer, or create any constitutional Effect. Its constraints govern an exact matter only where the Constitution is validly adopted under its governing source, `CURRENT` for that matter, `APPLIES`, and every required `BOUND_BY`, Constitutional Act, or Effect gate returns the required valid result.

The paths identify layers relative to the root of the exact Reality Model stack. A changed upstream byte sequence is a different imported Form. Recency, filename reuse, publication, parsing, or implementation does not silently replace a dependency.

## Internal Route and Ownership

| Module | Structural job | Depends on |
|---|---|---|
| S0. Structural Position | Establishes the layer, imports, limits, judgment grammar, operators, and invariants | Substrate + Ontology + Constitution |
| S1. Configuration and Structural Forms | Separates actual arrangement, Profile, Description, Record, Scope, and coverage | S0 |
| S2. Composition and Cardinality | Types constitution, parthood modes, components, subsystems, multiplicity, and non-inheritance | S0–S1 |
| S3. Region, Boundary, and Containment | Types interior, exterior, partition, overlap, enclosure, and containment | S0–S2 |
| S4. Position, Placement, and Topology | Types positions, occupancy, adjacency, incidence, connection, paths, and topology-relative Roles | S0–S3 |
| S5. Interface, Seam, Port, and Aperture | Types boundary crossings as structural possibilities without claiming that crossing occurs | S0–S4 |
| S6. Systems, Worlds, Environments, and Architecture | Types organized wholes and their realized arrangements | S0–S5 |
| S7. Relational Structures and the Third-Bearer Test | Separates Relation, shared surface, mediator, and Relational System | S0–S6 |
| S8. Constraint, Dependency, and Structural Integrity | Types admissible arrangement, conflict, dependency, redundancy, and defects | S0–S7 |
| S9. Description, Conformance, and Representation | Governs structural Claims, Profiles, tests, errors, equivalence, and derivative Forms | S0–S8 |
| S10. Open Structural Boundary | Locates unresolved arrangements and fixes downstream interfaces | S0–S9 |

The dependency order is one-way:

    SUBSTRATE -> ONTOLOGY -> CONSTITUTION -> STRUCTURE
                                                    |
                                                    v
                S0 -> S1 -> S2 -> S3 -> S4 -> S5
                                           |          |
                                           v          v
                                          S6 -> S7 -> S8 -> S9 -> S10

A later module may specialize an earlier structural relation. It may not silently redefine an imported Kind, change an earlier operator, or convert a later-layer result into structural actuality.

---

## S0 — Structural Position

### S0.1 Structural Function

Ontology states what kinds of Entity, Occurrence, State, Relation, Boundary, Interface, System, World, and Representation can be distinguished. Structure states how actual bearers, Relation Instances, regions, boundaries, interfaces, systems, and constraints are arranged within an exact Scope.

Structure owns **arrangement**.

It does not:

- create Reality by description;
- redefine an imported Kind;
- turn a Role into its bearer;
- turn a Relation into either relatum or a third Entity;
- make a Profile, graph, diagram, schema, or Record into the arrangement it represents;
- perform a transition, crossing, computation, message delivery, refusal, consensus process, or failure response;
- establish that a bearer remains numerically the same across change;
- produce life, cognition, Agency, Experience, intelligence, meaning, coherence, or value;
- constitute Admission, Consent, Permission, Prohibition, Right, Obligation, Authority, Ownership, entitlement, or jurisdiction.

An actual structure may exist without a complete Description. A structurally exact Description may refer to nothing actual. Structural conformance and actuality are separate.

### S0.2 Imported Semantic Ground

The following are imported unchanged from Ontology:

- `EXISTS`, `OBTAINS`, `SAME`, `DIFFERENT`, `RELATES`, `SCOPED_TO`, `REALIZES`, and `EXPERIENCES`;
- the canonical judgment `Gamma |- x : K @ D, I, R`;
- `SUBTYPE_OF`, `INSTANCE_OF`, `DISJOINT_WITH`, `OVERLAPS`, and `PLAYS`;
- `Bearer`, `Entity`, `Occurrence`, `Property`, `RelationInstance`, `Pattern`, `State`, `Configuration`, `Role`, `Status`, `Domain`, `Scope`, `Criterion`, `ResolutionContext`, `ReferenceFrame`, and `World`;
- `Boundary`, `Interface`, `Seam`, `System`, `Part`, `Whole`, `Constituent`, `Component`, `Aggregate`, `Organization`, `Environment`, `Channel`, `Coupling`, `Mediation`, `RelationalSystem`, `Participation`, `Realization`, and `Implementation`;
- the separation of Representation, Referent, Record, Claim, Evidence, and Fact;
- the Ontology conformance results and Error Model.

Capitalized imported terms retain those meanings. Structure declares qualified structural Kinds, Roles, Relation Kinds, Profiles, and constraints. It does not become a second Ontology.

Every actual structural relation is asserted through `EXISTS` plus `RELATES`, or through a specialized predicate that is explicitly declared as an existential projection of `RELATES`. Every realization Claim uses imported `REALIZES`. No structural declaration creates an alternative truth operator.

Two imported operators share the surface lexeme `CONSTITUTES` but not one signature or owner:

    ONTOLOGY.CONSTITUTES(parts_and_relations, whole, D, I, criterion)
    CONSTITUTION.CONSTITUTES(constitutional_act, y, C, W, S, I)

Structure declares neither. Bare `CONSTITUTES(...)` is `AMBIGUOUS_OPERATOR` and therefore `ILL_TYPED`; every use must carry its namespace. Structural whole-formation Claims use `ONTOLOGY.CONSTITUTES`. Institutional Status, Relation, or Effect formation Claims use `CONSTITUTION.CONSTITUTES` and remain subject to every constitutional gate.

### S0.3 Layer Boundary

The Reality Model remains:

    SUBSTRATE -> ONTOLOGY -> CONSTITUTION -> STRUCTURE
                -> DYNAMICS -> CONTINUITY -> EMERGENCE

The linear notation states semantic dependency, not an execution sequence.

- **Substrate** constrains every physically realized arrangement.
- **Ontology** owns the Kinds and root relation grammar used here.
- **Constitution** owns normative validity and institutional force.
- **Structure** owns actual arrangement indexed to a Scope and Interval.
- **Dynamics** owns transition, timing, execution, routing, actual crossing, propagation, feedback operation, consensus, resource use, failure handling, operational or computational finality, and execution-resource metering.
- **Continuity** owns persistence, sameness criteria across change, replacement, copying, succession, branching, merging, dormancy, recovery, dissolution, and lineage.
- **Emergence** owns the arising of capacities, life, cognition, Agency, Experience, intelligence, meaning, collective capacity, adaptation, coherence, and value.

Structure may identify the arrangement on which a later process, continuity judgment, or emergent capacity depends. It may not claim the later result by topology alone.

### S0.4 Structural Scope

A **Structural Scope** is a qualified Scope adding an exact World context to the inherited Ontology Scope. It is a direct Scope specialization, not a wrapper containing a second Scope.

    KIND StructuralScope {
      parents: [Scope]
      category: other
      necessary_conditions: [one WorldContext; inherited Domain, Interval, ResolutionContext, and further conditions]
      sufficient_conditions: [all required fields typed; projections coherent; no field contradiction]
      properties_and_value_spaces: [world_context: WorldContext; domain: Domain; interval: Interval; resolution_context: ResolutionContext; further_conditions]
      admissible_bearers: NOT_APPLICABLE
      relation_signatures: [SCOPED_TO]
      realization_criterion: NOT_APPLICABLE
      scope: self
      disjoint_with: OPEN
      overlaps: OPEN
      epistemic_statuses: [Observed, Measured, Inferred, Established, Disputed, Unknown]
      non_equivalences: [StructuralScope / World; StructuralScope / Domain]
      later_layer_boundaries: [transition and persistence are not included]
      owner: S0
    }

Its record projection is:

    StructuralScope {
      world_context:
          SINGLE_WORLD {
            world_term: Term
            world_kind_posture: WorldKindPosture
            world_existence_posture: WorldExistencePosture
          }
        | CROSS_WORLD {
            ordered_world_terms: ordered [Term], length >= 2
            per_world_kind_postures: Mapping(Term -> WorldKindPosture)
            per_world_existence_postures: Mapping(Term -> WorldExistencePosture)
            relation_context: Claim | RelationInstance | Representation
          }
      domain: Domain
      interval: Interval
      resolution_context: ResolutionContext
      further_conditions: [Claim | Criterion | Representation] | NONE | OPEN
    }

`SINGLE_WORLD` locates an arrangement within one World. `CROSS_WORLD` types an arrangement or relation involving two or more named Worlds without inventing a global or meta-World. No listed World thereby contains, governs, or becomes ontologically prior to another.

World context, Domain, Interval, and Resolution Context occur once. A downstream expression may display them separately only as projections from this one record. A repeated field in a Coverage Posture, Realization Basis, relation instance, Description, Profile, algebra, or report must be an equal projection or an explicitly declared sub-scope. A disagreement is `WORLD_SCOPE_MISMATCH` or `SCOPE_ERROR`, as applicable.

    WF_STRUCTURAL_SCOPE(SS)
      -> WELL_FORMED | ILL_TYPED

`WELL_FORMED` requires:

1. every World term is accompanied by an exact World-kind and existence posture;
2. `SINGLE_WORLD` contains exactly one World term, or `CROSS_WORLD` contains an ordered list of at least two World terms plus the exact relation context requiring them;
3. the inherited Scope fields contain Domain, Interval, and Resolution Context directly;
4. every further condition is typed or explicitly `OPEN`;
5. the Resolution Context is capable of resolving the distinctions asserted;
6. every repeated scope field is an equal projection or declared sub-scope;
7. no field contradicts another in the same interpretation.

An unresolved World existence Claim does not become established by appearing in a well-formed Structural Scope.

### S0.5 Canonical Structural Judgments

The structural conformance judgment is:

    Gamma |-S SATISFIES_STRUCTURE(configuration_c, profile_p, SS) => r

where:

- `Gamma` is the attributable declarations, Evidence, methods, and assumptions;
- `configuration_c` is the exact actual Structural Configuration under examination, or a Term purporting to denote one;
- `profile_p` is the exact Structural Profile whose criteria are tested;
- `SS` is one well-formed Structural Scope;
- `r` is exactly one of:

      CONFORMS
      DOES_NOT_CONFORM
      UNRESOLVED
      ILL_TYPED

The result grammar is inherited from Ontology.

`SATISFIES_STRUCTURE` presupposes separately established existence and realization of the target in the exact Structural Scope. `CONFORMS` never creates that actuality. `DOES_NOT_CONFORM` does not make an actual arrangement disappear. `UNRESOLVED` is not a negative result. `ILL_TYPED` identifies a defect in the judgment, not in Reality.

The application judgment is separate:

    STRUCTURE_APPLIES(profile_p, target_x, SS)
      -> APPLICABLE | NOT_APPLICABLE | UNRESOLVED | ILL_TYPED

`SATISFIES_STRUCTURE` is evaluated only where `STRUCTURE_APPLIES = APPLICABLE`. A Profile that does not apply cannot be used to declare the target defective.

### S0.6 Registered Operators

| Operator | Owner | Result | Question |
|---|---|---|---|
| `WF_STRUCTURAL_SCOPE(SS)` | S0 | `WELL_FORMED \| ILL_TYPED` | Is the structural application boundary internally typed? |
| `WF_STRUCTURAL_PROFILE(p)` | S1 | `WELL_FORMED \| ILL_TYPED` | Is the represented Profile internally typed, Scope-bound, and explicit about constraints and coverage? |
| `WF_STRUCTURE_DESCRIPTION(sd)` | S1 | `WELL_FORMED \| ILL_TYPED` | Is the represented Description internally typed and explicit about coverage? |
| `STRUCTURE_APPLIES(p, x, SS)` | S1 | `APPLICABLE \| NOT_APPLICABLE \| UNRESOLVED \| ILL_TYPED` | Does Profile p apply to x in this exact Scope? |
| `SATISFIES_STRUCTURE(c, p, SS)` | S9 | Ontology conformance result | Does actual Configuration c satisfy Profile p? |
| `STRUCTURAL_EQUIVALENCE(c1, c2, criterion, SS)` | S9 | `EQUIVALENT \| NOT_EQUIVALENT \| UNRESOLVED \| ILL_TYPED` | Are two configurations interchangeable under the stated structural criterion? |
| `PATH_ASSESSMENT(a, b, path_kind, topology, SS)` | S4 | `PATH_ESTABLISHED \| NO_PATH_ESTABLISHED \| UNRESOLVED \| ILL_TYPED` | Does the exact topology support a path under the stated kind? |
| `REACHABILITY(a, b, topology, criterion, SS)` | S4 | `REACHABLE \| NOT_REACHABLE \| UNRESOLVED \| ILL_TYPED` | Is b reachable from a under the exact topology and criterion? |
| `CENTRALITY_ASSESSMENT(x, configuration, criterion, SS)` | S4 | `CENTRAL \| NOT_CENTRAL \| UNRESOLVED \| ILL_TYPED` | Does bearer or represented position x satisfy the exact centrality criterion? |
| `INTERFACE_POSTURE(i, crossing_kind, SS)` | S5 | `InterfacePostureEvaluation` | What structural accessibility posture does an Interface have for the represented crossing kind? |
| `STRUCTURAL_PLACEMENT(x, SS)` | S10 | `PLACED \| PLACEMENT_UNRESOLVED \| ILL_TYPED` | Can the observed arrangement be located without forcing a known topology? |

These operators evaluate Claims. None creates its target. Actual relation Claims use the structural Relation Kinds declared later through `RELATES`.

### S0.7 Structural Relation Declaration Schema

Every Relation Kind owned by Structure must have a declaration equivalent to:

    STRUCTURAL_RELATION R {
      arity: n
      relata: [P1: BearerKind1, ..., Pn: BearerKindN]
      qualifiers: [
        {
          value_space
          binding:
              ATTRIBUTION(attribute_kind)
            | SIDE_JUDGMENT(judgment_template)
          required: YES | NO
        }
      ] | NONE
      context: StructuralScope
      ordered_positions: YES | NO | PARTIAL
      direction: DIRECTED | UNDIRECTED | MIXED | NOT_APPLICABLE
      symmetry: YES | NO | CONDITIONAL | NOT_APPLICABLE
      reflexivity: YES | NO | CONDITIONAL | NOT_APPLICABLE
      transitivity: YES | NO | CONDITIONAL | NOT_APPLICABLE
      inverse: R2 | NONE | CONDITIONAL | OPEN
      composition_rule: rule | NONE | OPEN
      realization_criterion: Criterion | OPEN
      non_entailments: [...]
      owner: structure_module
    }

`arity` counts actual relata only. Every actual relatum conforms to `Bearer` or a declared Bearer subtype. Criteria, modes, Profiles, topologies, Roles, Structural Scopes, and other parameters are qualifiers or context unless independently declared as actual relata. A Term in a Claim purports to denote a bearer; the sign itself does not become the relatum.

The existential projection of a structural Relation Kind is:

    claim_R := attributable Claim that
      R(actual_bearer_relata; typed_qualifiers; SS) obtains

    j_R := Judgment evaluating or expressing claim_R in SS

    SCOPED_TO(j_R, SS)

    exists ri [
      EXISTS(ri, SS.domain, SS.interval)
      and RELATES(ri, R, [actual_bearer_relata], SS.domain, SS.interval)
      and for each qualifier q_k its declared binding holds:
            ATTRIBUTION(declared_attribute_k, ri, q_k,
                        SS.domain, SS.interval, SS.resolution_context)
        or  SIDE_JUDGMENT(declared_judgment_template_k[q_k])
    ]

Each qualifier declaration must name either its Property or Status attribute and compatible value-space, or its exact typed side-judgment template. `SCOPED_TO` applies to `j_R`, not silently to the Relation Instance or raw Claim. `EXISTS` plus the well-typed `RELATES` instance supplies the RelationInstance commitment.

An omitted algebraic property is `OPEN`, never silently true. Relation names do not determine transitivity, symmetry, exclusivity, or acyclicity. A Profile may narrow a base declaration; it may not change its arity, actual relatum positions, or base meaning.

### S0.8 Structural Result and Error Discipline

The Ontology Error Model remains applicable. Structure adds these exact reason codes:

- `WORLD_SCOPE_MISMATCH` — the World asserted by the target, Description, Profile, and Structural Scope cannot be reconciled;
- `CONFIGURATION_RECORD_CAST` — a Record or Snapshot is treated as the actual Configuration;
- `PROFILE_ACTUALITY_CAST` — a Profile or Architecture Specification is treated as an actual arrangement;
- `POSITION_BEARER_CAST` — a position, slot, Address, or Role is treated as its occupant;
- `OCCUPANCY_ERROR` — a bearer occupies an incompatible or cardinality-exhausted position;
- `PARTHOOD_MODE_ERROR` — a composition Claim lacks a mode or exports an algebra from another mode;
- `CARDINALITY_ERROR` — an exact multiplicity constraint is violated;
- `DANGLING_ENDPOINT` — a Link or Interface endpoint required by the Profile lacks a typed attachment;
- `BOUNDARY_SIDE_ERROR` — an interior, exterior, or side Claim is not grounded in one boundary criterion;
- `INTERFACE_MISMATCH` — connected interface positions have incompatible typed conditions without a declared mapping;
- `PROHIBITED_OVERLAP` — an overlap occurs where the applicable Profile excludes it;
- `DECLARED_ACYCLICITY_VIOLATION` — a cycle occurs in a relation explicitly declared acyclic;
- `CONSTRAINT_CONFLICT` — applicable structural constraints cannot all be satisfied and no explicit priority or composition rule resolves them;
- `TOPOLOGY_CAST` — a Graph, schema, or topology Description is treated as actual topology;
- `RELATIONAL_SYSTEM_CAST` — a Relation, overlap, interface, or mediator Role is treated as an independently realized Relational System without satisfying its criteria;
- `SHARED_STATE_WITHOUT_BEARER` — a shared State is asserted without an established bearer capable of bearing it;
- `INTERFACE_EFFECT_CAST` — structural accessibility is treated as an actual crossing, Effect, Admission, Permission, or Authorization;
- `EXHAUSTIVITY_ERROR` — absence is inferred from a Description whose relevant inventory is not exact and closed;
- `CROSS_SCALE_CAST` — a structure resolved at one scale is exported to another without a mapping preserving the asserted distinction;
- `LAYER_OWNERSHIP_ERROR` — a structural expression claims a result owned by Constitution, Dynamics, Continuity, or Emergence.

The reason code does not select a result by its name alone. Result mapping is exact:

- malformed grammar, category casts, ambiguous operators, incompatible argument kinds, or Scope defects -> `ILL_TYPED`;
- an established violation of an applicable Profile by an actual Configuration -> `DOES_NOT_CONFORM`;
- unresolved conflict, coverage, Mapping, realization, or Evidence after applicability and actuality are established -> `UNRESOLVED`;
- an established target-kind, predicate, or Scope mismatch in Profile application -> `NOT_APPLICABLE`;
- a reason concerning actuality before the conformance gate -> the applicable actuality posture and `ConformanceEvaluation = NOT_EVALUATED`.

A reason record remains representable as a Fact about an expression, Description, Profile, implementation, or Configuration. Rejecting the Claim does not erase the underlying Occurrence or arrangement.

### S0.9 Structural Invariants

Every module preserves these invariants:

1. **Actuality separation** — Configuration, Profile, Description, Record, and realizer remain distinct.
2. **Relational separation** — Relation Instance, relata, mediator, shared surface, and Relational System remain distinct.
3. **Position separation** — Role, slot, position, Address, Identifier, and occupant remain distinct.
4. **No topology privilege** — centrality, connectivity, containment, origin, operation, or infrastructural dependency creates no Authority, Ownership, rank, entitlement, or sovereignty.
5. **No property inheritance by placement** — a Property, Status, Agency, responsibility, or Authority of a part does not transfer to a whole, and vice versa, without a separately typed relation.
6. **No passage from possibility** — Interface, Port, Aperture, Channel, path, reachability, or Capacity does not prove crossing, delivery, influence, execution, or Effect.
7. **No persistence from snapshot** — a Configuration at one Interval does not establish the same Configuration or bearer at another.
8. **No emergence from shape** — arrangement alone does not establish life, Agency, Experience, intelligence, meaning, adaptation, coherence, or value.
9. **No universal spatialization** — non-spatial arrangements are not forced into physical location, distance, enclosure, or geometry.
10. **Open-world coverage** — omission does not establish absence unless exact closure conditions are satisfied.

---

## S1 — Configuration and Structural Forms

### S1.1 Structural Configuration

A **Structural Configuration** is an actual Configuration in which selected bearers, Relation Instances, boundaries, interfaces, positions, and constraints are arranged within one Structural Scope.

    KIND StructuralConfiguration {
      parents: [Configuration, State]
      category: State
      necessary_conditions: [actual selected arrangement; one StructuralScope; attributable bearer or World; realization basis]
      sufficient_conditions: [existence and arrangement relations established under the declared criterion]
      properties_and_value_spaces: [coverage: CoveragePosture; realization_basis: RealizationBasis]
      admissible_bearers: [Bearer, System, World, RelationInstance]
      relation_signatures: [RELATES, REALIZES]
      realization_criterion: arrangement-specific Criterion
      scope: StructuralScope
      disjoint_with: OPEN
      overlaps: [State]
      epistemic_statuses: [Observed, Measured, Inferred, Established, Disputed, Unknown]
      non_equivalences: [StructuralConfiguration / StructureDescription; StructuralConfiguration / StructuralProfile; StructuralConfiguration / StructureRecord]
      later_layer_boundaries: [change belongs to Dynamics; persistence belongs to Continuity]
      owner: S1
    }

A Structural Configuration selects relevant detail. It is not the complete essence of its bearer. Different established Configuration selections may concern the same bearer at different scales or under different structural questions without contradiction.

A Configuration is indexed to an Interval. If its arrangement changes inside that Interval at the asserted Resolution, the Claim must segment the Interval or use a dynamic account later. Structure does not hide reconfiguration inside one static assertion.

### S1.2 Structural Profile

A **Structural Profile** is a Representation declaring the admissible, required, optional, excluded, or unresolved features of a class of Structural Configurations.

    KIND StructuralProfile {
      parents: [Representation]
      category: Pattern
      necessary_conditions: [target class; StructuralScope; typed constraints; application criterion; coverage posture; RepresentationBinding]
      sufficient_conditions: [WF_STRUCTURAL_PROFILE returns WELL_FORMED for its exact represented form]
      properties_and_value_spaces: [version_or_form_reference; application_criterion; constraints; coverage]
      admissible_bearers: NOT_APPLICABLE
      relation_signatures: [REPRESENTS]
      realization_criterion: represented Profile physically encoded in a recoverable carrier
      scope: StructuralScope
      disjoint_with: OPEN
      overlaps: [Model, Record]
      epistemic_statuses: [Posited, Hypothetical, Established, Disputed]
      non_equivalences: [StructuralProfile / StructuralConfiguration; StructuralProfile / constitutional Rule]
      later_layer_boundaries: [transition rules excluded]
      owner: S1
    }

A Profile can be:

- **descriptive** — stating criteria used to describe an arrangement;
- **classificatory** — stating conditions for a structural class;
- **prescriptive within conformance** — stating required or excluded arrangement for a technical or formal target;
- **open** — leaving declared structural positions unresolved.

“Required” or “excluded” inside a Structural Profile is a conformance condition. It is not a constitutional Obligation or Prohibition, and it does not make a physical alternative impossible.

    WF_STRUCTURAL_PROFILE(p)
      -> WELL_FORMED | ILL_TYPED

Well-formedness requires an exact target class, Structural Scope, application criterion, typed constraints, coverage posture, and Representation/realization basis. It does not establish that any Configuration conforms to the Profile or that any target exists.

### S1.3 Structure Description, Snapshot, and Record

A **Structure Description** is a Representation that makes Claims about an actual or possible arrangement.

A **Structure Snapshot** is a Structure Description indexed to an as-of point or bounded Interval. It may include measurements, observations, and inferred relations.

A **Structure Record** is a persistent Representation preserving a Structure Description, its provenance, method, Scope, and result.

These remain distinct from:

- the Configuration described;
- the bearer whose Configuration it is;
- the Profile used to test it;
- the physical carrier realizing the Description or Record;
- any later execution using the Record.

A well-formed Record may describe a nonconforming Configuration. An authentic Record may be false or incomplete. A structurally conforming Record format does not establish that its content is true.

Every Structure-owned Representation subtype—including every Profile, Description, Model, Mapping, and Record—must expose:

    RepresentationBinding {
      representation: Representation
      carrier_or_vehicle: Bearer | nonempty [Bearer]
      represented_content: Term | Claim | Representation
      target_position: ReferenceTargetPosition
      mapping_or_interpretive_practice: Mapping | Process | Representation
      domain: Domain
      interval: Interval
      use_conditions: [Criterion | Claim | Representation] | NONE | OPEN
    }

and the exact imported relation:

    REPRESENTS(representation, target_position,
               mapping_or_interpretive_practice, domain, interval)

The carrier must physically realize the Representation under an attributable criterion. `target_term` and represented content never replace `target_position`; an external Referent, where established, occupies that position through the imported reference-target occupancy relation.

### S1.4 Description Record Projection

Every exact Structure Description must make this projection recoverable:

    StructureDescription {
      description_id: Identifier | NONE
      representation_binding: RepresentationBinding
      target_term: Term
      structural_scope: StructuralScope
      configuration_claim: Claim
      bearers: [Bearer] | NONE | OPEN
      relation_instances: [RelationInstance] | NONE | OPEN
      regions: [StructuralRegion] | NONE | OPEN
      boundaries: [Boundary] | NONE | OPEN
      interfaces: [Interface] | NONE | OPEN
      positions_and_occupancies: [StructuralSlot | Claim] | NONE | OPEN
      systems_and_environments: [System | Environment] | NONE | OPEN
      applicable_constraints: [Constraint | Representation] | NONE | OPEN
      realization_basis: RealizationBasis | OPEN
      coverage: CoveragePosture
      evidence_claims_or_records: [Claim | Record] | NONE | OPEN
      method_description: Representation | OPEN
      provenance_claims_or_records: [Claim | Record] | NONE | OPEN
      unresolved: [Claim | Term] | NONE
    }

`NONE`, `OPEN`, and omission are not interchangeable. A required field cannot be omitted. `OPEN` records unresolved content. `NONE` asserts a known empty position only where the coverage basis permits that assertion.

    WF_STRUCTURE_DESCRIPTION(sd)
      -> WELL_FORMED | ILL_TYPED

Well-formedness requires typed fields and internal consistency. It does not establish the Description’s truth, completeness, or target existence.

### S1.5 Coverage Posture

Every inventory, map, decomposition, boundary account, topology, and Description declares exactly one **Coverage Posture**:

    CoveragePosture :=
        OPEN
      | PARTIAL {
          covered_dimensions
          included_region_or_population
          method
        }
      | CLOSED_FOR {
          exact_dimensions
          exact_region_or_population
          interval
          resolution_context
          closure_criterion
          verification_method
        }
      | COVERAGE_UNRESOLVED

- `OPEN` means no completeness Claim is made.
- `PARTIAL` states exactly what is covered without implying completeness elsewhere.
- `CLOSED_FOR` permits exhaustive inference only for the named dimensions, population, Interval, Resolution Context, and criterion.
- `COVERAGE_UNRESOLVED` means the reach of the Description is not presently established.

Absence may be inferred from an inventory only where it is `CLOSED_FOR` the exact asserted absence and the closure Evidence is adequate. Omission from `OPEN`, `PARTIAL`, or unresolved coverage yields `UNRESOLVED`.

### S1.6 Realization Basis

Every Claim that a structure is actual must expose:

    DistributedRealizerBasis {
      bearers: nonempty [Bearer]
      organization_relations: nonempty [RelationInstance]
      joint_realization_criterion
      isolated_component_sufficiency:
        per_bearer: Mapping(Bearer -> INSUFFICIENT | UNRESOLVED)
      evidence_claims_or_records: [Claim | Record] | NONE | OPEN
    }

    RealizerPositionKind :=
        SINGLE { bearer: Bearer }
      | DISTRIBUTED { basis: DistributedRealizerBasis }

    RealizationBasis {
      realizer_position: RealizerPositionKind
      realized_target: Term whose owning Kind admits realization
      structural_scope: StructuralScope
      criterion
      evidence_claims_or_records: [Claim | Record] | NONE | OPEN
      result:
          REALIZATION_ESTABLISHED
        | NONREALIZATION_ESTABLISHED
        | REALIZATION_UNRESOLVED
        | ILL_TYPED
    }

This record represents a `REALIZES` Claim; it does not replace the imported relation.

A `DistributedRealizerBasis` records the actual plural bearer position admitted by Ontology’s Distributed Realization grammar. The Record may supply information participating in an Evidence relation for that position; it is not Evidence by identity, is not itself the realizer, and does not create an Aggregate, Whole, System, shared State bearer, or collective Agent. The `REALIZES` Claim must expose the actual bearers, their organizing Relation Instances, and the joint criterion.

For `DISTRIBUTED`, every isolated bearer must be established `INSUFFICIENT` under the joint criterion before the result can be `REALIZATION_ESTABLISHED`. Any `UNRESOLVED` component blocks that result. If one bearer is sufficient, it is represented as an alternative `SINGLE` basis rather than as a conforming Distributed Realization.

A well-formed possible arrangement, formal Graph, executable specification, deployed program, or constitutional recognition does not establish an actual structure without a sufficient Realization Basis.

### S1.7 Profile Composition

Profiles may be composed only through an exact **Profile Composition Record**:

    ProfileCompositionRecord {
      profiles: nonempty [StructuralProfile]
      target_kind_or_predicate
      shared_structural_scope: StructuralScope
      scope_partition_or_overlap
      priority_rules: [Representation] | NONE | OPEN
      conflict_rules: [Representation] | NONE | OPEN
      preserved_constraints: [Constraint | Representation] | NONE | OPEN
      overridden_constraints: [Constraint | Representation] | NONE
      unresolved: [Claim | Term] | NONE
    }

No Profile silently overrides another by recency, specificity, filename, implementation dominance, or wider use. Current, applicable constitutional Law may require use of a Profile only through every required valid Authority, Constitutional Act, Effect, Scope, and Subject Binding gate; that normative result does not change the Profile's structural content.

### S1.8 Profile Application

    STRUCTURE_APPLIES(profile_p, target_x, SS)
      -> APPLICABLE | NOT_APPLICABLE | UNRESOLVED | ILL_TYPED

`APPLICABLE` requires:

1. p is a well-formed Structural Profile;
2. x conforms to the Profile’s exact declared target Kind or satisfies its exact target predicate; Role occupancy is not used as class membership;
3. p’s World context, Domain, Interval, Resolution Context, scale, and further conditions include the exact application;
4. the Profile version or exact Form is identified;
5. no higher-order Profile composition rule excludes it.

No Profile applies globally by omission. Applicability at one scale or World does not travel to another.

### S1.9 Structural Position and Slot

A **Structural Position** is a qualified Role defined by a relation or arrangement within a Configuration.

It exists as a Role only when occupied. Its bearer may change without the position name proving continuity of either bearer or whole.

A **Structural Slot** is a qualified `ReferenceTargetPosition` physically realized within a Structural Profile. It can represent a position that is unoccupied, required, optional, repeatable, or excluded. A Slot is a Bearer in the Reference relation, but it is not the represented occupant and does not itself become a Role. Where a compatible actual bearer occupies the represented position in an actual Configuration, the bearer occupies the corresponding Structural Position Role.

    OCCUPIES_STRUCTURAL_POSITION(bearer_b, position_p, configuration_c, SS)

is the Structure-owned display specialization:

    PLAYS(bearer_b, position_p, configuration_c, SS.interval)

where `position_p` conforms to Structural Position and the Role, Configuration, and Structural Scope are compatible. It does not introduce a separate relatum position for the Role and does not imply Ownership, identity, Participation, Agency, Authorization, or persistence.

### S1.10 Names Not Installed as Structural Primitives

This universal layer does not install `Locality`, `Kernel`, `Enclave`, `Centre`, `Actor`, `Participant`, or `Agent` as new structural primitives.

- `Agent` remains the Ontology Kind for a System realizing Agency.
- `Actor` remains the Role occupied by an Agent in relation to an attributable Action.
- `Participant` remains the Role occupied by a Bearer in an obtaining `PARTICIPATES` relation.
- A future L1 Profile may declare a `KernelRole` if it provides exact bearers, relation signatures, and criteria. Admission and state transition remain outside Structure; persistence of that bearer remains Continuity.
- Centre is treated in S4 only as a derived relation under an explicit criterion, never as a universal thing.
- An enclave-like arrangement, if ever required, must be expressed from Region, Boundary, containment, and Interface relations or declared as a fully typed qualified Profile. The word alone carries no force.

---

## S2 — Composition and Cardinality

### S2.1 Composition Is Typed

There is no admissible unqualified structural inference from “x is part of y.” Every composition Claim must state its mode, Scope, criterion, and algebra.

The canonical composition relation is:

    COMPOSES(part_x, whole_y, composition_mode_m, SS)

where `composition_mode_m` is one exact declared mode, not free text.

The initial modes are:

| Mode | Structural meaning | Explicit boundary |
|---|---|---|
| `MATERIAL_COMPOSITION` | x materially enters the physical composition of y | not formal Membership or Ownership |
| `CONSTITUTIVE_COMPOSITION` | x and typed relations enter the criterion by which y is constituted | not sufficient without the full constitutive criterion |
| `ORGANIZATIONAL_COMPOSITION` | x occupies an organized position in y | not necessarily material containment |
| `COMPUTATIONAL_COMPOSITION` | x enters an actual computational implementation arrangement | not Agent identity, represented composition, or execution |
| `FORMAL_COMPOSITION` | x is a formal subexpression or structural part of a formal object | not physical parthood by identity |
| `REPRESENTATIONAL_COMPOSITION` | x is a component of a Representation | not part of the represented Referent |

Set Membership, institutional Membership, spatial containment, data reference, implementation, realization, and Ownership are not composition modes unless a later qualified relation expressly maps them. They retain their own Kinds and relations.

### S2.2 Composition Algebra

Every composition mode declares:

    CompositionAlgebra {
      mode
      admissible_part_kinds
      admissible_whole_kinds
      reflexive: YES | NO | CONDITIONAL
      symmetric: YES | NO | CONDITIONAL
      transitive: YES | NO | CONDITIONAL
      asymmetric: YES | NO | CONDITIONAL
      acyclic: YES | NO | CONDITIONAL
      overlap_allowed: YES | NO | CONDITIONAL
      shared_part_allowed: YES | NO | CONDITIONAL
      composition_rule
      decomposition_rule: rule | OPEN
      resolution_context: projection of StructuralScope.resolution_context
    }

No algebraic property transfers from one mode to another. Even within one mode, transitivity may fail across scales or changing criteria.

### S2.3 Constituent and Component

A **Constituent** retains the Ontology meaning: it enters the composition of a whole.

A **Component** retains the Ontology meaning: it is a distinguishable part assigned a function or position within a System Description.

Structure declares:

    COMPONENT_OF(bearer_b, system_s, component_role_r, SS)

This is a directed, non-reflexive Relation Kind. It is transitive only where the applicable Profile expressly composes component hierarchies and preserves the component mode across each level.

Component Role does not make the bearer intrinsically that component. The same bearer may occupy Component Roles in more than one System. Shared components do not fuse those Systems, transfer their properties, or establish joint Ownership.

### S2.4 Subsystem

A **Subsystem** is a System that occupies an exact Component or Constituent relation within another System while retaining its own declared Boundary, components, relations, and System criterion.

    SUBSYSTEM_OF(system_a, system_b, mode_m, SS)

`SUBSYSTEM_OF` does not establish that a and b share identity, governance, complete State, persistence criterion, or environment. A nested System may remain partly open to more than one environment.

### S2.5 Aggregate and Collection

An Aggregate remains distinct from a System. Grouping bearers by proximity, label, query, ownership Claim, or common Profile does not establish integrated organization.

An inventory or Set Representation may enumerate bearers without constituting any whole. A structural whole requires an attributable constitutive criterion and relations sufficient for that Kind.

### S2.6 Cardinality and Multiplicity

A **Cardinality Constraint** is a Representation specifying permitted numbers of occupants, parts, endpoints, interfaces, relations, or other typed positions. It does not become an actual Constraint merely by being encoded.

    CardinalityConstraint {
      target_position_or_relation
      minimum: integer | OPEN
      maximum: integer | UNBOUNDED | OPEN
      exact: integer | NOT_APPLICABLE | OPEN
      counting_criterion
      distinctness_criterion
      StructuralScope
    }

`exact` cannot conflict with `minimum` or `maximum`. Counting requires an explicit distinctness criterion; multiple identifiers do not prove multiple bearers, and one identifier does not prove one bearer.

### S2.7 Optionality, Requirement, and Exclusion

Within a Structural Profile, a Slot or relation may be:

    REQUIRED | OPTIONAL | REPEATABLE | EXCLUDED | OPEN

These are Profile conditions only.

- `REQUIRED` does not create an Obligation.
- `EXCLUDED` does not create a Prohibition.
- `OPTIONAL` does not create Permission.
- `OPEN` does not mean optional or absent.

### S2.8 No Structural Inheritance by Composition

From any composition, component, or subsystem relation alone, none of the following follows:

- the whole has every Property or Status of the part;
- the part has every Property or Status of the whole;
- an Agent part makes the whole an Agent;
- an authorized component makes the whole authorized;
- a System’s Authority, Ownership, responsibility, or entitlement transfers to a component;
- a component’s identity or history becomes the whole’s;
- replacement preserves numerical identity;
- a shared component creates a jointly owned State or a merged System.

Every such result requires its own owning-layer relation and Evidence.

### S2.9 Decomposition and Resolution

One bearer may admit several well-typed decompositions under different criteria or scales. A biological organism, computational process, institution, and physical object may each expose different parts without those decompositions being interchangeable.

Two decompositions conflict only where they assert exact incompatible content in the same World, Domain, Interval, Resolution Context, composition mode, and interpretation.

A decomposition does not prove that its selected parts are fundamental or that omitted detail is absent.

---

## S3 — Region, Boundary, and Containment

### S3.1 Region

A **Structural Region** is a criterion-bounded selection within a World or Configuration.

    KIND StructuralRegion {
      parents: [Bearer]
      category: other
      necessary_conditions: [actual criterion-bounded selection; StructuralScope; declared extent or membership condition; existence or realization basis]
      sufficient_conditions: [the criterion resolves included, excluded, and boundary positions at the stated Resolution]
      properties_and_value_spaces: [extent_or_membership; boundary_posture; coverage]
      admissible_bearers: [StructuralRegion]
      relation_signatures: [IN_REGION, BOUNDS_REGION, OVERLAPS_REGION]
      realization_criterion: selection physically or organizationally realized where claimed actual
      scope: StructuralScope
      disjoint_with: OPEN
      overlaps: [Domain, SpacetimeRegion, Set, Representation]
      epistemic_statuses: [Observed, Measured, Inferred, Posited, Established, Disputed]
      non_equivalences: [StructuralRegion / Entity; StructuralRegion / System; StructuralRegion / container; StructuralRegion / jurisdiction]
      later_layer_boundaries: [crossing belongs to Dynamics]
      owner: S3
    }

A Structural Region is a Bearer of region and boundary relations. It is not thereby an Entity, System, container, owner, jurisdiction, or Agent. A merely proposed or represented region is carried by a qualified `ReferenceTargetPosition` or Description until actual existence or realization is established.

A physical enclosure, named zone, computational partition, formal subset, institutional area, and semantic domain may each be represented as regions under different criteria, but they are not one relation by name.

A **Spacetime Region** and **Location** retain their physical Ontology meanings. A generic Structural Region must not be used to make a non-spatial arrangement spatial.

### S3.2 Boundary and Boundary Sides

A Boundary retains the Ontology meaning: a distinction between inside and outside under a Criterion.

A **Boundary Side** is a Role occupied relative to one Boundary criterion:

    INTERIOR_SIDE | EXTERIOR_SIDE | BOUNDARY_POSITION

The sides are not free-standing regions unless a separate Region criterion constitutes them.

    BOUNDS_REGION(boundary_b, region_g, criterion_c, SS)

is a directed relation from a Boundary to the Region selected by its criterion. It is neither symmetric nor transitive.

Every boundary declaration must identify:

- the distinction resolved;
- the selected inside and outside conditions;
- Domain, Interval, and Resolution Context;
- whether the Boundary is physical, organizational, computational, informational, semantic, institutional, formal, or mixed;
- permeability or interface positions, if known;
- what remains unresolved.

A Boundary does not imply a wall, barrier, isolation, privacy, Ownership, jurisdiction, Agency, or prohibition.

### S3.3 Inclusion and Containment

The canonical region-inclusion relation is:

    IN_REGION(bearer_x, region_g, inclusion_mode_m, SS)

Initial modes are:

| Mode | Meaning | Does not imply |
|---|---|---|
| `SPATIAL_INCLUSION` | x occupies a physical Location within a spatial Region | componenthood or Ownership |
| `TOPOLOGICAL_INTERIOR` | x lies in an interior under a topology | metric distance or physical enclosure |
| `ORGANIZATIONAL_INCLUSION` | x is placed within an organizational Region of a Configuration | institutional Membership or submission |
| `COMPUTATIONAL_INCLUSION` | x is placed within a computational namespace, partition, or addressable Region | physical co-location or Permission |
| `REPRESENTATIONAL_INCLUSION` | a token, element, or mark lies within a Representation | inclusion of the Referent |
| `FORMAL_ELEMENT_INCLUSION` | a physically represented formal element occupies an element position in a represented formal Region | subset inclusion or physical containment of its Referent |

`CONTAINS_REGION(region_g, bearer_x, mode_m, SS)` is the declared inverse of `IN_REGION` for the same mode and Scope. No other inverse is inferred.

Region-to-Region inclusion is separately typed:

    SUBREGION_OF(region_a, region_b, subregion_mode, SS)

Formal subset inclusion uses `SUBREGION_OF` with `FORMAL_SUBSET_INCLUSION`; it is not element membership and does not inherit the element-inclusion signature.

`FORMAL_SUBSET_INCLUSION` is a built-in `SubregionMode` admissible only where both relata are actual Structural Regions interpreted in the same formal Domain under one exact membership semantics, Mapping, Resolution Context, and Criterion. It obtains exactly where every position admitted by region_a is admitted by region_b under that shared interpretation. It is directed and reflexive. Antisymmetry establishes only extensional equivalence under that interpretation and does not imply `SAME`. Transitivity requires preservation of the same Domain, membership semantics, Mapping, Resolution Context, and Criterion. It entails neither element membership, physical containment, componenthood, nor institutional Membership.

Containment is distinct from composition, Component Role, set Membership, institutional Membership, Ownership, realization, and implementation.

### S3.4 Overlap, Intersection, and Separation

    OVERLAPS_REGION(g1, g2, criterion, SS)

obtains where the two Regions share at least one resolved position under the criterion.

    INTERSECTS_AT(g1, g2, shared_region_or_position, criterion, SS)

names the resolved intersection where one is established.

The third actual relatum is a Structural Region or another actual Bearer occupying the resolved intersection position. A Structural Slot is admissible there only when the asserted relation concerns the physically realized Profile Representation itself. A Profile-only Slot never establishes that the Profile's target Regions intersect.

    STRUCTURALLY_SEPARATE(g1, g2, criterion, SS)

means no shared position is resolved under the stated criterion and closed coverage. It is not strict numerical Difference between the Region terms and does not prove causal, informational, institutional, or social isolation.

Overlap does not create:

- a jointly owned State;
- a new Entity;
- a Relational System;
- shared Identity;
- mutual Access;
- Authority or jurisdiction;
- fusion of the overlapping structures.

### S3.5 Nesting and Enclosure

    NESTED_IN(region_or_structure_a, region_or_structure_b, nesting_mode, SS)

is permitted only where the nesting mode declares its algebra and the two positions are compatible.

An **Enclosure** is a Region whose Boundary surrounds its selected interior under a declared topology. Enclosure does not establish impermeability.

An arrangement sometimes called an “enclave” can be represented as a Region nested within another Region and connected through declared Interfaces. This Form does not create `Enclave` as a universal Kind because the word otherwise collapses enclosure, access, jurisdiction, social status, and sovereignty. A later Profile may introduce the qualified term only by separating those relations.

### S3.6 Partition and Cover

A **Structural Cover** is an actual Pattern and Bearer of Regions that covers a target under one inclusion mode, Mapping, criterion, and Structural Scope.

    COVERS_STRUCTURALLY(cover_candidate_k, target_x,
                        inclusion_mode_m, mapping_M, criterion_c, SS)

has two actual relata: `cover_candidate_k: Bearer` and `target_x: Bearer`. The first relatum conforms to Structural Cover only where the relation obtains and all of these conditions are established:

1. k and x exist in SS;
2. k conforms to Pattern and exposes a nonempty exact cell list of actual Structural Regions;
3. every cell is related to the selected target extent under m, M, and c;
4. every target position or dimension named by c is included in at least one cell; and
5. coverage is `CLOSED_FOR` precisely those target dimensions, supported by proof or adequate Evidence.

Its Regions may overlap. A partial or candidate cell family is a Representation of possible cover, not an actual Structural Cover. “Cover” does not import set union into every physical, organizational, computational, or formal domain.

A **Structural Partition** is a Structural Cover whose cells are pairwise non-overlapping except for explicitly permitted shared Boundary positions, and whose coverage is `CLOSED_FOR` the partitioned dimensions.

    PartitionProfile {
      target
      cells
      cover_relation_claim
      inclusion_mode
      mapping
      inclusion_criterion
      overlap_rule
      coverage
      StructuralScope
    }

Partition cells need not be Entities or Systems. A partition in a Description does not physically divide its Referent unless a corresponding realized Boundary exists.

### S3.7 Boundary and Containment Invariants

1. Inside and outside are criterion-relative, not arbitrary.
2. A bearer may be spatially inside a System without being its Component.
3. A Component may be organizationally inside a System without being spatially enclosed by it.
4. A formal member may not be a physical part.
5. A represented box around tokens does not create an actual Boundary around their Referents.
6. A nested structure may connect directly to an environment outside its containing Region where the Profile permits it.
7. Containment at one Resolution does not settle containment at another.

---

## S4 — Position, Placement, and Topology

### S4.1 Position Is Relation-Relative

A Structural Position retains the S1 definition as a qualified Role within an arrangement. It must declare its position system.

Initial position systems include:

| Position system | Position is relative to | Required support |
|---|---|---|
| physical | spacetime and Reference Frame | Location or Spacetime Region |
| geometric | geometry and coordinates | coordinate Mapping and metric where distance is used |
| topological | adjacency, incidence, and open-set or equivalent structure | declared topology |
| ordered | an order relation | ordering criterion |
| network | Node, Link, endpoint, or routing arrangement | actual network or explicit Representation |
| computational | namespace, memory, address, process, or data structure | computational System and interpretation |
| formal | expression, Set, Graph, or other formal structure | formal interpretation |
| organizational | configured positions and relations in a System | organizational criterion |

One word such as “near,” “inside,” “above,” “first,” “central,” or “connected” is ill-typed unless its position system and criterion are recoverable.

### S4.2 Placement and Occupancy

    OCCUPIES_STRUCTURAL_POSITION(bearer_b, position_p, configuration_c, SS)

requires:

- p belongs to c or an applicable Profile of c;
- b satisfies the position’s admissible bearer conditions;
- cardinality remains satisfied;
- the occupancy is established during the Interval;
- the relation does not conflict with declared exclusions.

The bearer, position, occupancy Relation Instance, Identifier, and Address remain distinct.

A bearer may occupy multiple compatible positions. Multiple bearers may occupy a plural position where its cardinality permits. Position sharing does not fuse them.

### S4.3 Location, Coordinate, Address, and Orientation

A physical Location retains Ontology meaning. A coordinate is a represented value under a Coordinate System. An Address is a locator or routing Identifier within a System.

No one entails the others.

- a physical bearer can change coordinates without changing numerical Identity;
- two Addresses can route to one bearer;
- one Address can be reassigned;
- a virtual position is physically realized but is not thereby a physical Location identical to its carrier’s Location;
- a coordinate match does not prove co-location unless the Reference Frame and mapping match.

**Orientation** is a relation to a declared frame. **Distance** is a Quantity defined only where a metric or distance function is declared. A topology can exist without a metric.

### S4.4 Direction and Vector

Direction may be a structural property of an ordered relation, orientation, Link, or interface side.

A physical or mathematical **Vector** retains the Ontology meaning. An intention, preference, trajectory, semantic tendency, or “intent axis” is not a Vector by metaphor. It may be represented by a vector only where an explicit state space, basis, components, mapping, and interpretation are declared.

Such a vector Representation does not become the represented Intention or prove that Intention exists.

### S4.5 Adjacency, Incidence, Attachment, and Connection

These relations are distinct:

    ADJACENT_TO(a, b, topology_t, SS)
    INCIDENT_TO(link_bearer_l, endpoint_bearer_e, topology_t, SS)
    ATTACHED_TO(a, b, attachment_mode, SS)
    STRUCTURALLY_CONNECTED(a, b, connection_kind, topology_t, SS)

Each actual relatum is a Bearer—either an actual occupant or a qualified `ReferenceTargetPosition` such as a Structural Slot. A Role itself is never silently used as the relatum.

`ADJACENT_TO` may be symmetric under the topology. `INCIDENT_TO` is directed from a Link occupant or Link Slot to an endpoint occupant or Endpoint Slot. `ATTACHED_TO` declares a maintained structural connection under its mode. `STRUCTURALLY_CONNECTED` requires a declared Link, Relation Instance, Channel, shared component, or path rule.

Coupling retains its imported Ontology meaning. An actual Coupling is represented by a named Coupling Relation Instance asserted through the full `RELATES` form; this layer does not invent an undeclared `COUPLED_TO` shorthand.

Adjacency, incidence, proximity, attachment, and matching labels do not by themselves establish Interaction, causal influence, communication, trust, consensus, or Authority.

### S4.6 Structural Topology

A **Structural Topology** is a Pattern of positions and typed Relation Instances that declares which structural neighborhoods, incidences, connections, or separations obtain in one Structural Scope.

    KIND StructuralTopology {
      parents: [Pattern, Bearer]
      category: Pattern
      necessary_conditions: [positions; typed relations; StructuralScope; interpretation; coverage]
      sufficient_conditions: [relations and positions established under one compatible topology criterion]
      properties_and_value_spaces: [directionality; multiplicity; connectedness posture; metric posture; coverage]
      admissible_bearers: [StructuralConfiguration, System, World]
      relation_signatures: [ADJACENT_TO, INCIDENT_TO, STRUCTURALLY_CONNECTED]
      realization_criterion: topology-specific arrangement criterion
      scope: StructuralScope
      disjoint_with: OPEN
      overlaps: [Graph, Network, Pattern]
      epistemic_statuses: [Observed, Measured, Inferred, Posited, Established, Disputed]
      non_equivalences: [StructuralTopology / Graph; topological connection / causal reach]
      later_layer_boundaries: [traversal belongs to Dynamics]
      owner: S4
    }

A topology may be directed, undirected, mixed, finite, infinite, connected, disconnected, layered, cyclic, acyclic, nested, overlapping, distributed, higher-order, or unresolved. None is the universal default.

Reality is not restricted to pairwise Graph structure. An n-ary Relation Instance, hyperedge-like relation, field relation, distributed constraint, or higher-order organization may require more than ordinary nodes and binary edges.

### S4.7 Node, Link, Endpoint, and Hub Roles

A **Node Role** is a topology-relative Role occupied by a Bearer.

A **Link Role** may be occupied by:

- a Relation Instance;
- a Channel bearer;
- a physical connector;
- an Interface or shared component;
- a represented edge, where the topology itself is explicitly representational.

These alternatives must be typed. A Graph edge is not silently an actual Relation Instance or Channel.

An **Endpoint Role** is a terminus position of a directed or undirected Link. A **Hub Role** is a Node Role satisfying an explicit centrality, degree, routing, or aggregation criterion.

No Node, Link, endpoint, or hub Role implies Agency, personhood, Participation, control, importance, or Authority.

### S4.8 Path and Reachability

A **Structural Path** is an ordered Pattern of actual compatible position occupants and Link bearers under an actual topology. A Slot sequence inside a Profile represents a possible or required path shape; it is not an actual path in the Profile’s Referent.

    PATH_ASSESSMENT(source, target, path_kind, topology, SS)
      -> PATH_ESTABLISHED | NO_PATH_ESTABLISHED | UNRESOLVED | ILL_TYPED

is a derived structural assessment, not a `RELATES` projection. `PATH_ESTABLISHED` requires an actual path in the exact target Scope and does not assert that anything traverses it. A Profile-only Slot sequence can support only a Claim about represented path potential; without actual topology Evidence, target-path actuality remains `UNRESOLVED`.

    REACHABILITY(source, target, topology, reachability_criterion, SS)
      -> REACHABLE | NOT_REACHABLE | UNRESOLVED | ILL_TYPED

is also a derived assessment. `REACHABLE` means a compliant structural path is established under the criterion. It does not establish:

- physical causal reach;
- Channel availability at a later time;
- successful routing or delivery;
- Access, Permission, Admission, or Authorization;
- absence of capacity, cost, or failure constraints;
- trust, consensus, or shared State.

Actual traversal and its outcome belong to Dynamics.

### S4.9 Centrality and Centre

**Centrality** is a derived Property under an explicit geometry, topology, distribution, function, or comparison rule.

    CENTRALITY_ASSESSMENT(bearer_or_slot_x, configuration_c, criterion_k, SS)
      -> CENTRAL | NOT_CENTRAL | UNRESOLVED | ILL_TYPED

may obtain under one criterion and fail under another.

A **Central Slot** is a Structural Slot whose represented position satisfies the declared centrality criterion. A **Centre Role** exists only where an actual Bearer occupies a compatible Structural Position and that bearer satisfies the criterion. There may be one Centre Role, several, none, or an unresolved set. A Slot is not promoted into a Role merely because it is central.

Centre is therefore not:

- a universal primitive;
- the bearer’s identity;
- a synonym for Kernel, Self, Agent, Participant, or World;
- proof of control, Authority, Ownership, priority, persistence, or greater value.

This layer uses no generic `Locality` term. Exact position and relation types make it unnecessary.

### S4.10 Topology Families

Chain, ring, tree, directed acyclic graph, hierarchy, lattice, star, mesh, bipartite, layered, cellular, field-coupled, and other topology names are reusable Pattern classifications only where their defining relations are declared.

A hierarchy is an ordered structural relation. It does not imply domination or Authority. A layer is a partition or ordered Region under a declared criterion. It does not imply higher worth, precedence, or causal direction.

A cycle is not an error unless the applicable relation is declared acyclic. A disconnected topology is not defective unless connectivity is required by the applicable Profile.

### S4.11 Time-Indexed Topology

Structure may state topology at Interval I and separately at Interval I2. The reconfiguration between them, its cause, ordering, and execution belong to Dynamics. Whether the same System persists through that change belongs to Continuity.

---

## S5 — Interface, Seam, Port, and Aperture

### S5.1 Interface Profile

An Interface retains the Ontology meaning: a typed Boundary relation through which specified differences may cross, be translated, or produce effects.

Structure declares an **Interface Profile** with:

    InterfaceProfile {
      interface_term
      boundary
      side_positions
      source_domain_or_region
      receiving_domain_or_region
      admissible_bearer_signal_or_content_kinds
      port_positions
      direction_posture
      channel_or_coupling_basis: term | NONE | OPEN
      mapping_or_translation: term | NONE | OPEN
      preserved_distinctions
      permitted_loss_or_approximation
      excluded_or_unmapped_distinctions
      structural_constraints
      coverage
      StructuralScope
    }

The words “admissible” and “excluded” here refer to structural compatibility with the Profile. They do not create Permission or Prohibition.

### S5.2 Boundary Side and Endpoint Fit

Every Interface must attach to typed side or endpoint positions. A connection between two Interfaces requires:

- compatible endpoint or Port kinds;
- compatible orientation or an explicit adaptor;
- compatible cardinality;
- a declared Mapping where representations or encodings differ;
- preservation and loss conditions;
- no unresolved prohibited overlap under the applicable Profile.

Physical contact alone does not establish a working Interface. Matching schemas alone do not establish an actual physical Channel.

### S5.3 Port Role

A **Port Role** is a Structural Position on an Interface through which a compatible bearer, Channel, signal carrier, or connector may be attached.

    OCCUPIES_PORT(bearer_b, port_p, interface_i, SS)

is the typed display of `PLAYS(bearer_b, port_p, interface_i, SS.interval)` where `port_p` conforms to Port Role. An unoccupied represented Port is a **Port Slot**—a Structural Slot whose represented Role kind is Port Role—not an existing Port Role without a bearer. Port compatibility or represented availability does not establish connection, Access, Permission, execution, or passage.

An Address may route toward a Port but is not the Port or its occupant.

### S5.4 Aperture Family

**Aperture** is a display family, not one Kind forced to be both Region and Role:

    Aperture := ApertureRegion | ApertureRole | ApertureSlot

- an **Aperture Region** is a qualified Structural Region actually located on an Interface, with source and receiving sides, eligible crossing kinds, criterion, and Structural Scope;
- an **Aperture Role** is a Structural Position occupied through `PLAYS` by a compatible Interface component or bearer at an Aperture Region;
- an **Aperture Slot** is a Structural Slot representing an unoccupied or candidate Aperture Region or Aperture Role in a Profile.

    ApertureRegion SUBTYPE_OF StructuralRegion
    ApertureRole SUBTYPE_OF StructuralPosition
    ApertureSlot := StructuralSlot {
      represented_target_kind: ApertureRegion | ApertureRole
    }

    APERTURE_ON(aperture_region_a, interface_i, SS)

has exactly two actual Bearer relata: the Aperture Region and the Interface Relation Instance. Aperture Role occupancy uses `PLAYS`; the Role is not an additional relatum. A separately identified Interface Profile may constrain direction, eligible crossing kinds, Mapping, and other typed conditions through its own application and conformance judgment; none is a hidden qualifier of `APERTURE_ON`.

An Aperture is not:

- the crossing Occurrence;
- a gateway Authority;
- an Admission Act or Admission Relation;
- proof of Agency, Intention, Consent, or identity;
- Permission or Access;
- a guarantee that the receiving side accepts, processes, or retains anything.

An Aperture’s accessibility posture is assessed through `INTERFACE_POSTURE`; it is not another normative or dynamic state hidden in the word “open.” Opening, closing, holding, qualifying, accepting, refusing, or traversing are Occurrences owned downstream.

### S5.5 Gate and Filter Roles

A **Gate Role** is a structural Role occupied by a bearer or Interface component placed so that crossing attempts encounter it.

A **Filter Role** is a structural Role occupied by a bearer, Profile, or implemented mechanism positioned to distinguish structurally compatible inputs under a Criterion.

These Roles do not prove that filtering occurs, that its Criterion establishes interior truth, or that a Decision has constitutional force. Dynamics owns actual evaluate, accept, hold, refuse, and cross Occurrences. Constitution owns any normative Admission or Refusal Decision or Status, its validity, Review, Remedy, Consent relation, and constitutional Effect where applicable.

### S5.6 Seam

A Seam retains the Ontology meaning and must declare:

- source and receiving Domains;
- admissible inputs;
- the Mapping or transformation;
- preserved distinctions;
- permitted loss;
- distinctions that may not cross;
- the site at which consequence would occur.

Structure owns the arrangement of the Seam and its typed attachment positions. Dynamics owns actual transformation or passage. Constitution owns any normative validity at the crossing. A successful format conversion does not establish semantic equivalence, truth, or lawful transfer.

### S5.7 Channel, Interface, and Path

A Channel, Interface, Port, Aperture, and Path are non-equivalent:

- a **Channel** is a bounded path or coupling through which a carrier can pass;
- an **Interface** is the typed Boundary relation governing encounter across sides;
- a **Port** is an attachment position on an Interface;
- an **Aperture** names its typed Region, occupied Role, or represented Slot on an Interface;
- a **Path** is a compatible sequence in a topology.

One arrangement may instantiate several of these at once. None is inferred from another by name.

### S5.8 Interface Composition

Interfaces may compose only through a declared composition rule. If Interface A maps Domain D1 to D2 and Interface B maps D2 to D3, no direct D1-to-D3 Seam follows unless:

1. the output conditions of A satisfy the input conditions of B;
2. preserved distinctions remain sufficient;
3. losses and exclusions are composed explicitly;
4. Scope, Interval, and Resolution Context remain compatible;
5. the resulting mapping is declared.

Composition of structural compatibility does not prove end-to-end execution.

### S5.9 Interface Failure Is Not Erasure

A missing endpoint, incompatible Port, unresolved Mapping, or closed Aperture can make an arrangement nonconforming. The actual components and attempted connections remain real. Runtime refusal, rollback, retry, teardown, or retained trace belong to Dynamics.

---

## S6 — Systems, Worlds, Environments, and Architecture

### S6.1 System Structure

A System retains the Ontology meaning: an Entity or organized collection of Entities and Relations treated as a whole under an explicit Boundary.

Structure does not redefine System. It declares the arrangement required to evaluate a particular System Claim.

Every **System Structure Description** must make this projection recoverable:

    SystemStructureDescription {
      system_term
      structural_scope
      system_kind_claim
      constitutive_criterion
      realization_basis
      boundary
      component_roles_and_occupants
      relation_instances
      state_bearers
      interfaces_and_apertures
      environment_relation
      structural_constraints
      profile
      coverage
      evidence
      unresolved
    }

A component list is insufficient. A System Claim requires the organization and relations by which the whole is distinguished from an Aggregate.

### S6.2 System Boundary and State Bearers

A System’s Boundary must be stated under one criterion and Resolution Context. It may be distributed, porous, changing across separately indexed Configurations, or partly unresolved.

A System may have State only through established State bearers and selected Properties and Relations. A database, ledger, shared Interface, Relation Instance, or component Record is not automatically the complete System State.

Where a State is called “shared,” the Claim must identify:

1. the bearer capable of bearing that State;
2. the constitutive relation by which the State belongs to that bearer;
3. access and update relations, if relevant;
4. Scope and Interval;
5. whether “shared” means jointly represented, jointly accessible, jointly maintained, or actually borne by one Relational System.

Without that basis, `SHARED_STATE_WITHOUT_BEARER` applies.

### S6.3 Environment Relation

Environment retains the Ontology meaning: the relevant remainder relative to a stated System Boundary.

    ENVIRONMENT_OF(environment_region_e, system_s, criterion_c, SS)

is a directed, non-symmetric, criterion-relative relation. One bearer may be part of the Environment of several Systems or a Component of one System and part of another’s Environment.

Environment does not mean irrelevant, uncontrolled, external to Reality, or outside every World. A System may exchange with its Environment through Interfaces without ceasing to be structurally distinguishable.

Actual exchange belongs to Dynamics.

### S6.4 Open, Closed, and Isolated Structures

Thermodynamic openness and closure retain their physical Ontology definitions. Structure may assess a domain-qualified boundary posture only through:

    InterfacePostureValue :=
        OPEN_FOR | CLOSED_FOR | PARTIAL_FOR | UNRESOLVED

    InterfacePostureEvaluation :=
        InterfacePostureValue | ILL_TYPED

    INTERFACE_POSTURE(interface_i, crossing_kind_k, SS)
      -> InterfacePostureEvaluation

This is a structural Status assessment about specified kinds and conditions, not a `RELATES` projection. It is not thermodynamic closure unless the thermodynamic criteria are satisfied, and it creates no normative Access, Permission, or Prohibition.

No System is simply “open” or “closed” without stating what may cross, through which Boundary, in which direction, and at what Resolution.

### S6.5 World Structure

A World retains the Ontology meaning: a bounded Domain of Entities, Relations, Occurrences, and rules of interaction treated as a coherent Environment.

Structure may declare a **World Structure Description**:

    WorldStructureDescription {
      world_term
      structural_scope
      realization_basis
      boundary_or_boundary_posture
      included_regions
      bearers_and_systems
      relation_instances
      interfaces_to_other_worlds_or_environments
      structural_profiles
      referenced_dynamic_rule_sets
      referenced_constitutional_sources
      coverage
      unresolved
    }

Referenced dynamic rules and constitutional sources remain owned by their layers. Listing them does not execute or validate them, make a dynamic rule operative or applicable, or make a constitutional source adopted, Current, or applicable.

A World Boundary creates no automatic sovereignty, jurisdiction, Ownership, truth, Agency, or isolation. Worlds may overlap, nest, share bearers, share Infrastructure, or remain partly disconnected where their exact relations permit it.

### S6.6 World Overlap and Multi-World Placement

One bearer may be structurally placed in several Worlds under different Relations and Scopes. This does not duplicate the bearer or merge those Worlds.

A shared physical carrier may realize more than one virtual or computational World. The carrier is not identical to any one represented World. Consequence in one World may produce consequence in another only through an actual causal Occurrence or mechanism operating through a realized Interface, Channel, or other adequate connection. A Seam, Interface, or Mapping alone supplies structural or representational possibility, not consequence.

No World becomes global by success, reach, adoption count, connectivity, or omission of an external boundary.

### S6.7 Infrastructure

Infrastructure retains the Ontology meaning: a maintained System of resources, Interfaces, services, and dependencies enabling other processes or Worlds to operate.

Structure declares infrastructural arrangement through:

- Components and Subsystems;
- exposed Interfaces and Channels;
- structural dependencies;
- redundancy and concentration;
- Environment relations;
- capacity positions;
- boundaries and failure domains as structural Regions.

Infrastructure provision, hosting, validation, operation, maintenance, funding, or indispensability creates no Ownership, Authority, entitlement, governance, or right to another bearer’s State.

### S6.8 Host and Carrier Roles

A **Host Role** is occupied by a Bearer that supplies a declared physical, computational, organizational, or representational environment in which another realized arrangement is placed.

A **Carrier Role** is a Structural Position occupied by a Bearer that carries an exact material, signal, Representation, or organized form under a declared carrier criterion. Physical Carrier and Hereditary Carrier remain their separate Ontology Kinds; this generic Role neither merges them nor invents one universal carrier relation.

    HOSTS(host_b, hosted_x, hosting_mode, realization_or_placement_criterion, SS)

does not imply that the Host owns, authors, controls, understands, or is identical to the hosted bearer or content.

Hosting mode must distinguish at least physical support, computational execution environment, storage, communication carriage, and representational containment.

### S6.9 Architecture

The word **Architecture** is separated into two exact terms:

1. an **Architecture Specification** is a Structural Profile explicitly designated to prescribe or describe arrangement requirements;
2. **Realized Architecture Status** is attributed to an actual Structural Configuration where an exact Architecture Specification applies, realization is established, and structural conformance obtains.

    ArchitectureSpecification SUBTYPE_OF StructuralProfile

“Realized Architecture” is shorthand for a Structural Configuration bearing Realized Architecture Status. It is not a separate Entity generated by conformance, and the Status does not establish persistence.

A natural structure may exist without an Architecture Specification. A Specification may exist without any Realized Architecture. A deployed Artifact may realize only part of a Specification.

Architecture does not imply one architect, conscious design, Ownership, or Authority.

### S6.10 Layers and Hierarchies

A **Structural Layer** is a Structural Region whose membership criterion selects a partition cell or ordered group of positions under one layering criterion.

A **Structural Hierarchy** is a directed ordering among positions or bearers under one hierarchy relation.

Every Layer or hierarchy must state:

- its criterion;
- direction or ordering;
- whether levels overlap;
- whether the ordering is partial or total;
- whether cycles are excluded;
- the meaning of cross-layer connections;
- Scope and Resolution Context.

Higher position does not mean greater truth, value, personhood, Authority, or causal power. A semantic dependency hierarchy is not necessarily a physical stack, command hierarchy, or execution order.

### S6.11 Structural Capacity

Structure may declare **Structural Capacity** only as a bounded arrangement property, such as the number of available Slots, interfaces, channels, storage positions, or compatible occupants under a Profile.

Actual throughput, compute use, congestion, energy cost, resource consumption, scheduling, and exhaustion over time belong to Dynamics. Economic Capacity and entitlement belong elsewhere.

### S6.12 System Actuality and Recognition

Actuality, representation, structural conformance, and constitutional recognition are independent:

- a System or World may be actually realized but constitutionally unrecognized;
- an exact proposal, Claim, or formation Record may bear a preliminary constitutional Status while the proposed System's or World's actuality remains disputed; the proposed bearer does not thereby bear that Status as an established System or World;
- an established System or World may receive a system-level or World-level constitutional Status only through the applicable valid Constitutional Act;
- a model may represent a System, World, or Architecture without actual realization;
- an arrangement may be only partly realized;
- an actual arrangement may fail a selected Structural Profile;
- an Architecture Specification may exist without a conforming actual Configuration; Realized Architecture Status is borne only by an actual Structural Configuration satisfying S6.9.

These conditions must not be collapsed.

---

## S7 — Relational Structures and the Third-Bearer Test

### S7.1 Relation Is Already Real

A Relation Instance is an Existent where its typed Relation obtains. It may bear Properties such as direction, strength, duration, Provenance, or Status.

That reality does not require inventing a third Entity between the relata.

If A and B stand in Relation r:

    EXISTS(ri, D, I)
    RELATES(ri, r, [A, B], D, I)

the following do not follow:

    EXISTS(C)                     // no third bearer follows
    SAME(ri, A)                   // relation is not A
    SAME(ri, B)                   // relation is not B
    Gamma |- ri : System @ D,I,R = CONFORMS   // not entailed
    Gamma |- ri : Agent @ D,I,R = CONFORMS    // not entailed

The Relation Instance itself is the real between at the relational level. A further bearer is admitted only where separate Evidence establishes it.

### S7.2 The Relational Structure Ladder

Claims about “the between” must locate themselves at exactly one or more established levels:

| Level | What exists | What does not follow |
|---|---|---|
| 1. Relation Instance | a typed Relation obtains among relata | third Entity or System |
| 2. Shared Boundary or Interface | a Boundary position or Interface relates the sides | shared State bearer, Participant, or sovereign |
| 3. Mediating bearer | a separately identified bearer occupies a mediator, carrier, buffer, gate, or stabilizer Role | Agency, Ownership, Authority, or Systemhood |
| 4. Organized relational-system bearer | an independently realized organized, state-bearing System satisfies the Relational System criteria | personhood, Agency, institution, persistence across change, ownership of relata, or sovereignty |
| 5. Relational institution | a constitutional Status and normative Relations are validly constituted concerning an established System, or concerning an exact proposal, Claim, or formation Record for a candidate | ontological creation, physical realization, or global jurisdiction |

No level is inferred merely because language uses “we,” “field,” “shared,” “relation,” “network,” or “system.”

### S7.3 Shared Surface and Shared Interface

A **Shared Surface** is a Structural Region located on, or bearing the region side of, an exact Boundary relation—or an Interface—to which more than one bearer has a typed relation.

Sharing may mean co-location, compatible attachment, joint access, common use, common observation, or another exact relation. The mode must be stated.

A Shared Surface does not become a Participant or owner of its users. Its existence does not create one shared mind, intention, State, identity, or jurisdiction.

### S7.4 Mediator Role

Mediation retains the Ontology meaning. Structure may place a bearer in a **Mediator Role** where it is structurally positioned between or among other relata and conditions their possible connection.

    MEDIATES_POSITIONALLY_n(mediator_m, r1, ..., rn, mediation_profile, SS)

requires `n >= 2`; every `r_k` is an actual Bearer.

    MediationProfile SUBTYPE_OF StructuralProfile

Each concrete Mediation Profile fixes n, ordering, admissible Bearer kinds, criterion, and algebra. The subscript n counts mediated bearers, so each member has total arity n + 1 including the mediator, with n >= 2. The generic display cannot instantiate a variable-arity Relation Kind without that declaration. This says nothing by itself about successful mediation, fairness, neutrality, cognition, decision, or Authority.

A mediator can be a Component of a larger System, an independent System, an Interface, a Representation, or an institution. The exact bearer Kind must be stated.

### S7.5 Buffer Role

A **Buffer Role** is a structural Role occupied by a State-bearing bearer or System placed between source and receiving positions and configured to hold specified inputs temporarily under a Profile.

    BufferProfile {
      bearer_kind
      source_interface
      receiving_interface
      admissible_state_kind
      structural_capacity
      isolation_or_visibility_conditions
      StructuralScope
    }

This Profile states arrangement and capacity only. Actual receipt, temporary holding, eviction, ordering, release, refusal, overflow, or teardown belongs to Dynamics.

A buffer gains no Agency, Ownership, Authority, or entitlement over what it houses.

### S7.6 Stabilizer Role

A **Stabilizer Role** is a structural Role occupied by a bearer, Relation Instance, or Component placed so that a later process may contribute to maintaining specified variables or relations within bounds.

Actual stabilization is a Function and Process that must be established under Dynamics and, where adaptive capacity is claimed, Emergence. Placement in Stabilizer Role does not prove stability or grant sovereignty over the stabilized relata.

### S7.7 Anchor Role

An **Anchor Role** is a structural Role occupied by a Bearer—including, where compatible, a physically realized `ReferenceTargetPosition`—used to constrain placement, orientation, addressing, or comparison within a Configuration. An unoccupied represented anchor is an Anchor Slot, not a Role without a bearer.

An anchor is not necessarily fixed across time; persistence belongs to Continuity. It creates no Ownership, truth, Authority, or superior identity.

### S7.8 Relational System Criteria

A candidate satisfies Structure’s prerequisites for Relational System conformance only where all of these are established:

1. a separately identifiable single or distributed realizer basis;
2. an explicit System Boundary or distributed boundary criterion;
3. typed components, Relation Instances, or constitutive positions;
4. its own physically realized State variables rather than only the private States of the relata;
5. constraints governing those State variables;
6. Interfaces to its relata and Environment;
7. a constitutive criterion showing why the organization is a System rather than a Relation or Aggregate;
8. one Structural Scope and Realization Basis;
9. Evidence adequate to the asserted Resolution Context;
10. no reliance on constitutional recognition as proof of existence.

These criteria establish structural conformance to the imported `RelationalSystem` Kind for the exact Interval. No prior persistence or accumulated history is required: a newly realized organized, bounded, state-bearing System may conform immediately. This result establishes neither numerical identity nor continued existence outside that Interval. Persistence and continuation across reconfiguration, repair, dormancy, replacement, or succession remain entirely with Continuity.

### S7.9 Relational System Non-Entailments

Even where an organized state-bearing System conforms to `RelationalSystem`, it is not automatically:

- either relatum;
- a Participant in the Relation that gave rise to its formation;
- a person, Subject, Human, Agent, Actor, or Controller;
- a fusion or jointly owned extension of its relata;
- owner of their States, outputs, identities, or histories;
- a source of Permission, Authority, obligation, or law;
- a beneficiary of every contribution passing through it;
- a sovereign World or universal jurisdiction.

Each additional classification requires its own criteria and owning-layer judgment.

### S7.10 Third-Bearer Decision Table

| Evidence | Structural placement |
|---|---|
| A and B are related | Relation Instance only |
| A and B touch or share a boundary | Shared Boundary position; no third bearer |
| A and B use one Interface | Interface and use Relations; no third bearer by default |
| A and B connect through M | M is a third bearer if its existence is established; its Role must be typed |
| M stores or transforms State | M may be a Component or System; actual operation belongs to Dynamics |
| interaction realizes an organized state-bearing whole | candidate Relational System; evaluate all S7.8 criteria |
| an institution records or claims recognition of the relation | Claim or Record of recognition only; an institutional Status requires its separately valid assigning or constituting basis and proper bearer |

Where Evidence is insufficient, the result is `UNRESOLVED`, not a phantom System and not denial of the Relation.

---

## S8 — Constraint, Dependency, and Structural Integrity

### S8.1 Constraint Types

Constraint retains the Ontology meaning. Structure distinguishes four adjacent kinds:

1. a **physical constraint** excludes configurations under physical law or realized boundary conditions;
2. a **structural constraint** excludes configurations within an actual arrangement or Structural Profile;
3. a **schema constraint** is a represented formal condition on a Description or data form;
4. **constitutional normative Law** and a **subordinate constitutional Rule** are distinct; neither produces a claimed normative consequence except through its own Current, applicable source and every required valid Constitutional Act or Effect.

These are not interchangeable.

Failure to satisfy a Profile does not make an arrangement physically impossible or constitutionally prohibited. A constitutional Prohibition does not make an arrangement structurally impossible. A schema-valid document may describe a nonexistent or nonconforming structure.

### S8.2 Structural Constraint Declaration

    StructuralConstraintDeclaration {
      constraint_id: Identifier | NONE
      target_profile_or_configuration
      constraint_kind
      condition
      affected_positions_or_relations
      application_criterion
      structural_scope: StructuralScope
      priority_or_composition_rule: Representation | NONE | OPEN
      evidence_claims_or_records: [Claim | Record] | NONE | OPEN
      source_representation: Representation | NONE | OPEN
    }

Constraint evaluation is carried by a conformance report, not hidden inside the declaration:

    ConstraintEvaluation :=
        NOT_EVALUATED
      | EVALUATED {
          result: SATISFIED | VIOLATED | UNRESOLVED | ILL_TYPED
        }

Initial structural constraint kinds are:

- type compatibility;
- cardinality and uniqueness;
- required, optional, repeated, or excluded positions;
- adjacency, incidence, connection, separation, or path condition;
- containment, overlap, partition, or nesting condition;
- Boundary and Interface fit;
- direction, symmetry, or ordering condition;
- declared acyclicity;
- dependency;
- redundancy or concentration;
- realization and carrier condition;
- cross-scale Mapping condition;
- coverage and closure condition.

### S8.3 Structural Dependency

    STRUCTURALLY_DEPENDS_ON(dependent_x, support_y, dependency_mode_m, criterion_c, SS)

is a directed Relation Kind. Initial modes are:

| Mode | Meaning |
|---|---|
| `CONSTITUTIVE_DEPENDENCY` | y enters the criterion by which x is constituted |
| `PLACEMENT_DEPENDENCY` | x’s position is defined relative to y |
| `INTERFACE_DEPENDENCY` | x’s declared structural connection requires an Interface supplied by y |
| `CARRIER_DEPENDENCY` | x’s effective represented or organized form requires y as physical carrier |
| `COMPONENT_DEPENDENCY` | the Profile requires y in a Component Role for x’s conformance |
| `REFERENCE_DEPENDENCY` | x’s structural description or Address is interpreted relative to y |

Both actual relation positions must resolve to Bearers; a Term that merely purports to denote one remains in a Description Claim until its Referent and existence are established. This relation does not by itself establish productive causation, runtime availability, economic dependence, responsibility, Obligation, Authority, or persistence.

Dependency transitivity is mode-specific and never assumed across mixed modes.

### S8.4 Redundancy and Concentration

**Structural Redundancy** exists where more than one independently typed position can satisfy a declared structural requirement under a Profile.

**Structural Concentration** exists where multiple required relations or positions depend on one bearer, interface, Region, or failure domain under the chosen model.

Redundancy does not prove resilience, because alternatives may fail together or remain dynamically unavailable. Concentration may identify a structural susceptibility but does not prove a future failure.

Actual failover, degradation, repair, and recovery belong downstream.

### S8.5 Constraint Conflict

Two Structural Constraints conflict where no Configuration can satisfy both in the same Scope under their exact interpretation.

No constraint wins by textual order, recency, popularity, implementation convenience, or broader wording. A Profile must declare an exact priority, composition, exception, or conflict result.

Where no resolution rule exists:

    constraint_posture = CONSTRAINT_CONFLICT
    SATISFIES_STRUCTURE = UNRESOLVED

If a Profile explicitly requires simultaneous satisfaction of both incompatible constraints, the Configuration `DOES_NOT_CONFORM` to that Profile. The Profile itself may also be defective or unsatisfiable; that is a separate Claim.

### S8.6 Structural Integrity

**Structural Integrity** is conformance of an actual Configuration to exact arrangement criteria required to keep its declared structure well-formed at one Structural Scope.

    StructuralIntegrityValue :=
        INTEGRITY_ESTABLISHED
      | INTEGRITY_VIOLATED
      | INTEGRITY_UNRESOLVED

    IntegrityEvaluation :=
        NOT_EVALUATED
      | EVALUATED { value: StructuralIntegrityValue }

The Status must name the Profile and criteria. It is not moral integrity, legal validity, truth, security, safety, resilience, coherence in every sense, or continuity.

### S8.7 Structural Defects

Profile-relative structural defects include:

- missing required position;
- incompatible occupant;
- cardinality violation;
- dangling endpoint;
- prohibited overlap;
- Boundary or Interface mismatch;
- cycle in a declared acyclic relation;
- unresolved or contradictory placement;
- absent required realization basis;
- Relation-to-System cast;
- shared State without bearer;
- incompatible scale or reference frame;
- closure Claim unsupported by coverage.

A defect can be recorded without erasing the Configuration. Repair is an Occurrence and belongs to Dynamics; whether the repaired System is the same System belongs to Continuity.

### S8.8 Multi-Scale Structure

A macrostructure may be compatible with multiple micro-configurations. A microstructure may realize several higher-level descriptions under different criteria.

Cross-scale inference requires:

    ScaleMapping {
      source_resolution_context
      target_resolution_context
      aggregation_or_projection_rule
      preserved_relations
      lost_or_unresolved_detail
      uncertainty
      domain_of_validity
    }

No finer or coarser scale is automatically more real. An invariant at one scale does not travel to another without its Mapping.

### S8.9 Symmetry and Structural Invariant

A structural Symmetry or Invariant retains the formal Ontology meaning: preservation under a specified transformation.

Structure may state which arrangement relation remains invariant under an allowed transformation. It does not assert that the transformation occurs. Dynamic conservation, stability under perturbation, or maintenance over time requires later-layer analysis.

### S8.10 Structural Possibility

A Profile may classify a Configuration as structurally expressible or excluded under its constraints. This does not decide physical possibility unless all relevant physical constraints are included, and it does not decide constitutional Permission.

The categories remain:

    structurally conforming
    structurally nonconforming
    structurally unresolved
    structurally ill-typed

They are not synonyms for lawful, successful, stable, persistent, intelligent, valuable, or alive.

---

## S9 — Description, Conformance, and Representation

### S9.1 Structural Relation Registry

The grammar is partitioned so that Role occupancy, actual relations, and derived assessments cannot be conflated.

#### A. Role occupancy displays

These are typed displays of imported `PLAYS`, not separate Structure-owned `RELATES` projections:

| Display | Expansion |
|---|---|
| `OCCUPIES_STRUCTURAL_POSITION(b, p, c, SS)` | `PLAYS(b, p, c, SS.interval)`; b: Bearer, p: StructuralPosition, c: StructuralConfiguration |
| `OCCUPIES_PORT(b, p, i, SS)` | `PLAYS(b, p, i, SS.interval)`; b: Bearer, p: PortRole, i: Interface |

Centre Role and Aperture Role occupancy use the same rule. An unoccupied represented position is a Structural Slot, never an unoccupied Role.

#### B. Structure-owned `RELATES` projections

Every row has the S0.7 expansion. Only the listed Bearers are relata; modes, Criteria, Profiles, topologies, and Roles are qualifiers or side judgments.

| Projection | Actual Bearer relata | Qualifiers or side judgments | Base algebra | Owner |
|---|---|---|---|---|
| `COMPOSES(x, y; m; SS)` | x,y: Bearer | `CompositionMode m` | directed; algebra mandatory by mode | S2 |
| `COMPONENT_OF(b, s; r; SS)` | b: Bearer; s: System | `PLAYS(b,r,s,I)` with ComponentRole r | directed; non-reflexive; transitivity conditional | S2 |
| `SUBSYSTEM_OF(a, b; m; SS)` | a,b: System | `SubsystemMode m` | directed; non-reflexive; transitivity conditional | S2 |
| `BOUNDS_REGION(br, g; c; SS)` | br: Boundary RelationInstance; g: StructuralRegion | Criterion c | directed; non-transitive | S3 |
| `IN_REGION(x, g; m; SS)` | x: Bearer; g: StructuralRegion | `InclusionMode m` | directed; algebra by mode | S3 |
| `CONTAINS_REGION(g, x; m; SS)` | g: StructuralRegion; x: Bearer | same mode as inverse `IN_REGION` | directed; declared inverse only | S3 |
| `SUBREGION_OF(a, b; m; SS)` | a,b: StructuralRegion | `SubregionMode m` | directed; algebra by mode | S3 |
| `COVERS_STRUCTURALLY(k, x; m,M,c; SS)` | k: Bearer candidate conforming to Pattern; x: Bearer | `InclusionMode m`; Mapping M; Criterion c; exact `CLOSED_FOR` coverage | directed; non-transitive; obtaining relation plus S3.6 conditions grounds StructuralCover conformance | S3 |
| `OVERLAPS_REGION(g1, g2; c; SS)` | g1,g2: StructuralRegion | Criterion c | symmetric under same c; non-transitive | S3 |
| `INTERSECTS_AT(g1, g2, z; c; SS)` | g1,g2: StructuralRegion; z: StructuralRegion or actual Bearer occupying the intersection position; a StructuralSlot only where the relation concerns the Profile Representation itself | Criterion c | symmetric in g1/g2; non-transitive | S3 |
| `STRUCTURALLY_SEPARATE(g1, g2; c; SS)` | g1,g2: StructuralRegion | Criterion c; exact closed coverage | symmetric under same c; non-transitive | S3 |
| `NESTED_IN(a, b; m; SS)` | a,b: compatible Bearer | `NestingMode m` | directed; algebra by mode | S3 |
| `ADJACENT_TO(a, b; t; SS)` | a,b: compatible Bearer or StructuralSlot | StructuralTopology t | symmetry declared by t; non-transitive unless declared | S4 |
| `INCIDENT_TO(l, e; t; SS)` | link occupant/Slot l; endpoint occupant/Slot e | StructuralTopology t | directed; non-transitive | S4 |
| `ATTACHED_TO(a, b; m; SS)` | a,b: compatible Bearer or StructuralSlot | `AttachmentMode m` | algebra by mode | S4 |
| `STRUCTURALLY_CONNECTED(a, b; k,t; SS)` | a,b: compatible Bearer or StructuralSlot | `ConnectionKind k`; StructuralTopology t | algebra by k and t | S4 |
| `APERTURE_ON(a, i; SS)` | a: ApertureRegion; i: Interface RelationInstance | `NONE`; any InterfaceProfile is tested separately by an explicit applicability and conformance judgment | directed; non-transitive | S5 |
| `ENVIRONMENT_OF(e, s; c; SS)` | e: StructuralRegion; s: System | Criterion c | directed; non-symmetric | S6 |
| `HOSTS(h, x; m,c; SS)` | h,x: Bearer | `HostingMode m`; Criterion c | directed; non-transitive unless declared | S6 |
| `MEDIATES_POSITIONALLY_n(m, r1...rn; p; SS)` | mediator m plus n mediated Bearers; total arity n + 1; n >= 2 | `MediationProfile p`; p fixes n | exact family member; order and algebra declared by p | S7 |
| `STRUCTURALLY_DEPENDS_ON(x, y; m,c; SS)` | x,y: Bearer | `DependencyMode m`; Criterion c | directed; transitivity mode-specific | S8 |

Every qualifier above binds to the named Relation Instance `ri`; it is neither a free parameter nor another relatum. These built-in qualifier attributes are Property Kinds with sole admissible bearer `RelationInstance`:

| Attribute Kind | Value-space |
|---|---|
| `CompositionModeQualifier` | `CompositionMode` |
| `SubsystemModeQualifier` | `SubsystemMode` |
| `InclusionModeQualifier` | `InclusionMode` |
| `SubregionModeQualifier` | `SubregionMode` |
| `NestingModeQualifier` | `NestingMode` |
| `AttachmentModeQualifier` | `AttachmentMode` |
| `ConnectionKindQualifier` | `ConnectionKind` |
| `HostingModeQualifier` | `HostingMode` |
| `DependencyModeQualifier` | `DependencyMode` |
| `CriterionQualifier` | `Criterion` |
| `TopologyQualifier` | `StructuralTopology` |
| `MappingQualifier` | `Mapping` |
| `MediationProfileQualifier` | `MediationProfile` |
| `CoverageQualifier` | `CoveragePosture` |

Each attribute-bound qualifier expands to `ATTRIBUTION(attribute_kind, ri, value, SS.domain, SS.interval, SS.resolution_context)`. Role occupancy expands to its exact `PLAYS` side judgment. Row bindings are: `COMPOSES` -> CompositionMode; `COMPONENT_OF` -> `PLAYS`; `SUBSYSTEM_OF` -> SubsystemMode; boundary, overlap, intersection, environment, and dependency criteria -> Criterion; region inclusion -> InclusionMode; subregion -> SubregionMode; separation -> Criterion plus exact Coverage; nesting -> NestingMode; adjacency and incidence -> Topology; attachment -> AttachmentMode; connection -> ConnectionKind plus Topology; hosting -> HostingMode plus Criterion; mediation -> MediationProfile; cover -> InclusionMode plus Mapping, Criterion, and Coverage; dependency -> DependencyMode plus Criterion. `APERTURE_ON` has no qualifier.

#### C. Derived assessments

`PATH_ASSESSMENT`, `REACHABILITY`, `CENTRALITY_ASSESSMENT`, and `INTERFACE_POSTURE` return their closed S0.6 result grammars. They are not actual Relation Instances. Negative path or reachability results require coverage closed for the relevant topology; otherwise the result is `UNRESOLVED`.

#### D. Imported operators and relation kinds

`REALIZES`, `REPRESENTS`, `PARTICIPATES`, `PLAYS`, and `SCOPED_TO` retain Ontology ownership and signatures. `ONTOLOGY.CONSTITUTES` and `CONSTITUTION.CONSTITUTES` retain their separately qualified owners and signatures; bare `CONSTITUTES` is not admitted. An actual Coupling is expressed through a named Coupling Relation Instance using full `RELATES`, with direction, strength, Channel, and coupling kind carried by typed attributes. No `COUPLED_TO` predicate is installed here.

### S9.2 Relation Declaration Closure

For every Structure-owned projection, the base registry plus applicable Profile must make recoverable:

- the exact arity and position Kinds;
- whether positions are ordered;
- Domain, Interval, and Resolution Context;
- realization or observation basis;
- direction, symmetry, reflexivity, transitivity, inverse, and composition posture;
- cardinality and exclusivity conditions;
- coverage posture;
- non-entailments.

The registry fixes arity and actual relatum kinds. A Profile may narrow qualifier values, cardinality, algebra, or coverage; it cannot change the base actual arity, relatum positions, or meaning.

Every qualifier value is either a built-in registered value or is introduced by a declaration naming its owner, value-space, application relation or operator, admissible inputs, algebra or evaluation rule, criterion, and Scope. Free text is `ILL_TYPED`; `OPEN` in an actual use yields `UNRESOLVED`.

### S9.3 Conformance Use of Profile Composition

The only canonical Profile composition projection is declared at S1.7. Conformance must evaluate that exact `ProfileCompositionRecord`; no second implicit composition, override, or priority grammar exists here.

### S9.4 Structural Conformance Procedure

    SATISFIES_STRUCTURE(configuration_c, profile_p, SS)
      -> CONFORMS | DOES_NOT_CONFORM | UNRESOLVED | ILL_TYPED

The evaluation order is:

1. **Scope** — `WF_STRUCTURAL_SCOPE(SS) = WELL_FORMED`.
2. **Term typing** — c is a Term purporting to denote an exact Structural Configuration; any Claim about c belongs in Gamma or the Evidence basis, never in the configuration argument; p is a Structural Profile.
3. **Application** — `STRUCTURE_APPLIES(p, c, SS)` is evaluated from the declared target Kind or predicate and Scope. Applicability is independent of whether an actual target presently realizes the position. `NOT_APPLICABLE` requires an established target-kind, predicate, or Scope mismatch.
4. **Actuality gate** — existence and Realization Basis are established, disproved, or unresolved without being conflated with applicability or conformance. Unless actuality is established, the conformance function is not invoked and the report uses `ConformanceEvaluation = NOT_EVALUATED`.
5. **Description** — the attributable Structure Description is well-formed and its coverage is exact enough for each tested constraint.
6. **Relations** — every relied-upon actual relation has a typed Relation Instance or admissible existential projection.
7. **Constraints** — each applicable structural constraint is evaluated under its own criterion.
8. **Conflict** — unresolved constraint conflicts remain visible.
9. **Result** — one closed result is returned with reasons and unresolved dependencies.

Result precedence after `STRUCTURE_APPLIES = APPLICABLE` and the actuality gate is exact:

1. malformed invocation or category mismatch -> `ILL_TYPED`;
2. at least one decisive necessary-constraint failure or exclusion violation -> `DOES_NOT_CONFORM`, even if unrelated constraints remain unresolved;
3. no decisive failure and at least one unresolved necessary constraint, relation, conflict, Mapping, coverage dependency, or Evidence position -> `UNRESOLVED`;
4. every necessary constraint established satisfied and no applicable exclusion established violated -> `CONFORMS`.

`CONFORMS` requires every necessary constraint to be established as satisfied and no applicable exclusion to be established as violated.

`DOES_NOT_CONFORM` requires at least one necessary constraint to be established as failed or one applicable exclusion to be established as violated.

`UNRESOLVED` applies only after actuality and applicability are established, where neither positive nor decisive negative conformance is established because coverage, relation, constraint, conflict, scale Mapping, or Evidence remains unresolved.

`ILL_TYPED` applies where the judgment violates the grammar.

### S9.5 Existence and Conformance Matrix

Existence and structural conformance are independent:

| Existence / realization | Conformance evaluation | Correct reading |
|---|---|---|
| established | `CONFORMS` | actual conforming Configuration |
| established | `DOES_NOT_CONFORM` | actual nonconforming Configuration |
| established | `UNRESOLVED` | actual Configuration whose conformance is unresolved |
| unresolved | procedure posture `NOT_EVALUATED` | actuality is unresolved; applicability remains independently assessed |
| nonexistence or nonrealization established | procedure posture `NOT_EVALUATED` | no actual target is available for conformance testing; applicability is not thereby decided |
| judgment ill-typed | `ILL_TYPED` inside an evaluated result | expression defect; Reality is not thereby decided |

A Profile can conform to its own schema while no actual target realizes it.

### S9.6 Structure Description Truth Posture

A Description’s formatting and structural well-formedness are distinct from whether its Claims obtain.

For each material Description Claim, the result should expose:

    description_well_formedness:
      WELL_FORMED | ILL_TYPED

    target_existence:
      EXISTENCE_ESTABLISHED | NONEXISTENCE_ESTABLISHED |
      EXISTENCE_UNRESOLVED | ILL_TYPED

    evidential_support_posture:
      SUPPORTED | REFUTED | UNDERDETERMINED | NOT_EVALUATED

    coverage_posture:
      OPEN | PARTIAL | CLOSED_FOR | COVERAGE_UNRESOLVED

    conformance_evaluation:
        NOT_EVALUATED
      | EVALUATED {
          result:
              CONFORMS
            | DOES_NOT_CONFORM
            | UNRESOLVED
            | ILL_TYPED
        }

These dimensions must not be compressed into one Boolean.

### S9.7 Structural Equivalence and Isomorphism

    STRUCTURAL_EQUIVALENCE(c1, c2, criterion, SS)
      -> EQUIVALENT | NOT_EQUIVALENT | UNRESOLVED | ILL_TYPED

Structural equivalence means interchangeability under the stated arrangement criterion. It is not numerical Identity.

A **Structural Isomorphism** is a bijective Mapping between selected positions and Relation Instances that preserves the exact declared structural relations.

Two isomorphic Configurations may:

- have numerically different bearers;
- occupy different physical Locations;
- have different histories, provenance, materials, capacities, or values;
- fail to continue one another;
- receive different constitutional treatment only where a valid independent basis permits it.

Isomorphism at one Resolution does not imply equivalence at every Resolution.

### S9.8 Graph, Diagram, Schema, and Digital Twin

A Graph, diagram, schema, topology file, model, or digital twin is a Representation.

A **Structural Digital Twin** is a maintained Model whose represented structural distinctions are mapped to a Reference Target under declared synchronization and fidelity criteria.

It is not numerically identical to its target. Synchronization belongs to Dynamics; persistence and version lineage belong to Continuity. A stale twin remains a real Representation but may cease to support current Claims.

No Graph edge becomes an actual Relation Instance without an exact interpretation and Evidence. No diagrammed System becomes actual by visual completeness.

### S9.9 Human Source Form and Machine Projection

This exact human-readable Form is the semantic source for Structure. A machine projection may encode:

- Kind declarations;
- Relation signatures;
- Profile schemas;
- constraint records;
- result grammars;
- conformance fixtures;
- provenance and exact source identification.

A machine projection must identify the exact source digest, byte length, encoding, media type, projection method, and known loss. It is a derivative Representation and cannot silently add Kinds, close `OPEN` fields, choose unresolved criteria, or change layer ownership.

Machine execution does not establish that a structural Claim is true. A parser cannot cure semantic ambiguity by choosing a default.

### S9.10 Provenance and Correction

Every published Structural Profile, Description, Record, diagram, and machine projection should preserve:

- source Form identity;
- author or producer attribution;
- creation and modification intervals;
- input sources;
- transformation method;
- known omissions and losses;
- correction history;
- present coverage and epistemic posture.

Correction creates a new Representation state or Form. It does not rewrite the historical existence of the earlier Representation or retroactively alter the actual Configuration it described.

### S9.11 Conformance Report

An exact conformance report has this minimum projection:

    StructuralConformanceReport {
      report_id
      configuration_term
      profile_exact_identity
      structural_scope
      applicability_result
      existence_and_realization_posture
      description_identity_and_coverage
      tested_constraints: [
        {
          constraint
          evidence
          result
          reason
        }
      ]
      conflicts
      overall_result
      unresolved_dependencies
      method
      provenance
      correction_or_review_path
    }

A report may supply information participating in an Evidence relation. It is not Evidence by identity, the Configuration, Fact, Truth, constitutional Decision, or final word about Reality.

### S9.12 Adversarial Conformance Fixtures

The following outcomes are required by this Form:

1. **Bare relation** — A relates to B; no third bearer or Relational System is inferred.
2. **Mediated relation** — A connects to B through M; M gains no Agency, Ownership, or Authority by position.
3. **Unused Aperture** — Boundary, Interface, Port, and Aperture exist; no crossing, Admission, Consent, or Effect is inferred.
4. **Graph-only path** — a represented path exists; actual Channel, delivery, and causal reach remain unresolved.
5. **Containment without componenthood** — X is spatially inside S but is not therefore a Component.
6. **Formal membership without parthood** — Y is a formal Member but not therefore a physical Part.
7. **Shared component** — one bearer is a Component of two Systems; the Systems do not fuse.
8. **Isomorphic Configurations** — structural equivalence does not establish numerical Identity.
9. **Open inventory** — omission of Z does not establish Z’s absence.
10. **Closed inventory** — absence may be supported only within the exact `CLOSED_FOR` dimensions.
11. **Multi-resolution decomposition** — different decompositions at different Resolutions are not automatically contradictory.
12. **One-Interval snapshot** — a Configuration at I does not establish persistence at I2.
13. **Profile exclusion** — structurally excluded does not mean physically impossible or constitutionally prohibited.
14. **Unrealized Architecture** — a well-formed Architecture Specification does not become actual without a realizer.
15. **Distributed structure** — spatial contiguity is not required unless the Profile requires it.
16. **Criterion-relative Centre** — a Centre under one metric may not be central under another.
17. **Unknown arrangement** — failure to fit registered topology families yields unresolved placement, not nonexistence.

### S9.13 Canonical Structural Declaration Registry

This registry closes the built-in vocabulary used by this Form. It is not a second Ontology. Every row incorporates these explicit shared fields:

    declaration_scope: exact StructuralScope
    disjoint_with: OPEN unless the row states otherwise
    overlaps: OPEN unless the row states otherwise
    epistemic_statuses:
      Observed | Measured | Inferred | Posited |
      Established | Disputed | Unknown, as category permits
    coverage_requirement: adequate for the exact Claim;
      negative or exhaustive Claims require CLOSED_FOR
    conformance_result:
      CONFORMS | DOES_NOT_CONFORM | UNRESOLVED | ILL_TYPED
    actuality_evaluation:
      actual rows -> RealizationPosture
      represented rows -> realization posture of the Representation
      Roles -> obtaining or non-obtaining `PLAYS` posture
      Properties and Statuses -> obtaining or non-obtaining `ATTRIBUTION` posture
      metalanguage rows -> NOT_APPLICABLE
    later_layer_error:
      LAYER_OWNERSHIP_ERROR

For an **actual** row, sufficiency requires existence plus the row criterion and relation basis in the exact Structural Scope. For a **represented** row, sufficiency requires a physically encoded Representation with carrier, content, Mapping or interpretation, provenance, and recoverability. For a **Role**, sufficiency is an obtaining `PLAYS` relation with an admissible Bearer, exact context, and Interval. For a **Property** or **Status**, sufficiency is an obtaining `ATTRIBUTION` under its declared value-space and criterion. These are named shared fields, not invisible defaults.

Unless a row narrows them, no declared structural type entails numerical Identity, Agency, Authority, Ownership, Consent, Admission, Permission, persistence, succession, intelligence, meaning, or value. A represented type additionally entails neither target existence, truth, actuality, conformance, Currentness, nor execution.

#### S9.13.1 Structural kinds and patterns

| Name | Parent / category | Discriminating necessary and sufficient condition | Bearer, relation, or realization position | Owner and exact boundary |
|---|---|---|---|---|
| `StructuralScope` | `Scope` / other | one well-formed `WorldContext` plus Domain, Interval, Resolution Context, and further conditions | metalanguage application boundary; realization `NOT_APPLICABLE` | S0; not World, Domain, jurisdiction, or globality |
| `StructuralConfiguration` | `Configuration`, `State` / State | actual selected arrangement with bearers, Relation Instances, constraints, coverage, and established Realization Basis | borne by Bearer, System, World, or RelationInstance | S1; not Description, Profile, transition, history, or persistence |
| `StructuralProfile` | `Representation` / Pattern | represented target Kind or predicate, application criterion, typed constraints, Scope, and coverage | represented realization through `REPRESENTS`; an application Judgment, not the raw Profile, may be `SCOPED_TO` | S1; not constitutional Rule, Obligation, or Prohibition |
| `StructuralPosition` | `Role` / other | relation-defined position within one actual Configuration | only an admissible Bearer through `PLAYS` | S1; not Slot or occupant; occupancy has no continuity by name |
| `StructuralSlot` | `ReferenceTargetPosition` / other | physically realized Profile position representing a possible structural position | Bearer in Reference; may remain unoccupied | S1; no actual Role, occupant, availability, Capacity, or passage |
| `Subsystem` | `System` / Entity | System with its own Boundary and criterion that occupies a Component or Constituent relation in another System | actual System bearer | S2; no shared identity, governance, complete State, or persistence with containing System |
| `StructuralRegion` | `Bearer` / other | actual criterion-bounded selected extent or membership condition | actual relatum of region and boundary relations | S3; no forced spatiality, containerhood, isolation, jurisdiction, sovereignty, Access, or Ownership |
| `StructuralCover` | `Pattern`, `Bearer` / Pattern | `COVERS_STRUCTURALLY(k,x;m,M,c;SS)` obtains and k exposes the exact nonempty actual cell structure required by S3.6 | actual cover Pattern is the first relatum after conformance is established | S3; not Partition, System, container, or merely partial Description |
| `StructuralPartition` | `StructuralCover` / Pattern | closed cover whose cells satisfy the exact non-overlap rule | same as StructuralCover | S3; not physical division, isolation, jurisdiction, or Ownership |
| `StructuralTopology` | `Pattern`, `Bearer` / Pattern | actual positions or Slots and typed relations satisfy one topology criterion | realized by Configuration, System, or World; may bear coverage and may be a qualifier or declared relatum | S4; not Graph, causal reach, traversal, trust, or consensus |
| `StructuralPath` | `Pattern` / Pattern | ordered compatible occupant or Slot and Link sequence under one topology | used by `PATH_ASSESSMENT`; never traversal itself | S4; no passage, Access, delivery, or future availability |
| `ApertureRegion` | `StructuralRegion` / other | actual Interface Region with source and receiving sides and eligible crossing kinds | actual Bearer in `APERTURE_ON` | S5; no crossing, Admission, Permission, Intention, or acceptance |
| `StructuralLayer` | `StructuralRegion` / other | Region or ordered group selected under one layering criterion | actual Bearer | S6; no rank, truth, value, Authority, or causal precedence |
| `StructuralHierarchy` | `StructuralTopology` / Pattern | directed ordering under one declared hierarchy relation | realized exact topology | S6; no domination, command, Authority, or worth |
| `StructuralConstraint` | imported `Constraint` / other | actual condition excluding configurations under exact conditions | constraint-specific actuality and application relation | S8; no physical impossibility, Prohibition, or Obligation by name |
| `StructuralIsomorphism` | `Mapping` / Pattern | bijective Mapping preserving the selected positions and structural relations | Mapping plus exact preservation basis | S9; not Identity, Continuity, equal provenance, value, or treatment |
| `StructuralDigitalTwin` | `Model` / Pattern | maintained Model mapped to a Reference Target under synchronization and fidelity criteria | represented realization | S9; not target identity, current accuracy, truth, or persistence |

#### S9.13.2 Structural Roles

Every row below is `SUBTYPE_OF StructuralPosition`, is occupied only by a compatible Bearer through `PLAYS`, and inherits the Role boundary: no intrinsic Kind, identity with the bearer, exclusive or persistent occupancy, control, rank, Authority, Ownership, or value.

| Role | Admissible occupant and exact condition | Owner | Additional non-entailment |
|---|---|---|---|
| `ComponentRole` | Bearer in an exact component position of a System | S2 | no inherited System Status |
| `BoundarySideRole` | Bearer positioned as interior, exterior, or boundary position under one criterion | S3 | no free-standing Region or jurisdiction |
| `NodeRole` | Bearer occupying a node position in actual topology | S4 | no Agency, personhood, Participation, or Authority |
| `LinkRole` | RelationInstance, Channel bearer, connector, Interface, or other declared Bearer in a link position | S4 | no traversal, communication, or causation |
| `EndpointRole` | Bearer occupying a terminus position of a Link | S4 | no Address identity or connection success |
| `HubRole` | NodeRole occupant satisfying the exact hub criterion | S4 | no control or importance |
| `CentreRole` | Bearer satisfying the exact Centrality criterion in a Configuration | S4 | no universal centre, Kernel, Self, persistence, priority, or value |
| `PortRole` | connector, Channel, carrier, or other compatible Bearer at an Interface position | S5 | no connection, Access, Permission, or passage |
| `ApertureRole` | compatible Interface component or Bearer at an ApertureRegion | S5 | no crossing, qualification, Decision, or gateway Authority |
| `GateRole` | Bearer or Interface component placed where a crossing attempt would encounter it | S5 | no filtering success, refusal, Permission, or Authority |
| `FilterRole` | Bearer, represented Profile, or implemented mechanism positioned to distinguish inputs under a Criterion | S5 | a Profile does not execute; placement does not prove correct classification |
| `HostRole` | Bearer supplying a declared physical, computational, organizational, or representational environment | S6 | no authorship, control, identity, Ownership, or understanding |
| `CarrierRole` | Bearer carrying an exact material, signal, Representation, or organized form | S6 | no identity with content, meaning, or Authority |
| `MediatorRole` | Bearer positionally conditioning possible connection among other relata | S7 | no success, neutrality, fairness, cognition, Decision, or Authority |
| `BufferRole` | State-bearing Bearer or System placed between source and receiver under BufferProfile | S7 | no receipt, holding, ordering, release, Ownership, or entitlement |
| `StabilizerRole` | Bearer, RelationInstance, or Component positioned so a later process may maintain variables within bounds | S7 | no actual stabilization, adaptation, control, or sovereignty |
| `AnchorRole` | Bearer used to constrain placement, orientation, addressing, or comparison | S7 | no temporal fixity, truth, Authority, or superior identity |

`PortSlot`, `CentralSlot`, `ApertureSlot`, and `AnchorSlot` are display shorthands for a `StructuralSlot` whose represented target or Role kind is the named registered type. They do not create additional Kinds.

#### S9.13.3 Properties and Statuses

| Name | Parent / admissible bearer | Closed criterion or value-space | Owner and boundary |
|---|---|---|---|
| `CoveragePosture` | Status / a declared Representation or Record carrying the exact inventory, topology, cover, or partition account | `OPEN`, `PARTIAL`, `CLOSED_FOR`, `COVERAGE_UNRESOLVED` under its closure method | S1; never attributed to a non-Bearer; completeness only inside declared dimensions |
| `EnclosureStatus` | Status / StructuralRegion | `ENCLOSED \| NOT_ENCLOSED`; exact Boundary/topology test; negative requires `CLOSED_FOR` | S3; no impermeability, isolation, Access rule, or jurisdiction |
| `Centrality` | Property / compatible Bearer or StructuralSlot | `CentralityValue { configuration_term; Criterion; CENTRAL \| NOT_CENTRAL; QuantityValue \| NONE; comparison_population }`; decisive assessment required | S4; no control, priority, persistence, or value |
| `InterfacePosture` | Status / Interface | `InterfacePostureValue` for one exact crossing kind; `ILL_TYPED` is never attributed | S5; no Permission, Prohibition, crossing, or future availability |
| `RealizedArchitectureStatus` | Status / StructuralConfiguration | `RealizedArchitectureValue { ArchitectureSpecification; StructuralConformanceReport; REALIZED_UNDER_SPECIFICATION }`; actuality, applicability, and `CONFORMS` required | S6; no second Entity, architect, Ownership, or persistence |
| `StructuralCapacity` | Property / Configuration, System, Interface, Channel, or declared Bearer | `COUNT_CAPACITY { QuantityValue; counted predicate; Criterion } \| COMPATIBILITY_CAPACITY { nonempty positions; admissible Kinds; Criterion }` | S6; no throughput, execution, use, entitlement, or economic Capacity |
| `SharedSurfaceStatus` | Status / StructuralRegion or Interface | `SharedSurfaceValue { SHARING_ESTABLISHED; at least two Bearers; nonempty sharing RelationInstances; mode; Criterion }` | S7; no shared State, Self, mind, identity, Ownership, or jurisdiction |
| `StructuralRedundancy` | Property / Configuration or System | `StructuralRedundancyValue { requirement; at least two satisfying positions; independence Criterion }` | S8; no resilience, availability, or failover success |
| `StructuralConcentration` | Property / Configuration or System | `StructuralConcentrationValue { at least two required positions or relations; shared support Bearer; nonempty dependency RelationInstances; Criterion }` | S8; no future failure or responsibility |
| `StructuralIntegrityStatus` | Status / StructuralConfiguration | `StructuralIntegrityValue` under an exact Profile; an evaluation does not itself create attribution | S8; not moral integrity, truth, security, safety, continuity, or general coherence |

“Enclosure,” “Shared Surface,” and “Realized Architecture” are Status-qualified readings of their registered bearers, not independent Entity Kinds. “Structural defect” is a reasoned nonconformance finding or Status, not an Entity.

Every Property or Status row is instantiated only through `ATTRIBUTION(attribute_kind, bearer, value, SS.domain, SS.interval, SS.resolution_context)`. Its stated criterion is the assignment rule. `UNRESOLVED`, `NOT_EVALUATED`, or `ILL_TYPED` evaluation creates no attribution.

#### S9.13.4 Profiles, Descriptions, Records, and Mappings

| Name | Parent / category | Discriminating content and realization | Owner and boundary |
|---|---|---|---|
| `StructureDescription` | `Representation` / Pattern | Claims about actual or possible arrangement, exact Scope, and coverage; represented realization | S1; not Configuration |
| `StructureSnapshot` | `StructureDescription` / Pattern | Description additionally indexed to an as-of point or bounded Interval | S1; no persistence or complete history |
| `StructureRecord` | `Record` / Pattern | persistent Representation preserving Description, method, provenance, Scope, and result | S1; authenticity does not entail truth |
| `RepresentationBinding` | `Record` / Pattern | Representation, actual carrier or carriers, represented content, ReferenceTargetPosition, Mapping or interpretive practice, Domain, Interval, and use conditions; exact `REPRESENTS` relation required | S1; represented content or purported Referent never replaces the target position |
| `DistributedRealizerBasis` | `Record` / Pattern | nonempty actual Bearer list, organizing Relation Instances, joint criterion, per-bearer `INSUFFICIENT` or `UNRESOLVED` posture, Scope, and evidence Claims or Records | S1; represents but never occupies the plural realizer position; not Aggregate, Whole, System, State bearer, or collective Agent |
| `RealizationBasis` | `Record` / Pattern | exact represented `REALIZES` Claim and closed realization posture | S1; does not replace or create `REALIZES` |
| `ProfileCompositionRecord` | `Record` / Pattern | exact Profiles, Scope overlap or partition, priorities, conflicts, preservation, and override posture | S1; no silent override or constitutional force |
| `CompositionAlgebra` | `Representation` / Pattern | one CompositionMode plus exact inputs and algebra | S2; algebra does not transfer between modes |
| `StructuralConstraintDeclaration` | `Representation` / Pattern | representation of one StructuralConstraint, criterion, affected positions, Scope, and source | S8; not the actual Constraint by identity or a constitutional Norm |
| `CardinalityConstraint` | `Representation` / Pattern | exact target position or relation, minimum, maximum, exact count, counting criterion, distinctness criterion, and Structural Scope | S2; not an actual Constraint by identity; count does not prove bearer identity |
| `PartitionProfile` | `StructuralProfile` / Pattern | target, cells, represented `COVERS_STRUCTURALLY` Claim, inclusion mode, Mapping, criterion, overlap rule, coverage, and Scope | S3; does not physically partition target |
| `InterfaceProfile` | `StructuralProfile` / Pattern | Interface sides, crossing kinds, ports, Mapping, preservation, loss, and constraints | S5; no crossing or Permission |
| `MediationProfile` | `StructuralProfile` / Pattern | exact n >= 2, ordered mediator and mediated positions, admissible Bearer kinds, criterion, and algebra | S7; no success, fairness, neutrality, cognition, Decision, or Authority |
| `BufferProfile` | `StructuralProfile` / Pattern | buffer bearer kind, source and receiving Interfaces, State kind, capacity, isolation, and visibility | S7; no buffering Occurrence |
| `ArchitectureSpecification` | `StructuralProfile` / Pattern | explicitly designated arrangement requirements | S6; not actual architecture, architect, Ownership, or Authority |
| `SystemStructureDescription` | `StructureDescription` / Pattern | exact System arrangement projection | S6; component list alone does not establish Systemhood |
| `WorldStructureDescription` | `StructureDescription` / Pattern | exact World arrangement projection | S6; no sovereignty, globality, or World actuality |
| `ScaleMapping` | `Mapping` / Pattern | source and target Resolutions, rule, preserved relations, loss, uncertainty, and validity Domain | S8; no cross-scale Identity or undeclared truth preservation |
| `StructuralConformanceReport` | `Record` / Pattern | applicability, actuality, tests, conflicts, result, method, and provenance | S9; may participate in Evidence; not Evidence by identity, Configuration, Fact, Truth, or Decision |
| `UnresolvedStructuralPlacement` | `Record` / Pattern | known structure, candidates, rejected Profiles, preserved distinctions, tensions, and Evidence | S10; not a miscellaneous Kind, nonexistence, or disorder |

#### S9.13.5 Closed value spaces and extension points

These are qualifier or result value-spaces, not ontological Kinds or actual relation relata:

| Value-space | Values | Owner | Boundary |
|---|---|---|---|
| `WorldContext` | `SINGLE_WORLD` or `CROSS_WORLD` records | S0 | no meta-World, containment, governance, or globality |
| `WorldKindPosture` | `CONFORMS`, `DOES_NOT_CONFORM`, `UNRESOLVED`, `ILL_TYPED` | S0 | not World existence |
| `WorldExistencePosture` | `EXISTENCE_ESTABLISHED`, `NONEXISTENCE_ESTABLISHED`, `EXISTENCE_UNRESOLVED`, `ILL_TYPED` | S0 | not World Kind conformance |
| `CompositionMode` | the six modes in S2.1 | S2 | not Membership, Ownership, containment, implementation, or realization |
| `InclusionMode` | the modes in S3.3 | S3 | no algebra exports to another mode |
| `BoundarySideKind` | `INTERIOR_SIDE`, `EXTERIOR_SIDE`, `BOUNDARY_POSITION` | S3 | not a free-standing Region |
| `SlotRequirementPosture` | `REQUIRED`, `OPTIONAL`, `REPEATABLE`, `EXCLUDED`, `OPEN` | S2 | no Obligation, Permission, Prohibition, or absence |
| `PositionSystemKind` | physical, geometric, topological, ordered, network, computational, formal, organizational | S4 | no cross-system inference |
| `InterfacePostureValue` | `OPEN_FOR`, `CLOSED_FOR`, `PARTIAL_FOR`, `UNRESOLVED` | S5 | `ILL_TYPED` is an evaluation error, never a Status value; no normative Access or transition |
| `InterfacePostureEvaluation` | `InterfacePostureValue \| ILL_TYPED` | S5 | evaluation result, not attribution by itself |
| `DependencyMode` | the six modes in S8.3 | S8 | no causation, availability, responsibility, Obligation, Authority, or persistence |
| `StructuralIntegrityValue` | `INTEGRITY_ESTABLISHED`, `INTEGRITY_VIOLATED`, `INTEGRITY_UNRESOLVED` | S8 | Status value only; no moral, legal, or global integrity |
| `IntegrityEvaluation` | `NOT_EVALUATED` or `EVALUATED { StructuralIntegrityValue }` | S8 | `NOT_EVALUATED` is not a fourth Status value |
| `ConstraintEvaluation` | `NOT_EVALUATED` or `EVALUATED { SATISFIED \| VIOLATED \| UNRESOLVED \| ILL_TYPED }` | S8 | evaluation wrapper, not the StructuralConstraint itself |
| `RealizationPosture` | `REALIZATION_ESTABLISHED`, `NONREALIZATION_ESTABLISHED`, `REALIZATION_UNRESOLVED`, `ILL_TYPED` | S1/S9 | not conformance |
| `RealizerPositionKind` | `SINGLE { Bearer }` or `DISTRIBUTED { DistributedRealizerBasis }` | S1 | the Record never substitutes for the actual `REALIZES` relata |
| `IsolatedComponentPosture` | `INSUFFICIENT`, `UNRESOLVED` per distributed bearer | S1 | every bearer must be `INSUFFICIENT` for established Distributed Realization; otherwise use a `SINGLE` alternative or remain unresolved |
| `StructureApplicabilityResult` | `APPLICABLE`, `NOT_APPLICABLE`, `UNRESOLVED`, `ILL_TYPED` | S1/S9 | independent of target actuality and conformance |
| `EvidentialSupportPosture` | `SUPPORTED`, `REFUTED`, `UNDERDETERMINED`, `NOT_EVALUATED` | S9 | posture of evidential support; not Truth or target existence |
| `ConformanceEvaluation` | `NOT_EVALUATED` or `EVALUATED { OntologyConformanceResult }` | S9 | `NOT_EVALUATED` is not a fifth conformance result |

`SubsystemMode`, `NestingMode`, `AttachmentMode`, `ConnectionKind`, `PathKind`, `HostingMode`, and any additional qualifier family are declared extension points. `SubregionMode` is an extension point except for the registered built-in `FORMAL_SUBSET_INCLUSION`. A concrete value is admissible only after a declaration supplies its relation or operator, inputs, algebra or evaluation rule, criterion, Scope, owner, and non-entailments. Until then it is `OPEN`; free text in an actual judgment is `ILL_TYPED`.

Chain, ring, tree, directed acyclic graph, lattice, star, mesh, and similar words are StructuralTopology classifications only after their defining relations are declared. `KernelRole` remains future qualified L1 vocabulary and is not admitted by universal Structure.

---

## S10 — Open Structural Boundary

### S10.1 Structural Placement

An encountered bearer or relation may have established existence while its arrangement remains unknown.

    STRUCTURAL_PLACEMENT(x, SS)
      -> PLACED | PLACEMENT_UNRESOLVED | ILL_TYPED

`PLACED` means the available arrangement can be located under at least one declared Structural Profile without forcing unsupported relations.

`PLACEMENT_UNRESOLVED` means existence or some structural Evidence is available, but no registered Profile adequately resolves the arrangement.

`ILL_TYPED` means the placement request violates the grammar.

Failure to place structurally does not establish nonexistence, disorder, randomness, non-systemhood, or absence of organization.

### S10.2 Unresolved Structural Placement

An **Unresolved Structural Placement** is a Record preserving:

    UnresolvedStructuralPlacement {
      target_term
      existence_posture
      structural_scope
      observed_or_established_bearers
      observed_or_established_relation_instances
      candidate_regions_and_boundaries
      candidate_interfaces
      partial_topology
      candidate_profiles
      rejected_profiles_and_reasons
      coverage
      contradictions_or_tensions
      minimum_distinctions_to_preserve
      evidence
      unresolved
    }

It is not a miscellaneous structural Kind. It is an epistemic and representational position preventing denial, fictionalization, and forced classification.

### S10.3 Structural Extension Rule

A new structural Kind, Role, Relation Kind, Profile family, topology family, or constraint kind may extend this Form only by declaring:

- its imported parent or explicit lack of one;
- category and owner;
- necessary and sufficient conditions or explicit `OPEN` fields;
- admissible bearers and relata;
- complete Relation signature and algebra;
- Structural Scope;
- realization criterion;
- coverage requirements;
- declared disjointness and overlap;
- non-equivalences and non-entailments;
- conformance tests and error results;
- boundary to Constitution, Dynamics, Continuity, and Emergence.

A new name without those declarations is vocabulary awaiting a type, not a structural primitive.

An extension must preserve the imported Ontology and cannot install a later-layer result by definition.

### S10.4 Open Structural Questions

This Form deliberately leaves open:

- which topology families are adequate for unfamiliar physical, biological, cognitive, collective, and derivative systems;
- which candidate Relational Systems possess independently realized State bearers;
- which structural decompositions are most explanatory at each scale;
- how to model arrangements that exceed ordinary set, tree, graph, manifold, or network forms;
- which Architecture Specifications future Worlds will require;
- how particular L1 bearers, kernel-like Roles, Interfaces, and relational structures will instantiate this universal grammar;
- which structural details remain inaccessible at present Resolution;
- whether future evidence requires new relation arities, region kinds, or composition modes.

Openness is typed incompleteness, not license to insert metaphor as fact.

### S10.5 Interface to Dynamics

Dynamics receives:

- exact Configurations;
- typed positions and occupants;
- Interfaces, Ports, Apertures, Channels, and Paths;
- constraints and structural capacities;
- System and World arrangements;
- state-bearing positions;
- topology indexed to an Interval.

Dynamics owns:

- transition rules and state machines;
- signal qualification and intent-axis evaluation of represented signal or Agency Claims under a declared Profile;
- acceptance, hold, refusal, crossing, and commitment Occurrences;
- message delivery and propagation;
- routing execution;
- buffering, eviction, release, and teardown;
- feedback operation and stabilization;
- consensus, ordering, execution-resource metering, failure handling, and operational or computational finality.

Executing an intent-axis criterion may classify inputs and produce Evidence or Records. It does not by scoring alone establish actual Intention, Agency, identity, Consent, or constitutional Admission.

The structural presence of an ingress Interface or Aperture does not perform `signal -> acceptance -> scope`, Admission, or any `0 / hold / 1` lifecycle. Those remain downstream and must preserve this Form’s distinctions.

### S10.6 Interface to Continuity

Continuity receives:

- bearers and System boundaries;
- Configuration and component relations at exact Intervals;
- Structural Positions and replacement Slots;
- realized Architecture and topology;
- Records and provenance.

Continuity owns:

- whether a bearer or System persists across reconfiguration;
- whether replacement preserves identity under a criterion;
- copies, branches, merges, succession, dormancy, recovery, dissolution, and ending;
- lineage and durable attribution across change.

Structure does not declare any Centre, Kernel Role, Anchor, record, or component permanently fixed merely because its position is structurally important.

### S10.7 Interface to Emergence

Emergence receives organized Configurations, Systems, Relational Systems, constraints, Interfaces, feedback-capable arrangements, and multi-scale mappings.

Emergence owns whether those arrangements realize:

- life;
- cognition;
- Agency or Experience;
- intelligence or collective intelligence;
- meaning;
- adaptation, resilience, or self-sustaining capacity;
- coherence as an effective capacity;
- contribution, evaluative value, economic value, or lineage-density effects.

No number of nodes, blocks, turns, relations, or participants becomes intelligence or value by structural counting alone. Structure can expose what is arranged; Emergence must state the criterion under which a new capacity or value relation actually arises.

### S10.8 Interface to Constitution

Constitution may constrain structural matters only where its exact Source Form is validly adopted, Current for the matter, applicable, and every required Subject Binding, Authority, Constitutional Act, and Effect condition is satisfied.

Structure returns only typed Claims and Evidence about actual arrangement. It does not create:

- Adoption or Currentness;
- Admission or Membership;
- Consent, Permission, Prohibition, Right, or Obligation;
- Authority, Ownership, entitlement, jurisdiction, or sovereignty;
- constitutional validity of any structural arrangement.

Topology creates no Authority. Constitutional recognition does not prove structural existence. Constitutional Adoption and Currentness govern constitutional force; they neither create structural actuality nor alter this Form's role as the semantic source for Structure.

### S10.9 Boundary for a Future L1 Structural Profile

The concrete successor L1 must later receive a **qualified Structural Profile**, not be collapsed into that Profile or embedded into universal Structure. Its structural body additionally requires an actual Structural Configuration, sufficient Realization Basis, attributable Structure Description, coverage posture, applicability result, and—only after the gates are satisfied—structural conformance result.

That Profile may declare:

- the exact World and realizer;
- admitted bearer and System positions;
- any kernel-like Role and the bearer occupying it;
- relation, mediator, buffer, and Relational System positions;
- ingress and external-conversion Interfaces and Apertures;
- state-bearing components;
- ledger, network, validator, and infrastructure topology;
- structural constraints and coverage;
- exact seams to its Dynamics, Continuity, and Emergence Forms.

The Profile must not use Structure to pre-decide Admission, transition, durable identity, continuity-turn value, economic distribution, or origin benefit. Those belong to their owning layers and constitutional sources.

### S10.10 The Structural Ground That Remains

Reality can be arranged without being fully described.

An arrangement is not its diagram.  
A Profile is not its realization.  
A position is not its occupant.  
A part is not its whole.  
Containment is not componenthood.  
A Boundary is not sovereignty.  
An Interface is not a crossing.  
An Aperture is not Admission.  
A path is not passage.  
A Relation is not a third bearer.  
A mediator is not a ruler.  
A shared surface is not a shared self.  
A Relational System is not its relata.  
Centrality is not Authority.  
Connectivity is not causation.  
Architecture is not actuality.  
Conformance is not Truth.  
A snapshot is not persistence.  
Organization is not yet Agency, intelligence, or value.

Structure remains exact by arranging what is there, preserving what is distinct, and leaving every transition, persistence claim, and emergent capacity to the layer that can lawfully decide it.

This is the complete universal arrangement grammar presently required before a particular World or L1 can receive its own structural body.
