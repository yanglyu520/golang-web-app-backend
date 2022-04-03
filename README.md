
# golang web app Playbook

---
## What we building

We eventually build a webapp with apigateway and multiple lambda functions serve as microservices and RDS as a database.
## Goal in study
1. Understand the tcp server
2. Understand how to use the http server, mux
3. Understand how to connect your golang web app to a database
4. Understand how to use the aws sdk go to create a serverless web app

## Users of this Playbook

There are 3 stages of learning a programming language. These set of exercises are only targeting for understanding the packages packages that make a golang web backend.

- Syntax
- Packages
- Idioms
Therefore, it is a prerequisites that you have already understand the syntax of golang languages.

## tips

When learning a particular package, it is good to look at it in this order
1. The types
2. The interfaces of the types, and what it inherits from standard library
3. How each function work within the interface

---

# Exercise Collections
## Exercise 1: play with postman
- learning why we use postman
- learning different usage to test API with postman

## Exercise 2: code a tcp server
- task 1: create a tcp server that manually responds to a http request and returns the URL and Method of a GET request
- task 2: create a mux manually on top of the tcp server in task1 to do the routing
## Exercise 3: code a http handler

- return the url
- Return the request method
- Write a hello world page

## Exercise 4: code a multiplexer
- use the mux package to pass a handler function to do the routing

## Exercise 5: connect db to webapp
- download postgres and install
- create a table
- You can get your little web service that we made in Golang to actually query a row in that table, and return the contents as the response body

## Exercise 6: work with AWS, put web app in AWS Lambda
- create a s3 bucket in AWS and put a text
- retrieve the text from ur local go web app
- put ur app in lambda and then connect to s3 bucket

## Exercise 7: Add API Gateway in AWS to invoke the lambda function
- create an API gateway for your lambda function from exercise 6
- Manipulate your lambda function to print these results on the browser
- Make sure your lambda function is invoked by api-gateway

### Study goals
- understand how to write a go web app using apigateway and lambda functions
- how to do these in the aws console
- learn to use sam cli to start a lambda function

## Exercise 8: Create microservices with lambda functions
- put together the web with multiple lambda functions and have it invoked from API Gateway
- Draw a diagram for your architecture

---
## Extra Tools

### Install the AWS SAM CLI on macOS
[sam cli](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install-mac.html)

Packages an AWS SAM application. This command creates a . zip file of your code and dependencies, and uploads the file to Amazon Simple Storage Service (Amazon S3).

