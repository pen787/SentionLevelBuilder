package emath

import (
	"sync/atomic"
)

type IDGenerator struct {
	counter atomic.Uint32
}

func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

func (g *IDGenerator) Next() uint32 {
	return g.counter.Add(1)
}
