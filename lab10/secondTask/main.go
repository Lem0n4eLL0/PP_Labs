package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

func main() {
	key := "string"
	info := "important info"
	fmt.Println("info: ", info)
	encriptInfo := encryptIt([]byte(info), key)
	fmt.Println("encriptInfo: ", string(encriptInfo))
	decriptInfo := decryptIt(encriptInfo, key)
	fmt.Println("decriptInfo: ", string(decriptInfo))
}

func encryptIt(value []byte, keyPhrase string) []byte {
	aesBlock, err := aes.NewCipher([]byte(md5Hashing(keyPhrase)))
	if err != nil {
		fmt.Println(err)
	}
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		fmt.Println(err)
	}
	nonce := make([]byte, gcmInstance.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)
	cipheredText := gcmInstance.Seal(nonce, nonce, value, nil)
	return cipheredText
}

func decryptIt(ciphered []byte, keyPhrase string) []byte {
	hashedPhrase := md5Hashing(keyPhrase)
	aesBlock, err := aes.NewCipher([]byte(hashedPhrase))
	if err != nil {
		log.Fatalln(err)
	}
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		log.Fatalln(err)
	}
	nonceSize := gcmInstance.NonceSize()
	nonce, cipheredText := ciphered[:nonceSize], ciphered[nonceSize:]
	originalText, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return originalText
}

func md5Hashing(str string) string {
	md5Hash := md5.Sum([]byte(str))
	return hex.EncodeToString(md5Hash[:])
}
