package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
  "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp, err := http.Get(DefaultHTTPGetAddress)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if resp.StatusCode != 200 {
		return events.APIGatewayProxyResponse{}, ErrNon200Response
	}

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if len(ip) == 0 {
		return events.APIGatewayProxyResponse{}, ErrNoIP
	}

	return events.APIGatewayProxyResponse{
		Body:       getObject(),
		StatusCode: 200,
	}, nil
}

//below are functions for s3
var s3Session *s3.S3

//key is the full path of the file(../filename)
const (
	BUCKET_NAME = "yang-go-web-app-test"
	REGION      = "ap-southeast-2"
	OBJECT_KEY  = "test.txt"
)

func init() {
	awsSession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	}))
	s3Session = s3.New(awsSession)
}

func main() {
	lambda.Start(handler)
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
