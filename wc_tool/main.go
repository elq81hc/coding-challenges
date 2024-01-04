package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	var byteCount, lineCount, charCount, wordCount bool
	flag.BoolVar(&byteCount, "c", false, "The number of bytes in each input file is written to the standard output")
	flag.BoolVar(&lineCount, "l", false, "The number of lines in each input file is written to the standard output")
	flag.BoolVar(&charCount, "m", false, "The number of characters in each input file is written to the standard output")
	flag.BoolVar(&wordCount, "w", false, "The number of words in each input file is written to the standard output")
	flag.Parse()

	var readFrom io.Reader
	var fileName string

	if fileName = flag.Arg(0); fileName != "" {
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
		defer f.Close()
		readFrom = f
	} else {
		readFrom = os.Stdin
	}

	nLines, nWords, nChars, nBytes, err := wc(readFrom)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	var output strings.Builder
	showAll := !byteCount && !lineCount && !wordCount && !charCount
	if showAll || lineCount {
		output.WriteString(fmt.Sprintf("  %d", nLines))
	}
	if showAll || wordCount {
		output.WriteString(fmt.Sprintf("  %d", nWords))
	}
	if showAll || byteCount {
		output.WriteString(fmt.Sprintf("  %d", nBytes))
	}
	if charCount {
		output.WriteString(fmt.Sprintf("  %d", nChars))
	}
	if len(fileName) > 0 {
		output.WriteString(fmt.Sprintf("  %s", fileName))
	}
	fmt.Println(output.String())
}

func wc(rd io.Reader) (nLines, nWords, nChars, nBytes uint64, err error) {
	inWord := false
	nLines, nWords, nChars, nBytes = 0, 0, 0, 0
	reader := bufio.NewReader(rd)
	r, sz, err := reader.ReadRune()
	for ; err == nil; r, sz, err = reader.ReadRune() {
		if unicode.IsSpace(r) {
			if r == '\n' {
				nLines++
			}
			if inWord {
				nWords++
				inWord = false
			}
		} else {
			inWord = true
		}
		nChars++
		nBytes += uint64(sz)
	}
	if err != io.EOF {
		return 0, 0, 0, 0, err
	}
	if inWord {
		nWords++
	}
	return nLines, nWords, nChars, nBytes, nil
}
