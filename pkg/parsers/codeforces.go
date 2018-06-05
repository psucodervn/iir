package parsers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// CodeforcesParser parser for Codeforces
type CodeforcesParser struct {
}

// ParseTaskFromHTML parse Codeforces task from html
func (*CodeforcesParser) ParseTaskFromHTML(html string) (*Task, error) {
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
	return &Task{
		Name:  title,
		Tests: testCases,
	}, nil
}

// ParseTaskFromURL parse Codeforces task from url
func (*CodeforcesParser) ParseTaskFromURL(url string) (*Task, error) {
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

// Codeforces is the Codeforces judger
type Codeforces struct {
}

// WriteTask write codeforces task
func (j *Codeforces) WriteTask(task Task) error {
	// make dir
	if len(task.Dirs) == 0 {
		task.Dirs = []string{task.Site, task.Group, task.Name}
	}
	dir := path.Join(append([]string{sourceDir}, task.Dirs...)...)
	if err := os.MkdirAll(dir, 0755|os.ModeDir); err != nil {
		return err
	}

	// create generator
	g := NewFromFileGenerator(j, templateDir)

	// generate main file
	mainCode, err := g.WriteTaskToString(task, "main.cc")
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(path.Join(dir, "main.cc"), []byte(mainCode), 0644); err != nil {
		return err
	}

	// generate test script
	testScript, err := g.WriteTaskToString(task, "test.sh")
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(path.Join(dir, "test.sh"), []byte(testScript), 0755); err != nil {
		return err
	}

	return j.writeTests(task.Tests, dir)
}

func (j *Codeforces) writeTests(tests []TestCase, dir string) error {
	lines := bytes.Buffer{}

	// write to files
	for _, test := range tests {
		if _, err := lines.WriteString(test.Input); err != nil {
			return err
		}
		lines.WriteString("\n")
		if _, err := lines.WriteString(test.Output); err != nil {
			return err
		}
		lines.WriteString("\n-=-=-=\n")
	}

	// write to tests.io
	err := ioutil.WriteFile(path.Join(dir, "test.data"), lines.Bytes(), 0644)
	return err
}

func (j *Codeforces) writeTestsOld(tests []TestCase, dir string) error {
	// make dir
	testDir := path.Join(dir, "tests")
	if err := os.Mkdir(testDir, 0755|os.ModeDir); err != nil && !os.IsExist(err) {
		return err
	}

	// write to files
	for idx, test := range tests {
		// write input
		if err := ioutil.WriteFile(
			path.Join(testDir, fmt.Sprintf("%d.in", idx+1)), []byte(test.Input), 0644,
		); err != nil {
			return err
		}
		// write output
		if err := ioutil.WriteFile(
			path.Join(testDir, fmt.Sprintf("%d.out", idx+1)), []byte(test.Output), 0644,
		); err != nil {
			return err
		}
	}

	return nil
}

// Name returns name of judger
func (j *Codeforces) Name() string {
	panic("implement me")
}
