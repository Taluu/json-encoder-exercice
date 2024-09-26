package hmac

import "testing"

func TestCreate(t *testing.T) {
	const expected = "08ecaf9b9a0e67e0ebb152fbef192b0d14612726f097ce603335d1191d5a2fe8"
	const in = "this is a test"

	handler := &hmacHandler{key: "signature key"}
	sum, err := handler.Create(in)

	if err != nil {
		t.Fatalf("unexpected error : %s", err)
	}

	if sum != expected {
		t.Fatalf("signature is not valid (had %v, expected %v)", sum, expected)
	}
}

func TestVerify(t *testing.T) {
	const in = "this is a test"

	handler := &hmacHandler{key: "signature key"}

	t.Run("invalid signature", func(t *testing.T) {
		const invalidSignature = "08ecaf9b9a0e67e0ebb152fbef192b0d14712726f097ce603335d1191d6a2fe8"

		err := handler.Verify(in, invalidSignature)
		if err == nil {
			t.Fatalf("signature should not match")
		}
	})

	t.Run("valid signature", func(t *testing.T) {
		// this is a test
		const validSignature = "08ecaf9b9a0e67e0ebb152fbef192b0d14612726f097ce603335d1191d5a2fe8"

		err := handler.Verify(in, validSignature)
		if err != nil {
			t.Fatalf("signature should match")
		}
	})
}
