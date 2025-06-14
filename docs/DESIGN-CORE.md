# ✨ Core Design Spec: dasmlab_tuto_generator

This document outlines the foundational design of the `dasmlab_tuto_generator` framework, which enables programmatically defined tutorial animations, narration, and transitions across all DASMLAB projects.

---

## 🧱 Overview

This system allows DASMLAB projects to include animated, narrated tutorials that play within a branded, consistent interface. Tutorials are driven by a YAML/JSON manifest and rendered into an animation timeline with voice overlay support.

---

## 🧠 Concepts

- **Manifest-Driven**: Each tutorial is defined declaratively via YAML or JSON.
- **Structured Slides**: Tutorials are made up of sequential slides.
- **Animation Primitives**: Standard actions like boxes, arrows, callouts define the flow.
- **Timing & Voice**: Timed or user-triggered steps with voice narration support.

---

## 📐 Slide Manifest Structure

```yaml
slides:
  - title: "Slide Title"
    description: "Short description of this step"
    voice: "Optional TTS or narration text"
    animations:
      - type: box
        id: component1
        label: "Git Repo"
        x: 100
        y: 100
        fadeIn: true
      - type: arrow
        from: component2
        to: component1
        at: 2.0
        duration: 1.5
🧩 Animation Primitives
Primitive	Purpose
box	Create a labeled rectangle
arrow	Draw an arrow from one component to another
moveTo	Move an existing element to a new location
fadeIn	Fade an element into view
fadeOut	Fade an element out
highlight	Highlight or pulse an element
callout	Show floating annotation text box
group	Group multiple items for batch operations

⏱️ Animation Timing & Control
Each animation can optionally include:

Field	Meaning
at	When the animation begins (in seconds)
duration	How long the animation runs
pauseAfter	Pause duration after this step
trigger	Optional trigger (e.g., onClick, onVoiceEnd)

🎙️ Voice/Narration
voice: field per slide (plain text)

Used to auto-generate voiceover via TTS

Timeline syncs with animation pacing

📦 Schema Fields Reference
Field	Type	Description
type	string	One of the animation primitives
id	string	Unique identifier for reference
x, y	number	Position on screen
label	string	Optional visible label
from, to	string	Used for arrow type animations
attachedTo	string	Used for callout or highlight
style	string	Optional styling directive

🔖 Example: "FluxCD Polls Git"
yaml
Copy
Edit
- title: "FluxCD Watches Git"
  description: "FluxCD continuously polls the Git repository."
  voice: "FluxCD monitors Git for changes and reacts automatically."
  animations:
    - type: box
      id: repo
      label: "Git Repo"
      x: 100
      y: 100
      fadeIn: true
    - type: box
      id: flux
      label: "FluxCD"
      x: 300
      y: 100
      fadeIn: true
    - type: arrow
      from: flux
      to: repo
      style: dashed
      at: 2.0
    - type: callout
      text: "Polling every 1m"
      attachedTo: flux
