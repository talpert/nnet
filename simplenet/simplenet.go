package simplenet

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"source.datanerd.us/talpert/nnet/neuron"
)

type SimpleNet struct {
	NumInputs  int
	NumLayers  int
	InputLayer []*neuron.InputNode
	OutputNode *neuron.Neuron
	Neu        *neuron.Neuron
}

func New(numInputs int, weights [][]float64) (*SimpleNet, error) {
	log.Debug("New network")
	log.Debugf("Network structure: %v", weights)
	if numInputs != len(weights[0]) {
		return nil, fmt.Errorf("Number of inputs: %d does not match number of weights: %d",
			numInputs, len(weights[0]))
	}
	s := &SimpleNet{}
	s.NumInputs = numInputs
	s.NumLayers = len(weights)
	s.InputLayer, s.OutputNode = ConstructNet(numInputs, weights)

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
	return <-s.OutputNode.Outputs[0], nil
}

func ConstructNet(numInputs int, weights [][]float64) ([]*neuron.InputNode, *neuron.Neuron) {
	log.Debugf("Constructing new net with %d inputs", numInputs)
	// create input nodes
	inNodes := make([]*neuron.InputNode, numInputs)
	numNodes := len(weights)
	// listenTo := make([]chan float64, numInputs)
	for i := 0; i < numInputs; i++ {
		log.Debugf("creating input node #%d with %d listeners", i, numNodes)
		inNodes[i] = neuron.NewInputNode(numNodes)
		// listenTo[i] = inNodes[i].Listeners[0]
	}

	outNeuron := neuron.New(inNodes, 1, weights[0])
	go outNeuron.Run()
	//create other layers
	return inNodes, outNeuron
}
