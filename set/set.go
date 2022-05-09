/**
 * @Time : 2022/5/9 3:03 下午
 * @Author : frankj
 * @Email : kuangxj@dustess.com
 * @Description : 不可重复集合
 * @Revise : --
 */

package set

import "sync"

type Set struct {
	// 空结构体并且不占用内存
	m            map[int]struct{} // 使用map实现，因为map的key不能重复
	len          int              // 集合的大小
	cap          int              // 集合的容量
	sync.RWMutex                  // 并发安全锁
}

// New 初始化容量为cap的集合
func New(cap int) *Set {
	m := make(map[int]struct{}, cap)
	return &Set{
		m: m,
		cap: cap,
	}
}

// Add 增加元素
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	// 这里可以不要限制，扩容就没有作用了
	if s.len == s.cap {
		panic("out of range")
	}
	s.m[item] = struct{}{} // 更新map
	s.len = len(s.m)       // 更新大小
}

// Remove 删除元素
func (s *Set) Remove(item int) {
	s.Lock()
	defer s.Unlock()
	// 同 Add
	delete(s.m, item)
	s.len = len(s.m)
}

// Has 检查元素是否存在
func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// Len 检查集合的大小
func (s *Set) Len() int {
	return s.len
}

// RemoveAll 删除所有集合元素
func (s *Set) RemoveAll() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]struct{}{}
	s.len = 0
}

// List 将集合转为数组
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, s.len)
	for k := range s.m {
		list = append(list, k)
	}
	return list
}