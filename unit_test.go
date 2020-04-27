package namegen

import (
	"testing"
	"math/rand"
)

func TestIsVowel(t *testing.T) {
	got := isVowel(0)
	if got != false {
		t.Errorf("isVowel(0) = %t; want false", got)
	}
}

func TestNname(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	min_length := 3
	max_length := 9
	got, err := Nname(r, min_length, max_length, 50)
	if err != nil {
		t.Errorf("error while generating 50 names: %e", err)
	}
}
