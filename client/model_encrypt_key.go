package openapi

import (
	"crypto/rsa"
)

// EncryptKey struct for EncryptKey
type EncryptKey struct {
	PublicKey rsa.PublicKey `json:"public_key"`
}