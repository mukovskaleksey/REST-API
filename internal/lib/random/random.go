package random

import (
	"math/rand"
)

func NewRandomString(aliasLength int) string {
	rng := rand.New(rand.NewSource(int64(rand.Uint64())))

	const alphabet = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"

	b := make([]byte, aliasLength)

	for i := range b {
		b[i] = alphabet[rng.Intn(len(alphabet))]
	}

	return string(b)

}
