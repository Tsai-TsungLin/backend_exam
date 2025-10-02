package main

import (
	"fmt"
	"reflect"
)

func swap[T any](a, b T) {
	aVal := reflect.ValueOf(a)
	bVal := reflect.ValueOf(b)

	// 檢查是否為指標
	if aVal.Kind() != reflect.Ptr {
		panic("swap: first argument must be a pointer")
	}
	if bVal.Kind() != reflect.Ptr {
		panic("swap: second argument must be a pointer")
	}

	// 檢查是否為 nil
	if aVal.IsNil() {
		panic("swap: first argument is nil")
	}
	if bVal.IsNil() {
		panic("swap: second argument is nil")
	}

	// 取得指標指向的值
	aElem := aVal.Elem()
	bElem := bVal.Elem()

	// 檢查是否可以設定值
	if !aElem.CanSet() {
		panic("swap: first argument cannot be set")
	}
	if !bElem.CanSet() {
		panic("swap: second argument cannot be set")
	}

	// 檢查類型是否一致
	if aElem.Type() != bElem.Type() {
		panic(fmt.Sprintf("swap: type mismatch: %v != %v", aElem.Type(), bElem.Type()))
	}

	// 交換值
	tmp := reflect.New(aElem.Type()).Elem()
	tmp.Set(aElem)
	aElem.Set(bElem)
	bElem.Set(tmp)
}

func main() {
	a := 10
	b := 20

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)

	swap(&a, &b)

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)
}
