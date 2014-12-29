package summary

import (
	"regexp"
	"strings"
)

// DefaultTokenizer implements the default tokenizing functions
type DefaultTokenizer struct{}

// NewDefaultTokenizer returns new default Tokenizer
func NewDefaultTokenizer() *DefaultTokenizer {
	return &DefaultTokenizer{}
}

// GetParagraphs splits text to paragraphs
func (t DefaultTokenizer) GetParagraphs(text string) []string {
	return strings.Split(text, "\n\n")
}

// GetSentences splites text to sentences
func (t DefaultTokenizer) GetSentences(paragraph string) []string {
	return regexp.MustCompile(`(?s)\pL.*?[\.\?!]`).FindAllString(paragraph, -1)
}

func removeDuplicates(a []string) []string {
	result := []string{}
	seen := map[string]string{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}

// GetWords returns array of unique words
func (t DefaultTokenizer) GetWords(sentence string) []string {
	re := regexp.MustCompile(`\W `)
	sentence = re.ReplaceAllString(strings.ToLower(sentence), "")
	return removeDuplicates(strings.Fields(sentence))
}

// FormatSentence Generates pesudo hash of sentence
func (t DefaultTokenizer) FormatSentence(sentence string) string {
	re := regexp.MustCompile(`\W`)
	return re.ReplaceAllString(strings.ToLower(sentence), "")
}
