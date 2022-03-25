# Lambdas in AWS

## create a S3 bucket
- [doc](https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-bucket-overview.html) The easiest way to create a S3 bucket is to use the UI.

- put a text and a photo into the S3 bucket

Note: Try not to dive too deep into AWS, as it quickly drifts you from the focus of this exercise which focuses more on the golang backend.

## Learn to use AWS serverless: Lambda
- understand how the handler function work with events and response
- [aws sdk go](https://github.com/aws/aws-sdk-go)
- write a simple handler go function that handles a simple event

## Write a golang code that retrieves data from the S3 bucket
upload go binary to AWS console
- GOOS=linux go build -o main main.go && zip archive.zip main

---
More study note I have on S3 is [here](https://github.com/yanglyu520/aws-study/blob/main/S3.md), on lambda is [here](https://github.com/yanglyu520/aws-study/blob/main/Lambda.md)