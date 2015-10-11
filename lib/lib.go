package lib

import (
	"fmt"
)

func ValidateWeights(weights [][][]float64, numInputs int) error {
	for i, w := range weights[0] {
		if numInputs != len(w) {
			return fmt.Errorf(
				"Number of inputs: %d does not match number of weights: %d on node %d",
				numInputs, len(w), i)
		}
	}
	for i := 1; i < len(weights); i++ {
		for j, w := range weights[i] {
			if len(w) != len(weights[i-1]) {
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
