package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer // A Buffer needs no initialization.
	b.Write([]byte("Hello"))
	b.Write([]byte(" World!"))
	//fmt.Fprintf(&b, "world!")
	b.WriteTo(os.Stdout)
	fmt.Print(b.String())
	fmt.Print(b.String())
}
