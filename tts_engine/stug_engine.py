# tts_engine/stub_engine.py
import sys
import yaml

def run_voice_engine(manifest_file):
    with open(manifest_file, 'r') as f:
        data = yaml.safe_load(f)

    for slide in data:
        voice = slide.get('voice')
        if voice:
            print(f"[🔈] Voice: {voice}")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python stub_engine.py <manifest.yaml>")
        sys.exit(1)
    run_voice_engine(sys.argv[1])

