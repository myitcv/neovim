## neovim

[![Build Status](https://travis-ci.org/myitcv/neovim.svg?branch=master)](https://travis-ci.org/myitcv/neovim)

Go package for writing [Neovim](http://neovim.org/) plugins

```bash
go get github.com/myitcv/neovim
```

This package is very much in alpha as [the Neovim API itsef](https://github.com/neovim/neovim/issues/973)
is still in flux. Therefore, expect changes to this API.

## Overview and Writing plugins

See [the wiki](https://github.com/myitcv/neovim/wiki/Overview-of-writing-and-using-Go-packages-with-Neovim) for more
details.

## Supported platforms

At the time of writing this package has only been written for/tested against Linux.

Support welcomed on other platforms

## Tests

```bash
go test
```

## Credit

* [@tarruda](https://github.com/tarruda) for leading the way with his [python-client](https://github.com/neovim/python-client)
* [@vmihailenco](https://github.com/vmihailenco) for [msgpack](https://github.com/vmihailenco/msgpack)

## Todo list

See [the wiki](https://github.com/myitcv/neovim/wiki/Overview-of-writing-and-using-Go-packages-with-Neovim)
