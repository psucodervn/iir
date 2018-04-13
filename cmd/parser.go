package cmd

// TestCase interface
type TestCase interface {
	Input() string
	Output() string
}

// Task interface
type Task interface {
	Name() string
	Description() string
	TestCases() []TestCase
	TimeLimit() int
	MemoryLimit() int
}
