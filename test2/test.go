package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

type Apple struct {
	Color  string
	Weight int
}

func (p Person) GetName() string {
	return p.Name
}
func (p Person) SetName(s string) {
	p.Name = s
}
func FillStruct(obj interface{}) {
	fmt.Println("t.value of field by name", reflect.ValueOf(obj).FieldByName("Color"))
	fmt.Printf("typeof fieldbyname")
	test, _ := reflect.TypeOf(obj).FieldByName("Color")
	fmt.Printf("field by name is %+v", test)
	t := reflect.TypeOf(obj)                            //反射出一个interface{}的类型
	fmt.Println("t.Name()", t.Name())                   //类型名
	fmt.Println("t.Kind().String()", t.Kind().String()) //Type类型表示的具体分类
	fmt.Println("t.PkgPath()", t.PkgPath())             //反射对象所在的短包名
	fmt.Println("t.String()", t.String())               //包名.类型名
	fmt.Println("t.Size()", t.Size())                   //要保存一个该类型要多少个字节
	fmt.Println("t.Align()", t.Align())                 //返回当从内存中申请一个该类型值时，会对齐的字节数
	fmt.Println("t.FieldAlign()", t.FieldAlign())       //返回当该类型作为结构体的字段时，会对齐的字节数
	fmt.Printf("t.fieldbyname color")
	fmt.Println(t.FieldByName("Color"))

	var u Person
	fmt.Println("t.AssignableTo(reflect.TypeOf(u))", t.AssignableTo(reflect.TypeOf(u)))   // 如果该类型的值可以直接赋值给u代表的类型，返回真
	fmt.Println("t.ConvertibleTo(reflect.TypeOf(u))", t.ConvertibleTo(reflect.TypeOf(u))) // 如该类型的值可以转换为u代表的类型，返回真
	t.FieldByName("name")
	fmt.Println("t.NumField()", t.NumField())                         // 返回struct类型的字段数（匿名字段算作一个字段），如非结构体类型将panic
	fmt.Println("t.Field(0).Name", t.Field(0).Name)                   // 返回struct类型的第i个字段的类型，如非结构体或者i不在[0, NumField())内将会panic
	fmt.Println(t.FieldByName("name"))                                // 返回该类型名为name的字段（会查找匿名字段及其子字段），布尔值说明是否找到，如非结构体将panic
	fmt.Println("t.FieldByIndex([]int{0})", t.FieldByIndex([]int{0})) // 返回索引序列指定的嵌套字段的类型，等价于用索引中每个值链式调用本方法，如非结构体将会panic
}
func main() {
	var a interface{} = nil
	a = Person{"gooduone", 20}
	a = Apple{"red", 20}
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a))
	FillStruct(a)
}
