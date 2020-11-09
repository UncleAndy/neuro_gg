package local

import ngg "github.com/UncleAndy/neurogg"

const (
	OutputSignal = int64(32000)
	PeakLevel    = int64(160000)
	MaxWeight    = int64(32000)
)

type BackLink struct {
	From       *Neuron
	Connection *Connection
}

type BackLinks []BackLink

type Neuron struct {
	connectionCluster ngg.ConnectorCluster
	backLinks         BackLinks

	ExcitePeekLevel int64

	exciteBuffer int64

	lastTick int64
}

func (n *Neuron) SetConnectorCluster(cc ngg.ConnectorCluster) {
	n.connectionCluster = cc
}

func (n *Neuron) GetConnectorCluster() ngg.ConnectorCluster {
	return n.connectionCluster
}

func (n *Neuron) Excite(e ngg.Excitation) {
	s, ok := e.(*Signal)
	if ok {
		n.lastTick += s.Value
	}
}

func (n *Neuron) Tick(tick interface{}) {
	t, ok := tick.(int64)
	if ok {
		if t != n.lastTick {
			n.processExcite()
			n.lastTick = t
		}
	}
}

func (n *Neuron) processExcite() {
	if n.exciteBuffer >= n.ExcitePeekLevel {
		signal := Signal{Value: OutputSignal}
		n.connectionCluster.ProcessExcitation(&signal)
	}
	n.exciteBuffer = 0
}
