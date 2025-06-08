package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type Config interface {
	GetCipherKey() string
}
type Encrypter struct {
	Key string
}

func NewEncrypter(cfg Config) *Encrypter {
	return &Encrypter{Key: cfg.GetCipherKey()}
}

func (enc *Encrypter) Encrypt(plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}
	return aesGCM.Seal(nonce, nonce, plainText, nil), nil
}

func (enc *Encrypter) Decrypt(encryptedText []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := aesGCM.NonceSize()
	nonce, encryptedText := encryptedText[:nonceSize], encryptedText[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, encryptedText, nil)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}
