package bfv

import (
	"github.com/ldsec/lattigo/v2/ring"
	"github.com/ldsec/lattigo/v2/rlwe"
)

// Decryptor is an interface for decryptors
type Decryptor interface {
	// DecryptNew decrypts the input ciphertext and returns the result on a new
	// plaintext.
	DecryptNew(ciphertext *Ciphertext) *Plaintext

	// Decrypt decrypts the input ciphertext and returns the result on the
	// provided receiver plaintext.
	Decrypt(ciphertext *Ciphertext, plaintext *Plaintext)
}

// decryptor is a structure used to decrypt ciphertexts. It stores the secret-key.
type decryptor struct {
	params   Parameters
	ringQ    *ring.Ring
	sk       *rlwe.SecretKey
	polypool *ring.Poly
}

// NewDecryptor creates a new Decryptor from the parameters with the secret-key
// given as input.
func NewDecryptor(params Parameters, sk *rlwe.SecretKey) Decryptor {

	ringQ := params.RingQ()

	return &decryptor{
		params:   params,
		ringQ:    ringQ,
		sk:       sk,
		polypool: ringQ.NewPoly(),
	}
}

func (decryptor *decryptor) DecryptNew(ciphertext *Ciphertext) *Plaintext {
	p := NewPlaintext(decryptor.params)
	decryptor.Decrypt(ciphertext, p)
	return p
}

func (decryptor *decryptor) Decrypt(ciphertext *Ciphertext, p *Plaintext) {

	ringQ := decryptor.ringQ
	tmp := decryptor.polypool

	ringQ.NTTLazy(ciphertext.Value[ciphertext.Degree()], p.value)

	for i := ciphertext.Degree(); i > 0; i-- {
		ringQ.MulCoeffsMontgomery(p.value, decryptor.sk.Value, p.value)
		ringQ.NTTLazy(ciphertext.Value[i-1], tmp)
		ringQ.Add(p.value, tmp, p.value)

		if i&3 == 3 {
			ringQ.Reduce(p.value, p.value)
		}
	}

	if (ciphertext.Degree())&3 != 3 {
		ringQ.Reduce(p.value, p.value)
	}

	ringQ.InvNTT(p.value, p.value)
}
