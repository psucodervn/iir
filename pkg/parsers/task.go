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

// DataFormat format of input or output data
type DataFormat struct {
	Type     string
	FileName string `json:"fileName"`
	Pattern  string
}

// Task task in json format
type Task struct {
	Site        string
	ID          string `json:"id"`
	Name        string
	Group       string
	URL         string `json:"url"`
	Input       DataFormat
	Output      DataFormat
	MemoryLimit int
	TimeLimit   int
	TestType    string `json:"testType"`
	Tests       []TestCase
	Dirs        []string
	Meta        map[string]interface{}
}
