---
name: image-sourcing-policy
description: "Decide and execute how each interface image is sourced. Trigger: an image is needed and no client asset exists — stock, generation, or placeholder."
license: Apache-2.0
metadata:
  author: gentleman-programming
  version: "1.0"
---

# Image Sourcing Policy

There is no good shortcut for "I don't have the client's real photo". This skill decides, case by case, which of the 3 options to use, and executes it.

## Decision tree

1. **Does the client already have the asset (logo, product shot, team photo)?** Use it, always, no exceptions. Never substitute a generic image for a real asset that is available.
2. **Is it a generic "real life" photo (office, person working, ambience)?** → Real stock API (Unsplash or Pexels, see Setup). Never use "people" generators here — it shows, and it can read as misleading for a real client.
3. **Is it a conceptual illustration or hero art with no photographic equivalent (diagrams, abstract iconography, fictional product mockups)?** → Generation (FLUX via Replicate/fal.ai), or build it with SVG/CSS when that is simpler and more editable. Prefer SVG/CSS when the element is part of the design system (e.g. the "control panel" with toggles) — it is more controllable and does not depend on an external API.
4. **Is no API configured, and is this a mockup/proposal delivery rather than production?** → Explicit placeholder (see the placeholder policy below). Never deliver it as if it were final without saying so.

## API setup (expected environment variables)

```
UNSPLASH_ACCESS_KEY=...      # api.unsplash.com — free tier, attribution required in production
PEXELS_API_KEY=...           # api.pexels.com — free tier, no attribution required
REPLICATE_API_TOKEN=...      # for FLUX.1 [dev]/[pro] and other image models
FAL_KEY=...                  # alternative to Replicate, lower latency
```

The agent must check which of these are present before choosing route 2 or 3. If the matching key is missing, degrade to a placeholder and tell the user explicitly — never hide it.

Gentle-AI does not make these calls for you: it is a configurator, not an HTTP runtime. `gentle-ai doctor` only reports which keys are present in the environment; the API call is made by the agent from this skill.

## Placeholder policy

- Use a deterministic, stable service (e.g. `picsum.photos/seed/<seed>/<w>/<h>`), never broken images or URLs invented from memory.
- Flag it in the delivery copy ("these photos are placeholders, replace before publishing") — never leave it implicit.
- Do not use placeholders for logos, real product screenshots, or photos of identifiable people from the brand — there, always request the real asset or state it explicitly as pending.

## Licensing and ethics rules

- Never generate or use photos that appear to depict a real, identifiable person.
- Respect attribution when the API requires it (Unsplash does, Pexels does not).
- Do not mix generic stock photos into the same layout as the client's real data or figures in a way that reads as evidence (e.g. do not put a random "call center" photo next to the client's real metrics, implying it is their team).
