package main

import (
	"bufio"
	"io"
	"strconv"
	"unicode"
)

type FileStats struct {
	LineCount  int64
	WordCount  int64
	BytesCount int64
	CharCount  int64
}

func (fileStats FileStats) GetCountersAsStringSlice(lineCount, wordCount, charsCount, bytesCount bool) []string {
	stats := [4]int64{fileStats.LineCount, fileStats.WordCount, fileStats.CharCount, fileStats.BytesCount}
	var statsString []string

	for idx, flag := range [4]bool{lineCount, wordCount, charsCount, bytesCount} {
		if !flag {
			continue
		}

		statsString = append(statsString, strconv.FormatInt(stats[idx], 10))
	}

	return statsString
}

func Count(input io.Reader) (*FileStats, error) {
	//defer input.Close()

	fileStats := &FileStats{}

	r := bufio.NewReader(input)

	inWord := false
	for {
		char, size, err := r.ReadRune()
		if err == io.EOF {
			if inWord {
				fileStats.LineCount += 1
			}
			break
		}

		if err != nil {
			return nil, err
		}

		fileStats.CharCount += 1
		fileStats.BytesCount += int64(size)

		if unicode.IsSpace(char) {
			if inWord {
				fileStats.WordCount += 1
			}

			if string(char) == "\n" {
				fileStats.LineCount += 1
			}

			inWord = false
			continue
		}

		inWord = true

	}

	return fileStats, nil
}
