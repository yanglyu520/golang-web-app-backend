# Lambdas in AWS

## create a S3 bucket in the aws console
- [doc](https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-bucket-overview.html)
- put a text and a photo into the S3 bucket
## Learn to use AWS serverless: Lambda
- understand how the handler function work with events and response
- [aws sdk go](https://github.com/aws/aws-sdk-go)
- write a simple handler go function that handles a simple event and test in aws lambda

## Write a golang app using aws go sdk that retrieves data from the S3 bucket, test locally first
- understand how the s3 aws package work
- write a function to list all buckets, list objects in the bucket you created, and get the txt object inside the bucket

## wrap the above app using the lambda function handler and test in aws lambda
- The Lambda function handler is the method in your function code that processes events. When your function is invoked, Lambda runs the handler method. When the handler exits or returns a response, it becomes available to handle another event.

```go
package main

import (
"fmt"
"context"
"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
Name string `json:"name"`
}

//ctx 
func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
return fmt.Sprintf("Hello %s!", name.Name ), nil
}

func main() {
lambda.Start(HandleRequest)
}
```

- give your lambda function the right IAM role, can be customised with policies to cloudwatch log and permissions to s3 buckets
- GOOS=linux go build -o main main.go && zip archive.zip main
- upload go binary to AWS console and test

[AWS Lambda context object in Go](https://docs.aws.amazon.com/lambda/latest/dg/golang-context.html)
- When Lambda runs your function, it passes a context object to the handler. This object provides methods and properties with information about the **invocation, function, and execution environment.**
- The Lambda context library provides the following global variables, methods, and properties.


---
More study note I have on S3 is [here](https://github.com/yanglyu520/aws-study/blob/main/S3.md), on lambda is [here](https://github.com/yanglyu520/aws-study/blob/main/Lambda.md)