package nameWriter

import (
	FontColors "go-server/NameWriter/Colors"
	"maps"
	"strings"
)

func WriteName(name string, style string, color string) string {
	if len(style) == 0 {
		style = "BIG"
	}
	if len(color) == 0 {
		color = FontColors.FontColors["RESET"]
	}

	var st strings.Builder
	currentLine := 0
	currentAlphabet := alphabets[strings.ToUpper(style)]
	iterations := len(currentAlphabet["A"])

	for i := 0; i < iterations; i++ {
		for _, letter := range name {

			st.WriteString(
				FontColors.FontColors[strings.ToUpper(color)] +
					currentAlphabet[strings.ToUpper(string(letter))][currentLine] +
					"   ")
		}
		currentLine++
		st.WriteString("\n")
	}
	return st.String()
}

func GetFonts() string {
	var sb strings.Builder
	fonts := maps.Keys(alphabets)

	for font := range fonts {
		sb.WriteString("Fonte: " + font + " \n\n")
		sb.WriteString(WriteName("A", font, ""))
		sb.WriteString("\n\n\n")
	}
	return sb.String()
}
