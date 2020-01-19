package main

import (
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
