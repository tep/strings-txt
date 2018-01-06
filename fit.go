package txt

import (
	"bufio"
	"bytes"
	"strings"
	"unicode"
)

var (
	TabWidth   = 2
	LineJoiner = "\n"
)

// Fit returns a transformed version of s with leading blank lines removed and
// a common width of leading whitespace removed from each subsequent.
// The intended use is to allow long, multi-line, raw strings in Go source code
// (using `back-quoted text`) to be indented in a normal fashion (as gofmt is
// wont to do) but the resulting string value to not be indented.
func Fit(s string) string {
	var (
		lc, ind int
		lines   []string
	)

	sc := bufio.NewScanner(bytes.NewBufferString(s))

	for sc.Scan() {
		l := strings.Replace(sc.Text(), "\t", strings.Repeat(" ", TabWidth), -1)
		if lc == 0 && isOnlySpaces(l) {
			continue
		}

		lc++
		ns := indexFirstNonSpace(l)

		if lc == 1 {
			ind = ns
		}

		li := ind
		if ns >= 0 && ns < li {
			li = ns
		}

		if len(l) >= li {
			l = l[li:]
		}

		lines = append(lines, strings.TrimRightFunc(l, unicode.IsSpace))
	}

	return strings.TrimSpace(strings.Join(lines, LineJoiner)) + "\n"
}

func isOnlySpaces(s string) bool {
	return (indexFirstNonSpace(s) == -1)
}

func indexFirstNonSpace(s string) int {
	return strings.IndexFunc(s, func(r rune) bool { return !unicode.IsSpace(r) })
}
