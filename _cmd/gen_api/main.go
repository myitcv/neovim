package main

import (
	"encoding/json"
	"flag"
	"fmt"
	_log "log"
	"os"
	"strings"
	"text/template"

	"github.com/myitcv/neovim"
	sstrings "github.com/myitcv/strings"
)

var generated_functions = map[string]bool{
	"vim_get_buffers":        true,
	"vim_get_current_buffer": true,
	"buffer_get_length":      true,
	"buffer_get_line":        true,
}

var log = _log.New(os.Stdout, "", _log.Lshortfile)
var elog = _log.New(os.Stderr, "", _log.Lshortfile)

var known_classes = []string{"Buffer", "Window", "Tabpage"}

var f_cprint = flag.Bool("c", false, "custom print")
var f_print = flag.Bool("p", false, "print the API")
var f_gen = flag.Bool("g", false, "generate code from the API")

func showUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %v [-p] [-g]\n\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nOne of -p or -g must be supplied\n")
	os.Exit(1)
}

func main() {
	flag.Usage = showUsage
	flag.Parse()

	if !*f_print && !*f_gen && !*f_cprint {
		showUsage()
	}

	client, err := neovim.NewUnixClient("/tmp/neovim", "unix")
	if err != nil {
		elog.Fatalf("Could not create neovim client: %v\n", err)
	}
	api, err := client.API()
	if err != nil {
		elog.Fatalf("Could not get API from client: %v\n", err)
	}

	switch {
	case *f_print:
		j, err := json.MarshalIndent(api, "", "  ")
		if err != nil {
			elog.Fatalf("Could not marshall JSON: %v\n", err)
		}

		os.Stdout.Write(j)
		os.Stdout.WriteString("\n")
	case *f_gen:
		genAPI(api)
	case *f_cprint:
		for _, v := range api.Functions {
			splits := strings.SplitN(v.Name, "_", 2)
			fmt.Printf("%v: ", splits[0])
			if len(v.Parameters) > 0 {
				fmt.Print(v.Parameters[0])
			} else {
				fmt.Print("nil")
			}
			fmt.Print("\n")
		}
	}
}

type methodTemplate struct {
	Id     uint32
	Name   string
	Rec    variable
	Ret    *variable
	Params []variable
}

type _type struct {
	name      string
	enc       string
	dec       string
	primitive bool
}

func (t *_type) Name() string {
	return t.name
}

func (t *_type) CanEnc() bool {
	return t.enc != ""
}

func (t *_type) Enc() string {
	if t.enc == "" {
		elog.Fatalln("Tried to get Enc method for type that can't be encoded")
	}
	return t.enc
}

func (t *_type) CanDec() bool {
	return t.dec != ""
}

func (t *_type) Dec() string {
	if t.dec == "" {
		elog.Fatalln("Tried to get Dec method for type that can't be decoded")
	}
	return t.dec
}

func (t *_type) Primitive() bool {
	return t.primitive
}

type variable struct {
	name string
	Type _type
}

func (v *variable) Name() string {
	return v.name
}

func (v *variable) Client() string {
	switch v.Type.Name() {
	case "Client":
		return v.Name()
	case "Buffer", "Window", "Tabpage":
		return v.Name() + ".client"
	default:
		log.Fatalf("Don't know how to handle a client request for type %v", v.Type.Name())
	}
	return ""
}

type api struct {
	Methods []methodTemplate
}

func getType(s string) _type {
	res, ok := type_map[s]
	if !ok {
		log.Fatalf("Could not find type for %v", s)
	}
	return res
}

func genMethodTemplates(fs []neovim.APIFunction) []methodTemplate {
	res := make([]methodTemplate, len(fs))

	for i, f := range fs {
		m := methodTemplate{}
		splits := strings.SplitN(f.Name, "_", 2)

		if len(splits) != 2 {
			elog.Fatalf("Function name not as expected: %v\n", f.Name)
		}

		// name
		m.Name = sstrings.Camelize(splits[1])
		m.Id = f.Id

		// receiver
		rec_id := splits[0]
		var rec_type _type
		switch rec_id {
		case "vim":
			rec_type = getType("Client")
		case "buffer", "window", "tabpage":
			rec_type = getType(sstrings.Camelize(rec_id))
		default:
			elog.Fatalf("Do not know how to deal with receiver type %v\n", rec_id)
		}
		m.Rec = variable{
			Type: rec_type,
			name: "recv",
		}

		// return
		ret_type := getType(f.ReturnType)
		m.Ret = &variable{
			Type: ret_type,
			name: "ret_val",
		}

		// params
		// TODO this could be improved
		var of_interest []neovim.APIFunctionParameter
		switch m.Rec.Type.Name() {
		case "Client":
			of_interest = f.Parameters
		case "Buffer", "Window", "Tabpage":
			of_interest = f.Parameters[1:]
		default:
			elog.Fatalf("Don't know how to handle receiver of type %v\n", m.Rec.Type.Name())
		}

		m.Params = make([]variable, len(of_interest))
		for i, v := range of_interest {
			p := getType(v.Type)
			m.Params[i].name = "i_" + v.Name
			m.Params[i].Type = p
		}

		res[i] = m
	}

	return res
}

func genAPI(a *neovim.API) {
	// ensure we only have classes we know about
	comp := make(map[string]int, len(known_classes))
	for _, k := range known_classes {
		comp[k] += 1
	}
	for _, k := range a.Classes {
		if comp[k.Name] != 1 {
			elog.Fatalf("We got an unexpected class: %v\n", k.Name)
		}
	}

	funcs_of_interest := make([]neovim.APIFunction, 0)
	for i, _ := range a.Functions {
		if _, ok := generated_functions[a.Functions[i].Name]; ok {
			funcs_of_interest = append(funcs_of_interest, a.Functions[i])
		}
	}

	if funcs_of_interest == nil {
		elog.Fatalln("Could not find functions of interest")
	}

	t := template.New("api")
	_, err := t.Parse(clientAPITemplate)
	if err != nil {
		elog.Fatalf("Could not parse client API template: %v\n", err)
	}

	api := api{}
	api.Methods = genMethodTemplates(funcs_of_interest)

	err = t.Execute(os.Stdout, api)
	if err != nil {

		// ensure we are on a newline
		fmt.Println()
		elog.Fatalf("Error generating API: %v\n", err)
	}
}

var clientAPITemplate = `
// **** THIS FILE IS GENERATED - DO NOT EDIT BY HAND
package neovim

import "github.com/juju/errgo"

{{range .Methods }}
{{template "meth" .}}
{{end}}

{{define "meth"}}
func {{template "meth_rec" .}} {{ .Name }}({{template "meth_params" .Params}}) {{template "meth_ret" .Ret}} {
	enc := func() (_err error) {
		{{if .Rec.Type.CanEnc}}
		_err = {{.Rec.Name}}.encode()
		if _err != nil {
			return
		}
		{{end}}
		{{ $client := .Rec.Client }}

		{{range .Params}}
		{{if .Type.Primitive}}
		_err = {{$client}}.enc.{{.Type.Enc}}({{.Name}})
		{{else}}
		_err = {{.Name}}.{{.Type.Enc}}
		{{end}}
		if _err != nil {
			return
		}
		{{end}}

		return
	}
	{{if .Ret}}
	dec := func() (_i interface{}, _err error) {
		{{if .Ret.Type.Primitive}} _i, _err = {{.Rec.Client}}.dec.{{.Ret.Type.Dec}}()
		{{else}} _i, _err = {{.Rec.Client}}.{{.Ret.Type.Dec}}(){{end}}
		return
	}
	{{end}}
	resp_chan, err := {{.Rec.Client}}.makeCall({{.Id}}, enc, dec)
	if err != nil {
		return {{.Ret.Name}}, errgo.NoteMask(err, "Could not make call to {{.Rec.Type.Name}}.{{.Name}}")
	}
	resp := <-resp_chan
	if resp == nil {
		return {{.Ret.Name}}, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return {{.Ret.Name}}, errgo.NoteMask(err, "We got a non-nil error in our response")
	}
	{{if .Ret}}
	{{.Ret.Name}} = resp.obj.({{.Ret.Type.Name}})
	return {{.Ret.Name}}, ret_err
	{{else}}
	return ret_err
	{{end}}

}
{{end}}

{{define "meth_rec"}}({{.Rec.Name}} *{{.Rec.Type.Name}}){{end}}

{{define "meth_params"}}{{$join := ""}}{{range .}}{{ $join }}{{ .Name }} {{.Type.Name}}{{$join := ", "}}{{end}}{{end}}

{{define "meth_ret"}}({{if .}}{{.Name}} {{.Type.Name}}, {{end}}ret_err error){{end}}
`

var type_map = map[string]_type{
	"String": {
		name:      "string",
		enc:       "EncodeString",
		dec:       "DecodeString",
		primitive: true,
	},
	"Integer": {
		name:      "int",
		enc:       "EncodeInt",
		dec:       "DecodeInt",
		primitive: true,
	},
	"Buffer": {
		name: "Buffer",
		enc:  "encode",
		dec:  "decodeBuffer",
	},
	"Window": {
		name: "Window",
		enc:  "encode",
		dec:  "decodeWindow",
	},
	"Client": {
		name: "Client",
	},
	"BufferArray": {
		name: "[]Buffer",
		enc:  "encode",
		dec:  "decodeBufferSlice",
	},
}
