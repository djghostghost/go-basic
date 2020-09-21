package set

import "github.com/djghostghost/go-basic/collection"

type Set interface {
	collection.Collection
	Each(f func(k, v interface{}) bool)
	Add(e interface{})
	Remove(e interface{})
	ToSlice()
}
