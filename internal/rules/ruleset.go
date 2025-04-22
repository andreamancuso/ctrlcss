package rules

import (
	"strings"
)

type RuleSet struct {
	Spacing map[string]string
	Colors  map[string]string
}

func New(spacing, colors map[string]string) RuleSet {
	return RuleSet{Spacing: spacing, Colors: colors}
}

func (r RuleSet) Lookup(class string) string {
	base := class
	if idx := strings.LastIndex(class, ":"); idx != -1 {
		base = class[idx+1:]
	}

	var rule string
	if css := spacingRule(base, r.Spacing); css != "" {
		rule = css
	}
	if css := colorRule(base, r.Colors); css != "" {
		rule = css
	}
	if css := layoutRule(base, r.Spacing); css != "" {
		rule = css
	}

	if rule == "" {
		return ""
	}

	if strings.Contains(class, ":") {
		return WrapWithVariants(class, rule)
	}
	return rule
}