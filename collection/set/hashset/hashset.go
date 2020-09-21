package hashset

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// 内部依然用hmap实现
type HashSet struct {
	hash map[interface{}]nothing
}

type nothing struct {

}

func New() HashSet {
	return HashSet{hash: make(map[interface{}]nothing)}
}

func (hs *HashSet) Add(elem interface{}) bool {
	_, exists := hs.hash[elem]
	hs.hash[elem] = nothing{}
	return exists
}

func (hs *HashSet) Remove(elem interface{}) {
	delete(hs.hash, elem)
}

func (hs *HashSet) Contains(elem interface{}) bool {
	_, exists := hs.hash[elem]
	return exists
}

func (hs *HashSet) IsEmpty() bool {
	return len(hs.hash) == 0
}

func (hs *HashSet) ToSlice() []interface{} {
	var slice []interface{}
	for k := range hs.hash {
		slice = append(slice, k)
	}
	return slice
}

func (hs *HashSet) ToIntSlice() ([]int, error) {
	var slice []int
	for k := range hs.hash {
		var v int
		ki := reflect.TypeOf(k).Kind()
		switch ki {
		case reflect.String:
			vv, err := strconv.Atoi(k.(string))
			if err != nil {
				return []int{}, err
			}
			v = vv
			break
		case reflect.Int:
			v = k.(int)
			break
		default:
			return []int{}, errors.New(fmt.Sprint("cant convert to int , kind is ", ki))
		}
		slice = append(slice, v)
	}
	return slice, nil
}

func (hs *HashSet) ToStringSlice() ([]string, error) {
	var slice []string
	for k := range hs.hash {
		var v string
		ki := reflect.TypeOf(k).Kind()
		switch ki {
		case reflect.String:
			v = k.(string)
			break
		default:
			return []string{}, errors.New(fmt.Sprint("element kind only string, but this is ", ki))
		}
		slice = append(slice, v)
	}
	return slice, nil
}

/**
if f return false, break the for each!
 */
func (hs *HashSet) Each(f func(kk interface{}) bool) {
	for k := range hs.hash {
		if !f(k) {
			break
		}
	}
}

func (hs *HashSet) Size() int {
	return len(hs.hash)
}

func (hs *HashSet) Clear() {
	hs.hash = make(map[interface{}]nothing)
}



