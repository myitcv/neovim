package main

func main() {

}

type pluginType struct {
	Package, Type string
}

func install(plugins ...string) {
	/*
		we need to collate a list plugins and their exported types
		things to report (and test)
		1. plugins that don't export any types
		2. plugins that don't compile (or pass tests)

		We should output for info the types we have found on a per package basis

		The steps required to achieve this:

		1. go get -d the plugin
		2. go test the plugin (will check it compiles)
	*/

}
