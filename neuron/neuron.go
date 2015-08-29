package neuron

import (
	log "github.com/Sirupsen/logrus"
)

type Neuron struct {
	Weights []float64
	Inputs  []chan float64
	Output  chan float64
}

func New(numInputs int, weights []float64) *Neuron {
	n := &Neuron{}
	log.Info("Creating new neuron")
	n.Inputs = make([]chan float64, numInputs)
	for i := 0; i < numInputs; i++ {
		n.Inputs[i] = make(chan float64)
	}
	n.Weights = weights

	n.Output = make(chan float64)
	return n
}

func (n *Neuron) Run() {
	log.Info("Neuron started")
	for {
		sumIn := 0.0
		log.Debug("starting to wait for inputs")
		// each input only read once
		for i, in := range n.Inputs {
			log.Debug("waiting for input")
			sumIn += n.Weights[i] * <-in
			log.Debugf("got input. sum is: %f", sumIn)
		}
		log.Debugf("Sum of inputs is: %f", sumIn)
		n.Output <- Transfer(sumIn)
	}

}

func Transfer(in float64) float64 {
	if in > 0.5 {
		return 1.0
	}
	return 0.0
}
