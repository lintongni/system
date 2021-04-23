package hashmap

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_HashMap(t *testing.T) {
	fnAssertTrue := func(b bool, msg string) {
		if !b {
			t.Fatal(msg)
		}
	}

	fnEquals := func(a interface{}, b interface{}) bool {
		i1, b1 := a.(int)
		if b1 {
			i2, b2 := b.(int)
			if b2 {
				return i1 == i2
			}
		}

		s1, b1 := a.(string)
		if b1 {
			s2, b2 := b.(string)
			if b2 {
				return s1 == s2
			}
		}

		return a == b
	}

	hasher := NewCrc64Hashful(fnEquals)
	hm := NewHashMap(hasher, 2)

	hm.Put(1, "10")
	t.Log(hm)
	fnAssertTrue(hm.Size() == 1, "expecting size == 1")
	fnAssertTrue(hm.IsNotEmpty(), "expecting not empty")

	ok, v := hm.Get(1)
	fnAssertTrue(ok, "expecting ok")
	fnAssertTrue(v == "10", "expecting 10")

	hm.Put(2, "20")
	hm.Put(3, "30")
	hm.Put(4, "40")
	hm.Put(5, "50")
	t.Log(hm)
	fnAssertTrue(hm.Size() == 5, "expecting size == 5")
	ok, v = hm.Get(2)
	fnAssertTrue(ok == true && v == "20", "expecting true and 20")

	hm.Clear()
	t.Log(hm)
	fnAssertTrue(hm.Size() == 0, "expecting size == 0")
	fnAssertTrue(hm.IsEmpty(), "expecting empty")

	iter := hm.Iterator()
	fnAssertTrue(!iter.More(), "expecting no more")

	hm.Put(1, "10")
	hm.Put(2, "20")
	hm.Put(3, "30")
	t.Log(hm)
	fnAssertTrue(hm.Has(1) && hm.Has(2) && hm.Has(3) && !hm.Has(4), "expecting has 1,2,3")

	hm.Put(4, "40")
	hm.Put(5, "50")
	hm.Put(6, "60")
	t.Log(hm)
	iter = hm.Iterator()
	fnAssertTrue(iter.More(), "expecting more")
	e, k, v := iter.Next()
	t.Logf("%v>%s", k, v)
	fnAssertTrue(e == nil, "e == nil")
	e, k, v = iter.Next()
	t.Logf("%v>%s", k, v)
	fnAssertTrue(e == nil, "e == nil")
	e, k, v = iter.Next()
	t.Logf("%v>%s", k, v)
	fnAssertTrue(e == nil, "e == nil")
	e, k, v = iter.Next()
	t.Logf("%v>%s", k, v)
	fnAssertTrue(e == nil, "e == nil")
	e, k, v = iter.Next()
	t.Logf("%v>%s", k, v)
	fnAssertTrue(e == nil, "e == nil")
	e, k, v = iter.Next()
	t.Logf("%v>%s", k, v)
	fnAssertTrue(e == nil, "e == nil")
	e, k, v = iter.Next()
	fnAssertTrue(e != nil, "expecting e != nil")

	ok, v = hm.Remove(3)
	t.Log(hm)
	fnAssertTrue(ok && v == "30" && hm.Size() == 5, "expecting remove 30")

	ok, v = hm.Remove(2)
	t.Log(hm)
	fnAssertTrue(ok && v == "20" && hm.Size() == 4, "expecting remove 20")

	t0 := time.Now().UnixNano()
	hm.Clear()
	size := 10000 * 1000
	for i := 0; i < size; i++ {
		hm.Put(strconv.Itoa(i), i)
	}
	millis := (time.Now().UnixNano() - t0) / 1000000
	fmt.Println("百万存储时间：", millis)
	t.Logf("putting %v string>int = %v ms", size, millis)
	fnAssertTrue(hm.Size() == size, fmt.Sprintf("expecting %v", size))

	for i := 0; i < size; i++ {
		ok, v = hm.Get(strconv.Itoa(i))
		fnAssertTrue(ok == true && v == i, "expecting i")
	}
}
