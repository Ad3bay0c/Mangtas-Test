package main

import (
	"fmt"
	"github.com/Ad3bay0c/golang-testing/service"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("words.txt")
	if err != nil {
        panic(err)
    }
	result := service.GetMostUsedWords(string(file)) // returns a slice of service.Mapper

	for idx, word := range result {
		fmt.Printf("%d. %s\t%d\n", idx+1, word.Key, word.Value)
	}
}
