# dasmlab_tuto_generator

🎓 A programmatic tutorial generator for `dasmlab` projects, blending animation, narration, and branding — all defined in YAML.

> Each tutorial helps explain complex devops, GitOps, or AI workflows in a visual, animated, and reusable way.

---

## 📦 What This Project Does

- Defines structured tutorials using YAML
- Renders animated slides via SVG + D3
- Narrates concepts with TTS voice-over (optional)
- Exports tutorials to an embeddable branded player

---

## 🧪 Current Status

- ✅ Tutorial manifest format defined (`YAML`)
- ⏳ Renderer scaffold (in progress)
- ⏳ First tutorial: `fluxcd_gitops.yaml` in development

---

## 🚀 Usage

### 1. Define a Tutorial

```yaml
title: "FluxCD GitOps Flow"
slides:
  - title: "Step 1: Git Commit"
    animation:
      - type: box
        id: github
        label: GitHub Repo
        pos: [100, 100]
      - type: arrow
        from: github
        to: runner
    narration: "Your GitHub repo triggers GitHub Actions"
```

### 2. Run the Generator (coming soon)
./scripts/render_tuto.sh tutorials/fluxcd_gitops.yaml

# 🛠 Requirements
NodeJS + Yarn (for player UI)

Python (optional for voice generation)

Git (for fetching templates)

# 📚 Related Projects
dasmlab-home – Public homepage

dasmlab-live-cicd – GitOps delivery backbone

# 🧭 Roadmap
See full phase breakdown in PROJECT.md.

# ⚖️ License
Apache License 2.0 with Attribution (AAL format) - See licenses/DASMLAB-LICENSE.md


codign something fo an exampple:
