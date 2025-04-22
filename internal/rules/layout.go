package rules

import (
	"fmt"
	"strings"
)


var displayClasses = map[string]string{
	"block":       "display: block;",
	"inline":      "display: inline;",
	"inline-block": "display: inline-block;",
	"flex":        "display: flex;",
	"inline-flex": "display: inline-flex;",
	"grid":        "display: grid;",
	"inline-grid": "display: inline-grid;",
	"hidden":      "display: none;",
}

var flexDirection = map[string]string{
	"flex-row": "flex-direction: row;",
	"flex-col": "flex-direction: column;",
}

var flexWrap = map[string]string{
	"flex-wrap":   "flex-wrap: wrap;",
	"flex-nowrap": "flex-wrap: nowrap;",
}

var justifyContent = map[string]string{
	"justify-start":   "justify-content: flex-start;",
	"justify-center":  "justify-content: center;",
	"justify-between": "justify-content: space-between;",
	"justify-end":     "justify-content: flex-end;",
}

var alignItems = map[string]string{
	"items-start":   "align-items: flex-start;",
	"items-center":  "align-items: center;",
	"items-end":     "align-items: flex-end;",
	"items-stretch": "align-items: stretch;",
}

var contentAlign = map[string]string{
	"content-start":   "align-content: flex-start;",
	"content-center":  "align-content: center;",
	"content-between": "align-content: space-between;",
	"content-around":  "align-content: space-around;",
}


func layoutRule(class string, spacing map[string]string) string {
	if decl, ok := displayClasses[class]; ok {
		return fmt.Sprintf(".%s { %s }", class, decl)
	}
	if decl, ok := flexDirection[class]; ok {
		return fmt.Sprintf(".%s { %s }", class, decl)
	}
	if decl, ok := flexWrap[class]; ok {
		return fmt.Sprintf(".%s { %s }", class, decl)
	}
	if decl, ok := justifyContent[class]; ok {
		return fmt.Sprintf(".%s { %s }", class, decl)
	}
	if decl, ok := alignItems[class]; ok {
		return fmt.Sprintf(".%s { %s }", class, decl)
	}
	if decl, ok := contentAlign[class]; ok {
		return fmt.Sprintf(".%s { %s }", class, decl)
	}
	if strings.HasPrefix(class, "gap-") {
		key := strings.TrimPrefix(class, "gap-")
		if val, ok := spacing[key]; ok {
			return fmt.Sprintf(".%s { gap: %s; }", class, val)
		}
	}
	if strings.HasPrefix(class, "gap-x-") {
		key := strings.TrimPrefix(class, "gap-x-")
		if val, ok := spacing[key]; ok {
			return fmt.Sprintf(".%s { column-gap: %s; }", class, val)
		}
	}
	if strings.HasPrefix(class, "gap-y-") {
		key := strings.TrimPrefix(class, "gap-y-")
		if val, ok := spacing[key]; ok {
			return fmt.Sprintf(".%s { row-gap: %s; }", class, val)
		}
	}
	return ""
}
