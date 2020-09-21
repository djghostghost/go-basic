package b_encrypt

import (
	"encoding/hex"
	"testing"
)

func TestAESECBPKCS5(t *testing.T) {
	plaintext := []byte("hello world")
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")

	ciphertext, err := AESECBPKCS5Encrypt(plaintext, key)
	if err != nil {
		t.Fatal("encrypt error", err)
	}
	t.Logf("ciphertext: %v", string(ciphertext))

	plaintext1, err := AESECBPKCS5Decrypt(ciphertext, key)
	if err != nil {
		t.Fatal("decrypt error", err)
	}

	t.Logf("plaintext1: %v", string(plaintext1))

	if string(plaintext) != string(plaintext1) {
		t.Errorf("wrong result: %v %v", string(plaintext), string(plaintext1))
	}
}

func TestAESCBCPKCS5(t *testing.T) {
	plaintext := []byte("hello world")
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	iv := []byte("0123456789abcdef")

	ciphertext, err := AESCBCPKCS5Encrypt(plaintext, key, iv)
	if err != nil {
		t.Fatal("encrypt error", err)
	}
	t.Logf("ciphertext: %v", string(ciphertext))

	plaintext1, err := AESCBCPKCS5Decrypt(ciphertext, key, iv)

	if err != nil {
		t.Fatal("decrypt error", err)
	}

	t.Logf("plaintext1: %v", string(plaintext1))

	if string(plaintext) != string(plaintext1) {
		t.Errorf("wrong result: %v %v", string(plaintext), string(plaintext1))
	}

}
