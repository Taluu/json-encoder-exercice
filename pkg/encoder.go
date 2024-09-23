package domain

// Encoder represents a method of encoding (and decoding) any string fed to it.
type Encoder interface {
	// Encode encodes the passed string into the choosen algorithm.
	// If the encoder could not encode the data, an error will be sent.
	Encode(in string) (string, error)

	// Decode decodes the passed encoded string with the chosen algorithm.
	// If the decoder could not decode the data, an error will be returned.
	Decode(in string) (string, error)
}
