package main

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"testing"
)

func validateFilename(filename string) (bool, error) {
	matched, err := regexp.MatchString(patternFilename, filename)
	if err != nil {
		return false, err
	}
	return matched, nil
}

func TestAValidFilename(t *testing.T) {
	filename := "hello.templ"
	isValid, err := validateFilename(filename)
	if err != nil {
		log.Fatal(err)
	}
	if !isValid {
		log.Fatal(fmt.Sprintf("isvalid: %v", isValid))
	}

}

var testFileName = "test.templ"

func TestPuttingNewTemplate(t *testing.T) {
	html := "<html><body><h1>Hello</h1></body></html>"
	buff := bytes.NewBuffer([]byte(html))

	err := PutTemplate(buff, testFileName, int64(buff.Len()))
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeletingNewTemplate(t *testing.T) {
	err := DeleteTemplate(testFileName)
	if err != nil {
		t.Fatal(err)
	}
}
