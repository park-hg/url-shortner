package pkg_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"traffic-reporter/internal/pkg"
)

func TestEncodeAndDecode(t *testing.T) {
	for _, ts := range []struct {
		original int64
		encoded  string
	}{
		{
			original: int64(0),
			encoded:  "0",
		},
		{
			original: int64(1152921504606846975),
			encoded:  "1NAOLcol8qV",
		},
	} {
		enc := pkg.NewBase62Encoder()
		encoded := enc.Encode(ts.original)
		assert.Equal(t, ts.encoded, encoded)

		decoded, err := enc.Decode(encoded)
		assert.NoError(t, err)
		assert.Equal(t, ts.original, decoded)
	}
}

func TestDecodeInvalidChar(t *testing.T) {
	invalid := "8919ABDC-="
	enc := pkg.NewBase62Encoder()
	_, err := enc.Decode(invalid)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, pkg.ErrInvalidCharacter))
}
