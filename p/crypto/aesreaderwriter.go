package crypto

import (
	"code.google.com/p/go.crypto/pbkdf2"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"io"
	"math/rand"
)

type AesReaderWriter struct {
	Block cipher.Block
	IV    []byte
}

func NewAesReaderWriter(derivationelements []string, salt []byte) *AesReaderWriter {

	offset := 0
	hashedvalues := make([]byte, 64*len(derivationelements))
	for _, de := range derivationelements {
		copy(hashedvalues[offset:], hashValue(de))
		offset += 64
	}

	pwd := hashValue(string(hashedvalues))
	key := pbkdf2.Key(pwd, salt, 2048, 32, sha512.New)

	arw := AesReaderWriter{}

	arw.Block, _ = aes.NewCipher(key)
	arw.IV = salt[:aes.BlockSize]

	return &arw
}

func (a *AesReaderWriter) Encrypt(plaintext []byte) []byte {
	// encrypt
	encrypter := cipher.NewCFBEncrypter(a.Block, a.IV)
	encrypted := make([]byte, len(plaintext))
	encrypter.XORKeyStream(encrypted, plaintext)
	return encrypted
}

func (a *AesReaderWriter) Decrypt(ciphertext []byte) []byte {
	// encrypt
	decrypter := cipher.NewCFBDecrypter(a.Block, a.IV)
	decrypted := make([]byte, len(ciphertext))
	decrypter.XORKeyStream(decrypted, ciphertext)
	return decrypted
}

func hashValue(s string) []byte {
	h512 := sha512.New()
	io.WriteString(h512, s)
	return h512.Sum(nil)
}

func randomBytes(l int) []byte {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return bytes
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
