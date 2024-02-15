package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getFile(filename string) (io.ReadCloser, error) {
	stdinStat, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if stdinStat.Mode()&os.ModeCharDevice == 0 {
		return os.Stdin, nil
	}

	return os.OpenFile(filename, os.O_RDONLY, 0)
}

func main() {
	chars := flag.Bool("m", false, "count the chars from the file")
	lines := flag.Bool("l", false, "count the lines from the file")
	bytes := flag.Bool("c", false, "count the bytes from the file")
	words := flag.Bool("w", false, "count the words from the file")

	flag.Parse()

	filePath := flag.Arg(0)

	fileReader, err := getFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer fileReader.Close()

	stats, err := Count(fileReader)
	if err != nil {
		log.Fatal(err)
	}

	if !*chars && !*lines && !*bytes && !*words {
		*words = true
		*lines = true
		*bytes = true
	}

	var statsPrint []string
	statsPrint = append(statsPrint, stats.GetCountersAsStringSlice(*lines, *words, *chars, *bytes)...)

	fmt.Printf("    %s %s\n", strings.Join(statsPrint, "  "), filePath)
}
