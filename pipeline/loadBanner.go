package pipeline

import (
	"bufio"
	"fmt"
	"os"
)

func LoadBanner(name string) (map[string][]string, error) {
	file, err := os.Open("banners/" + name + ".txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	banner := make(map[string][]string)

	const (
		startChar = 32
		endChar   = 126
		height    = 8
		blockSize = height + 1 // ⬅️ ΚΕΝΗ ΓΡΑΜΜΗ + 8 γραμμές
	)

	expected := (endChar - startChar + 1) * blockSize
	if len(lines) < expected {
		return nil, fmt.Errorf("invalid banner file: wrong line count")
	}

	index := 0
	for c := startChar; c <= endChar; c++ {
		index++ // ⬅️ ΠΑΡΑΛΕΙΠΟΥΜΕ ΤΗΝ ΚΕΝΗ ΓΡΑΜΜΗ
		banner[string(rune(c))] = append(
			[]string(nil),
			lines[index:index+height]...,
		)
		index += height
	}

	return banner, nil
}
