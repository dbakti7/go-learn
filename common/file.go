package common

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadLines reads a file and returns a slice of strings of its content.
func ReadLines(path string) (lines []string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimRight(scanner.Text(), "\n"))
	}
	return
}

// WriteLines writes strings into a file.
// The file will be created if it doesn't exist, else it will be overwritten.
func WriteLines(lines []string, path string) (err error) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	err = writer.Flush()
	return
}
