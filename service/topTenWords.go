package service

import (
	"encoding/json"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type TopWord struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

type WordMapContainer struct {
	mu                   sync.Mutex
	wg                   sync.WaitGroup
	WordOccurrenceMap    map[string]int
	WordOccurrenceStruct []TopWord
}

type ITopTenWords interface {
	TopTenWords(textInput string) []byte
	WordCount() int
	Reset()
}

func InitTopTenWordsService() ITopTenWords {
	m := make(map[string]int)
	return &WordMapContainer{
		WordOccurrenceMap: m,
	}
}

func (c *WordMapContainer) Reset() {
	c.WordOccurrenceMap = make(map[string]int)
	c.WordOccurrenceStruct = make([]TopWord, 0)
}

func (c *WordMapContainer) TopTenWords(textInput string) []byte {
	// m := make(map[string]int)
	// var wg sync.WaitGroup
	// wmContainer := WordMapContainer{
	// 	wg:                wg,
	// 	WordOccurrenceMap: m,
	// }

	r := regexp.MustCompile(`[^a-zA-Z\-'’]`)

	textInput = strings.ToLower(textInput)
	textInput = r.ReplaceAllString(textInput, " ")
	wordCandidates := strings.Split(textInput, " ")
	firstHalf := len(wordCandidates) / 2

	pushFunc := func(words []string) {
		for _, word := range words {
			if word != "" {
				c.PushWordToMap(word)
			}
		}
		c.wg.Done()
	}

	c.wg.Add(2)
	go pushFunc(wordCandidates[:firstHalf])
	go pushFunc(wordCandidates[firstHalf:])
	c.wg.Wait()

	c.MapToStruct()
	c.Sort()

	return c.ToJson()
}

func (c *WordMapContainer) PushWordToMap(word string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.WordOccurrenceMap[word] > 0 {
		c.WordOccurrenceMap[word] = c.WordOccurrenceMap[word] + 1
	} else {
		c.WordOccurrenceMap[word] = 1
	}
}

func (c *WordMapContainer) MapToStruct() {
	for key, val := range c.WordOccurrenceMap {
		c.WordOccurrenceStruct = append(c.WordOccurrenceStruct, TopWord{
			Word:  key,
			Count: val,
		})
	}
}

func (c *WordMapContainer) Sort() {
	sort.SliceStable(c.WordOccurrenceStruct, func(i, j int) bool {
		return c.WordOccurrenceStruct[i].Count > c.WordOccurrenceStruct[j].Count
	})
}

func (c *WordMapContainer) ToJson() []byte {
	wordJson, _ := json.MarshalIndent(c.WordOccurrenceStruct[:10], "", " ")

	return wordJson
}

func (c *WordMapContainer) WordCount() int {
	var count int

	for _, val := range c.WordOccurrenceStruct {
		count = count + val.Count
	}

	return count
}
