# 🧭 DESIGN: Who Does What?

This document explains the architectural separation between the `tutorial-generator` project and frontend clients like `dasmlab-home`.

---

## 🧱 tutorial-generator

### Purpose:
Builds **self-contained animated tutorials** from structured manifests.

### Responsibilities:
- Defines the animation manifest schema (YAML/JSON)
- Parses tutorial slides and animation steps
- Renders tutorials using a JavaScript/Canvas engine
- Supports:
  - `Box`, `Arrow`, `Callout`, `FadeIn`, `MoveTo`
  - Timed timeline execution (`at`, `duration`)
  - Voice string capture
  - Subtitle export to `.vtt` via Go utility
- Stubs out CLI (e.g. `render_local.sh`) and TTS support

### Output:
A standalone HTML/JS+Canvas engine that plays animations driven by a manifest.

---

## 🖥️ dasmlab-home

### Purpose:
**Displays and hosts tutorials** in the UI/UX context (e.g., user clicks “Tutorial” on a project card).

### Responsibilities:
- Handles routing and triggering tutorial display
- Opens tutorials in modals or full-screen embedded canvas
- Supplies the correct manifest (`sample_flux_demo.yaml`, etc.)
- Integrates branding, watermark, and user controls (optional)
- Optionally uses subtitles or audio from generated `.vtt`

### How it interacts:
- Embeds the animation player (iframe, JS module, or static asset)
- Passes the manifest URL or data blob
- Does **not** manage rendering logic — that stays in `tutorial-generator`

---

## 🔁 Why this Separation?

Keeping rendering logic in `tutorial-generator` ensures:
- Reusability across frontend apps
- Independent testing and export
- Clear ownership boundaries

The frontend (`dasmlab-home`) focuses purely on **delivery and experience**.

---

✅ **Next**: Continue embedding the generated player into frontend UIs and polish playback UX.

