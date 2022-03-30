package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"io/ioutil"
	//"os"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	//"strings"
	"context"
)

var s3Session *s3.S3

//key is the full path of the file(../filename)
const (
	BUCKET_NAME = "yang-go-web-app-test"
	REGION      = "ap-southeast-2"
	OBJECT_KEY  = "test.txt"
)

//a struct for lambda response
type Response struct {
	Message1 *s3.ListBucketsOutput
	Message2 *s3.ListObjectsV2Output
	Message3 string
}

func init() {
	awsSession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	}))
	s3Session = s3.New(awsSession)
}

func listBuckets() (resp *s3.ListBucketsOutput) {
	resp, err := s3Session.ListBuckets(&s3.ListBucketsInput{})
	checkErr(err)
	return resp
}

func listObjects() (resp *s3.ListObjectsV2Output) {
	resp, err := s3Session.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(BUCKET_NAME),
	})

	checkErr(err)
	return resp
}

func getObject() string {
	resp, err := s3Session.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(OBJECT_KEY),
	})

	checkErr(err)

	//print out the resp
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	return fmt.Sprintf(string(body))
}

func Handler(ctx context.Context) (Response, error) {
	return Response{
		Message1: listBuckets(),
		Message2: listObjects(),
		Message3: getObject(),
	}, nil
}

func main() {
	lambda.Start(Handler)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
