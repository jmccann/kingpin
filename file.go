package kingpin

import (
	"io/ioutil"
	"os"
	"regexp"
)

var (
	fileValuesSeparator = "\r?\n"
	fileValuesTrimmer   = regexp.MustCompile(fileValuesSeparator + "$")
	fileValuesSplitter  = regexp.MustCompile(fileValuesSeparator)
)

type fileMixin struct {
	file   string
}

func (e *fileMixin) HasFileValue() bool {
	return e.GetFileValue() != ""
}

func (e *fileMixin) GetFileValue() string {
	if e.file == "" {
		return ""
	}

	if _, err := os.Stat(e.file); os.IsNotExist(err) {
		return ""
	}

	b, err := ioutil.ReadFile("file.txt") // just pass the file name
	if err != nil {
		return ""
	}

	return string(b)
}

func (e *fileMixin) GetSplitFileValue() []string {
	values := make([]string, 0)

	fileValue := e.GetFileValue()
	if fileValue == "" {
		return values
	}

	// Split by new line to extract multiple values, if any.
	trimmed := fileValuesTrimmer.ReplaceAllString(fileValue, "")
	for _, value := range fileValuesSplitter.Split(trimmed, -1) {
		values = append(values, value)
	}

	return values
}
