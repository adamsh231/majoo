package helper

import (
	"strconv"
	"unicode"
	"unicode/utf8"
)

func StringToInt(data string) int {
	res, err := strconv.Atoi(data)
	if err != nil {
		res = 0
	}

	return res
}

type buffer struct {
	r         []byte
	runeBytes [utf8.UTFMax]byte
}

func (b *buffer) indent() {
	if len(b.r) > 0 {
		b.r = append(b.r, '_')
	}
}

func (b *buffer) write(r rune) {
	if r < utf8.RuneSelf {
		b.r = append(b.r, byte(r))
		return
	}
	n := utf8.EncodeRune(b.runeBytes[0:], r)
	b.r = append(b.r, b.runeBytes[0:n]...)
}

func Underscore(s string) string {
	b := buffer{
		r: make([]byte, 0, len(s)),
	}
	var m rune
	var w bool
	for _, ch := range s {
		if unicode.IsUpper(ch) {
			if m != 0 {
				if !w {
					b.indent()
					w = true
				}
				b.write(m)
			}
			m = unicode.ToLower(ch)
		} else {
			if m != 0 {
				b.indent()
				b.write(m)
				m = 0
				w = false
			}
			b.write(ch)
		}
	}
	if m != 0 {
		if !w {
			b.indent()
		}
		b.write(m)
	}

	return string(b.r)
}