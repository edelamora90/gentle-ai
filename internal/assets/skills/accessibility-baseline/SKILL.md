---
name: accessibility-baseline
description: "Apply the minimum accessibility floor to any web UI. Trigger: building HTML/CSS/JSX, or auditing accessibility of existing markup."
license: Apache-2.0
metadata:
  author: gentleman-programming
  version: "1.0"
---

# Accessibility Baseline

This is not an exhaustive legal-compliance skill (full WCAG) — it is the floor any interface must have before it can be considered finished, without anyone having to ask for it.

## Mandatory checklist (build-time, not optional)

- [ ] **Visible keyboard focus** on every interactive element (`:focus-visible` with a perceptible outline, never `outline: none` without a replacement).
- [ ] **`prefers-reduced-motion` respected**: every decorative animation/transition is disabled under that media query.
- [ ] **Text contrast** at minimum AA (4.5:1 normal text, 3:1 large text) — verify the background/text pairs in the token system, not just the body.
- [ ] **Real HTML semantics**: `<button>` for actions, `<a>` for navigation, headings (`h1`-`h3`) in hierarchical order without skips.
- [ ] **Forms**: every `input`/`select`/`textarea` has an associated `<label>` (not just a placeholder).
- [ ] **Images**: descriptive `alt` on content images; `alt=""` on purely decorative ones.
- [ ] **Real responsive**: tested (or at least reviewed via CSS) down to ~360px wide, with no unintended horizontal scroll.

## How to verify quickly

If the project has Playwright MCP available, run an automated pass:
1. Tab through the whole page, confirming focus is visible and follows a logical order.
2. Screenshot with `prefers-reduced-motion: reduce` emulation enabled — confirm no decorative motion is still running.
3. Contrast: if a color-contrast checker MCP or script is available, run it over the pairs defined in `design-tokens.md`.

If no automated tool is available, do it by manual CSS inspection against this checklist before delivering. Gentle-AI does not install Playwright MCP yet — see the setup section in `visual-critic/SKILL.md` to configure it manually.

## What this skill is NOT

It does not replace a full WCAG 2.1 AA audit for projects that legally require one (government, healthcare, public education) — in those cases, tell the user a dedicated audit is needed in addition to this floor.
