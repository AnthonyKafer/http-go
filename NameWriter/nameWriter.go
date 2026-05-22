package nameWriter

import "strings"

func WriteName(name string, style string) string {
	if len(style) == 0 {
		style = "BIG"
	}

	var st strings.Builder
	currentLine := 0
	currentAlphabet := alphabets[strings.ToUpper(style)]
	iterations := len(currentAlphabet["A"])

	for i := 0; i < iterations; i++ {
		for _, letter := range name {
			st.WriteString(currentAlphabet[strings.ToUpper(string(letter))][currentLine] + "   ")
		}
		currentLine++
		st.WriteString("\n")
	}
	return st.String()
}
