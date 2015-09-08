package neuron

import (
	log "github.com/Sirupsen/logrus"
)

type InputNode struct {
	Name      string
	Listeners []chan float64
}

func NewInputNode(name string) *InputNode {
	log.Debugf("create new input node %s", name)
	n := &InputNode{}
	n.Name = name
	n.Listeners = []chan float64{}
	return n
}

func (n *InputNode) AddListener(listener chan float64) {
	n.Listeners = append(n.Listeners, listener)
}

func (n *InputNode) In(value float64) {
	log.Debugf("%s: posting input %f to %d Listeners.", n.Name, value, len(n.Listeners))
	for _, ch := range n.Listeners {
		// log.Debugf("%s: posting %f to listener", n.Name, value)
		ch <- value
	}
}
