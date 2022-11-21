package cli

import "strings"

type flagSet map[string]string

func (f flagSet) has(name string) bool {
	_, ok := f[name]

	return ok
}

func (f flagSet) get(name string, def string) string {
	if f.has(name) {
		return f[name]
	}

	return def
}

func parseFlagSet(args []string) flagSet {
	set := flagSet{}

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			parts := strings.Split(strings.TrimLeft(arg, "-"), "=")

			set[parts[0]] = parts[1]
		}
	}

	return set
}
