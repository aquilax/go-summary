package summary

import (
	"regexp"
	"strings"
)

type DefaultTokenizer struct{}

func NewDefaultTokenizer() *DefaultTokenizer {
	return &DefaultTokenizer{}
}

func (t DefaultTokenizer) GetParagraphs(text string) []string {
	return strings.Split(text, "\n\n")
}

func (t DefaultTokenizer) GetSentences(paragraph string) []string {
	return regexp.MustCompile(`(?s)\pL.*?[\.\?!]`).FindAllString(paragraph, -1)
}

func (t DefaultTokenizer) GetWords(sentence string) []string {
	re := regexp.MustCompile(`\W `)
	sentence = re.ReplaceAllString(strings.ToLower(sentence), "")
	return strings.Fields(sentence)
}

func (t DefaultTokenizer) FormatSentence(sentence string) string {
	re := regexp.MustCompile(`\W`)
	return re.ReplaceAllString(strings.ToLower(sentence), "")
}
