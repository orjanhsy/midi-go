package clrconv

import (
	"errors"
	"image/color"
	"strings"
)

func GetRGBAFromNote(note string) (color.RGBA, error) {
	note = strings.ToLower(note)

	colors := map[string]color.RGBA{
		"a": {R: 255, G: 0, B: 0, A: 255},
		"b": {R: 255, G: 165, B: 0, A: 255},
		"c": {R: 255, G: 255, B: 0, A: 255},
		"d": {R: 0, G: 255, B: 0, A: 255},
		"e": {R: 0, G: 0, B: 255, A: 255},
		"f": {R: 75, G: 0, B: 130, A: 255},
		"g": {R: 238, G: 130, B: 238, A: 255},
		"":  {R: 0, G: 0, B: 0, A: 255},
	}
	if rgba, ok := colors[note]; ok {
		return rgba, nil
	}

	return color.RGBA{}, errors.New("color name not recognized")
}
