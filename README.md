## neovim

Go package for writing [neovim](http://neovim.org/) plugins in Go

## Writing plugins

## Tests

## Todo list

* Get full test coverage; can we auto-generate certain basic tests?
* Check out use of types; e.g. should we be using `msgpack.DecodeUint64` everywhere instead of plain `Int`
* More tests around concurrent use of a `Client`
* Look into semantics of current decision to make `sub_events` a buffered channel
* Ability to cleanly terminate a client; along with associated tests
* Look at how we verify at runtime whether the Neovim instance to which we have connected exposes an API
with which we (a calling client) was compiled. This handshake will probably need some work on the Neovim
side (versioning of the API perhaps?)
* Benchmark tests for performance
* Look into whether notification decoding can be made more efficient (current uses `DecodeInterface`)
* Make API generator more robust

