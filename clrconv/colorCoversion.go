package clrconv

import (
	"errors"
	"image/color"
	"strings"
)

func NoteToColor(note string) string {
	switch note {
	case "A": // A
		return "Red" // Red
	case "B": // B
		return "Orange" // Orange
	case "C": // C
		return "yellow" // Yellow
	case "D": // D
		return "green" // Green
	case "E": // E
		return "blue" // Blue
	case "F": // F
		return "indigo" // Indigo
	case "G": // G
		return "violet" // Violet
	default:
		return ""
	}
}

func GetRGBAFromReadableColor(name string) (color.RGBA, error) {
	name = strings.ToLower(name)

	colors := map[string]color.RGBA{
		"red":    {R: 255, G: 0, B: 0, A: 255},
		"orange": {R: 255, G: 165, B: 0, A: 255},
		"yellow": {R: 255, G: 255, B: 0, A: 255},
		"green":  {R: 0, G: 255, B: 0, A: 255},
		"blue":   {R: 0, G: 0, B: 255, A: 255},
		"indigo": {R: 75, G: 0, B: 130, A: 255},
		"violet": {R: 238, G: 130, B: 238, A: 255},
		"black":  {R: 0, G: 0, B: 0, A: 255},
		// add more if needed for extended support
	}

	if rgba, ok := colors[name]; ok {
		return rgba, nil
	}

	return color.RGBA{}, errors.New("color name not recognized")
}
