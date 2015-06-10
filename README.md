# substring [![Build Status](https://travis-ci.org/toqueteos/substring.png?branch=master)](https://travis-ci.org/toqueteos/substring) [![GoDoc](http://godoc.org/github.com/toqueteos/substring?status.png)](http://godoc.org/github.com/toqueteos/substring) [![GitHub release](https://img.shields.io/github/release/toqueteos/substring.svg)](https://github.com/toqueteos/substring/releases)

Simple and composable alternative to [regexp](http://golang.org/pkg/regexp/) package for fast substring searches.

## Installation

The recommended way to install substring

```
go get gopkg.in/toqueteos/substring.v1
```

Examples
--------

A basic example:

```go
package main

import (
    "fmt"
    "regexp"

    "gopkg.in/toqueteos/substring.v1"
)

func main() {
    m1 := substring.After("assets/", substring.Or(
        substring.Has("jquery"),
        substring.Has("angular"),
        substring.Suffixes(
            ".js",
            ".css",
            ".html",
        ),
    ))
    fmt.Println(m1.Match("assets/angular/foo/bar")) //Prints: true
    fmt.Println(m1.Match("assets/js/file.js"))      //Prints: true
    fmt.Println(m1.Match("assets/style/bar.css"))   //Prints: true
    fmt.Println(m1.Match("assets/foo/bar.html"))    //Prints: false
    fmt.Println(m1.Match("assets/js/qux.json"))     //Prints: false
    fmt.Println(m1.Match("core/file.html"))         //Prints: false
    fmt.Println(m1.Match("foobar/that.jsx"))        //Prints: false

    m2 := substring.After("vendor/", substring.Suffixes(
        ".css",
        ".js",
        ".less",
    ))

    fmt.Println(m2.Match("foo/vendor/bar/qux.css")) //Prints: true
    fmt.Println(m2.Match("foo/var/qux.less"))       //Prints: false

    // Is way faster than...
    re := regexp.MustCompile(`vendor\/[^\.]*\.(css|js|less)`)
    fmt.Println(re.MatchString("foo/vendor/bar/qux.css")) //Prints: true
    fmt.Println(re.MatchString("foo/var/qux.less"))       //Prints: false
}
```

License
-------

MIT, see [LICENSE](LICENSE)
