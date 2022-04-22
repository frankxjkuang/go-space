/**
 * @Time : 2022/4/22 10:44 上午
 * @Author : 旷祥军
 * @Email : kuangxj@dustess.com
 * @Description : --
 * @Revise : --
 */

package array

import (
	"testing"
)

func TestMake(t *testing.T) {
	array := Make(10, 20)
	t.Logf("len: %v, cap: %v", array.Len(), array.Cap())

	array2 := Make(-1, 20)
	t.Logf("len: %v, cap: %v", array2.Len(), array2.Cap())

	array3 := Make(20, 10)
	t.Logf("len: %v, cap: %v", array3.Len(), array3.Cap())
}

func TestAppend(t *testing.T) {
	array := Make(1, 2)
	t.Logf("array: %+v, len: %v, cap: %v", array.array, array.Len(), array.Cap())
	array.Append(1)
	t.Logf("array: %+v, len: %v, cap: %v", array.array, array.Len(), array.Cap())
	array.Append(2)
	t.Logf("array: %+v, len: %v, cap: %v", array.array, array.Len(), array.Cap())
	array.Append(3)
	t.Logf("array: %+v, len: %v, cap: %v", array.array, array.Len(), array.Cap())
	array.Append(4)
	t.Logf("array: %+v, len: %v, cap: %v", array.array, array.Len(), array.Cap())
	array.Append(5)
	t.Logf("array: %+v, len: %v, cap: %v", array.array, array.Len(), array.Cap())
}

func TestAppendMany(t *testing.T) {
	array := Make(1, 2)
	t.Logf("array: %+v, len: %v, cap: %v", array.array, array.Len(), array.Cap())
	array.AppendMany(1, 2)
	t.Logf("array: %+v, len: %v, cap: %v", array.array, array.Len(), array.Cap())
	array.AppendMany(2, 3, 4, 5)
	t.Logf("array: %+v, len: %v, cap: %v", array.array, array.Len(), array.Cap())
}

func TestGet(t *testing.T) {
	array := Make(2, 2)
	idx := 1
	val := array.Get(idx)
	t.Logf("array[%d] = %d", idx, val)
}

func TestSet(t *testing.T) {
	array := Make(2, 2)
	idx := 1
	oriVal := array.Set(idx, 999)
	newVal := array.Get(idx)
	t.Logf("array oriVal=%d, newVal=%d", oriVal, newVal)
}