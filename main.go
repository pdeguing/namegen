/*
Package namegen implements a simple library to generate random names.
*/
package namegen

import (
	"math/rand"
	"strings"
)

func isVowel(c byte) bool {
	return strings.Contains("aeiouy", string(c))
}

func isConsonant(c byte) bool {
	return strings.Contains("bcdfghjklmnpqrstvxz", string(c))
}

func getVowel(r *rand.Rand) byte {
	vowel := "aeiouy"
	index := r.Intn(len(vowel))
	return vowel[index]
}

func getConsonant(r *rand.Rand) byte {
	consonant := "bcdfghjklmnpqrstvxz"
	index := r.Intn(len(consonant))
	return consonant[index]
}

func Name(r *rand.Rand, length int) (string, error) {
	if length <= 0 {
		return "", nil
	}

	buffer := make([]byte, length)

	if r.Intn(9) < 4 {
		buffer[0] = getVowel(r)
	} else {
		buffer[0] = getConsonant(r)
	}

	for i := 1; i < length; i++ {
		if i > 1 {
			if isVowel(buffer[i-1]) && isVowel(buffer[i-2]) {
				buffer[i] = getConsonant(r)
				continue
			}
			if isConsonant(buffer[i-1]) && isConsonant(buffer[i-2]) {
				buffer[i] = getVowel(r)
				continue
			}
		}
		if r.Intn(9) < 4 {
			buffer[i] = getVowel(r)
		} else {
			buffer[i] = getConsonant(r)
		}
	}
	generated := string(buffer)
	return generated, nil
}

func Nname(r *rand.Rand, min_length, max_length, n int) ([]string, error) {
	var length int
	var names	[]string

	for i := 0; i < n; i++ {
		length = r.Intn(max_length-min_length) + min_length
		generated, err := Name(r, length)
		if err != nil {
			return names, err
		}
		names = append(names, generated)
	}
	return names, nil
}
