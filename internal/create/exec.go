package create

import (
	"fmt"
	"os"
	"sync"

	"github.com/go-git/go-git/v5"
	"github.com/rudolfoborges/create-go-api/internal/create/strategy"
)

type Input struct {
	name    string
	handler string
	logger  string
}

func NewInput(name, handler, logger string) *Input {
	return &Input{
		name:    name,
		handler: handler,
		logger:  logger,
	}
}

// Execute create a new project with
func Execute(input *Input) error {
	if err := cloneTemplate(input.name); err != nil {
		return err
	}

	if err := os.RemoveAll(fmt.Sprintf("./%s/.git", input.name)); err != nil {
		return err
	}

	var commands []strategy.CreateStrategy

	handler, err := strategy.NewHandlerStrategy(input.handler)
	if err != nil {
		return err
	}

	logger, err := strategy.NewLoggerStrategy(input.logger)
	if err != nil {
		return err
	}

	commands = append(commands, handler)
	commands = append(commands, logger)

	var wg sync.WaitGroup

	for _, c := range commands {
		go func(command strategy.CreateStrategy) {
			defer wg.Done()
			fmt.Println(command)
		}(c)
	}

	wg.Add(len(commands))
	wg.Wait()

	return nil
}

func cloneTemplate(name string) error {
	_, err := git.PlainClone(fmt.Sprintf("./%s", name), false, &git.CloneOptions{
		URL:               "https://github.com/rudolfoborges/create-go-api-template-blank.git",
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          os.Stdout,
	})

	if err != nil {
		return err
	}

	return nil
}
