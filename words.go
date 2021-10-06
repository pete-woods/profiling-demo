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
	for {
		r, err := readbyte(r)
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, err
		}

		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	return words, nil
}

func readbyte(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}
