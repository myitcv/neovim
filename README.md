## neovim

[![Build Status](https://travis-ci.org/myitcv/neovim.svg?branch=master)](https://travis-ci.org/myitcv/neovim)

Go package for writing [Neovim](http://neovim.org/) plugins

```bash
go get github.com/myitcv/neovim
```

This package is very much in alpha. Therefore, expect changes to this API. Stuff will break.

## Writing plugins

See:

* the [`Example`](https://github.com/myitcv/neovim/tree/master/example) plugin for a brief `README` on how to
implement your own Go plugin
* [`neogo`](), a proof of concept Neovim plugin written against the `neovim` Go package to support Go development in Neovim. This uses a `go/parser` AST generated from the current buffer to highlight code using `matchaddpos` (as opposed to Neovim's default [regex-based syntax definitions](https://github.com/neovim/neovim/blob/master/runtime/syntax/go.vim))

As explained in the [TODO](https://github.com/myitcv/neovim/wiki/TODO), many of the steps required to create a plugin are not currently automated, but very much on the roadmap.

## Supported platforms

At the time of writing this package has only been written for/tested against Linux.

Support welcomed on other platforms

## Tests

```bash
go test ./...
```

## Credit

* The entire [Neovim](https://github.com/neovim/neovim) team for their work
* [@tarruda](https://github.com/tarruda) for leading the way with his [python-client](https://github.com/neovim/python-client)
* [@philhofer](https://github.com/philhofer) et al for the excellent [`msgp`](https://github.com/tinylib/msgp)

## Todo list

See [the wiki](https://github.com/myitcv/neovim/wiki/TODO)
