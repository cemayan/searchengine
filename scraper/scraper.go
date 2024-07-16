package scraper

import (
	"github.com/cemayan/searchengine/trie"
)

type Scraper interface {
	getRecords() ([]string, error)
	Start() error
}

type scraper struct {
	data string
}

func (s scraper) Start() error {
	_, err := s.getRecords()
	if err != nil {
		return err
	}

	return nil
}

func (s scraper) getRecords() ([]string, error) {
	_trie := trie.New()
	indexing := _trie.ConvertForIndexing(s.data)

	arr := make([]string, 0)

	for key, _ := range indexing {
		arr = append(arr, key)
	}

	return arr, nil
}

func New(data string) Scraper {
	return &scraper{data: data}
}
