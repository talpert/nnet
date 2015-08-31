package neuron

import (
	log "github.com/Sirupsen/logrus"
)

type Node interface {
	AddListener(chan float64)
}

type Neuron struct {
	Weights []float64
	Inputs  []chan float64
	Outputs []chan float64
}

type InputNode struct {
	Listeners []chan float64
}

func New(inputNodes []*InputNode, numOutputs int, weights []float64) *Neuron {
	n := &Neuron{}
	log.Info("Creating new neuron")
	n.Inputs = make([]chan float64, len(inputNodes))
	n.Outputs = make([]chan float64, numOutputs)
	for i, input := range inputNodes {
		n.Inputs[i] = make(chan float64)
		input.AddListener(n.Inputs[i])
	}
	for i := 0; i < numOutputs; i++ {
		n.Outputs[i] = make(chan float64)
	}
	n.Weights = weights

	return n
}

func NewInputNode(numListeners int) *InputNode {
	log.Debugf("create new input node with %d listeners", numListeners)
	n := &InputNode{}
	n.Listeners = []chan float64{}
	//make([]chan float64, numListeners)
	// for i := 0; i < numListeners; i++ {
	// 	n.Listeners[i] = make(chan float64)
	// }
	return n
}

func (n *InputNode) AddListener(listener chan float64) {
	n.Listeners = append(n.Listeners, listener)
}

func (n *InputNode) In(value float64) {
	log.Debugf("posting input %f to %d Listeners.", value, len(n.Listeners))
	for _, ch := range n.Listeners {
		log.Debugf("posting %f to listener", value)
		ch <- value
	}
}

func (n *Neuron) Run() {
	log.Info("Neuron started")
	for {
		sumIn := 0.0
		log.Debugf("starting to wait for inputs. Expecting %d", len(n.Inputs))
		// each input only read once
		for i, in := range n.Inputs {
			log.Debug("waiting for input")
			sumIn += n.Weights[i] * <-in
			log.Debugf("got input. sum is: %f", sumIn)
		}
		log.Debugf("Sum of inputs is: %f", sumIn)
		output := Transfer(sumIn)
		for _, out := range n.Outputs {
			out <- output
		}
	}
}

func Transfer(in float64) float64 {
	if in > 0.5 {
		return 1.0
	}
	return 0.0
}
