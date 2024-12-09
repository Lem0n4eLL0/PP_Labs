package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	str := "Hello world!"
	fmt.Printf("Base string:\n%s\n", str)
	encryptStr := EncryptIt(str)
	fmt.Printf("Encrypt:\n%x\n", *encryptStr)
	decriptStr := DecryptIt(encryptStr)
	fmt.Printf("Decrypted: \n%s\n", decriptStr)
}

func EncryptIt(info string) *[]byte {
	publicKeyPEM, err := os.ReadFile("public.pem")
	if err != nil {
		panic(err)
	}
	publicKeyBlock, _ := pem.Decode(publicKeyPEM)
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}

	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), []byte(info))
	if err != nil {
		panic(err)
	}

	return &ciphertext
}

func DecryptIt(info *[]byte) string {
	privateKeyPEM, err := os.ReadFile("private.pem")
	if err != nil {
		panic(err)
	}
	privateKeyBlock, _ := pem.Decode(privateKeyPEM)
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}

	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, *info)
	if err != nil {
		panic(err)
	}
	return string(plaintext)
}
