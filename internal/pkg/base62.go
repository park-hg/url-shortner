package pkg

import (
	"errors"
	"fmt"
	"strings"
)

type Encoder interface {
	Encode(int64) string
	Decode(string) (int64, error)
}

type Base62Encoder struct{}

const (
	base          = int64(62)
	base62Charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

var (
	ErrInvalidCharacter = errors.New("invalid character")
)

func NewBase62Encoder() *Base62Encoder {
	return &Base62Encoder{}
}

func (e *Base62Encoder) Encode(n int64) string {
	if n == 0 {
		return "0"
	}

	var encoded string
	for n > 0 {
		r := n % base
		encoded += string([]rune(base62Charset)[r])
	}

	return reverse(encoded)
}

func (e *Base62Encoder) Decode(encoded string) (int64, error) {
	var n int64

	for _, char := range encoded {
		index := strings.IndexRune(base62Charset, char)
		if index == -1 {
			return 0, errors.Join(ErrInvalidCharacter, fmt.Errorf("input char: %s", string(char)))
		}
		n = n*base + int64(index)
	}

	return n, nil
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
