package confdir

import (
	"testing"

	"github.com/flimzy/testify/assert"
)

func TestConfDir(t *testing.T) {
	assert := assert.New(t)
	var testPath = "./t/a"
	var expected = []string{
		"Testing",
		"Testing",
		"1",
		"2",
		"3",
	}
	result, err := ReadConfDir(testPath)
	assert.DeepEqual(expected, result, "Expected output")
	assert.Nil(err, "No error")
}

func TestConfDir_MultipleFiles(t *testing.T) {
	assert := assert.New(t)
	var testPath = "./t/b"
	var expected = []string{
		"Baz",
		"Foo",
	}
	result, err := ReadConfDir(testPath)
	assert.DeepEqual(expected, result, "Expected output")
	assert.Nil(err, "No error")
}

func TestConfDir_DoesntExist(t *testing.T) {
	assert := assert.New(t)
	var testPath = "./t/c"
	var expected []string
	var expectedErr = "1 error occurred:\n\n* open ./t/c: no such file or directory"
	result, err := ReadConfDir(testPath)
	assert.DeepEqual(expected, result, "Expected output")
	assert.Equal(expectedErr, err.Error(), "Expected error")
}
