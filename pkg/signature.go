package domain

type SignatureHandler interface {
	// Creates from a string its hexadecimal signature
	Create(data string) (string, error)

	// Verify that for the given data, the hexadecimal signature matches
	// An error is returned if it's not the case.
	Verify(data string, signature string) error
}
