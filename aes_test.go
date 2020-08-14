package aes

import (
	"log"
	"testing"
)

func TestAES(t *testing.T) {
	key := "321423u9y8d2fwfl"
	pass := "vdncloud1234526"
	xpass, err := AesEncryptString(pass, key)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("加密后:%s\n",xpass)

	tpass, err := AesDecryptString(xpass, key)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("解密后:%s\n", tpass)
}
