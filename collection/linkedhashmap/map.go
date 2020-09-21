package linkedhashmap

import "github.com/djghostghost/go-basic/collection"

type Map interface {
	collection.Collection
	Each(f func(k, v interface{}) bool)
	Put(k, v interface{})
	Get(k interface{}) (interface{}, bool)
	Remove(k interface{}) (interface{}, bool)
	Keys() []interface{}
}
