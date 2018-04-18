package parsers

import (
	"errors"
	"strings"
)

// Site type
type Site string

var (
	// SiteCodeforces site codeforces
	SiteCodeforces Site = "codeforces"
	// SiteKattis site kattis
	SiteKattis Site = "kattis"
)

func (s Site) String() string {
	return string(s)
}

var (
	templateDir string
	workingDir  string
)

// SetTemplatesDir set templates dir
func SetTemplatesDir(tplDir string) {
	templateDir = tplDir
}

// SetWorkingDir set working dir
func SetWorkingDir(workDir string) {
	workingDir = workDir
}

// Parser is the parser interface
type Parser interface {
	ParseTaskFromHTML(html string) (*Task, error)
	ParseTaskFromURL(url string) (*Task, error)
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
	WriteTask(task Task) error
}

var (
	// ErrInvalidHTMLFormat is invalid HTML format error
	ErrInvalidHTMLFormat = errors.New("invalid HTML format")
	// ErrSiteNotSupport is site not support error
	ErrSiteNotSupport = errors.New("site not support")
)

// AddTask add a ask with json format
func AddTask(task Task) error {
	var judger Judger
	switch task.Site {
	case SiteCodeforces:
		judger = new(Codeforces)
	default:
		return ErrSiteNotSupport
	}

	return judger.WriteTask(task)
}
