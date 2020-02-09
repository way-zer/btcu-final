package clientSDK

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

type PublicKey string
type PrivateKey string

var (
	Base64Encoding = base64.StdEncoding
)

func GenerateKeys() (*PrivateKey, *PublicKey, error) {
	private, err := rsa.GenerateKey(rand.Reader, keyBit)
	if err != nil {
		return nil, nil, err
	}
	err = private.Validate()
	if err != nil {
		return nil, nil, err
	}
	privateBs := x509.MarshalPKCS1PrivateKey(private)
	privateKey := PrivateKey(Base64Encoding.EncodeToString(privateBs))
	publicBs := x509.MarshalPKCS1PublicKey(&private.PublicKey)
	publicKey := PublicKey(Base64Encoding.EncodeToString(publicBs))
	return &privateKey, &publicKey, nil
}

func DecodePrivateKey(key PrivateKey) (*rsa.PrivateKey, error) {
	bs, err := Base64Encoding.DecodeString(string(key))
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PrivateKey(bs)
}

func DecodePublicKey(key PublicKey) (*rsa.PublicKey, error) {
	bs, err := Base64Encoding.DecodeString(string(key))
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PublicKey(bs)
}
