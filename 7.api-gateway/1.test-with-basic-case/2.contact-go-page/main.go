package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Body string
	Status int
}

func Handler () (Response, error) {
	return Response {
		Body: fmt.Sprint("Hello from contact page"),
		Status: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}