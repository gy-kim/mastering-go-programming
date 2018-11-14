package main

import "fmt"

type magicStore struct {
	value interface{}
	name  string
}

func (ms *magicStore) SetValue(v interface{}) {
	ms.value = v
}

func (ms *magicStore) GetValue() interface{} {
	return ms.value
}

func NewMagicStore(nm string) *magicStore {
	return &magicStore{name: nm}
}

func main() {
	istore := NewMagicStore("Integer Store")
	istore.SetValue(4)
	if v, ok := istore.GetValue().(int); ok {
		v *= 4
		fmt.Println(v)
	}

	sstore := NewMagicStore("String store")
	sstore.SetValue("Hello")
	if v, ok := sstore.GetValue().(string); ok {
		v += " World"
		fmt.Println(v)
	}
}
