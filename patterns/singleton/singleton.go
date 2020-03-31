// Package singleton represent implementation of singleton pattern.
package singleton

var single *Singleton

// Singleton stands for for singleton pattern base struct.
type Singleton struct {
	Number int
}

// New returns new singleton pattern.
func New() *Singleton {
	if single == nil {
		single = &Singleton{}
	}
	return single
}
