package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	ID float64
	Value string
}

type Response struct {
	Message string
	OK bool
}

func Handler(event Event) (Response, error) {
	return Response {
		Message: fmt.Sprintf("Process Event ID %f", event.ID),
		OK: true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}