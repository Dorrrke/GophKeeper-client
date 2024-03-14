package coder

import (
	"errors"

	errText "github.com/Dorrrke/GophKeeper-client/internal/domain/errors"
	"github.com/fernet/fernet-go"
)

const DecodeKey = "cx_0x654RpI-jtBZ7oE8h_eQsKImvJlKFeSbXpwM7e4="

var ErrDecodeData = errors.New(errText.DataDecryptError)

func Decoder(encData []byte) ([]byte, error) {
	key := fernet.MustDecodeKeys(DecodeKey)
	data := fernet.VerifyAndDecrypt(encData, -1, key)
	if len(data) == 0 {
		return nil, ErrDecodeData
	}
	return data, nil
}

func Encoder(data []byte) ([]byte, error) {
	key := fernet.MustDecodeKeys(DecodeKey)
	encData, err := fernet.EncryptAndSign(data, key[0])
	if err != nil {
		return nil, err
	}
	return encData, nil
}
