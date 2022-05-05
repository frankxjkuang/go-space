/**
 * @Time : 2022/5/5 3:23 下午
 * @Author : frankj
 * @Email : frankxjkuang@gmail.com
 * @Description : --
 * @Revise : --
 */

package sort

import "testing"

var (
	arrays = [][]int{
		{5, 9, 1, 111},
		{6},
		{},
		{1, 1, 3, 3, 3, 8, 2},
	}
)

func TestBubbleSort(t *testing.T) {
	for i, v := range arrays {
		t.Logf("%d ===> %v", i, BubbleSort(v))
	}
}

func TestSelectSort(t *testing.T) {
	for i, v := range arrays {
		t.Logf("%d ===> %v", i, BubbleSort(v))
	}
}