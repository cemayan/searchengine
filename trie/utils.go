package trie

import "strings"

func (t *trie) ConvertForIndexing(data string) map[string]map[string]int {

	data = strings.ToLower(data)
	data = strings.TrimSpace(data)

	t.Insert(data)

	lastMap := make(map[string]map[string]int)

	for i := 0; i < len(data); i++ {
		s := string([]rune(data)[0 : i+1])

		recordsMap := make(map[string]int)
		array := t.SearchByPrefix(s)
		for _, v := range array {
			recordsMap[v] = 0
		}

		lastMap[s] = recordsMap
	}

	return lastMap
}
