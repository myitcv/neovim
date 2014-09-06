## neovim

[![Build Status](https://travis-ci.org/myitcv/neovim.svg?branch=master)](https://travis-ci.org/myitcv/neovim)

Go package for writing [Neovim](http://neovim.org/) plugins

```bash
go get github.com/myitcv/neovim
```

This package is very much in alpha as [the Neovim API itsef](https://github.com/neovim/neovim/issues/973)
is still in flux. Therefore, expect changes to this API.

## Writing plugins

A current, idiomatic example of a plugin written against `neovim` is [`neovim-go`](https://github.com/myitcv/neovim-go),
a plugin that is designed to help support the editing of Go files.

A further example can be found [here](https://github.com/myitcv/neovim_example); this example demonstrates on-the-fly
compiling of a Go file.

## Supported platforms

At the time of writing this package has only been written for/tested against Linux.

Support welcomed on other platforms

## Tests

Tests currently rely on the `NEOVIM_LISTEN_ADDRESS` environment variable being set

```bash
NEOVIM_LISTEN_ADDRESS=/tmp/neovim go test
```

## Regenerating the API

First get the generator:

```bash
go get github.com/myitcv/neovim/_cmd/gen_neovim_api
```

The API generator relies on `nvim` being in your `$PATH`. Here are the various options:

```bash
$ ./gen_neovim_api -h
Usage: ./gen_neovim_api [-p] [-g] [-f filename]

  -c=false: custom print
  -f="": file containing the list of API functions to generate
  -g=false: generate code from the API
  -p=false: print the API

One of -p or -g must be supplied

If -g is supplied, -f may also be supplied to provide a list of functions to generate
```

The `-g` flag currently outputs to stdout; future work will provide a flag to have it
place the generated, formatted source in `$GOPATH` as appropriate.

## Credit

* [@tarruda](https://github.com/tarruda) for leading the way with his [python-client](https://github.com/neovim/python-client)
* [@vmihailenco](https://github.com/vmihailenco) for [msgpack](https://github.com/vmihailenco/msgpack)

## Todo list

* Delete remote `plugin-host` branch
* Neovim hooks up stdin and stdout back to front
* Remove `example_test.go` - it is superseded by `example` and associated tests
* Ensure that example gives examples (!) of testing API requests, handling requests and events
* Consider using `internal` in Go 1.4 for better layout of package
* Better definition of logging level-based interface for plugins
* Stop using KillChannel in generated plugin-host; move to Tomb
* Need some way for plugins to be loaded at startup...
* Support Go package acting as a server, i.e. receiving requests from Neovim
* Cleanly handle Neovim instances quitting - need to kill goroutines for reading, subscription manager etc
* Need a test to confirm we have cleaned up all goroutines
* Move from `log.Fatal` to something better....will work now that we have `tomb`
* Tidy up use of `PanicOnError` and ensure all errors that would otherwise have been return use `panic` if this is set
* Ensure tests panic if there is an error so we can see the stack
* Add support for functions that `CanFail`
* Get full test coverage; can we auto-generate certain basic tests?
* Test on more platforms
* Tidy up the API generator - it's very messy. Very messy
* Make the API generator optionally write the generated file into the appropriate place in `$GOPATH`
* Check our use of types; e.g. what does `Integer` in the API really map to? `uint64`?
* Look into semantics of current decision to make certain channels buffered
* Look at how we verify at runtime whether the Neovim instance to which we have connected exposes an API
with which we (a calling client) was compiled. This handshake will probably need some work on the Neovim
side (versioning of the API perhaps?). See comments in [this
thread](https://github.com/neovim/neovim/pull/1121#issuecomment-53873091)
* More benchmark tests for performance
* Make notification decoding can be made more efficient (currently uses generic `DecodeSlice`)
* Fix function and variable names (there are a few anomalies like `Client.GetVvar`)

