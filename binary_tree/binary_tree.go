/**
 * @Time : 2022/4/26 3:57 下午
 * @Author : frankj
 * @Email : frankxjkuang@gmail.com
 * @Description : 链表实现二叉树
 * @Revise : --
 */

package binary_tree

import (
	"github.com/frankxjkuang/go-space/link"
)

// TreeNode 二叉树
type TreeNode struct {
	Data  interface{} // 值
	Left  *TreeNode   // 左子数
	Right *TreeNode   // 右子数
}

// PreOrder 前序遍历：先访问根节点，再访问左子树，最后访问右子树
func (t *TreeNode) PreOrder() []interface{} {
	result := make([]interface{}, 0)
	if t == nil {
		return result
	}

	result = append(result, t.Data)
	result = append(result, t.Left.PreOrder()...)
	result = append(result, t.Right.PreOrder()...)
	return result
}

// MidOrder 中序遍历：先访问左子树，再访问根节点，最后访问右子树
func (t *TreeNode) MidOrder() []interface{} {
	result := make([]interface{}, 0)
	if t == nil {
		return result
	}
	result = append(result, t.Left.MidOrder()...)
	result = append(result, t.Data)
	result = append(result, t.Right.MidOrder()...)
	return result
}

// PostOrder 后序遍历：先访问左子树，再访问右子树，最后访问根节点
func (t *TreeNode) PostOrder() []interface{} {
	result := make([]interface{}, 0)
	if t == nil {
		return result
	}
	result = append(result, t.Left.PostOrder()...)
	result = append(result, t.Right.PostOrder()...)
	result = append(result, t.Data)
	return result
}

// 层次遍历：每一层从左到右访问每一个节点
func (t *TreeNode) LayerOrder() []interface{} {
	result := make([]interface{}, 0)
	if t == nil {
		return result
	}

	// 使用队列实现压栈
	linkAueue := new(link.LinkQueue)

	// 根节点先入栈
	linkAueue.Append(t)
	for linkAueue.Size > 0 {
		// 不断的出栈
		element := linkAueue.Shift()
		// 出栈元素是linkNode，需要转为二叉树节点（与入栈的时候类型一致）
		node, _ := (element.Value).(*TreeNode)

		result = append(result, node.Data)
		// 左子数非空继续入队列
		if node.Left != nil {
			linkAueue.Append(node.Left)
		}
		// 右子数非空继续入队列
		if node.Right != nil {
			linkAueue.Append(node.Right)
		}
	}

	return result
}
