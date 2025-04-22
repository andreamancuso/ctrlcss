// ctrlcss.go — TOML-powered utility class to CSS rule emitter
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
)

type Theme struct {
	Spacing map[string]string
	Colors  map[string]string
}

func loadTheme(path string) Theme {
	var theme Theme
	_, err := toml.DecodeFile(path, &theme)
	if err != nil {
		log.Fatalf("Error loading theme: %v", err)
	}
	return theme
}

func lookupRule(class string, spacing map[string]string, colors map[string]string) string {
	switch {
	case strings.HasPrefix(class, "mt-"):
		key := strings.TrimPrefix(class, "mt-")
		if val, ok := spacing[key]; ok {
			return fmt.Sprintf(".%s { margin-top: %s; }", class, val)
		}
	case strings.HasPrefix(class, "py-"):
		key := strings.TrimPrefix(class, "py-")
		if val, ok := spacing[key]; ok {
			return fmt.Sprintf(".%s { padding-top: %s; padding-bottom: %s; }", class, val, val)
		}
	case strings.HasPrefix(class, "text-"):
		key := strings.TrimPrefix(class, "text-")
		if val, ok := colors[key]; ok {
			return fmt.Sprintf(".%s { color: %s; }", class, val)
		}
	case strings.HasPrefix(class, "bg-"):
		key := strings.TrimPrefix(class, "bg-")
		if val, ok := colors[key]; ok {
			return fmt.Sprintf(".%s { background-color: %s; }", class, val)
		}
	}
	return ""
}

func extractClasses(line string) []string {
	re := regexp.MustCompile(`class(Name)?="([^"]+)"`)
	matches := re.FindAllStringSubmatch(line, -1)
	var classes []string
	for _, m := range matches {
		parts := strings.Fields(m[2])
		classes = append(classes, parts...)
	}
	return classes
}

func main() {
	theme := loadTheme("ctrlcss.theme.toml")
	seen := make(map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		for _, cls := range extractClasses(line) {
			if !seen[cls] {
				seen[cls] = true
				if rule := lookupRule(cls, theme.Spacing, theme.Colors); rule != "" {
					fmt.Println(rule)
				}
			}
		}
	}
}
