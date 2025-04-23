package pw

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	EncryptionKey []byte
}

var service *Service

func NewService(key string) (*Service, error) {
	if service == nil {
		k, err := hex.DecodeString(key)
		if err != nil {
			return nil, err
		}
		if len(k) != 16 && len(k) != 24 && len(k) != 32 {
			return nil, errors.New("key must be 16, 24, or 32 bytes")
		}
		service = &Service{EncryptionKey: k}
	}

	return service, nil
}

func Instance() *Service {
	return service
}

func (s *Service) HashAndEncrypt(password string) (*PasswordPair, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(s.EncryptionKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, hashed, nil)

	return &PasswordPair{
		EncryptedHash: base64.StdEncoding.EncodeToString(ciphertext),
		Nonce:         base64.StdEncoding.EncodeToString(nonce),
	}, nil
}

func (s *Service) Verify(password string, pair *PasswordPair) error {
	block, err := aes.NewCipher(s.EncryptionKey)
	if err != nil {
		return err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce, err := base64.StdEncoding.DecodeString(pair.Nonce)
	if err != nil {
		return err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(pair.EncryptedHash)
	if err != nil {
		return err
	}

	plaintextHash, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return errors.New("failed to decrypt or tampered data")
	}

	return bcrypt.CompareHashAndPassword(plaintextHash, []byte(password))
}
