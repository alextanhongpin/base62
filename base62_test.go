package base62_test

import (
	"log"
	"math/rand"
	"testing"
	"testing/quick"

	base62 "github.com/alextanhongpin/go-base62"
)

func TestBase62(t *testing.T) {
	f := func(key uint64) bool {
		return base62.Decode(base62.Encode(key)) == key
	}

	if err := quick.Check(f, nil); err != nil {
		log.Fatal(err)
	}
}

func TestCompare(t *testing.T) {
	f := func(i uint64) bool {
		return Encode(i) == EncodePerf(i)
	}

	if err := quick.Check(f, nil); err != nil {
		log.Fatal(err)
	}
}

// go test -bench=. -benchmem -cpuprofile cpu.out -memprofile mem.out
func BenchmarkEncode(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	for i := 0; i < b.N; i++ {
		Encode(r.Uint64())
	}
}

func BenchmarkEncodePerf(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	for i := 0; i < b.N; i++ {
		EncodePerf(r.Uint64())
	}
}
