package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func TrimAllStrings(a any) {
	visited := make(map[uintptr]bool)
	trimAllStringsRecursive(reflect.ValueOf(a), visited)
}

func trimAllStringsRecursive(v reflect.Value, visited map[uintptr]bool) {
	if !v.IsValid() {
		return
	}

	// 處理 interface，取得實際值
	if v.Kind() == reflect.Interface && !v.IsNil() {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Ptr:
		// 檢查 nil
		if v.IsNil() {
			return
		}

		// 檢查循環引用
		ptr := v.Pointer()
		if visited[ptr] {
			return
		}
		visited[ptr] = true

		// 遞迴處理指標指向的值
		trimAllStringsRecursive(v.Elem(), visited)

	case reflect.Struct:
		// 遍歷所有欄位
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if field.CanSet() {
				trimAllStringsRecursive(field, visited)
			}
		}

	case reflect.String:
		// Trim 字串
		if v.CanSet() {
			trimmed := strings.TrimSpace(v.String())
			v.SetString(trimmed)
		}

	case reflect.Slice, reflect.Array:
		// 處理每個元素
		for i := 0; i < v.Len(); i++ {
			trimAllStringsRecursive(v.Index(i), visited)
		}

	case reflect.Map:
		// 處理 map 的每個 value
		if v.IsNil() {
			return
		}
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			// Map value 不能直接 Set，需要創建新值
			if val.Kind() == reflect.String {
				trimmed := strings.TrimSpace(val.String())
				v.SetMapIndex(key, reflect.ValueOf(trimmed))
			} else if val.CanAddr() {
				trimAllStringsRecursive(val, visited)
			} else {
				// 對於不可尋址的值，需要複製
				newVal := reflect.New(val.Type()).Elem()
				newVal.Set(val)
				trimAllStringsRecursive(newVal, visited)
				v.SetMapIndex(key, newVal)
			}
		}
	}
}

func main() {
	type Person struct {
		Name string
		Age  int
		Next *Person
	}

	a := &Person{
		Name: " name ",
		Age:  20,
		Next: &Person{
			Name: " name2 ",
			Age:  21,
			Next: &Person{
				Name: " name3 ",
				Age:  22,
			},
		},
	}

	TrimAllStrings(&a)

	m, _ := json.Marshal(a)

	fmt.Println(string(m))

	a.Next = a

	TrimAllStrings(&a)

	fmt.Println(a.Next.Next.Name == "name")
}
