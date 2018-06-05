package parsers

import (
	"errors"
	"path"
	"strings"
)

var (
	templateDir string
	workingDir  string
	outDir      string
	sourceDir   string
)

// SetTemplatesDir set templates dir
func SetTemplatesDir(tplDir string) {
	templateDir = tplDir
}

// SetWorkingDir set working dir
func SetWorkingDir(workDir string) {
	workingDir = workDir
	outDir = path.Join(workingDir, "out")
}

// SetSourceDir set sources dir
func SetSourceDir(srcDir string) {
	sourceDir = srcDir
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
	var judger Judger = new(Codeforces)
	return judger.WriteTask(task)
}
