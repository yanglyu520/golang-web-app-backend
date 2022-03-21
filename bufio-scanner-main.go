package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "sss dfdfa dsf"
	//strings.NewReader will take a string and return a pointer to a reader
	scanner := bufio.NewScanner(strings.NewReader(s))

	//Scanner provides a convenient interface for reading data such as a file of newline-delimited lines of text. Successive calls to the Scan method will step through the 'tokens' of a file, skipping the bytes between the tokens. The specification of a token is defined by a split function of type SplitFunc;
	scanner.Split(bufio.ScanWords)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
