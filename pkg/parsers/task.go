package parsers

// TestCase struct
type TestCase struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

// NewTestCase create new TestCase with input and output
func NewTestCase(input, output string) *TestCase {
	tc := &TestCase{
		Input:  input,
		Output: output,
	}
	return tc
}

// Task task in json format
type Task struct {
	Site        Site
	Name        string
	URL         string `json:"url"`
	Input       string
	Output      string
	MemoryLimit string
	TimeLimit   string
	Tests       []TestCase
	Dirs        []string
	Meta        map[string]interface{}
}
