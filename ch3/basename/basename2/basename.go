package basename2

import "strings"

// Basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func Basename(s string) string {
	i := strings.LastIndex(s, "/") // index of slash
	s = s[i+1:]
	i = strings.LastIndex(s, ".") // index of dot
	s = s[:i]

	return s
}
