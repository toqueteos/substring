package substring

import (
	"testing"

	"github.com/matryer/is"
)

func TestBytesAny(t *testing.T) {
	is := is.New(t)

	a := BytesAny("foo") // search s in foo
	is.Equal(a.MatchIndex([]byte("f")), 1)
	is.Equal(a.MatchIndex([]byte("foo")), 3)
	is.Equal(a.MatchIndex([]byte("foobar")), -1)
	is.Equal(a.MatchIndex([]byte("p")), -1)
}

func TestBytesHas(t *testing.T) {
	is := is.New(t)

	h := BytesHas("foo") // search foo in s
	is.Equal(h.MatchIndex([]byte("foo")), 3)
	is.Equal(h.MatchIndex([]byte("foobar")), 3)
	is.Equal(h.MatchIndex([]byte("f")), -1)
}

func TestBytesPrefix(t *testing.T) {
	is := is.New(t)

	p := BytesPrefix("foo")
	is.True(p.Match([]byte("foo")))
	is.True(p.Match([]byte("foobar")))
	is.Equal(p.Match([]byte("barfoo")), false)
	is.Equal(p.Match([]byte(" foo")), false)
	is.Equal(p.Match([]byte("bar")), false)
	is.Equal(p.MatchIndex([]byte("foo")), 3)
	is.Equal(p.MatchIndex([]byte("foobar")), 3)
	is.Equal(p.MatchIndex([]byte("barfoo")), -1)
	is.Equal(p.MatchIndex([]byte(" foo")), -1)
	is.Equal(p.MatchIndex([]byte("bar")), -1)
	ps := BytesPrefixes("foo", "barfoo")
	is.True(ps.Match([]byte("foo")))
	is.True(ps.Match([]byte("barfoo")))
	is.Equal(ps.Match([]byte("qux")), false)
	is.Equal(ps.MatchIndex([]byte("foo")), 2)
	is.Equal(ps.MatchIndex([]byte("barfoo")), 5)
	is.Equal(ps.MatchIndex([]byte("qux")), -1)
}

func TestBytesSuffix(t *testing.T) {
	is := is.New(t)

	p := BytesSuffix("foo")
	is.True(p.Match([]byte("foo")))
	is.True(p.Match([]byte("barfoo")))
	is.Equal(p.Match([]byte("foobar")), false)
	is.Equal(p.Match([]byte("foo ")), false)
	is.Equal(p.Match([]byte("bar")), false)
	is.Equal(p.MatchIndex([]byte("foo")), 3)
	is.Equal(p.MatchIndex([]byte("barfoo")), 3)
	is.Equal(p.MatchIndex([]byte("foobar")), -1)
	is.Equal(p.MatchIndex([]byte("foo ")), -1)
	is.Equal(p.MatchIndex([]byte("bar")), -1)
	ps := BytesSuffixes("foo", "foobar")
	is.True(ps.Match([]byte("foo")))
	is.True(ps.Match([]byte("foobar")))
	is.Equal(ps.Match([]byte("qux")), false)
	is.Equal(ps.MatchIndex([]byte("foo")), 2)
	is.Equal(ps.MatchIndex([]byte("foobar")), 5)
	is.Equal(ps.MatchIndex([]byte("qux")), -1)
	ps2 := BytesSuffixes(".foo", ".bar", ".qux")
	is.True(ps2.Match([]byte("bar.foo")))
	is.Equal(ps2.Match([]byte("bar.js")), false)
	is.True(ps2.Match([]byte("foo/foo.bar")))
	is.Equal(ps2.Match([]byte("foo/foo.js")), false)
	is.True(ps2.Match([]byte("foo/foo/bar.qux")))
	is.Equal(ps2.Match([]byte("foo/foo/bar.css")), false)
}

func TestBytesExact(t *testing.T) {
	is := is.New(t)

	a := BytesExact("foo")
	is.True(a.Match([]byte("foo")))
	is.Equal(a.Match([]byte("bar")), false)
	is.Equal(a.Match([]byte("qux")), false)
}

func TestBytesAfter(t *testing.T) {
	is := is.New(t)

	a1 := BytesAfter("foo", BytesExact("bar"))
	is.True(a1.Match([]byte("foobar")))
	is.Equal(a1.Match([]byte("foo_bar")), false)
	a2 := BytesAfter("foo", BytesHas("bar"))
	is.True(a2.Match([]byte("foobar")))
	is.True(a2.Match([]byte("foo_bar")))
	is.True(a2.Match([]byte("_foo_bar")))
	is.Equal(a2.Match([]byte("foo_nope")), false)
	is.Equal(a2.Match([]byte("qux")), false)
	a3 := BytesAfter("foo", BytesPrefixes("bar", "qux"))
	is.True(a3.Match([]byte("foobar")))
	is.True(a3.Match([]byte("fooqux")))
	is.Equal(a3.Match([]byte("foo bar")), false)
	is.Equal(a3.Match([]byte("foo_qux")), false)
}

func TestBytesSuffixGroup(t *testing.T) {
	is := is.New(t)

	sg1 := BytesSuffixGroup(".foo", BytesHas("bar"))
	is.True(sg1.Match([]byte("bar.foo")))
	is.True(sg1.Match([]byte("barqux.foo")))
	is.Equal(sg1.Match([]byte(".foo.bar")), false)
	sg2 := BytesSuffixGroup(`.foo`,
		BytesAfter(`bar`, BytesHas("qux")),
	)
	is.True(sg2.Match([]byte("barqux.foo")))
	is.True(sg2.Match([]byte("barbarqux.foo")))
	is.Equal(sg2.Match([]byte("bar.foo")), false)
	is.Equal(sg2.Match([]byte("foo.foo")), false)
	sg3 := BytesSuffixGroup(`.foo`,
		BytesAfter(`bar`, BytesRegexp(`\d+`)),
	)
	is.True(sg3.Match([]byte("bar0.foo")))
	is.Equal(sg3.Match([]byte("bar.foo")), false)
	is.Equal(sg3.Match([]byte("bar0.qux")), false)
}
