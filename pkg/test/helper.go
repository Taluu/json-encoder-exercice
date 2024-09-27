package test

import (
	domain "github.com/Taluu/json-encoder-exercise/pkg"
	"github.com/Taluu/json-encoder-exercise/pkg/di"
	_ "github.com/Taluu/json-encoder-exercise/pkg/encoder/base64"
	"github.com/Taluu/json-encoder-exercise/pkg/signature/hmac"
)

func init() {
	di.Provide[domain.SignatureHandler](func() domain.SignatureHandler {
		return hmac.NewSignatureHandler("test")
	})
}
