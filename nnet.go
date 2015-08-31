package main

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"source.datanerd.us/talpert/nnet/simplenet"
)

type Network interface {
	New(int)
	Run()
}

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
	sampleData := [][]float64{
		[]float64{1, 1},
		[]float64{1, 0},
		[]float64{0, 1},
		[]float64{0, 0},
	}
	// layout := [][][]float64{
	// 	{{0.5, 0.5}, {0.5, 0.5}},
	// 	{{0.5, 0.5}},
	// }
	net, err := simplenet.New(2, [][]float64{
		{0.5, 0.5},
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, inputVals := range sampleData {
		log.Infof("Inputs: %v", inputVals)
		ret, err := net.Run(inputVals)
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("Result: %f", ret)
	}
	os.Exit(0)
}
