package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

var key = []byte("e78d3c82-126e7-447d-14ee5ea1a21d")

func Encrypt(message string) (encoded string, err error) {
	plainText := []byte(message)

	// Create a new AES cipher using the key
	block, err := aes.NewCipher(key)
	CheckErr(err)

	// Make the cipher text a byte array of size BlockSize + the length of the message
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	// iv is the ciphertext up to the blocksize (16)
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return base64.RawStdEncoding.EncodeToString(cipherText), err
}

func Decrypt(secure string) (decoded string, err error) {
	// Remove base64 encoding
	cipherText, err := base64.RawStdEncoding.DecodeString(secure)
	CheckErr(err)

	block, err := aes.NewCipher(key)
	CheckErr(err)

	// If the length of the cipherText is less than 16 Bytes
	if len(cipherText) < aes.BlockSize {
		err = errors.New("ciphertext block size is too short")
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// Decrypt the message
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), err
}
