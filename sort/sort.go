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
		for j := 0; j < len(arr)-i-1; j++ {
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
	for i := 0; i < len(arr)-1; i++ {
		min := arr[i]
		idx := i
		for j := i + 1; j < len(arr); j++ {
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
			for ; j >= 0 && curr < arr[j]; j-- {
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

// MergeSort 归并排序
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	leftArr := arr[0:mid]
	rightArr := arr[mid:]
	// 递归切分数组
	leftArr = MergeSort(leftArr)
	rightArr = MergeSort(rightArr)

	lLen := len(leftArr)
	rLen := len(rightArr)
	l, r := 0, 0

	// 重排的结果容量应该是跟输入的数组长度一样
	result := make([]int, 0, len(arr))
	for {
		// 左边部分先遍历结束，直接将右边剩余部分追加到结果中
		if l >= lLen {
			result = append(result, rightArr[r:]...)
			break
		}
		// 右边部分先遍历结束，直接将左边剩余部分追加到结果中
		if r >= rLen {
			result = append(result, leftArr[l:]...)
			break
		}
		// 比较左右两边元素，将较小的追加到结果中
		// 增加下标，比较下一个元素的情况
		if leftArr[l] < rightArr[r] {
			result = append(result, leftArr[l])
			l++
		} else {
			result = append(result, rightArr[r])
			r++
		}
	}

	return result
}

// QuickSort 快速排序
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// 取第一个数作为标杆
	flag := arr[0]
	leftArr, rightArr := make([]int, 0), make([]int, 0)
	for i := 1; i < len(arr); i++ {
		// 比第一个数小的放左边
		if arr[i] < flag {
			leftArr = append(leftArr, arr[i])
		} else {
			// 反之放右边
			rightArr = append(rightArr, arr[i])
		}
	}
	// 左右数组递归排序
	lResult := QuickSort(leftArr)
	rResult := QuickSort(rightArr)
	// 最终将左右数组连接起来就是排好的结果
	return append(append(lResult, flag), rResult...)
}
