// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	_log "log"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/template"

	"bitbucket.org/pkg/inflect"

	"github.com/juju/errors"
	"github.com/myitcv/neovim/apidef"
	"github.com/tinylib/msgp/msgp"
)

var generatedFunctions map[string]bool

var log = _log.New(os.Stdout, "", _log.Lshortfile)

var fCustomPrint = flag.Bool("c", false, "custom print")
var fPrintAPI = flag.Bool("p", false, "print the API")
var fGenAPI = flag.Bool("g", false, "generate code from the API")
var fAPIOutFile = flag.String("o", "", "the output file for the generated API")
var fTestOutFile = flag.String("t", "", "the output file for the generated test interface")

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

	output, err := exec.Command(os.Getenv("NEOVIM_BIN"), "--api-info").CombinedOutput()
	if err != nil {
		log.Fatalf("Could not get current API dump: %v", errors.Details(err))
	}

	br := bytes.NewReader(output)
	ad := msgp.NewReader(br)

	api, err := apidef.GetAPI(ad)
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
	Methods  []methodTemplate
	Types    []_type
	APITypes []apidef.APIClass
}

func getType(s string) _type {
	res, ok := typeMap[s]
	if !ok {
		log.Fatalf("Could not find type for %v", s)
	}
	return res
}

func genMethodTemplates(fs []apidef.APIFunction) []methodTemplate {
	res := make([]methodTemplate, len(fs))

	for i, f := range fs {
		m := methodTemplate{}
		splits := strings.SplitN(f.Name, "_", 2)

		if len(splits) != 2 {
			log.Fatalf("Function name not as expected: %v\n", f.Name)
		}

		// name
		m.RawName = f.Name
		m.Name = inflect.Camelize(splits[1])

		m.ID = f.ID

		// receiver
		recID := splits[0]
		var recType _type
		switch recID {
		case "vim":
			recType = getType("Client")
		case "buffer", "window", "tabpage":
			recType = getType(inflect.Camelize(recID))
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
		var ofInterest []apidef.APIFunctionParameter
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
			m.Params[i].name = inflect.Camelize(v.Name)
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

func genTypeTemplates(ts []apidef.APIClass) []_type {
	var res []_type
	for _, v := range typeMap {
		if v.genHelper {
			res = append(res, v)
		}
	}
	sort.Sort(byType(res))
	return res
}

func genAPI(a *apidef.API) {
	var funcsOfInterest []apidef.APIFunction
	if generatedFunctions != nil {
		funcsOfInterest = make([]apidef.APIFunction, 0)
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
	fm["camelize"] = inflect.Camelize
	fm["to_lower"] = strings.ToLower

	// generate the API
	t := template.New("api")
	t.Funcs(fm)
	_, err := t.Parse(clientAPITemplate)
	if err != nil {
		log.Fatalf("Could not parse client API template: %v\n", err)
	}

	api := api{}
	api.Methods = genMethodTemplates(funcsOfInterest)
	api.Types = genTypeTemplates(a.Types)
	api.APITypes = a.Types

	var apiOutFile io.Writer
	if *fAPIOutFile != "" {
		apiOutFile, err = os.OpenFile(*fAPIOutFile, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalf("Could not create API out file: %v\n", err)
		}
	} else {
		apiOutFile = os.Stdout
	}
	err = t.Execute(apiOutFile, api)
	if err != nil {
		fmt.Println()
		log.Fatalf("Error generating API: %v\n", err)
	}

	// generate the test interface
	tt := template.New("test interface")
	_, err = tt.Parse(clientAPITestInferface)
	if err != nil {
		log.Fatalf("Could not parse client API test interface template: %v\n", err)
	}
	var testOutFile io.Writer
	if *fTestOutFile != "" {
		testOutFile, err = os.OpenFile(*fTestOutFile, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalf("Could not create API out file: %v\n", err)
		}
	} else {
		testOutFile = os.Stdout
	}
	err = tt.Execute(testOutFile, api)
	if err != nil {
		fmt.Println()
		log.Fatalf("Error generating API: %v\n", err)
	}
}

var clientAPITestInferface = `
package neovim_test

import "gopkg.in/check.v1"

type neovimTester interface {
	{{range .Methods }}Test{{.Rec.Type.Name}}{{.Name}}(*check.C)
	Benchmark{{.Rec.Type.Name}}{{.Name}}(*check.C)
	{{end}}
}
`

var clientAPITemplate = `
package neovim

// **** THIS FILE IS GENERATED - DO NOT EDIT BY HAND

import "github.com/juju/errors"

// constants representing method ids

const (
	{{range .APITypes }}type{{.Name}} int8 = {{.Id}}
	{{end}}
)

const (
	{{range .Methods }}{{.Rec.Type.Name | to_lower}}{{.Name}} = "{{.RawName}}"
	{{end}}
)

// methods on the API

{{range .Methods }}
{{template "meth" .}}
{{end}}

{{range .APITypes }}
{{template "type_enc_dec_meth" .}}
{{end}}

// helper functions for types

{{range .Types}}
{{template "type" .}}
{{end}}

{{define "type"}}
func (c *Client) encode{{.Name | camelize }}Slice(s []{{.Name}}) error {
	err := c.enc.WriteArrayHeader(uint32(len(s)))
	if err != nil {
		return errors.Annotate(err, "Could not encode slice length")
	}

	for i := 0; i < len(s); i++ {
		{{if .Primitive}}
		err := c.enc.{{.Enc}}(s[i])
		{{else}}
		err := c.{{.Enc}}(s[i])
		{{end}}
		if err != nil {
			return errors.Annotatef(err, "Could not encode {{.Name}} at index %v", i)
		}
	}

	return  nil
}

func (c *Client) decode{{.Name | camelize }}Slice() ([]{{.Name}}, error) {
	l, err := c.dec.ReadArrayHeader()
	if err != nil {
		return nil, errors.Annotate(err, "Could not decode slice length")
	}

	res := make([]{{.Name}}, l)

	var i uint32
	for i = 0; i < l; i++ {
		{{if .Primitive}}
		b, err := c.dec.{{.Dec}}()
		{{else}}
		b, err := c.{{.Dec}}()
		{{end}}
		if err != nil {
			return nil, errors.Annotatef(err, "Could not decode {{.Name}} at index %v", i)
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
		_err = {{.Rec.Client}}.enc.WriteArrayHeader({{.NumParams}})
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
		_err = {{.Rec.Client}}.dec.ReadNil()
		{{end}}
		return
	}
	respChan, err := {{.Rec.Client}}.makeCall({{.Rec.Type.Name | to_lower}}{{.Name}}, enc, dec)
	if err != nil {
		return {{if .Ret}}{{.Ret.Name}}, {{end}}{{.Rec.Client}}.panicOrReturn(errors.Annotate(err, "Could not make call to {{.Rec.Type.Name}}.{{.Name}}"))
	}
	resp := <-respChan
	if resp == nil {
		return {{if .Ret}}{{.Ret.Name}}, {{end}}{{.Rec.Client}}.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return {{if .Ret}}{{.Ret.Name}}, {{end}}{{.Rec.Client}}.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
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
{{define "type_enc_dec_meth"}}
func (c *Client) decode{{.Name}}() (retVal {{.Name}}, retErr error) {
	retVal.client = c
	err := c.dec.ReadExtension(&retVal)
	if err != nil {
		return retVal, errors.Annotatef(err, "Could not decode extension {{.Name}}")
	}

	return
}

func (c *Client) encode{{.Name}}(b {{.Name}}) error {
	err := c.enc.WriteExtension(&b)
	if err != nil {
		return errors.Annotatef(err, "Could not encode {{.Name}}")
	}
	return nil
}
{{end}}
`

var typeMap = map[string]_type{
	"String": {
		name:      "string",
		enc:       "WriteString",
		dec:       "ReadString",
		primitive: true,
		genHelper: true,
	},
	"ArrayOf(String)": {
		name: "[]string",
		enc:  "encodeStringSlice",
		dec:  "decodeStringSlice",
	},
	"ArrayOf(Integer, 2)": {
		name: "[]int",
		enc:  "encodeIntSlice",
		dec:  "decodeIntSlice",
	},
	"Integer": {
		name:      "int",
		enc:       "WriteInt",
		dec:       "ReadInt",
		primitive: true,
		genHelper: true,
	},
	"Boolean": {
		name:      "bool",
		enc:       "WriteBool",
		dec:       "ReadBool",
		primitive: true,
	},
	"Array": {
		name:      "[]interface{}",
		enc:       "WriteIntf",
		dec:       "ReadIntf",
		primitive: true,
	},
	"Dictionary": {
		name: "map[string]interface{}",
		enc:  "encodeDictionary",
		dec:  "decodeDictionary",
	},
	"Object": {
		name:      "interface{}",
		enc:       "WriteIntf",
		dec:       "ReadIntf",
		primitive: true,
	},
	"Buffer": {
		name:      "Buffer",
		enc:       "encodeBuffer",
		dec:       "decodeBuffer",
		genHelper: true,
	},
	"ArrayOf(Buffer)": {
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
	"ArrayOf(Window)": {
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
	"ArrayOf(Tabpage)": {
		name: "[]Tabpage",
		enc:  "encodeTabpageSlice",
		dec:  "decodeTabpageSlice",
	},
	"Client": {
		name: "Client",
	},
}
