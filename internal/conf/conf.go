package conf

import (
	"github.com/hashicorp/hcl"
	"os"
)

func ReadFromFile(name string) (*Config, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = hcl.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	normalize(cfg)

	return cfg, nil
}

func normalize(config *Config) {
	for name, entry := range config.Entries {
		entry.Name = name

		for _, mkdir := range entry.MakeDirs {
			mkdir.Mode = zeroIntTo(mkdir.Mode, 0755)
		}
	}
}

func zeroIntTo(old int, new int) int {
	if old == 0 {
		return new
	}

	return old
}
