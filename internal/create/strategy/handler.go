package strategy

import "fmt"

type Fiber struct{}
type NetHttp struct{}
type Chi struct{}

func NewHandlerStrategy(handler string) (CreateStrategy, error) {
	switch handler {
	case FIBER:
		return &Fiber{}, nil
	case NET_HTTP:
		return &NetHttp{}, nil
	case CHI:
		return &Chi{}, nil
	default:
		return nil, fmt.Errorf("Handler %s not found", handler)
	}
}

func (c *Fiber) Execute() error {
	return nil
}

func (c *NetHttp) Execute() error {
	return nil
}

func (c *Chi) Execute() error {
	return nil
}
