package namegen

import "testing"

func TestIsVowel(t *testing.T) {
	got := isVowel(0)
	if got != false {
		t.Errorf("isVowel(0) = %t; want false", got)
	}
}
