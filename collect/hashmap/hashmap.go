package hashmap

import (
	"fmt"
	"strings"
)

type tHashMap struct {
	hasher     IHasher
	partitions *tPartition
	size       int
	version    int
}

func NewHashMap(hasher IHasher, capacity int) IMap {
	if capacity < 4 {
		capacity = 4
	}

	part := newPartition(hasher, capacity)
	return &tHashMap{
		hasher:     hasher,
		partitions: part,
		size:       0,
		version:    0,
	}
}

func (me *tHashMap) Size() int {
	return me.size
}

func (me *tHashMap) IsEmpty() bool {
	return me.Size() <= 0
}

func (me *tHashMap) IsNotEmpty() bool {
	return !me.IsEmpty()
}

func (me *tHashMap) Put(key interface{}, value interface{}) {
	hash := me.hasher.Hash(key)
	ok, _, bucket, node, _ := me.findByKeyAndHash(key, hash)
	if ok {
		bucket.putAt(node, key, value)

	} else {
		if me.partitions.nearlyFull() {
			// create new partition
			part := newPartition(me.hasher, int(me.partitions.bucketCount*2))
			part.next = me.partitions
			me.partitions.prev = part
			me.partitions = part
			part.appendByKeyAndHash(key, value, hash)
		} else {
			me.partitions.appendByKeyAndHash(key, value, hash)
		}

		me.size++
	}

	me.version++
}

func (me *tHashMap) findByKey(key interface{}) (ok bool, part *tPartition, bucket *tBucket, node *tLinkedNode, prev *tLinkedNode) {
	hash := me.hasher.Hash(key)
	return me.findByKeyAndHash(key, hash)
}

func (me *tHashMap) findByKeyAndHash(key interface{}, hash uint64) (ok bool, part *tPartition, bucket *tBucket, node *tLinkedNode, prev *tLinkedNode) {
	for part = me.partitions; part != nil; part = part.next {
		ok, bucket, node, prev = part.findByKeyAndHash(key, hash)
		if ok {
			return ok, part, bucket, node, prev
		}
	}

	return false, nil, nil, nil, nil
}

func (me *tHashMap) Get(key interface{}) (bool, interface{}) {
	ok, _, _, node, _ := me.findByKey(key)
	if ok {
		return true, node.value
	} else {
		return false, nil
	}
}

func (me *tHashMap) Has(key interface{}) bool {
	ok, _, _, _, _ := me.findByKey(key)
	return ok
}

func (me *tHashMap) Remove(key interface{}) (ok bool, value interface{}) {
	ok, part, bucket, node, prev := me.findByKey(key)
	if ok {
		value = node.value
		bucket.removeAt(node, prev)

		// 缩容
		if part.size <= 0 && part != me.partitions {
			me.removePartition(part)
		}

		me.size--
		me.version++
		return ok, value

	} else {
		return false, nil
	}
}

func (me *tHashMap) removePartition(part *tPartition) {
	prev := part.prev
	next := part.next

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}

	if me.partitions == part {
		me.partitions = next
	}
}

func (me *tHashMap) Clear() {
	if me.IsEmpty() {
		return
	}

	part := newPartition(me.hasher, len(me.partitions.buckets))
	me.partitions = part
	me.size = 0

	me.version++
}

func (me *tHashMap) Iterator() IMapIterator {
	return newHashMapIterator(me)
}

func (me *tHashMap) String() string {
	itemStrings := make([]string, 0)
	for p := me.partitions; p != nil; p = p.next {
		itemStrings = append(itemStrings, p.String())
	}
	return fmt.Sprintf("s=%v,v=%v %s", me.size, me.version, strings.Join(itemStrings, " "))
}
