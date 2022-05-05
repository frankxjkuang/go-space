/**
 * @Time : 2022/5/5 3:22 下午
 * @Author : frankj
 * @Email : frankxjkuang@gmail.com
 * @Description : --
 * @Revise : --
 */

package sort

// BubbleSort 冒泡排序
func BubbleSort(arr []int) []int {
	// 记录是否有过交换
	flag := false
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr) - i - 1; j ++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		// 走完一轮没有交换说明现在已经是有序数组了
		if !flag {
			return arr
		}
	}
	return arr
}

// SelectSort 选择排序
func SelectSort(arr []int) []int {
	for i := 0; i < len(arr) - 1; i++ {
		min := arr[i]
		idx := i
		for j := i+1; j < len(arr); j ++ {
			if arr[j] < min {
				min = arr[j]
				idx = j
			}
		}
		if idx != i {
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}
	return arr
}