package utils

import (
	"testing"
)

func TestGenePwd(t *testing.T) {
	str1 := GenePwd()
	str2 := GenePwd()

	t.Logf("str1: %s |---| str2: %s", str1, str2)

	if len(str1) != 8 || len(str1) != len(str2) || str1 == str2 {
		t.Error("GenePwd function error!")
	}
}
