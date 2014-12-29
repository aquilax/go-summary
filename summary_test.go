package summary

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSummary(t *testing.T) {
	Convey("Given Summary", t, func() {
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
		Convey("Generating summary works", func() {
			var testCases = []struct {
				given    string
				expected string
			}{
				{
					``,
					``,
				},
				{
					`Действието започва една вечер в двора на чорбаджи Марко, който вечеря с многолюдното си семейство. След вечерята той изпраща жена си при болния си син Асен, който е трескав, защото видял трупа на едно обезглавено от страшния бандит Емексиз Пехливан дете. В този момент се чува тропот при дувара, а слугинята се разкрещява. Марко отива да види какво е станало и там се натъква на Иван Краличът, който му се представя. Като най-близък приятел на баща му той решава да го скрие у дома си, но виковете на слугинята вече са предизвикали хората да повикат онбашия и турците вече тропат на вратата. В това време беглецът прескача зида и се измъква от заптиетата, но при сборичкването си с един от тях връхната му дреха остава в ръцете му.

След дивия си бяг, подпомогнат от тъмнината и настъпващата буря, той се измъква от турците и се скрива в една воденица на Манастирската река. Той смята, че тя е пуста, но скоро се прибира стопанинът дядо Стоян с дъщеря си Марийка и беглецът се укрива. Скоро дохождат двама турци на връщане от лов. Това са страшният Емексиз Пехливан и другарят му. Те принуждават стопанина да ги прислони във воденицата, а след като виждат младата му дъщеря, решават да го завържат и да се гаврят с момичето. В този момент Иван излиза от скривалището си, напада ги и ги убива — Емексиз Пехливан с една брадва, а на другаря му забива нож в гърлото. С това той печели вечната признателност и преданост на бащата и дъщерята. С общи усилия Иван и стопанинът заличават следите и погребват турците зад воденицата, а убитото куче пускат във водата. Решават да се приютят в манастира за през нощта и тръгват натам, но никой не подозира, че освен тях има още един свидетел на случката — идиотът Мунчо.

В манастира беглецът е добре приет от дякон Викентий, който се отнася гостоприемно с бунтовниците. На другия ден се връща и игуменът Натанаил, който носи лоши новини от града. Доктор Соколов през нощта е бил арестуван и откаран. Когато Иван го питал предния ден за Марковата къща, докторът му подарил дрехата си, защото беглецът бил дрипав. Иван сложил в джоба си бунтовнически книжа, а дрехата останала у турците. След тия известия той решава да отиде да се предаде, за да избави този добър човек.`,
					`Действието започва една вечер в двора на чорбаджи Марко, който вечеря с многолюдното си семейство.
След дивия си бяг, подпомогнат от тъмнината и настъпващата буря, той се измъква от турците и се скрива в една воденица на Манастирската река.
В манастира беглецът е добре приет от дякон Викентий, който се отнася гостоприемно с бунтовниците.
`,
				},
			}
			for _, tc := range testCases {
				So(Summarize(tc.given, NewDefaultTokenizer()), ShouldEqual, tc.expected)
			}
		})
	})
}

func BenchmarkGetRanks(b *testing.B) {
	text := `Three Beauties of the Present Day is a nishiki-e colour woodblock print of c. 1792–93 by Japanese ukiyo-e artist Kitagawa Utamaro (c. 1753–1806). The triangular composition depicts the busts of three celebrity beauties of the time: geisha Tomimoto Toyohina (middle), and teahouse waitresses Takashima Hisa (left) and Naniwa Kita (right), each adorned with an identifying family crest. Subtle differences can be detected in the faces of the subjects—a level of individualized realism at the time unusual in ukiyo-e, and a contrast with the stereotyped beauties in earlier masters such as Harunobu and Kiyonaga. The triangular positioning became a vogue in the 1790s. Utamaro produced several other pictures with this arrangement of the same three beauties, and each appeared in numerous other portraits by Utamaro and other artists. Utamaro was the leading ukiyo-e artist in the 1790s in the bijin-ga genre of pictures of female beauties, and was known in particular for his ōkubi-e, which focus on the heads. The luxurious print was published by Tsutaya Jūzaburō and made with multiple woodblocks—one for each colour—and the background was dusted with muscovite to produce a glimmering effect. `
	t := NewDefaultTokenizer()
	for i := 0; i < b.N; i++ {
		getRanks(text, t)
	}
}
