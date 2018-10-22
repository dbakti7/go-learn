package common

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadWrite(t *testing.T) {
	lines := []string{"line1", "line2", "abc"}
	err := WriteLines(lines, "test.txt")
	defer os.Remove("test.txt")
	assert.Equal(t, nil, err, "There should not be error on Write Operation")

	readLines, err := ReadLines("test.txt")
	assert.Equal(t, nil, err, "There should not be error on Read Operation")
	assert.Equal(t, lines, readLines, "Wrong file content")

	assert.Equal(t, nil, err, "There should not be error on Cleanup Operation")
}
