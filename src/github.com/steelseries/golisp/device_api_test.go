// Copyright 2013 SteelSeries ApS. All rights reserved.
// No license is given for the use of this source code.

// This package impliments modular device support
// This file tests device description and serialization
package golisp

import (
    //"fmt"
    . "launchpad.net/gocheck"
)

type DeviceApiSuite struct {
}

var _ = Suite(&DeviceApiSuite{})

func (s *DeviceApiSuite) SetUpSuite(c *C) {
    Global = NewSymbolTableFrameBelow(nil)
    InitBuiltins()
    InitDeviceBuiltins()
}

func (s *DeviceApiSuite) TestByteArrayWrapping(c *C) {
    data := []byte{1, 2, 3, 4}
    wrapped := WrapByteArray(42, &data)
    c.Assert(len(*wrapped), Equals, 12)

    // type tag
    c.Assert((*wrapped)[0], Equals, byte(42))
    c.Assert((*wrapped)[1], Equals, byte(0))
    c.Assert((*wrapped)[2], Equals, byte(0))
    c.Assert((*wrapped)[3], Equals, byte(0))

    // size
    c.Assert((*wrapped)[4], Equals, byte(4))
    c.Assert((*wrapped)[5], Equals, byte(0))
    c.Assert((*wrapped)[6], Equals, byte(0))
    c.Assert((*wrapped)[7], Equals, byte(0))

    // data
    c.Assert((*wrapped)[8], Equals, byte(1))
    c.Assert((*wrapped)[9], Equals, byte(2))
    c.Assert((*wrapped)[10], Equals, byte(3))
    c.Assert((*wrapped)[11], Equals, byte(4))
}

func (s *DeviceApiSuite) TestListToByteArray(c *C) {
    var a []*Data
    d := []uint32{uint32(1), uint32(2), uint32(3), uint32(4), uint32(5)}
    for _, d := range d {
        a = append(a, NumberWithValue(d))
    }
    l := ArrayToList(a)
    b := ListToByteArray(l, Global)
    c.Assert((*b)[0], Equals, byte(1))
    c.Assert((*b)[1], Equals, byte(2))
    c.Assert((*b)[2], Equals, byte(3))
    c.Assert((*b)[3], Equals, byte(4))
    c.Assert((*b)[4], Equals, byte(5))
}

func (s *DeviceApiSuite) TestSimpleDefChunk(c *C) {
    code := "(def-chunk 10 8 (list-to-bytearray '(1 2 3 4 5 6 7 8)))"
    sexpr, err := Parse(code)
    chunkObj, err := Eval(sexpr, Global)
    c.Assert(err, IsNil)
    c.Assert(TypeOfObject(chunkObj), Equals, "ApiChunk")
    chunk := (*ApiChunk)(ObjectValue(chunkObj))
    c.Assert(chunk.DataType, Equals, uint32(10))
    c.Assert(chunk.DataSize, Equals, uint32(8))
}
