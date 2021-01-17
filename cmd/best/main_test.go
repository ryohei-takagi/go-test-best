package main

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"strings"
	"testing"
)

type testSuite struct {
	suite.Suite
	service *objectGetter
}

func (s *testSuite) SetUpTest() {
	s.service.Bucket = "dummy"
	s.service.Key = "dummy"
}

func TestExecution(t *testing.T) {
	suite.Run(t, new(testSuite))
}

type objectGetterMock struct{}

func (m objectGetterMock) getObject() (*s3.GetObjectOutput, error) {
	b := ioutil.NopCloser(strings.NewReader("hoge"))
	return &s3.GetObjectOutput{
		Body: b,
	}, nil
}

func (s *testSuite) Test() {
	mock := objectGetterMock{}
	res, _ := readObject(mock)
	assert.Equal(s.T(), "hoge", string(res))
}
