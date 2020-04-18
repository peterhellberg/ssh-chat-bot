flip (╯°□°）╯︵ʇxǝʇ
===================

Go library used to flip text.

[![Build Status](https://travis-ci.org/peterhellberg/flip.svg?branch=master)](https://travis-ci.org/peterhellberg/flip)
[![Go Report Card](https://goreportcard.com/badge/github.com/peterhellberg/flip)](https://goreportcard.com/report/github.com/peterhellberg/flip)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/flip)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/flip#license-mit)

## Command line tool

### Installation

```bash
go get -u github.com/peterhellberg/flip/cmd/flip
```

### Usage

You can flip a string on its own or decorate it with a named emoticon

```bash
flip foo        #=> ooɟ
flip table bar  #=> (╯°□°）╯︵ɹɐq
flip gopher baz #=> ʕ╯◔ϖ◔ʔ╯︵zɐq
```

You can also specify a custom emoticon

```bash
flip '(＃｀д´)ﾉ︵' baz  #=> (＃｀д´)ﾉ︵zɐq
```

## Examples

### table.go

```go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/peterhellberg/flip"
)

func main() {
	fmt.Println(flip.Table(strings.Join(os.Args[1:], " ")))
}
```

## License (MIT)

*Copyright (C) 2014-2019 [Peter Hellberg](https://c7.se)*

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the "Software"),
> to deal in the Software without restriction, including without limitation
> the rights to use, copy, modify, merge, publish, distribute, sublicense,
> and/or sell copies of the Software, and to permit persons to whom the
> Software is furnished to do so, subject to the following conditions:
>
> The above copyright notice and this permission notice shall be included
> in all copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
> OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
> IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
> DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
> TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
> OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
