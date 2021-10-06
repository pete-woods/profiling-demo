package profiling_demo

import (
	"bufio"
	"io"
	"unicode"
)

func words(r io.Reader) (int, error) {
	words := 0
	inword := false
	r = bufio.NewReader(r)
	var buf [1]byte

	for {
		_, err := r.Read(buf[:])
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, err
		}

		r := rune(buf[0])

		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	return words, nil
}
