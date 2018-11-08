package base62

import (
	"fmt"
)

// Base62 represents the base62 const.
const Base62 uint64 = 62

var (
	// Base62 character set, [A-Za-z0-9]. https://tools.ietf.org/html/rfc4648
	// We follow the orders similar to base64.
	base62Chars = [...]rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}

	base62Map map[rune]int
)

func init() {
	base62Map = make(map[rune]int)
	for k, v := range base62Chars {
		base62Map[v] = k
	}
}

func main() {
	fmt.Println("encode:", Encode(2024))
	fmt.Println("decode:", Decode("go"))
}

// EncodePerf converts the given integer into a base 62 string.
func EncodePerf(in uint64) string {
	i, tmp := 0, in
	for {
		tmp /= Base62
		i++
		if tmp == 0 {
			break
		}
	}
	out := make([]rune, i)
	for {
		i--
		if i < 0 {
			break
		}
		out[i] = base62Chars[in%62]
		in /= Base62
	}
	return string(out)
}

// Encode converts the given integer into a base 62 string.
func Encode(in uint64) string {
	var out []rune
	// If set to >0, Encode(0) will return empty string.
	for {
		char := base62Chars[in%62]
		out = append([]rune{char}, out...)
		in /= Base62
		if in == 0 {
			break
		}
	}
	return string(out)
}

// Decode attempts to convert a base 62 string back into an integer.
func Decode(in string) uint64 {
	var sum uint64
	for _, v := range in {
		sum = sum*Base62 + uint64(base62Map[v])
	}
	return sum
}
