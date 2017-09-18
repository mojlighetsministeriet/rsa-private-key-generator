package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func panicValue(callback func()) (recovered interface{}) {
	defer func() {
		recovered = recover()
	}()
	callback()
	return
}

func pemStringToPrivateKey(pemString string) (privateKey *rsa.PrivateKey, err error) {
	block, _ := pem.Decode([]byte(pemString))
	if block == nil {
		err = errors.New("Unable to find private key in pemString")
		return
	} else if block.Type != "RSA PRIVATE KEY" {
		err = errors.New("Unable to find private key in pemString, type found " + block.Type)
		return
	}

	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	return
}

func TestMain(test *testing.T) {
	os.Args = []string{"rsa-private-key-generator", "128"}

	rescueStdout := os.Stdout
	readStream, writeStream, _ := os.Pipe()
	os.Stdout = writeStream

	got := panicValue(func() { main() })
	err, ok := got.(error)
	assert.Equal(test, false, ok)
	assert.NoError(test, err)

	writeStream.Close()
	output, _ := ioutil.ReadAll(readStream)
	os.Stdout = rescueStdout

	assert.Equal(test, true, len(output) > 0)

	_, err = pemStringToPrivateKey(string(output))
	assert.NoError(test, err)
}

func TestMainInvalidKeyGeneration(test *testing.T) {
	os.Args = []string{"rsa-private-key-generator", "invalid1bits"}

	rescueStdout := os.Stdout
	_, writeStream, _ := os.Pipe()
	os.Stdout = writeStream

	got := panicValue(func() { main() })
	err, ok := got.(error)
	assert.Equal(test, true, ok)
	assert.Error(test, err)

	writeStream.Close()
	os.Stdout = rescueStdout
}

func TestMainInvalidEmptyBits(test *testing.T) {
	os.Args = []string{"rsa-private-key-generator", ""}

	rescueStdout := os.Stdout
	_, writeStream, _ := os.Pipe()
	os.Stdout = writeStream

	got := panicValue(func() { main() })
	err, ok := got.(error)
	assert.Equal(test, true, ok)
	assert.Error(test, err)

	writeStream.Close()
	os.Stdout = rescueStdout
}
