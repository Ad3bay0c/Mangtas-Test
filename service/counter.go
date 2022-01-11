package service

import (
	"regexp"
	"sort"
	"strings"
)

type Mapper struct {
	Key   string `json:"word"`
	Value int    `json:"No_of_appearance"`
}

// GetMostUsedWords function to map the words to the number of appearance and returns 10 most used words
func GetMostUsedWords(text string) []Mapper {
	var mapper []Mapper

	arrayText := regexp.MustCompile("[^0-9a-zA-Z]+").Split(text, -1) // split the text and returns slice of substrings

	m := map[string]int{}
	for _, word := range arrayText {
		word = strings.ToLower(word)
		m[word]++
	}

	// since order of map is not guaranteed, loop through and
	// append the result to the mapper slice
	for key, val := range m {
		mapper = append(mapper, Mapper{Key: key, Value: val})
	}

	// sorts the mapper value in descending order (i.e. from highest to lowest)
	sort.SliceStable(mapper, func(i, j int) bool {
		return mapper[i].Value > mapper[j].Value
	})

	length := len(mapper)
	if length > 10 {
		length = 10
	}
	return mapper[:length]
}
