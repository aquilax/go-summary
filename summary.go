package summary

import (
	"regexp"
	"sort"
	"strings"
)

const minSentences = 2

type rankMap map[string]int

type Tokenizer interface {
	GetParagraphs(text string) []string
	GetSentences(paragraph string) []string
	GetWords(sentence string) []string
	FormatSentece(sentence string) string
}

func sentencesIntersection(s1, s2 []string) int {
	result := 0
	if len(s1) == 0 || len(s2) == 0 {
		return result
	}
	sort.StringSlice(s2).Sort()
	for _, w := range s1 {
		i := sort.SearchStrings(s2, w)
		if s2[i] == w {
			result++
		}
	}
	return result
}

func formatSentence(s string) string {
	re := regexp.MustCompile(`\W`)
	return re.ReplaceAllString(strings.ToLower(s), "")
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
		ranks[formatSentence(sentence)] = sum(matrix[i])
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
		stripped := formatSentence(s)
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
