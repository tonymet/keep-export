package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	stateSkip = iota
	stateStart
	stateHeading
	stateReading
)

var (
	patternHeading = regexp.MustCompile("class.*c5")
	patternStart   = regexp.MustCompile("<body class=\"c12\"")
)

func scanFile() error {
	var (
		docBuf   []string = make([]string, 0, 5)
		curState int      = stateSkip
		fileNum  int      = 0
		fileName string
	)

	//for m, i, err := 0, 0, error(nil); err != io.EOF; i++ {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if patternStart.Match([]byte(line)) {
			curState = stateStart
		} else if patternHeading.Match([]byte(line)) {
			curState = stateHeading
		}

		switch curState {
		case stateSkip:
		case stateStart:
			fmt.Printf("start: %s\n", line)
		case stateHeading:
			if len(docBuf) > 0 {
				// if buf is not empty, write out the file
				fileName = fmt.Sprintf("output/%04d.html", fileNum)
				if err := writeFile(docBuf, fileName); err != nil {
					panic(err)
				}
				fileNum++
			}
			// init buf
			docBuf = make([]string, 0, 5)
			// start appending
			docBuf = append(docBuf, line)
			fmt.Printf("heading: %s\n", line)
			// next state
			curState = stateReading
		case stateReading:
			docBuf = append(docBuf, line)
		}

	}
	return nil
}

func writeFile(docBuf []string, fileName string) error {
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	docBuf = append([]string{"<html>\n<body>\n"}, docBuf...)
	docBuf = append(docBuf, "</body>\n</html>\n")

	if _, err2 := f.WriteString(strings.Join(docBuf, "\n")); err2 != nil {
		panic(err2)
	}
	return nil
}

func main() {
	scanFile()
}
