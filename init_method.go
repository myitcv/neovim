package neovim

import "github.com/tinylib/msgp/msgp"

type InitMethodWrapper struct {
	*Client
	InitMethod
	args    *InitMethodArgs
	results *InitMethodRetVals
}

func (i *InitMethodWrapper) Args() msgp.Decodable {
	return i.args
}

func (i *InitMethodWrapper) Results() msgp.Encodable {
	return i.results
}

func (i *InitMethodWrapper) Params() *MethodOptionParams {
	return nil
}

func (i *InitMethodWrapper) Eval() msgp.Decodable {
	return nil
}

type InitMethodArgs struct {
	hostName string
}

type InitMethodRetVals struct {
	InitMethod
}

func (z *InitMethodArgs) DecodeMsg(dc *msgp.Reader) (err error) {
	i, err := dc.ReadString()
	if err != nil {
		return
	}
	z.hostName = string(i)

	return
}

func (z *InitMethodRetVals) EncodeMsg(en *msgp.Writer) (err error) {

	err = en.WriteArrayHeader(0)
	if err != nil {
		return err
	}
	return nil
}

func (i *InitMethodWrapper) Run() (error, error) {
	// TODO gross remove this
	i.Client.HostName = i.args.hostName
	return nil, i.InitMethod()
}
