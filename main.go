package main

import (
	"os"
)

var (
	port,
	mjApiKeyPublic,
	mjApiKeyPrivate,
	mjSenderName,
	mjSenderEmail string
)

func main() {
	port = os.Getenv("PORT")

	mjApiKeyPublic = os.Getenv("MJ_APIKEY_PUBLIC")
	mjApiKeyPrivate = os.Getenv("MJ_APIKEY_PRIVATE")
	mjSenderName = os.Getenv("MJ_SENDER_NAME")
	mjSenderEmail = os.Getenv("MJ_SENDER_EMAIL")

	if port == "" ||
		// MJ
		mjApiKeyPublic == "" ||
		mjApiKeyPrivate == "" ||
		mjSenderName == "" ||
		mjSenderEmail == "" {

		panic("env variables should be provided")
	}

	RunServer()
}
