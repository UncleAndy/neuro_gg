package local

import (
	"log"
	"math/rand"
	"time"
)

type TestRewardFunc func() int64

type TrainMethod int

const (
	TrainMethodBackLinks = iota
	TrainMethodPowerFunc
)

type PowerFunc func(inputs, outputs []*Neuron) int64

type Net struct {
	inputs  []*Neuron
	layers  [][]*Neuron
	outputs []*Neuron

	trainFinishFlag bool

	trainMethod TrainMethod

	powerFunc PowerFunc
}

func NewLayerNet(input int, layers []int, output int) *Net {
	rand.Seed(time.Now().UnixNano())

	net := &Net{
		inputs:  make([]*Neuron, input),
		layers:  make([][]*Neuron, len(layers)),
		outputs: make([]*Neuron, output),
	}

	// Инициализация слоев нейронов
	for i := range net.layers {
		net.layers[i] = make([]*Neuron, layers[i])
		for j := range net.layers[i] {
			net.layers[i][j] = &Neuron{
				connectionCluster: NewCluster(),
				ExcitePeekLevel:   PeakLevel,
			}
		}
	}

	for j := range net.inputs {
		net.inputs[j] = &Neuron{
			connectionCluster: NewCluster(),
			ExcitePeekLevel:   PeakLevel,
		}
	}

	for j := range net.outputs {
		net.outputs[j] = &Neuron{
			connectionCluster: NewCluster(),
			ExcitePeekLevel:   PeakLevel,
		}
	}

	// Инициализация связей между слоями

	// От входного на первый слой
	fromInput := net.outputs
	if len(layers) > 0 {
		fromInput = net.layers[0]
	}

	net.fillRandomConnections(net.inputs, fromInput)

	// От последнего слоя на выходной слой
	if len(layers) > 0 {
		net.fillRandomConnections(net.layers[len(layers)-1], net.outputs)
	}

	// Между слоями
	if len(layers) > 1 {
		for l := 0; l < len(layers)-1; l++ {
			net.fillRandomConnections(net.layers[l], net.layers[l+1])
		}
	}

	return net
}

func (n *Net) SetPowerFunc(f PowerFunc) {
	n.powerFunc = f
}

func (n *Net) fillRandomConnections(fromLayer, toLayer []*Neuron) {
	for _, n := range fromLayer {
		for _, ln := range toLayer {
			conn := Connection{
				target: ln,
				Weight: int64(float64(MaxWeight*2)*rand.Float64()) - MaxWeight,
			}
			n.connectionCluster.Append(&conn)
			ln.backLinks = append(ln.backLinks, BackLink{
				From:       n,
				Connection: &conn,
			})
		}
	}
}

// Метод, реализующий процесс обучения
func (n *Net) Train() {
	switch n.trainMethod {
	case TrainMethodBackLinks:
		for !n.trainFinishFlag {
			n.TrainStepBackSignal()
		}
	case TrainMethodPowerFunc:
		if n.powerFunc == nil {
			log.Fatalln("Power func not set with using method train by power func")
		}
		for !n.trainFinishFlag {
			n.TrainStepByPower()
		}
	}
}

// Метод обучения на примерах с использованием метода обратного распространения ошибки
func (n *Net) TrainStepBackSignal() {

}

// Метод обучения по
func (n *Net) TrainStepByPower() {

}

// Метод, реализующий процесс работы сети (обработки и выдачи результата)
func (n *Net) Predict() {

}
