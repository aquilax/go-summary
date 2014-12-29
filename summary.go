// This packages generates naive text summay
package summary

import (
	"sort"
)

const minSentences = 2

type rankMap map[string]int

// Tokenizer defines the interface for the tokenizing functions
type Tokenizer interface {
	GetParagraphs(text string) []string
	GetSentences(paragraph string) []string
	GetWords(sentence string) []string
	FormatSentence(sentence string) string
}

func sentencesIntersection(s1, s2 []string) int {
	result := 0
	if len(s1) == 0 || len(s2) == 0 {
		return result
	}
	sort.StringSlice(s2).Sort()
	lenS2 := len(s2)
	for _, w := range s1 {
		i := sort.SearchStrings(s2, w)
		if i < lenS2 && s2[i] == w {
			result++
		}
	}
	return result
}

func sum(a []int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum
}

func getRanks(text string, t Tokenizer) rankMap {
	sentences := t.GetSentences(text)
	n := len(sentences)
	ranks := make(rankMap)
	for i := 0; i < n; i++ {
		rank := 0
		for j := 0; j < n; j++ {
			rank += sentencesIntersection(t.GetWords(sentences[i]), t.GetWords(sentences[j]))
		}
		ranks[t.FormatSentence(sentences[i])] = rank
	}
	return ranks
}

func getBestSentence(text string, ranks rankMap, t Tokenizer) string {
	bestSentence := ""
	sentences := t.GetSentences(text)
	if len(sentences) < minSentences {
		return bestSentence
	}
	maxRank := 0
	for _, s := range sentences {
		stripped := t.FormatSentence(s)
		rank, ok := ranks[stripped]
		if ok && rank > maxRank {
			maxRank = rank
			bestSentence = s
		}
	}
	return bestSentence
}

func getSummary(text string, ranks rankMap, t Tokenizer) string {
	paragraphs := t.GetParagraphs(text)
	var result string
	for _, p := range paragraphs {
		if sentence := getBestSentence(p, ranks, t); sentence != "" {
			result += sentence + "\n"
		}
	}
	return result
}

// Summarize generates summary for the text using the provided tokenizer
func Summarize(text string, t Tokenizer) string {
	ranks := getRanks(text, t)
	return getSummary(text, ranks, t)
}
