package checker

import (
	"strings"

	"github.com/dbakti7/go-learn/common"
)

func StrictCheck(expectedPath string, actualPath string) (accepted bool, err error) {
	expected, err := common.ReadLines(expectedPath)
	if err != nil {
		return
	}

	actual, err := common.ReadLines(actualPath)
	if err != nil {
		return
	}

	if len(expected) != len(actual) {
		return
	}

	for i := range expected {
		// strings retrieved from Read function might contain newline character
		if expected[i] != actual[i] {
			return
		}
	}
	accepted = true
	return
}

func LenientCheck(expectedPath string, actualPath string) (accepted bool, err error) {
	expected, err := common.ReadLines(expectedPath)
	if err != nil {
		return
	}

	actual, err := common.ReadLines(actualPath)
	if err != nil {
		return
	}

	preprocess := func(lines []string) []string {
		// remove leading and trailing white spaces
		for i := range lines {
			lines[i] = strings.TrimSpace(lines[i])
		}

		// remove trailing empty lines
		for i := len(lines) - 1; i >= 0; i-- {
			if lines[i] != "" {
				lines = lines[:i+1]
				break
			}
		}
		return lines
	}

	expected = preprocess(expected)
	actual = preprocess(actual)

	if len(expected) != len(actual) {
		return
	}

	for i := range expected {
		// strings retrieved from Read function might contain newline character
		if expected[i] != actual[i] {
			return
		}
	}
	accepted = true
	return
}
