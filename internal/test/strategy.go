package test

import (
	"github.com/utkuozdemir/pv-migrate/internal/task"
)

type Strategy struct {
	NameVal     string
	PriorityVal int
	CanDoVal    func(task.Task) bool
	RunFunc     func(task.Task) error
	CleanupFunc func(task.Task) error
}

func (t *Strategy) Name() string {
	return t.NameVal
}

func (t *Strategy) Priority() int {
	return t.PriorityVal
}

func (t *Strategy) CanDo(task task.Task) bool {
	return t.CanDoVal(task)
}

func (t *Strategy) Run(task task.Task) error {
	return t.RunFunc(task)
}

func (t *Strategy) Cleanup(task task.Task) error {
	return t.CleanupFunc(task)
}