package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const top = 10

func Top10(t string) []string {
	if t == "" {
		return nil
	}

	var wordsSort []struct {
		w string
		n int
	}
	var sliceWords []string
	words := strings.Fields(t)

	for i := 0; i < len(words); i++ {
		d := words[0]
		a := 0
		for u := 0; u < len(words); u++ {
			if d == words[u] {
				a++
				copy(words[u:], words[u+1:])
				words[len(words)-1] = ""
			}
		}
		if d != "" {
			wordsSort = append(wordsSort, struct {
				w string
				n int
			}{d, a})
		}
	}
	sort.SliceStable(wordsSort, func(i, j int) bool {
		if wordsSort[i].n == wordsSort[j].n {
			return wordsSort[i].w < wordsSort[j].w
		}
		return wordsSort[i].n > wordsSort[j].n
	})

	for _, i := range wordsSort[:top] {
		sliceWords = append(sliceWords, i.w)
	}

	return sliceWords
}
