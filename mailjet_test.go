package main

import (
	"fmt"
	"log"
	"testing"
)

func TestSendEmail(t *testing.T) {
	htmlPart := "<html><body><h1>Hello</h1></body></html>"
	inputMJTest := MJInput{
		From: &RecipientInputPart{
			Name:  "Parallel",
			Email: "hello@prllel.co",
		},
		To: []RecipientInputPart{
			{Name: "Anthony", Email: "aarostegui.utec@gmail.com"},
		},
		Subject:  "Hello World",
		HTMLPart: htmlPart,
	}

	err := SendEmail(&inputMJTest)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("*** Email Sent \n")
	}
}
