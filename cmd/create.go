/*
Copyright © 2023 Rudolfo Borges <oliveira.rudolfo@gmail.com>
*/
package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/rudolfoborges/create-go-api/internal/create"
	"github.com/rudolfoborges/create-go-api/internal/create/strategy"
	"github.com/spf13/cobra"
)

type answers struct {
	Name          string
	Handler       string
	Database      string
	Configuration string
	Logger        string
	DI            string
	WithDocker    bool
	Agree         bool
}

func (a *answers) input() *create.Input {
	return create.NewInput(a.Name, a.Handler, a.Logger)
}

var questions = []*survey.Question{
	{
		Name: "name",
		Prompt: &survey.Input{
			Message: "Name of the project",
		},
		Validate: survey.Required,
	},
	{
		Name: "handler",
		Prompt: &survey.Select{
			Message: "Choose a http handler",
			Options: []string{
				strategy.FIBER,
				strategy.NET_HTTP,
				strategy.CHI,
			},
			Default:  strategy.FIBER,
			PageSize: 3,
		},
		Validate: survey.Required,
	},
	{
		Name: "di",
		Prompt: &survey.Select{
			Message: "Choose a dependency injection strategy",
			Options: []string{
				"fx",
				"wire",
				"none",
			},
			Default:  "none",
			PageSize: 3,
		},
		Validate: survey.Required,
	},
	{
		Name: "database",
		Prompt: &survey.Select{
			Message: "Choose a database",
			Options: []string{
				"mysql",
				"postgres",
				"sqlite",
			},
			Default:  "mysql",
			PageSize: 3,
		},
		Validate: survey.Required,
	},
	{
		Name: "configuration",
		Prompt: &survey.Select{
			Message: "choose a configuration strategy",
			Options: []string{
				"viper",
				"godotenv",
			},
			Default:  "viper",
			PageSize: 2,
		},
		Validate: survey.Required,
	},
	{
		Name: "logger",
		Prompt: &survey.Select{
			Message: "choose a logger strategy",
			Options: []string{
				strategy.ZAP,
				strategy.ZEROLOG,
			},
			Default:  strategy.ZAP,
			PageSize: 2,
		},
		Validate: survey.Required,
	},
	{
		Name: "withDocker",
		Prompt: &survey.Confirm{
			Message: "create Dockerfile?",
			Default: true,
		},
	},
	{
		Name: "agree",
		Prompt: &survey.Confirm{
			Message: "If everything is okay, can I create this project for you?",
			Default: true,
		},
	},
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new golang api project",
	Long:  `Create a new golang api project`,
	RunE:  runCommand,
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func runCommand(cmd *cobra.Command, args []string) error {
	var a answers

	if err := survey.Ask(questions, &a); err != nil {
		panic(err)
	}

	return create.Execute(a.input())
}
