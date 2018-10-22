package scale

import (
	"container/ring"
	"strings"
)

var keyTranslationMap = map[string]string{
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

const (
	// WholeStep is an interval between two notes
	WholeStep byte = 'M'

	// AugmentedFirst is an interval with two interceding notes
	AugmentedFirst byte = 'A'
)

// Scale takes a tonic and an interval to return a collection of notes
func Scale(tonic string, interval string) []string {

	keyboard := initializeKeyboard(tonic)
	renderFlatNotes := shouldRenderFlatNotes(tonic)
	notesToReturn := keyboard.Len()

	if interval != "" {
		notesToReturn = len(interval)
	}

	scale := make([]string, notesToReturn)
	for i := 0; i < notesToReturn; i++ {

		currentNote := keyboard.Value.(string)

		if isSharpNote(currentNote) && renderFlatNotes {
			scale[i] = keyTranslationMap[keyboard.Value.(string)]
		} else {
			scale[i] = keyboard.Value.(string)
		}

		if interval != "" && interval[i] == AugmentedFirst {
			keyboard = keyboard.Move(3)
		} else if interval != "" && interval[i] == WholeStep {
			keyboard = keyboard.Move(2)
		} else {
			keyboard = keyboard.Next()
		}
	}

	return scale
}

func initializeKeyboard(tonic string) (keyboard *ring.Ring) {
	keyboard = ring.New(12)
	keyboardKeys := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

	for _, value := range keyboardKeys {
		keyboard.Value = value
		keyboard = keyboard.Next()
	}

	for strings.EqualFold(keyboard.Value.(string), tonic) == false && strings.EqualFold(keyTranslationMap[keyboard.Value.(string)], tonic) == false {
		keyboard = keyboard.Next()
	}

	return
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
