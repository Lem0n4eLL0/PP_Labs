package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println("First task:\n-----------------")
	var test string = "string"
	fmt.Println("md5: ", md5Hashing(test))
	fmt.Println("sha256: ", sha256Hashing(test))
	fmt.Println("sha512 ", sha512Hashing(test))
	fmt.Println(checkHash(test, md5Hashing(test), "md5"))
	fmt.Println(checkHash(test, md5Hashing(test), "sha256"))
}

func md5Hashing(str string) string {
	md5Hash := md5.Sum([]byte(str))
	return hex.EncodeToString(md5Hash[:])
}

func sha256Hashing(str string) string {
	sha256Hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(sha256Hash[:])
}

func sha512Hashing(str string) string {
	sha512Hash := sha512.Sum512([]byte(str))
	return hex.EncodeToString(sha512Hash[:])
}

func checkHash(str string, hash string, algorithm string) bool {
	switch algorithm {
	case "md5":
		h := md5.Sum([]byte(str))
		return hex.EncodeToString(h[:]) == hash
	case "sha256":
		h := sha256.Sum256([]byte(str))
		return hex.EncodeToString(h[:]) == hash
	case "sha512":
		h := sha256.Sum256([]byte(str))
		return hex.EncodeToString(h[:]) == hash
	default:
		return false
	}
}
