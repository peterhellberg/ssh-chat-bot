# hi

[![Build Status](https://travis-ci.org/peterhellberg/hi.svg?branch=master)](https://travis-ci.org/peterhellberg/hi)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/hi)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/hi#license-mit)

Find images for a given hashtag

## Installation

    go get -u github.com/peterhellberg/hi

If you want to install the command line application:

    go get -u github.com/peterhellberg/hi/cmd/hi

## Usage

```go
package main

import (
	"fmt"

	"github.com/peterhellberg/hi"
)

func main() {
	image, err := hi.FindShuffledImage("pixel_dailies")

	if err == nil {
		fmt.Println(image.URL)
	}
}
```

## Dependencies

The scraping is performed using the lovely [github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)

## License (MIT)

Copyright (c) 2015 [Peter Hellberg](http://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
