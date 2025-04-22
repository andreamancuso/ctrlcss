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

var position = map[string]string{
	"absolute": "position: absolute;",
	"relative": "position: relative;",
	"fixed":    "position: fixed;",
	"sticky":   "position: sticky;",
	"static":   "position: static;",
}

var overflowClasses = map[string]string{
	"overflow-hidden":   "overflow: hidden;",
	"overflow-auto":     "overflow: auto;",
	"overflow-scroll":   "overflow: scroll;",
	"overflow-x-hidden": "overflow-x: hidden;",
	"overflow-y-auto":   "overflow-y: auto;",
}

var textOverflowClasses = map[string]string{
	"truncate":     "overflow: hidden; text-overflow: ellipsis; white-space: nowrap;",
	"text-ellipsis": "text-overflow: ellipsis;",
	"text-clip":     "text-overflow: clip;",
}

var objectFitClasses = map[string]string{
	"object-contain": "object-fit: contain;",
	"object-cover":   "object-fit: cover;",
	"object-fill":    "object-fit: fill;",
	"object-none":    "object-fit: none;",
	"object-scale-down": "object-fit: scale-down;",
}

var autoMap = map[string]string{
	"auto-cols-fr":  "grid-auto-columns: minmax(0, 1fr);",
	"auto-cols-min": "grid-auto-columns: min-content;",
	"auto-cols-max": "grid-auto-columns: max-content;",
	"auto-rows-fr":  "grid-auto-rows: minmax(0, 1fr);",
	"auto-rows-min": "grid-auto-rows: min-content;",
	"auto-rows-max": "grid-auto-rows: max-content;",
}

var flowMap = map[string]string{
	"grid-flow-row":        "grid-auto-flow: row;",
	"grid-flow-col":        "grid-auto-flow: column;",
	"grid-flow-row-dense":  "grid-auto-flow: row dense;",
	"grid-flow-col-dense":  "grid-auto-flow: column dense;",
}

var selfAlign = map[string]string{
	"justify-self-start":  "justify-self: start;",
	"justify-self-center": "justify-self: center;",
	"justify-self-end":    "justify-self: end;",
	"justify-self-stretch": "justify-self: stretch;",

	"self-auto":   "align-self: auto;",
	"self-start":  "align-self: flex-start;",
	"self-center": "align-self: center;",
	"self-end":    "align-self: flex-end;",
	"self-stretch": "align-self: stretch;",
}

var placeMap = map[string]string{
	"place-items-start":   "place-items: start;",
	"place-items-center":  "place-items: center;",
	"place-items-end":     "place-items: end;",
	"place-items-stretch": "place-items: stretch;",

	"place-content-start":   "place-content: start;",
	"place-content-center":  "place-content: center;",
	"place-content-end":     "place-content: end;",
	"place-content-between": "place-content: space-between;",
	"place-content-around":  "place-content: space-around;",
	"place-content-evenly":  "place-content: space-evenly;",
}

var cursorMap = map[string]string{
	"cursor-pointer": "pointer",
	"cursor-default": "default",
	"cursor-wait":    "wait",
	"cursor-text":    "text",
	"cursor-move":    "move",
	"cursor-not-allowed": "not-allowed",
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
	if decl, ok := position[class]; ok {
		return fmt.Sprintf(".%s { %s }", class, decl)
	}

	if decl, ok := overflowClasses[class]; ok {
		return fmt.Sprintf(".%s { %s }", class, decl)
	}

	if decl, ok := textOverflowClasses[class]; ok {
		return fmt.Sprintf(".%s { %s }", class, decl)
	}

	if decl, ok := objectFitClasses[class]; ok {
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

	if strings.HasPrefix(class, "w-") {
		key := strings.TrimPrefix(class, "w-")
		switch key {
		case "full":
			return fmt.Sprintf(".%s { width: 100%%; }", class)
		case "screen":
			return fmt.Sprintf(".%s { width: 100vw; }", class)
		default:
			if val, ok := spacing[key]; ok {
				return fmt.Sprintf(".%s { width: %s; }", class, val)
			}
		}
	}
	
	if strings.HasPrefix(class, "h-") {
		key := strings.TrimPrefix(class, "h-")
		switch key {
		case "full":
			return fmt.Sprintf(".%s { height: 100%%; }", class)
		case "screen":
			return fmt.Sprintf(".%s { height: 100vh; }", class)
		default:
			if val, ok := spacing[key]; ok {
				return fmt.Sprintf(".%s { height: %s; }", class, val)
			}
		}
	}

	if strings.HasPrefix(class, "z-") {
		key := strings.TrimPrefix(class, "z-")
		return fmt.Sprintf(".%s { z-index: %s; }", class, key)
	}

	if strings.HasPrefix(class, "top-") {
		key := strings.TrimPrefix(class, "top-")
		if val, ok := spacing[key]; ok {
			return fmt.Sprintf(".%s { top: %s; }", class, val)
		}
	}
	if strings.HasPrefix(class, "bottom-") {
		key := strings.TrimPrefix(class, "bottom-")
		if val, ok := spacing[key]; ok {
			return fmt.Sprintf(".%s { bottom: %s; }", class, val)
		}
	}
	if strings.HasPrefix(class, "left-") {
		key := strings.TrimPrefix(class, "left-")
		if val, ok := spacing[key]; ok {
			return fmt.Sprintf(".%s { left: %s; }", class, val)
		}
	}
	if strings.HasPrefix(class, "right-") {
		key := strings.TrimPrefix(class, "right-")
		if val, ok := spacing[key]; ok {
			return fmt.Sprintf(".%s { right: %s; }", class, val)
		}
	}
	if strings.HasPrefix(class, "inset-") {
		key := strings.TrimPrefix(class, "inset-")
		if val, ok := spacing[key]; ok {
			return fmt.Sprintf(".%s { top: %s; right: %s; bottom: %s; left: %s; }", class, val, val, val, val)
		}
	}
	if strings.HasPrefix(class, "inset-x-") {
		key := strings.TrimPrefix(class, "inset-x-")
		if val, ok := spacing[key]; ok {
			return fmt.Sprintf(".%s { left: %s; right: %s; }", class, val, val)
		}
	}
	if strings.HasPrefix(class, "inset-y-") {
		key := strings.TrimPrefix(class, "inset-y-")
		if val, ok := spacing[key]; ok {
			return fmt.Sprintf(".%s { top: %s; bottom: %s; }", class, val, val)
		}
	}

	if strings.HasPrefix(class, "min-w-") {
		key := strings.TrimPrefix(class, "min-w-")
		if val, ok := spacing[key]; ok {
			return fmt.Sprintf(".%s { min-width: %s; }", class, val)
		}
	}
	if strings.HasPrefix(class, "max-w-") {
		key := strings.TrimPrefix(class, "max-w-")
		if val, ok := spacing[key]; ok {
			return fmt.Sprintf(".%s { max-width: %s; }", class, val)
		}
	}
	if strings.HasPrefix(class, "min-h-") {
		key := strings.TrimPrefix(class, "min-h-")
		if val, ok := spacing[key]; ok {
			return fmt.Sprintf(".%s { min-height: %s; }", class, val)
		}
	}
	if strings.HasPrefix(class, "max-h-") {
		key := strings.TrimPrefix(class, "max-h-")
		if val, ok := spacing[key]; ok {
			return fmt.Sprintf(".%s { max-height: %s; }", class, val)
		}
	}

	if strings.HasPrefix(class, "grid-cols-") {
		n := strings.TrimPrefix(class, "grid-cols-")
		return fmt.Sprintf(".%s { grid-template-columns: repeat(%s, minmax(0, 1fr)); }", class, n)
	}

	if strings.HasPrefix(class, "col-span-") {
		n := strings.TrimPrefix(class, "col-span-")
		return fmt.Sprintf(".%s { grid-column: span %s / span %s; }", class, n, n)
	}
	
	if strings.HasPrefix(class, "row-span-") {
		n := strings.TrimPrefix(class, "row-span-")
		return fmt.Sprintf(".%s { grid-row: span %s / span %s; }", class, n, n)
	}

	if strings.HasPrefix(class, "grid-rows-") {
		n := strings.TrimPrefix(class, "grid-rows-")
		return fmt.Sprintf(".%s { grid-template-rows: repeat(%s, minmax(0, 1fr)); }", class, n)
	}
	
	if decl, ok := autoMap[class]; ok {
		return fmt.Sprintf(".%s { %s }", class, decl)
	}

	if decl, ok := flowMap[class]; ok {
		return fmt.Sprintf(".%s { %s }", class, decl)
	}

	if decl, ok := selfAlign[class]; ok {
		return fmt.Sprintf(".%s { %s }", class, decl)
	}
	
	if decl, ok := placeMap[class]; ok {
		return fmt.Sprintf(".%s { %s }", class, decl)
	}

	if strings.HasPrefix(class, "order-") {
		key := strings.TrimPrefix(class, "order-")
		switch key {
		case "first":
			return fmt.Sprintf(".%s { order: -9999; }", class)
		case "last":
			return fmt.Sprintf(".%s { order: 9999; }", class)
		case "none":
			return fmt.Sprintf(".%s { order: 0; }", class)
		default:
			return fmt.Sprintf(".%s { order: %s; }", class, key)
		}
	}

	if val, ok := cursorMap[class]; ok {
		return fmt.Sprintf(".%s { cursor: %s; }", class, val)
	}

	return ""
}
