package rules

import (
	"regexp"
	"strings"
)

var classRe = regexp.MustCompile(`class(Name)?="([^"]+)"`)

func Extract(line string) []string {
	matches := classRe.FindAllStringSubmatch(line, -1)
	var classes []string
	for _, m := range matches {
		parts := strings.Fields(m[2])
		classes = append(classes, parts...)
	}
	return classes
}
