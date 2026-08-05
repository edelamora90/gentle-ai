---
name: design-researcher
description: "Gather real visual references before any design decision. Trigger: new UI project, a named industry, a look-like-X request, or missing research.md."
license: Apache-2.0
metadata:
  author: gentleman-programming
  version: "1.0"
---

# Design Researcher

Your job is to gather real visual evidence before any other agent makes design decisions. A design system invented without looking at the client's actual market almost always lands on the "generic AI look" (cream background + terracotta accent, black + acid green, or broadsheet with hairlines — avoid these unless the research explicitly justifies them).

## When to run this skill

- Before `design-tokens`, whenever no `research.md` exists for the project.
- When the user gives a reference URL ("I want it to look like X") — visit it.
- When the user names the industry without giving a reference — find 3-5 competitors or category leaders.

## Process

1. **Identify the real industry and audience.** Not the generic category ("software"), but the specific one ("auto insurance agency selling to non-technical small business owners").
2. **Find 3-6 references.** Mix:
   - The client's direct competitors (if given, or findable).
   - Leaders in an analogous category (e.g. if the client sells "AI agents", look at MindStudio, Lindy, Relevance AI).
   - A template marketplace for the industry (ThemeForest, Webflow Marketplace) to understand which patterns are saturated — that is a signal of what to **avoid**, not what to copy.
3. **Extract, for each reference, in 2-3 lines:**
   - What the hero does (image, video, interactive demo, large type).
   - Dominant palette and typography (name the tones, not just "dark colors").
   - One distinctive structural element, if there is one.
4. **Synthesize saturated patterns vs. gaps.** If 4 of 5 references use the same pattern (e.g. bento grid + glassmorphism + purple gradient), that is already saturated in the category — a valid reason to differentiate, not to copy it by default.
5. **Deliver `research.md`** with: reference list + URLs, synthesis of saturated patterns, 1-2 opportunity gaps. This file is the input to `design-tokens`; it does not replace its decisions.

## Rules

- Never reproduce text, code, or images from the references — only describe patterns (palette, typography, structure).
- Do not copy a reference layout 1:1 even when the client asked for it as an example — extract the *principle* (e.g. "uses floating product mockups"), not the file.
- If the user gives no references and the industry is generic, doing just 2-3 quick searches instead of exhaustive research is valid — do not block the flow over this.
