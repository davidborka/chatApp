package auth

import (
	"bytes"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
)

var (
	err                   error
	privKey               *rsa.PrivateKey
	pubKey                *rsa.PublicKey
	pubKeyBytes           []byte
	SigningKey, VerifyKey []byte
)

//InitKeys init public and private RSA key.
func InitKeys() ([]byte, []byte) {

	privKey, err = rsa.GenerateKey(cryptorand.Reader, 2048)
	if err != nil {
		log.Fatal("Error generating private key")
	}
	pubKey = &privKey.PublicKey //hmm, this is stdlib manner...

	// Create signingKey from privKey
	// prepare PEM block
	var privPEMBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey), // serialize private key bytes
	}
	// serialize pem
	privKeyPEMBuffer := new(bytes.Buffer)
	pem.Encode(privKeyPEMBuffer, privPEMBlock)
	//done
	signingKeyb := privKeyPEMBuffer.Bytes()
	//SigningKey, err := jwt.ParseRSAPrivateKeyFromPEM(signingKeyb)
	if err != nil {
		log.Fatal(err)
	}

	// create verificationKey from pubKey. Also in PEM-format
	pubKeyBytes, err = x509.MarshalPKIXPublicKey(pubKey) //serialize key bytes
	if err != nil {
		// heh, fatality
		log.Fatal("Error marshalling public key")
	}

	var pubPEMBlock = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	}
	// serialize pem
	pubKeyPEMBuffer := new(bytes.Buffer)
	pem.Encode(pubKeyPEMBuffer, pubPEMBlock)
	verifyKeyB := pubKeyPEMBuffer.Bytes()

	//VerifyKey, _ := jwt.ParseECPublicKeyFromPEM(verifyKeyB)
	// done
	return signingKeyb, verifyKeyB
}
