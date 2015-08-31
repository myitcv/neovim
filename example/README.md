## `github.com/myitcv/neovim/example`

A bare-bones example of a Neovim plugin written in Go

Writing a plugin currently involves a number of steps which will be automated with time:

1. Defining the equivalent of `Example` (`example.go`), a type that implements `neovim.Plugin`
2. Defining sync and async plugin methods, such as `GetTwoNumbers` and `DoSomethingAsync`
3. Generating (by hand for now) wrappers around these types (`gen_example.go`)
4. Generating (using `msgp`) MSGPACK wrappers around the wrappers (!) (`gen_example_gen.go`)
5. Testing that this all works (`example_test.go`) with Neovim: `go test`
6. Generating (by hand for now) the necessary bootstrap/wrappers to register these functions with Neovim

Clearly all the steps after step 2 should be automated in some fashion. See the
[TODO](https://github.com/myitcv/neovim/wiki/TODO) for a more up-to-date list of what remains to be done.

## Installing and using the `Example` plugin

0. `mkdir -p $HOME/.nvim/plugins/go/`
1. `go get -u github.com/myitcv/neovim/cmd/neovim-go-plugin-manager`
3. `$GOPATH/bin/neovim-go-plugin-manager github.com/myitcv/neovim/example`

Now launch `nvim`

```
nvim -u $GOPATH/src/github.com/myitcv/neovim/example/special.vimrc
```

and try:

```
:echo GetTwoNumbers(5)
```

The output should read:

```
[47, '42']
```

If you tail the equivalent of `/tmp/neovim_go_plugin_host` then you should also see output:

```
$ tail -f /tmp/neovim_go_plugin_host
2015/08/25 11:45:09 /tmp/neovim-go-plugin-manager_1440497199966172499/plugin_host.go:66: Successfully connected to Neovim
2015/08/25 11:45:09 /tmp/neovim-go-plugin-manager_1440497199966172499/plugin_host.go:81: Connecting *example.Example
2015/08/25 11:45:09 /tmp/neovim-go-plugin-manager_1440497199966172499/plugin_host.go:86: Successfully called Init on *example.Example
```

You can also call the async method `DoSomethingAsync` from `nvim`:

```
:call DoSomethingAsync("test_string")
```

This results in the following line being appended to the host log file:

```
2015/08/25 11:45:56 /tmp/neovim-go-plugin-manager_1440497199966172499/plugin_host.go:152: *example.ExampleGot an event: test_string
```
