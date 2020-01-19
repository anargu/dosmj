package main

import (
	"os"
)

var (
	port,
	mjApiKeyPublic,
	mjApiKeyPrivate,
	mjSenderName,
	mjSenderEmail,

	doSpacesKey,
	doSpacesSecret,
	doSpacesEndpoint,
	doSpacesBucketName string
)

func main() {
	port = os.Getenv("PORT")

	mjApiKeyPublic = os.Getenv("MJ_APIKEY_PUBLIC")
	mjApiKeyPrivate = os.Getenv("MJ_APIKEY_PRIVATE")
	mjSenderName = os.Getenv("MJ_SENDER_NAME")
	mjSenderEmail = os.Getenv("MJ_SENDER_EMAIL")

	doSpacesKey = os.Getenv("DO_SPACES_KEY")
	doSpacesSecret = os.Getenv("DO_SPACES_SECRET")
	doSpacesEndpoint = os.Getenv("DO_SPACES_ENDPOINT")
	doSpacesBucketName = os.Getenv("DO_SPACES_NAME")

	if port == "" ||
		// MJ
		mjApiKeyPublic == "" ||
		mjApiKeyPrivate == "" ||
		mjSenderName == "" ||
		mjSenderEmail == "" ||
		// DO Spaces
		doSpacesKey == "" ||
		doSpacesSecret == "" ||
		doSpacesEndpoint == "" ||
		doSpacesBucketName == "" {

		panic("env variables should be provided")
	}

	RunServer()
}
