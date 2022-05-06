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

// InsertSort 插入排序
func InsertSort(arr []int) []int {
	// 从第二个数开始
	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		// 待排序左边第一个数索引
		j := i - 1
		// 当前数比左边第一个数小，需要往前插入
		if curr < arr[j] {
			// 注意 j >=0 需要作为前置条件，避免数组越界
			for ;j >=0 && curr < arr[j]; j-- {
				// 比待排序大的数都往后挪一位
				arr[j+1] = arr[j]
			}
			// 插入待排数
			arr[j+1] = curr
		}
	}
	return arr
}

// ShellSort 希尔排序
func ShellSort(arr []int) []int {
	// 数组长度
	n := len(arr)

	// 每次减半，直到步长为 1
	for step := n / 2; step >= 1; step /= 2 {
		// fmt.Printf("step: %v \n", step)
		// 开始插入排序，每一轮的步长为 step
		for i := step; i < n; i += step {
			for j := i - step; j >= 0; j -= step {
				// 待排序的数
				curr := arr[i]
				// 待排序的数的左边最近一个数的下标
				j := i - step
				// 在有序列表中寻找待排序数的插入位置
				for ; j >= 0 && curr < arr[j]; j -= step {
					// 比待排序数大的数往后移
					arr[j+step] = arr[j]
				}
				// 插入待排数
				arr[j+step] = curr
			}
		}
	}
	return arr
}