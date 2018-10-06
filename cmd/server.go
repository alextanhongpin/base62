package main

import (
	"fmt"
	"math/rand"
	"time"

	base62 "github.com/alextanhongpin/go-base62"
)

type URLEntry struct {
	CreatedAt time.Time
	ID        uint64
	LongURL   string
	ShortURL  string
}

type URLShortener interface {
	Shorten(longURL string) *URLEntry
}

type Generator func() uint64

type shortenerImpl struct {
	generator Generator
}

func NewURLShortener() *shortenerImpl {
	return &shortenerImpl{
		generator: func() uint64 {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			return r.Uint64()
		},
	}
}

func (s *shortenerImpl) Shorten(longURL string) *URLEntry {
	// Generate auto-incremented id.
	id := s.generator()
	shortPath := base62.Encode(id)
	return &URLEntry{
		CreatedAt: time.Now().UTC(),
		ID:        id,
		LongURL:   longURL,
		ShortURL:  fmt.Sprintf("http://go/%s", shortPath),
	}
}

func main() {
	hash := base62.Encode(0)
	fmt.Println(hash)
	fmt.Println(base62.Decode(hash))

	s := NewURLShortener()
	entry := s.Shorten("http://google.com")
	fmt.Println(entry)
}
