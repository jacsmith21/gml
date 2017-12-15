package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

var exp mat.Matrix
var act mat.Matrix

func initCostStuff() {
	exp = mat.Init(vec.Init(0, 1, 2))
	act = mat.Init(vec.Init(0, 0, 0))
}

func TestSECost(t *testing.T) {
	se := SE{}
	cost := se.Cost(exp, act)

	assert.Equal(t, 1, cost.Len())
	assert.Equal(t, (0. + 1. + 4.), cost.At(0))
}

func TestSEDef(t *testing.T) {
	mse := SE{}
	cost := mse.Der(exp, act)

	assert.Equal(t, 1, cost.Len())
	assert.Equal(t, -(0.+1.+2.)*2, cost.At(0))
}
