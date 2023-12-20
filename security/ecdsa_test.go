package security

import (
	"testing"
)

func TestReadPrivateKeyFromFile(t *testing.T) {
	_, err := ReadPrivateKeyFromFile("$GOPATH/src/github.com/hyfco/oasis/backend/deploy/configs/ecc-private-key.pem")
	if err != nil {
		t.Errorf("err:%v", err)
	}
}

func TestReadPublicKeyFromFile(t *testing.T) {
	_, err := ReadPublicKeyFromFile("$GOPATH/src/github.com/hyfco/oasis/backend/deploy/configs/ecc-public-key.pem")
	if err != nil {
		t.Errorf("%v", err)
	}
}
