package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
)

func Encrypt(data string, secretKey string) (string, error) {
	// Derive a 32-byte key from the secret key using SHA-256
	key := sha256.Sum256([]byte(secretKey))

	// Create AES cipher block
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}
	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create a random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	// Encrypt and authenticate
	ciphertext := gcm.Seal(nonce, nonce, []byte(data), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(data string, secretKey string) (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	key := sha256.Sum256([]byte(secretKey))

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	if len(decodedData) < nonceSize {
		return "", errors.New("ciphertext too short")
	}
	nonce, ciphertext := decodedData[:nonceSize], decodedData[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, []byte(ciphertext), nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
