package io

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestIo(t *testing.T) {

	if _, err := io.WriteString(os.Stdout, "Hello World\n"); err != nil {
		log.Println(err)
	}

	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 68)

	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}

	r1 := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.\n")
	b, err := io.ReadAll(r1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

}

