# substring [![Build Status](https://travis-ci.org/toqueteos/substring.png?branch=master)](https://travis-ci.org/toqueteos/substring) [![GoDoc](http://godoc.org/github.com/toqueteos/substring?status.png)](http://godoc.org/github.com/toqueteos/substring)

Very fast **one-time string searches** in Go. Simple and composable.

Interop with [regexp](http://golang.org/pkg/regexp/) for backwards compatibility (easy migration from your current system to `substring`).

## Installation

The recommended way to install substring is by using `go get`:

```
go get github.com/toqueteos/substring
```

Go Modules are supported!

## Examples

A basic example with two matchers:

```go
package main

import (
	"fmt"
	"regexp"

	"github.com/toqueteos/substring/v2"
)

func main() {
	m1 := substring.After("assets/", substring.Or(
		substring.Has("jquery"),
		substring.Has("angular"),
		substring.Suffixes(".js", ".css", ".html"),
	))
	fmt.Println(m1.Match("assets/angular/foo/bar")) // Prints: true
	fmt.Println(m1.Match("assets/js/file.js"))      // Prints: true
	fmt.Println(m1.Match("assets/style/bar.css"))   // Prints: true
	fmt.Println(m1.Match("assets/foo/bar.html"))    // Prints: true
	fmt.Println(m1.Match("assets/js/qux.json"))     // Prints: false
	fmt.Println(m1.Match("core/file.html"))         // Prints: false
	fmt.Println(m1.Match("foobar/that.jsx"))        // Prints: false
	fmt.Println()

	m2 := substring.After("vendor/", substring.Suffixes(".css", ".js", ".less"))
	fmt.Println(m2.Match("foo/vendor/bar/qux.css")) // Prints: true
	fmt.Println(m2.Match("foo/var/qux.less"))       // Prints: false
	fmt.Println()

	re := regexp.MustCompile(`vendor\/.*\.(css|js|less)$`)
	fmt.Println(re.MatchString("foo/vendor/bar/qux.css")) // Prints: true
	fmt.Println(re.MatchString("foo/var/qux.less"))       // Prints: false
}
```

## How fast?

It may vary depending on your use case but 1~2 orders of magnitude faster than `regexp` is pretty common.

Test it out for yourself by running `go test -bench .`!

```
$ go test -bench .
pkg: github.com/toqueteos/substring
BenchmarkExample1-16            30759529                38.4 ns/op
BenchmarkExample2-16            26659675                40.0 ns/op
BenchmarkExample3-16            30760317                37.7 ns/op
BenchmarkExample4-16            31566652                36.8 ns/op
BenchmarkExample5-16           123704845                9.70 ns/op
BenchmarkExampleRe1-16           2739574                 436 ns/op
BenchmarkExampleRe2-16           2494791                 480 ns/op
BenchmarkExampleRe3-16           1681654                 713 ns/op
BenchmarkExampleRe4-16           2205490                 540 ns/op
BenchmarkExampleRe5-16          19673001                55.0 ns/op
PASS
ok      github.com/toqueteos/substring  15.016s
```

## License

MIT, see [LICENSE](LICENSE)
