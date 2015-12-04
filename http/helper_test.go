package http

import (
	"testing"
)

func TestGenerateClientID(t *testing.T) {
	str1 := generateClientID("str1")
	str2 := generateClientID("str1")

	if str1 != str2 {
		t.Error("Did not generate same value for key")
	}
}
