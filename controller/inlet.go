package controller

import (
	"encoding/json"
	"github.com/kidoman/embd"
)

const (
	InletBucket = "inlets"
)

type Inlet struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Pin  int    `json:"pin"`
	Type string `json:"type"`
}

func (c *Controller) ReadFromInlet(i *Inlet) (int, error) {
	switch i.Type {
	case "digital":
		p, err := embd.NewDigitalPin(i.Pin)
		if err != nil {
			return 0, err
		}
		if err := p.SetDirection(embd.In); err != nil {
			return 0, err
		}
		return p.Read()
	case "analog":
		return c.state.adc.Read(i.Pin)
	}
	return 0, nil
}

func (c *Controller) GetInlet(id string) (Inlet, error) {
	var inlet Inlet
	return inlet, c.store.Get(InletBucket, id, &inlet)
}

func (c *Controller) ListInlets() (*[]interface{}, error) {
	fn := func(v []byte) (interface{}, error) {
		var inlet Inlet
		if err := json.Unmarshal(v, &inlet); err != nil {
			return nil, err
		}
		return inlet, nil
	}
	return c.store.List(InletBucket, fn)
}

func (c *Controller) CreateInlet(inlet Inlet) error {
	fn := func(id string) interface{} {
		inlet.ID = id
		return inlet
	}
	return c.store.Create(InletBucket, fn)
}

func (c *Controller) UpdateInlet(id string, payload interface{}) error {
	return c.store.Update(InletBucket, id, payload)
}

func (c *Controller) DeleteInlet(id string) error {
	return c.store.Delete(InletBucket, id)
}
