### This project takes in a text and returns the top 10 most used words
## using this project
To use this project, you can run the following commands below:
1. [Install Go](https://golang.org/doc/install)
2. clone this repository
3. Install all dependencies using:
```
go mod tidy
```
## This Projects has two implementations:
The first Implementation is a simple implementation that reads input
from a file and returns the top 10 most used words. The file (words.txt) can be found in 
the root directory of this project.<br>
**_You may want to test the implementation with your own input, just 
open the file (words.txt), clear the content, copy and paste your text and run the application
._**

Run the code using:
```
make run2
or
go run main2.go
```
The second implementation is a more advanced implementation that accept a post request from client, performs
the same operation as the first implementation and returns the result in a json format. This implementation
needs to start a server and can be found in the `main.go` file.

Run the code using:
```
make run
or
go run main.go
```
The Server starts on localhost:8080 containing only one route which is a POST request
`/` that takes in a text and returns the top 10 most used words.
It will be advisable you consume the endpoint with a POSTMAN client.