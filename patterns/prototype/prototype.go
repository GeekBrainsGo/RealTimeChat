// Package prototype represent implementation of prototype pattern.
package prototype

// Proto stands for basic prototype object.
type Proto struct {
	N *int
}

// Copy return copy of prototype object.
func (p *Proto) Copy() Proto {
	return Proto{&(*p.N)}
}
