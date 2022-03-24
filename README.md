
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

## Exercise 7: API Gateway in AWS

## Exercise 8: 

---
Left to do:
1. Add the diagram to how net package work for tcp server
2. Add a diagram for api gateway to lamda to RDS in AWS architect diagram


---
Good to have, need to be trimmed as many are not relevant
- understand even better and clearer idea of the tcp 3 handshakes
- good to understand the http 1.0 vs 2.0 vs 3.0
- understand more on the apigateway
- understand better with the microservices

