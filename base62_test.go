package base62_test

import (
	"log"
	"math/rand"
	"testing"
	"testing/quick"

	"github.com/alextanhongpin/base62"
)

type Phonetic struct {
	Text string
	ID   uint64
}

func TestBase62Defaults(t *testing.T) {
	for _, c := range base62.DEFAULT_CHARS {
		s := string(c)
		if got := base62.Encode(base62.Decode(s)); got != s {
			t.Fatalf("got %s, want %c", got, c)
		}
	}
}
func TestBase62Conversion(t *testing.T) {
	tests := []Phonetic{
		{"Alfa", 0x5e56b},
		{"Bravo", 0x26493f5},
		{"Charlie", 0x2f05af0401},
		{"Delta", 0x3f8e21f},
		{"Echo", 0x13eab1},
		{"Foxtrot", 0x5844e26adc},
		{"Golf", 0x1be5c0},
		{"Hotel", 0x7a39498},
		{"India", 0x8807a5d},
		{"Juliet", 0x24c0017bc},
		{"Kilo", 0x2a1791},
		{"Lima", 0x2dbab9},
		{"Mike", 0x315d39},
		{"November", 0x2effb90a3716},
		{"Oscar", 0xddb68b2},
		{"Papa", 0x3bcf33},
		{"Quebec", 0x3ca23c4e7},
		{"Romeo", 0x1071dc1f},
		{"Sierra", 0x42cc93f4b},
		{"Tango", 0x1201f22f},
		{"Uniform", 0x11e603c6be3},
		{"Victor", 0x4d093525a},
		{"Whiskey", 0x1378be69c91},
		{"Xray", 0x59e2ad},
		{"Yankee", 0x56d804485},
		{"Zulu", 0x61584f},
		{"One", 0xeb0b},
		{"Two", 0x13857},
		{"Three", 0x121ba281},
		{"Four", 0x184502},
		{"Five", 0x17eb1b},
		{"Six", 0x125f8},
		{"Seven", 0x112f7da2},
		{"Eight", 0x4e89826},
		{"Nine", 0x3500eb},
		{"Zero", 0x60697d}}

	for _, tt := range tests {
		if got := base62.Decode(tt.Text); got != tt.ID {
			t.Fatalf("want %v, got %v", tt.ID, got)
		}
		if got := base62.Encode(tt.ID); got != tt.Text {
			t.Fatalf("want %v, got %v for input %v", tt.Text, got, tt.ID)
		}
	}
}

func TestBase62(t *testing.T) {
	f := func(key uint64) bool {
		return base62.Decode(base62.Encode(key)) == key
	}

	if err := quick.Check(f, nil); err != nil {
		log.Fatal(err)
	}
}

func TestIncrementID(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		id := uint64(i)
		if got := base62.Decode(base62.Encode(id)); got != id {
			log.Fatalf("want %v, got %v", id, got)
		}
	}
}

// func TestCompare(t *testing.T) {
//         f := func(i uint64) bool {
//                 return base62.Encode(i) == base62.EncodeLegacy(i)
//         }
//
//         if err := quick.Check(f, nil); err != nil {
//                 log.Fatal(err)
//         }
// }

// go test -bench=. -benchmem -cpuprofile cpu.out -memprofile mem.out
func BenchmarkEncode(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	for i := 0; i < b.N; i++ {
		base62.Encode(r.Uint64())
	}
}

// func BenchmarkEncodeLegacy(b *testing.B) {
//         r := rand.New(rand.NewSource(0))
//         for i := 0; i < b.N; i++ {
//                 base62.EncodeLegacy(r.Uint64())
//         }
// }
