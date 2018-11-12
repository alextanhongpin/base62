# go-base62

[![](https://godoc.org/github.com/alextanhongpin/base62?status.svg)](http://godoc.org/github.com/alextanhongpin/base62)

Base62 encoder/decoder with golang. Uses the variant `[A-Za-z0-9]`, but can be initialized with custom variant. The test cases covered the variant mentioned previously.

## Installation

```bash
$ go get github.com/alextanhongpin/base62
```

## Usage

Encode/Decode:
```go
import base62 "github.com/alextanhongpin/base62"

func main() {
  fmt.Println(base62.Decode("golang")) // Outputs: 30847375997
  fmt.Println(base62.Encode(30847375997)) // Outputs: "golang"
}
```

Factory:

```go
import base62 "github.com/alextanhongpin/base62"

func main() {
  b62 := base62.New(base62.DEFAULT_CHARS) // Or use your own variant
  fmt.Println(b62.Decode("golang")) // Outputs: 30847375997
  fmt.Println(b62.Encode(30847375997)) // Outputs: "golang"
}
```

## Test Table

To test the encoding/decoding, we use the [NATO Phonetic Alphabets](https://en.wikipedia.org/wiki/NATO_phonetic_alphabet). If the encoded values matches the decoded values, the test passes. There's a `quickcheck` test too to capture unexpected input.

| Text | Decoded |
|------|---------|
| Alfa | 386411 |
| Bravo | 40145909 |
| Charlie | 201958818817 |
| Delta | 66642463 |
| Echo | 1305265 |
| Foxtrot | 379112811228 |
| Golf | 1828288 |
| Hotel | 128160920 |
| India | 142637661 |
| Juliett | 611630563254 |
| Kilo | 2758545 |
| Lima | 2996921 |
| Mike | 3235129 |
| November | 51675855992598 |
| Oscar | 232482994 |
| Papa | 3919667 |
| Quebec | 16276243687 |
| Romeo | 275897375 |
| Sierra | 17931255627 |
| Tango | 302117423 |
| Uniform | 1229975219171 |
| Victor | 20679184986 |
| Whiskey | 1338081975441 |
| Xray | 5890733 |
| Yankee | 23311959173 |
| Zulu | 6379599 |

## Performance test

It's fast. Period. 

```
$ go test -bench=. -benchmem -memprofile mem.out -cpuprofile cpu.out
```

Output:

```
goos: darwin
goarch: amd64
pkg: github.com/alextanhongpin/base62
BenchmarkEncode-4   	10000000	       200 ns/op	      64 B/op	       2 allocs/op
PASS
ok  	github.com/alextanhongpin/base62	2.378s
```
