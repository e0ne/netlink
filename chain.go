package netlink

import (
	"fmt"
)

// ChainAttrs are the attributes of a Chain
type Chain struct {
	Parent uint32
	Chain  uint16
}

func (c Chain) String() string {
	return fmt.Sprintf("{Parent: %d, Chain: %d}", c.Parent, c.Chain)
}

func NewChain(parent uint32, chain uint16) Chain {
	return Chain{
		Parent: parent,
		Chain:  chain,
	}
}
