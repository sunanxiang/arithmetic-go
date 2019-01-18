package gosnippet

import (
	"fmt"
	"reflect"
)

// introduce how to judge a interface is nil
type TestInterface interface {
	TestFunc ()
}

type TestObject struct {
	Name string
}

func (t *TestObject) TestFunc() {
	fmt.Println(t.Name)
}

func TestMain()  {
	// 1、
	var a TestInterface
	fmt.Println(a == nil) // true
	//fmt.Println(reflect.ValueOf(a).IsNil()) // panic

	// 2、
	var aa TestInterface = nil
	fmt.Println(aa == nil) // true
	//fmt.Println(reflect.ValueOf(aa).IsNil()) // panic

	// 3、
	var b *TestObject
	fmt.Println(b == nil) // true
	fmt.Println(reflect.ValueOf(b).IsNil()) // true

	var bb TestInterface = new(TestObject)
	fmt.Println(bb == nil) // false
	fmt.Println(reflect.ValueOf(bb).IsNil()) // false

	var bbb TestInterface = &TestObject{}
	fmt.Println(bbb == nil) // false
	fmt.Println(reflect.ValueOf(bbb).IsNil()) // false

	var bbbb = new(TestObject)
	fmt.Println(bbbb == nil) // false
	fmt.Println(reflect.ValueOf(bbbb).IsNil()) // false

	var bbbbb = &TestObject{}
	fmt.Println(bbbbb == nil) // false
	fmt.Println(reflect.ValueOf(bbbbb).IsNil()) // false

	// 4、
	var c *TestObject = nil
	var cc TestInterface = c
	fmt.Println(cc == nil) // false
	fmt.Println(reflect.ValueOf(cc).IsNil()) //true

	// 5、
	var d interface{}
	fmt.Println(d == nil) // true
	// fmt.Println(reflect.ValueOf(d).IsNil()) // panic

	//conclusion:
	// 1、if a interface which including methods is direct nil, interface == nil is true.
	// 2、if a object implemented interfaces but not explicitly invoked at definition time,
	// according to object's way to judge
	// 3、if a object implemented interfaces and explicitly invoked at definition time,
	// is not nil, must use reflect to judge
}
