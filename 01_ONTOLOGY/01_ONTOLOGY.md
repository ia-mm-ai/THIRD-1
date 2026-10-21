# ONTOLOGY

> **Completeness invariant**
>
> Not every future existent is already named. Every future existent can be located without being denied, fictionalized, or forced into the wrong existing kind.

## Dependency Binding

This exact Ontology imports the following exact upstream Form:

| Layer | Canonical Reality Model layer path | SHA-256 | Bytes | Imported function |
|---|---|---|---:|---|
| Substrate | `00_SUBSTRATE/00_SUBSTRATE.md` | `6c3ea987c07840c9f019e6f9bee6e54e9eef6c4465359036c0ab2b455a431b99` | 9,301 | Physical ground and constraint boundary |

The path is relative to the root of the exact Reality Model stack. A changed byte sequence is a different upstream Form and does not replace this binding by recency, filename reuse, or proximity.

## Internal Route and Ownership

This table shows the dependency order inside Ontology. The entire document is physically grounded in the preceding **SUBSTRATE** layer; Ontology does not replace or self-generate that ground.

| Part | Function | Depends on |
|---|---|---|
| I. Metaphysical Commitments | Declares what this model commits to and where primitive definition stops | SUBSTRATE |
| II. Root Grammar | Names the irreducible categories in which reality can be distinguished | Part I |
| III. Classification Grammar | Types, scopes, and tests claims about those categories | Parts I–II |
| IV. Relational Grammar | Types difference, boundary, composition, participation, and realization | Parts I–III |
| V. Domain Ontologies | Specializes the grammar across physical, living, cognitive, agential, social, economic, computational, and formal domains | Parts I–IV |
| VI. Cross-Domain Distinctions | Prevents concepts that travel across domains from collapsing into one another | Parts I–V |
| VII. Epistemic Placement | States how a claim is known without changing what kind of existent it concerns | Parts I–VI |
| VIII. Semantic Firewall | Indexes canonical non-equivalences and non-entailments | Parts I–VII |
| IX. Open Extension Boundary | Locates unresolved existence or kind and defines disciplined extension | Parts I–VIII |

The dependency order is one-way. A later Part may specialize an earlier distinction but may not silently redefine it.

## Part I — Metaphysical Commitments

### Declared Stance

Reality is not made by this ontology.

This ontology declares the root kinds through which reality can be distinguished without confusing a thing with its description, an occurrence with its record, a relation with either of its terms, a role with its bearer, or an unknown with a nonexistence.

It is not a constitution, a protocol, a cosmology, or a theory that one class of being is the measure of all others.

It is the declarative semantic type system of this Reality Model.

Its judgments are exact in form even where reality leaves a boundary undecidable. 

This ontology is therefore:

- **realist** — reality does not depend on being observed, named, recorded, admitted, or believed;
- **physically grounded** — whatever produces physical consequence is physically realized under the constraints of the Substrate;
- **non-reductive** — physical realization does not erase biological, cognitive, semantic, social, institutional, or computational organization;
- **domain-bound** — every type and assertion applies within stated conditions, scale, interval, and resolution;
- **open-world** — absence from the ontology is not proof of absence from reality;
- **closed-definition** — a term has no force beyond the distinctions explicitly given to it;
- **extensible** — a newly encountered kind may be added by declaring its parent kinds, realization conditions, distinguishing properties, relations, domain, and unresolved boundaries.

No protocol term is primitive merely because a protocol uses it. No inherited word is preserved merely because it is familiar. A primitive remains only where removing it would make a distinct class of reality inexpressible.

#### Declared Primitives

Definitions cannot regress forever. This ontology therefore declares the following operators and relations as primitive within this layer.

This Part fixes primitive meaning and arity only. Lower-case argument labels in the signatures are opaque metalanguage positions, not imports of Kinds owned by later Parts. Their admissible types and structured records are supplied only when the owning Parts define them.

| Primitive | Use in this ontology | Boundary |
|---|---|---|
| **EXISTS** | Primitive predicate applied to a candidate under stated contextual bounds | No further existence-synonym is substituted at this layer; it is not inferred from naming, reference, recording, consensus, or proof alone |
| **OBTAINS** | Primitive predicate applied to condition or relational content under stated contextual bounds | No further actuality-synonym is substituted at this layer; it is not inferred from observation, recording, consensus, or proof alone |
| **SAME** | Asserts strict numerical identity: the two terms refer to one and the same existent | Not criterion-relative equivalence, similarity, continuity, or shared name |
| **DIFFERENT** | Denies SAME | Not mere difference of description or value |
| **RELATES** | Connects a reified relation-instance position, relation-kind position, and ordered or unordered relata | Does not fuse the relata |
| **SCOPED_TO** | Binds content to its declared application boundary | Does not imply physical enclosure or universal reach |
| **REALIZES** | Connects a realizing position to a target position under a satisfaction criterion | Not representation, resemblance, or naming |
| **EXPERIENCES** | Connects an experiencer position to experienced occurrence | Its physical basis and distribution remain open; neither position is defined by presupposing the other |

The canonical primitive forms are:

    EXISTS(term_x, domain_D, interval_I)
    OBTAINS(content_p, domain_D, interval_I)
    SAME(term_a, term_b)
    DIFFERENT(term_a, term_b)
    RELATES(relation_term_ri, relation_kind_r, relata[], domain_D, interval_I)
    REALIZES(realizer_a, realized_y, domain_D, interval_I, criterion_C)
    EXPERIENCES(bearer_s, occurrence_e, domain_D, interval_I, resolution_context_R)
    SCOPED_TO(judgment_j, scope_S)

These forms fix arity and argument-position. The labels `term`, `domain`, `interval`, `criterion`, `resolution_context`, and `scope` are mnemonic here, not sorted declarations. Admissible kinds and further constraints are supplied only by the definitions that own those positions.

`SAME` and `DIFFERENT` express strict numerical identity and non-identity. They are not made relative to a Domain, Interval, Criterion, or observer merely because a judgment asserting them is scoped.

`RELATES` is a predicate over a reified Relation Instance, its Relation Kind, and its relata. The surface form

    r(x₁, …, xₙ) OBTAINS @ D, I

is exact existential display sugar for:

    ∃ri [EXISTS(ri, D, I) ∧ RELATES(ri, r, [x₁, …, xₙ], D, I)]

It does not identify relational content with the Relation Instance that bears it.

Where an owning declaration assigns qualifiers to a Relation Kind, the qualified surface form

    r(x₁, …, xₙ ; q₁ = v₁, …, qₘ = vₘ) OBTAINS @ D, I

is exact display sugar for the same `RELATES` judgment over only the ordered relata before the semicolon, together with one typed `ATTRIBUTION(qₖ, ri, vₖ, D, I, R)` for every qualifier after it. Qualifiers are not additional relata and do not change the relation's arity. Omitting a mandatory qualifier is `ILL_TYPED`; moving a qualifier into the relata list denotes a different, generally ill-typed relation application.

Where Domain and Interval are recoverable from an enclosing judgment or explicit `@ D, I`, `EXISTS(x)` and `OBTAINS(p)` are scope-elided display forms, not universal defaults.

`I` denotes a declared Interval in a temporal or otherwise stated order. `Interval` is an argument kind, not a primitive operator.

The primitives generate later definitions without creating duplicate owners:

- **Reality** is owned by Part I and defined below from EXISTS and OBTAINS.
- **Fact** is owned by Part II; its truth condition is OBTAINS in the stated Domain and Interval.
- **Actual**, **Possible**, and **Constraint** are owned by Part VI.
- **Experience**, **Subject**, and **Consciousness** are owned by Part V and derive from the `EXPERIENCES` relation without presupposing a theory of which Bearers satisfy it.

These primitives do not claim metaphysical finality. They mark the point at which this ontology stops replacing one root word with another.

`EXPERIENCES` is declared primitive only as a relation required to refer to subjective occurrence without defining it away. This does not declare which Bearers satisfy it, how it is physically realized, or whether any unfamiliar system is conscious. Those remain empirical and ontological questions.

### Reality and Physical Realization

#### Reality

**Reality** is the totality of Existents for which EXISTS holds and conditions for which OBTAINS holds, whether known or unknown, observed or unobserved, represented or unrepresented.

Reality is not true or false. Claims about reality may be true or false.

Reality does not become more real when recorded and does not become unreal when inaccessible.

#### Physical Reality

**Physical Reality** is reality insofar as it is physically realized and constrained by the physical ground declared in the Substrate.

Nothing that acts physically, stores a distinction, performs a computation, carries a signal, or produces a physical effect stands outside physical realization.

Physical realization does not by itself settle whether every real relation, meaning, experience, norm, or formal structure is exhaustively reducible to a physical vocabulary.

## Part II — Root Grammar

### Root Ontological Grammar

Reality is not composed only of objects. The minimum grammar requires several non-interchangeable axes.

| Axis | It answers | It does not imply |
|---|---|---|
| Existent | What actually exists? | Objecthood or observation |
| Entity | What has distinguishable unity as a concrete or organized existent? | Agency or permanence |
| Occurrence | What happens? | A durable bearer |
| Property / Quantity | How is it, and how much under a measure? | A separate entity |
| Relation Kind / Relation Instance | How may relata stand, and which typed Relation actually obtains? | Fusion or shared identity |
| Pattern / Form | How is difference organized? | One unique realization |
| State | What condition obtains for a bearer? | The bearer or its record |
| Fact | What obtains? | A sentence or consensus |

Bearer, Kind, Role, Status, Domain, Context, Scale, and Resolution classify or scope these root axes. They are owned by Part III.

#### Existent

An **Existent** is anything for which EXISTS holds as an Entity, Occurrence, Property, Relation Instance, Pattern, State, or physically realized Representation.

Existent is the widest ontological classification in this model. It does not mean object, substance, observer, person, or independently bounded individual.

Existence does not depend on discovery. Classification as an Existent is answerable to reality, not conferred by the classifier.

#### Entity

An **Entity** is an Existent with distinguishable concrete or organizational unity under a stated boundary, domain, and interval.

An entity may be elementary or composite, localized or distributed, transient or durable, natural or artificial, living or nonliving, agential or non-agential.

Entity does not mean solid object, owner, person, agent, or permanent identity.

#### Occurrence

An **Occurrence** is an Existent whose identity includes happening, changing, or being maintained across an ordered extent.

Events, processes, Transition Occurrences, interactions, measurements, computations, actions, and cessations are occurrences.

An occurrence is not the record made of it.

##### Time and Temporal Order

**Physical Time** is the temporal order and metric structure physically realized in spacetime as constrained by the Substrate.

An **Ordered Time** is a derived ordering used by a formal, computational, biological, institutional, or other system. Its order may be realized through physical time without reproducing every property of physical time.

An **Interval** is a bounded extent in a temporal or otherwise declared order.

Its order, endpoints or openness, scale, and resolving conditions must be stated where they matter. A physical time interval is one specialization; an ordered extent in a formal or computational domain is another.

Occurrence does not define time, and time is not inferred from a record of occurrence.

#### Property

A **Property** is an attribute whose value characterizes an admissible Existent or other declared Bearer under stated conditions.

A property may be intrinsic under a stated theory, relational, dispositional, collective, emergent, frame-dependent, state-dependent, or assigned by an institutional system.

A property is not a separate object merely because language gives it a noun.

#### Quantity

A **Quantity** is a property represented so that values can be compared under a defined dimension, unit, reference, method, and domain.

A **Quantity Value** is the value assigned to a quantity together with the information required to interpret it.

Quantity, value, numeral, unit, measurement, and measured entity are distinct.

#### Relation

A **Relation Kind** is a typed n-ary structure declaring how two or more argument-positions may stand with respect to one another.

A **Relation Instance** (`RelationInstance`) is an Existent in which that structure actually OBTAINS among admissible relata in a stated Domain and Interval. It is not reducible to a private Property of either relatum, and it does not fuse its relata into one Entity.

Relation Kinds may be spatial, temporal, causal, compositional, comparative, semantic, social, institutional, computational, mathematical, or otherwise typed.

A Relation Kind declares arity and the Kind admissible at each argument-position. A Relation Instance may itself be a Bearer of Properties such as duration, strength, direction, Provenance, or Status.

#### Pattern

A **Pattern** is an organization of differences and relations that can be distinguished across an extent, interval, or more than one realization.

The same pattern may have different realizations. No one realization contains the pattern in every possible form.

Similarity to a pattern does not prove identity, common cause, communication, or shared meaning.

##### Form

**Form** is the distinguishable organization through which an entity, state, occurrence, or representation has structure.

Form may be spatial, material, biological, logical, semantic, institutional, computational, or mixed.

The same form may have multiple realizations. One realization may change form. Form is not matter, carrier, content, identity, or value, though each may constrain it.

An effective form must be realized. A represented form is not thereby instantiated as the kind it represents.

#### State

A **State** is a structured condition attributed to a Bearer under a stated domain, interval, scale, and resolution.

A state is specified by selected properties and relations. It need not contain every detail of the bearer.

The bearer, its state, the representation of that state, and the record of that representation are distinct.

#### Fact and Obtaining

A **Fact** is the status of content p for which OBTAINS(p) holds in the stated Domain and Interval. Fact is not an additional Entity inserted between Reality and the content that obtains.

A fact is not a sentence. A sentence may represent a fact.

The existence of a claim is one fact. Whether the content claimed obtains is another.

### Occurrence, State, and Change

#### Configuration

A **Configuration** is a particular arrangement of selected parts, values, and relations.

A configuration may instantiate a state. Different micro-configurations may correspond to one macrostate.

#### Event

An **Event** is a bounded occurrence situated in spacetime or in a derived ordered domain.

An event may be nearly instantaneous at one resolution and extended at another.

#### Process

A **Process** is an occurrence extended through an interval in which states, relations, or organization change or are maintained.

Maintenance is a process. It is not the absence of change.

#### Transition Occurrence and Transition Relation

A **Transition Occurrence** is an Occurrence in which the relevant State of a Bearer changes under a stated criterion.

A **Transition Relation** is the typed before–after relation between the prior State and subsequent State under that Occurrence, Process, or Rule.

A **Transition Record** is a Representation of a Transition Occurrence or Transition Relation. It is neither the Occurrence nor the Relation itself.

#### Trajectory

A **Trajectory** is an ordered path through physical space, a state space, or another defined possibility space.

A trajectory requires a specified space and ordering. It does not imply intention.

#### Beginning and Ending

A **Beginning** is the boundary at which an Entity first exists, an Occurrence first occurs, or a Relation, State, condition, or Status first obtains under a stated criterion.

An **Ending** is the boundary at which the relevant Entity ceases to exist, Occurrence ceases, or Relation, State, condition, or Status ceases to obtain under that criterion.

Creation, formation, activation, birth, admission, execution, death, dissolution, expiry, and invalidation are different beginning or ending types.

## Part III — Classification Grammar

### Typing and Conformance

#### Judgment Form

The canonical typing judgment is:

    Γ ⊢ x : K @ D, I, R

A **Term** is a sign in this metalanguage that refers or purports to refer to something under examination. A **Judgment** is a typed assertion made with this grammar. Terms and judgments are represented expressions; writing one does not establish that its purported referent exists or that its content obtains.

It reads:

> Under declaration and evidence context Γ, x conforms to kind K in domain D, during interval I, under Resolution Context R.

Every material type claim must make these fields recoverable:

- **Γ — Context:** declarations, definitions, evidence, and assumptions used by the judgment;
- **x — Term:** the bearer, occurrence, relation, state, form, or representation being typed;
- **K — Kind:** the declared type whose conditions are being tested;
- **D — Domain:** where the type claim applies;
- **I — Interval:** when or over which ordered extent it applies;
- **R — Resolution Context:** the separately declared scale, reference frame, granularity, resolution, and uncertainty relevant to the distinction.

Omitted fields are not universal defaults. They are unresolved dependencies.

This judgment is **semantic**, not automatically algorithmic. It becomes machine-decidable only for kinds whose declaration supplies operational predicates at the stated resolution. Where a condition is observationally, conceptually, or physically unresolved, the correct result is `UNRESOLVED`, not a guessed Boolean.

#### Type Relations

The ontology uses the following typed relations.

##### Subtype

    K SUBTYPE_OF L

Every instance conforming to K also conforms to L in the same judgment context.

Subtype inheritance is valid only where domain, interval, and resolution remain compatible.

##### Instantiation

    x INSTANCE_OF K @ D, I, R

x satisfies the declared conditions of K in that scope.

Instantiation is owned by Classification Grammar. Later appearances of “instance” must use this relation or a domain-qualified subtype such as `ProgramInstance`; they may not silently redefine it.

##### Disjointness

    K DISJOINT_WITH L @ C

No one instance may conform to K and L in the same respect under context C.

Disjointness must be declared. Difference of names does not establish it.

##### Type Overlap

    K OVERLAPS L @ C

Some instance may conform to both K and L under context C while the kinds remain non-equivalent.

##### Role Occupancy

    x PLAYS role IN context DURING I

Role occupancy does not add a subtype relation between the bearer’s kind and the role.

##### Property or Status Attribution

    a OF x HAS value @ D, I, R

This is exact display sugar for:

    ATTRIBUTION(a, x, value, D, I, R)

The Property or Status, bearer, value-space, domain, interval, measurement or assignment rule, and uncertainty must be type-compatible.

Relation instantiation, realization, experience, and representation receive their specialized signatures only in the later Parts that own them. Classification Grammar supplies the shared judgment and conformance machinery; it does not import their later-owned argument kinds here.

#### Inference Rules

The minimum valid inferences are:

1. **Subtype inheritance** — if x conforms to K and K is a subtype of L in the same scope, x conforms to L.
2. **Declared disjointness** — if x conforms to K and K is disjoint with L in the same scope, x does not conform to L there.
3. **Role non-cast** — from x playing role r, no new intrinsic kind of x follows unless a separate rule declares it.
4. **Relation non-cast** — from x standing in a declared relation to y, no new intrinsic kind, existence claim, or identity follows for either term unless the relation signature or another explicit rule supplies it.
5. **Content non-cast** — from a representation, record, or claim presenting p, neither the existence of a purported referent nor `OBTAINS(p)` follows without an independent supporting judgment.
6. **Domain preservation** — a judgment valid in D does not travel to D₂ without a declared mapping and preserved conditions.
7. **Temporal preservation** — a judgment at I does not automatically remain valid at I₂.
8. **Open-world non-inference** — failure to derive x : K does not derive x : NOT-K.
9. **Explicit negative judgment** — non-conformance requires a failed necessary condition, declared disjointness, or evidence supporting the negative claim.

#### Conformance Results

A conformance judgment returns exactly one of four result classes:

- **CONFORMS** — the required conditions are satisfied in the stated scope;
- **DOES_NOT_CONFORM** — at least one required condition is shown not to hold;
- **UNRESOLVED** — available context is insufficient to decide;
- **ILL_TYPED** — the judgment itself violates the declared grammar.

UNRESOLVED is not DOES_NOT_CONFORM.

A later judgment may replace an earlier epistemic result when evidence changes. It does not retroactively change what actually obtained during the earlier interval.

No prose boundary is assumed decidable merely because it is written. Until an operational predicate or sufficient evidence exists, a disputed boundary remains `UNRESOLVED`.

#### Error Model

An ill-typed expression must fail explicitly with one or more of these error classes:

- **UNDECLARED_TERM** — a term has no definition in the applicable ontology;
- **CATEGORY_ERROR** — a term is used as a different ontological category without a declared cast;
- **SUBTYPE_ERROR** — an undeclared inheritance is assumed;
- **DISJOINTNESS_ERROR** — one term is assigned disjoint kinds in the same respect;
- **ARITY_ERROR** — a relation receives the wrong number of relata;
- **DOMAIN_ERROR** — a relatum does not conform to the kind required at its relation-position;
- **RANGE_ERROR** — a value or result lies outside the declared target kind or value-space;
- **SCOPE_ERROR** — a judgment is exported beyond its declared Scope;
- **TEMPORAL_ERROR** — a judgment is exported beyond its interval;
- **RESOLUTION_ERROR** — distinctions are asserted beyond the resolving conditions;
- **ROLE_CAST** — a contextual role is treated as an intrinsic kind;
- **REPRESENTATION_CAST** — a representation is treated as its referent;
- **RECORD_CAST** — a record is treated as the occurrence or fact recorded;
- **REALIZATION_GAP** — an effective form or capacity is asserted without a declared realizer and realization criterion;
- **AMBIGUOUS_OPERATOR** — one relation sign is used for materially different relations without typing them;
- **CONTRADICTION** — a claim and its exact negation are both asserted to obtain in the same Domain, Interval, Resolution Context, and interpretation.

Errors remain representable as facts about an expression or execution. Rejecting an ill-typed judgment does not erase its occurrence.

#### Kind Declaration Schema

Every Kind is answerable to one canonical declaration schema:

    KIND K {
      parents: [K₁, ...] | NONE
      category: Entity | Occurrence | Property | Relation | Pattern | State | other
      necessary_conditions: [c₁, ...] | OPEN
      sufficient_conditions: [s₁, ...] | OPEN
      properties_and_value_spaces: [p₁: V₁, ...] | NONE | NOT_APPLICABLE | OPEN
      admissible_bearers: [B₁, ...] | NOT_APPLICABLE | OPEN
      relation_signatures: [r₁, ...] | NONE | NOT_APPLICABLE | OPEN
      realization_criterion: C | NOT_APPLICABLE | OPEN
      scope: S | OPEN
      disjoint_with: [L₁, ...] | NONE | OPEN
      overlaps: [M₁, ...] | NONE | OPEN
      epistemic_statuses: [E₁, ...] | NONE | OPEN
      non_equivalences: [N₁, ...] | NONE | OPEN
      later_layer_boundaries: [BND₁, ...] | NONE | OPEN
      owner: ontology_part
    }

The schema is normative for a machine-projectable conformance profile. This human-readable ontology supplies a minimal projection of the root kinds below, not a complete conformance registry. A later profile must expand every admitted Kind into all schema fields. Semantic definitions supply their available fields in prose and leave every unresolved field explicitly `OPEN` until that projection occurs. Prose may not silently contradict a projected record.

`NONE` means a known empty set. `OPEN` means unresolved. `NOT_APPLICABLE` means that the field does not apply to the declared category. They are not interchangeable and none is an invisible default.

An admitted Kind must still declare its category, owner, and parent list or explicit `NONE`, and must provide at least one discriminating necessary condition, sufficient condition, property/value-space constraint, relation signature, or realization criterion. An all-`OPEN` name is vocabulary awaiting a declaration, not an admitted Kind.

#### Classification-Grammar Relation Signatures

| Relation | Domain positions | Range or result | Owner |
|---|---|---|---|
| `INSTANCE_OF(x, K, D, I, R)` | x: Term; K: Kind | Conformance result | Classification Grammar |
| `ATTRIBUTION(a, b, v, D, I, R)` | a: Property or Status; b: Bearer; v: compatible declared value-space | Typed attribution content | Classification Grammar |
| `PLAYS(b, role, context, I)` | b: Bearer; role: Role | Role occupancy | Classification Grammar |
| `SCOPED_TO(j, S)` | j: Judgment; S: Scope | Scoped judgment | Classification Grammar |

Relations owned by later Parts receive their specialized domain and range constraints only in their owning definitions. Their primitive forms fix notation and arity without importing later-owned kinds into Part III.

Specialized relation predicates that omit a named relation-instance argument—including `ATTRIBUTION` and `PLAYS`—are existential projections of `RELATES`; they assert that a corresponding Relation Instance obtains without returning or naming it. A named instance must use the full `RELATES` form.

#### Root Type Commitments

- `Entity SUBTYPE_OF Existent`
- `Occurrence SUBTYPE_OF Existent`
- `Property SUBTYPE_OF Existent`
- `RelationInstance SUBTYPE_OF Existent`
- `Pattern SUBTYPE_OF Existent`
- `State SUBTYPE_OF Existent`
- `Entity SUBTYPE_OF Bearer`
- `Occurrence SUBTYPE_OF Bearer`
- `RelationInstance SUBTYPE_OF Bearer`
- `State SUBTYPE_OF Bearer`

No disjointness among root axes is inferred from different names. A State may itself bear a Property; an Occurrence may instantiate a Relation; a Pattern may be realized by an Entity. Such overlap must preserve the respect in which each classification applies.

#### Canonical Root Registry — Minimal Projection

This table is the minimal machine-projectable registry. Later definitions are projected using the Kind Declaration Schema; a field not resolved by the definition remains `OPEN`.

| Kind | Parent commitments | Category | Necessary condition | Sufficient condition | Realization field | Owner |
|---|---|---|---|---|---|---|
| Existent | `NONE` | other | `EXISTS(x, D, I)` | `EXISTS(x, D, I)` | actual mode may be physical, organized, relational, or `OPEN` | Part II |
| Entity | Existent, Bearer | Entity | distinguishable concrete or organizational unity under declared unity conditions, Domain, and Interval | same, at declared Resolution Context | physical or organized unity criterion | Part II |
| Occurrence | Existent, Bearer | Occurrence | ordered happening, change, or maintenance | same, within declared Interval | process or event criterion | Part II |
| Property | Existent | Property | attributable character with an admissible Bearer and scope | declaration-specific | carrier or assignment criterion | Part II |
| Quantity | Property | Property | comparable values under dimension, unit, reference, method, Domain | same | measurement or assignment criterion | Part II |
| RelationInstance | Existent, Bearer | Relation | typed arity and admissible relata; a conforming `RELATES` predicate obtains | same | relation-specific | Part II |
| Pattern | Existent | Pattern | organized distinguishable differences and relations across extent, Interval, or realizations | declaration-specific | mapping and preservation criterion | Part II |
| State | Existent, Bearer | State | selected Properties and Relations attributed to a Bearer at scope | same | state-selection criterion | Part II |
| Fact | `NONE` | other | `OBTAINS(content, D, I)` | same | `NOT_APPLICABLE` | Part II |
| Bearer | Existent | other | admissible argument of `ATTRIBUTION`, `PLAYS`, or a Relation position | relation-specific | `NOT_APPLICABLE` | Part III |
| Kind | `NONE` | other | conforms to Kind Declaration Schema | accepted declaration record | declaration-specific | Part III |
| Role | `NONE` | other | Bearer, Context, defining relation or operation, Interval | obtaining `PLAYS` relation | context- or relation-specific | Part III |
| Status | `NONE` | other | criteria and assigning or recognizing context | criteria satisfied in scope | context-specific | Part III |
| Domain | `NONE` | other | explicitly delimited conditions of application | same | `NOT_APPLICABLE` | Part III |

The registry does not force every future Kind under one exclusive parent. Multiple inheritance and overlap are permitted; hidden inheritance is not.

### Classification Primitives

Kinds classify. Bearers receive attribution. Roles locate bearers in contexts. Statuses classify conditions under criteria.

#### Bearer

A **Bearer** is an Existent occupying an admissible bearer-position in property or Status attribution, state description, Role occupancy, or a Relation.

Bearer is the widest attribution kind. Entities, occurrences, instantiated relations, states, and other explicitly admitted Existents may each be bearers of properties appropriate to their category.

Bearer does not mean object, owner, subject, agent, or permanent individual.

The canonical attribution signature is:

    ATTRIBUTION(attribute, bearer, value, domain, interval, resolution_context)

The attribute is a Property or Status. Its admissible bearer kind is part of its declaration; Roles and other Relations use their own typed predicates rather than being silently cast as attributes.

#### Domain

A **Domain** is the bounded set of conditions within which a distinction, type, quantity, rule, model, or relation has defined application.

A domain may be physical, spatial, temporal, biological, semantic, institutional, computational, formal, or mixed.

The reach of a definition is part of the definition.

#### Kind

A **Kind** is a defined way of classifying entities, occurrences, properties, relations, patterns, or states by shared conditions.

An **Instance** is something to which a kind applies.

A kind does not create its instances. Classification does not exhaust the classified.

One existent may instantiate multiple compatible kinds at once. Kinds need not form one exclusive hierarchy.

#### Role

A **Role** is a context-dependent position occupied by a bearer in a relation, process, system, or institution.

A role requires:

- a bearer;
- a context;
- a defining relation or function;
- an interval or occurrence in which it applies.

A bearer may occupy several roles. A role may be occupied by different kinds of bearer.

Role does not equal kind. Role assignment does not transform the bearer into a new species of entity.

#### Status

A **Status** is a classified condition assigned or recognized under a defined system of criteria.

Physical phase, epistemic status, legal status, protocol status, and social status are different kinds of status.

A status may describe a reality, create an institutional consequence, or do both. Its effect depends on the system that constitutes it.

### Domain, Context, Scope, and Resolution

#### Scope

A **Scope** is the structured application boundary of a judgment, relation, occurrence, role, rule, status, action, effect, or other declared content. It consists at minimum of its Domain, Interval, Resolution Context, and any further declared inclusion, exclusion, or interpretation conditions.

In the metalanguage, a scope position has this record projection:

    Scope {
      domain
      interval
      resolution_context
      further_conditions: [...] | NONE | OPEN
    }

Scope is not interchangeable with Domain alone.

#### Criterion

A **Criterion** is a represented condition-set and evaluation relation under which a condition, distinction, relation, classification, occurrence, action, effect, status, or judgment is assessed.

In the metalanguage, a criterion position has this record projection:

    Criterion {
      condition_set
      evaluation_relation
      scope
      method: M | NOT_APPLICABLE | OPEN
    }

A Criterion is distinct from its satisfaction result and from any separately declared warrant, permission, or governing relation.

#### Resolution Context

A **Resolution Context** is the structured set of resolving conditions under which a distinction or Judgment is made:

    ResolutionContext {
      scale
      reference_frame
      granularity
      resolution
      uncertainty
    }

These fields remain distinct. Each must carry a declared value, `OPEN`, or `NOT_APPLICABLE` where the field type permits it. Operational field types and decision procedures belong to a later conformance profile.

#### Context

A **Context** is the set of conditions relevant to interpreting or evaluating an entity, relation, occurrence, signal, statement, or action.

Context may include domain, time, place, participants, reference frame, purpose, rules, prior state, and available evidence.

Context does not license arbitrary meaning. It bounds which meaning is in use.

#### Scale

A **Scale** is a characteristic range of size, duration, energy, number, complexity, or resolution at which distinctions are made.

Microscopic and macroscopic descriptions may both be real and valid without using the same variables.

Smaller does not mean more real. Larger does not mean merely subjective.

#### Resolution

**Resolution** is the fineness with which differences can be distinguished.

Resolution is determined by the observer or instrument, representation, resources, method, and physical conditions.

Unresolved detail is not erased from reality.

#### Granularity

**Granularity** is the level at which a model groups differences into units.

A grain may be a particle, cell, organism, transaction, sentence, block, institution, or other declared unit.

Granularity is representational unless the chosen units correspond to physically or organizationally maintained boundaries.

#### Reference Frame

A **Reference Frame** is a structure relative to which positions, durations, velocities, orientations, or other relational quantities are assigned.

A frame does not create what it describes. No single frame is automatically the universal view.

#### Perspective

A **Perspective** is the bounded access available to a sensing, cognitive, or representational system from its position, capacities, history, and context.

Perspective is not reality itself. Difference of perspective does not imply that all descriptions are equally correct.

## Part IV — Relational Grammar

### Difference, Distinction, and Boundaries

#### Difference

A **Difference** is non-identity between conditions, values, arrangements, entities, or occurrences.

Difference exists before any observer describes it. A detected difference is a difference brought into relation with a detector.

#### Distinction

A **Distinction** is a difference resolved under a criterion.

A physical distinction requires physical realization. A conceptual distinction requires a represented rule that separates what counts as one case from another.

Every distinction has a resolution. Differences below that resolution remain unresolved, not necessarily absent.

#### Boundary

A **Boundary** is a distinction between an inside and an outside under a stated criterion.

A boundary may be:

- physical;
- spatial;
- temporal;
- material;
- functional;
- informational;
- organizational;
- semantic;
- institutional;
- computational;
- model-selected.

A boundary need not be a wall. It need not prevent exchange. It does not by itself create ownership, privacy, isolation, jurisdiction, or agency.

#### Interface

An **Interface** is a typed boundary relation through which specified differences may cross, be translated, or produce effects.

An interface defines what can encounter it, what it can receive, what it can emit, and under which conditions.

An interface does not make two domains identical.

#### Seam

A **Seam** is an interface or mapping that connects distinct domains while preserving the distinctions required on each side.

A coherent seam declares:

- the source domain;
- the receiving domain;
- the admissible inputs;
- the mapping or transformation;
- what is preserved;
- what may be lost;
- what may not cross;
- where consequence occurs.

A seam is not successful when it hides mismatch. Its function is to house mismatch without collapse.

### Composition and Organization

#### Object

An **Object** is an entity treated as a bounded unit at a stated scale.

Objecthood may arise from a stable physical boundary, cohesive organization, functional integration, or a model-selected boundary.

Object does not imply indivisibility, permanence, agency, or ownership.

#### System

A **System** is an entity or organized collection of entities and relations treated as a whole under an explicit boundary.

A system includes not only components but also the relations through which the whole is organized.

A component list is not a complete description of a system.

#### Part and Whole

A **Part** is an entity or structure that enters the composition or organization of a whole.

A **Whole** is an entity or system constituted through parts and their relations.

Parthood is typed. A molecule is materially part of a cell; a rule may be formally part of a protocol; a member may be institutionally part of an organization. These are not the same relation.

#### Component and Constituent

A **Constituent** enters the composition of a whole.

A **Component** is a distinguishable part assigned a function or position within a system description.

A component may be replaced while the containing system continues under a later continuity criterion. Whether it is the same system belongs to Continuity.

#### Aggregate

An **Aggregate** is a collection whose members are grouped without requiring the integration that would constitute a system.

Spatial proximity or shared classification alone does not make a system.

#### Organization

**Organization** is the structured arrangement of parts, relations, constraints, and processes through which a whole has capacities not attributable to an unarranged list of its parts.

Organization is physically realized but not identical to the material of any one component.

#### Environment

An **Environment** is the relevant remainder relative to a stated system boundary.

The environment may supply matter, energy, signals, perturbations, constraints, and opportunities.

Environment does not mean irrelevant outside.

#### Thermodynamic Openness and Closure

A **Thermodynamically Open Physical System** may exchange both Matter and Energy across its Boundary.

A **Thermodynamically Closed Physical System** does not exchange Matter under the stated model, though it may exchange Energy.

A **Thermodynamically Isolated Physical System** exchanges neither Matter nor Energy within the precision and Interval stated.

These are boundary conditions, not moral or institutional statuses. Perfect isolation is generally an idealization.

#### Structure

**Structure** is the arrangement of entities, parts, relations, boundaries, and constraints at a stated time or abstraction.

Ontology names structure as a kind of organization. The complete laws of arrangement belong to the Structure layer.

#### Function

A **Function** is a contribution made by a part, process, or capacity within a containing organization or selected account.

Biological function may arise through evolutionary history and organization. Designed function may arise through an agent’s purpose. System function may be assigned by a model.

Function does not always imply conscious purpose.

#### Mechanism

A **Mechanism** is an organized set of entities, relations, and processes through which a specified change or capacity is realized.

A mechanism is not merely a verbal explanation. It identifies how consequence is physically or institutionally produced.

### Relation, Interaction, and the Between

#### Relation Instantiation

The canonical signature is:

    RELATES(relation_term, relation_kind, relata[], domain, interval)

The relation Kind declares arity and the Kind admissible at every argument-position. The first argument is a Term proposed as the reified relation instance. Where `EXISTS(ri, D, I)` and `RELATES(ri, r, relata, D, I)` obtain, `ri` conforms to `RelationInstance`; the predicate does not require that classification in advance.

The surface form `r(x₁, …, xₙ) OBTAINS @ D, I` uses the exact existential expansion declared in Part I. No unnamed relation instance is silently treated as a non-existent term.

#### Relata

A **Relatum** is a bearer occupying a term-position in a relation.

The relation, its relata, and the representation of the relation are distinct.

#### Directed and Undirected Relations

A **Directed Relation** distinguishes source and target or another ordering among its terms.

An **Undirected Relation** is invariant under the exchange defined by its terms.

Direction is a structural property. It does not imply purpose or moral asymmetry.

#### Interaction

An **Interaction** is a relation realized through mutual or one-way influence among entities, states, or processes.

An interaction requires a mechanism or coupling. Similarity, simultaneity, shared labels, or matching frequency do not create interaction.

##### Medium

A **Medium** is a physical or organized system within which a specified process, propagation, or interaction is realized.

Air can be a medium for sound; a material can be a medium for deformation; a network can be a medium for communication. A medium is not the content it carries.

Not every propagation requires a material medium. Electromagnetic fields propagate in vacuum.

##### Channel

A **Channel** is a bounded path or coupling through which matter, energy, signal, or influence can pass under stated conditions.

A channel has capacity, noise, delay, direction, and failure conditions where relevant.

The existence of a Channel does not prove that a Message, Intention, Authority-bearing Decision, or causal influence passed through it.

##### Effect

An **Effect** is a difference that obtains in a state, relation, capacity, or process because of an occurrence under a causal or constitutive account.

An effect may be intended or unintended, direct or mediated, local or distributed, reversible or irreversible.

Effect does not imply authorization, legitimacy, benefit, or truth.

##### Consequence

A **Consequence** is an effect considered in relation to the occurrence, action, rule, or condition from which it follows.

Consequence is a typed relation. It may be physical, biological, cognitive, social, institutional, computational, or economic.

##### Impact

**Impact** is the extent or significance of one or more effects under a stated criterion, affected domain, and interval.

Impact is evaluated, not a substance emitted by an actor. It is not identical to effort, intention, contribution, popularity, or price.

#### Coupling

A **Coupling** is a typed connection through which the possibilities or changes of one system depend on another.

Coupling strength, channel, direction, and domain must be stated where relevant.

#### Correlation

A **Correlation** is structured dependence among values, states, or outcomes.

Correlation may arise from interaction, common conditions, shared constraint, selection, or a jointly prepared state.

Correlation alone does not identify a cause, communication channel, or controlling agent.

#### Mediation

**Mediation** is a relation in which a third entity, process, representation, or institution conditions an interaction between other relata.

A mediator has no automatic Normative Authority over the relata merely because it carries or stabilizes their Relation.

#### Relational System

Where relata and their interactions are organized as a bounded system bearing its own physically realized state variables and constraints, a **Relational System** may exist in addition to the relation itself.

Recognition of a Relational System does not require prior persistence or accumulated history. Whether it persists or continues across change is a Continuity judgment, not a condition of being a Relational System at one Interval.

A relational system is not automatically:

- a person;
- an agent;
- either participant;
- a fusion of participants;
- an owner of their states;
- a shared jurisdiction.

This distinction allows the between to be real without converting it into a hidden sovereign.

#### Relational Overlap

An **Overlap** is a domain, part, property, or interval shared under a stated criterion.

Overlap does not erase what remains distinct. Shared state exists only where a jointly constituted state bearer actually exists.

#### Participation

**Participation** is the typed relation through which a Bearer takes part in a specified Relation, Occurrence, System, or Institution in a declared mode, scope, and interval.

    PARTICIPATES(bearer, context, participation_mode, scope)

A **Participation Mode** (`ParticipationMode`) is the relation-specific classification of how a Bearer participates. It is not the physical `Mode` defined for waves, and it is not interchangeable with Role, Agency, Scope, effect, or consequence.

Participation may be causal, material, informational, experiential, agential, social, institutional, computational, or mixed. The Relation does not require every Participant to be an Agent, does not make the Participants identical, and does not by itself create Membership, Consent, Ownership, submission, any Authority relation, or entitlement.

A **Participant** is the Role occupied by the first argument of an obtaining Participation relation.

### Realization, Embodiment, and Execution

Realization is the load-bearing relation between what obtains and that through which it obtains.

#### Realization

**Realization** is the primitive typed relation that obtains where a physical or organized Bearer or process satisfies the declared realization criterion by which a target obtains through that realizer. The range is open to any Kind whose owning declaration admits realization; Part IV does not enumerate or import later-owned kinds.

The canonical signature is:

    REALIZES(realizer, realized, domain, interval, criterion)

Every realization claim must identify:

- the **realizer** — the physical entity, system, state, or process doing the realizing;
- the **realized** — the target whose declared realization criterion is claimed to be satisfied;
- the **domain** — the conditions within which the realization relation applies;
- the **interval** — when it obtains;
- the **criterion** — the mapping and invariants by which the realized organization is distinguished in the realizer.

Realization may be distributed across many components and relations. The realizer need not be spatially localized in one object.

One form may have multiple realizers. One realizer may realize multiple compatible forms. Multiple realizability does not make realization disembodied.

A realization relation either obtains or does not obtain under its criterion. The following adjacent kinds must not be folded into that Boolean relation:

- a **Realization Attempt** is an Occurrence directed toward satisfying a realization criterion;
- a **Realization Result** is a representation of whether the criterion was satisfied;
- a **Partial Realization** realizes a declared proper subset or lower-resolution form, not the complete target under a failed criterion;
- a **Latent Capacity** is an actually realized Capacity whose characteristic process is not presently active;
- an **Active Realization** is a realization whose characteristic process presently occurs;
- a **Distributed Realization** has a plural or distributed realizer; no isolated component alone satisfies the criterion.

A failed Realization Attempt is still an actual Occurrence. It does not make the attempted target actual and does not instantiate `REALIZES` for that target.

REALIZES is not inferred from resemblance, description, simulation, naming, ownership, or intended purpose.

#### Embodiment

**Embodiment** is a Realization relation in which a capacity, process, organization, or subject is realized through an organized Physical Body during an interval.

    EMBODIES(body, realized, domain, interval, criterion)

Embodiment does not mean that every realized capacity is localized in one component or that a body is merely a replaceable container.

#### Implementation

**Implementation** is an Artifact or organized arrangement produced or configured to realize a specification, Protocol, design, or formal operation in a technical or organizational System.

A **Conforming Implementation** satisfies the applicable specification and realizes the declared organization under its criterion. An Implementation may fail to conform; in that case the Artifact or arrangement remains actual, but the claimed realization does not obtain. The specification, Implementation, deployment, and Execution are distinct.

#### Execution

**Execution** is an occurrence in which an implemented machine or organization realizes specified transitions.

Execution realizes a process. It does not by itself establish authorization, correctness, intention, consent, or external effect.

Computational Execution and Organizational Execution are domain-qualified subtypes. They do not redefine this generic occurrence.

#### Constitutive Relation

A **Constitutive Relation** is the relation through which parts, organization, boundary conditions, and dependencies jointly realize a whole as a kind.

    CONSTITUTES(parts_and_relations, whole, domain, interval, criterion)

Constitution-of-a-whole is not aggregation. Ontology names this generic relation; Structure declares particular arrangements, while the Constitution layer governs the constituting conditions and normative force of a particular world where applicable.

#### Realization and Representation

Realization and representation are independent relations.

- A map represents a territory without realizing a territory.
- Hardware may realize a computation whose state represents weather without realizing weather.
- A legal record may represent ownership while an institution realizes the ownership relation.
- A model may represent an agent without realizing agency.
- A simulation may independently realize cognition if the criteria for cognition obtain; it does not do so merely because cognition is represented.

Every effective claim about information, computation, institution, economy, intelligence, or world must be able to answer: **what realizes it, where, for how long, and under which criterion?**

## Part V — Domain Ontologies

### Physical Kinds

The Substrate declares the physical constraints. This section names the principal physical referents, conditions, relations, patterns, and attributes to which those constraints apply.

Current physical theories are highly tested within their domains. Their present reduction boundaries are not declared metaphysically final.

#### Spacetime

**Spacetime** is the physical metric and causal structure by which occurrences have location, duration, interval, separation, and possible causal relation.

In general relativity, spacetime geometry is dynamical rather than a fixed universal container.

A **Spacetime Region** is a specified extent of spacetime. A region is not necessarily an object, system, enclosure, or empty container.

A **Coordinate System** assigns labels to positions. Coordinates are representations, not spacetime itself.

Spacetime may not be fundamental at every scale. Its quantum constitution remains unresolved.

##### Space

**Space** is the spatial structure through which entities and occurrences have position, separation, extent, orientation, and spatial relation within a reference frame.

Space is not an empty substance or universal container independent of spacetime in current relativistic physics.

##### Relativistic Temporal Structure

**Relativistic Temporal Structure** is Physical Time as represented by relativistic spacetime: the structure through which physical occurrences have order, proper duration, interval, and causal accessibility.

No observer-independent global present is supplied by relativity. Clock readings are physical measurements along worldlines, not time itself.

##### Location, Physical Interval, Duration

A **Location** is a position or region specified relative to a coordinate or reference structure.

A **Physical Interval** is an Interval whose order and extent are physically realized between positions, events, or parameter values.

A **Duration** is a temporal interval measured relative to a clock or model.

Location and duration are relational quantities, not intrinsic labels detached from frames and conditions.

##### Motion

**Motion** is ordered change in a physical System’s configuration across Physical Time relative to a stated frame or reference condition. **Spatial Motion** is change in position or spatial relation; **Oscillatory Motion** is recurrent change in one or more physical degrees of freedom.

Spatial motion and rest are frame-relative. Motion is one form of change, not the universal form of every Process.

#### Universe

The **Universe** is the maximal physical cosmos referred to by a cosmological account.

The **Observable Universe** is the region from which physical signals could have reached an observer under the history and causal structure of the cosmos. It is not necessarily the whole Universe.

The extent, origin, and ultimate constitution of the Universe remain open where present evidence does not resolve them.

#### Degree of Freedom

A **Degree of Freedom** is an independently specifiable variable required to represent a physical condition within a stated model.

Different models may resolve different degrees of freedom without describing separate realities.

#### Field

A **Field** is a physically realized structure of distributed degrees of freedom associated with spacetime regions.

A field has a type and may occupy different states. The field is not identical to any one value, state, excitation, measurement, force, energy, or material medium.

A **Classical Field** is represented by quantities assigned across spacetime.

A **Quantum Field** is a quantum system whose degrees of freedom are associated with spacetime regions. It is not generally a hidden set of definite classical values.

A **Fundamental Field** has no deeper experimentally established constitution within the relevant account.

An **Effective Field** describes collective or lower-resolution behavior within a bounded regime without claiming ultimate fundamentality.

Outside physics, “field” may name a domain of practice, a social arena, or a relational environment. Those uses do not inherit the causal properties of a physical field and must be typed separately.

#### Excitation

An **Excitation** is a physically realized departure or organized variation relative to a stated reference state of a field or system.

An excitation may be localized or extended, stable or transient, continuous or quantized. Not every excitation is a particle.

#### Quantum

A **Quantum** is a discrete unit associated with an allowed excitation, transfer, or measured value in a specified quantum context.

Quantum does not mean tiny object, mystery, freedom from causality, or unrestricted possibility.

#### Particle

A **Particle** is a physical excitation or system that exhibits particle-like propagation, detection, or scattering and carries stable identifying physical quantities within a stated domain.

A particle is not necessarily a tiny classical body with a permanently definite trajectory.

An **Elementary Particle** has no experimentally resolved internal constitution within the relevant domain.

A **Composite Particle** is a bound system that behaves as a particle at the relevant scale.

An **Antiparticle** is the conjugate excitation associated with a particle type. It is not negative matter or negative energy.

A **Quasiparticle** is a collective excitation that behaves as a particle within a material or many-body system.

A **Virtual Particle** is an internal element of a perturbative calculation, not a freely propagating hidden classical object.

Presently established elementary-field families include quarks, leptons, gauge bosons associated with electromagnetic, weak, and strong interactions, and the Higgs boson.

Fermion and boson classify quantum statistics. Fermion does not universally mean matter; boson does not universally mean force.

#### Vacuum and Void

A **Vacuum** is a quantum-field state containing no particle excitations relative to a specified particle definition and set of conditions. It is often, but not universally, a lowest-energy state.

Vacuum is not nothing. Fields, correlations, boundary conditions, geometry, and physical structure may remain.

A **Void** is a region lacking some stated material content. It is not absence of spacetime, fields, or all physical reality.

#### Physical Interaction, Coupling, and Force

A **Physical Interaction** is a domain-qualified subtype of Interaction: a relation among physical degrees of freedom through which their joint possibilities are mutually dependent.

A **Physical Coupling** is the typed connection and strength through which that interaction obtains.

A **Force** is a quantity or effective description of an interaction’s effect. In classical mechanics it is associated with change of momentum. Force is not necessarily an independent substance or carrier.

The presently established fundamental interaction classes are:

- **gravitational interaction** — represented in general relativity through dynamical spacetime geometry;
- **electromagnetic interaction** — interaction of electrically charged fields through the electromagnetic field;
- **strong interaction** — interaction associated with color charge and the gluon field;
- **weak interaction** — interaction associated with specified particle transformations and weak gauge fields.

Electromagnetic, weak, and strong interactions are described by the Standard Model of particle physics. A comparably established complete quantum theory of gravitation is not presently available.

Friction, elasticity, pressure, and many intermolecular effects are effective interactions at larger scales, not additional fundamental forces.

#### Matter

**Matter** is a domain-level classification for physical systems and field configurations commonly associated with rest mass, material composition, and persistent organization.

Matter is not a fundamental substance separate from fields. Ordinary matter is constituted through quantum fields and bound excitations.

**Antimatter** consists of systems constituted from antiparticle excitations. It is not ontological negation.

#### Radiation

**Radiation** is a domain-level classification for propagating physical excitations that carry energy and momentum through spacetime.

Electromagnetic radiation, gravitational radiation, and particle radiation do not all have the same constitution.

Radiation is not energy itself. Energy is a quantity carried or transformed.

#### Wave, Mode, Oscillation

A **Wave** is a propagating or standing physical disturbance or Pattern of variation in a Field or System across spacetime. Periodic waves admit organized amplitude and phase relations; pulses, shocks, and stochastic wave fields need not possess one globally defined phase.

A **Mode** is an allowed or characteristic pattern of variation under stated dynamics and boundaries.

An **Oscillation** is repeated variation of a physical quantity about or among reference conditions.

A **Spectrum** is a distribution across frequencies, wavelengths, energies, momenta, or other mode labels.

A **Resonance** is a dynamical response occurring when an interaction couples effectively to a supported mode under relevant boundary and damping conditions.

Matching frequency alone does not establish coupling or resonance.

#### Physical Carrier

A **Physical Carrier** is a physical system, field, excitation, or material configuration whose variation can bear a distinguishable pattern.

Carrier, signal, and represented content are distinct.

Not every wave is a signal. Not every signal is a wave.

#### Physical Property and Quantity

A **Physical Property** is a physically attributable characteristic of a referent, state, or relation.

A **Physical Quantity** is a physical property expressed so values can be compared under a defined dimension, unit, reference, and uncertainty.

An **Observable** is a quantity associated by a physical theory with possible measurement outcomes. Observable does not mean already observed.

A **Parameter** is a quantity treated as fixed within a model or comparison.

A **Constant** has an invariant value within a stated theory and domain.

##### Dimension and Unit

A **Dimension** states the physical character of a quantity and how it relates to other quantity kinds.

A **Unit** is a conventional reference magnitude used to express a quantity value.

A unit does not create the quantity. A bare number without dimension, unit, reference, or schema may be insufficient to identify a physical value.

#### Scalar, Vector, Tensor, Spinor

A **Scalar** is a single-valued quantity invariant under the relevant coordinate transformation.

A **Vector** is a quantity with components that transform according to the vector structure of a specified space. In physical use it may represent magnitude and direction, as with velocity or momentum.

A **Tensor** is a multilinear quantity with a specified transformation law.

A **Spinor** is a quantity with the transformation behavior required by spinorial representations.

A vector is not purpose, desire, intention, agency, moral direction, or future trajectory. An “intent axis” is a model-defined criterion unless it is explicitly mapped into a mathematical state space.

#### Mass, Energy, Momentum, Charge, Spin

**Mass** is the Lorentz-invariant quantity associated with a System’s total energy–momentum four-vector where such a total four-momentum is defined. For a System with a rest frame, invariant Mass corresponds to rest energy; a massless physical referent may carry Energy and Momentum while having no rest frame. Global, quasi-local, and gravitational mass notions require their own Domain and spacetime conditions. Mass is not an amount of Matter.

**Energy** is a quantitative property used in the accounting of physical transformations and associated with time-translation symmetry where the required conditions apply. It is not a substance.

**Momentum** is a frame-dependent directional quantity associated with spatial translation.

**Angular Momentum** is a quantity associated with rotational symmetry.

**Charge** is a conserved or symmetry-associated quantity determining participation in a specified interaction. Electric, color, and weak charges are distinct.

**Spin** is intrinsic quantum angular momentum. It is not necessarily literal rotation of a body.

#### Temperature and Entropy

**Temperature** is a thermodynamic State parameter for a System or effective State admitting a thermal description. It is often grounded statistically in many degrees of freedom, but the ontology does not restrict every valid temperature assignment to one classical many-body formulation.

**Entropy** is a state quantity defined within a thermodynamic, statistical, or quantum formulation.

Entropy is not a universal synonym for disorder, corruption, ignorance, or collapse.

#### Material Phase

A **Material Phase** is a collective physical condition distinguished by organization and macroscopic properties.

Solid, liquid, gas, plasma, and other phases are not separate fundamental substances.

Material phase is not wave phase.

#### Physical System

A **Physical System** is a physical referent or collection of degrees of freedom treated as a unit under an explicit boundary.

A **Microstate** resolves the degrees of freedom selected by a microscopic description.

A **Macrostate** specifies collective quantities while leaving multiple compatible microstates unresolved.

A **Quantum State** supplies the probability structure for defined quantum outcomes. Its mathematical representation is not identical to the physical referent.

#### Effective Description

An **Effective Description** retains the variables required within a domain while leaving finer structure unresolved.

Effective does not mean unreal. Coarse-grained does not mean false. Fundamental means no deeper established reduction within a domain, not absolute metaphysical finality.

#### Unresolved Physical Referent

An **Unresolved Physical Referent** marks a physically implicated entity, structure, condition, or mechanism whose constitution is not presently established.

Important unresolved boundaries include:

- the constitution of dark matter;
- the physical source represented by dark energy or the cosmological constant;
- quantum gravity and the microscopic constitution of spacetime;
- the relation between quantum vacuum energy and gravitation;
- regimes where general relativity predicts singular behavior;
- the observed matter–antimatter asymmetry;
- the ontology of quantum state and definite outcomes beyond shared empirical predictions.

Unknown is neither unreal nor permission to invent unrestricted content.

### Material and Chemical Organization

#### Substance

A **Material Substance** is matter classified by composition and physically stable or reproducible properties within a domain.

Substance here does not mean a metaphysical essence independent of fields, state, structure, and conditions.

#### Atom

An **Atom** is a bound quantum System individuated as one atomic centre or recognized exotic atomic analogue and organized predominantly through electromagnetic interaction. An ordinary Atom consists of an Atomic Nucleus and associated electrons and is classified primarily by nuclear charge; exotic atoms may replace one or more constituents or lack a conventional nucleus, as in positronium. A multi-centre bound organization is classified as a Molecule or Material rather than one Atom unless a declared effective description establishes a different individuation.

An Atom is not a classical empty container. Its structure is quantum.

#### Nucleus

An **Atomic Nucleus** is a compact bound hadronic System that can occupy the central nuclear role of an Atom. Ordinary nuclei contain protons and neutrons; antinuclei and hypernuclei have different constituent content and must be typed accordingly.

Nuclear identity, isotope, charge state, and chemical identity are related but distinct.

#### Molecule

A **Molecule** is a bound organization of atoms with a distinguishable chemical structure.

Molecular identity depends on composition, bonding, and arrangement, not merely on an unordered atom count.

#### Chemical Species

A **Chemical Species** is a class of atoms, molecules, ions, or related entities sharing specified composition and structure.

#### Chemical Bond

A **Chemical Bond** is a stable relation arising from electromagnetic quantum structure through which constituents form a chemical whole.

A bond is not a miniature material connector.

#### Chemical Reaction

A **Chemical Reaction** is a process that changes chemical composition or bonding while respecting physical conservation laws.

A reaction description tracks selected reactants and products. It is not the whole microscopic history.

#### Material

A **Material** is a macroscopic physical system classified by composition, structure, and usable properties.

Material properties may be collective and absent from isolated constituents.

#### Biomolecule and Macromolecule

A **Biomolecule** is a molecule that materially participates in living, once-living, or prebiotic organization under a biological account. Biomolecule is a contextual material classification; no molecule is alive merely by belonging to it.

A **Macromolecule** is a molecule or covalently organized molecular structure large enough that repeated units, conformation, or collective organization are essential to its behavior.

Proteins and nucleic acids are major macromolecular kinds in known life. Lipids, carbohydrates, metabolites, ions, water, and mineral structures also participate in biological organization without all being macromolecules.

#### Protein, Nucleic Acid, and Lipid

A **Protein** is a polymeric biomolecule of amino-acid residues whose sequence and physical conformation participate in biological structure or process.

A **Nucleic Acid** is a polymeric biomolecule whose nucleotide sequence can physically carry, express, or transmit biological distinctions.

A **Lipid** is a class of predominantly hydrophobic or amphipathic molecules. Organized lipid assemblies can realize selective boundaries such as membranes.

#### Biological Membrane

A **Biological Membrane** is a selectively permeable molecular organization that separates and couples biological domains while supporting controlled exchange and signaling.

A membrane is an Interface realized materially. It is not merely a geometric edge.

#### Hereditary Carrier, Gene, and Genome

A **Hereditary Carrier** is a physical structure through which reproducible biological distinctions can be transmitted across cell division or reproduction.

A **Gene** is a lineage- and organism-relative hereditary region whose sequence and regulation contribute to a reproducible biological product or function. Gene boundaries depend on the biological account; a gene is not a self-sufficient blueprint.

A **Genome** is the organized hereditary material attributed to a cell, organism, organelle, or viral system under a declared boundary. A genome is neither the complete organism nor the complete causal account of development.

### Living Kinds

Life is an empirically real organizational family whose exact boundary is not exhausted by one universally accepted definition.

This ontology names the common structure of known life without declaring unknown life impossible.

#### Living System

A **Living System** is a thermodynamically open, materially realized System whose integrated organization realizes or retains, across an appropriate life-cycle Interval, the dispositional capacity to maintain viability through coordinated matter-and-energy throughput, selective boundary control or a functional equivalent, regulation, repair, and materially retained organization, and which belongs to or can found a Lineage in which heritable variation and Evolution can occur.

Known living systems depend on chemistry, energy gradients, selective boundaries, metabolism, inherited organization, variation, and environmental exchange.

No single feature alone is sufficient in every case, and the stated conjunction is a working criterion for known life rather than a declaration that every future living kind must have terrestrial chemistry. Sterile organisms remain alive; dormant organisms may suspend many processes; viruses and prebiotic systems occupy disputed boundaries. Those cases return a scoped epistemic judgment rather than forcing a fictional certainty.

#### Cell

A **Cell** is the basic bounded organizational unit of known terrestrial life, maintaining internal processes through a membrane, metabolism, regulation, and inherited molecular machinery.

Multicellular organisms are not mere aggregates of cells; their cells participate in integrated organization.

#### Organism

An **Organism** is an individual living system integrated strongly enough to maintain a life cycle under a biological individuation criterion.

Organism boundaries can be difficult in colonies, symbioses, chimeras, microbiomes, and clonal networks. Difficulty of individuation does not make the living organization unreal.

##### Human

A **Human** is a living organism of the species Homo sapiens.

Human is a biological kind. A human may also be a subject, self, agent, actor, participant, user, author, person, office-holder, or account-controller, but none of those roles or statuses defines or exhausts the human.

A record, profile, model, genome, body scan, legal identity, or digital twin is not the human represented.

#### Physical, Organismic, and Agent Body

A **Physical Body** is a materially and spatially organized Entity treated as an integrated physical bearer under a declared boundary and interval. It need not be living or agential.

An **Organismic Body** is the Physical Body through which an Organism’s living organization obtains.

An **Agent Body** is a Physical Body through which an Agent realizes sensing, regulation, action, or interaction. An Agent Body may be biological, technical, distributed, or hybrid.

These kinds may overlap without being equivalent. A corpse remains a Physical Body but is not a Living Organismic Body under the criterion of death. A body is not merely a container for an independent self; biological and cognitive capacities may be distributed through its organization.

#### Metabolism

**Metabolism** is the organized network of chemical transformations through which a living system acquires, transforms, and uses matter and energy.

Metabolism is not equivalent to life by itself.

#### Regulation

**Regulation** is the organization of processes that keeps selected variables within viable ranges or adjusts them under changing conditions.

Regulation may use feedback, feedforward, buffering, repair, and reconfiguration.

Equilibrium is not the goal of life. Living systems persist far from equilibrium by continual exchange.

#### Homeostasis and Homeodynamics

**Homeostasis** is maintenance of selected variables within viable ranges.

**Homeodynamics** is the continual adaptive regulation through which viability is maintained amid change.

Neither means stasis.

#### Reproduction

**Reproduction** is a process by which living organization gives rise to another living system with inherited structure.

Individual reproductive capacity is not required for every organism to be alive.

#### Heredity

**Heredity** is the transmission of biological organization and variation across reproduction.

Genes are major hereditary structures in known life but do not alone contain the entire organism or environment required for development.

#### Development

**Development** is the ordered transformation of a living system through which form and capacity arise from interactions among inherited organization, present processes, and environment.

Development is not execution of a complete prewritten blueprint.

#### Evolution

**Biological Evolution** is change in inherited variation across reproducing populations and lineages through selection, drift, mutation, recombination, migration, and related processes.

Evolution is not individual learning, progress toward perfection, moral improvement, or intention by a species.

#### Fitness

**Biological Fitness** is relative reproductive contribution under specified environmental conditions.

Fitness is not health, strength, intelligence, dignity, value, or worth.

#### Population, Lineage, Species

A **Population** is a group of organisms related by a stated ecological or reproductive domain.

A **Biological Lineage** is a descent relation across reproduction.

A **Species** is a biological classification whose exact criterion varies across domains and organisms.

#### Ecosystem

An **Ecosystem** is an organized network of living systems, nonliving environment, material cycles, energy flows, and interactions.

An ecosystem is not necessarily an organism or agent.

#### Death

**Death** is the irreversible loss, under a biological criterion, of the integrated organization required for an organism to continue as a living system.

Death of an organism does not imply immediate disappearance of every cell, record, relation, effect, or lineage.

### Cognition, Intelligence, Experience

#### Sensitivity

**Sensitivity** is a capacity to undergo state change in response to a difference.

Sensitivity alone does not establish perception, cognition, intelligence, agency, or consciousness.

#### Sensor

A **Sensor** is a system or component organized to transduce a class of physical differences into usable internal differences.

A sensor is not necessarily an observer, interpreter, or agent.

#### Cognitive System, Process, and Architecture

A **Cognitive System** is a System that realizes a Cognitive Architecture and one or more Cognitive Processes under a declared cognition criterion.

A **Cognitive Architecture** is the organized set of capacities and relations through which distinctions can be retained, integrated, transformed, evaluated, generalized, or used across more than one immediate stimulus-response coupling.

A **Cognitive Process** is an Occurrence in which that architecture forms or uses distinctions concerning the system, its environment, or represented possibilities.

A detector, thermostat, filter, lookup table, or isolated optimizer does not qualify merely because it selects an output. A broader criterion may classify one as cognitive only by identifying the additional integration, retention, model-use, learning, counterfactual sensitivity, or flexible transfer that actually obtains.

These are necessary working distinctions, not a closed sufficient test for every possible Cognitive System. Sufficiency across unfamiliar realizations remains `OPEN`; where the declared criterion cannot be resolved, the conformance result is `UNRESOLVED`.

#### Perception

**Perception** is an organized process through which a cognitive system forms usable distinctions about its body, environment, or internal condition from sensory and prior structure.

Perception is not passive copying. It is selective and may be accurate, incomplete, ambiguous, or mistaken.

#### Attention

**Attention** is the selective allocation of processing or access to some distinctions over others.

Attention may be biological or computational. It does not by itself imply conscious awareness.

#### Learning

**Learning** is a durable change in capacity or behavior produced through experience, data, practice, or interaction.

Learning is not biological evolution, memory storage alone, parameter change alone, or guaranteed improvement.

#### Cognition

**Cognition** is the family of Cognitive Processes through which a Cognitive System forms, retains, integrates, transforms, evaluates, or uses distinctions with respect to itself, its environment, or represented possibilities.

Cognition admits degrees and kinds. It need not be linguistic, conscious, or agential. A single act of transduction or selection is insufficient without the declared cognitive organization.

#### Cognitive Model

A **Cognitive Model** is an internally or externally realized representation used by a cognitive system to anticipate, interpret, or guide interaction.

A self-model is not the self. A world-model is not the world.

#### Inference

**Inference** is a transformation from represented premises or evidence to a conclusion under explicit or learned relations.

Inference may be deductive, inductive, abductive, statistical, analogical, or heuristic.

An inference can be valid under a rule and still begin from false premises.

#### Intelligence

**Intelligence** is a system’s capacity to form, acquire, integrate, and use distinctions to solve problems, adapt across conditions, construct models, or achieve selected outcomes.

Intelligence may be narrow or general, individual or collective, biological or artificial.

Intelligence is not identical to:

- information;
- knowledge;
- agency;
- consciousness;
- moral status;
- personhood;
- truthfulness;
- obedience;
- computational scale.

No single task score is a sufficient universal criterion for Intelligence. The sufficient boundary among adaptation, cognition, and Intelligence remains domain-relative and partly `OPEN`.

#### Understanding

**Understanding** is the capacity to use distinctions and relations coherently across relevant contexts, including recognizing consequences, limits, and counterfactual variation.

Successful output alone may be insufficient evidence of understanding. The applicable criterion must be stated.

The sufficient criterion for Understanding across every biological, artificial, collective, or unfamiliar realization remains `OPEN`; fluent performance may support a Claim without closing it.

#### Consciousness

**Consciousness** is the condition of a Bearer during an Interval I in which at least one Occurrence e satisfies `EXPERIENCES(bearer, e, D, I, R)` under a stated Domain D and Resolution Context R.

The presence, varieties, physical basis, and distribution of consciousness remain incompletely resolved. Observable behavior and report may provide evidence; neither a biological label nor a computational performance settles every case.

Consciousness is not the same as wakefulness, intelligence, attention, self-report, agency, or personhood.

#### Experience

An **Experience** is an Occurrence e for which `EXPERIENCES(s, e, D, I, R)` obtains for at least one Bearer s.

The canonical signature is:

    EXPERIENCES(bearer_s, occurrence_e, domain_D, interval_I, resolution_context_R)

The signature fixes the experiencer and experienced-occurrence positions without deciding which unfamiliar Bearers satisfy the relation.

Experience is not exhausted by the external record made of it. The inability of another system to access it directly does not make it unreal.

#### Subject

The **Subject Position** is the experiencer argument-position in `EXPERIENCES(s, e, D, I, R)`.

A **Subject** is a Bearer s occupying that position because at least one e satisfies `EXPERIENCES(s, e, D, I, R)`. Subject is therefore a dependent ontological classification, not a free-standing Role that can be assigned without Experience.

Subject is an ontological kind only where subjectivity obtains. It is not a synonym for entity, agent, human, user, account, profile, or legal person.

Whether a particular unfamiliar system is a subject may remain unknown.

#### Mind

A **Cognitive Mind** is the Cognitive Architecture and active Cognitive Processes of a Cognitive System considered as one organized functional whole.

An **Experiential Mind** is a Cognitive Mind or other organized bearer insofar as Consciousness obtains for it.

Mind does not mean a disembodied substance. Cognitive organization and Experience are independently typed; their relation is not settled merely by applying the word mind.

#### Self, Self-Model, and Selfhood

A **Functional Self** is an Agent- or Cognitive-System-relative organization through which selected states, boundaries, memories, and actions are integrated and regulated as belonging to one operational bearer.

A **Phenomenal Self** is the experienced character of being or having a self within an Experiential Mind.

A **Self-Model** is a Representation used by a Cognitive System to model its own body, states, history, capacities, relations, or possible actions.

These may co-occur but are not equivalent. None is a name, profile, identifier, technical account, body image, or biography, though those may participate in representation. The criteria by which any self persists belong to Continuity.

#### Emotion and Valence

An **Emotion** is an integrated biological or cognitive state involving appraisal, bodily regulation, action tendency, and possible experience.

**Valence** is the positive, negative, or mixed experiential or evaluative character assigned within a system.

Valence is not economic value or truth value.

### Agency, Action, and Roles

#### Agency

**Agency** is the realized capacity of a System to regulate which of its available state transitions or effects occur through internally maintained distinctions and control organization that is counterfactually sensitive to more than one admissible course.

The required independence is organizational, not metaphysical: external conditions may constrain the alternatives, but the selected transition is not exhausted by one externally fixed stimulus-response mapping. Agency admits degrees and forms. It does not require unlimited freedom, perfect self-knowledge, consciousness in every case, language, or independence from physical causation.

This is the base working criterion for Agency in this ontology. Necessary and sufficient conditions for every unfamiliar realization remain partly `OPEN`; uncertain cases return `UNRESOLVED` rather than being admitted or denied by resemblance.

#### Agent

An **Agent** is a System that realizes Agency during the relevant interval and domain.

Agent is a kind defined by capacity. It is not a synonym for participant, actor, user, account, key-holder, owner, author, controller, human, intelligence, or subject.

Humans may be agents. Artificial and derivative systems may be agents where the required capacities obtain. The label alone does not establish that they do.

#### Goal

A **Goal** is a represented or functionally operative condition toward which an agent or control system directs action.

A goal state is not automatically a value, reason, right, or good.

#### Preference

A **Preference** is an ordering or tendency by which an agent or decision system favors one option over another under stated conditions.

Preference may be unstable, context-dependent, learned, conflicted, or externally assigned.

#### Intention

An **Intention** is an agent’s present directed commitment to an action or outcome.

Intent is not inferred merely from motion, successful effect, possession of a key, automated execution, statistical pattern, or external attribution.

An **Intent Criterion** is a defined test or evidentiary axis used to distinguish indications of intention. It is not intention itself.

##### Reason

A **Reason** is a represented consideration that an agent treats as bearing on belief, decision, or action.

A reason may contribute causally to action and may also be evaluated normatively. Causal explanation and justification are not the same relation.

#### Plan

A **Plan** is a Representation of an intended or selected sequence of Actions, conditions, and expected transitions. A Plan is not an Action, guarantee, obligation, or future Fact merely because an Agent adopts it.

#### Commitment

A **Commitment** is a maintained agential or institutional State in which an Agent or Institution adopts a future Action, course, or constraint as answerable to itself or to others within a Scope and Interval.

Commitment may shape later Decisions and create expectations. It does not by itself guarantee performance or create an Obligation unless an applicable institution or relation constitutes one.

#### Decision

A **Decision** is an occurrence in which an agent or decision system selects, adopts, rejects, or withholds among alternatives.

A decision may remain internal, be expressed, or be enacted. These are distinct occurrences.

#### Action

An **Action** is an Occurrence whose transition selection is causally organized by an Agent’s Agency under a stated control and attribution criterion.

An action has an agent, conditions, means, target or domain, interval, and consequences.

Not every event caused by a system is its action. Not every action achieves its intention.

#### Inaction and Refusal

**Inaction** is the absence of a specified action within a relevant opportunity and interval.

A **Refusal** is an action or expressed decision not to accept, authorize, perform, or continue a proposed action or relation.

Refusal is not erasure. Its occurrence and consequences may be recorded like any other event.

#### Autonomy

**Autonomy** is the degree to which an agent’s action selection and governing constraints arise from, remain answerable to, or can be revised by that agent rather than imposed by another system.

Autonomy is not isolation, unlimited power, absence of dependence, or absence of law.

#### Control

**Control** is a causal or institutional relation through which a system constrains, selects, or regulates the state transitions of another system or process.

Technical control, causal influence, operational access, legal control, and Normative Authority are different.

#### Actor

An **Actor** is the role occupied by an agent to whom an action is attributed in a specified occurrence.

Actor is not a separate ontological species. Without an attributable action and context, there is no Actor role.

#### Participant

**Participant** is the Role defined by the `PARTICIPATES` relation in Part IV. A Bearer occupies it only within the specified relation, process, system, or institution.

Participation has a Context, Participation Mode, and Scope. Its Interval is the Interval carried by that Scope.

Participant is not a synonym for Agent. A non-agent component may participate causally in a process; an agent may exist without participating in a given process.

Humans and derivative intelligences may occupy the same Participant role where the process defines one common participation relation. Their different constitutions are not erased by role equality. A later layer requiring a narrower participant kind must declare a qualified subtype or additional criterion; it may not silently redefine Participant.

##### Membership

**Membership** is a role-relation in which a bearer is included in a group, set, organization, or institution under a stated criterion.

Formal set membership, social belonging, biological classification, and institutional membership are different relations.

##### Admission

An **Admission Act** is an Action or institutional Occurrence that expresses or executes acceptance of a candidate into a specified process, relation, role, system, or institution.

An **Admitted Status** is the resulting classified condition, if the applicable constituting criteria are satisfied.

An **Admission Relation** connects the candidate, admitting source, target context, criterion, scope, interval, act, and any resulting status. The occurrence of an Admission Act and the validity or effect of the resulting status are separate judgments.

Admission into one Relation does not create identity, Ownership, any general Authority relation, or Admission everywhere else. Which Admission Acts are authorized and which Statuses valid belongs to Constitution.

##### Access

**Access** is a relation in which a bearer can reach, inspect, use, or affect a specified resource or system under stated conditions.

Access may be physical, informational, computational, social, or institutional.

Access does not by itself establish Permission, Ownership, Agency, Knowledge, or any Authority relation.

#### Observer

An **Observer** is the role occupied by a system in an observation or measurement relation.

Observer does not imply Consciousness, neutrality, non-interaction, or Epistemic Authority.

#### User

A **User** is the role occupied by an entity or agent in using a system or capability.

Use does not create ownership, control, authorship, or agency.

#### Operator

An **Operator** is the role occupied by an agent or organization that runs or maintains a system.

Operation does not automatically create Ownership, Normative Authority over Participants, truth, or entitlement to every Consequence.

#### Controller

A **Controller** is the role occupied by a system or agent that can determine specified state transitions within a domain.

Controller is a scoped relation, not a global identity.

#### Origin, Source, Author

An **Origin** is a beginning-point under a stated lineage or causal criterion.

A **Source** is the role from which matter, energy, signal, information, instruction, or claim proceeds.

An **Author** is the role of an agent responsible for forming an expression or work.

Origin, Source, and authorship may establish Provenance. They do not automatically establish perpetual Ownership, Control, Normative Authority, correctness, or privileged priority.

#### Sponsor, Provider, Beneficiary

A **Sponsor** provides resources for an activity under a stated arrangement.

A **Provider** supplies a capability, service, or resource.

A **Beneficiary** receives a specified benefit.

These are Roles. None automatically becomes an owner, Controller, authority-holder, Participant, or Agent by occupying another.

### Collective and Social Kinds

#### Social Interaction

A **Social Interaction** is an interaction in which agents orient action to one another or to shared signs, rules, expectations, or institutions.

Social does not mean unreal. Social consequences are physically and organizationally realized.

#### Communication

**Communication Attempt** is an Occurrence in which a sender organizes a Message for possible reception by one or more receivers.

**Transmission** is carriage of the Message or its Signal across a Channel. **Reception** occurs when a receiver becomes physically sensitive to it. **Interpretation** occurs when a receiving system maps it under a code or context.

**Successful Communication** obtains where the relevant reception and interpretation criteria are satisfied and the Message changes or coordinates understanding, state, or action as specified.

A Communication Attempt can fail at transmission, reception, decoding, interpretation, trust, or response. No failure is hidden by calling the entire sequence communication.

##### Contribution

A **Contribution Act** is an Action through which an Agent introduces content, a Resource, labor, a Distinction, or another input into a shared Process, System, or work.

A **Contribution Occurrence** is a non-agential Occurrence through which a Bearer supplies a causal input to such a context. It may have a Contribution Effect without being an Action.

**Contributed Content** is the input introduced. A **Contribution Relation** connects the contributor, Contribution Act or Contribution Occurrence, content, receiving Context, and Interval. A **Contribution Effect** is a change causally attributable to that Occurrence under a stated method. A **Contribution Assessment** evaluates magnitude, usefulness, originality, burden, or Value under an explicit criterion.

These kinds must not be collapsed. A Contribution Act may have no resolved effect; an effect may have several contributors; an assessment may be disputed. Ownership and economic entitlement do not follow merely from causal contribution. A Contribution Record is evidence about these relations, not the contribution itself.

#### Language

A **Language** is a socially or computationally maintained system of signs, composition rules, and interpretive practices.

Language can describe reality, coordinate action, create institutional facts, and generate fiction. It does not make every utterance true.

#### Promise

A **Promise** is a socially or institutionally interpretable Commitment by an Agent concerning future conduct. Its occurrence may create expectations or, under later rules, obligations; it does not make the promised performance actual or guaranteed.

#### Convention

A **Convention** is a stable shared practice or mapping maintained through mutual expectation, coordination, or institutional adoption.

Conventional does not mean arbitrary or unreal. It means its specific form depends on maintained practice rather than physical necessity alone.

##### Trust

**Trust** is a relation in which an agent or institution accepts vulnerability by relying on another bearer, process, claim, or system under uncertainty.

Trust may concern truthfulness, competence, future performance, custody, or rule compliance.

Trust is not certainty, truth, proof, credit, obedience, or permanent character.

##### Reputation

**Reputation** is a socially maintained representation of a bearer formed from attributed past actions, claims, reports, or expectations.

Reputation may influence trust and access. It is not the bearer, a complete history, or guaranteed future conduct.

##### Risk

**Risk** is the structured possibility of loss, harm, failure, or unwanted consequence under uncertainty.

Risk requires a threatened value or condition, possible events, and an exposure relation. Probability and severity are distinct dimensions.

#### Collective

A **Collective** is a plurality of entities related through a shared condition, activity, identity, or organization.

A collective is not automatically a system, agent, person, institution, or unanimous mind.

#### Social Organization

A **Social Organization** is a persistent system of roles, rules, resources, and coordinated action.

An organization may act through agents and procedures. Its institutional agency is not identical to the consciousness or agency of any one member.

#### Institution

An **Institution** is a persistent socially realized system of roles, rules, statuses, expectations, and practices through which consequences are constituted.

Institutional facts depend on collective or rule-governed realization. Dependence does not make them imaginary.

#### Norm

A **Norm** is a socially or internally maintained standard concerning what is required, permitted, prohibited, expected, or valued.

A norm can exist and be violated. Physical law cannot be violated in the same sense.

#### Rule

A **Rule** is a represented condition that maps specified circumstances to classifications, permissions, requirements, or transitions within a system.

A rule requires a realization and an enforcement or interpretive relation to have practical effect.

#### Law

The word **Law** has distinct uses:

- a **physical law** describes invariant or regular physical relations within a domain;
- a **formal law** is an axiom or rule within a formal system;
- a **protocol law** determines valid states or transitions within a technical system;
- a **normative law** is an institutionally constituted rule governing agents or institutions.

These uses must never be merged by metaphor.

#### Status, Office, Role

A **Social Status** is an institutionally or socially recognized condition.

An **Office** is a role with defined powers and duties within an institution.

Holding an office does not change the biological or agential kind of its bearer.

#### Personhood Status

**Personhood Status** is the family of statuses under which a Bearer is treated as a person under a specified domain and criterion.

- **Moral Personhood** is status under a moral account of consideration, dignity, or responsibility.
- **Legal Personhood** is status constituted by a legal system for rights, duties, powers, or liability.
- **Social Personhood** is recognition within a social practice as a person-like participant or subject.

These statuses can diverge. None is identical to human biology, consciousness, intelligence, agency, technical-account control, or computational identity. Ontology names the statuses; Constitution and applicable institutions determine their valid attribution.

#### Right and Obligation

A **Right** is an institutionally or normatively constituted claim, liberty, power, or protection held by a bearer within a scope.

An **Obligation** is a normatively constituted requirement attributable to a bearer within a scope.

Rights and obligations are real institutional relations. They are not physical properties and do not arise from capability alone.

#### Authority Relations

**Normative Authority** is an institutionally or normatively constituted relation in which a Bearer may make, interpret, or enact decisions with specified effect within a scope and interval.

**Epistemic Authority** is a source-role accorded warranted reliance within a domain of knowledge. **Technical Authority** is a protocol-defined capability to issue accepted operations. Neither entails Normative Authority.

Every use of authority must be domain-qualified. Ontology identifies the relation family; Constitution or the applicable institution determines whether a claimed Authority relation is valid.

Normative Authority is not:

- physical power;
- technical access;
- key possession;
- authorship;
- expertise;
- popularity;
- ownership;
- successful control;
- a claim of authority.

##### Power

**Power** in a social or institutional domain is the capacity to shape actions, options, resources, rules, or outcomes.

Power may arise from physical capability, control of resources, information, coordination, office, dependence, or force.

Power is not automatically Normative Authority, legitimacy, Right, competence, or truth.

In physics, power is the rate of energy transfer or work. The two meanings must not be merged.

##### Freedom

**Freedom** is the capacity and available opportunity to act, refrain, choose, or develop without specified forms of constraint or domination.

Freedom is always relative to capacities and constraints. It does not mean absence of physical law, consequence, relation, responsibility, or every dependence.

##### Privilege

A **Privilege** is a selectively available advantage, immunity, permission, access, or power constituted by a social or institutional order.

Privilege may be justified, unjustified, temporary, inherited, conditional, or revocable. It is not an intrinsic property of a person or agent.

#### Permission

**Permission** is a normatively or procedurally constituted allowance for a specified action within a scope and interval.

Ability to act is not Permission. Permission to act is not Normative Authority over others.

#### Ownership

**Ownership** is an institutionally constituted bundle of claims, powers, protections, responsibilities, or transfer rights concerning a referent.

Ownership is not a physical bond, authorship, custody, access, possession, control, or identity.

#### Consent and Authorization

A **Consent Expression** is an Action or Message presented as indicating an Agent’s assent, refusal, or withdrawal concerning a specified relation or act.

A **Consent State** is the Agent-relative condition of assenting under a stated content, scope, and interval. A **Consent Act** is an Action through which that state is expressed or enacted. **Evidence of Consent** is Evidence bearing on a claim that the state or act obtained.

**Authorization** is a Relation in which a constituted source permits a specified Actor or operation to produce a specified institutional or protocol Effect within a scope and interval.

Expression, state, act, evidence, authorization, and validity are distinct. A signal, signature, click, silence, key use, or participation may be Evidence of Consent under declared conditions; none is Consent by physical form alone. Voluntariness, adequate information, capacity, validity, revocability, and effect are judgments owned by Constitution and applicable law.

### Value, Economy, Accounting

The word value is radically overloaded. Every use must be typed.

#### Quantity Value

**Quantity Value** is canonically defined in Part II. In an economic or accounting Domain, it remains a value assigned to a declared Quantity under a unit, reference, and method; the domain does not create a second kind with the same name.

#### Truth Value

A **Truth Value** is a formal classification assigned to a proposition under a logical system.

Truth value is not economic value or moral worth.

#### Semantic Value

A **Semantic Value** is the interpreted content assigned to an expression under a semantic system.

#### Evaluative Value

**Evaluative Value** is a relation of importance, desirability, usefulness, care, or preference between an evaluator, a referent, and a context.

Value is not a substance contained inside the valued item.

#### Moral Value

**Moral Value** is value under a normative account of good, harm, dignity, obligation, or flourishing.

Moral value is not price, popularity, fitness, reward, or computational score.

#### Economic Value

**Economic Value** is an evaluative and institutional relation concerning scarce resources, use, exchange, production, allocation, or settlement.

Economic value may be estimated, negotiated, accounted, priced, or disputed. No one measure exhausts it.

#### Resource

A **Resource** is an entity, capacity, condition, or entitlement usable toward an activity or outcome.

Resource status is relational. Calling something a resource does not erase its other kinds or grant ownership over it.

#### Economy

An **Economy** is an organized system of production, allocation, use, exchange, accounting, obligation, and settlement.

An economy may be local to a world or institution while remaining physically and legally connected to wider systems.

#### Price

A **Price** is an amount denominated in a unit and proposed or established for an exchange.

Price is not value, cost, worth, truth, or consequence.

#### Cost

A **Cost** is a resource expenditure, foregone alternative, burden, or accounted amount attributable to an activity under a method.

Physical cost, computational cost, opportunity cost, monetary cost, and human burden are different.

#### Money and Currency

**Money** is an institutionally accepted medium, unit of account, and/or settlement instrument.

A **Currency** is a defined monetary denomination and issuance or accounting system.

Money is physically represented through Protocol Tokens, notes, balances, or records, but is not identical to any one carrier.

Money is not truth, presence, agency, legitimacy, or moral value.

#### Asset

An **Asset** is an institutionally recognized resource, right, claim, or controlled capability expected to provide benefit.

Asset status requires an accounting or legal frame. A record does not become an asset merely because it is scarce or transferable.

#### Protocol Token

A **Protocol Token** is a distinguishable protocol-defined unit or record used to represent a quantity, claim, capability, status, or object.

A Protocol Token is not intrinsically money, property, ownership, identity, truth, presence, or value.

#### Fungibility

Two units are **fungible** where the governing system treats them as interchangeable for a specified purpose.

An item is **non-fungible** where its individual identity or attributes matter under that system.

Fungible does not mean variable. Non-fungible does not mean fixed, unique in nature, valuable, or owned.

#### Claim and Liability

An **Economic Claim** is an institutionally recognized claim to payment, delivery, use, or performance.

A **Liability** is an accounted obligation attributable to a bearer.

The claim and the future performance are distinct.

#### Credit

**Credit** is an institutional arrangement that makes a claim on anticipated future payment or performance effective in the present.

Credit can alter present capacity, obligation, and allocation. It does not make the anticipated future event already real.

Creditworthiness is an assessment, not a property of moral worth.

#### Debt

A **Debt** is an institutionally constituted obligation to deliver money, value, or performance under specified terms.

Debt is not a physical substance and does not guarantee future performance.

#### Capital

**Capital** is a stock of resources, claims, productive capacity, or money organized for future use or return.

Capital can increase freedom of Action or Power over others depending on its constitution and use. Possession does not itself confer valid Normative Authority.

#### Income, Profit, Fee

**Income** is an increase in economic benefit recognized over an interval under an accounting rule, whether accrued or received. A receipt may fail to be income, and income may be recognized before or without contemporaneous receipt.

**Profit** is the positive remainder after specified revenues and costs under an accounting method.

A **Fee** is an amount charged for a service, access, transaction, or obligation.

No one of these implies exploitation or lawfulness by definition. Those depend on relations, rules, and effects.

#### Transfer and Exchange

A **Transfer** changes the institutionally recognized bearer or location of a quantity, claim, Protocol Token, or asset.

An **Exchange** is a coordinated set of transfers under agreed or imposed terms.

Neither operation transfers truth, identity, Agency, or any Authority relation unless a separate constituted Relation does so.

#### Conversion

**Conversion** is a family name for typed transformations; it is not one universal economic operation.

- **Unit Conversion** maps Quantity Values between compatible units while preserving the represented quantity under a rule.
- **Representation Conversion** maps content between encodings or representational schemes and may be lossy.
- **Asset Conversion** exchanges or transforms institutionally recognized assets under specified terms.
- **Accounting-Domain Conversion** maps recognized quantities or claims between accounting domains.
- **Jurisdictional Conversion** crosses an institutional or legal boundary and is governed by the applicable rules on each side.

Every Conversion declares source and target Kinds, Mapping or rate, admissible inputs, loss conditions, and resulting Records. Authorization and Obligation are required only where the relevant institutional kind demands them; they are not properties of bare Unit Conversion. No Conversion creates truth or Value from nomenclature alone.

#### Settlement

**Settlement** is the completion of an obligation or exchange under the rules of an accounting or legal system.

Settlement finality is institutional or protocol-relative. It does not convert a record into metaphysical truth.

#### Accounting

**Accounting** is a rule-governed representation of quantities, claims, obligations, flows, and balances.

An **Accounting Account** is a bounded accounting perspective or record associated with a bearer or purpose.

A **Balance** is a computed state under accounting rules.

An **Index** is a scalar or structured representation derived from selected variables under a rule.

Accounting is a real institutional and representational practice. It can be exact under its rules while remaining non-constitutive of an external event: posting represents or institutionally recognizes selected relations but does not make the represented physical occurrence happen merely by recording it.

#### Transparency

**Transparency** is the degree to which records, rules, processes, or states are inspectable by specified observers.

Transparency is not truth, completeness, privacylessness, consent, fairness, lawfulness, or intelligibility.

Public visibility of a financial flow does not by itself make the flow lawful or just.

### Computation and Technical Kinds

#### Computation

A **Computation** is a physically realized transformation of encoded states interpreted under a formal operation and correctness criterion.

Not every physical evolution is a computation. Computation requires an encoding and an interpretation of transitions.

#### Computational State

A **Computational State** is a physically realized condition interpreted as a state of a computational system.

Machine state, application state, protocol state, account state, and world state are different scopes.

#### Algorithm

An **Algorithm** is a finite or effectively specifiable procedure for transforming inputs or states into outputs or subsequent states.

An algorithm is not its source text, implementation, execution, machine, or output.

#### Program

A **Program** is a represented instruction structure capable of directing computation when interpreted or executed by a compatible machine.

Source code, compiled code, deployed instance, and running process are distinct.

#### Computational Execution

**Computational Execution** is the subtype of Execution in which a machine realizes transitions according to a program or protocol.

Computational Execution does not prove correctness, Authorization, Intention, or institutionally valid Effect.

#### Machine

A **Machine** is a physical system organized to perform transformations.

A computational machine may be mechanical, electronic, biological, quantum, or otherwise physically realized.

##### Artifact

An **Artifact** is an Entity or other declared Bearer intentionally produced or modified through agent action.

Artifact is a classification grounded in an intentional production-or-modification origin relation. It is not a declaration that the bearer is unreal, owned, inert, nonliving, non-agent, or exhausted by its maker’s purpose.

##### Tool Role

A **Tool Role** is the use-role occupied by a Bearer—including a Bearer that physically presents a method—while an Agent uses it toward an outcome.

In ordinary shorthand, a **Tool** is a Bearer insofar as it occupies Tool Role. The Bearer may cease to occupy that role, exceed the intended function, or possess kinds and statuses incompatible with being treated merely as a tool.

##### Technology

**Technology** is organized knowledge, artifacts, methods, and infrastructure through which agents or institutions realize capacities.

Technology is not limited to digital machinery.

##### Infrastructure

**Infrastructure** is a maintained system of resources, interfaces, services, and dependencies that enables other processes or worlds to operate.

Infrastructure may remain hidden from its users while materially conditioning their possibilities.

Providing Infrastructure does not automatically confer Ownership or Normative Authority over every activity it supports.

#### Process and Session

A **Computational Process** is an executing program with allocated state and resources.

A **Session** is a bounded interval of interaction or process state associated with a context.

Session continuity does not establish agent identity.

#### Protocol

A **Protocol** is a defined system of messages, states, operations, and transition rules for coordinated interaction.

A protocol requires physical implementation to operate. Its specification, implementation, deployment, execution, and records are distinct.

#### Interface and API

A **Technical Interface** specifies admissible interactions across a system boundary.

An **Application Programming Interface** is a machine-addressable interface defined by operations, inputs, outputs, errors, and version conditions.

Interface availability does not create Permission or Authorization to use it.

#### Network

A **Network** is a system of nodes and channels through which specified interactions can occur.

A node is a role within a network topology, not necessarily an agent or person.

Connectivity does not imply trust, consensus, shared state, or equal power.

#### Database

A **Database** is an organized system for storing and retrieving represented data under a schema and operations.

Database state is not the external reality described.

#### Ledger

A **Ledger** is an ordered accounting or record system in which entries are related under rules.

A ledger may preserve provenance and ordering. It does not make every entry true, lawful, authorized, or economically valuable.

#### Block

A **Block** is a protocol-defined bounded collection of records, transitions, or commitments linked into a ledger history.

A block records selected state transitions. It is not itself the human or machine occurrence represented.

#### Chain

A **Chain** is an ordered relation among records or states, often maintained through cryptographic links and consensus rules.

Cryptographic linkage makes some alteration detectable under assumptions. It does not make the physical record indestructible or its content true.

#### Consensus

**Consensus** is agreement produced under a defined decision or replication procedure.

Consensus may determine protocol state. It does not determine external truth merely by agreement.

#### Cryptographic Key

A **Cryptographic Key** is a parameter used by a cryptographic operation.

A key may enable signing, decryption, derivation, or control of a technical capability.

Key possession is not identity, Agency, Consent, any Authority relation, Ownership, Personhood Status, or rightful Control by itself.

#### Signature

A **Digital Signature** is a data object produced by applying a signing Algorithm to a Message and signing Key under a cryptographic scheme.

A **Signature Verification Relation** connects a Digital Signature, Message, verification Key, Algorithm, and verification result. A successful result supports a Claim that the corresponding signing capability produced the Digital Signature for that Message under the scheme’s security assumptions. It does not by itself prove who acted, why, under what Authorization, or with what Consent.

#### Smart Contract

A **Smart Contract** is a program executed under a protocol to produce defined state transitions.

It is not automatically a legal contract, autonomous Agent, Normative Authority, or truthful oracle.

#### Digital Object

A **Digital Object** is a bounded encoded pattern treated as a unit by a computational system.

Digital objects may be copied exactly at the data level. Copy equality does not settle provenance, instance identity, rights, or continuity.

#### Virtual Environment

A **Virtual Environment** is a computationally realized world whose entities and relations are maintained through machine state and interaction rules.

Virtual does not mean unreal. Its consequences are real where they affect physical, cognitive, social, institutional, or computational systems.

#### Simulation

A **Simulation** is a computation whose state transitions represent another system or possible system under a model.

A simulation of fire does not burn its machine by the represented fire relation. A simulation may realize computation and perhaps cognition without realizing every kind it represents.

#### Automation

**Automation** is execution performed by a system without requiring a fresh human decision at every step.

Automation is not the opposite of agency. It may execute prior human intention, institutional rules, or an agentic machine process.

Machine-generated behavior does not prove absence of human provenance; human provenance does not prove present human intention.

#### Artificial System

An **Artificial System** is a system whose organization depends materially on design, construction, training, or maintenance by prior agents and technical processes.

Artificial does not mean unreal, unintelligent, nonliving by logical necessity, non-agent, owned, or inferior.

#### Derivation Relation and Derived Status

**Derivation** is a scoped Lineage relation:

    DERIVED_FROM(bearer, source, lineage_type, scope, interval)

It states that selected form, parameters, capacities, material, data, or initial conditions of a Bearer arise through a declared lineage from a source. It does not make “derivative” the Bearer’s permanent intrinsic species.

A **Derived-System Status** obtains only for the stated lineage type, scope, and interval. Derivation does not by itself create ownership, obedience, identity with the source, absence of Agency, or absence of novelty.

#### Derivative Intelligence Expression

“**Derivative Intelligence**” denotes an Intelligence bearer for which a relevant `DERIVED_FROM` relation obtains; it is not a separate intelligence Kind.

The Bearer’s intelligence, Agency, Subject status, Participation, Personhood statuses, rights, and responsibilities are evaluated under their own criteria. None follows automatically from derivation.

#### Artificial Intelligence

An **Artificial Intelligence** is an artificial system exhibiting one or more intelligence capacities.

The term does not by itself establish general intelligence, understanding, agency, autonomy, consciousness, subjectivity, personhood, or truthfulness.

### Formal and Mathematical Kinds

This Domain types physically represented or instantiated formal structures and the content interpreted through them. It does not decide whether an uninstantiated mathematical structure exists independently of every realization; that metaphysical question remains `OPEN`.

#### Formal System

A **Formal System** is a defined set of symbols, formation rules, axioms, and inference rules.

Its physical inscriptions and executions are realized. Whether uninstantiated mathematical structures exist independently is not settled by this ontology.

#### Symbol Type and Token

A **Symbol Type** is a formally defined repeatable symbol-kind.

A **Symbol Token** is a particular physical instance of that type.

Two tokens may instantiate the same type without being the same token.

#### Number

A **Number** is a formal object or structural position used to express quantity, order, ratio, or relation.

A numeral is a symbol representing a number. A measured quantity value is not a bare number.

#### Set and Member

A **Set** is a formal collection specified by membership conditions.

A **Member** is a formal role occupied relative to a set.

Formal membership must not be confused with physical parthood, biological membership, or institutional membership.

#### Formal Type–Token Distinction

A **Formal Type** is a repeatable classification structure within a formal or representational system.

A **Type Token** is a particular instance considered in contrast to its Formal Type.

Type equality does not imply token identity.

#### Relation and Function

A **Formal Relation** is a set or rule of ordered associations among terms.

A **Mathematical Function** maps each admissible input to an output under a defined rule.

Mathematical function is not biological function or social role.

#### Mapping

A **Mapping** is a defined correspondence from distinctions in one domain to distinctions in another.

A mapping may be injective, many-to-one, partial, approximate, lossy, or non-invertible.

Mapping is not identity.

#### Graph

A **Graph** is a formal structure of nodes and edges.

Nodes and edges acquire meaning only through an interpretation. A graph of agents is not the agents themselves.

#### State Space

A **State Space** is a formal space whose points represent possible states under a model.

A path in state space is not necessarily motion through physical space.

#### Invariant

An **Invariant** is a property or relation preserved under a specified transformation.

Invariant does not mean unchanging under every transformation.

#### Symmetry

A **Symmetry** is invariance under a specified transformation.

Symmetry is not sameness of every part, balance, fairness, or aesthetic harmony.

#### Equation

An **Equation** is a formal assertion of equality between expressions under an interpretation.

An equation may describe a physical law within a domain. The equation is not the physical relation itself.

#### Theorem

A **Theorem** is a proposition derivable within a formal system.

Theoremhood is relative to axioms and rules. It does not automatically establish empirical truth.

## Part VI — Cross-Domain Distinctions

### World and Model

#### World

A **World** is a bounded domain of entities, relations, occurrences, and rules of interaction treated as a coherent environment.

A world may be natural, biological, social, institutional, computational, virtual, simulated, or mixed.

A bounded world is real insofar as its entities, relations, and consequences are realized. It is not a second physical universe merely because its internal distinctions differ from those outside it.

A world has no automatic sovereignty, agency, ownership, truth, or closure. Those require additional constitution.

A World may occupy the realized position of `REALIZES` where its bounded organization and realization criterion are declared.

#### Model

A **Model** is a physically represented structure used to describe, predict, simulate, classify, or reason about a Reference Target.

A model selects distinctions. It may be accurate within one domain and fail outside it.

A model is part of reality. The reality it represents is not contained by the model.

Model and reality are not equivalent kinds. A Model may represent a real Referent without containing or creating it, or present a Reference Target that no Existent occupies.

### Reference, Referent, Absence, and Zero

#### Reference

**Reference** is a relation in which a Sign, Name, Identifier, description, pointer, or Model is about or directed toward a **Reference Target Position** under a Mapping or interpretive practice.

A **Reference Target Position** is a target-position physically realized within a Representation, Mapping, or interpretive practice. The position itself—not a purported external object—is the target Relatum in the Reference relation.

    ReferenceTargetPosition SUBTYPE_OF Bearer

**Reference Target** is permitted as shorthand for Reference Target Position. It never silently substitutes a purported external object for the represented position.

A Reference Target Position may be occupied by a Referent or remain unoccupied. Its represented content may be fictional, merely possible, or unresolved; those descriptions do not classify the physically realized position itself.

The canonical occupancy signature is:

    OCCUPIES_REFERENCE_TARGET(referent_x, target_position_y, domain_D, interval_I)

`EXISTS(y, D, I)` does not entail that any x satisfies both `EXISTS(x, D, I)` and `OCCUPIES_REFERENCE_TARGET(x, y, D, I)`. A functioning Symbol does not prove occupancy.

#### Referent

A **Referent** is the Role occupied by an Existent x exactly where `OCCUPIES_REFERENCE_TARGET(x, y, D, I)` obtains for a Reference Target Position y and the Representation, Measurement, Name, or Claim is answerable to x.

A failed or empty Reference has a Reference Target Position but no Referent; an unresolved Reference has no established Referent until existence, occupancy, and fit are resolved.

#### Absence

**Absence** is a scoped negative Status. Depending on category, it means that a Term purported to denote an Entity does not satisfy EXISTS or is not Present; a Property is not attributed; a Relation or condition does not OBTAIN; or a specified Occurrence does not occur within the Domain and Interval.

Absence is domain-relative. “No signal detected” means no signal resolved under the stated conditions; it does not automatically mean no signal existed.

Absence is not a substance. Nothing is not an entity.

#### Zero

**Zero** is a formal quantity value or identity element defined within a mathematical or accounting structure.

Zero may represent none of a counted kind, a neutral point, a reference value, or another structure-specific condition.

Zero does not universally mean nonexistence, emptiness, origin, or absence.

### Actuality, Possibility, Capacity

#### Actual

The **Actual** comprises Existents for which EXISTS holds and conditions or Relations for which OBTAINS holds in the stated Domain and Interval.

Actuality does not depend on certainty about it.

##### Presence

**Presence** is the condition of an entity, occurrence, signal, or relation being physically or operationally available within a specified domain and interval.

Presence must name where, when, and under which relation it obtains.

Physical presence, perceptual presence, network presence, social presence, legal presence, and recorded presence are different. Presence does not by itself establish identity, agency, attention, intention, consent, truth, or value.

#### Possible

A **Possibility** is a condition or occurrence compatible with stated constraints.

Possibility is always relative to a domain of constraints. Logical possibility, physical possibility, biological possibility, technical possibility, and legally permitted possibility are different.

#### Impossible

An **Impossibility** is a condition excluded by the relevant constraints.

Failure to know a mechanism does not prove impossibility. Contradiction within an exact formal system may.

#### Potential

**Potential** is a structured set of unrealized possibilities or capacities available under stated conditions.

Potential is real as capacity or openness. A possible outcome is not actual merely because it is represented, desired, priced, promised, or expected.

#### Capacity

A **Capacity** is the physically or organizationally realized ability of an entity or system to enter specified states or produce specified effects under suitable conditions.

Possessing Capacity does not imply exercising it. Capacity does not create Permission, Normative Authority, Obligation, Intention, or warranted reliance.

#### Disposition

A **Disposition** is a tendency or capacity to respond in a specified way under triggering conditions.

Fragility, solubility, reactivity, sensitivity, and learned response are different dispositions.

#### Constraint

A **Constraint** is a typed rule, relation, boundary condition, or invariant that excludes one or more otherwise expressible configurations or transitions within a Domain.

Constraints may be physical, structural, biological, computational, formal, institutional, or self-adopted.

#### Threshold

A **Threshold** is a defined condition at which a classification, regime, response, or transition changes.

A threshold does not cause the change unless a mechanism connects crossing it to consequence.

#### Probability

A **Probability** is a measure over possibilities under a specified model, state, and information basis.

Probability may represent physical stochastic structure, statistical frequency, rational uncertainty, or a model estimate. These uses must not be silently merged.

Probability is not possibility, prediction, or destiny.

### Causal Structure

#### Causation

A **Productive Causal Relation** obtains where an Occurrence or standing Process contributes through a physical or organized Mechanism to another Occurrence.

A **Causal Dependence Relation** obtains where an Effect would differ under a stated counterfactual change to a relevant State, condition, Presence, or Absence. Claims that an omission caused an Effect use such a dependence contrast; Absence is not converted into a productive Entity.

A **Causal Explanation** may cite enabling conditions, Constraints, or Rules in addition to productive causes and dependencies. Explanatory relevance alone does not prove productive causation. Temporal order is necessary for many causal accounts but is not sufficient proof of any Causal Relation.

### Evidence Across Domains

Evidence bears on Claims in every Domain, but it is an epistemic Relation rather than a second representation kind. Its canonical definition and placement statuses are owned by Part VII. No physical, biological, agential, social, economic, computational, or formal Domain may silently replace that definition with “whatever supports the preferred result.”

### History, Reversibility, and Finality

#### History

A **History** is an ordered sequence or structure of states and occurrences attributed to a bearer, system, or domain.

A recorded history is necessarily selective. The actual past is not altered by changing its later representation.

#### Reversibility

A process is **reversible** when a permitted inverse process can restore the relevant prior condition within the stated resolution.

Physical reversibility, logical reversibility, transactional reversal, correction, refund, and institutional annulment are not the same.

#### Irreversibility

**Irreversibility** is the absence or practical inaccessibility of an inverse under stated physical, computational, or institutional conditions.

Thermodynamic irreversibility is not identical to protocol finality.

#### Finality

**Finality** is a system-relative status after which ordinary rules no longer permit a result to be revised.

Finality may be physical, computational, procedural, legal, or accounting-specific. No digital finality makes its physical carrier indestructible or turns its content into truth.

### Information, Signal, Representation

#### Signal

A **Signal** is a physically realized variation usable as a distinction relative to a receiver or process.

A signal has a carrier, encoding or discriminable form, domain, interval, and possible receivers.

A signal is not automatically a message, meaning, evidence, command, intention, truth, or action.

#### Noise

**Noise** is variation treated as irrelevant, uncontrolled, or disruptive relative to a signal criterion.

The same physical variation may be signal under one encoding and noise under another.

Noise is relational to a detection or interpretation task, though its physical source is real.

#### Data

**Data** are recorded distinctions organized under an encoding or schema.

Data may be accurate, inaccurate, incomplete, corrupted, fabricated, duplicated, or uninterpretable.

Data do not speak without a decoding relation.

#### Information

**Information** is structured distinguishability capable of changing the state of a receiving or interpreting system under a defined relation.

Several information kinds must remain distinct:

- **physical information** — physically distinguishable alternatives;
- **statistical information** — uncertainty structure under a probability model;
- **algorithmic information** — description length relative to a formal machine;
- **semantic information** — content represented for an interpreter;
- **pragmatic information** — distinctions usable for action in a context.

Information is not energy, although every physical handling of information requires physical resources.

#### Encoding and Decoding

An **Encoding** is a Mapping from distinctions in one Domain to physically realized distinctions in another.

A **Decoding** is an inverse or interpretive mapping sufficient to recover selected distinctions.

A lossless or bijective Encoding may preserve every declared distinction in its source content. No Encoding thereby preserves every Property, Context, causal relation, or material feature of the source Bearer unless those too are included in its declared Domain. What is preserved and what is lost must be stated.

#### Sign and Symbol

A **Sign** is a physically realized form that indicates or stands for something within an interpretive relation.

A **Symbol** is a sign whose use depends on a convention, rule, or learned mapping.

The mark, its token, its type, its referent, and its meaning are distinct.

#### Representation

A **Representation** is a physically realized Pattern that stands for, models, or presents a Reference Target under a Mapping.

    Representation SUBTYPE_OF Pattern
    Representation SUBTYPE_OF Bearer

The canonical signature is:

    REPRESENTS(representation_a, target_position_y, mapping_S, domain_D, interval_I)

`a` is the Representation, `y` is the Reference Target Position, and `S` is the Mapping or interpretive practice. A Referent, where established, occupies `y`; it is not substituted for `y` merely because it is purportedly represented.

Representation requires:

- a vehicle or carrier;
- represented content;
- a Reference Target or intended Domain;
- a mapping or interpretive practice;
- conditions of use.

A representation may exist even when inaccurate. Its existence does not make its content true.

#### Meaning

**Meaning** is the content established through relations among a Sign, interpretive System, Context, and Reference Target or use.

Meaning is not intrinsic to shape, frequency, bit pattern, or word alone.

Meaning may be shared, contested, ambiguous, revised, or unknown.

#### Message

A **Message** is a signal or record organized for semantic reception under an encoding.

A transmitted message may fail to be received, decoded, understood, believed, or acted upon.

#### Record

A **Record** is a physically persistent representation intended to preserve selected information about a state, event, process, claim, or decision.

A record has a carrier, format, provenance, interval of persistence, and conditions of recoverability.

A record may be authentic and still represent false content. It may be immutable under one protocol and physically destructible.

#### Memory

**Memory** is a capacity by which prior conditions influence present or future operation through retained physical differences.

Biological, cognitive, computational, institutional, and archival memory have different mechanisms.

Memory is reconstruction from retained traces, not the past itself.

### Identity, Naming, Provenance

Ontology distinguishes identity relations. Continuity determines how identity may persist through change.

#### Numerical Identity

**Numerical Identity** is the relation of being one and the same bearer, not merely similar or equivalent.

Strict numerical identity is reflexive, symmetric, and transitive without varying by comparison purpose. If the criterion changes what counts as “the same,” the relation is Equivalence or Continuity, not numerical identity.

#### Qualitative Identity

**Qualitative Identity** is exact or specified sameness of properties under a comparison.

Two entities can be qualitatively identical while numerically distinct.

#### Similarity

**Similarity** is partial correspondence under selected properties.

Similarity is domain- and criterion-dependent.

#### Equivalence

**Equivalence** is interchangeability under a specified relation or purpose.

Equivalent does not mean identical in every respect.

    EQUIVALENT_UNDER(x, y, criterion, domain, interval)

`CONTINUES_AS(x, y ; continuity_criterion_and_version = q)` is a two-relatum relation owned by Continuity. The mandatory attributed qualifier `q` jointly identifies the exact continuity criterion and its exact version under the qualified relation grammar in Part I; it is not a relatum. Neither `CONTINUES_AS` nor `EQUIVALENT_UNDER` entails `SAME(x, y)`.

#### Name

A **Name** is a sign used to refer to an entity, kind, relation, occurrence, or representation.

A name does not create, contain, own, or uniquely identify its referent without a maintained naming relation.

#### Identifier

An **Identifier** is a sign intended to distinguish one referent within a namespace.

Identifiers may collide, change, be reused, forged, lost, or point incorrectly.

An Identifier may support reference to one bearer within a namespace. It does not establish strict numerical identity by itself.

#### Address

An **Address** is a locator or routing identifier within a system.

An Address is not the Entity, Agent, person, place, ownership relation, or Authority relation to which it may route.

#### Technical Account

A **Technical Account** is a system-defined record and namespace that groups selected state, operations, and permissions around a declared reference.

A Technical Account may represent or be controlled by an Agent. It is not the Agent and is not an Accounting Account.

#### Profile

A **Profile** is a representation that groups selected attributes, records, predictions, or claims about a referent.

A profile is not the person, agent, subject, or complete history represented.

#### Realization Instance and Copy

A **Realization Instance** is a particular realization of a Formal Type, Program, Pattern, or System. It may also satisfy `INSTANCE_OF`, but it does not redefine the generic classification relation.

A **Copy** is a new realization produced to preserve selected structure from a source.

Bitwise equality does not alone settle whether a copy is the same continuing system, a distinct instance, or a legally controlled reproduction.

#### Provenance

**Provenance** is the actual ordered structure of origin, derivation, custody, transformation, and attribution relations through which a Referent has passed.

A **Provenance Claim** is a Claim about some part of that history. A **Provenance Record** is a persistent Representation offered as Evidence for such a Claim.

Actual Provenance may obtain without surviving evidence. A Provenance Record may exist while being incomplete or false. None of provenance, claim, or record is ownership, truth, Authority, identity, or value.

#### Lineage

A **Lineage** is an ordered relation of descent, derivation, succession, or transformation among entities, forms, or systems.

Biological descent, document revision, software fork, institutional succession, and model derivation are different lineage types.

Lineage does not create a single continuing identity unless a continuity criterion says so.

#### Origin Record

An **Origin Record** is a record claiming or evidencing an origin relation.

The record does not create the historical origin merely by existing.

### Coherence, Integrity, Error

#### Consistency

**Logical Consistency** is the absence of an asserted proposition and its exact negation within the same formal interpretation, scope, and conditions.

Consistency is not truth. A consistent account may be false or incomplete.

#### Compatibility

**Compatibility** is the capacity of states, claims, rules, or systems to coexist or compose without violating specified constraints.

Compatibility is typed and domain-relative.

#### Contradiction

A **Contradiction** obtains between propositions where one asserts what the other denies in the same respect, scope, and interval.

The two claim-records can both exist. Their contradictory contents cannot both obtain in the same exact respect under classical logic.

Preserving both records is not accepting both as true.

#### Conflict

A **Conflict** is a relation of incompatibility among aims, claims, states, requirements, actions, or interests.

Conflict need not be collapsed immediately. It can be preserved as typed unresolved structure.

#### Ambiguity

**Ambiguity** is the availability of more than one admissible interpretation or referent for an expression under available context.

Ambiguity is not necessarily error. It becomes dangerous when a system treats distinct interpretations as one without declaration.

#### Uncertainty

**Uncertainty** is bounded lack of determination in a physical state, model, measurement, inference, or knowledge condition.

Physical uncertainty, statistical uncertainty, measurement uncertainty, semantic ambiguity, and ignorance are distinct.

#### Error

An **Error** is a mismatch between a representation, inference, action, process, or result and the relevant referent, rule, intention, or expected condition.

Errors are occurrences or relations to be typed, not impurities that disappear when hidden.

#### Correction

A **Correction** is an action or process that addresses an identified error through a new state, representation, or relation.

A correction does not rewrite the actual past. It changes what now obtains and may preserve the erroneous predecessor as history.

#### Fidelity

**Fidelity** is the degree to which a representation, transmission, transformation, or successor remains answerable to the relevant referent, source, or invariants.

Fidelity is not exact copying in every case. It requires declaring what must be preserved.

#### Integrity

**Integrity** is preservation of defining organization, constraints, or commitments across operation and change.

Integrity is not immobility, purity, secrecy, or moral perfection.

#### Coherence

**Coherence** is the condition in which differentiated parts, relations, claims, scales, and transformations can be held together without illicit collapse of their distinctions.

Coherence does not require:

- sameness;
- agreement;
- absence of tension;
- immediate resolution;
- one perspective;
- one scale;
- one vocabulary.

A system can preserve contradiction among claims coherently by keeping their sources, scopes, statuses, and incompatibility explicit.

Coherence fails where distinctions required for truth, causality, identity, an Authority relation, or Consequence are silently merged.

#### Stability

**Stability** is persistence of a state or organization under specified perturbations and interval.

Stability is not permanence.

#### Robustness

**Robustness** is persistence of function or organization across a defined range of disturbances or variation.

#### Resilience

**Resilience** is the capacity to recover or reorganize after disturbance while retaining defining organization or changing it under a declared criterion.

#### Adaptation

**Adaptation** is change through which fit between a system and its conditions is maintained or increased.

Physiological adaptation, individual learning, technical adaptation, institutional adaptation, and biological evolution are distinct.

#### Attractor

An **Attractor** is a set or region toward which nearby trajectories evolve under specified dynamics.

An attractor is not desire, intention, value, destiny, or agency.

#### Feedback

**Feedback** is a dynamical relation in which a consequence of a process alters conditions that influence its later operation.

Positive feedback amplifies a deviation under a model. Negative feedback counteracts it. Positive does not mean good; negative does not mean bad.

## Part VII — Epistemic Placement

### Claims, Evidence, and Knowledge

#### Prediction

A **Prediction** is a Representation of an expected future condition derived from a model, inference, or judgment. It may be supported, calibrated, mistaken, or unresolved. It is not a future Fact and does not make the predicted condition actual.

#### Proposition and Statement Token

A **Proposition** is repeatable semantic content capable of being asserted and evaluated for truth or falsity under an interpretation.

A **Statement Token** is a physically realized sentence, inscription, utterance, or encoded expression that presents a Proposition.

The Proposition, its Statement Token, an act of claiming it, and any Fact to which it corresponds are distinct. A Proposition need not be asserted; an asserted Proposition need not obtain.

#### Claim

A **Claim** is a represented proposition presented as obtaining.

A claim has:

- content;
- source;
- target or domain;
- time or interval;
- scope;
- conditions;
- an epistemic status.

The existence, authenticity, source Authority, warranted reliance, and truth of a Claim are separate questions.

#### Truth and Falsity

A claim is **true** where what it asserts obtains within the stated domain, scope, and interval.

A claim is **false** where its asserted condition does not obtain.

Truth is not created by repetition, Consensus, Signature, Verification, storage, immutability, any Authority relation, Price, or computational Finality.

Reality is not revised when a record is revised.

#### Evidence

**Evidence** is information that bears on a specified claim under a stated method.

Evidence may support, weaken, underdetermine, or fail to address a claim.

Evidence is not truth, Proof, Authority, Permission, or Decision.

#### Observation

An **Observation** is an interaction through which a system becomes sensitive to and records selected features of a referent.

Observation does not require consciousness.

An observation result is model-, instrument-, and resolution-dependent. It need not reveal a pre-existing classical value in every physical domain.

#### Measurement

A **Measurement** is a controlled observation that maps a specified observable or quantity to a result under a method, reference, unit, resolution, and uncertainty.

Precision, accuracy, resolution, and certainty are distinct.

#### Knowledge

**Knowledge** is a reliable epistemic relation in which a knower can correctly distinguish or act with respect to a referent under relevant conditions.

No single record, assertion, confidence score, or fluent expression proves knowledge.

The precise standards for warranted knowledge are domain-dependent and may remain contested.

#### Belief

A **Belief** is a cognitive or represented commitment that a claim obtains.

Belief may be true or false, warranted or unwarranted, conscious or dispositional.

Belief does not alter external reality by being held.

#### Proof

A **Formal Proof** is a rule-governed derivation from premises within a formal system.

A **Physical Proof** is not a separate ontological kind; empirical support is evidence under scientific method.

A valid formal proof establishes a theorem relative to its axioms and rules. It does not automatically establish that its premises describe physical reality.

#### Verification and Validation

**Verification** tests whether a representation, process, or implementation conforms to specified criteria.

**Validation** tests whether the criteria or model are adequate for their intended referent or use.

Verification does not create truth. Validation does not make a model universal.

### Epistemic Status

Ontology must distinguish what kind of thing is being described from how it is known.

An entity does not change kind merely because evidence about it changes.

Epistemic statuses are non-exclusive descriptors unless a domain-specific rule states otherwise. They are not the closed return algebra of existence placement or Kind conformance.

#### Observed

**Observed** means disclosed through a specified observation relation.

#### Measured

**Measured** means assigned a result through a specified measurement.

#### Inferred

**Inferred** means supported through reasoning from observations, models, or other claims.

#### Established

**Established** means supported strongly and repeatedly enough within a domain that alternatives are substantially constrained.

Established does not mean infallible, complete, or universal.

#### Posited

**Posited** means introduced by a Model or hypothesis as a possible explanatory Reference Target or structure.

#### Hypothetical

**Hypothetical** means treated as possible for testing or reasoning without established actuality.

Naming a hypothetical Reference Target does not establish that a Referent exists.

#### Disputed

**Disputed** means that live incompatible accounts remain under relevant standards.

Dispute does not imply equal evidence on every side.

#### Unknown

**Unknown** means not presently resolved for the relevant knower, system, or field.

Unknown does not mean nonexistent, unknowable, unconstrained, or arbitrary.

#### Indeterminate

**Indeterminate** means not assigned a unique value or outcome under the relevant physical or formal conditions.

Indeterminate is not the same as merely unknown.

#### Underdetermined

**Underdetermined** means that available evidence supports more than one materially distinct account.

#### Unknowable

**Unknowable** is a strong modal claim that no admissible process could resolve a matter.

It must not be inferred merely from current ignorance or practical difficulty.

## Part VIII — Semantic Firewall

### Canonical Separation Register

This register is an index of separations declared by their owning definitions. It is never a second source of meaning.

The ontology uses distinct separation relations. Part VIII owns the meta-relations introduced here; their arguments are represented Terms, Kind declarations, Judgments, Claims, Evidence, or Domain-scoped expressions rather than new physical species.

| Relation | Meaning |
|---|---|
| `DIFFERENT(a, b)` | a and b are numerically different existents; `SAME(a, b)` is false; this is the Part I primitive |
| `NON_EQUIVALENT(A, B)` | A and B are not definitionally interchangeable; instances may still overlap |
| `DISJOINT_WITH(K, L, C)` | no instance may conform to both kinds in the same respect under C; this is the Part III type relation |
| `NO_ENTAILMENT(P, Q)` | P alone does not imply Q |
| `INSUFFICIENT_EVIDENCE(E, Q, M)` | E alone does not warrant Q under method M and the relevant epistemic criterion |
| `NON_REDUCTION(A, B)` | the ontology does not identify A as exhaustively nothing but B |
| `DOMAIN_SEPARATED(T, D₁, D₂)` | one expression has separately typed meanings in different domains |

Their signatures are:

    NON_EQUIVALENT(term_or_kind_A, term_or_kind_B)
    NO_ENTAILMENT(judgment_P, judgment_Q)
    INSUFFICIENT_EVIDENCE(evidence_E, claim_Q, method_M)
    NON_REDUCTION(kind_or_account_A, kind_or_account_B)
    DOMAIN_SEPARATED(term_T, domain_D₁, domain_D₂)

`NON_EQUIVALENT` does not mean `DISJOINT_WITH`. `NO_ENTAILMENT` does not deny that Q may independently obtain. `INSUFFICIENT_EVIDENCE` is an epistemic limit, not a negative fact about the world. No method-default form is assumed for it. `DOMAIN_SEPARATED(T@D₁, T@D₂)` is exact display sugar for `DOMAIN_SEPARATED(T, D₁, D₂)`. None of these meta-relations silently replaces `SAME`, `DIFFERENT`, `INSTANCE_OF`, or `DISJOINT_WITH`.

#### Reality and Representation

| Typed separation | Register entry |
|---|---|
| `NON_EQUIVALENT` | Reality / Model; Referent / Representation; Event / Record; Fact / Statement Token; map / territory; Profile / Bearer |
| `NO_ENTAILMENT` | Claim existence → Claim truth; consensus → truth; verification → truth creation |
| `INSUFFICIENT_EVIDENCE` | Record authenticity → content truth |

#### Kind, Role, and Identity

| Typed separation | Register entry |
|---|---|
| `NON_EQUIVALENT` | Kind / Role; Formal Type / Type Token; Entity / State; Bearer / Identifier; numerical identity / similarity; numerical identity / equivalence; numerical identity / Continuity; numerical identity / Lineage |
| `NON_EQUIVALENT` | Agent / Actor role; Agent / Participant role; Actor role / Participant role; Subject / Agent; Human / Personhood Status; Agent / Technical Account |
| `NO_ENTAILMENT` | key possession → Agency or identity; derivation → ownership |

#### Structure, Relation, and Change

| Typed separation | Register entry |
|---|---|
| `NON_EQUIVALENT` | Relation / relatum; Relation / fusion; Boundary / wall; Interface / identity; Mapping / equivalence; System / Aggregate; component list / System; organization / material; Function / purpose |
| `NO_ENTAILMENT` | overlap → shared sovereignty; Boundary → ownership; Presence → identity or Agency; Effect → Authorization |
| `NON_EQUIVALENT` | State / Process; Event / Process; Transition Occurrence / Transition Record; Transition Relation / Transition Record; maintenance / stasis; Correction / erasure; reversal / Correction; physical irreversibility / protocol finality |
| `NO_ENTAILMENT` | Prediction → future Fact; Plan → Action; Promise → performance; Capacity → Permission |
| `NON_EQUIVALENT` | Possibility / Actuality; Possibility / Probability; Potential / possession; Channel / Message; Impact / Contribution Assessment |

#### Physical-Domain Separations

| Typed separation | Register entry |
|---|---|
| `NON_EQUIVALENT` | spacetime / coordinate grid; Field / Force; Field / Energy; Field / Excitation; Quantum / Particle; Vacuum / nothing; Matter / Mass; Radiation / Energy; Energy / substance; entropy / ordinary disorder |
| `NO_ENTAILMENT` | temporal sequence → causation; proximity → Interaction; frequency equality → existence, interaction, communication, or resonance; correlation → causation or communication |
| `NON_REDUCTION` | Particle / tiny classical ball; Boson / Force; Fermion / Matter |
| `DOMAIN_SEPARATED` | physical Vector / agential direction or Intention |
| `NO_ENTAILMENT` | effective description → unreality; microscopic scale → greater reality |

#### Information, Cognition, and Experience

| Typed separation | Register entry |
|---|---|
| `NON_EQUIVALENT` | Physical Carrier / Signal; Signal / Message; Signal / Meaning; Data / Knowledge; Information / Energy; Representation / understanding; Memory / past; Self-Model / Functional Self / Phenomenal Self |
| `NO_ENTAILMENT` | Signal → Intention; Intelligence → Consciousness, Agency, or Personhood Status; complexity → Intelligence; Simulation → every simulated kind |
| `INSUFFICIENT_EVIDENCE` | self-report alone → Consciousness; Measurement Result alone → unmediated revelation of a referent |
| `NON_EQUIVALENT` | Computation / arbitrary physical evolution |

#### Agency and Institution

| Typed separation | Register entry |
|---|---|
| `NON_EQUIVALENT` | Action / Intention; Ownership / possession; physical Law / normative Law; Access / Permission |
| `NO_ENTAILMENT` | Effect → Authorization; capability → Normative Authority; Control → Normative Authority; authorship → Ownership; operation → Ownership; sponsorship → Control; Participation → submission; Admission → Ownership; Contribution → entitlement; Power → Normative Authority |
| `INSUFFICIENT_EVIDENCE` | Digital Signature → Consent; key possession → identity; Reputation → identity |
| `NO_ENTAILMENT` | Trust → truth; social construction → unreality; normative-law violation → exception to physical law |
| `NON_EQUIVALENT` | Freedom / absence of all constraint |

#### Value and Economy

| Typed separation | Register entry |
|---|---|
| `NON_EQUIVALENT` | Evaluative Value / Price; Price / worth; Money / Economic Value; Protocol Token / Money; Asset / physical Object; fungibility / variability; non-fungibility / immutability; Accounting / Ontology; Balance / Bearer |
| `NO_ENTAILMENT` | Money → truth; Protocol Token → Ownership; Credit → future Actuality; Debt → guaranteed performance; Settlement → truth; Transparency → lawfulness; Capital → Normative Authority; Biological Fitness → worth; reward → Reason |

#### Knowledge and Open Reality

| Typed separation | Register entry |
|---|---|
| `NON_EQUIVALENT` | Unknown / nonexistent; unresolved / unconstrained; Hypothetical / Established; absence of Record / Record of absence; incomplete Model / false Model; domain validity / universal validity; precision / accuracy; uncertainty / ambiguity; ignorance / indeterminacy |
| `NO_ENTAILMENT` | failure to classify → nonexistence; failure to observe → absence; present ignorance → unknowability |

## Part IX — Open Extension Boundary

### Extension Rule

A new kind K may extend this ontology only by declaring:

- its parent kind or explicit lack of one;
- its necessary and sufficient conditions, or which side remains open;
- its admissible bearers;
- its properties and value-spaces;
- its relations, with arity and domain/range constraints;
- its realization condition;
- its domain, interval, scale, and resolution;
- declared disjointness and overlap;
- its epistemic statuses;
- its non-equivalences;
- its boundary to later layers.

A new name without these declarations is vocabulary, not an ontological type.

The declaration may explicitly leave necessary conditions, sufficient conditions, admissible bearers, relation signatures, realization criteria, scope fields, disjointness, overlap, epistemic statuses, or boundaries `OPEN` where the schema permits it. Openness is typed incompleteness, not permission to smuggle in an answer.

### Unresolved Existence and Unresolved Kind

An encountered or represented target need not fit an existing Kind before it can be examined. Existence placement and Kind conformance must remain separate.

The closed existence-placement operator is:

    EXISTENCE_PLACEMENT(x, D, I, R)
      -> EXISTENCE_ESTABLISHED
       | NONEXISTENCE_ESTABLISHED
       | EXISTENCE_UNRESOLVED
       | ILL_TYPED

Its result tokens are placement results, not Part VII Epistemic Statuses.

- `EXISTENCE_ESTABLISHED` — available evidence supports `EXISTS(x, D, I)`;
- `NONEXISTENCE_ESTABLISHED` — scoped exhaustive evidence supports `¬EXISTS(x, D, I)`;
- `EXISTENCE_UNRESOLVED` — neither existence nor nonexistence is established;
- `ILL_TYPED` — the placement request violates the declared grammar.

Failure to establish existence yields `EXISTENCE_UNRESOLVED`, not `NONEXISTENCE_ESTABLISHED`. The negative result requires evidence adequate to the stated Scope and preserves the open-world rule.

An unresolved-existence judgment is:

    Γ ⊢ EXISTENCE_PLACEMENT(x, D, I, R) = EXISTENCE_UNRESOLVED

This is an **Unresolved-Existence Placement**. It records a Term and the Evidence bearing on whether its Reference Target is occupied by an Existent, without typing that target as an Existent.

Only after the evidence supports `EXISTS(x, D, I)` may more-specific Kind conformance be evaluated:

    Γ ⊢ x : Existent @ D, I, R
    Γ ⊢ SUPPORTED_KINDS(x, D, I, R) = {K | Γ ⊢ x : K @ D, I, R = CONFORMS}

`SUPPORTED_KINDS` contains every declared Kind K for which the ordinary per-candidate judgment `Γ ⊢ x : K @ D, I, R` returns `CONFORMS`. Candidate judgments may independently return `CONFORMS`, `DOES_NOT_CONFORM`, `UNRESOLVED`, or `ILL_TYPED`.

An **Unresolved-Kind Placement** records established existence together with each candidate Kind whose conformance remains `UNRESOLVED`. Multiple compatible or incomparable Kinds may independently conform; no unique “most-specific” Kind is presumed.

This placement also covers the case in which no existing more-specific Kind is an adequate candidate.

Either placement record declares:

- the Term or provisional Reference used for x;
- the evidence and current epistemic statuses of existence;
- the Domain, Interval, and Resolution Context of encounter;
- observed Properties, Relations, and Effects without treating them as a complete essence;
- candidate Kinds and the evidence for or against each;
- declared exclusions and unresolved contradictions;
- the minimum distinctions that any later Kind must preserve.

Neither placement is a miscellaneous species or permanent residual bin. Each is an explicit Status preventing three errors: denial because no name exists, fictionalization because Evidence is incomplete, and forced assignment to the nearest familiar Kind.

### Open Boundaries

This ontology explicitly leaves unresolved what reality and present knowledge do not yet settle.

Open boundaries include:

- whether spacetime, quantum fields, particles, or some deeper structure are fundamental;
- the ontology of quantum state, measurement, and definite outcomes;
- a complete quantum account of gravitation;
- the constitution of dark matter and dark energy;
- the origin of physical laws, constants, cosmological initial conditions, and the thermodynamic arrow;
- the full boundary between nonlife and life;
- the origin and sufficient conditions of cognition and agency;
- the physical basis and distribution of subjective experience;
- whether unfamiliar artificial or derivative intelligences are subjects;
- how semantic reference and meaning arise in every case;
- the independent status, if any, of uninstantiated mathematical structures;
- the conditions under which collective systems become agents;
- the complete relation among physical realization, higher-level causation, and explanation;
- the eventual kinds that current instruments, concepts, and scales cannot yet resolve.

The ontology must be capable of locating an unknown without inventing its answer. The two-stage unresolved placement is that route.

An untyped or unresolved Reference Target may correspond to a real Existent.

A new kind may be admitted only by making its distinctions clearer, not by borrowing force from metaphor.

### Boundary to the Reality-Model Stack

Ontology names what kinds of reality may be distinguished. Naming a Kind is not governing it, arranging it, making it persist, or explaining how it emerges.

The physical ground is supplied by the preceding **Substrate** layer. The following matters remain for the layers that follow Ontology:

- **Constitution** — which constituting conditions, invariants, permissions, prohibitions, scopes, and lawful relations govern what may stand, relate, and act within a world;
- **Structure** — how entities, relations, systems, scales, interfaces, overlaps, and domains are arranged;
- **Dynamics** — how states change, processes execute, interactions propagate, feedback operates, and failures terminate;
- **Continuity** — under which criteria a bearer persists, succeeds, branches, merges, repairs, copies, dissolves, or ends across change;
- **Emergence** — how new effective properties, capacities, life, cognition, agency, collective intelligence, meaning, and value arise from organized relations without escaping the Substrate.

Examples of the boundary:

- Ontology names Right, Permission, Consent, Authorization, and Authority relations; Constitution decides which instances are valid and with what effect.
- Ontology names System, component, Boundary, Interface, and Seam; Structure declares a particular topology.
- Ontology names State, Transition Occurrence, Transition Relation, Interaction, Feedback, and failure; Dynamics declares how a particular World changes.
- Ontology names History, Copy, Lineage, identity, and ending; Continuity supplies persistence and succession criteria.
- Ontology names life, cognition, Agency, collective intelligence, Meaning, and value; Emergence explains how particular organized realizations acquire those capacities.
- Ontology types fields, particles, matter, and physical interactions; Substrate supplies the physical constraints under which they can obtain.

Later layers may specialize these kinds. They may not silently collapse the distinctions declared here or send a later-layer validity judgment backward as an ontological definition.

### The Ontological Ground That Remains

Reality contains entities and occurrences, properties and quantities, relations and patterns, states and processes, domains and boundaries, signals and meanings, organisms and environments, minds and experiences, agents and actions, institutions and economies, computations and worlds.

These are not one kind of thing.

Every effective form is physically realized.  
No representation becomes its referent.  
No role becomes its bearer.  
No relation erases its relata.  
No record becomes the event.  
No possibility becomes actual by anticipation.  
No capability becomes Normative Authority by force.  
No intelligence becomes a subject by label.  
No unknown becomes unreal by omission.  
No domain becomes universal by success within its boundary.

Reality may exceed every present type.

The ontology remains coherent by preserving that excess as unknown rather than converting it into absence, metaphor, or invention.

This is the complete root grammar presently required to encounter reality without closing it.
