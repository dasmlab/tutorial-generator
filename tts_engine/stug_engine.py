import sys
import logging

logging.basicConfig(level=logging.INFO)

def speak(text: str):
    """Stub TTS engine that logs the voice string."""
    logging.info(f"[VOICE] {text}")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: stub_engine.py '<text>'")
        sys.exit(1)
    speak(sys.argv[1])

