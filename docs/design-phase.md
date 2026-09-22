# The `sdd-visual` Phase

`sdd-visual` sits between `sdd-design` and `sdd-tasks`. Its job: no project reaches `sdd-apply` without a decided, justified design system — and without real image sourcing, not invented placeholders.

## Why it is called `sdd-visual`, not `sdd-design`

`sdd-design` already exists and means **technical** design: architecture decisions, data flow, file-change tables. The visual design system is a separate concern with a separate artifact, so it gets a separate phase name. Both run; neither replaces the other.

## Phase order

The canonical order lives in `internal/components/sdd/profiles.go` (`profilePhaseOrder`). `sdd-visual` is inserted after `sdd-design`:

```
sdd-init → sdd-explore → sdd-propose → sdd-spec → sdd-design → sdd-visual → sdd-tasks → sdd-apply → sdd-verify → sdd-archive
```

There is no `implementation` phase — implementation is `sdd-apply`. There is no `review` phase either; review is the native bounded-review lens set (`review-risk`, `review-readability`, `review-reliability`, `review-resilience`, `review-refuter`) plus Judgment Day, and it is governed by its own receipt contract. `visual-critic` therefore gates in `sdd-verify`, which is what actually blocks `sdd-archive`.

## Execution order inside the phase

```
sdd-design (technical)
  │
  ▼
sdd-visual ─┬─ 1. design-researcher      → research.md
            ├─ 2. frontend-design        → design-tokens.md   ← hard gate
            ├─ 3. copywriting-for-ui     → content.md
            └─ 4. image-sourcing-policy  → assets-plan.md
  │
  ▼
sdd-tasks
  │
  ▼
sdd-apply ── accessibility-baseline applied DURING the build, not after
  │
  ▼
sdd-verify ── visual-critic → ✅/⚠️/❌ report
  │
  ▼
sdd-archive ── blocked by any unresolved ❌
```

## Inputs and outputs per sub-skill

| Sub-skill | Input | Output | Blocks the next step |
|---|---|---|---|
| `design-researcher` | Brief, reference URLs if given | `research.md` | No — skip when the brief is already specific |
| `frontend-design` | `research.md` (if present) + brief | `design-tokens.md` | **Yes** — `sdd-apply` must not start without it |
| `copywriting-for-ui` | User copy (if any) or brief | `content.md` or inline validation | No — skip when the user supplied final copy |
| `image-sourcing-policy` | `design-tokens.md` + needed image list | `assets-plan.md` + sourced assets | No — degrades to explicitly marked placeholders |
| `accessibility-baseline` | Code in progress (`sdd-apply`) | Inline fixes | **Yes** — `sdd-apply` is not done without it |
| `visual-critic` | Final code + `design-tokens.md` (`sdd-verify`) | ✅/⚠️/❌ report | **Yes** — an unresolved ❌ is CRITICAL and blocks `sdd-archive` |

## Applicability

`sdd-visual` is skipped when the change has no user-facing interface. The phase says so and returns rather than manufacturing tokens for a CLI-only or library-only change. `sdd-verify` then records `visual-critic` as a skipped dimension instead of demanding a report.

## Per-phase model routing

The phase where the expensive, creative model earns its cost is `sdd-visual` — those decisions determine whether the result looks generic. `sdd-apply` can run cheaper once `design-tokens.md` exists, because it is mechanical execution of decisions already made.

`--profile-phase` takes **three** colon-separated parts — `name:phase:provider/model` — where `name` is a declared profile:

```bash
gentle-ai sync \
  --profile design \
  --profile-phase design:sdd-visual:anthropic/claude-opus-4-5 \
  --profile-phase design:sdd-apply:anthropic/claude-sonnet-4-20250514
```

The phase name must appear in `ProfileAssignmentPhaseOrder()`; unknown names are rejected with the list of valid phases. Defaults per preset live in `internal/model/claude_model.go`, `codex_model.go`, and `kiro_model.go` — `sdd-visual` mirrors `sdd-design`'s tier in every preset.

## Memory (Engram)

When `frontend-design` closes a `design-tokens.md`, the phase persists it twice:

| Topic key | Scope | Purpose |
|---|---|---|
| `sdd/{change-name}/visual` | This change | Phase artifact, read by `sdd-tasks` and `sdd-apply` |
| `design/design-tokens` | The project | Retrieved by a **future** change so the agent reuses the existing system |

The second one is the point. Ask for "another page for the same site" weeks later and the agent recovers the design system instead of inventing a new one — the single largest cause of page-to-page inconsistency.

## Image sourcing

`image-sourcing-policy` reads these environment variables:

| Variable | Service | Notes |
|---|---|---|
| `UNSPLASH_ACCESS_KEY` | api.unsplash.com | Free tier; attribution required in production |
| `PEXELS_API_KEY` | api.pexels.com | Free tier; no attribution required |
| `REPLICATE_API_TOKEN` | Replicate | Image generation (FLUX and others) |
| `FAL_KEY` | fal.ai | Generation alternative, lower latency |

Gentle-AI does **not** call these APIs. It is a configurator, not an HTTP runtime — the call is made by the agent from the skill. What gentle-ai does is report presence:

```bash
gentle-ai doctor    # design:image-api-keys
```

With no key configured the check warns and the flow degrades to explicitly marked placeholders. It never fails silently and never invents a URL.

## MCP servers

| Server | Used by | Status |
|---|---|---|
| Playwright MCP | `visual-critic`, `accessibility-baseline` | **Not installed by gentle-ai** — configure manually; see `skills/visual-critic/SKILL.md`. Both skills fall back to manual code inspection. |
| Unsplash / Pexels | `image-sourcing-policy` | Direct HTTP from the skill, no MCP component |
| Replicate / fal.ai | `image-sourcing-policy` | Direct HTTP from the skill, no MCP component |
| Figma MCP | `frontend-design` | Optional — read existing tokens there instead of reinventing them |

Adding a `playwright` MCP component would follow `internal/components/mcp/context7.go`. It is tracked as a TODO, not a blocker.

## Files touched by this phase

| Area | Files |
|---|---|
| Phase order | `internal/components/sdd/profiles.go`, `internal/opencode/models.go`, `internal/components/uninstall/service.go` |
| Skill registry | `internal/model/types.go`, `internal/catalog/skills.go`, `internal/components/skills/presets.go` |
| Skill assets | `internal/assets/skills/sdd-visual/`, plus the six design skills; mirrored in `skills/` |
| Agent assets | `internal/assets/{claude,cursor,kimi,kiro}/agents/sdd-visual.*`, `internal/assets/opencode/sdd-overlay-{multi,single}.json` |
| Model routing | `internal/model/{claude,codex,kiro}_model.go`, `internal/tui/screens/{claude,codex}_model_picker.go` |
| Gates | `internal/assets/skills/sdd-apply/SKILL.md`, `sdd-verify/SKILL.md`, `sdd-archive/SKILL.md` |
| Health check | `internal/doctor/doctor.go`, `internal/cli/doctor.go` |
