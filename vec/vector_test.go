package vec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var v Vector

func initVector() {
	v = Vector{[]float64{0, 1, 2, 3, 4}}
}

func TestAt(t *testing.T) {
	initVector()
	assert.Equal(t, 3., v.At(3))
}

func TestSet(t *testing.T) {
	initVector()
	v.Set(3, 6)
	assert.Equal(t, 6., v.At(3))
}

func TestSetData(t *testing.T) {
	initVector()
	list := []float64{4, 3, 2, 1, 0}
	v.SetData(list)
	assert.Equal(t, 4., v.At(0))
	assert.Equal(t, 1., v.At(3))
}

func TestScale(t *testing.T) {
	initVector()
	v.Scale(2)
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 6., v.At(3))
}

func TestExp(t *testing.T) {
	initVector()
	v.Exp()
	assert.Equal(t, 1., v.At(0))
	assert.InDelta(t, 2.71828183, v.At(1), 0.00001)
}

func TestPow(t *testing.T) {
	initVector()
	v.Pow(2)
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 9., v.At(3))
}

func TestSub(t *testing.T) {
	initVector()
	other := Vector{[]float64{0, 1, 2, 3, 4}}
	v.Sub(other)
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 0., v.At(4))
}

func TestAdd(t *testing.T) {
	initVector()
	other := Vector{[]float64{0, 1, 2, 3, 4}}
	v.Add(other)
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 6., v.At(3))
}

func TestAddScalar(t *testing.T) {
	initVector()
	v.AddScalar(5)
	assert.Equal(t, 5., v.At(0))
	assert.Equal(t, 8., v.At(3))
}

func TestLen(t *testing.T) {
	initVector()
	assert.Equal(t, 5, v.Len())
}

func TestSigmoid(t *testing.T) {
	initVector()
	v.Sigmoid()
	assert.Equal(t, 0.5, v.At(0))
	assert.InDelta(t, 0.952574126822, v.At(3), 0.000001)
}

func TestSigmoidDer(t *testing.T) {
	initVector()
	v.SigmoidDer()
	assert.Equal(t, 0.25, v.At(0))
	assert.InDelta(t, 0.196611933241481, v.At(1), 0.000001)
}

func TestSum(t *testing.T) {
	initVector()
	f := v.Sum()
	assert.InDelta(t, 10., f, 0.00000001)
}

func TestMul(t *testing.T) {
	initVector()
	v.Mul(Vector{[]float64{0, 1, 2, 3, 4}})
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 9., v.At(3))
}

func TestLn(t *testing.T) {
	initVector()
	v.slice[0] = 1 // ln(0) is undefined
	v.Ln()
	assert.Equal(t, 0., v.At(0))
	assert.InDelta(t, 1.09861228, v.At(3), 0.00001)
}

func TestSwap(t *testing.T) {
	initVector()
	v.Swap(0, 1)
	assert.Equal(t, 1., v.At(0))
	assert.Equal(t, 0., v.At(1))
}

func TestReLU(t *testing.T) {
	initVector()
	v.AddScalar(-2)
	v.ReLU()
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 0., v.At(1))
	assert.Equal(t, 0., v.At(2))
	assert.Equal(t, 1., v.At(3))
}

func TestReLUDer(t *testing.T) {
	initVector()
	v.AddScalar(-1)
	v.ReLUDer()
	assert.Equal(t, Init(0, 0, 1, 1, 1), v)
}

func TestIndex(t *testing.T) {
	initVector()
	assert.Equal(t, 3, v.Index(3))
}

func TestAppend(t *testing.T) {
	initVector()
	v.Append(1)
	assert.Equal(t, Init(0, 1, 2, 3, 4, 1), v)
}

func TestContains(t *testing.T) {
	initVector()
	assert.Equal(t, false, v.Contains(-1))
	assert.Equal(t, true, v.Contains(3))
}

func TestRemove(t *testing.T) {
	initVector()
	v.Remove(4)
	v.Remove(0)
	assert.Equal(t, Init(1, 2, 3), v)
}

func TestSoftmax(t *testing.T) {
	initVector()
	v.Softmax()
	assert.InDelta(t, 0.01165623, v.At(0), 0.000001)
	assert.InDelta(t, 0.03168492, v.At(1), 0.000001)
	assert.InDelta(t, 0.086128544, v.At(2), 0.000001)
	assert.InDelta(t, 0.234121657, v.At(3), 0.000001)
	assert.InDelta(t, 0.636408649, v.At(4), 0.000001)
	assert.InDelta(t, 1., v.Sum(), 0.000001)
}
