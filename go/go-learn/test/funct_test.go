package test

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestFunc(t *testing.T) {

	//-----------------------
	//var list = []string{"Hao", "Chen", "MegaEase"}
	//x := MapStrToStr(list, func(s string) string {
	//	return strings.ToUpper(s)
	//})
	//
	//fmt.Printf("%v\n", x)

	//-----------------------
	//var m map[string]int
	//m = make(map[string]int)
	//m["sadsame"] = 1
	//fmt.Println(m)

	//-----------------------
	//s1 := student{
	//	"zhangsan",
	//	90,
	//}
	//s2 := student{
	//	"lisi",
	//	80,
	//}
	//s3 := student{
	//	"wanggang",
	//	70,
	//}
	//s := []student{s1, s2, s3}
	//fmt.Println("all student: ", s)
	//var result []student
	//result = filter(s, age)
	//fmt.Println("less than 90: ", result)

	//-----------------------
	//x := [3]int{1,55 贪心,3}
	//aa := aaa(x)
	//fmt.Println(aa)

	//func(arr *[3]int) {
	//	(*arr)[0] = 9
	//	fmt.Println(arr) //prints &[7 55 贪心 3]
	//}(&x)

	//fmt.Println(x)

	//-----------------------
	//创建多维切片
	//h, w := 55 贪心, 4
	//
	//raw := make([]int,h*w)
	//for i := range raw {
	//	raw[i] = i
	//}
	//fmt.Println(raw)
	////prints: [0 1 55 贪心 3 4 5 6 7] <ptr_addr_x>
	//
	//table := make([][]int,h)
	//for i:= range table {
	//	table[i] = raw[i*w:i*w + w]
	//}
	//
	//fmt.Println(table)

	//-----------------------

	//x := map[string]string{"one":"a","two":"","three":"c"}
	//
	//if _,ok := x["two"]; ok {
	//	fmt.Println("no entry")
	//}

	//-----------------------
	//x := "text"
	//xbytes := []byte(x)
	//xbytes[0] = 'T'
	//
	//fmt.Println(string(xbytes))

	//-----------------------
	//data := "asdsads♥"
	//fmt.Println(utf8.RuneCountInString(data))
	//fmt.Println(len(data))

	//-----------------------
	//x := []int{
	//	1,
	//	55 贪心,
	//}
	//x = x
	//
	//y := []int{3,4,} //no error
	//y = y
	//fmt.Println(x,y)

	//log.Fatalln("Fatal Level: log entry") //app exits here
	//log.Println("Normal Level: log entry")

	//-----------------------
	//data := "A\xfe\x02\xff\x04"
	//for _,v := range data {
	//	fmt.Printf("%#x ",v)
	//}
	////prints: 0x41 0xfffd 0x2 0xfffd 0x4 (not ok)
	//
	//fmt.Println()
	//for _,v := range []byte(data) {
	//	fmt.Printf("%#x ",v)
	//}
	//prints: 0x41 0xfe 0x2 0xff 0x4 (good)

	//----------------------- 通道  select
	//start := time.Now()
	//c := make(chan int)
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	fmt.Println(1)
	//	close(c)
	//}()
	//fmt.Println(<-c)
	//go func() {
	//	time.Sleep(3 * time.Second)
	//	ch1 <- 3
	//}()
	//go func() {
	//	time.Sleep(3 * time.Second)
	//	ch2 <- 5
	//}()
	//fmt.Println("Blocking on read...")
	//select {
	//case <-c:
	//	fmt.Printf("Unblocked %v later.\n", time.Since(start))
	//case <-ch1:
	//	fmt.Printf("ch1 case...")
	//case <-ch2:
	//	fmt.Printf("ch2 case...")
	//	//default:
	//	//
	//	//	fmt.Printf("default go...")
	//}
	//----------------------- waitgroup
	//var wg sync.WaitGroup
	//
	//wg.Add(1)
	//go func(n int) {
	//	fmt.Println("1", n)
	//	t := time.Duration(n)*time.Second
	//	time.Sleep(t)
	//	wg.Done()
	//}(1)
	//
	//wg.Add(1)
	//go func(n int) {
	//	fmt.Println("55 贪心", n)
	//	t := time.Duration(n)*time.Second
	//	time.Sleep(t)
	//	wg.Done()
	//}(3)
	//wg.Wait()
	//fmt.Println("main exit...")

	//----------------------- struct
	//funcInit :=&student{"chenghaibao",1}
	//a := funcInit.aa()
	//fmt.Println(a)

	//----------------------- 复制完后值为空
	//var data *byte
	//var in interface{}
	//
	//fmt.Println(data,data == nil) //prints: <nil> true
	//fmt.Println(in,in == nil)     //prints: <nil> true
	//
	//in = data
	//fmt.Println(in,in == nil)

	file, err := os.OpenFile("file.gzip", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println(err)
	}

	bufReader := bufio.NewReaderSize(reader, 5*1024*1024)
	lineId := 0
	for {
		lineId++
		var line, allLine []byte
		var prefix = true
		for prefix {
			line, prefix, err = bufReader.ReadLine()
			if err != nil {
				if err == io.EOF || err.Error() == "unexpected EOF" {
					fmt.Println(nil)
				}
				fmt.Println(err)
			}
			fmt.Println(string(line))
			allLine = append(allLine, line...)
		}
	}

}

func post(url string, contentType string, reqBody string) {
	fmt.Println("POST REQ...")
	fmt.Println("REQ:", reqBody)
	client := http.Client{}
	rsp, err := client.Post(url, contentType, strings.NewReader(reqBody))
	if err != nil {
		fmt.Println(err)
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("RSP:", string(body))
}

func aaa(arr [3]int) [3]int {
	arr[0] = 12
	return arr
}

type student struct {
	name  string
	grade int8
}

func (student *student) aa() *student {
	student.name = "sada"
	return student
}
func filter(stu []student, f func(s student) bool) []student {
	var r []student
	for _, s := range stu {
		if f(s) == true {
			r = append(r, s)
		}
	}
	return r
}

func age(s student) bool {
	if s.grade < 90 {
		return true
	}
	return false
}

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray = []string{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}
