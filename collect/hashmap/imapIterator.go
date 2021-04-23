package hashmap

type IMapIterator interface {
	More() bool
	Next() (err error, key interface{}, value interface{})
}
