package create

type Input struct {
	Name string
}

func NewInput(name string) *Input {
	return &Input{
		Name: name,
	}
}

// Execute create a new project with
func Execute(input *Input) error {
	return nil
}
