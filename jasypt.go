package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	iterations := 1000
	password := "9b8739b5-42a4-4173-99af-e409201ef9e3"
	keylen := 16 //first 8 are key, second 8 are initialization vector
	salt := make([]byte, des.BlockSize)
	rand.Read(salt)

	pbekey := pbkdf2.Key([]byte(password), salt, iterations, keylen, md5.New)
	key := make([]byte, 8)
	iv := make([]byte, 8)
	copy(pbekey, key[0:7])
	copy(pbekey, key[7:])
	fmt.Printf("PBE key len: %d\n", len(key))

	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	encryptor := cipher.NewCBCEncrypter(block, iv)

	ciphertext := make([]byte, len("Hello"))

	encryptor.CryptBlocks(ciphertext, []byte("Hello"))

	encoded := base64.StdEncoding.EncodeToString(key)
	fmt.Printf("Hello World: %s\n", encoded)
}
