package newebpay

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

var (
	blockSize = 32
)

func KeyEncrypt(plaintext, hashKey, hashIV string) (string, string) {
	bKey := []byte(hashKey)
	bIV := []byte(hashIV)
	bPlaintext := AddPKCS7Padding([]byte(plaintext))
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	hexCipherText := hex.EncodeToString(ciphertext)

	h := sha256.New()
	h.Write([]byte("HashKey=" + hashKey + "&" + hexCipherText + "&HashIV=" + hashIV))
	code256 := fmt.Sprintf("%x", h.Sum(nil))
	code256 = strings.ToUpper(code256)

	return hexCipherText, code256
}

func AddPKCS7Padding(ciphertext []byte) []byte {
	padding := blockSize - (len(ciphertext) % blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func KeyDecrypt(deCodeText, hashKey, hashIV string) string {
	cipherTextDecoded, err := hex.DecodeString(deCodeText)
	if err != nil {
		panic(err)
	}
	block, err := aes.NewCipher([]byte(hashKey))
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCDecrypter(block, []byte(hashIV))
	mode.CryptBlocks(cipherTextDecoded, cipherTextDecoded)
	return string(cipherTextDecoded)
}

func GenerateCheckValue(plaintext, hashKey, hashIV string) string {
	//fmt.Println("IV="+HashIV+"&" + plaintext + "&Key="+HashKey)
	h := sha256.New()
	h.Write([]byte("IV=" + hashIV + "&" + plaintext + "&Key=" + hashKey))
	code256 := fmt.Sprintf("%x", h.Sum(nil))
	code256 = strings.ToUpper(code256)

	return code256
}
