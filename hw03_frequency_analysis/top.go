package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const top = 11

var (
	rn = regexp.MustCompile("\n")
	rt = regexp.MustCompile("\t")
)

func Top10(t string) []string {
	if t == "" {
		var e []string
		return e
	}
	var wordsSort []struct {
		w string
		n int
	}
	var sliceWords []string
	var words []string

	t = rn.ReplaceAllString(t, " ")
	t = rt.ReplaceAllString(t, "")
	w := strings.Split(t, " ")

	if len(w) < 10 {
		var e []string
		return e
	}

	for _, v := range w {
		if v != "" {
			words = append(words, v)
		}
	}
	for i := 0; i < len(words); i++ {
		d := words[0]
		a := 0

		for u := 0; u < len(words); u++ {
			if words[u] == "" {
				continue
			}
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
	sort.Slice(wordsSort, func(i, j int) bool { return wordsSort[i].w > wordsSort[j].w })
	sort.SliceStable(wordsSort, func(i, j int) bool { return wordsSort[i].n < wordsSort[j].n })

	for i := len(wordsSort) - 1; i > len(wordsSort)-top; i-- {
		sliceWords = append(sliceWords, wordsSort[i].w)
	}
	return sliceWords
}
