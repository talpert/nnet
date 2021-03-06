package lib

import (
	"fmt"
	"math"
)

func ValidateWeights(weights [][][]float64, numInputs int) error {
	for i, w := range weights[0] {
		// match number of inputs plus threshold
		if numInputs+1 != len(w) {
			return fmt.Errorf(
				"Number of inputs: %d does not match number of weights: %d on node %d",
				numInputs, len(w), i)
		}
	}
	for i := 1; i < len(weights); i++ {
		for j, w := range weights[i] {
			// match number of nodes in previous layer plus threshold
			if len(w) != len(weights[i-1])+1 {
				return fmt.Errorf(
					"Number of nodes in layer %d: %d does not match number of weights: %d on node %d of layer %d",
					i-1, len(weights[i-1]), len(w), j, i)
			}
		}
	}
	return nil
}

func Transfer(in float64) float64 {
	if in > 0.5 {
		return 1.0
	}
	return 0.0
}

func Sigmoid(in float64) float64 {
	return in / math.Sqrt(1+math.Pow(in, 2))
}
