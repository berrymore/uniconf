package renderer

import (
	"bytes"
	"fmt"
	"os/user"
	"strings"
	"text/template"
	"uniconf/internal/conf"
	"uniconf/internal/container"
)

type Renderer interface {
	Render(text string, bag *DataBag) (string, error)
}

type Os struct {
	User *user.User
}

type DataBag struct {
	Entry      *conf.Entry
	Os         *Os
	Containers map[string]*container.Container
	Vars       map[string]string
}

type BasicRenderer struct {
}

func (r *BasicRenderer) Render(text string, bag *DataBag) (string, error) {
	tpl := template.Must(template.New("").Funcs(getFuncMap(bag)).Parse(text))

	buf := bytes.NewBufferString("")

	err := tpl.Execute(buf, bag)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func CreateRenderer() Renderer {
	return &BasicRenderer{}
}

func getFuncMap(bag *DataBag) template.FuncMap {
	return map[string]any{
		"container": func(name string) *container.Container {
			if c, ok := bag.Containers[name]; ok {
				return c
			}

			fmt.Printf("Container \"%s\" not found\n", name)

			return nil
		},
		"var": func(name string) any {
			if c, ok := bag.Vars[name]; ok {
				return c
			}

			if c, ok := bag.Entry.Vars[name]; ok {
				return c
			}

			fmt.Printf("Variable \"%s\" not found\n", name)

			return ""
		},
		"replace": func(old string, new string, s string) string {
			return strings.Replace(s, old, new, -1)
		},
	}
}
