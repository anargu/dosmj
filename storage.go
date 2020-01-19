package main

import (
	"fmt"
	"github.com/minio/minio-go"
	"io"
	"regexp"
)

const (
	TemplateDirName = "dosmj_templates"
	patternFilename = "^(\\w|-)+\\.templ$"
)

var (
	storage *minio.Client
)

func init() {
	var err error
	storage, err = minio.New(doSpacesEndpoint, doSpacesKey, doSpacesSecret, false)
	if err != nil {
		panic(err)
	}
}

func ObjectPathName(filename string) string {
	return fmt.Sprintf("%s/%s", TemplateDirName, filename)
}

func ValidateFilename(filename string) (bool, error) {
	matched, err := regexp.MatchString("\\w", filename)
	if err != nil {
		return false, err
	}
	return matched, nil
}

func GetTemplate(templateName string) (io.Reader, error) {
	object, err := storage.GetObject(doSpacesBucketName, ObjectPathName(templateName), minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return object, nil
}

func PutTemplate(file io.Reader, filename string, size int64) error {
	_, err := storage.PutObject(doSpacesBucketName, ObjectPathName(filename), file, size, minio.PutObjectOptions{})
	return err
}

func DeleteTemplate(filename string) error {
	err := storage.RemoveObject(doSpacesBucketName, ObjectPathName(filename))
	return err
}
