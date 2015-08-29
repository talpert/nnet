package main

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"source.datanerd.us/talpert/nnet/neuron"
)

func init() {
	var (
		logDebug = kingpin.Flag("debug", "Enable debug level logging.").Short('d').Bool()
	)
	kingpin.Parse()
	if *logDebug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	log.Info("starting")
	// inputVals := []float64{1.0, 1.0}
	sampleData := [][]float64{
		[]float64{1.0, 1.0},
		[]float64{1.0, 0.0},
		[]float64{0.0, 1.0},
		[]float64{0.0, 0.0},
	}
	neu := neuron.New(2, []float64{0.5, 0.5})
	go neu.Run()
	for _, inputVals := range sampleData {
		for i, ch := range neu.Inputs {
			log.Debugf("Input: %f", inputVals[i])
			ch <- inputVals[i]
		}

		log.Debug("waiting for ouput")
		ret := <-neu.Output
		log.Infof("Inputs: %v", inputVals)
		log.Infof("Result: %f", ret)
	}
	os.Exit(0)
}
