// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	_log "log"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/template"

	"github.com/juju/errgo"
	sstrings "github.com/myitcv/strings"
	"github.com/vmihailenco/msgpack"
)

// An API represents the API as advertised by Neovim
type API struct {
	Classes   []APIClass
	Functions []APIFunction
}

// An APIClass represents a class as defined as part of the API
type APIClass struct {
	Name string
}

// An APIFunction represents a class as defined as part of the API
type APIFunction struct {
	Name              string
	ReturnType        string
	ID                uint32
	CanFail           bool
	ReceivesChannelID bool
	Parameters        []APIFunctionParameter
}

// An APIFunctionParameter represents a function parameters as defined by an APIFunction
type APIFunctionParameter struct {
	Type, Name string
}

var generatedFunctions map[string]bool

var log = _log.New(os.Stdout, "", _log.Lshortfile)

var knowClasses = []string{"Buffer", "Window", "Tabpage"}

var fCustomPrint = flag.Bool("c", false, "custom print")
var fPrintAPI = flag.Bool("p", false, "print the API")
var fGenAPI = flag.Bool("g", false, "generate code from the API")
var fGenList = flag.String("f", "", "file containing the list of API functions to generate")

func showUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %v [-p] [-g] [-f filename]\n\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nOne of -p or -g must be supplied\n")
	fmt.Fprintf(os.Stderr, "\nIf -g is supplied, -f may also be supplied to provide a list of functions to generate\n")
	os.Exit(1)
}

func main() {
	flag.Usage = showUsage
	flag.Parse()

	if !*fPrintAPI && !*fGenAPI && !*fCustomPrint {
		showUsage()
	}

	api, err := getAPI()
	if err != nil {
		log.Fatalf("Could not get API from client: %v\n", err)
	}

	switch {
	case *fPrintAPI:
		j, err := json.MarshalIndent(api, "", "  ")
		if err != nil {
			log.Fatalf("Could not marshall JSON: %v\n", err)
		}

		os.Stdout.Write(j)
		os.Stdout.WriteString("\n")
	case *fGenAPI:
		loadGeneratedFunctions()
		genAPI(api)
	case *fCustomPrint:
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

func getAPI() (*API, error) {
	output, err := exec.Command(os.Getenv("NEOVIM_BIN"), "--api-msgpack-metadata").CombinedOutput()
	if err != nil {
		log.Fatalf("Could not get current API dump: %v", errgo.Details(err))
	}

	br := bytes.NewReader(output)
	ad := msgpack.NewDecoder(br)

	ml, err := ad.DecodeMapLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode map length")

	}

	resp := &API{}

	for i := 0; i < ml; i++ {
		k, err := ad.DecodeString()
		if err != nil {
			return nil, errgo.NoteMask(err, "Could not decode key of top level api map")
		}

		switch k {
		case "classes":
			classes, err := decodeAPIClassSlice(ad)
			if err != nil {
				return nil, errgo.NoteMask(err, "Could not decode class slice")
			}
			resp.Classes = classes
		case "functions":
			functions, err := decodeAPIFunctionSlice(ad)
			if err != nil {
				return nil, errgo.NoteMask(err, "Could not decode function slice")
			}
			resp.Functions = functions
		}
	}

	return resp, nil
}

func decodeAPIClass(d *msgpack.Decoder) (APIClass, error) {
	resp := APIClass{}
	cn, err := d.DecodeString()
	if err != nil {
		return resp, errgo.NoteMask(err, "Could not decode class name")
	}

	resp.Name = cn
	return resp, nil
}

func decodeAPIClassSlice(d *msgpack.Decoder) ([]APIClass, error) {
	sl, err := d.DecodeSliceLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode slice length")
	}

	resp := make([]APIClass, sl)

	for i := 0; i < sl; i++ {
		nvc, err := decodeAPIClass(d)
		if err != nil {
			return nil, errgo.Notef(err, "Could not decode class at index %v", i)
		}
		resp[i] = nvc
	}
	return resp, nil
}

func decodeAPIFunction(d *msgpack.Decoder) (APIFunction, error) {
	resp := APIFunction{}
	ml, err := d.DecodeMapLen()
	if err != nil {
		return resp, errgo.NoteMask(err, "Could not decode map length")
	}

	for i := 0; i < ml; i++ {
		k, err := d.DecodeString()
		if err != nil {
			return resp, errgo.NoteMask(err, "Could not decode function property key")
		}

		switch k {
		case "name":
			s, err := d.DecodeString()
			if err != nil {
				return resp, errgo.NoteMask(err, "Could not decode function name")
			}
			resp.Name = s
		case "receives_channel_id":
			b, err := d.DecodeBool()
			if err != nil {
				return resp, errgo.NoteMask(err, "Could not decode function receives_channel_id")
			}
			resp.ReceivesChannelID = b
		case "can_fail":
			b, err := d.DecodeBool()
			if err != nil {
				return resp, errgo.NoteMask(err, "Could not decode function can_fail")
			}
			resp.CanFail = b
		case "return_type":
			s, err := d.DecodeString()
			if err != nil {
				return resp, errgo.NoteMask(err, "Could not decode function return type")
			}
			resp.ReturnType = s
		case "id":
			i, err := d.DecodeUint32()
			if err != nil {
				return resp, errgo.NoteMask(err, "Could not decode function id")
			}
			resp.ID = i
		case "parameters":
			ps, err := decodeAPIFunctionParameterSlice(d)
			if err != nil {
				return resp, errgo.NoteMask(err, "Could not decode function parameters")
			}
			resp.Parameters = ps
		default:
			return resp, errgo.Newf("Unknown function property %v", k)
		}
	}

	return resp, nil
}

func decodeAPIFunctionSlice(d *msgpack.Decoder) ([]APIFunction, error) {
	sl, err := d.DecodeSliceLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode slice length")
	}

	resp := make([]APIFunction, sl)

	for i := 0; i < sl; i++ {
		nvc, err := decodeAPIFunction(d)
		if err != nil {
			return nil, errgo.Notef(err, "Could not decode function at index %v", i)
		}
		resp[i] = nvc
	}
	return resp, nil
}

func decodeAPIFunctionParameter(d *msgpack.Decoder) (APIFunctionParameter, error) {
	resp := APIFunctionParameter{}

	// we should have a slice of length 2
	sl, err := d.DecodeSliceLen()
	if err != nil {
		return resp, errgo.NoteMask(err, "Could not decode slice length")
	}

	if sl != 2 {
		return resp, errgo.Newf("Expected lenght to be 2; got %v", sl)
	}

	pt, err := d.DecodeString()
	if err != nil {
		return resp, errgo.NoteMask(err, "Could not decode class name")
	}
	resp.Type = pt
	pn, err := d.DecodeString()
	if err != nil {
		return resp, errgo.NoteMask(err, "Could not decode class name")
	}
	resp.Name = pn
	return resp, nil
}

func decodeAPIFunctionParameterSlice(d *msgpack.Decoder) ([]APIFunctionParameter, error) {
	sl, err := d.DecodeSliceLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode slice length")
	}

	resp := make([]APIFunctionParameter, sl)

	for i := 0; i < sl; i++ {
		nvc, err := decodeAPIFunctionParameter(d)
		if err != nil {
			return nil, errgo.Notef(err, "Could not decode function parameter at index %v", i)
		}
		resp[i] = nvc
	}
	return resp, nil
}

func loadGeneratedFunctions() {
	if *fGenList != "" {
		generatedFunctions = make(map[string]bool)
		f, err := os.Open(*fGenList)
		if err != nil {
			log.Fatalf("Could not open -f supplied file to read list of API functions: %v\n", err)
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fn := strings.TrimSpace(scanner.Text())
			generatedFunctions[fn] = true
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading from the -f supplied file: %v\n", err)
		}
	}
}

type typeTemplate struct {
	Name string
}

type methodTemplate struct {
	ID      uint32
	Name    string
	RawName string
	Rec     variable
	Ret     *variable
	Params  []variable
}

func (m *methodTemplate) NumParams() (res int) {
	if m.Rec.Type.CanEnc() {
		res++
	}
	res += len(m.Params)
	return
}

type _type struct {
	name      string
	enc       string
	dec       string
	primitive bool
	genHelper bool
}

func (t *_type) Name() string {
	return t.name
}

func (t *_type) CanEnc() bool {
	return t.enc != ""
}

func (t *_type) Enc() string {
	if t.enc == "" {
		log.Fatalln("Tried to get Enc method for type that can't be encoded")
	}
	return t.enc
}

func (t *_type) CanDec() bool {
	return t.dec != ""
}

func (t *_type) Dec() string {
	if t.dec == "" {
		log.Fatalln("Tried to get Dec method for type that can't be decoded")
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
	Types   []_type
}

func getType(s string) _type {
	res, ok := typeMap[s]
	if !ok {
		log.Fatalf("Could not find type for %v", s)
	}
	return res
}

func genMethodTemplates(fs []APIFunction) []methodTemplate {
	res := make([]methodTemplate, len(fs))

	for i, f := range fs {
		m := methodTemplate{}
		splits := strings.SplitN(f.Name, "_", 2)

		if len(splits) != 2 {
			log.Fatalf("Function name not as expected: %v\n", f.Name)
		}

		// name
		m.RawName = f.Name
		m.Name = sstrings.Camelize(splits[1])
		m.ID = f.ID

		// receiver
		recID := splits[0]
		var recType _type
		switch recID {
		case "vim":
			recType = getType("Client")
		case "buffer", "window", "tabpage":
			recType = getType(sstrings.Camelize(recID))
		default:
			log.Fatalf("Do not know how to deal with receiver type %v\n", recID)
		}
		m.Rec = variable{
			Type: recType,
			name: strings.ToLower(string(recType.Name()[0])),
		}

		// return
		if f.ReturnType != "void" {
			retType := getType(f.ReturnType)
			m.Ret = &variable{
				Type: retType,
				name: "retVal",
			}
		}

		// params
		// TODO this could be improved
		var ofInterest []APIFunctionParameter
		switch m.Rec.Type.Name() {
		case "Client":
			ofInterest = f.Parameters
		case "Buffer", "Window", "Tabpage":
			ofInterest = f.Parameters[1:]
		default:
			log.Fatalf("Don't know how to handle receiver of type %v\n", m.Rec.Type.Name())
		}

		m.Params = make([]variable, len(ofInterest))
		for i, v := range ofInterest {
			p := getType(v.Type)
			m.Params[i].name = sstrings.Camelize(v.Name)
			m.Params[i].name = strings.ToLower(string(m.Params[i].name[0])) + m.Params[i].name[1:]
			m.Params[i].Type = p
		}

		res[i] = m
	}

	sort.Sort(byMethod(res))

	return res
}

type byType []_type

func (a byType) Len() int           { return len(a) }
func (a byType) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byType) Less(i, j int) bool { return a[i].name < a[j].name }

type byMethod []methodTemplate

func (a byMethod) Len() int           { return len(a) }
func (a byMethod) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byMethod) Less(i, j int) bool { return a[i].RawName < a[j].RawName }

func genTypeTemplates(ts []APIClass) []_type {
	var res []_type
	for _, v := range typeMap {
		if v.genHelper {
			res = append(res, v)
		}
	}
	sort.Sort(byType(res))
	return res
}

func genAPI(a *API) {
	// ensure we only have classes we know about
	comp := make(map[string]int, len(knowClasses))
	for _, k := range knowClasses {
		comp[k]++
	}
	for _, k := range a.Classes {
		if comp[k.Name] != 1 {
			log.Fatalf("We got an unexpected class: %v\n", k.Name)
		}
	}

	var funcsOfInterest []APIFunction
	if generatedFunctions != nil {
		funcsOfInterest = make([]APIFunction, 0)
		for i := range a.Functions {
			if _, ok := generatedFunctions[a.Functions[i].Name]; ok {
				funcsOfInterest = append(funcsOfInterest, a.Functions[i])
			}
		}
	} else {
		funcsOfInterest = a.Functions
	}

	if funcsOfInterest == nil {
		log.Fatalln("Could not find functions of interest")
	}

	fm := make(template.FuncMap)
	fm["camelize"] = sstrings.Camelize
	fm["to_lower"] = strings.ToLower

	t := template.New("api")
	t.Funcs(fm)
	_, err := t.Parse(clientAPITemplate)
	if err != nil {
		log.Fatalf("Could not parse client API template: %v\n", err)
	}

	api := api{}
	api.Methods = genMethodTemplates(funcsOfInterest)
	api.Types = genTypeTemplates(a.Classes)

	err = t.Execute(os.Stdout, api)
	if err != nil {

		// ensure we are on a newline
		fmt.Println()
		log.Fatalf("Error generating API: %v\n", err)
	}
}

var clientAPITemplate = `
package neovim

// **** THIS FILE IS GENERATED - DO NOT EDIT BY HAND

import "github.com/juju/errgo"

// constants representing method ids

const (
	{{range .Methods }}{{.Rec.Type.Name | to_lower}}{{.Name}} = "{{.RawName}}"
	{{end}}
)

// methods on the API

{{range .Methods }}
{{template "meth" .}}
{{end}}

// helper functions for types

{{range .Types}}
{{template "type" .}}
{{end}}

{{define "type"}}
func (c *Client) encode{{.Name | camelize }}Slice(s []{{.Name}}) error {
	err := c.enc.EncodeSliceLen(len(s))
	if err != nil {
		return errgo.NoteMask(err, "Could not encode slice length")
	}

	for i := 0; i < len(s); i++ {
		{{if .Primitive}}
		err := c.enc.{{.Enc}}(s[i])
		{{else}}
		err := c.{{.Enc}}(s[i])
		{{end}}
		if err != nil {
			return errgo.Notef(err, "Could not encode {{.Name}} at index %v", i)
		}
	}

	return  nil
}

func (c *Client) decode{{.Name | camelize }}Slice() ([]{{.Name}}, error) {
	l, err := c.dec.DecodeSliceLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode slice length")
	}

	res := make([]{{.Name}}, l)

	for i := 0; i < l; i++ {
		{{if .Primitive}}
		b, err := c.dec.{{.Dec}}()
		{{else}}
		b, err := c.{{.Dec}}()
		{{end}}
		if err != nil {
			return nil, errgo.Notef(err, "Could not decode {{.Name}} at index %v", i)
		}
		res[i] = b
	}

	return res, nil
}
{{end}}

{{define "meth"}}
// {{ .Name }} waiting for documentation from Neovim
func {{template "meth_rec" .}} {{ .Name }}({{template "meth_params" .Params}}) {{template "meth_ret" .Ret}} {
	{{if .Ret}}var {{.Ret.Name}} {{.Ret.Type.Name}}{{end}}
	enc := func() (_err error) {
		_err = {{.Rec.Client}}.enc.EncodeSliceLen({{.NumParams}})
		if _err != nil {
			return
		}
		{{if .Rec.Type.CanEnc}}
		_err = {{.Rec.Client}}.{{.Rec.Type.Enc}}(*{{.Rec.Name}})
		if _err != nil {
			return
		}
		{{end}}
		{{ $client := .Rec.Client }}

		{{range .Params}}
		{{if .Type.Primitive}}
		_err = {{$client}}.enc.{{.Type.Enc}}({{.Name}})
		{{else}}
		_err = {{$client}}.{{.Type.Enc}}({{.Name}})
		{{end}}
		if _err != nil {
			return
		}
		{{end}}

		return
	}
	dec := func() (_i interface{}, _err error) {
		{{if .Ret}}
		{{if .Ret.Type.CanDec}}
		{{if .Ret.Type.Primitive}} _i, _err = {{.Rec.Client}}.dec.{{.Ret.Type.Dec}}()
		{{else}} _i, _err = {{.Rec.Client}}.{{.Ret.Type.Dec}}(){{end}}
		{{end}}
		{{else}}
		_, _err = {{.Rec.Client}}.dec.DecodeBytes()
		{{end}}
		return
	}
	respChan, err := {{.Rec.Client}}.makeCall({{.Rec.Type.Name | to_lower}}{{.Name}}, enc, dec)
	if err != nil {
		return {{if .Ret}}{{.Ret.Name}}, {{end}}{{.Rec.Client}}.panicOrReturn(errgo.NoteMask(err, "Could not make call to {{.Rec.Type.Name}}.{{.Name}}"))
	}
	resp := <-respChan
	if resp == nil {
		return {{if .Ret}}{{.Ret.Name}}, {{end}}{{.Rec.Client}}.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return {{if .Ret}}{{.Ret.Name}}, {{end}}{{.Rec.Client}}.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}
	{{if .Ret}}
	{{.Ret.Name}} = resp.obj.({{.Ret.Type.Name}})
	return {{.Ret.Name}}, nil
	{{else}}
	return nil
	{{end}}

}
{{end}}

{{define "meth_rec"}}({{.Rec.Name}} *{{.Rec.Type.Name}}){{end}}

{{define "meth_params"}}{{range $index, $element := .}}{{if gt $index 0}}, {{end}}{{ .Name }} {{.Type.Name}}{{end}}{{end}}

{{define "meth_ret"}}({{if .}}{{.Type.Name}}, {{end}} error){{end}}
`

var typeMap = map[string]_type{
	"String": {
		name:      "string",
		enc:       "EncodeString",
		dec:       "DecodeString",
		primitive: true,
		genHelper: true,
	},
	"StringArray": {
		name: "[]string",
		enc:  "encodeStringSlice",
		dec:  "decodeStringSlice",
	},
	"Position": {
		name:      "uint32",
		enc:       "EncodeUint32",
		dec:       "DecodeUint32",
		primitive: true,
	},
	"Integer": {
		name:      "int",
		enc:       "EncodeInt",
		dec:       "DecodeInt",
		primitive: true,
	},
	"Boolean": {
		name:      "bool",
		enc:       "EncodeBool",
		dec:       "DecodeBool",
		primitive: true,
	},
	"Object": {
		name:      "interface{}",
		enc:       "Encode",
		dec:       "DecodeInterface",
		primitive: true,
	},
	"Buffer": {
		name:      "Buffer",
		enc:       "encodeBuffer",
		dec:       "decodeBuffer",
		genHelper: true,
	},
	"BufferArray": {
		name: "[]Buffer",
		enc:  "encodeBufferSlice",
		dec:  "decodeBufferSlice",
	},
	"Window": {
		name:      "Window",
		enc:       "encodeWindow",
		dec:       "decodeWindow",
		genHelper: true,
	},
	"WindowArray": {
		name: "[]Window",
		enc:  "encodeWindowSlice",
		dec:  "decodeWindowSlice",
	},
	"Tabpage": {
		name:      "Tabpage",
		enc:       "encodeTabpage",
		dec:       "decodeTabpage",
		genHelper: true,
	},
	"TabpageArray": {
		name: "[]Tabpage",
		enc:  "encodeTabpageSlice",
		dec:  "decodeTabpageSlice",
	},
	"Client": {
		name: "Client",
	},
}
