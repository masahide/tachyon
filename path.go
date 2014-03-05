package tachyon

import (
	"path/filepath"
)

type Paths interface {
	Role(name string) string
	Vars(name string) string
	Task(name string) string
	Handler(name string) string
}

type SimplePath struct {
	Root string
}

func (s SimplePath) Role(name string) string {
	return filepath.Join(s.Root, "roles", name)
}

func (s SimplePath) Vars(name string) string {
	return filepath.Join(s.Root, name)
}

func (s SimplePath) Task(name string) string {
	return filepath.Join(s.Root, name)
}

func (s SimplePath) Handler(name string) string {
	return filepath.Join(s.Root, name)
}

type SeparatePaths struct {
	Root string
}

func (s SeparatePaths) Role(name string) string {
	return filepath.Join(s.Root, "roles", name)
}

func (s SeparatePaths) Vars(name string) string {
	return filepath.Join(s.Root, "vars", name)
}

func (s SeparatePaths) Task(name string) string {
	return filepath.Join(s.Root, "tasks", name)
}

func (s SeparatePaths) Handler(name string) string {
	return filepath.Join(s.Root, "handlers", name)
}
