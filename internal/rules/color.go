package rules

import (
	"fmt"
	"strings"
)

func colorRule(class string, colors map[string]string) string {
	if strings.HasPrefix(class, "text-") {
		key := strings.TrimPrefix(class, "text-")
		if val, ok := colors[key]; ok {
			return fmt.Sprintf(".%s { color: %s; }", class, val)
		}
	}
	if strings.HasPrefix(class, "bg-") {
		key := strings.TrimPrefix(class, "bg-")
		if val, ok := colors[key]; ok {
			return fmt.Sprintf(".%s { background-color: %s; }", class, val)
		}
	}
	return ""
}
