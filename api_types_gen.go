package neovim

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Range) DecodeMsg(dc *msgp.Reader) (err error) {
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
	z.StartLine, err = dc.ReadInt()
	if err != nil {
		return
	}
	z.EndLine, err = dc.ReadInt()
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Range) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.StartLine)
	if err != nil {
		return
	}
	err = en.WriteInt(z.EndLine)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Range) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendInt(o, z.StartLine)
	o = msgp.AppendInt(o, z.EndLine)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Range) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
	z.StartLine, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		return
	}
	z.EndLine, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

func (z Range) Msgsize() (s int) {
	s = 1 + msgp.IntSize + msgp.IntSize
	return
}
