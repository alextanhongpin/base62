package base62

import (
	"math"
)

// Base62 character set, [a-zA-Z0-9].
var base62Chars = [...]rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}

var base62Map map[rune]int

// Base62 represents the base52 const.
const Base62 int64 = 62

func init() {
	base62Map = make(map[rune]int)
	for k, v := range base62Chars {
		base62Map[v] = k
	}
}

// Encode converts the given integer into a base 62 string.
func Encode(in int64) string {
	var out []rune
	for in > 0 {
		char := base62Chars[in%62]
		out = append([]rune{char}, out...)
		in /= 62
	}
	return string(out)
}

// Decode attempts to convert a base 62 string back into an integer.
func Decode(in string) int64 {
	var sum int64
	for i, v := range in {
		pow := int64(len(in) - 1 - i)
		p := math.Pow(float64(Base62), float64(pow))
		sum += int64(p) * int64(base62Map[v])
	}
	return sum
}
