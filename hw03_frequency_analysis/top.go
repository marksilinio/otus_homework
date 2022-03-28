package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	wordstat := make(map[string]int)
	words := strings.Fields(text)
	result := make([]string, 10)

	if text == "" {
		return []string{}
	}

	// calculating
	for _, word := range words {
		num, err := wordstat[word]
		if err {
			wordstat[word] = num + 1
		} else {
			wordstat[word] = 1
		}
	}

	// sorting
	keys := make([]string, 0)
	for k := range wordstat {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// make top 10
	var maxvalue int
	for i := 0; i < 10; i++ {
		maxvalue = 0
		for _, k := range keys {
			if wordstat[k] > maxvalue {
				maxvalue = wordstat[k]
				result[i] = k
			}
		}
		delete(wordstat, result[i])
	}

	return result
}
