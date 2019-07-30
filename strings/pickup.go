package strings

import (
	"math/rand"
)

const (
	FDigits = 1 << iota
	FLowerCase
	FUpperCase

	FLetters = FLowerCase | FUpperCase
	FAny     = FDigits | FLetters
)

// pickup string with length and flags.
func Pickup(length int, flags int) string {
	f := []func() string{}
	if flags&FDigits == FDigits {
		f = append(f, pickupDigits)
	}
	if flags&FLowerCase == FLowerCase {
		f = append(f, pickupLowerCase)
	}
	if flags&FUpperCase == FUpperCase {
		f = append(f, pickupUpperCase)
	}
	l := len(f)
	var res string
	for i := 0; i < length; i++ {
		res += f[rand.Intn(l)]()
	}
	return res
}

// pickup random integer.
func pickupDigits() string {
	return pickup(10, 48)
}

// pickup random ascii lowercase.
func pickupLowerCase() string {
	return pickup(26, 97)
}

// pickup random ascii uppercase.
func pickupUpperCase() string {
	return pickup(26, 65)
}

// pickup random string with range and base.
func pickup(r int, b int) string {
	return string(rand.Intn(r) + b)
}
