package domain

import (
	"errors"
	"fmt"
)

var (
	ErrCouldNotDecode           = errors.New("could not decode string")
	ErrCouldNotComputeSignature = errors.New("could not compute signature")
	ErrCouldNotDecodeSignature  = errors.New("could not decode the signature from hexadecimal string")
	ErrInvalidSignature         = errors.New("invalid signature")
)

func CouldNotDecode(err error) error {
	return fmt.Errorf("%w : %s", ErrCouldNotDecode, err)
}

func CouldNotComputeSignature(err error) error {
	return fmt.Errorf("%w : %s", ErrCouldNotComputeSignature, err)
}

func CouldNotDecodeSignature(err error) error {
	return fmt.Errorf("%w : %s", ErrCouldNotDecodeSignature, err)
}

func InvalidSignature(signature string) error {
	return fmt.Errorf("%w : %v is invalid", ErrInvalidSignature, signature)
}
