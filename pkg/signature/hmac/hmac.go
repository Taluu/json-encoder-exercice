package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	domain "github.com/Taluu/json-encoder-exercise/pkg"
)

type hmacHandler struct {
	key string
}

func (h *hmacHandler) Create(data string) (string, error) {
	hmacSum, err := h.compute(data)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hmacSum), nil
}

func (h *hmacHandler) Verify(data string, signature string) error {
	decodedSum, err := hex.DecodeString(signature)
	if err != nil {
		return domain.CouldNotDecodeSignature(err)
	}

	computedSum, err := h.compute(data)
	if err != nil {
		return err
	}

	if !hmac.Equal(decodedSum, computedSum) {
		return domain.InvalidSignature(signature)
	}

	return nil
}

func (h *hmacHandler) compute(data string) ([]byte, error) {
	hash := hmac.New(sha256.New, []byte(h.key))
	if _, err := hash.Write([]byte(data)); err != nil {
		return nil, domain.CouldNotComputeSignature(err)
	}

	return hash.Sum(nil), nil
}

func NewSignatureHandler(key string) domain.SignatureHandler {
	return &hmacHandler{
		key: key,
	}
}
