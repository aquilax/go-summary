package summary

import (
	"sort"
)

const minSentences = 2

type rankMap map[string]int

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
	var matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = sentencesIntersection(t.GetWords(sentences[i]), t.GetWords(sentences[j]))
		}
	}
	ranks := make(rankMap)
	for i, sentence := range sentences {
		ranks[t.FormatSentence(sentence)] = sum(matrix[i])
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

func Summarize(text string, t Tokenizer) string {
	ranks := getRanks(text, t)
	return getSummary(text, ranks, t)
}
