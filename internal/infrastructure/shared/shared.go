package shared

import (
	"regexp"
	"strings"
)

var spaceEater = regexp.MustCompile(`\s+`)

func compact(queryStmt string) string {
	return spaceEater.ReplaceAllString(strings.TrimSpace(queryStmt), "")
}
