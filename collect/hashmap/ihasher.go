package hashmap

type IHasher interface {
    Hash(key interface{}) uint64
    Equals(a interface{}, b interface{}) bool
}