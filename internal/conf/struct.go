package conf

type Config struct {
	Metadata  *Metadata            `hcl:"metadata"`
	Entries   map[string]*Entry    `hcl:"entry"`
	Templates map[string]*Template `hcl:"template"`
}

func (c *Config) HasEntries() bool {
	return len(c.Entries) > 0
}

func (c *Config) HasTemplate(name string) bool {
	_, ok := c.Templates[name]

	return ok
}

type Template struct {
	Destination string `hcl:"destination"`
	Data        string `hcl:"data"`
}

type Entry struct {
	Name      string
	MakeDirs  map[string]*MakeDir `hcl:"mkdir"`
	Templates []string            `hcl:"templates"`
	Vars      map[string]any      `hcl:"vars"`
}

type MakeDir struct {
	Mode  int    `hcl:"mode"`
	Owner string `hcl:"owner"`
	Group string `hcl:"group"`
}

func (e *Entry) HasTemplates() bool {
	return len(e.Templates) > 0
}

type Metadata struct {
	Namespace string `hcl:"namespace"`
}
