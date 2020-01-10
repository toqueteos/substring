package substring

import (
	"testing"

	"github.com/matryer/is"
)

func TestAny(t *testing.T) {
	is := is.New(t)

	a := Any("foo") // search s in foo
	is.Equal(a.MatchIndex("f"), 1)
	is.Equal(a.MatchIndex("foo"), 3)
	is.Equal(a.MatchIndex("foobar"), -1)
	is.Equal(a.MatchIndex("p"), -1)
}

func TestHas(t *testing.T) {
	is := is.New(t)

	h := Has("foo") // search foo in s
	is.Equal(h.MatchIndex("foo"), 3)
	is.Equal(h.MatchIndex("foobar"), 3)
	is.Equal(h.MatchIndex("f"), -1)
}

func TestPrefix(t *testing.T) {
	is := is.New(t)

	p := Prefix("foo")
	is.True(p.Match("foo"))
	is.True(p.Match("foobar"))
	is.Equal(p.Match("barfoo"), false)
	is.Equal(p.Match(" foo"), false)
	is.Equal(p.Match("bar"), false)
	is.Equal(p.MatchIndex("foo"), 3)
	is.Equal(p.MatchIndex("foobar"), 3)
	is.Equal(p.MatchIndex("barfoo"), -1)
	is.Equal(p.MatchIndex(" foo"), -1)
	is.Equal(p.MatchIndex("bar"), -1)
	ps := Prefixes("foo", "barfoo")
	is.True(ps.Match("foo"))
	is.True(ps.Match("barfoo"))
	is.Equal(ps.Match("qux"), false)
	is.Equal(ps.MatchIndex("foo"), 2)
	is.Equal(ps.MatchIndex("barfoo"), 5)
	is.Equal(ps.MatchIndex("qux"), -1)
}

func TestSuffix(t *testing.T) {
	is := is.New(t)

	p := Suffix("foo")
	is.True(p.Match("foo"))
	is.True(p.Match("barfoo"))
	is.Equal(p.Match("foobar"), false)
	is.Equal(p.Match("foo "), false)
	is.Equal(p.Match("bar"), false)
	is.Equal(p.MatchIndex("foo"), 3)
	is.Equal(p.MatchIndex("barfoo"), 3)
	is.Equal(p.MatchIndex("foobar"), -1)
	is.Equal(p.MatchIndex("foo "), -1)
	is.Equal(p.MatchIndex("bar"), -1)
	ps1 := Suffixes("foo", "foobar")
	is.True(ps1.Match("foo"))
	is.True(ps1.Match("foobar"))
	is.Equal(ps1.Match("qux"), false)
	is.Equal(ps1.MatchIndex("foo"), 2)
	is.Equal(ps1.MatchIndex("foobar"), 5)
	is.Equal(ps1.MatchIndex("qux"), -1)
	ps2 := Suffixes(".foo", ".bar", ".qux")
	is.True(ps2.Match("bar.foo"))
	is.Equal(ps2.Match("bar.js"), false)
	is.True(ps2.Match("foo/foo.bar"))
	is.Equal(ps2.Match("foo/foo.js"), false)
	is.True(ps2.Match("foo/foo/bar.qux"))
	is.Equal(ps2.Match("foo/foo/bar.css"), false)
}

func TestExact(t *testing.T) {
	is := is.New(t)

	a := Exact("foo")
	is.True(a.Match("foo"))
	is.Equal(a.Match("bar"), false)
	is.Equal(a.Match("qux"), false)
}

func TestAfter(t *testing.T) {
	is := is.New(t)

	a1 := After("foo", Exact("bar"))
	is.True(a1.Match("foobar"))
	is.Equal(a1.Match("foo_bar"), false)
	a2 := After("foo", Has("bar"))
	is.True(a2.Match("foobar"))
	is.True(a2.Match("foo_bar"))
	is.True(a2.Match("_foo_bar"))
	is.Equal(a2.Match("foo_nope"), false)
	is.Equal(a2.Match("qux"), false)
	a3 := After("foo", Prefixes("bar", "qux"))
	is.True(a3.Match("foobar"))
	is.True(a3.Match("fooqux"))
	is.Equal(a3.Match("foo bar"), false)
	is.Equal(a3.Match("foo_qux"), false)
}

func TestSuffixGroup(t *testing.T) {
	is := is.New(t)

	sg1 := SuffixGroup(".foo", Has("bar"))
	is.True(sg1.Match("bar.foo"))
	is.True(sg1.Match("barqux.foo"))
	is.Equal(sg1.Match(".foo.bar"), false)
	sg2 := SuffixGroup(`.foo`,
		After(`bar`, Has("qux")),
	)
	is.True(sg2.Match("barqux.foo"))
	is.True(sg2.Match("barbarqux.foo"))
	is.Equal(sg2.Match("bar.foo"), false)
	is.Equal(sg2.Match("foo.foo"), false)
	sg3 := SuffixGroup(`.foo`,
		After(`bar`, Regexp(`\d+`)),
	)
	is.True(sg3.Match("bar0.foo"))
	is.Equal(sg3.Match("bar.foo"), false)
	is.Equal(sg3.Match("bar0.qux"), false)
}
