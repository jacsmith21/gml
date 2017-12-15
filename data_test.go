package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var d DataSet

func initDataSet() {
	x1 := vec.Init([]float64{0, 0, 0})
	x2 := vec.Init([]float64{1, 1, 1})
	x3 := vec.Init([]float64{2, 2, 2})
	x4 := vec.Init([]float64{3, 3, 3})
	y1 := vec.Init([]float64{0})
	y2 := vec.Init([]float64{1})
	y3 := vec.Init([]float64{2})
	y4 := vec.Init([]float64{3})
	d = DataSet{
		data:   []vec.Vector{x1, x2, x3, x4},
		labels: []vec.Vector{y1, y2, y3, y4},
	}
}

type RanderMock struct {
	count int
	mock.Mock
}

func (m *RanderMock) Intn(n int) int {
	args := m.Called(n)
	return args.Int(0)

}

func TestSampleCount(t *testing.T) {
	initDataSet()
	assert.Equal(t, 4, d.SampleCount())
}

func TestShuffle(t *testing.T) {
	initDataSet()
	rander := new(RanderMock)
	rander.On("Intn", 2).Return(2)
	rander.On("Intn", 3).Return(3)
	rander.On("Intn", 4).Return(0)

	d.Shuffle(rander)

	assert.Equal(t, 2., d.labels[3].At(0))
	assert.Equal(t, 2., d.data[3].At(0))
	assert.Equal(t, 3., d.labels[0].At(0))
	assert.Equal(t, 3., d.data[0].At(0))
	rander.AssertExpectations(t)
}

func TestGenerateBatches(t *testing.T) {

}
