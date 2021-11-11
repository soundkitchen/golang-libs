package strings

import (
	"fmt"
	"math/rand"
)

const (
	FDigits = 1 << iota
	FLowerCase
	FUpperCase
	FPunctuation

	FLetters = FLowerCase | FUpperCase
	FAny     = FDigits | FLetters | FPunctuation
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
	if flags&FPunctuation == FPunctuation {
		f = append(f, pickupPunctuation)
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

// pickup random ascii punctuation.
func pickupPunctuation() string {
	var res string
	switch rand.Intn(4) {
	case 0:
		// !"#$%&'()*+,-./
		res = pickup(15, 33)
	case 1:
		// :;<=>?@
		res = pickup(7, 58)
	case 2:
		// [\]^_`
		res = pickup(6, 91)
	case 3:
		// {|}~
		res = pickup(4, 123)
	}
	return res
}

// pickup random string with range and base.
func pickup(r int, b int) string {
	return fmt.Sprintf("%c", rand.Intn(r)+b)
}
