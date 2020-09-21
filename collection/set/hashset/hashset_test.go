package hashset

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	s.Each(func(k interface{}) bool {
		fmt.Println(k)
		return true
	})

	if s.Size() != 3 {
		t.Error()
	}

	if s.IsEmpty() {
		t.Error()
	}

	s.Remove(1)
	if s.Size() != 2 {
		t.Error()
	}

	s.Remove("1")
	if s.Size() != 2 {
		t.Error()
	}

	arr, err := s.ToIntSlice()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(arr)

	arrI := s.ToSlice()
	if len(arrI) != 2 {
		t.Error()
	}

	e := s.Contains("2")
	if e {
		t.Error()
	}
	e = s.Contains(2)
	if !e {
		t.Error()
	}
	s.Clear()
	if !s.IsEmpty() {
		t.Error()
	}
	if s.Size() != 0 {
		t.Error()
	}

	ss := New()
	ss.Add("1")
	ss.Add("2")

	as, err := ss.ToStringSlice()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(as)
}
