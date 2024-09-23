package domain

import (
	"errors"
	"fmt"
)

var (
	ErrCouldNotDecode = errors.New("could not decode string")
)

func CouldNotDecode(err error) error {
	return fmt.Errorf("%w : %s", ErrCouldNotDecode, err)
}
