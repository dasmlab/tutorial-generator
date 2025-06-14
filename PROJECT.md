# dasmlab_tuto_generator – Project Overview & Phased Plan

This document outlines the development lifecycle of the `dasmlab_tuto_generator`, a framework to programmatically generate animated, narrated tutorials for `dasmlab` projects. These tutorials will serve as interactive explainers and lightweight demos embedded across project landing pages.

---

## 🎯 Vision

Enable consistent, branded educational flows for each `dasmlab` project via a declarative system that supports:
- Visual animations (boxes, links, transitions)
- Step-by-step concepts tied to code or deployment flows
- Text-to-speech (TTS) voice-over narration
- Reusable branded player embedded in web environments

---

## 📌 Key Development Phases

### 🧱 Phase 1: Core Architecture & Input Format
- Define tutorial manifest structure (YAML/JSON)
- Slide schema includes:
  - `title`, `description`, `animation`, `voice`
- Animation primitives: `box`, `arrow`, `fadeIn`, `moveTo`

### 🧩 Phase 2: Animation Renderer
- Choose SVG + D3.js for portability + simplicity
- Build a Quasar/Vue-based canvas engine that:
  - Renders from input YAML
  - Supports sequencing and animation types
- Add pause/play + slide navigation controls

### 🎙️ Phase 3: Narration Engine
- Integrate subtitle renderer
- Add optional TTS voice-over (OpenAI, Coqui, gTTS)
- Allow override of voice files per slide (mp3)

### 📼 Phase 4: Branded Tutorial Player
- Wrap animations in a branded player (Dasmlab watermark)
- Allow autoplay, theme switching (light/dark), and embedding
- Add tutorial metadata + credits footer

### 🔁 Phase 5: Authoring CLI Tools
- Generate tutorial stubs (`tuto init`, `tuto render`)
- Validate slide manifests
- Compile to static HTML bundle

### 🌍 Phase 6: Hosting & Embedding
- Export player as embeddable HTML5 widget
- Auto-deploy to GitHub Pages for preview
- Generate social sharing links

---

## 📁 Example Structure

dasmlab_tuto_generator/
   	|  
	|- tutorials/
	|	|- fluxcd_gitops.yaml
        |
	|- renderer/ # D3/SVG slide engine
        |
	|- player/ # Vue/Quasar frontend
        |
	|- tts_engine/ # TTS backend integration
        |
	|- scripts/ # CLI tools for building
        |
	|- assets/ # Shared media (icons, audio, font)
        |
	|- PROJECT.md
        |
	|- README.md

---

## 🛠️ Future Extensions
- Support keyboard navigation
- Multi-language subtitle support
- AI-generated diagrams from textual descriptions

---

## 📅 Current Status

> We are currently in **Phase 1: Manifest + Renderer Setup**
- **GOAL** - Render a basic manifest-driven tutorial in a browser - just static shapes, arroes and a timeline-sequenced movement based on our YAML


See [README.md](./README.md) for usage instructions.
