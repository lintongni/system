package hashmap

import (
	"fmt"
	"strings"
)

type tBucket struct {
	hasher IHasher
	nodes  *tLinkedNode
	size   int
}

type tLinkedNode struct {
	key   interface{}
	value interface{}
	next  *tLinkedNode
}

func newBucket(hasher IHasher) *tBucket {
	return &tBucket{
		hasher: hasher,
		nodes:  nil,
		size:   0,
	}
}

func newLinkedNode(key interface{}, value interface{}) *tLinkedNode {
	return &tLinkedNode{
		key:   key,
		value: value,
		next:  nil,
	}
}

func (me *tBucket) put(key interface{}, value interface{}) bool {
	existed, node, _ := me.find(key)
	me.putAt(node, key, value)
	return !existed
}

func (me *tBucket) append(key interface{}, value interface{}) {
	me.putAt(nil, key, value)
}

func (me *tBucket) putAt(node *tLinkedNode, key interface{}, value interface{}) {
	if node != nil {
		node.value = value
		return
	}

	it := newLinkedNode(key, value)
	if me.nodes == nil {
		me.nodes = it

	} else {
		it.next = me.nodes
		me.nodes = it
	}

	me.size++
}

func (me *tBucket) get(key interface{}) (bool, interface{}) {
	ok, node, _ := me.find(key)
	if ok {
		return true, node.value
	}
	return false, nil
}

func (me *tBucket) find(key interface{}) (ok bool, node *tLinkedNode, prev *tLinkedNode) {
	prev = nil
	for it := me.nodes; it != nil; it = it.next {
		if me.hasher.Equals(it.key, key) {
			return true, it, prev
		}
		prev = it
	}
	return false, nil, nil
}

func (me *tBucket) remove(key interface{}) (bool, *tLinkedNode) {
	ok, node, prev := me.find(key)
	if !ok {
		return false, nil
	}

	me.removeAt(node, prev)
	return true, node
}

func (me *tBucket) removeAt(node *tLinkedNode, prev *tLinkedNode) {
	if prev == nil {
		me.nodes = node.next
	} else {
		prev.next = node.next
	}
	me.size--
}

func (me *tBucket) String() string {
	itemStrings := make([]string, me.size)
	i := 0
	for it := me.nodes; it != nil; it = it.next {
		itemStrings[i] = fmt.Sprintf("%v", it.key)
		i++
	}
	return fmt.Sprintf("b[%v %s]", me.size, strings.Join(itemStrings, ","))
}
