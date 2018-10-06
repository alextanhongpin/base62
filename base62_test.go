package base62_test

import (
	"log"
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
