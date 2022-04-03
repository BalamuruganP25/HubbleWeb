package service

import (
	"sort"
	"strings"
)

type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

type ToatalWordCount struct {
	Topwords []WordCount `json:"topwords"`
}

func FindToptenwords(input string) ToatalWordCount {
	var Toptencount []WordCount
	var Resposedata ToatalWordCount

	replacer := strings.NewReplacer(",", "", "!", "", "?", "", ".", "")
	res := replacer.Replace(input)
	words := strings.Split(res, " ")
	m := make(map[string]int)
	for _, word := range words {
		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	// create and fill slice of word-count pairs for sorting by count
	wordCounts := make([]WordCount, 0, len(m))
	for key, val := range m {
		wordCounts = append(wordCounts, WordCount{Word: key, Count: val})
	}

	// sort wordCount slice by decreasing count number
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].Count > wordCounts[j].Count
	})

	for i := 0; i < len(wordCounts) && i < 10; i++ {
		var s WordCount
		s.Count = wordCounts[i].Count
		s.Word = wordCounts[i].Word
		Toptencount = append(Toptencount, s)
	}
	Resposedata.Topwords = Toptencount
	return Resposedata
}
