package reflect

import(
	"fmt"
	"reflect"
	"testing"
)

// https://www.jianshu.com/p/26a284e69586
type order struct{
	ordId int
	customerId int
}

func query(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
}


func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Kind ", k)


}


func createQueryField(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}

}

func TestMaintainer(t *testing.T){
	o := order{
		ordId: 456,
		customerId: 56,
	}
	createQuery(o)
	query(o)
	createQueryField(o)

	//a := 56
	//x := reflect.ValueOf(a).Int()
	//fmt.Printf("type:%T value:%v\n", x, x)
	//b := "Naveen"
	//y := reflect.ValueOf(b).String()
	//fmt.Printf("type:%T value:%v\n", y, y)
}