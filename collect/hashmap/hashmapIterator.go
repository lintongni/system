package hashmap

import "errors"

type tHashMapIterator struct {
	hashMap *tHashMap
	pindex  *tPartition
	bindex  int
	nindex  *tLinkedNode
	version int
	visited int
}

func newHashMapIterator(hashMap *tHashMap) IMapIterator {
	return &tHashMapIterator{
		hashMap: hashMap,
		pindex:  hashMap.partitions,
		bindex:  -1,
		nindex:  nil,
		version: hashMap.version,
		visited: 0,
	}
}

func (me *tHashMapIterator) nextNode() *tLinkedNode {
	node := me.nindex
	for {
		if node == nil {
			bkt := me.nextBucket()
			if bkt == nil {
				return nil
			} else {
				me.nindex = bkt.nodes
				return me.nindex
			}

		} else {
			node = node.next
			if node != nil {
				me.nindex = node
				return node
			}
		}
	}
}

func (me *tHashMapIterator) nextBucket() *tBucket {
	part := me.pindex
	bi := me.bindex + 1

	for {
		if bi >= len(part.buckets) {
			part = me.nextPartition()
			if part == nil {
				return nil
			}

			bi = 0
		}

		bkt := part.buckets[bi]
		if bkt.nodes != nil {
			me.bindex = bi
			return bkt
		}

		bi++
	}
}

func (me *tHashMapIterator) nextPartition() *tPartition {
	if me.pindex == nil {
		return nil
	}
	me.pindex = me.pindex.next
	return me.pindex
}

func (me *tHashMapIterator) More() bool {
	if me.version != me.hashMap.version {
		return false
	}
	return me.visited < me.hashMap.size
}

func (me *tHashMapIterator) Next() (err error, key interface{}, value interface{}) {
	if me.version != me.hashMap.version {
		return gConcurrentModificationError, nil, nil
	}

	if !me.More() {
		return gNoMoreElementsError, nil, nil
	}

	node := me.nextNode()
	if node == nil {
		return gNoMoreElementsError, nil, nil

	} else {
		me.visited++
		return nil, node.key, node.value
	}
}

var gConcurrentModificationError = errors.New("concurrent modification error")
var gNoMoreElementsError = errors.New("no more elements")
