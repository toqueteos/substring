package substring

import (
	"regexp"
	"testing"
)

var matcher = After("vendor/", Suffixes(".css", ".js", ".less"))

func BenchmarkExample1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matcher.Match("foo/vendor/bar/qux.css")
	}
}
func BenchmarkExample2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matcher.Match("foo/vendor/bar.foo/qux.css")
	}
}
func BenchmarkExample3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matcher.Match("foo/vendor/bar.foo/qux.jsx")
	}
}
func BenchmarkExample4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matcher.Match("foo/vendor/bar/qux.jsx")
	}
}
func BenchmarkExample5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matcher.Match("foo/var/qux.less")
	}
}

var re = regexp.MustCompile(`vendor\/.*\.(css|js|less)$`)

func BenchmarkExampleRe1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re.MatchString("foo/vendor/bar/qux.css")
	}
}
func BenchmarkExampleRe2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re.MatchString("foo/vendor/bar.foo/qux.css")
	}
}
func BenchmarkExampleRe3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re.MatchString("foo/vendor/bar.foo/qux.jsx")
	}
}
func BenchmarkExampleRe4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re.MatchString("foo/vendor/bar/qux.jsx")
	}
}
func BenchmarkExampleRe5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re.MatchString("foo/var/qux.less")
	}
}
