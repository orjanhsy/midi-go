package clrconv

import (
	"errors"
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2/data/binding"
)

func NoteToColor(note string) string {
	switch note {
	case "A": // A
		return "#FF0000" // Red
	case "B": // B
		return "#FFA500" // Orange
	case "C": // C
		return "#FFFF00" // Yellow
	case "D": // D
		return "#00FF00" // Green
	case "E": // E
		return "#0000FF" // Blue
	case "F": // F
		return "#4B0082" // Indigo
	case "G": // G
		return "#EE82EE" // Violet
	default:
		return ""
	}
}

// Map note letters (A-G) to a visible color
func NoteToRGBColor(note string) color.Color {
	switch note {
	case "A":
		return color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red
	case "B":
		return color.RGBA{R: 255, G: 165, B: 0, A: 255} // Orange
	case "Bb":
		return color.RGBA{R: 255, G: 165, B: 0, A: 255} // Orange
	case "C":
		return color.RGBA{R: 255, G: 255, B: 0, A: 255} // Yellow
	case "D":
		return color.RGBA{R: 0, G: 255, B: 0, A: 255} // Green
	case "E":
		return color.RGBA{R: 0, G: 0, B: 255, A: 255} // Blue
	case "F":
		return color.RGBA{R: 75, G: 0, B: 130, A: 255} // Indigo
	case "G":
		return color.RGBA{R: 238, G: 130, B: 238, A: 255} // Violet
	default:
		return color.Black
	}
}

func GetReadableColorFromHexa(hexString string) (string, error) {
	colors := map[string]string{
		"#FF0000": "Red",
		"#00FF00": "Green",
		"#0000FF": "Blue",
		"#FFFF00": "Yellow",
		"#00FFFF": "Cyan",
		"#FF00FF": "Magenta",
		"#FFFFFF": "White",
		"#000000": "Black",
		"#808080": "Gray",
		"#800000": "Maroon",
		"#808000": "Olive",
		"#008000": "Dark Green",
		"#800080": "Purple",
		"#FFC0CB": "Pink",
		"#A52A2A": "Brown",
		"#FFD700": "Gold",
		"#F0F8FF": "Alice Blue",
		"#FAEBD7": "Antique White",
		"#EE82EE": "Violet",
	}

	color, ok := colors[hexString]

	if ok {
		return color, nil
	}
	return "", errors.New("color not found")
}

func GetReadableColorFromRGB(c color.Color) (string, error) {
	r, g, b, _ := c.RGBA()

	// Convert from 16-bit (0–65535) to 8-bit (0–255)
	r8 := uint8(r >> 8)
	g8 := uint8(g >> 8)
	b8 := uint8(b >> 8)

	colors := map[[3]uint8]string{
		{255, 0, 0}:     "Red",
		{0, 255, 0}:     "Green",
		{0, 0, 255}:     "Blue",
		{255, 255, 0}:   "Yellow",
		{0, 255, 255}:   "Cyan",
		{255, 0, 255}:   "Magenta",
		{255, 255, 255}: "White",
		{0, 0, 0}:       "Black",
		{128, 128, 128}: "Gray",
		{128, 0, 0}:     "Maroon",
		{128, 128, 0}:   "Olive",
		{0, 128, 0}:     "Dark Green",
		{128, 0, 128}:   "Purple",
		{255, 192, 203}: "Pink",
		{165, 42, 42}:   "Brown",
		{255, 215, 0}:   "Gold",
		{240, 248, 255}: "Alice Blue",
		{250, 235, 215}: "Antique White",
		{238, 130, 238}: "Violet",
		{255, 165, 0}:   "Orange",
		{75, 0, 130}:    "Indigo",
	}

	if name, ok := colors[[3]uint8{r8, g8, b8}]; ok {
		return name, nil
	}
	return "", errors.New("color not found")
}

func GetRGBAFromReadableColor(clr binding.String) (color.RGBA, error) {
	name, err := clr.Get()
	if err != nil {
		fmt.Println("Failed to get string value of binding")
		return color.RGBA{}, err
	}
	name = strings.ToLower(name)

	colors := map[string]color.RGBA{
		"red":    {R: 255, G: 0, B: 0, A: 255},
		"orange": {R: 255, G: 165, B: 0, A: 255},
		"yellow": {R: 255, G: 255, B: 0, A: 255},
		"green":  {R: 0, G: 255, B: 0, A: 255},
		"blue":   {R: 0, G: 0, B: 255, A: 255},
		"indigo": {R: 75, G: 0, B: 130, A: 255},
		"violet": {R: 238, G: 130, B: 238, A: 255},
		// add more if needed for extended support
	}

	if rgba, ok := colors[name]; ok {
		return rgba, nil
	}

	return color.RGBA{}, errors.New("color name not recognized")
}
