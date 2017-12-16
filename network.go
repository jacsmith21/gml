package gnn

import (
	"github.com/jacsmith21/gnn/mat"
)

// Layer a layer in a neural network
type Layer interface {
	Forward(x mat.Matrix) mat.Matrix
	BackProp(t Trainer, partial mat.Matrix) mat.Matrix
}

// Net a neural network
type Net []Layer

// Forward forwards the network
func (n Net) Forward(input mat.Matrix) mat.Matrix {
	output := input
	for _, layer := range n {
		output = layer.Forward(output)
	}

	return output
}

// BackProp updates the network parameters given the cost partial
func (n Net) BackProp(t Trainer, partial mat.Matrix) {
	for i := len(n)-1; i >= 0; i-- {
		partial = n[i].BackProp(t, partial)
	}
}
