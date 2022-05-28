package words

import (
	"searchengineqxy/searcher/utils"
	"strings"

	"github.com/wangbin/jiebago"
)

type Tokenizer struct {
	seg jiebago.Segmenter
}

func NewTokenizer(dictionaryPath string) *Tokenizer {

	tokenizer := &Tokenizer{}
	err := tokenizer.seg.LoadDictionary(dictionaryPath)
	if err != nil {
		panic(err)
	}

	return tokenizer
}

func (t *Tokenizer) Cut(text string) []string {

	text = strings.ToLower(text)
	text = utils.RemovePunctuation(text) //移除标点符号todo
	text = utils.RemoveSpace(text)       //移除空格todo

	var wordmap = make(map[string]int)

	resultChan := t.seg.CutForSearch(text, true)

	for {
		w, ok := <-resultChan
		if !ok {
			break
		}

		//去除重复分词
		_, found := wordmap[w]
		if !found {
			wordmap[w] = 1
		}
	}

	var wordSlice []string
	for k := range wordmap {
		wordSlice = append(wordSlice, k)
	}
	return wordSlice
}
