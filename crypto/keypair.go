package crypto

import (
	"encoding/hex"
	"errors"
)

// KeyPair represents a private and public key pair.
type KeyPair struct {
	PrivateKey []byte
	PublicKey  []byte

	PrivKey PrivKey
	PubKey  PubKey
}

var (
	ErrPrivateKeySize = errors.New("private key length does not equal expected key length")
)

// Sign returns a cryptographic signature that is a signed hash of the message.
func (k *KeyPair) Sign(sp SignaturePolicy, hp HashPolicy, message []byte) ([]byte, error) {
	// if len(k.PrivateKey) != sp.PrivateKeySize() {
	// 	return nil, ErrPrivateKeySize
	// }

	message = hp.HashBytes(message)
	signature, err := k.PrivKey.Sign(message)

	if err != nil {
		return nil, err
	}

	//signature := sp.Sign(k.PrivateKey, message)
	return signature, nil
}

// PrivateKeyHex returns the hex representation of the private key.
func (k *KeyPair) PrivateKeyHex() string {
	return hex.EncodeToString(k.PrivKey.Bytes())
}

// PublicKeyHex returns the hex representation of the public key.
func (k *KeyPair) PublicKeyHex() string {
	return hex.EncodeToString(k.PubKey.Bytes())
}

// String returns the private and public key pair.
func (k *KeyPair) String() (string, string) {
	return k.PrivateKeyHex(), k.PublicKeyHex()
}

// FromPrivateKey returns a KeyPair given a signature policy and private key.
func FromPrivateKey(sp SignaturePolicy, privateKey string) (*KeyPair, error) {
	rawPrivateKey, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	return fromPrivateKeyBytes(sp, rawPrivateKey)
}

func fromPrivateKeyBytes(sp SignaturePolicy, rawPrivateKey []byte) (*KeyPair, error) {
	if len(rawPrivateKey) != sp.PrivateKeySize() {
		return nil, ErrPrivateKeySize
	}

	rawPublicKey, err := sp.PrivateToPublic(rawPrivateKey)
	if err != nil {
		return nil, err
	}

	keyPair := &KeyPair{
		PrivateKey: rawPrivateKey,
		PublicKey:  rawPublicKey,
	}

	return keyPair, nil
}

// Verify returns true if the given signature was generated using the given public key, message, signature policy, and hash policy.
func Verify(sp SignaturePolicy, hp HashPolicy, publicKey []byte, message []byte, signature []byte) bool {
	// Public key must be a set size.
	if len(publicKey) != sp.PublicKeySize() {
		return false
	}

	message = hp.HashBytes(message)
	return sp.Verify(publicKey, message, signature)
}
