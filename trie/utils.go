package trie

import "strings"

func (t *trie) ConvertForIndexing(data string) map[string]map[string]int {

	data = strings.ToLower(data)
	data = strings.TrimSpace(data)

	t.Insert(data)

	lastMap := make(map[string]map[string]int)

	for i := 1; i <= len(data); i++ {
		recordsMap := make(map[string]int)
		array := t.SearchByPrefix(data[0:i])
		for _, v := range array {
			recordsMap[v] = 0
		}

		lastMap[data[0:i]] = recordsMap
	}

	return lastMap
}
