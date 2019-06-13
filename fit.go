// Copyright © 2018 Timothy E. Peoples <eng@toolman.org>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package txt

import (
	"bufio"
	"bytes"
	"strings"
	"unicode"
)

var (
	TabWidth      = 2
	LineSeparator = "\n"
)

// Fit the results of transforming the string s, on a line-by-line basis, in
// the following manner:
//
//   * All TAB characters are replaced with a TabWidth length sequence
//     of SPACE characters.
//
//   * The length of the leading whitespace for the first non-blank
//     line is recorded then the whitespace characters are removed.
//
//   * Leading whitespace, up to the recorded length)  on subsequent
//     lines is also removed.
//
//   * Trailing whitespace on ALL lines is removed.
//
//   * All lines are then joined using LineSeparator and a final
//     LineSeparator is appended to the end.
//
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

	return strings.TrimSpace(strings.Join(lines, LineSeparator)) + LineSeparator
}

func isOnlySpaces(s string) bool {
	return (indexFirstNonSpace(s) == -1)
}

func indexFirstNonSpace(s string) int {
	return strings.IndexFunc(s, func(r rune) bool { return !unicode.IsSpace(r) })
}
