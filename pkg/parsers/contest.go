package parsers

// Contest is the contest interface
type Contest interface {
	Name() string
	Tasks() []Task
}
