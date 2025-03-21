package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	splitStr := strings.Fields(str)
	countWords := make(map[string]int)
	for _, s := range splitStr {
		countWords[s]++
	}

	sortedMap := make([]struct {
		key   string
		count int
	}, 0, len(countWords))

	for k, v := range countWords {
		sortedMap = append(sortedMap, struct {
			key   string
			count int
		}{k, v})
	}

	sort.Slice(sortedMap, func(i, j int) bool {
		return sortedMap[i].count > sortedMap[j].count ||
			(sortedMap[i].count == sortedMap[j].count && sortedMap[i].key < sortedMap[j].key)
	})

	limit := 10
	if len(sortedMap) < 10 {
		limit = len(sortedMap)
	}

	strFinal := make([]string, limit)
	for i := 0; i < limit; i++ {
		strFinal[i] = sortedMap[i].key
	}

	return strFinal
}
