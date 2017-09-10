package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	bitsArgument := "4096"
	if len(os.Args) > 1 {
		bitsArgument = os.Args[1]
	}

	pattern := regexp.MustCompile("[^\\d]+")
	bits, err := strconv.Atoi(pattern.ReplaceAllString(bitsArgument, ""))
	if err != nil {
		panic("First argument must be the number of bits that should be used for the key e.g. 4096")
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}

	block := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}

	key := string(pem.EncodeToMemory(block))
	fmt.Print(key)
}
