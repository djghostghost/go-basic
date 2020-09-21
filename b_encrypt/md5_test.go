package b_encrypt

import "testing"

func TestMd5Hex(t *testing.T) {
	v := Md5Hex("testmd5hex")
	if v != "cf4c16a6e02526b9f0c19260101ad5ab" {
		t.Fail()
	}
}
