package base64

import (
	"encoding/base64"

	"github.com/Taluu/json-encoder-exercise/pkg/di"

	//lint:ignore ST1001 shared DSL
	. "github.com/Taluu/json-encoder-exercise/pkg"
)

func init() {
	di.Provide[Encoder](func() Encoder {
		return &base64Encoder{}
	})
}

type base64Encoder struct {
}

func (b *base64Encoder) Encode(in string) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(in)), nil
}

func (b *base64Encoder) Decode(in string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return "", CouldNotDecode(err)
	}

	return string(decoded), err
}

var _ Encoder = (*base64Encoder)(nil)
