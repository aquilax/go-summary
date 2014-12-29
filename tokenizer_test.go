package summary

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTokenizer(t *testing.T) {
	Convey("Given default tokenizer", t, func() {
		dt := NewDefaultTokenizer()
		Convey("Formatting Sentence works", func() {
			var testCases = []struct {
				given    string
				expected string
			}{
				{
					` a nice day in the park  `,
					`anicedayinthepark`,
				},
				{
					` a Nice day in,.#$^% tHe park  `,
					`anicedayinthepark`,
				},
			}
			for _, tc := range testCases {
				So(dt.FormatSentence(tc.given), ShouldEqual, tc.expected)
			}
		})
		Convey("Getting sentences works", func() {
			text := "I like turtles. Do you? Awesome! Hahaha. Lol!!! What's going on????"
			expected := []string{
				"I like turtles.",
				"Do you?",
				"Awesome!",
				"Hahaha.",
				"Lol!",
				"What's going on?",
			}
			So(dt.GetSentences(text), ShouldResemble, expected)
		})
	})
}
