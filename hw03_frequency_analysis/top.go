package hw03frequencyanalysis

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

var reg = regexp.MustCompile(`(([a-zA-Zа-яА-Я]+\S[a-zA-Zа-яА-Я]+)|[a-zA-Zа-яА-Я]+)`)

type wordData struct {
	Word  string
	Count int
}

func Top10(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}

	words := reg.FindAllString(strings.ToLower(s), -1)

	wordsCount := make(map[string]int)
	for _, w := range words {
		wordsCount[w]++
	}

	wordsData := make([]wordData, 0, len(wordsCount))
	for key, val := range wordsCount {
		wordsData = append(wordsData, wordData{key, val})
	}

	sort.Slice(wordsData, func(first, second int) bool {
		if wordsData[first].Count == wordsData[second].Count {
			return wordsData[first].Word < wordsData[second].Word
		}
		return wordsData[first].Count > wordsData[second].Count
	})

	len := int(math.Min(10, float64(len(wordsData))))
	result := make([]string, len)
	for i := 0; i < len; i++ {
		result[i] = wordsData[i].Word
	}

	return result
}
