package base64

import (
	"errors"
	"testing"

	. "github.com/Taluu/json-encoder-exercise/pkg"
)

func TestEncode(t *testing.T) {
	const in = "this is a test"
	const expected = "dGhpcyBpcyBhIHRlc3Q="

	encoder := &base64Encoder{}
	result, err := encoder.Encode(in)

	if err != nil {
		t.Errorf("could not property encode the text : %s", err)
		t.FailNow()
	}

	if result != expected {
		t.Fatalf("did not get the expected result %v, got %v", expected, result)
	}
}

func TestDecode(t *testing.T) {
	const in = "dGhpcyBpcyBhIHRlc3Q="
	const expected = "this is a test"

	encoder := &base64Encoder{}
	result, err := encoder.Decode(in)

	if err != nil {
		t.Errorf("could not property decode the text : %s", err)
		t.FailNow()
	}

	if result != expected {
		t.Fatalf("did not get the expected result %v, got %v", expected, result)
	}
}

func TestDecodeFailure(t *testing.T) {
	const in = "dGhpcyBpcyBhIHRlc3Q" // missing last =

	encoder := &base64Encoder{}
	_, err := encoder.Decode(in)

	if err == nil {
		t.Fatalf("an error should have been triggered")
	}

	if !errors.Is(err, ErrCouldNotDecode) {
		t.Fatalf("unexpected error %s", err)
	}
}
