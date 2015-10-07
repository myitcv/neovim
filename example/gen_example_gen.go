package example

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *DoSomethingAsyncArgs) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var ssz uint32
		ssz, err = dc.ReadArrayHeader()
		if err != nil {
			return
		}
		if ssz != 1 {
			err = msgp.ArrayError{Wanted: 1, Got: ssz}
			return
		}
	}
	z.Arg0, err = dc.ReadString()
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z DoSomethingAsyncArgs) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 1
	err = en.Append(0x91)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Arg0)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z DoSomethingAsyncArgs) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 1
	o = append(o, 0x91)
	o = msgp.AppendString(o, z.Arg0)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DoSomethingAsyncArgs) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var ssz uint32
		ssz, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			return
		}
		if ssz != 1 {
			err = msgp.ArrayError{Wanted: 1, Got: ssz}
			return
		}
	}
	z.Arg0, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

func (z DoSomethingAsyncArgs) Msgsize() (s int) {
	s = 1 + msgp.StringPrefixSize + len(z.Arg0)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GetTwoNumbersArgs) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var ssz uint32
		ssz, err = dc.ReadArrayHeader()
		if err != nil {
			return
		}
		if ssz != 1 {
			err = msgp.ArrayError{Wanted: 1, Got: ssz}
			return
		}
	}
	z.Arg0, err = dc.ReadInt64()
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z GetTwoNumbersArgs) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 1
	err = en.Append(0x91)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Arg0)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z GetTwoNumbersArgs) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 1
	o = append(o, 0x91)
	o = msgp.AppendInt64(o, z.Arg0)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GetTwoNumbersArgs) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var ssz uint32
		ssz, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			return
		}
		if ssz != 1 {
			err = msgp.ArrayError{Wanted: 1, Got: ssz}
			return
		}
	}
	z.Arg0, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

func (z GetTwoNumbersArgs) Msgsize() (s int) {
	s = 1 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GetTwoNumbersResults) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var ssz uint32
		ssz, err = dc.ReadArrayHeader()
		if err != nil {
			return
		}
		if ssz != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: ssz}
			return
		}
	}
	z.Ret0, err = dc.ReadInt64()
	if err != nil {
		return
	}
	z.Ret1, err = dc.ReadString()
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z GetTwoNumbersResults) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Ret0)
	if err != nil {
		return
	}
	err = en.WriteString(z.Ret1)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z GetTwoNumbersResults) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendInt64(o, z.Ret0)
	o = msgp.AppendString(o, z.Ret1)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GetTwoNumbersResults) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var ssz uint32
		ssz, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			return
		}
		if ssz != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: ssz}
			return
		}
	}
	z.Ret0, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		return
	}
	z.Ret1, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

func (z GetTwoNumbersResults) Msgsize() (s int) {
	s = 1 + msgp.Int64Size + msgp.StringPrefixSize + len(z.Ret1)
	return
}
