package main

import (
	"fmt"
	"reflect"
)

type OldInt int
type OldInt32 int32

type Int interface {
	~int | ~int32 | ~int64
}

// type Generic[T int | int32] T // error

// ptr、slice、struct、map、channel、interface
type Ptr[T int | int32] *T // ptr

type SliceInt[T Int] []T // slice

type StructInt[T Int] struct { // struct
	Data T
}

func (s *StructInt[T]) Val() T {
	return s.Data
}

type MapInt[K Int, V any] map[K]V // map

type ChannelInt[T Int] chan T // chan

// 只要满足有方法 Val() T，T的约束为 ~int | ~int32 | ~int64的都为这个接口的实现
type InterfaceInt[T Int] interface { // interface
	Val() T
}

type InterfaceIntImpl1 struct{}

func (i *InterfaceIntImpl1) Val() int {
	return 1
}

type InterfaceIntImpl2 int32

func (i InterfaceIntImpl2) Val() int32 {
	return int32(i)
}

func main() {
	var interfaceInt InterfaceInt[int]

	interfaceInt = &InterfaceIntImpl1{}
	val := interfaceInt.Val()
	fmt.Println(val)
	fmt.Println(reflect.TypeOf(val)) // int

	interfaceInt = &StructInt[int]{Data: 2}
	val2 := interfaceInt.Val()
	fmt.Println(val2)
	fmt.Println(reflect.TypeOf(val2)) // int

	var interfaceInt32 InterfaceInt[int32] = InterfaceIntImpl2(int32(2))
	valInt32 := interfaceInt32.Val()
	fmt.Println(valInt32)
	fmt.Println(reflect.TypeOf(valInt32)) // int32
}
