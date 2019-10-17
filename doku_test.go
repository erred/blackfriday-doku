package doku

import (
	"testing"

	"github.com/russross/blackfriday"
)

func TestRenderer(t *testing.T) {
	cases := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			"paragraph",
			[]byte(`
para 1

this is 
another paragraph


third one
                        `),
			[]byte(`
para 1

this is
another paragraph

third one
`),
		}, {
			"styling",
			[]byte(`
this is **bold** and this is *emphasized*,
this is ~~strike~~ and this is **_~~all in one~~_**
                        `),
			[]byte(`
this is **bold** and this is //emphasized//,
this is <del>strike</del> and this is **//<del>all in one</del>//**
`),
		}, {
			"code",
			[]byte(
				"this is some `inline code` here" +
					"\n```\n" +
					`and
        some
        multiline
        ~~code~~
` + "```"),
			[]byte(`
this is some ''inline code'' here

<code>
and
        some
        multiline
        ~~code~~
</code>
`),
		}, {
			"ruler",
			[]byte("---"),
			[]byte("\n----\n"),
		}, {
			"link",
			[]byte(`
[link text](https://example.com)

https://example.com
`),
			[]byte(`
[[https://example.com|link text]]

[[https://example.com|https://example.com]]
`),
		}, {
			"lists",
			[]byte(`
- unordered 1
  - sub unordered 1
  - sub unordered 2
- unordered 2
  1. sub ordered 1
  2. sub ordered 2

new list

1. ordered 1
   - sub unordered 1
   - sub unordered 2
2. ordered 2
   1. sub ordered 1
   2. sub ordered 2
`),
			[]byte(`
* unordered 1
  * sub unordered 1
  * sub unordered 2
* unordered 2
  - sub ordered 1
  - sub ordered 2

new list

- ordered 1
  * sub unordered 1
  * sub unordered 2
- ordered 2
  - sub ordered 1
  - sub ordered 2
`),
		}, {
			"quotes",
			[]byte(`
unquoted

> but this
> is on 2 lines
> > 2 levels
> > > 3 deep
> > > 3x2=6
> back to one

out
`),
			[]byte(`
unquoted

> but this
> is on 2 lines
> > 2 levels
> > > 3 deep
> > > 3x2=6
> back to one

out
`),
		}, {
			"table",
			[]byte(`
| Tables        | Are           | Cool  |
| ------------- |:-------------:| -----:|
| col 3 is      | right-aligned | $1600 |
| col 2 is      | centered      |   $12 |
| zebra stripes | are neat      |    $1 |
                        `),
			[]byte(`

^Tables ^ Are ^ Cool^
|col 3 is | right-aligned | $1600|
|col 2 is | centered | $12|
|zebra stripes | are neat | $1|
`),
		}, {
			"headings",
			[]byte(`
# heading 1
## h2
`),
			[]byte(`
====== heading 1 ======

===== h2 =====
`),
		},
	}

	for i, c := range cases {
		o := blackfriday.Run(c.input, blackfriday.WithRenderer(NewRenderer()), blackfriday.WithExtensions(Exts))
		if len(c.expected) != len(o) {
			t.Errorf("%d: %s failed len mismatch expected %d got %d", i, c.name, len(c.expected), len(o))
		}
		s := len(c.expected)
		if len(o) < s {
			s = len(o)
		}
		for j := 0; j < s; j++ {
			b := c.expected[j]
			if b != o[j] {
				t.Errorf("%d: %s byte mismatch at %d expected >%s< got >%s<\n%s\n======\n%s\n", i, c.name, j, string(b), string(o[j]), string(c.expected), string(o))

				break
			}
		}
	}
}
