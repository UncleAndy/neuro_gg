package local

import ngg "github.com/UncleAndy/neurogg"

type Cluster struct {
	connectors []ngg.Connector
}

func NewCluster() *Cluster {
	return &Cluster{connectors: make([]ngg.Connector, 0)}
}

func (c *Cluster) Append(cn ngg.Connector) {
	c.connectors = append(c.connectors, cn)
}

func (c *Cluster) ProcessExcitation(e ngg.Excitation) {
	for _, c := range c.connectors {
		c.ProcessExcitation(e)
	}
}

type Connection struct {
	target ngg.Excitable
	Weight int64
}

func (c *Connection) SetTarget(e ngg.Excitable) {
	c.target = e
}

func (c *Connection) ProcessExcitation(e ngg.Excitation) {
	if c.target != nil {
		c.ProcessExcitation(e)
	}
}
