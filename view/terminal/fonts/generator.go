package fonts

import "strings"

func Generate(f Font, str string) []string {
	font := fontMap[f]
	lines := make([]strings.Builder, font.height)
	for _, r := range str {
		c := font.letters[r]
		for i, cl := range c {
			lines[i].WriteString(cl)
		}
	}

	text := make([]string, len(lines))
	for i, l := range lines {
		text[i] = l.String()
	}

	return text
}
