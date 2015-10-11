package simplenet

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	//	"reflect"
	"source.datanerd.us/talpert/nnet/lib"
	"source.datanerd.us/talpert/nnet/neuron"
)

type SimpleNet struct {
	NumInputs  int
	InputLayer []neuron.Node
	Output     chan float64
}

func New(numInputs int, weights [][][]float64) (*SimpleNet, error) {
	log.Debug("New network")
	log.Debugf("Network structure: %v", weights)
	if validEr := lib.ValidateWeights(weights, numInputs); validEr != nil {
		return nil, validEr
	}
	s := &SimpleNet{}
	s.NumInputs = numInputs
	log.Debugf("Layers found: %d", len(weights))
	s.InputLayer, s.Output = ConstructNet(numInputs, weights)

	return s, nil
}

func (s *SimpleNet) Run(inputVals []float64) (float64, error) {
	if len(inputVals) != s.NumInputs {
		return 0.0, fmt.Errorf("Input data size: %d does not match expected: %d",
			len(inputVals), s.NumInputs)
	}

	for i, inNode := range s.InputLayer {
		log.Debugf("Input %d is: %f", i, inputVals[i])
		inNode.In(inputVals[i])
	}

	log.Debug("waiting for ouput")
	return <-s.Output, nil
}

func ConstructNet(numInputs int, weights [][][]float64) ([]neuron.Node, chan float64) {
	log.Infof("Constructing new net with %d inputs and %d layers", numInputs, len(weights))
	// create input nodes
	inNodes := make([]neuron.Node, numInputs)
	log.Debug("Creating input layer")
	for i := 0; i < numInputs; i++ {
		log.Debugf("creating input node #%d", i)
		inNodes[i] = neuron.NewInputNode(fmt.Sprintf("I(%d)", i))
	}

	// create middle layers
	previous := inNodes
	for i := 0; i < len(weights); i++ { // for each layer
		log.Debugf("Creating layer %d", i)
		layer := []neuron.Node{}
		for j := 0; j < len(weights[i]); j++ { // for each node in layer
			log.Debugf("Creating node %d in layer %d with weights %v", j, i, weights[i][j])
			current := neuron.New(previous, weights[i][j], fmt.Sprintf("N(%d:%d)", i, j))
			layer = append(layer, current)
			go current.Run()
		}
		previous = layer
	}

	log.Debugf("Net has %d outputs", len(previous))
	output := make(chan float64)
	previous[0].AddListener(output)
	return inNodes, output
}
