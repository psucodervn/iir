package cmd

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

// Task interface
type Task interface {
	Name() string
	Description() string
	TestCases() []TestCase
	TimeLimit() int64
	MemoryLimit() int64
}

// DefaultTask is the default implement of Task interface
type DefaultTask struct {
	name        string
	description string
	testCases   []TestCase
	timeLimit   int64
	memoryLimit int64
}

// SetName set task name
func (task *DefaultTask) SetName(name string) {
	task.name = name
}

// SetDescription set task description
func (task *DefaultTask) SetDescription(description string) {
	task.description = description
}

// SetTestCases set task test cases
func (task *DefaultTask) SetTestCases(testCases []TestCase) {
	task.testCases = testCases
}

//SetTimeLimit set task time limit
func (task *DefaultTask) SetTimeLimit(timeLimit int64) {
	task.timeLimit = timeLimit
}

//SetMemoryLimit set task memory limit
func (task *DefaultTask) SetMemoryLimit(memoryLimit int64) {
	task.memoryLimit = memoryLimit
}

// Name return task name
func (task *DefaultTask) Name() string {
	return task.name
}

// Description return task description
func (task *DefaultTask) Description() string {
	panic("implement me")
}

// TestCases return all test case
func (task *DefaultTask) TestCases() []TestCase {
	return task.testCases
}

// TimeLimit return time limit in nanosecond unit
func (task *DefaultTask) TimeLimit() int64 {
	return task.timeLimit
}

// MemoryLimit return memory limit in Byte unit
func (task *DefaultTask) MemoryLimit() int64 {
	return task.memoryLimit
}
