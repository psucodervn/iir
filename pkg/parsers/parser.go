package parsers

import (
	"errors"
	"strings"
)

// Parser is the parser interface
type Parser interface {
	ParseTaskFromHTML(html string) (Task, error)
	ParseTaskFromURL(url string) (Task, error)
	ParseContestFromHTML(html string) (Contest, error)
	ParseContestFromURL(url string) (Contest, error)
}

// NewParser return new Parser instance depends on site
func NewParser(site string) Parser {
	switch strings.TrimSpace(strings.ToLower(site)) {
	case "codeforces":
		return new(CodeforcesParser)
	default:
		return nil
	}
}

// Judger judger interface
type Judger interface {
	Name() string
	TaskDir(task Task) string
	ContestURLPatterns() []string
	TaskURLPatterns() []string
}

var (
	// ErrInvalidHTMLFormat is invalid HTML format error
	ErrInvalidHTMLFormat = errors.New("invalid HTML format")
)
