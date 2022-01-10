### This project takes in a text and returns the top 10 most used words
## using this project
To use this project, you can run the following commands below:
1. [Install Go](https://golang.org/doc/install)
2. Install all dependencies using:
```
go mod tidy
```
3. Run the code using:
```
make run
or
go run main.go
```
The Server starts on localhost:8080 and contains only one route which is a POST request
`/` which takes in a text and returns the top 10 most used words.
It will be advisable you consume the endpoint with a POSTMAN client.