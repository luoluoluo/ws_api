package library

import "testing"

func TestNewCBCEncrypter(t *testing.T) {
	aes := &AES{Key: "1ba7e3787f99fb13"}
	text, err := aes.NewCBCEncrypter("test123")
	if err != nil {
		t.Error(err)
	}
	text, err = aes.NewCBCDecrypter(text)
	if err != nil {
		t.Error(err)
	}
	if text != "test123" {
		t.Error(text)
	}
}
