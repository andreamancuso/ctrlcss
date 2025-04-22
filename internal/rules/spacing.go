package rules

import (
	"fmt"
	"strings"
)

func spacingRule(class string, spacing map[string]string) string {
	switch {
	case strings.HasPrefix(class, "m-"):
		return prop(class, "m-", "margin", spacing)
	case strings.HasPrefix(class, "mt-"):
		return prop(class, "mt-", "margin-top", spacing)
	case strings.HasPrefix(class, "mr-"):
		return prop(class, "mr-", "margin-right", spacing)
	case strings.HasPrefix(class, "mb-"):
		return prop(class, "mb-", "margin-bottom", spacing)
	case strings.HasPrefix(class, "ml-"):
		return prop(class, "ml-", "margin-left", spacing)
	case strings.HasPrefix(class, "my-"):
		return multiProp(class, "my-", []string{"margin-top", "margin-bottom"}, spacing)
	case strings.HasPrefix(class, "mx-"):
		return multiProp(class, "mx-", []string{"margin-left", "margin-right"}, spacing)
	case strings.HasPrefix(class, "p-"):
		return prop(class, "p-", "padding", spacing)
	case strings.HasPrefix(class, "pt-"):
		return prop(class, "pt-", "padding-top", spacing)
	case strings.HasPrefix(class, "pr-"):
		return prop(class, "pr-", "padding-right", spacing)
	case strings.HasPrefix(class, "pb-"):
		return prop(class, "pb-", "padding-bottom", spacing)
	case strings.HasPrefix(class, "pl-"):
		return prop(class, "pl-", "padding-left", spacing)
	case strings.HasPrefix(class, "py-"):
		return multiProp(class, "py-", []string{"padding-top", "padding-bottom"}, spacing)
	case strings.HasPrefix(class, "px-"):
		return multiProp(class, "px-", []string{"padding-left", "padding-right"}, spacing)
	}
	return ""
}

func prop(class, prefix, cssProp string, spacing map[string]string) string {
	key := strings.TrimPrefix(class, prefix)
	if val, ok := spacing[key]; ok {
		return fmt.Sprintf(".%s { %s: %s; }", class, cssProp, val)
	}
	return ""
}

func multiProp(class, prefix string, props []string, spacing map[string]string) string {
	key := strings.TrimPrefix(class, prefix)
	if val, ok := spacing[key]; ok {
		decl := ""
		for _, p := range props {
			decl += fmt.Sprintf("%s: %s; ", p, val)
		}
		return fmt.Sprintf(".%s { %s }", class, strings.TrimSpace(decl))
	}
	return ""
}
