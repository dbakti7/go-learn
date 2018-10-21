package checker

import (
	"os"
	"testing"

	"github.com/dbakti7/go-learn/common"

	"github.com/stretchr/testify/assert"
)

func TestStrictCheck(t *testing.T) {
	expected := []string{"aa", "bb", "", "123"}
	common.WriteLines(expected, "expected.txt")
	common.WriteLines(expected, "actualAccepted.txt")
	defer os.Remove("expected.txt")
	defer os.Remove("actualAccepted.txt")

	accepted, err := StrictCheck("expected.txt", "actualAccepted.txt")
	assert.Equal(t, nil, err, "There should not be error.")
	assert.Equal(t, true, accepted, "Should be accepted.")

	common.WriteLines(append(expected, ""), "actualWrongAnswer.txt")
	defer os.Remove("actualWrongAnswer.txt")

	accepted, err = StrictCheck("expected.txt", "actualWrongAnswer.txt")
	assert.Equal(t, nil, err, "There should not be error.")
	assert.Equal(t, false, accepted, "Extra line should not be accepted.")

	expected[0] = "ab"
	common.WriteLines(expected, "actualWrongAnswer.txt")
	accepted, err = StrictCheck("expected.txt", "actualWrongAnswer.txt")
	assert.Equal(t, nil, err, "There should not be error.")
	assert.Equal(t, false, accepted, "Wrong output should not be accepted")

	expected[0] = "aa "
	common.WriteLines(expected, "actualWrongAnswer.txt")
	accepted, err = StrictCheck("expected.txt", "actualWrongAnswer.txt")
	assert.Equal(t, nil, err, "There should not be error.")
	assert.Equal(t, false, accepted, "Trailing spaces should not be accepted")

	expected[0] = " aa"
	common.WriteLines(expected, "actualWrongAnswer.txt")
	accepted, err = StrictCheck("expected.txt", "actualWrongAnswer.txt")
	assert.Equal(t, nil, err, "There should not be error.")
	assert.Equal(t, false, accepted, "Leading spaces should not be accepted")
}

func TestLenientCheck(t *testing.T) {
	expected := []string{"aa", "bb", "", "123"}
	common.WriteLines(expected, "expected.txt")
	common.WriteLines(expected, "actualAccepted.txt")
	defer os.Remove("expected.txt")
	defer os.Remove("actualAccepted.txt")

	accepted, err := LenientCheck("expected.txt", "actualAccepted.txt")
	assert.Equal(t, nil, err, "There should not be error.")
	assert.Equal(t, true, accepted, "Should be accepted.")

	common.WriteLines(append(expected, "  ", "\n"), "actualAccepted.txt")
	accepted, err = LenientCheck("expected.txt", "actualAccepted.txt")
	assert.Equal(t, nil, err, "There should not be error.")
	assert.Equal(t, true, accepted, "Extra empty lines should be accepted.")

	expected[0] = "ab"
	common.WriteLines(expected, "actualWrongAnswer.txt")
	defer os.Remove("actualWrongAnswer.txt")
	accepted, err = LenientCheck("expected.txt", "actualWrongAnswer.txt")
	assert.Equal(t, nil, err, "There should not be error.")
	assert.Equal(t, false, accepted, "Wrong output should not be accepted")

	expected[0] = "aa "
	common.WriteLines(expected, "actualAccepted.txt")
	accepted, err = LenientCheck("expected.txt", "actualAccepted.txt")
	assert.Equal(t, nil, err, "There should not be error.")
	assert.Equal(t, true, accepted, "Trailing spaces should be accepted")

	expected[0] = " aa"
	common.WriteLines(expected, "actualAccepted.txt")
	accepted, err = LenientCheck("expected.txt", "actualAccepted.txt")
	assert.Equal(t, nil, err, "There should not be error.")
	assert.Equal(t, true, accepted, "Leading spaces should be accepted")
}
