/**
 * @Time : 2022/5/9 3:06 下午
 * @Author : frankj
 * @Email : kuangxj@dustess.com
 * @Description : --
 * @Revise : --
 */

package set

import "testing"

func TestNew(t *testing.T) {
	arr := []int{-19, 1, 0, 5}
	for i, v := range arr {
		set := New(v)
		t.Logf("New(%d) set is: %v, len=%d", i, set.m, set.len)
	}
}

func TestAdd(t *testing.T) {
	arr := []int{-19, 1, 0, 5}
	for i, v := range arr {
		set := New(v)
		set.Add(999)
		t.Logf("New(%d) set is: %v, len=%d", i, set.m, set.len)
	}
}

func TestRemove(t *testing.T) {
	arr := []int{-19, 1, 0, 5}
	for i, v := range arr {
		set := New(v)
		set.Add(999)
		set.Remove(1)
		set.Remove(999)
		t.Logf("New(%d) set is: %v, len=%d", i, set.m, set.len)
	}
}

func TestHas(t *testing.T) {
	set := New(2)
	t.Logf("set has 999? %v", set.Has(999))
	set.Add(999)
	t.Logf("set has 999? %v", set.Has(999))
}

func TestRemoveAll(t *testing.T) {
	set := New(2)
	set.Add(999)
	set.Add(88)
	set.RemoveAll()
	t.Logf("New set is: %v, len=%d", set.m, set.len)
}

func TestList(t *testing.T) {
	set := New(2)
	set.Add(999)
	set.Add(88)
	//set.Add(7)
	list := set.List()
	t.Logf("New set list is: %v, len=%d", list, set.len)
}
