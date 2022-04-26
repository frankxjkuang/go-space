/**
 * @Time : 2022/4/26 4:57 下午
 * @Author : 旷祥军
 * @Email : frankxjkuang@gamil.com
 * @Description : 链表实现
 * @Revise : --
 */

package link

import "sync"

// LinkNode 单项链表节点
type LinkNode struct {
	Next  *LinkNode   // 下一个节点
	Value interface{} // 值
}

// LinkQueue 单项链表
type LinkQueue struct {
	root *LinkNode  // 根
	Size int        // 大小
	lock sync.Mutex // 安全锁
}

// Append 入队列：添加到队列末尾
func (lq *LinkQueue) Append(v interface{}) {
	lq.lock.Lock()
	defer lq.lock.Unlock()

	// 新节点，待添加到队列中
	node := &LinkNode{
		Value: v,
	}

	currNode := lq.root
	// 根节点为空直接赋值为根节点
	if currNode == nil {
		lq.root = node
	} else {
		// 否则添加到末尾
		for currNode.Next != nil {
			currNode = currNode.Next
		}
		currNode.Next = node
	}
	// 增加size
	lq.Size++
}

// Shift 出队列：删除头部节点（先进先出）
func (lq *LinkQueue) Shift() *LinkNode {
	lq.lock.Lock()
	defer lq.lock.Unlock()

	if lq.Size == 0 {
		panic("LinkQueue is empty")
	}
	// 顶部出栈
	root := lq.root
	// 下一个节点变长根节点
	lq.root = root.Next
	lq.Size--
	// 返回栈顶的节点
	return &LinkNode{Value: root.Value}
}