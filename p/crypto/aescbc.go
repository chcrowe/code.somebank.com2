package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func CreateAesCipherFromBase64Key(b64key string) (cipher.Block, error) {
	b64keyBytes, _ := base64.StdEncoding.DecodeString(b64key)
	return aes.NewCipher(b64keyBytes)
}

func Decrypt(block cipher.Block, iv, ciphered []byte) []byte {
	decrypter := cipher.NewCBCDecrypter(block, iv)
	deciphered := make([]byte, len(ciphered))
	decrypter.CryptBlocks(deciphered, ciphered)

	unpadding := int(deciphered[len(deciphered)-1])
	for i := len(deciphered) - 1; i >= len(deciphered)-unpadding; i-- {
		deciphered[i] = 0
	}

	return deciphered
}

func DecryptFromBase64String(block cipher.Block, iv []byte, b64ciphered string) []byte {
	ciphered, _ := base64.StdEncoding.DecodeString(b64ciphered)
	return Decrypt(block, iv, ciphered)
}

func Encrypt(block cipher.Block, iv []byte, b64plaintext string) []byte {
	plain, _ := base64.StdEncoding.DecodeString(b64plaintext)
	encrypter := cipher.NewCBCEncrypter(block, iv)
	enciphered := make([]byte, len(plain))
	encrypter.CryptBlocks(enciphered, plain)
	return RemoveCbcPadding(enciphered)
}

func EncryptToBase64String(block cipher.Block, iv []byte, b64plaintext string) string {
	return base64.StdEncoding.EncodeToString(Encrypt(block, iv, b64plaintext))
}

func RemoveCbcPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
