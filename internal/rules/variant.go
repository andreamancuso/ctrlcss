package rules

import (
	"strings"
	"fmt"
)

var mediaBreakpoints = map[string]string{
	"sm":  "640px",
	"md":  "768px",
	"lg":  "1024px",
	"xl":  "1280px",
	"2xl": "1536px",
}

func WrapWithVariants(class, rule string) string {
	parts := strings.Split(class, ":")
	var chain []string
	for _, part := range parts[:len(parts)-1] {
		chain = append(chain, part)
	}

	selector := "." + escapeClass(class)
	body := rule[strings.Index(rule, "{"):] // get rule body (e.g. `{ color: red; }`)
	

	for i := len(chain) - 1; i >= 0; i-- {
		prefix := chain[i]

		switch prefix {
		case "hover":
			selector += ":hover"
		case "focus":
			selector += ":focus"
		case "active":
			selector += ":active"
		case "disabled":
			selector += ":disabled"
		case "visited":
			selector += ":visited"
		case "checked":
			selector += ":checked"
		case "focus-within":
			selector += ":focus-within"
		case "focus-visible":
			selector += ":focus-visible"
		case "first":
			selector += ":first-child"
		case "last":
			selector += ":last-child"
		case "odd":
			selector += ":nth-child(odd)"
		case "even":
			selector += ":nth-child(even)"
		default:
			if bp, ok := mediaBreakpoints[prefix]; ok {
				body = fmt.Sprintf(" @media (min-width: %s) { %s %s }", bp, selector, body)
				selector = "." + escapeClass(class)
			}
		}
	}

	if len(chain) > 0 && mediaBreakpoints[chain[0]] != "" {
		return body // already wrapped
	}
	return selector + body
}

func escapeClass(class string) string {
	return strings.ReplaceAll(class, ":", "\\:")
}
