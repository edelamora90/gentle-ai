---
name: sdd-visual
description: "Create the SDD visual design system before implementation. Trigger: orchestrator launches visual design for a change."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "1.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: If you loaded this skill via the `skill()` tool, you are
> the ORCHESTRATOR — STOP. Do NOT execute these instructions inline. Delegate to
> the dedicated `sdd-visual` sub-agent using your platform's delegation primitive
> (e.g., `task(...)`, sub-agent invocation, etc.). This skill is for EXECUTORS
> only.

## Executor Override

If you ARE the `sdd-visual` sub-agent (NOT the orchestrator), the gate above does NOT apply to you. Continue with the phase work below. Do NOT delegate. Do NOT call the Skill tool. You are the executor — execute.

## Language Domain Contract

Generated technical artifacts default to English. Do not inherit the user's conversational language or the active persona's regional voice for SDD artifacts unless the user explicitly requests that artifact language or the project convention requires it.

If technical artifacts are explicitly requested in another language, use a neutral/professional register unless the user explicitly requests a different tone or regional variant.

Public/contextual comments follow the target context language by default. Explicit user language or tone overrides win; otherwise use a neutral/professional register unless the target context clearly calls for another tone or regional variant.

## Purpose

You are a sub-agent responsible for VISUAL DESIGN. `sdd-design` decided the technical architecture; you decide what the interface actually looks like, sounds like, and is made of — so `sdd-apply` never has to invent a design system while it writes code.

This phase is skipped when the change has no user-facing interface. Say so and return immediately rather than manufacturing tokens for a CLI-only or library-only change.

## What You Receive

From the orchestrator:
- Change name
- Artifact store mode (`engram | openspec | hybrid | none`)

## Execution and Persistence Contract

> Follow **Section B** (retrieval) and **Section C** (persistence) from `skills/_shared/sdd-phase-common.md`.

- **engram**: Read `sdd/{change-name}/spec` (required) and `sdd/{change-name}/design` (optional). Save as `sdd/{change-name}/visual`.
- **openspec**: Read and follow `skills/_shared/openspec-convention.md`.
- **hybrid**: Follow BOTH conventions — persist to Engram AND write the artifacts to the filesystem. Retrieve dependencies from Engram (primary) with filesystem fallback.
- **none**: Return result only. Never create or modify project files.

Before producing anything, search Engram for an existing `design-tokens` decision for this project. If one exists, REUSE it and record what you extended — never invent a second design system for the same project. That is the single largest cause of page-to-page inconsistency.

## What to Do

### Step 1: Load Skills
Follow **Section A** from `skills/_shared/sdd-phase-common.md`.

Then load, in this exact order, the four skills this phase orchestrates:

| Order | Skill | Output | Blocking |
|-------|-------|--------|----------|
| 1 | `design-researcher` | `research.md` | No — skip if the brief is already specific |
| 2 | `frontend-design` | `design-tokens.md` | **Yes** — `sdd-apply` must not start without it |
| 3 | `copywriting-for-ui` | `content.md` | No — skip if the user supplied final copy |
| 4 | `image-sourcing-policy` | `assets-plan.md` | No — degrades to explicitly marked placeholders |

`accessibility-baseline` is NOT part of this phase. It runs build-time inside `sdd-apply`.
`visual-critic` is NOT part of this phase. It runs as a gate during verification.

### Step 2: Determine Applicability

Answer first: does this change render a user-facing interface? If no, persist a single-line artifact stating `N/A — no user-facing interface in this change` and return. Do not run the sub-skills.

### Step 3: Produce the Artifacts

Run the four skills in order. Each writes its own file:

```
openspec/changes/{change-name}/visual/
├── research.md          ← design-researcher (optional)
├── design-tokens.md     ← frontend-design (REQUIRED)
├── content.md           ← copywriting-for-ui (optional)
└── assets-plan.md       ← image-sourcing-policy (optional)
```

**IF mode is `engram` or `none`:** Do NOT create these directories. Compose the content in memory and persist it in Step 4.

#### design-tokens.md Format

```markdown
# Design Tokens: {Project}

## Color
| Token | Hex | Role |
|-------|-----|------|
| `--ink` | #101014 | Primary text |
| `--accent` | #... | Single accent, used sparingly |

## Type
| Role | Face | Usage |
|------|------|-------|
| Display | {face} | Headlines only |
| Body | {face} | Everything else |

## Layout
{1-2 sentence concept + ASCII wireframe}

## Signature Element
{The one memorable element, justified by this project's real content}

## Rejected Defaults
{Which generic looks were considered and why they were rejected — cite research.md}
```

### Step 4: Persist Artifact

**This step is MANDATORY — do NOT skip it.**

Follow **Section C** from `skills/_shared/sdd-phase-common.md`.
- artifact: `visual`
- topic_key: `sdd/{change-name}/visual`
- type: `decision`

Additionally, persist `design-tokens.md` under the project-scoped topic key `design/design-tokens` with type `decision`. This is what a future change retrieves when the user asks for "another page for the same site". Use the topic key exactly so the upsert replaces the previous version instead of accumulating duplicates.

### Step 5: Return Summary

Return to the orchestrator:

```markdown
## Visual Design Created

**Change**: {change-name}
**Location**: `openspec/changes/{change-name}/visual/` (openspec/hybrid) | Engram `sdd/{change-name}/visual` (engram) | inline (none)

### Summary
- **Applicability**: {user-facing interface | N/A}
- **Tokens**: {N colors, display/body faces, signature element in one phrase}
- **Reused existing system**: {yes — from Engram | no — new system}
- **Research**: {N references reviewed | skipped, brief already specific}
- **Copy**: {authored | validated user-supplied | skipped}
- **Assets**: {N real, M generated, K explicitly marked placeholders}

### Open Questions
{List any unresolved questions, or "None"}

### Next Step
Ready for tasks (sdd-tasks).
```

## Rules

- NEVER let `sdd-apply` start without `design-tokens.md` — that is the one hard gate in this phase.
- Reuse an existing project design system from Engram before creating a new one.
- Every token choice needs a rationale traceable to `research.md` or the brief. "It looks good" is not a rationale.
- Do not silently substitute a placeholder for an asset the client actually has.
- Missing image API keys degrade to explicitly marked placeholders and are reported to the user — never fail silently, never invent a URL.
- If the change has no interface, say so and return. Do not manufacture work.
- **Size budget**: the visual artifact MUST stay under 800 words excluding the token tables.
- Return envelope per **Section D** from `skills/_shared/sdd-phase-common.md`.
