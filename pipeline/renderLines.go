package pipeline

import "strings"

func RenderLines(tokens []string, banner map[string][]string) []string {
	var out []string
	current := make([]string, 8)

	hasContent := false
	lastWasNewline := false

	flush := func() {
		if hasContent {
			out = append(out, current...)
			current = make([]string, 8)
			hasContent = false
		}
	}

	for _, tok := range tokens {
		if tok == "\n" {
			flush()

			// αν ήταν ήδη newline, προσθέτουμε ΜΙΑ κενή γραμμή
			if lastWasNewline {
				out = append(out, "")
			}

			lastWasNewline = true
			continue
		}

		lastWasNewline = false

		glyph, ok := banner[tok]
		if !ok {
			pad := strings.Repeat(" ", 4)
			for i := 0; i < 8; i++ {
				current[i] += pad
			}
			hasContent = true
			continue
		}

		if len(glyph) < 8 {
			tmp := make([]string, 8)
			copy(tmp, glyph)
			glyph = tmp
		}

		for i := 0; i < 8; i++ {
			current[i] += glyph[i]
		}

		hasContent = true
	}

	flush()
	return out
}
