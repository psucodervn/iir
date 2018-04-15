package parsers

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// CodeforcesParser parser for Codeforces
type CodeforcesParser struct {
}

// ParseTaskFromHTML parse Codeforces task from html
func (*CodeforcesParser) ParseTaskFromHTML(html string) (Task, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	title := doc.Find(".problem-statement .title").First().Text()
	if title == "" {
		return nil, ErrInvalidHTMLFormat
	}
	sampleTests := doc.Find(".sample-tests").First()
	inputs := sampleTests.Find("div.input pre")
	outputs := sampleTests.Find("div.output pre")
	testCases := make([]TestCase, inputs.Length())
	for i := 0; i < inputs.Length(); i++ {
		input, _ := inputs.Eq(i).Html()
		output, _ := outputs.Eq(i).Html()
		testCases[i] = TestCase{
			Input:  strings.Replace(input, "<br/>", "\n", -1),
			Output: strings.Replace(output, "<br/>", "\n", -1),
		}
	}
	return &DefaultTask{
		title:     title,
		testCases: testCases,
	}, nil
}

// ParseTaskFromURL parse Codeforces task from url
func (*CodeforcesParser) ParseTaskFromURL(url string) (Task, error) {
	panic("implement me")
}

// ParseContestFromHTML parse Codeforces contest from html
func (*CodeforcesParser) ParseContestFromHTML(html string) (Contest, error) {
	panic("implement me")
}

// ParseContestFromURL parse Codeforces contest from url
func (*CodeforcesParser) ParseContestFromURL(url string) (Contest, error) {
	panic("implement me")
}

type CodeforcesJudger struct {
}

func (*CodeforcesJudger) Name() string {
	panic("implement me")
}

func (*CodeforcesJudger) TaskDir(task Task) string {
	panic("implement me")
}

func (*CodeforcesJudger) ContestURLPatterns() []string {
	panic("implement me")
}

func (*CodeforcesJudger) TaskURLPatterns() []string {
	panic("implement me")
}