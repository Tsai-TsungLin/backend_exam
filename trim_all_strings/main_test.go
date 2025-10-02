package main

import (
	"testing"
)

// TestTrimSimpleString 測試簡單字串 trim
func TestTrimSimpleString(t *testing.T) {
	type Data struct {
		Value string
	}

	d := &Data{Value: "  hello  "}
	TrimAllStrings(&d)

	if d.Value != "hello" {
		t.Errorf("Value = %q, want %q", d.Value, "hello")
	}
}

// TestTrimNestedStruct 測試巢狀結構
func TestTrimNestedStruct(t *testing.T) {
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

	if a.Name != "name" {
		t.Errorf("a.Name = %q, want %q", a.Name, "name")
	}
	if a.Next.Name != "name2" {
		t.Errorf("a.Next.Name = %q, want %q", a.Next.Name, "name2")
	}
	if a.Next.Next.Name != "name3" {
		t.Errorf("a.Next.Next.Name = %q, want %q", a.Next.Next.Name, "name3")
	}
}

// TestTrimCircularReference 測試循環引用
func TestTrimCircularReference(t *testing.T) {
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
		},
	}

	// 創建循環引用
	a.Next.Next = a

	TrimAllStrings(&a)

	if a.Name != "name" {
		t.Errorf("a.Name = %q, want %q", a.Name, "name")
	}
	if a.Next.Name != "name2" {
		t.Errorf("a.Next.Name = %q, want %q", a.Next.Name, "name2")
	}

	// 驗證循環引用仍然存在
	if a.Next.Next != a {
		t.Errorf("circular reference broken")
	}

	// 驗證循環引用中的值也被 trim
	if a.Next.Next.Name != "name" {
		t.Errorf("a.Next.Next.Name = %q, want %q", a.Next.Next.Name, "name")
	}
}

// TestTrimSelfReference 測試自我引用
func TestTrimSelfReference(t *testing.T) {
	type Person struct {
		Name string
		Next *Person
	}

	a := &Person{
		Name: " self ",
	}
	a.Next = a

	TrimAllStrings(&a)

	if a.Name != "self" {
		t.Errorf("a.Name = %q, want %q", a.Name, "self")
	}
	if a.Next != a {
		t.Errorf("self reference broken")
	}
}

// TestTrimSliceOfStrings 測試字串 slice
func TestTrimSliceOfStrings(t *testing.T) {
	type Data struct {
		Items []string
	}

	d := &Data{
		Items: []string{" one ", " two ", " three "},
	}

	TrimAllStrings(&d)

	expected := []string{"one", "two", "three"}
	for i, item := range d.Items {
		if item != expected[i] {
			t.Errorf("Items[%d] = %q, want %q", i, item, expected[i])
		}
	}
}

// TestTrimSliceOfStructs 測試結構體 slice
func TestTrimSliceOfStructs(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	type Group struct {
		Members []Person
	}

	g := &Group{
		Members: []Person{
			{Name: " Alice ", Age: 30},
			{Name: " Bob ", Age: 25},
			{Name: " Charlie ", Age: 35},
		},
	}

	TrimAllStrings(&g)

	expected := []string{"Alice", "Bob", "Charlie"}
	for i, member := range g.Members {
		if member.Name != expected[i] {
			t.Errorf("Members[%d].Name = %q, want %q", i, member.Name, expected[i])
		}
	}
}

// TestTrimMap 測試 map
func TestTrimMap(t *testing.T) {
	type Data struct {
		Values map[string]string
	}

	d := &Data{
		Values: map[string]string{
			"key1": " value1 ",
			"key2": " value2 ",
			"key3": " value3 ",
		},
	}

	TrimAllStrings(&d)

	expected := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	for key, val := range d.Values {
		if val != expected[key] {
			t.Errorf("Values[%s] = %q, want %q", key, val, expected[key])
		}
	}
}

// TestTrimMapOfStructs 測試包含結構體的 map
func TestTrimMapOfStructs(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	type Data struct {
		People map[string]Person
	}

	d := &Data{
		People: map[string]Person{
			"p1": {Name: " Alice ", Age: 30},
			"p2": {Name: " Bob ", Age: 25},
		},
	}

	TrimAllStrings(&d)

	if d.People["p1"].Name != "Alice" {
		t.Errorf("People[p1].Name = %q, want %q", d.People["p1"].Name, "Alice")
	}
	if d.People["p2"].Name != "Bob" {
		t.Errorf("People[p2].Name = %q, want %q", d.People["p2"].Name, "Bob")
	}
}

// TestTrimArray 測試陣列
func TestTrimArray(t *testing.T) {
	type Data struct {
		Items [3]string
	}

	d := &Data{
		Items: [3]string{" one ", " two ", " three "},
	}

	TrimAllStrings(&d)

	expected := [3]string{"one", "two", "three"}
	if d.Items != expected {
		t.Errorf("Items = %v, want %v", d.Items, expected)
	}
}

// TestTrimNilPointer 測試 nil 指標
func TestTrimNilPointer(t *testing.T) {
	type Person struct {
		Name string
		Next *Person
	}

	a := &Person{
		Name: " name ",
		Next: nil,
	}

	TrimAllStrings(&a)

	if a.Name != "name" {
		t.Errorf("a.Name = %q, want %q", a.Name, "name")
	}
	if a.Next != nil {
		t.Errorf("a.Next should be nil")
	}
}

// TestTrimEmptyString 測試空字串
func TestTrimEmptyString(t *testing.T) {
	type Data struct {
		Value string
	}

	d := &Data{Value: "   "}
	TrimAllStrings(&d)

	if d.Value != "" {
		t.Errorf("Value = %q, want empty string", d.Value)
	}
}

// TestTrimMultipleSpaces 測試多個空白字元
func TestTrimMultipleSpaces(t *testing.T) {
	type Data struct {
		Value string
	}

	d := &Data{Value: "  \t\n  hello  \t\n  "}
	TrimAllStrings(&d)

	if d.Value != "hello" {
		t.Errorf("Value = %q, want %q", d.Value, "hello")
	}
}

// TestTrimComplexStructure 測試複雜結構
func TestTrimComplexStructure(t *testing.T) {
	type Address struct {
		Street string
		City   string
	}

	type Person struct {
		Name      string
		Addresses []Address
		Metadata  map[string]string
	}

	p := &Person{
		Name: " John Doe ",
		Addresses: []Address{
			{Street: " 123 Main St ", City: " New York "},
			{Street: " 456 Oak Ave ", City: " Los Angeles "},
		},
		Metadata: map[string]string{
			"title": " Engineer ",
			"dept":  " IT ",
		},
	}

	TrimAllStrings(&p)

	if p.Name != "John Doe" {
		t.Errorf("Name = %q, want %q", p.Name, "John Doe")
	}
	if p.Addresses[0].Street != "123 Main St" {
		t.Errorf("Addresses[0].Street = %q, want %q", p.Addresses[0].Street, "123 Main St")
	}
	if p.Addresses[0].City != "New York" {
		t.Errorf("Addresses[0].City = %q, want %q", p.Addresses[0].City, "New York")
	}
	if p.Addresses[1].Street != "456 Oak Ave" {
		t.Errorf("Addresses[1].Street = %q, want %q", p.Addresses[1].Street, "456 Oak Ave")
	}
	if p.Addresses[1].City != "Los Angeles" {
		t.Errorf("Addresses[1].City = %q, want %q", p.Addresses[1].City, "Los Angeles")
	}
	if p.Metadata["title"] != "Engineer" {
		t.Errorf("Metadata[title] = %q, want %q", p.Metadata["title"], "Engineer")
	}
	if p.Metadata["dept"] != "IT" {
		t.Errorf("Metadata[dept] = %q, want %q", p.Metadata["dept"], "IT")
	}
}

// TestTrimDeepCircularReference 測試深層循環引用
func TestTrimDeepCircularReference(t *testing.T) {
	type Person struct {
		Name string
		Next *Person
	}

	a := &Person{Name: " a "}
	b := &Person{Name: " b "}
	c := &Person{Name: " c "}

	a.Next = b
	b.Next = c
	c.Next = a // 循環回到 a

	TrimAllStrings(&a)

	if a.Name != "a" {
		t.Errorf("a.Name = %q, want %q", a.Name, "a")
	}
	if b.Name != "b" {
		t.Errorf("b.Name = %q, want %q", b.Name, "b")
	}
	if c.Name != "c" {
		t.Errorf("c.Name = %q, want %q", c.Name, "c")
	}

	// 驗證循環引用仍然存在
	if a.Next != b || b.Next != c || c.Next != a {
		t.Errorf("circular reference chain broken")
	}
}

// TestTrimNoChanges 測試不需要 trim 的字串
func TestTrimNoChanges(t *testing.T) {
	type Data struct {
		Value string
	}

	d := &Data{Value: "hello"}
	TrimAllStrings(&d)

	if d.Value != "hello" {
		t.Errorf("Value = %q, want %q", d.Value, "hello")
	}
}

// TestTrimSliceOfPointers 測試指標 slice
func TestTrimSliceOfPointers(t *testing.T) {
	type Person struct {
		Name string
	}

	type Group struct {
		Members []*Person
	}

	g := &Group{
		Members: []*Person{
			{Name: " Alice "},
			{Name: " Bob "},
		},
	}

	TrimAllStrings(&g)

	if g.Members[0].Name != "Alice" {
		t.Errorf("Members[0].Name = %q, want %q", g.Members[0].Name, "Alice")
	}
	if g.Members[1].Name != "Bob" {
		t.Errorf("Members[1].Name = %q, want %q", g.Members[1].Name, "Bob")
	}
}
