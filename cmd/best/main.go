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

type objectGetterInterface interface {
	getObject() (*s3.GetObjectOutput, error)
}

type objectGetter struct {
	Bucket string
	Key    string
}

func newObjectGetter(bucket, key string) *objectGetter {
	return &objectGetter{
		Bucket: bucket,
		Key:    key,
	}
}

func (getter *objectGetter) getObject() (*s3.GetObjectOutput, error) {
	obj, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(getter.Bucket),
		Key:    aws.String(getter.Key),
	})
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func readObject(t objectGetterInterface) ([]byte, error) {
	obj, err := t.getObject()
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
	t := newObjectGetter("", "")
	res, err := readObject(t)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(res))

	return
}