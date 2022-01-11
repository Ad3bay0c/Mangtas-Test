package service

import (
	"regexp"
	"sort"
	"strings"
)

type Mapper struct {
	Key 	string	`json:"word"`
	Value 	int	    `json:"No_of_appearance"`
}

func GetMostUsedWords(text string) []Mapper {
	var mapper []Mapper
	arrayText := regexp.MustCompile("[^0-9a-zA-Z]+").Split(text, -1)

	m := map[string]int{}
	for _, word := range arrayText {
		word = strings.ToLower(word)
		m[word]++
	}

	for key, val := range m {
		mapper = append(mapper, Mapper{Key: key, Value: val})
    }

	sort.SliceStable(mapper, func(i, j int) bool {
		return mapper[i].Value > mapper[j].Value
	})

	length := len(mapper)
	if length > 10 {
		length = 10
	}
	return mapper[:length]
}
