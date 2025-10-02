package main

import (
	"testing"
)

// TestSwapInt 測試整數交換
func TestSwapInt(t *testing.T) {
	a := 10
	b := 20

	addrA := &a
	addrB := &b

	swap(&a, &b)

	if a != 20 {
		t.Errorf("a = %d, want 20", a)
	}
	if b != 10 {
		t.Errorf("b = %d, want 10", b)
	}

	// 驗證地址不變
	if &a != addrA {
		t.Errorf("address of a changed: got %p, want %p", &a, addrA)
	}
	if &b != addrB {
		t.Errorf("address of b changed: got %p, want %p", &b, addrB)
	}
}

// TestSwapString 測試字串交換
func TestSwapString(t *testing.T) {
	a := "hello"
	b := "world"

	addrA := &a
	addrB := &b

	swap(&a, &b)

	if a != "world" {
		t.Errorf("a = %s, want world", a)
	}
	if b != "hello" {
		t.Errorf("b = %s, want hello", b)
	}

	// 驗證地址不變
	if &a != addrA {
		t.Errorf("address of a changed")
	}
	if &b != addrB {
		t.Errorf("address of b changed")
	}
}

// TestSwapFloat64 測試浮點數交換
func TestSwapFloat64(t *testing.T) {
	a := 3.14
	b := 2.71

	addrA := &a
	addrB := &b

	swap(&a, &b)

	if a != 2.71 {
		t.Errorf("a = %f, want 2.71", a)
	}
	if b != 3.14 {
		t.Errorf("b = %f, want 3.14", b)
	}

	// 驗證地址不變
	if &a != addrA {
		t.Errorf("address of a changed")
	}
	if &b != addrB {
		t.Errorf("address of b changed")
	}
}

// TestSwapBool 測試布林值交換
func TestSwapBool(t *testing.T) {
	a := true
	b := false

	swap(&a, &b)

	if a != false {
		t.Errorf("a = %v, want false", a)
	}
	if b != true {
		t.Errorf("b = %v, want true", b)
	}
}

// TestSwapStruct 測試結構體交換
func TestSwapStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	a := Person{Name: "Alice", Age: 30}
	b := Person{Name: "Bob", Age: 25}

	addrA := &a
	addrB := &b

	swap(&a, &b)

	if a.Name != "Bob" || a.Age != 25 {
		t.Errorf("a = %+v, want {Name:Bob Age:25}", a)
	}
	if b.Name != "Alice" || b.Age != 30 {
		t.Errorf("b = %+v, want {Name:Alice Age:30}", b)
	}

	// 驗證地址不變
	if &a != addrA {
		t.Errorf("address of a changed")
	}
	if &b != addrB {
		t.Errorf("address of b changed")
	}
}

// TestSwapSlice 測試切片交換
func TestSwapSlice(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	addrA := &a
	addrB := &b

	swap(&a, &b)

	if len(a) != 3 || a[0] != 4 || a[1] != 5 || a[2] != 6 {
		t.Errorf("a = %v, want [4 5 6]", a)
	}
	if len(b) != 3 || b[0] != 1 || b[1] != 2 || b[2] != 3 {
		t.Errorf("b = %v, want [1 2 3]", b)
	}

	// 驗證地址不變
	if &a != addrA {
		t.Errorf("address of a changed")
	}
	if &b != addrB {
		t.Errorf("address of b changed")
	}
}

// TestSwapMap 測試 map 交換
func TestSwapMap(t *testing.T) {
	a := map[string]int{"one": 1, "two": 2}
	b := map[string]int{"three": 3}

	addrA := &a
	addrB := &b

	swap(&a, &b)

	if len(a) != 1 || a["three"] != 3 {
		t.Errorf("a = %v, want map[three:3]", a)
	}
	if len(b) != 2 || b["one"] != 1 || b["two"] != 2 {
		t.Errorf("b = %v, want map[one:1 two:2]", b)
	}

	// 驗證地址不變
	if &a != addrA {
		t.Errorf("address of a changed")
	}
	if &b != addrB {
		t.Errorf("address of b changed")
	}
}

// TestSwapPointer 測試指標交換
func TestSwapPointer(t *testing.T) {
	val1 := 100
	val2 := 200

	a := &val1
	b := &val2

	addrA := &a
	addrB := &b

	swap(&a, &b)

	if *a != 200 {
		t.Errorf("*a = %d, want 200", *a)
	}
	if *b != 100 {
		t.Errorf("*b = %d, want 100", *b)
	}
	if a != &val2 {
		t.Errorf("a should point to val2")
	}
	if b != &val1 {
		t.Errorf("b should point to val1")
	}

	// 驗證地址不變
	if &a != addrA {
		t.Errorf("address of a changed")
	}
	if &b != addrB {
		t.Errorf("address of b changed")
	}
}

// TestSwapArray 測試陣列交換
func TestSwapArray(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := [3]int{4, 5, 6}

	swap(&a, &b)

	if a != [3]int{4, 5, 6} {
		t.Errorf("a = %v, want [4 5 6]", a)
	}
	if b != [3]int{1, 2, 3} {
		t.Errorf("b = %v, want [1 2 3]", b)
	}
}

// TestSwapInterface 測試 interface 交換
func TestSwapInterface(t *testing.T) {
	var a interface{} = 42
	var b interface{} = "hello"

	swap(&a, &b)

	if a != "hello" {
		t.Errorf("a = %v, want hello", a)
	}
	if b != 42 {
		t.Errorf("b = %v, want 42", b)
	}
}

// TestSwapPanicNonPointer 測試非指標參數會 panic
func TestSwapPanicNonPointer(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("swap should panic with non-pointer argument")
		} else if r != "swap: first argument must be a pointer" {
			t.Errorf("unexpected panic message: %v", r)
		}
	}()

	a := 10
	b := 20
	swap(a, b) // 傳值而非指標
}

// TestSwapPanicNil 測試 nil 指標會 panic
func TestSwapPanicNil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("swap should panic with nil pointer")
		}
	}()

	var a *int = nil
	b := 20
	swap(a, &b)
}

// TestSwapPanicTypeMismatch 測試類型不匹配的情況
// 注意：由於使用泛型，類型不匹配會在編譯時被捕獲，而非運行時
// 此測試主要用於文檔目的，說明泛型提供了編譯時類型安全
func TestSwapPanicTypeMismatch(t *testing.T) {
	// 由於泛型的類型推斷，無法傳入不同類型的參數
	// 例如：swap(&a, &b) 其中 a 是 int, b 是 string
	// 這會導致編譯錯誤，而非運行時 panic
	// 因此此測試被跳過
	t.Skip("Type mismatch is caught at compile time due to generics")
}

// TestSwapSameValue 測試交換相同值
func TestSwapSameValue(t *testing.T) {
	a := 42
	b := 42

	swap(&a, &b)

	if a != 42 {
		t.Errorf("a = %d, want 42", a)
	}
	if b != 42 {
		t.Errorf("b = %d, want 42", b)
	}
}

// TestSwapZeroValues 測試零值交換
func TestSwapZeroValues(t *testing.T) {
	var a int
	var b int

	swap(&a, &b)

	if a != 0 {
		t.Errorf("a = %d, want 0", a)
	}
	if b != 0 {
		t.Errorf("b = %d, want 0", b)
	}
}
