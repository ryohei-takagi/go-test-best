package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"log"
)

var awssess = newSession()
var s3Client = newS3Client()

const defaultRegion = "ap-northeast-1"

func newSession() *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}

func newS3Client() *s3.S3 {
	return s3.New(awssess, &aws.Config{
		Region: aws.String(defaultRegion),
	})
}

func readObject(bucket, key string) ([]byte, error) {
	obj, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	defer obj.Body.Close()
	res, err := ioutil.ReadAll(obj.Body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func main() {
	res, err := readObject("", "")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(res))

	return
}