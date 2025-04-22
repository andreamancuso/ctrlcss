package rules

type RuleSet struct {
	Spacing map[string]string
	Colors  map[string]string
}

func New(spacing, colors map[string]string) RuleSet {
	return RuleSet{Spacing: spacing, Colors: colors}
}

func (r RuleSet) Lookup(class string) string {
	if css := spacingRule(class, r.Spacing); css != "" {
		return css
	}
	if css := colorRule(class, r.Colors); css != "" {
		return css
	}
	return ""
}
