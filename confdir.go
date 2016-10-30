// Package confdir reads a directory as though it were a config file
// It is meant to be used as a primitive for reading conf.d/ style
// directories.
package confdir

import (
	"bufio"
	"io/ioutil"
	"os"

	"github.com/hashicorp/go-multierror"
)

// ReadConfDir recursively reads all files contained in the passed directory,
// and returns an array of un-parsed lines read. Any errors are accumulated
// and returned together as a multierror. This means you can receive conf
// data and an error in response.
func ReadConfDir(path string) ([][]byte, error) {
	var result = make([][]byte, 0)
	var resultErr error
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, multierror.Append(resultErr, err)
	}
	for _, file := range files {
		filePath := path + "/" + file.Name()
		if file.IsDir() {
			subResult, subErr := ReadConfDir(filePath)
			resultErr = multierror.Append(resultErr, subErr)
			result = append(result, subResult...)
			continue
		}
		f, err := os.Open(filePath)
		if err != nil {
			resultErr = multierror.Append(resultErr, err)
			continue
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			result = append(result, scanner.Bytes())
		}
		if err := scanner.Err(); err != nil {
			resultErr = multierror.Append(resultErr, err)
		}
	}
	return result, resultErr
}
