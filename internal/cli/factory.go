package cli

import (
	"github.com/mitchellh/cli"
	"uniconf/internal/generator"
)

func CreateCommandFactory(gen *generator.Generator) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"gen": func() (cli.Command, error) {
			return &Gen{generator: gen}, nil
		},
	}
}
