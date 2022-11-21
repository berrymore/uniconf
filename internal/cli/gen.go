package cli

import (
	"fmt"
	"strings"
	"uniconf/internal/conf"
	"uniconf/internal/generator"
)

type Gen struct {
	generator *generator.Generator
}

func (g *Gen) Help() string {
	panic("implement me")
}

func (g *Gen) Run(args []string) int {
	if len(args) < 1 {
		fmt.Println("You should specify HCL file as a first argument")
		return -1
	}

	flagSet := parseFlagSet(args)

	config, err := conf.ReadFromFile(args[0])
	if err != nil {
		fmt.Println(err)
		return -1
	}

	outDir := strings.TrimRight(flagSet.get("out", "."), "/") + "/"

	err = g.generator.Run(config, createVars(flagSet), outDir)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	return 0
}

func (g *Gen) Synopsis() string {
	return "Generate configurations"
}

func createVars(flagSet flagSet) map[string]string {
	vars := map[string]string{}

	for key, val := range flagSet {
		if strings.HasPrefix(key, "var-") {
			vars[strings.TrimPrefix(key, "var-")] = val
		}
	}

	return vars
}
