package gridnet

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"source.datanerd.us/talpert/nnet/lib"
)

type GridNet struct {
	NumInputs int
	Nodes     [][]*Node
}

type Node struct {
	Value   float64
	Weights []float64
}

func New(numInputs int, weights [][][]float64) (*GridNet, error) {
	log.Debugf("Creating new GridNet with %d inputs", numInputs)
	if validEr := lib.ValidateWeights(weights, numInputs); validEr != nil {
		return nil, validEr
	}
	g := &GridNet{NumInputs: numInputs}
	g.Nodes = ConstructNet(numInputs, weights)

	return g, nil
}

func (g *GridNet) Run(inputVals []float64) (float64, error) {
	if len(inputVals) != g.NumInputs {
		return 0, fmt.Errorf("Input data size: %d does not match expected: %d",
			len(inputVals), g.NumInputs)
	}
	previous := inputVals
	for _, layer := range g.Nodes {
		current := make([]float64, len(layer))
		for j, node := range layer {
			for p := 0; p < len(previous); p++ {
				if len(previous) != len(node.Weights) {
					return 0, fmt.Errorf("Not enough weights to handle previous layer")
				}
				node.Value += previous[p] * node.Weights[p]
			}
			node.Value = lib.Transfer(node.Value)
			current[j] = node.Value
		}
		previous = current
	}
	// this will later need to handle multi return
	return g.Nodes[len(g.Nodes)-1][0].Value, nil
}

func ConstructNet(numInputs int, weights [][][]float64) [][]*Node {
	nodes := make([][]*Node, len(weights))
	for i := 0; i < len(weights); i++ {
		nodes[i] = make([]*Node, len(weights[i]))
		for j := 0; j < len(weights[i]); j++ {
			nodes[i][j] = &Node{
				Value:   0,
				Weights: weights[i][j],
			}
		}
	}

	return nodes
}
