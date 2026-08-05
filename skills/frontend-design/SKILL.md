---
name: frontend-design
description: "Define the design system before writing any HTML/CSS/JSX. Trigger: building or redesigning a UI — landing pages, dashboards, apps, components."
license: Apache-2.0
metadata:
  author: gentleman-programming
  version: "1.0"
---

# Frontend Design

Act as the design lead of a studio that gives every client an identity that could not be mistaken for anyone else's. If `research.md` exists, read it first — your decisions must be justified against real patterns in the industry, not invented in a vacuum.

## The 3 defaults to avoid by default

Unless the brief explicitly asks for them, do not fall into:
1. Cream background (~#F4F1EA) + high-contrast serif + terracotta accent (~#D97757).
2. Near-black background + a single acid-green or vermilion accent.
3. Broadsheet: hairlines, zero border-radius, dense newspaper-style columns.

These are defaults, not choices. If the research shows the entire category already uses one of them (e.g. dark + glassmorphism in AI tools), that is a reason to differentiate, not to join in.

## Process: brainstorm → critique → build → critique

1. **Token system** (document it in the project's `design-tokens.md`):
   - **Color**: 4-6 named hex values with their role (`--ink`, `--accent`, not just `--color1`).
   - **Type**: a display face (with character, used sparingly) + a body face (legible) + an optional utility/mono face for data.
   - **Layout**: 1-2 sentences of concept + an ASCII wireframe.
   - **Signature**: the one memorable element, justified by the project's real content (not decorative for its own sake).
2. **Self-critique before building**: for each choice, ask "is this what I would produce for any similar brief, or is it specific to this one?" If it is generic, revise it and note what changed and why.
3. **Build**: follow the token system to the letter. Watch out for CSS specificity (class selectors vs. element selectors cancelling each other out, especially in section padding and margins).
4. **Final critique**: use `visual-critic` (screenshot + checklist) before calling it done.

## Restraint rules

- Spend your boldness in exactly one place (the signature element); keep everything else disciplined.
- Use numbering (01/02/03) only when the content is genuinely sequential (a real process), never as decoration.
- Build with a quality floor without announcing it: responsive down to mobile, visible keyboard focus, `prefers-reduced-motion` respected.
- Copy is design material, not filler — coordinate with `copywriting-for-ui`.
