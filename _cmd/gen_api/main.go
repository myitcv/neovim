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

type variable struct {
	Name      string
	Type      string
	Enc       string
	Dec       string
	EncNeeded bool
}

type methodTemplate struct {
	Name   string
	Rec    variable
	Ret    *variable
	Params []variable
}

type api struct {
	Methods []methodTemplate
}

func genMethodTemplates(fs []neovim.APIFunction) []methodTemplate {
	res := make([]methodTemplate, len(fs))

	for i, f := range fs {
		r := methodTemplate{}
		splits := strings.SplitN(f.Name, "_", 2)

		if len(splits) != 2 {
			elog.Fatalf("Function name not as expected: %v\n", f.Name)
		}

		// name
		r.Name = sstrings.Camelize(splits[1])

		// receiver
		switch splits[0] {
		case "vim":
			r.Rec = variable{
				EncNeeded: false,
				Name:      "rec_c",
				Type:      "*Client",
			}
		case "buffer":
			r.Rec = variable{
				EncNeeded: true,
				Name:      "rec_b",
				Type:      "*Buffer",
				Enc:       "encodeBuffer",
				Dec:       "decodeBuffer",
			}
		case "window":
			r.Rec = variable{
				EncNeeded: true,
				Name:      "rec_w",
				Type:      "*Window",
				Enc:       "encodeWindow",
				Dec:       "decodeWindow",
			}
		default:
			elog.Fatalf("Do not know how to deal with %v\n", splits[0])
		}

		// return
		switch f.ReturnType {
		case "Buffer":
			r.Ret = &variable{
				EncNeeded: true,
				Name:      "ret_b",
				Type:      "Buffer",
				Enc:       "encodeBuffer",
				Dec:       "decodeBuffer",
			}
		case "BufferArray":
			r.Ret = &variable{
				EncNeeded: true,
				Name:      "ret_b",
				Type:      "[]Buffer",
				Enc:       "encodeBufferSlice",
				Dec:       "decodeBufferSlice",
			}
		case "void":
			// we do nothing; Ret is nil
		default:
			elog.Fatalf("Do not know how to deal with %v\n", f.ReturnType)
		}

		// params
		var of_interest []neovim.APIFunctionParameter
		switch r.Rec.Type {
		case "*Client":
			of_interest = f.Parameters
		case "*Buffer", "*Window":
			// we don't need the receiver
			of_interest = f.Parameters[1:]
		default:
			elog.Fatalf("Don't know how to handle receiver of type %v\n", r.Rec.Type)
		}

		r.Params = make([]variable, len(of_interest))
		for i, v := range of_interest {
			switch v.Type {
			case "String":
				r.Params[i] = variable{
					EncNeeded: true,
					Name:      v.Name,
					Type:      "string",
					Enc:       "encodeString",
					Dec:       "decodeString",
				}
			case "Integer":
				r.Params[i] = variable{
					EncNeeded: true,
					Name:      v.Name,
					Type:      "int",
					Enc:       "encodeInt",
					Dec:       "decodeInt",
				}
			default:
				elog.Fatalf("Do not know how to handle parameter type %v\n", v.Type)
			}
		}

		res[i] = r
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

	// gen vim_get_buffers for now
	funcs_of_interest := make([]neovim.APIFunction, 0)
	for i, _ := range a.Functions {
		switch a.Functions[i].Name {
		case "vim_get_buffers", "vim_err_write", "window_set_height":
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
package neovim

import "github.com/juju/errgo"

{{range .Methods }}
{{template "meth" .}}
{{end}}

{{define "meth"}}
func {{template "meth_rec" .}} {{ .Name }}({{template "meth_params" .Params}}) {{template "meth_ret" .Ret}} {
	{{if .Rec.EncNeeded}}
	{{end}}

}
{{end}}

{{define "meth_rec"}}({{.Rec.Name}} {{.Rec.Type}}){{end}}

{{define "meth_params"}}{{$join := ""}}{{range .}}{{ $join }}{{ .Name }} {{.Type}}{{$join := ", "}}{{end}}{{end}}

{{define "meth_ret"}}({{if .}}{{.Type}}, {{end}}error){{end}}
`

/*

   {
     "Name": "vim_get_buffers",
     "ReturnType": "BufferArray",
     "Id": 40,
     "CanFail": false,
     "ReceivesChannelId": false,
     "Parameters": []
   },

*/
