package parsers

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// CodeforcesParser parser for Codeforces
type CodeforcesParser struct {
}

func (*CodeforcesParser) ParseTaskFromHTML(html string) (Task, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	title := doc.Find(".problem-statement .title").First().Text()
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
		name:      title,
		testCases: testCases,
	}, nil
}

func (*CodeforcesParser) ParseTaskFromURL(url string) (Task, error) {
	panic("implement me")
}

func (*CodeforcesParser) ParseContestFromHTML(html string) (Contest, error) {
	panic("implement me")
}

func (*CodeforcesParser) ParseContestFromURL(url string) (Contest, error) {
	panic("implement me")
}
