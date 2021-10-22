package byte

import (
	"bytes"
	"fmt"
	"testing"
)

type b struct {
	byt *bytes.Buffer
	read *bytes.Reader
}
func TestByte(t *testing.T) {
	stringByt :=  bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	reader :=	bytes.NewReader([]byte("asdsa"))
	b :=&b{
		byt: stringByt,
		read: reader,
	}
	b.byt.Write([]byte("haibao"))

	fmt.Println(b.byt.Len())
	fmt.Println(b.read.Len())
}