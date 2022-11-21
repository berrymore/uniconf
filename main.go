package main

import (
	"fmt"
	"github.com/docker/docker/client"
	"github.com/mitchellh/cli"
	"os"
	intCli "uniconf/internal/cli"
	"uniconf/internal/docker"
	"uniconf/internal/generator"
	"uniconf/internal/renderer"
)

func main() {
	c := cli.NewCLI("uniconf", "1.0.0")

	dock, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	gen := &generator.Generator{
		Renderer: renderer.CreateRenderer(),
		Docker:   docker.CreateDocker(dock),
	}

	c.Args = os.Args[1:]
	c.Commands = intCli.CreateCommandFactory(gen)

	exitCode, err := c.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	os.Exit(exitCode)
}
