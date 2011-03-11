// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/ and
// http://creativecommons.org/publicdomain/zero/1.0/legalcode

package optarg

import (
	"strings"
	"bytes"
	"regexp"
)

const (
	_ALIGN_LEFT = iota
	_ALIGN_CENTER
	_ALIGN_RIGHT
	_ALIGN_JUSTIFY
)

var reg_multilinewrap = regexp.MustCompile("[^a-zA-Z0-9,.]")

func multilineWrap(text string, linesize, leftmargin, rightmargin, alignment int) []string {
	lines := make([]string, 0)
	pad := ""

	for n := 0; n < leftmargin; n++ {
		pad += " "
	}

	linesize--

	if linesize < 1 {
		linesize = 80
	}

	wordboundary := 0
	size := linesize - leftmargin - rightmargin

	if len(text) <= size {
		lines = []string{align(text, pad, linesize, size, alignment)}
		return lines
	}

	for n := 0; n < len(text); n++ {
		if reg_multilinewrap.MatchString(text[n : n+1]) {
			wordboundary = n
		}

		if n > size {
			lines = append(lines,
				align(strings.TrimSpace(text[0:wordboundary]),
					pad, linesize, size, alignment))
			text = text[wordboundary:len(text)]
			n = 0
		}
	}

	lines = append(lines, align(strings.TrimSpace(text), pad, linesize, size, alignment))
	return lines
}

func align(v, pad string, linesize, size, alignment int) string {
	var data []byte
	buf := bytes.NewBuffer(data)

	switch alignment {
	case _ALIGN_LEFT:
		buf.WriteString(pad)
		buf.WriteString(v)

	case _ALIGN_RIGHT:
		diff := linesize - len(v) - len(pad)
		for n := 0; n < diff; n++ {
			buf.WriteByte(' ')
		}
		buf.WriteString(v)

	case _ALIGN_CENTER:
		diff := (size - len(v)) / 2
		buf.WriteString(pad)
		for n := 0; n < diff; n++ {
			buf.WriteByte(' ')
		}
		buf.WriteString(v)

	case _ALIGN_JUSTIFY:
		if strings.Index(v, " ") == -1 {
			buf.WriteString(pad)
			buf.WriteString(v)
			return buf.String()
		}

		diff := size - len(v)
		if diff == 0 {
			buf.WriteString(pad)
			buf.WriteString(v)
			break
		}

		spread := "  "
		for {
			if v = strings.Replace(v, spread[0:len(spread)-1], spread, -1); len(v) > size {
				break
			}
			spread += " "
		}

		for {
			if strings.Index(v, spread) == -1 {
				spread = spread[0 : len(spread)-1]
			}

			if v = strings.Replace(v, spread, spread[0:len(spread)-1], 1); len(v) <= size {
				break
			}
		}

		buf.WriteString(pad)
		buf.WriteString(v)
	}

	return buf.String()
}
