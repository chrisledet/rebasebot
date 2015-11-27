package http

import (
	"testing"
)

func TestGenerateClientId(t *testing.T) {
	str1 := generateClientId("str1")
	str2 := generateClientId("str1")

	if str1 != str2 {
		t.Error("Did not generate same value for key")
	}
}
