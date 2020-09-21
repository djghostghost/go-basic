package linkedhashmap

import (
	"fmt"
	"testing"
)

func TestLinkedHashMap_Put(t *testing.T) {
	m := New()
	if !m.IsEmpty() {
		t.Error()
	}

	m.Put("1", 1)
	m.Put("2", 2)
	m.Put("3", 3)
	m.Put("4", 4)

	if m.Size() != 4 {
		t.Error()
	}

	keys := m.Keys()

	if keys[0] != "1" || keys[1] != "2" || keys[2] != "3" || keys[3] != "4" {
		t.Error()
	}

	fmt.Println(keys)

	m.Remove("4")
	keys = m.Keys()

	if keys[0] != "1" || keys[1] != "2" || keys[2] != "3"{
		t.Error()
	}
	fmt.Println(keys)

	m.Put("4", 4)

	m.Remove("1")
	keys = m.Keys()

	if keys[0] != "2" || keys[1] != "3" || keys[2] != "4"{
		t.Error()
	}
	fmt.Println(keys)

	m.Each(func(k, v interface{}) bool {
		fmt.Println("k:", k, " ---> v:", v)
		return true
	})

	m.Clear()

	if !m.IsEmpty() || m.Size() > 0 {
		t.Error()
	}

	m.Each(func(k, v interface{}) bool {
		fmt.Println("k:", k, " ---> v:", v)
		t.Error()
		return true
	})

}
