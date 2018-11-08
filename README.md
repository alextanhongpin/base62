# go-base62

[![](https://godoc.org/github.com/alextanhongpin/go-base62?status.svg)](http://godoc.org/github.com/alextanhongpin/go-base62)

A simple base62 url shortener implementation

## Performance test

```
$ go test -bench=. -benchmem -memprofile mem.out -cpuprofile cpu.out
```

Output:

```
goos: darwin
goarch: amd64
BenchmarkEncode-4                2000000               979 ns/op             373 B/op         21 allocs/op
BenchmarkEncodePerf-4            5000000               261 ns/op              64 B/op          2 allocs/op
PASS
ok      _/Users/alextanhongpin/Documents/all/base62     4.651s
```
