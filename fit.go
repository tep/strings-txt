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

func Fit(s string) string {
	var (
		lc, ind int
		lines   []string
	)

	sc := bufio.NewScanner(bytes.NewBufferString(s))

	for sc.Scan() {
		l := strings.Replace(sc.Text(), "\t", strings.Repeat(" ", TabWidth), -1)
		if lc == 0 && IsOnlySpaces(l) {
			continue
		}

		lc++
		ns := IndexFirstNonSpace(l)

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

func IsOnlySpaces(s string) bool {
	return (IndexFirstNonSpace(s) == -1)
}

func IndexFirstNonSpace(s string) int {
	return strings.IndexFunc(s, func(r rune) bool { return !unicode.IsSpace(r) })
}
