// Package builder stands for basic builder pattern.
package builder

// New stands for basic builder struct.
type New struct {
	N int
}

// Build builds a builder.
func (a *New) Build(n int) New {
	return New{
		N: n,
	}
}

// Build stands for builder constructor.
func Build(n int) New {
	return New{
		N: n,
	}
}
