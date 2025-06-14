# tts_engine/stub_engine.py
# This module plays voice strings from a YAML tutorial manifest using pyttsx3 (offline TTS engine)

import sys
import yaml
import pyttsx3

# Function: Convert a single text string to speech using pyttsx3
def play_voice(text: str):
    engine = pyttsx3.init()        # Initialize the TTS engine
    engine.say(text)               # Queue the text to be spoken
    engine.runAndWait()            # Block until all queued commands are finished

# Function: Load YAML manifest and play each slide's voice string
def run_voice_engine(manifest_file):
    with open(manifest_file, 'r') as f:
        data = yaml.safe_load(f)   # Load YAML content as Python list/dict

    for slide in data:
        voice = slide.get('voice') # Extract 'voice' string from each slide (if present)
        if voice:
            play_voice(voice)      # Play the voice using TTS engine

# Entry point: Parse CLI args and run engine
if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python stub_engine.py <manifest.yaml>")
        sys.exit(1)

    run_voice_engine(sys.argv[1])
