package ubix

import (
	"log"
	"testing"
)

const aes_enc_str = "1f594e603731f8b8e6ec710463378a70"
const aes_dec_str = "1800"

var aes_key = "test202020202020"

func TestCBCEncrypter(t *testing.T) {
	enc := AesCBCEncrypte([]byte(aes_dec_str), aes_key)

	t.Log(enc)
	dec, err := AesCBCDecrypte(aes_enc_str, aes_key)
	if err != nil {
		t.Error(err)
	}
	if dec != aes_dec_str {
		t.Fatal(dec)
	}
}

func TestCBCDecrypter(t *testing.T) {
	dec, err := AesCBCDecrypte(aes_enc_str, aes_key)
	if err != nil {
		t.Error(err)
	}
	log.Println(string(dec))
	if aes_dec_str == string(dec) {
		t.Log(string(dec))
	} else {
		t.Error("aes_decode_error")
	}
}
