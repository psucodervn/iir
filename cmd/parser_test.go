package cmd

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTestCase(t *testing.T) {
	as := assert.New(t)

	tc := NewTestCase("input", "output")
	as.NotNil(tc)
	as.Equal(tc.Input, "input")
	as.Equal(tc.Output, "output")
}

func TestDefaultTaskImplementsTask(t *testing.T) {
	as := assert.New(t)

	as.Implements((*Task)(nil), new(DefaultTask))
}

func TestDefaultTask_SetName(t *testing.T) {
	type fields struct {
		name        string
		description string
		testCases   []TestCase
		timeLimit   int64
		memoryLimit int64
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{{
		name:   "set name to Task1",
		fields: fields{name: "Task1"},
		args:   args{name: "Task1"},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &DefaultTask{
				name:        tt.fields.name,
				description: tt.fields.description,
				testCases:   tt.fields.testCases,
				timeLimit:   tt.fields.timeLimit,
				memoryLimit: tt.fields.memoryLimit,
			}
			task.SetName(tt.args.name)
		})
	}
}

func TestDefaultTask_SetDescription(t *testing.T) {
	type fields struct {
		name        string
		description string
		testCases   []TestCase
		timeLimit   int64
		memoryLimit int64
	}
	type args struct {
		description string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &DefaultTask{
				name:        tt.fields.name,
				description: tt.fields.description,
				testCases:   tt.fields.testCases,
				timeLimit:   tt.fields.timeLimit,
				memoryLimit: tt.fields.memoryLimit,
			}
			task.SetDescription(tt.args.description)
		})
	}
}

func TestDefaultTask_SetTestCases(t *testing.T) {
	type fields struct {
		name        string
		description string
		testCases   []TestCase
		timeLimit   int64
		memoryLimit int64
	}
	type args struct {
		testCases []TestCase
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &DefaultTask{
				name:        tt.fields.name,
				description: tt.fields.description,
				testCases:   tt.fields.testCases,
				timeLimit:   tt.fields.timeLimit,
				memoryLimit: tt.fields.memoryLimit,
			}
			task.SetTestCases(tt.args.testCases)
		})
	}
}

func TestDefaultTask_SetTimeLimit(t *testing.T) {
	type fields struct {
		name        string
		description string
		testCases   []TestCase
		timeLimit   int64
		memoryLimit int64
	}
	type args struct {
		timeLimit int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &DefaultTask{
				name:        tt.fields.name,
				description: tt.fields.description,
				testCases:   tt.fields.testCases,
				timeLimit:   tt.fields.timeLimit,
				memoryLimit: tt.fields.memoryLimit,
			}
			task.SetTimeLimit(tt.args.timeLimit)
		})
	}
}

func TestDefaultTask_SetMemoryLimit(t *testing.T) {
	type fields struct {
		name        string
		description string
		testCases   []TestCase
		timeLimit   int64
		memoryLimit int64
	}
	type args struct {
		memoryLimit int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &DefaultTask{
				name:        tt.fields.name,
				description: tt.fields.description,
				testCases:   tt.fields.testCases,
				timeLimit:   tt.fields.timeLimit,
				memoryLimit: tt.fields.memoryLimit,
			}
			task.SetMemoryLimit(tt.args.memoryLimit)
		})
	}
}

func TestDefaultTask_Name(t *testing.T) {
	type fields struct {
		name        string
		description string
		testCases   []TestCase
		timeLimit   int64
		memoryLimit int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{{
		name:   "task name should be `Task1`",
		fields: fields{name: "Task1"},
		want:   "Task1",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &DefaultTask{
				name:        tt.fields.name,
				description: tt.fields.description,
				testCases:   tt.fields.testCases,
				timeLimit:   tt.fields.timeLimit,
				memoryLimit: tt.fields.memoryLimit,
			}
			if got := task.Name(); got != tt.want {
				t.Errorf("DefaultTask.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultTask_Description(t *testing.T) {
	type fields struct {
		name        string
		description string
		testCases   []TestCase
		timeLimit   int64
		memoryLimit int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &DefaultTask{
				name:        tt.fields.name,
				description: tt.fields.description,
				testCases:   tt.fields.testCases,
				timeLimit:   tt.fields.timeLimit,
				memoryLimit: tt.fields.memoryLimit,
			}
			if got := task.Description(); got != tt.want {
				t.Errorf("DefaultTask.Description() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCases(t *testing.T) {
	type fields struct {
		name        string
		description string
		testCases   []TestCase
		timeLimit   int64
		memoryLimit int64
	}
	testCases := []TestCase{
		TestCase{},
	}
	tests := []struct {
		name   string
		fields fields
		want   []TestCase
	}{{
		name: "set testCases to empty",
		fields: fields{
			testCases: []TestCase{},
		},
		want: []TestCase{},
	}, {
		name: "set testCases to an array",
		fields: fields{
			testCases: testCases,
		},
		want: testCases,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &DefaultTask{
				name:        tt.fields.name,
				description: tt.fields.description,
				testCases:   tt.fields.testCases,
				timeLimit:   tt.fields.timeLimit,
				memoryLimit: tt.fields.memoryLimit,
			}
			if got := task.TestCases(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultTask.TestCases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultTask_TimeLimit(t *testing.T) {
	type fields struct {
		name        string
		description string
		testCases   []TestCase
		timeLimit   int64
		memoryLimit int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{{
		name: "set timeLimit to 1s",
		fields: fields{
			timeLimit: 1000000,
		},
		want: 1000000,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &DefaultTask{
				name:        tt.fields.name,
				description: tt.fields.description,
				testCases:   tt.fields.testCases,
				timeLimit:   tt.fields.timeLimit,
				memoryLimit: tt.fields.memoryLimit,
			}
			if got := task.TimeLimit(); got != tt.want {
				t.Errorf("DefaultTask.TimeLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultTask_MemoryLimit(t *testing.T) {
	type fields struct {
		name        string
		description string
		testCases   []TestCase
		timeLimit   int64
		memoryLimit int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{{
		"1000 mem",
		fields{
			memoryLimit: 1000,
		},
		1000,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &DefaultTask{
				name:        tt.fields.name,
				description: tt.fields.description,
				testCases:   tt.fields.testCases,
				timeLimit:   tt.fields.timeLimit,
				memoryLimit: tt.fields.memoryLimit,
			}
			if got := task.MemoryLimit(); got != tt.want {
				t.Errorf("DefaultTask.MemoryLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}
