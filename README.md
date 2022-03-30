
# golang web app Playbook
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
- create an API gateway for your lambda function
- Invoke your function through API gateway
## Exercise 8: Create microservices with lambda functions
- put together the web with multiple lambda functions and have it invoked from API Gateway
- Draw a diagram for your architecture
