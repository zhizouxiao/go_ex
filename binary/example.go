// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
)

type abc struct {
	a int32
	b int
	c int
}

func sizeof(t reflect.Type) int {
	fmt.Println("=========size", t.Kind())
	switch t.Kind() {
	case reflect.Array:
		if s := sizeof(t.Elem()); s >= 0 {
			return s * t.Len()
		}

	case reflect.Struct:
		sum := 0
		for i, n := 0, t.NumField(); i < n; i++ {
			s := sizeof(t.Field(i).Type)
			fmt.Println("=========", t.Field(i).Type, t)
			if s < 0 {
				return -1
			}
			sum += s
		}
		return sum

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return int(t.Size())

	}

	return -1
}

func dataSize(v reflect.Value) int {
	if v.Kind() == reflect.Slice {
		fmt.Println(v.Len())

		if s := sizeof(v.Type().Elem()); s >= 0 {
			return s * v.Len()
		}
		return -1
	}
	//fmt.Println(v.Len())
	return sizeof(v.Type())
}

func printLen(data interface{}) int {
	v := reflect.Indirect(reflect.ValueOf(data))
	size := dataSize(v)

	return size

}

func main() {
	_abc := new(abc)
	_abc.a = 1
	_abc.b = 1
	_abc.c = 1

	fmt.Println("abc len=", printLen(_abc))

	buf := new(bytes.Buffer)

	var pi float64 = math.Pi
	err := binary.Write(buf, binary.LittleEndian, pi)
	fmt.Println("Pi len=", printLen(pi))

	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.Bytes())
	// Output: 18 2d 44 54 fb 21 09 40
}

func ExampleWrite_multi() {
	buf := new(bytes.Buffer)
	var data = []interface{}{
		uint16(61374),
		int8(-54),
		uint8(254),
	}
	for _, v := range data {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	fmt.Printf("%x", buf.Bytes())
	// Output: beefcafe
}

func ExampleRead() {
	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Print(pi)
	// Output: 3.141592653589793
}
