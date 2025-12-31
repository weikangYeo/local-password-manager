package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
)

func Encrypt(data string, secretKey string) (string, error) {
	// Generate random salt (16 bytes)
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}

	// Use PBKDF2 to derive key from password + salt
	key, err := pbkdf2.Key(sha256.New, secretKey, salt, 100000, 32)
	if err != nil {
		return "", err
	}

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	// Encrypt and authenticate
	ciphertext := gcm.Seal(nonce, nonce, []byte(data), nil)

	// Prepend salt to ciphertext: [salt || ciphertext]
	result := append(salt, ciphertext...)
	return base64.StdEncoding.EncodeToString(result), nil
}

func Decrypt(data string, secretKey string) (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	// Extract salt (first 16 bytes)
	if len(decodedData) < 16 {
		return "", errors.New("invalid ciphertext")
	}
	salt := decodedData[:16]
	ciphertextWithNonce := decodedData[16:]

	// Derive same key using extracted salt
	key, err := pbkdf2.Key(sha256.New, secretKey, salt, 100000, 32)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertextWithNonce) < nonceSize {
		return "", errors.New("ciphertext too short")
	}
	nonce, ciphertext := ciphertextWithNonce[:nonceSize], ciphertextWithNonce[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
