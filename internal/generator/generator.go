package generator

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"syscall"
	"uniconf/internal/conf"
	"uniconf/internal/docker"
	"uniconf/internal/renderer"
)

type Generator struct {
	Renderer renderer.Renderer
	Docker   docker.Docker
}

func (g *Generator) Run(config *conf.Config, vars map[string]string, outDir string) error {
	if !config.HasEntries() {
		return errors.New("config should have entries")
	}

	for _, entry := range config.Entries {
		fmt.Printf("Processing \"%s\" entry\n", entry.Name)

		if !entry.HasTemplates() {
			return errors.New("entry should have at least one template")
		}

		for path, mkdir := range entry.MakeDirs {
			syscall.Umask(0)
			err := os.MkdirAll(outDir+path, os.FileMode(mkdir.Mode))
			syscall.Umask(0022)
			if err != nil {
				fmt.Printf("Cannot create directory \"%s\"\n", path)
				continue
			}
		}

		for _, tplName := range entry.Templates {
			if !config.HasTemplate(tplName) {
				return errors.New(fmt.Sprintf("template \"%s\" is not defined", tplName))
			}

			containers, err := g.Docker.GetContainers()
			if err != nil {
				return err
			}

			osBag, err := createOsBag()
			if err != nil {
				return err
			}

			template := config.Templates[tplName]
			bag := &renderer.DataBag{
				Entry:      entry,
				Os:         osBag,
				Containers: containers,
				Vars:       vars,
			}

			out, err := g.Renderer.Render(template.Data, bag)
			if err != nil {
				return err
			}

			destOut, err := g.Renderer.Render(template.Destination, bag)
			if err != nil {
				return err
			}

			err = os.WriteFile(outDir+destOut, []byte(out), os.FileMode(0755))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func createOsBag() (*renderer.Os, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	return &renderer.Os{User: u}, nil
}
