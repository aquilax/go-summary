package summary

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSummary(t *testing.T) {
	Convey("Given Summary", t, func() {
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
				So(formatSentence(tc.given), ShouldEqual, tc.expected)
			}
		})
		Convey("Intersection works", func() {
			var testCases = []struct {
				s1       []string
				s2       []string
				expected int
			}{
				{
					[]string{},
					[]string{},
					0,
				},
				{
					[]string{"one", "two", "three"},
					[]string{"two", "four", "six"},
					1,
				},
			}
			for _, tc := range testCases {
				So(sentencesIntersection(tc.s1, tc.s2), ShouldEqual, tc.expected)
			}

		})
		Convey("Sum works", func() {
			var testCases = []struct {
				arr      []int
				expected int
			}{
				{
					[]int{},
					0,
				},
				{
					[]int{1, 2, 3},
					6,
				},
			}
			for _, tc := range testCases {
				So(sum(tc.arr), ShouldEqual, tc.expected)
			}
		})
	})
}
