package hashmap

import (
	"fmt"
	"strings"
)

type tPartition struct {
	hasher      IHasher
	buckets     []*tBucket
	bucketCount uint64
	prev        *tPartition
	next        *tPartition
	size        int
	threshhold  int
}

func newPartition(hasher IHasher, bucketCount int) *tPartition {
	it := &tPartition{
		hasher:      hasher,
		buckets:     make([]*tBucket, bucketCount),
		bucketCount: uint64(bucketCount),
		prev:        nil,
		next:        nil,
		size:        0,
		threshhold:  bucketCount * 3 / 4,
	}
	for i, _ := range it.buckets {
		it.buckets[i] = newBucket(hasher)
	}
	return it
}

func (me *tPartition) putByKey(key interface{}, value interface{}) bool {
	hash := me.hasher.Hash(key)
	return me.putByKeyAndHash(key, value, hash)
}

func (me *tPartition) putByKeyAndHash(key interface{}, value interface{}, hash uint64) bool {
	if me.getBucketByHash(hash).put(key, value) {
		me.size++
		return true
	}
	return false
}

func (me *tPartition) appendByKeyAndHash(key interface{}, value interface{}, hash uint64) {
	me.getBucketByHash(hash).append(key, value)
	me.size++
}

func (me *tPartition) getBucketByKey(key interface{}) *tBucket {
	hash := me.hasher.Hash(key)
	return me.getBucketByHash(hash)
}

func (me *tPartition) getBucketByHash(hash uint64) *tBucket {
	return me.buckets[int(hash%me.bucketCount)]
}

func (me *tPartition) get(key interface{}) (bool, interface{}) {
	return me.getBucketByKey(key).get(key)
}

func (me *tPartition) findByKey(key interface{}) (ok bool, bucket *tBucket, node *tLinkedNode, prev *tLinkedNode) {
	bucket = me.getBucketByKey(key)
	ok, node, prev = bucket.find(key)
	return ok, bucket, node, prev
}

func (me *tPartition) findByKeyAndHash(key interface{}, hash uint64) (ok bool, bucket *tBucket, node *tLinkedNode, prev *tLinkedNode) {
	bucket = me.getBucketByHash(hash)
	ok, node, prev = bucket.find(key)
	return ok, bucket, node, prev
}

func (me *tPartition) remove(key interface{}) (bool, value interface{}) {
	ok, node := me.getBucketByKey(key).remove(key)
	if ok {
		me.size--
		return true, node.value

	} else {
		return false, nil
	}
}

func (me *tPartition) nearlyFull() bool {
	return me.size >= me.threshhold
}

func (me *tPartition) String() string {
	itemStrings := make([]string, 0)
	for i, b := range me.buckets {
		if b.size > 0 {
			itemStrings = append(itemStrings, fmt.Sprintf("%v:%s", i, b.String()))
		}
	}
	return fmt.Sprintf("p[%s]", strings.Join(itemStrings, ","))
}
