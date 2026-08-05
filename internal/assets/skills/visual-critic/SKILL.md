---
name: visual-critic
description: "Screenshot and audit the built result before calling it done. Trigger: final step after the UI is built, before reporting a task complete."
license: Apache-2.0
metadata:
  author: gentleman-programming
  version: "1.0"
---

# Visual Critic

A person with taste looks in the mirror before leaving and takes off one accessory. This skill is that mirror.

## Process

1. **Screenshot** desktop (~1440px) and mobile (~375px) of the result, with Playwright MCP if available.
2. **Compare against `design-tokens.md`**: do the defined palette, typography, and signature element actually appear as planned, or did the build drift?
3. **Generic-detection question**: "If I showed this to someone without telling them the brief, could they guess it belongs to this specific client, or could it be any AI startup's landing page?" If the answer is "any startup", go back to `frontend-design`.
4. **Check decoration density**: is more than one "bold" element competing for attention? If so, remove one (the Chanel rule).
5. **Run `accessibility-baseline`** if it has not run yet.
6. **Report findings as a short actionable list**, not an essay: what is good, what gets fixed before delivery, what stays an open decision for the user.

## Expected report format

```
✅ Meets token system: palette, typography, and signature element correct
⚠️  Adjust: the hero reads generic on mobile (the signature element gets lost)
⚠️  Adjust: 2 elements competing for attention in the pricing section
❌ Missing: keyboard focus invisible on the navbar buttons
```

## Playwright MCP setup

Gentle-AI does NOT install Playwright MCP yet: there is no `playwright` component the way there is a `context7` one. Configure it by hand in your agent, following the same format gentle-ai uses for its other MCP servers:

| Agent | File | Root key |
|---|---|---|
| Claude Code | `~/.claude.json` | `mcpServers` |
| OpenCode | `~/.config/opencode/opencode.json` | `mcp` |
| VS Code | `mcp.json` | `servers` |
| Kimi / Antigravity | agent config | `mcpServers` |

```json
{
  "mcpServers": {
    "playwright": {
      "command": "npx",
      "args": ["-y", "@playwright/mcp@latest"]
    }
  }
}
```

**Fallback without MCP**: if Playwright is unavailable, this skill is NOT skipped — it runs in manual inspection mode: read the CSS and rendered markup against `design-tokens.md` and the checklist, and note in the report that there was no screenshot. A report without a screenshot is still a valid gate; a missing report is not.

> TODO: add a `playwright` MCP component under `internal/components/mcp/` following the `context7.go` pattern, so `gentle-ai install` configures it automatically.

## Rules

- Do not mark a task "done" in the memory/Engram system until this skill has run and the ⚠️/❌ findings are resolved or explicitly accepted by the user.
- If the user explicitly asked for "quick, unpolished", this skill may run in abbreviated mode (only the accessibility checklist + the generic question), but it is never skipped entirely.
- This skill runs in `sdd-verify`, not in `sdd-visual` or `sdd-apply`. An unresolved ❌ is a CRITICAL issue and blocks `sdd-archive`.
