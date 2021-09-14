package main

//利用反射，根据结构体自动生成insert语句
import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) {
	v := reflect.ValueOf(q)
	t := reflect.TypeOf(q)
	if t.Kind() != reflect.Struct{
		fmt.Println("unsupported type")
		return
	}
	var values []string
	var fields []string
	for i :=0; i < t.NumField(); i++{
		vF := v.Field(i)
		fields = append(fields, t.Field(i).Name)
		switch(vF.Kind()){
		case reflect.String:
			values = append(values,"'"+vF.String()+"'")
		case reflect.Int:
			values = append(values, strconv.Itoa(int(vF.Int())))
		}
	}
	fmt.Printf("insert into %s(%s) values(%s)\n", t.Name(),strings.Join(fields,","),strings.Join(values,","))
}


func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	i := 90
	createQuery(i)
}