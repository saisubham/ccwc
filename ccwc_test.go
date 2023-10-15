package main

import (
	"bufio"
	"os"
	"testing"
)

func TestCcwc(t *testing.T) {
	filename := "test.txt"
	expectedLineCount := 7145
	expectedByteCount := 58164
	expectedWordCount := 327900
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf(`Failed to read file %q`, filename)
	}
	scanner := bufio.NewScanner(file)
	res := countUtil(scanner)
	if res[0] != expectedLineCount || res[1] != expectedByteCount || res[2] != expectedWordCount {
		t.Fatalf(`expected [lines=%d, bytes=%d, words=%d] but found [lines=%d, bytes=%d, words=%d]`,
			expectedLineCount, expectedByteCount, expectedWordCount, res[0], res[1], res[2])
	}
}
