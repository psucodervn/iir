package parsers

import "strings"

//go:generate stringer -type=Site -output=parser_strings.go

// Site site
type Site int

const (
	// Unknown unknown site
	Unknown Site = iota
	// Codeforces Codeforces site
	Codeforces
)

// Parser is the parser interface
type Parser interface {
	ParseTaskFromHTML(html string) (Task, error)
	ParseTaskFromURL(url string) (Task, error)
	ParseContestFromHTML(html string) (Contest, error)
	ParseContestFromURL(url string) (Contest, error)
}

func NewParser(site string) Parser {
	switch strings.TrimSpace(strings.ToLower(site)) {
	case "codeforces":
		return new(CodeforcesParser)
	default:
		return nil
	}
}
