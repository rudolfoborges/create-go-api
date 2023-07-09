package strategy

import "fmt"

type Zap struct{}
type Zerolog struct{}

func NewLoggerStrategy(logger string) (CreateStrategy, error) {
	switch logger {
	case ZAP:
		return &Zap{}, nil
	case ZEROLOG:
		return &Zerolog{}, nil
	default:
		return nil, fmt.Errorf("Logger %s not found", logger)
	}
}

func (c *Zap) Execute() error {
	return nil
}

func (c *Zerolog) Execute() error {
	return nil
}
