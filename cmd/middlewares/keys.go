package middlewares

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

func GenerateKeyPair() (*rsa.PrivateKey, error) {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        return nil, err
    }
    return privatekey, nil
}

func Decrypt(cipherText string, privKey rsa.PrivateKey) (string, error) {
    ct, _ := base64.StdEncoding.DecodeString(cipherText)
    label := []byte("OAEP Encrypted")
    rng := rand.Reader
    plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, &privKey, ct, label)
    if err != nil {
        return "", err
    }
    return string(plaintext), nil
}