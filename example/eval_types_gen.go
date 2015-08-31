package example

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *MyEvalResult) DecodeMsg(dc *msgp.Reader) (err error) {
	var ssz uint32
	ssz, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if ssz != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: ssz}
		return
	}
	{
		z.S, err = dc.ReadBytes(z.S)
		if err != nil {
			return
		}
		z.I, err = dc.ReadInt()
		if err != nil {
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MyEvalResult) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.S)
	if err != nil {
		return
	}
	err = en.WriteInt(z.I)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MyEvalResult) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendBytes(o, z.S)
	o = msgp.AppendInt(o, z.I)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MyEvalResult) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
	z.S, bts, err = msgp.ReadBytesBytes(bts, z.S)
	if err != nil {
		return
	}
	z.I, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

func (z *MyEvalResult) Msgsize() (s int) {
	s = 1 + msgp.BytesPrefixSize + len(z.S) + msgp.IntSize
	return
}
