# Exporter Tool (Subtitle Export)

This tool extracts the `voice` field from DASMLAB tutorial manifests and generates `.vtt` subtitle files for syncing narration.

## 🔗 Context

The exporter is part of the larger [`tutorial-generator`](../../) system that creates animated, narrated educational tutorials from structured YAML.

- It fits into the **post-processing** step of the toolchain.
- It parses YAML manifests from `tutorials/*.yaml`
- It outputs `.vtt` subtitle files used by the player or export-to-video pipelines.

## 🛠 Usage

```bash
cd tools/exporter
go run main.go ../../tutorials/sample_flux_demo.yaml > output.vtt
```

## 📂 File Layout

main.go – Entry point

parser/manifest.go – Parses YAML tutorials

writer/vtt.go – Generates valid .vtt subtitles

go.mod – Module definition

## 🧩 Related Modules

player/ – Runs the frontend visual timeline

tts_engine/ – Converts voice string into spoken audio
