package Algorithm

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	t.Log(obj.Get(1))
	obj.Put(3, 3)
	t.Log(obj.Get(2))
	obj.Put(4, 4)
	t.Log(obj.Get(1))
	t.Log(obj.Get(3))
	t.Log(obj.Get(4))

	obj.Put(3, 3)
	t.Log(obj.Get(3))
	obj.Put(2, 2)
	obj.Put(4, 4)
	t.Log(obj.Get(4))
	t.Log(obj.Get(2))
	t.Log(obj.Get(3))
}
