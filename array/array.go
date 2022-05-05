/**
 * @Time : 2022/4/22 10:40 上午
 * @Author : frankj
 * @Email : kuangxj@dustess.com
 * @Description : --
 * @Revise : --
 */
package array

import "sync"

// array 可变长int数组
type array struct {
	array []int      // 基于固定大小数组实现
	len   int        // 长度
	cap   int        // 容量
	lock  sync.Mutex // 并发安全锁
}

// Make 创建一个 array 数组
func Make(len, cap int) *array {
	if len < 0 || cap < 0 {
		panic("len and cap cannot be less than 0")
	}
	if len > cap {
		panic("cannot create slice with len greater than cap")
	}

	return &array{
		array: make([]int, cap, cap),
		len:   len,
		cap:   cap,
	}
}

// Append 增加单个元素
func (a *array) Append(ele int) {
	// 并发锁
	a.lock.Lock()
	defer a.lock.Unlock()

	// 元素满了
	if a.len == a.cap {
		// 扩容
		newCap := 2 * a.cap

		// 如果之前容量为0，那么新的容量设置为1
		if a.cap == 0 {
			newCap = 1
		}

		newArray := make([]int, newCap, newCap)

		// 原数组元素移动到新数组
		for i, v := range a.array {
			newArray[i] = v
		}
		a.array = newArray
		a.cap = newCap
	}
	// 新增元素
	a.array[a.len] = ele
	// 长度自增
	a.len += 1
}

// AppendMany 增加多个元素
func (a *array) AppendMany(ele ...int) {
	for _, v := range ele {
		a.Append(v)
	}
}

// Get 按下标获取元素
func (a *array) Get(idx int) int {
	// 数据越界
	if a.len == 0 || idx >= a.len {
		panic("array out of bounds")
	}
	return a.array[idx]
}

// Set 按下标赋值元素
// 返回原来的值
func (a *array) Set(idx, val int) (origin int) {
	// 数据越界
	if a.len == 0 || idx >= a.len {
		panic("array out of bounds")
	}
	origin = a.Get(idx)
	a.array[idx] = val
	return origin
}

// Len 获取长度
func (a *array) Len() int {
	return a.len
}

// Cap 获取容量
func (a *array) Cap() int {
	return a.cap
}

