package neuron

import (
	log "github.com/Sirupsen/logrus"
)

type Node interface {
	AddListener(chan float64)
	In(float64)
}

type Neuron struct {
	Name    string
	Weights []float64
	Inputs  []chan float64
	Outputs []chan float64
}

func New(inputNodes []Node, weights []float64, name string) *Neuron {
	n := &Neuron{}
	log.Debugf("Creating new neuron %s", name)
	n.Name = name
	n.Inputs = make([]chan float64, len(inputNodes))
	n.Outputs = []chan float64{}
	// hook into the previous layer
	for i, input := range inputNodes {
		log.Debugf("%s: hooking into input %d", n.Name, i)
		n.Inputs[i] = make(chan float64)
		input.AddListener(n.Inputs[i])
	}
	n.Weights = weights

	return n
}

func (n *Neuron) AddListener(listener chan float64) {
	n.Outputs = append(n.Outputs, listener)
}

func (n *Neuron) In(v float64) {
	// meet the interface requirement
}

func (n *Neuron) Run() {
	log.Infof("Neuron %s started", n.Name)
	for {
		sumIn := 0.0
		// log.Debugf("starting to wait for inputs. Expecting %d", len(n.Inputs))
		// each input only read once
		for i, in := range n.Inputs {
			log.Debugf("%s: waiting for input %d", n.Name, i)
			sumIn += n.Weights[i] * <-in
			log.Debugf("%s: got input %d. sum is: %f", n.Name, i, sumIn)
		}
		log.Debugf("%s: Sum of %d inputs is: %f", n.Name, len(n.Inputs), sumIn)
		output := Transfer(sumIn)
		// try to parallelize this so calculations are not blocked
		log.Debugf("%s: posting %f to %d Listeners.", n.Name, output, len(n.Outputs))
		for _, out := range n.Outputs {
			// log.Debugf("%s: posting %f to listener", n.Name, output)
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
