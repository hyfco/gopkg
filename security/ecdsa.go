package security

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"math/big"
	"os"
	"strings"
)

// ReadPrivateKeyFromFile 从文件中读取私钥
func ReadPrivateKeyFromFile(filename string) (*ecdsa.PrivateKey, error) {
	filename = os.ExpandEnv(filename)
	keyData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	privKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}

// ReadPublicKeyFromFile 从文件读取公钥
func ReadPublicKeyFromFile(filename string) (*ecdsa.PublicKey, error) {
	filename = os.ExpandEnv(filename)
	keyData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pubKey.(*ecdsa.PublicKey), nil
}

func BuildPrivateKey(in string) (privateKey *ecdsa.PrivateKey, err error) {
	bytes, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return nil, err
	}
	privateKey, err = x509.ParseECPrivateKey(bytes)
	if err != nil {
		return nil, err
	}
	return
}

func BuildPublicKey(in string) (publicKey *ecdsa.PublicKey, err error) {
	bytes, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(bytes), "+")
	xStr := split[0]
	yStr := split[1]
	x := new(big.Int)
	y := new(big.Int)
	err = x.UnmarshalText([]byte(xStr))
	if err != nil {
		return nil, err
	}
	err = y.UnmarshalText([]byte(yStr))
	if err != nil {
		return nil, err
	}
	publicKey = &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}
	return
}
