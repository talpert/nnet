package main

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"source.datanerd.us/talpert/nnet/gridnet"
	"source.datanerd.us/talpert/nnet/simplenet"
	"time"
)

type Network interface {
	New(int, [][][]float64)
	Run([]float64) (float64, error)
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
		[]float64{1, 1, 1},
		[]float64{0, 1, 2},
		[]float64{2, 3, 3},
		[]float64{0, 0, 0},
	}
	// layout := [][][]float64{
	// 	{{0.5, 0.5}, {0.5, 0.5}},
	// 	{{0.5, 0.5}},
	// }
	weights := [][][]float64{
		{{0.1, 0.2, 0.0, 0.3}, {0.3, 0.4, 0.1, 0.2}},
		{{0.1, 0.2, 0.1}, {0.3, 0.4, 0.6}, {0.5, 0.6, 0.3}},
		{{0.7, 0.8, 0.9, 0.4}},
	}
	snet, err := simplenet.New(3, weights)
	if err != nil {
		log.Fatal(err)
	}
	for _, inputVals := range sampleData {
		start := time.Now()
		log.Infof("Inputs: %v", inputVals)
		ret, err := snet.Run(inputVals)
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("Result: %f", ret)
		log.Infof("SimpleNet Time: %v", time.Since(start))
	}

	gnet, gErr := gridnet.New(3, weights)
	if gErr != nil {
		log.Fatal(err)
	}
	for _, inputVals := range sampleData {
		start := time.Now()
		log.Infof("Inputs: %v", inputVals)
		ret, err := gnet.Run(inputVals)
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("Result: %f", ret)
		log.Infof("GridNet Time: %v", time.Since(start))
	}

	os.Exit(0)
}
