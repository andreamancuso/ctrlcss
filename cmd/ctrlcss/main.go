// cmd/ctrlcss/main.go
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/andreamancuso/ctrlcss/internal/rules"
	"github.com/andreamancuso/ctrlcss/internal/theme"
)

func main() {
	t := theme.Load("ctrlcss.theme.toml")
	ruleSet := rules.New(t.Spacing, t.Colors)

	seen := make(map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		for _, cls := range rules.Extract(line) {
			if !seen[cls] {
				seen[cls] = true
				if css := ruleSet.Lookup(cls); css != "" {
					fmt.Println(css)
				}
			}
		}
	}

	for _, cls := range t.Safelist.Classes {
		if !seen[cls] {
			seen[cls] = true
			if css := ruleSet.Lookup(cls); css != "" {
				fmt.Println(css)
			}
		}
	}
}
