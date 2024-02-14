package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func countFileChars(file io.Reader) (int, error) {

	b := bufio.NewReader(file)
	defer b.Reset(file)

	count := 0

	for {
		_, _, err := b.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, fmt.Errorf("error reading line")
		}

		count += 1
	}

	return count, nil
}

func countFileLines(file io.Reader) (int, error) {
	b := bufio.NewReader(file)
	defer b.Reset(file)

	count := 0
	for {
		_, _, err := b.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, fmt.Errorf("error reading line")
		}

		count += 1
	}

	return count, nil
}

func countFileBytes(file io.Reader) (int, error) {
	b := bufio.NewReader(file)
	defer b.Reset(file)

	count := 0
	for {
		bytes, err := b.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, fmt.Errorf("error reading file: %v", err)
		}

		count += len(bytes)
	}

	return count, nil
}

func countFileWords(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
	}

	return count, nil
}

func openFile(filePath string) (io.ReadCloser, error) {
	return os.OpenFile(filePath, os.O_RDONLY, 0)
}

type Counter struct {
	flag    bool
	countFn func(f io.Reader) (int, error)
}

func newCounter(flag bool, countFn func(f io.Reader) (int, error)) *Counter {
	return &Counter{
		flag,
		countFn,
	}
}

func main() {
	countBytes := flag.Bool("c", false, "Print the number of bytes")
	countLines := flag.Bool("l", false, "Print the number of lines")
	countWords := flag.Bool("w", false, "Print the number of words")
	countChars := flag.Bool("m", false, "Print the number of characters")

	flag.Parse()

	noFlags := !*countBytes && !*countLines && !*countWords && !*countChars

	args := os.Args[1:]
	if !noFlags {
		args = os.Args[2:]
	}

	stdinFileStat, errStdin := os.Stdin.Stat()

	if len(args) == 0 && stdinFileStat.Size() == 0 {
		panic(fmt.Errorf("no files to count"))
	}

	filePath := ""
	if len(args) > 0 {
		filePath = args[0]
	}
	file, err := openFile(filePath)
	if err != nil && (errStdin != nil || stdinFileStat.Size() == 0) {
		panic(err)
	}

	if stdinFileStat.Size() > 0 {
		file = os.Stdin
	}

	var buf bytes.Buffer
	fileReader := io.TeeReader(file, &buf)

	flags := [4]*Counter{
		newCounter(*countLines || noFlags, countFileLines),
		newCounter(*countWords || noFlags, countFileWords),
		newCounter(*countBytes || noFlags, countFileBytes),
		newCounter(*countChars, countFileChars),
	}

	results := make([]int, 0)

	for _, flag := range flags {
		if !flag.flag {
			continue
		}

		result, err := flag.countFn(fileReader)
		if err != nil {
			panic(err)
		}

		results = append(results, result)

		fileReader, err = os.OpenFile(filePath, os.O_RDONLY, 0)
		if err != nil {
			panic(err)
		}
	}

	var numbers []string

	for _, n := range results {
		numbers = append(numbers, strconv.Itoa(n))
	}

	fmt.Printf("  %s %s\n", strings.Join(numbers, "  "), filePath)
}
