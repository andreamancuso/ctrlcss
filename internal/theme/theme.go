package theme

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Theme struct {
	Spacing map[string]string
	Colors  map[string]string
	Safelist struct {
		Classes []string
	}
}

func Load(path string) Theme {
	var theme Theme
	if _, err := toml.DecodeFile(path, &theme); err != nil {
		log.Fatalf("failed to load theme: %v", err)
	}
	return theme
}
