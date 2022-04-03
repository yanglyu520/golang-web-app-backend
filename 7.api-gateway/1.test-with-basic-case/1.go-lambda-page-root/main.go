//Lambda Function Go Code
package main

import "github.com/aws/aws-lambda-go/lambda"
import "github.com/aws/aws-lambda-go/events"

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var stringResponse string = "Yay a successful Response!!"
	ApiResponse := events.APIGatewayProxyResponse{Body: stringResponse, StatusCode: 200}
	return ApiResponse, nil
}
