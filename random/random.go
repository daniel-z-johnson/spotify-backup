package random

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

const sessionTokenBytes = 24

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	bRead, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("bytes: %w", err)
	}
	if bRead < n {
		return nil, fmt.Errorf("didn't read enough bytes wanted: '%d' got: '%d'", n, bRead)
	}
	return b, nil
}

func String(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", fmt.Errorf("issue creating rand string: '%s'", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil

}

func SessionToken() (string, error) {
	return String(sessionTokenBytes)
}
