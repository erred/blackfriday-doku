# blackfriday-doku

a renderer for [blackfriday](https://github.com/russross/blackfriday) to output dokuwiki

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/seankhliao/blackfriday-doku)
[![License](https://img.shields.io/github/license/seankhliao/blackfriday-doku.svg?style=flat-square)](LICENSE)
![Version](https://img.shields.io/github/v/tag/seankhliao/blackfriday-doku?sort=semver&style=flat-square)

## usage

```
import (
        "github.com/russross/blackfriday"
        doku "github.com/seankhliao/blackfriday-dokuwiki"
)

func main() {
        b := ...
        b = blackfriday.Run(b, blackfriday.WithRenderer(doku.NewRenderer()))
}
```
