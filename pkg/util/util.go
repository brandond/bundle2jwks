package util

import (
	"bytes"
	"crypto"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base32"
	"strings"

	"github.com/go-jose/go-jose/v4"
	"k8s.io/client-go/util/cert"
)

func GetKeySet(path string) (*jose.JSONWebKeySet, error) {
	rootCerts, err := cert.CertsFromFile(path)
	if err != nil {
		return nil, err
	}

	keySet := &jose.JSONWebKeySet{}
	for _, rootCert := range rootCerts {
		keySet.Keys = append(keySet.Keys, jose.JSONWebKey{
			Key:   rootCert.PublicKey,
			KeyID: LibtrustKeyID(rootCert.PublicKey),
		})
	}
	return keySet, nil
}

func LibtrustKeyID(publickey crypto.PublicKey) string {
	keyBytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		return ""
	}

	sum := sha256.Sum256(keyBytes)
	b64 := strings.TrimRight(base32.StdEncoding.EncodeToString(sum[:30]), "=")

	var buf bytes.Buffer
	var i int
	for i = 0; i < len(b64)/4-1; i++ {
		start := i * 4
		end := start + 4
		buf.WriteString(b64[start:end] + ":")
	}
	buf.WriteString(b64[i*4:])

	return buf.String()
}
