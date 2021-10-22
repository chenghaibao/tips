package sort

import (
	"fmt"
	"sort"
	"testing"
)

type Person struct {
	Name string
	Age  int
}
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func TestSort(t *testing.T) {

	// 常用sort
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))

	fmt.Println(people)
	a := []int{55, 45, 36, 28, 21, 15, 10, 6, 3, 1}
	x := 6
	sort.Ints(a)
	fmt.Println(a)

	i := sort.Search(len(a), func(i int) bool { return a[i] <= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}

	fmt.Println(i)
}
