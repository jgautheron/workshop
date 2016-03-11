package vatcheck_test

import (
	"math/rand"
	"testing"

	"github.com/jgautheron/workshop/vat"
)

var numbers = []struct {
	vatNumber   string
	countryCode string
	isValid     bool
}{
	{"0" + randStringAlpha(9), "BE", true},
	{"1" + randStringAlpha(9), "BE", false},
	{"0" + randStringAlpha(10), "BE", false},
	{"0" + randStringAlpha(5), "BE", false},
	{"A" + randStringAlpha(9), "BE", false},
	{"0" + randStringAlpha(8) + "A", "BE", false},
	{randStringAlpha(9) + "A", "BE", false},

	{randStringAlpha(9), "BG", true},
	{randStringAlpha(10), "BG", true},
	{randStringAlpha(11), "BG", false},
	{randStringAlpha(5), "BG", false},
	{"A" + randStringAlpha(9), "BG", false},
	{randStringAlpha(8) + "A", "BG", false},

	{randStringAlpha(8), "CZ", true},
	{randStringAlpha(9), "CZ", true},
	{randStringAlpha(10), "CZ", true},
	{randStringAlpha(11), "CZ", false},
	{randStringAlpha(5), "CZ", false},
	{"A" + randStringAlpha(9), "CZ", false},
	{randStringAlpha(8) + "A", "CZ", false},

	{randStringAlpha(8), "DK", true}, //See if we must take space
	{randStringAlpha(11), "DK", false},
	{randStringAlpha(5), "DK", false},
	{"A" + randStringAlpha(9), "DK", false},
	{randStringAlpha(8) + "A", "DK", false},

	{randStringAlpha(9), "DE", true},
	{randStringAlpha(11), "DE", false},
	{randStringAlpha(5), "DE", false},
	{"A" + randStringAlpha(9), "DE", false},
	{randStringAlpha(8) + "A", "DE", false},

	{randStringAlpha(9), "EE", true},
	{randStringAlpha(11), "EE", false},
	{randStringAlpha(5), "EE", false},
	{"A" + randStringAlpha(9), "EE", false},
	{randStringAlpha(8) + "A", "EE", false},

	{randStringAlpha(9), "EL", true},
	{randStringAlpha(11), "EL", false},
	{randStringAlpha(5), "EL", false},
	{"A" + randStringAlpha(9), "EL", false},
	{randStringAlpha(8) + "A", "EL", false},

	{randStringAlpha(11), "HR", true},
	{randStringAlpha(15), "HR", false},
	{randStringAlpha(5), "HR", false},
	{"A" + randStringAlpha(9), "HR", false},
	{randStringAlpha(8) + "A", "HR", false},

	{randStringAlpha(11), "IT", true},
	{randStringAlpha(15), "IT", false},
	{randStringAlpha(5), "IT", false},
	{"A" + randStringAlpha(9), "IT", false},
	{randStringAlpha(8) + "A", "IT", false},

	{randStringAlpha(11), "LV", true},
	{randStringAlpha(15), "LV", false},
	{randStringAlpha(5), "LV", false},
	{"A" + randStringAlpha(9), "LV", false},
	{randStringAlpha(8) + "A", "LV", false},

	{randStringAlpha(8), "LU", true},
	{randStringAlpha(11), "LU", false},
	{randStringAlpha(5), "LU", false},
	{"A" + randStringAlpha(9), "LU", false},
	{randStringAlpha(8) + "A", "LU", false},

	{randStringAlpha(8), "HU", true},
	{randStringAlpha(11), "HU", false},
	{randStringAlpha(5), "HU", false},
	{"A" + randStringAlpha(9), "HU", false},
	{randStringAlpha(8) + "A", "HU", false},

	{randStringAlpha(8), "MT", true},
	{randStringAlpha(11), "MT", false},
	{randStringAlpha(5), "MT", false},
	{"A" + randStringAlpha(9), "MT", false},
	{randStringAlpha(8) + "A", "MT", false},

	{randStringAlpha(8), "FI", true},
	{randStringAlpha(11), "FI", false},
	{randStringAlpha(5), "FI", false},
	{"A" + randStringAlpha(9), "FI", false},
	{randStringAlpha(8) + "A", "FI", false},

	{randStringAlpha(8), "SI", true},
	{randStringAlpha(11), "SI", false},
	{randStringAlpha(5), "SI", false},
	{"A" + randStringAlpha(9), "SI", false},
	{randStringAlpha(8) + "A", "SI", false},

	{randStringAlpha(12), "SE", true},
	{randStringAlpha(15), "SE", false},
	{randStringAlpha(5), "SE", false},
	{"A" + randStringAlpha(9), "SE", false},
	{randStringAlpha(8) + "A", "SE", false},

	{randStringAlpha(10), "SK", true},
	{randStringAlpha(15), "SK", false},
	{randStringAlpha(5), "SK", false},
	{"A" + randStringAlpha(9), "SK", false},
	{randStringAlpha(8) + "A", "SK", false},

	{randStringAlpha(10), "PL", true},
	{randStringAlpha(15), "PL", false},
	{randStringAlpha(5), "PL", false},
	{"A" + randStringAlpha(9), "PL", false},
	{randStringAlpha(8) + "A", "PL", false},

	{randStringAlpha(9), "PT", true},
	{randStringAlpha(15), "PT", false},
	{randStringAlpha(5), "PT", false},
	{"A" + randStringAlpha(9), "PT", false},
	{randStringAlpha(8) + "A", "PT", false},

	{randStringAlpha(2), "RO", true},
	{randStringAlpha(3), "RO", true},
	{randStringAlpha(4), "RO", true},
	{randStringAlpha(5), "RO", true},
	{randStringAlpha(6), "RO", true},
	{randStringAlpha(7), "RO", true},
	{randStringAlpha(8), "RO", true},
	{randStringAlpha(9), "RO", true},
	{randStringAlpha(10), "RO", true},
	{randStringAlpha(15), "RO", false},
	{randStringAlpha(1), "RO", false},
	{"A" + randStringAlpha(9), "RO", false},
	{randStringAlpha(8) + "A", "RO", false},

	{randStringLettersMaj(1) + randStringAlpha(7) + randStringLettersMaj(1), "ES", true},
	{randStringAlpha(8) + randStringLettersMaj(1), "ES", true},
	{randStringLettersMaj(1) + randStringAlpha(8), "ES", true},
	{randStringLettersMaj(2) + randStringAlpha(6) + randStringLettersMaj(1), "ES", false},
	{randStringAlpha(9), "ES", true},
	{randStringAlpha(15), "ES", false},
	{randStringAlpha(5), "ES", false},

	{randStringLettersMaj(1) + randStringAlpha(8) + randStringLettersMaj(1), "FR", false}, //See if we must take space
	{randStringLettersMaj(2) + randStringAlpha(9), "FR", true},
	{randStringLettersMaj(1) + randStringAlpha(10), "FR", true},
	{randStringLettersMaj(2) + randStringAlpha(6) + randStringLettersMaj(1), "FR", false},
	{randStringAlpha(11), "FR", true},
	{randStringAlpha(15), "FR", false},
	{randStringAlpha(5), "FR", false},

	{randStringLettersMaj(1) + randStringAlpha(8) + randStringLettersMaj(1), "CY", false},
	{randStringLettersMaj(2) + randStringAlpha(9), "CY", false},
	{randStringAlpha(8) + randStringLettersMaj(1), "CY", true},
	{randStringLettersMaj(2) + randStringAlpha(6) + randStringLettersMaj(1), "CY", false},
	{randStringAlpha(9), "CY", false},

	{randStringAlpha(9), "LT", true},
	{randStringAlpha(12), "LT", true},
	{randStringAlpha(10), "LT", false},
	{randStringAlpha(15), "LT", false},
	{randStringAlpha(5), "LT", false},
	{"A" + randStringAlpha(9), "LT", false},
	{randStringAlpha(8) + "A", "LT", false},

	{randStringAlpha(9) + "B" + randStringAlpha(2), "NL", true},
	{randStringAlpha(9) + "C" + randStringAlpha(2), "NL", false},
	{randStringAlpha(12), "NL", false},
	{randStringAlpha(15), "NL", false},
	{randStringAlpha(5), "NL", false},
	{"A" + randStringAlpha(9), "NL", false},
	{randStringAlpha(8) + "A", "NL", false},

	{randStringAlpha(1) + randStringLettersMaj(1) + randStringAlpha(5) + randStringLettersMaj(1), "IE", true},
	{randStringAlpha(2) + randStringLettersMaj(1) + randStringAlpha(5) + randStringLettersMaj(1), "IE", false},
	{randStringAlpha(1) + randStringLettersMaj(2), "IE", false},
	{randStringAlpha(7) + "C" + randStringAlpha(2), "IE", false},
	{randStringAlpha(7) + randStringLettersMaj(1), "IE", true},
	{randStringAlpha(7) + "WI", "IE", true},
	{randStringAlpha(12), "IE", false},
	{randStringAlpha(15), "IE", false},
	{randStringAlpha(5), "IE", false},
	{"A" + randStringAlpha(9), "IE", false},
	{randStringAlpha(8) + "A", "IE", false},
}

func TestValidFormat(t *testing.T) {
	var isValid bool
	for _, tt := range numbers {
		isValid = vatcheck.IsValidFormat(tt.countryCode + tt.vatNumber)
		if tt.isValid != isValid {
			t.Errorf(`Test failed for "%s%s"`, tt.countryCode, tt.vatNumber)
		}
	}
}

func TestInvalidCountry(t *testing.T) {
	isValid := vatcheck.IsValidFormat("US123")
	if isValid {
		t.Error("A non-EU country should not validate")
	}
}

const (
	letterBytes      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterMajBytes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterAlphaBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	alphaBytes       = "0123456789"

	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringLetters(n int) string {
	return randStringBytesMaskImpr(n, letterBytes)
}

func randStringLettersMaj(n int) string {
	return randStringBytesMaskImpr(n, letterMajBytes)
}

func randStringLettersAlpha(n int) string {
	return randStringBytesMaskImpr(n, letterAlphaBytes)
}

func randStringAlpha(n int) string {
	return randStringBytesMaskImpr(n, alphaBytes)
}

func randStringBytesMaskImpr(n int, stringTorandom string) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(stringTorandom) {
			b[i] = stringTorandom[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
