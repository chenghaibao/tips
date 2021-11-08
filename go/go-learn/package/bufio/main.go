package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

type Bufio struct {
	reader *bufio.Reader
	write  *bufio.Writer
}

/**https://studygolang.com/pkgdoc*/
func main() {

	// writer
	s := bytes.NewBuffer([]byte{})
	w := bufio.NewWriter(s)
	w.WriteString("hello world")
	w.WriteString("你好")
	fmt.Printf("string--%s", s.String())
	fmt.Println()
	w.Flush()
	fmt.Printf("string--%s\n", s.String())

	// Scanner
	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}


	input := "abcdef\nghijkl"
	scanner1 := bufio.NewScanner(strings.NewReader(input))

	// 底层原理
	//split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	//	fmt.Printf("%t\t%d\t%s\n", atEOF, len(data), data)
	//	return 0, nil, nil
	//}
	//buf := make([]byte, 55 贪心)
	//scanner1.Buffer(buf, bufio.MaxScanTokenSize)

	scanner1.Split(bufio.ScanLines)
	for scanner1.Scan() {
		fmt.Printf("%s\n", scanner1.Text())
	}

	input1 := "foo   bar      baz"
	scanner2 := bufio.NewScanner(strings.NewReader(input1))
	scanner2.Split(bufio.ScanWords)
	for scanner2.Scan() {
		fmt.Println(scanner2.Text())
	}


	// Reader
	r := strings.NewReader("hello world !")
	// 缓冲区
	read := bufio.NewReader(r)
	rea := &Bufio{reader: read}
	// 清楚缓冲区
	//rea.reader.Reset(rea.reader)

	// 缓冲区字节大小
	//fmt.Println(rea.reader.Buffered())

	//str,_ :=rea.reader.ReadByte()
	//fmt.Println(str)

	for {
		str, err := rea.reader.ReadString(byte(' '))
		fmt.Println(str)

		if err != nil {
			return
		}
	}




}
