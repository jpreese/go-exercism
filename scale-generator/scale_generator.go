package scale

import (
	"container/ring"
	"strings"
)

var keys = ring.New(12)
var keyTranslationMap = map[string] string {
	"C#": "Db",
	"D#": "Eb",
	"F#": "Gb",
	"G#": "Ab",
	"A#": "Bb",
	"Db": "C#",
	"Eb": "D#",
	"Gb": "F#",
	"Ab": "G#",
	"Bb": "A#",
}

func Scale(tonic string, interval string) []string{

	initializeKeyboard()
	setStartingKey(tonic)

	renderFlatNotes := shouldRenderFlatNotes(tonic)

	keyCount := 12

	if interval != "" {
		keyCount = len(interval)
	}

	scale := make([]string, keyCount)
	for i := 0; i < keyCount; i++ {

		currentNote := keys.Value.(string)

		if isSharpNote(currentNote) && renderFlatNotes {
			scale[i] = keyTranslationMap[keys.Value.(string)]
		} else {
			scale[i] = keys.Value.(string)
		}

		if interval != "" && interval[i] == 'M' {
			keys = keys.Move(2)
		} else if interval != "" && interval[i] == 'A' {
			keys = keys.Move(3)
		} else {
			keys = keys.Next()
		}
	}

	return scale
}

func initializeKeyboard() {
	keyboardKeys := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

	for i := 0; i < keys.Len(); i++ {
		keys.Value = keyboardKeys[i]
		keys = keys.Next()
	}
}

func setStartingKey(tonic string)  {

	for strings.EqualFold(keys.Value.(string), tonic) == false && strings.EqualFold(keyTranslationMap[keys.Value.(string)], tonic) == false {
		keys = keys.Next()
	}
}

func shouldRenderFlatNotes(tonic string) bool {
	const FlatTonics = "Fdgcf"

	if len(tonic) > 1 && tonic[1] == 'b' {
		return true
	}

	if strings.Index(FlatTonics, tonic) > -1 {
		return true
	}

	return false
}

func isSharpNote(note string) bool {
	const Sharp = "#"
	if strings.Index(note, Sharp) > -1 {
		return true
	}

	return false
}
