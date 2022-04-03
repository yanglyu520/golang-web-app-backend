
## Events package
example that is usually used in lambda function together with apigateway:

```go
package main

import (
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"

  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)


func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    resp, err := http.Get(DefaultHTTPGetAddress)
    //check Err
    if err != nil {
      return events.APIGatewayProxyResponse{}, err
    }
    if resp.StatusCode != 200 {
      return events.APIGatewayProxyResponse{}, ErrNon200Response
    }
    
    //try to return the response body
    ip, err := ioutil.ReadAll(resp.Body)
    //check err
    if err != nil {
      return events.APIGatewayProxyResponse{}, err
    }

    if len(ip) == 0 {
      return events.APIGatewayProxyResponse{}, ErrNoIP
    }

    return events.APIGatewayProxyResponse{
      Body:       fmt.Sprintf("Hello, %v", string(ip)),
      StatusCode: 200,
    }, nil

}

func main() {
lambda.Start(handler)
}
```

1. Understand `request events.APIGatewayProxyRequest`:
```go
type APIGatewayProxyRequest struct {
  Resource string `json:"resource"` // The resource path defined in API Gateway
  Path string `json:"path"` // The url path for the caller
  HTTPMethod string 
  Headers map[string]string
  MultiValueHeaders map[string][]string 
  QueryStringParameters map[string]string 
  MultiValueQueryStringParameters map[string][]string 
  PathParameters map[string]string 
  StageVariables map[string]string
  RequestContext APIGatewayProxyRequestContext
  Body string
  IsBase64Encoded bool
}
```

2. Understand `events.APIGatewayProxyResponse`
```go
type APIGatewayProxyResponse struct {
  StatusCode int `json:"statusCode"`
  Headers map[string]string `json:"headers"`
  MultiValueHeaders map[string][]string `json:"multiValueHeaders"`
  Body string `json:"body"`
  IsBase64Encoded bool `json:"isBase64Encoded,omitempty"`
}
```