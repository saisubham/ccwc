package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	lineFlag := flag.Bool("l", false, "count the number of lines in input file")
	byteFlag := flag.Bool("c", false, "count the number of bytes in input file")
	wordFlag := flag.Bool("w", false, "count the number of words in input file")
	flag.Parse()

	paths := flag.Args()

	if len(paths) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		res := countUtil(scanner)
		displayResult(lineFlag, byteFlag, wordFlag, res)
	} else {
		for _, path := range paths {
			res := count(path)
			displayResult(lineFlag, byteFlag, wordFlag, res)
		}
	}
}

func displayResult(lineFlag *bool, byteFlag *bool, wordFlag *bool, res []int) {
	ans := "\t"
	allFlag := false
	if !*lineFlag && !*byteFlag && !*wordFlag {
		allFlag = true
	}
	if *lineFlag || allFlag {
		ans += fmt.Sprintf("%d\t", res[0])
	}
	if *byteFlag || allFlag {
		ans += fmt.Sprintf("%d\t", res[1])
	}
	if *wordFlag || allFlag {
		ans += fmt.Sprintf("%d\t", res[2])
	}
	fmt.Println(ans, "stdin")
}

func count(path string) []int {
	file, err := os.Open(path)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	return countUtil(scanner)
}

func countUtil(scanner *bufio.Scanner) []int {
	lineCount := 0
	wordCount := 0
	byteCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		lineCount++
		wordCount += len(words)
		byteCount += len(line)
	}
	return []int{lineCount, wordCount, byteCount}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
