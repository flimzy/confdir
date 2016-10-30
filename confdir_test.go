package confdir

import (
	"testing"

	"github.com/flimzy/testify/assert"
)

func TestConfDir(t *testing.T) {
	assert := assert.New(t)
	var testPath = "./t/a"
	var expected = [][]byte{
		[]byte("Testing"),
		[]byte("Testing"),
		[]byte("1"),
		[]byte("2"),
		[]byte("3"),
	}
	result, err := ReadConfDir(testPath)
	assert.DeepEqual(expected, result, "Expected output")
	assert.Nil(err, "No error")
}

func TestConfDir_MultipleFiles(t *testing.T) {
	assert := assert.New(t)
	var testPath = "./t/b"
	var expected = [][]byte{
		[]byte("Baz"),
		[]byte("Foo"),
	}
	result, err := ReadConfDir(testPath)
	assert.DeepEqual(expected, result, "Expected output")
	assert.Nil(err, "No error")
}

func TestConfDir_DoesntExist(t *testing.T) {
	assert := assert.New(t)
	var testPath = "./t/c"
	var expected [][]byte
	var expectedErr = "1 error(s) occurred:\n\n* open ./t/c: no such file or directory"
	result, err := ReadConfDir(testPath)
	assert.DeepEqual(expected, result, "Expected output")
	assert.Equal(expectedErr, err.Error(), "Expected error")
}
