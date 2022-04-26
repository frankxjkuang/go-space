/**
 * @Time : 2022/4/26 4:35 下午
 * @Author : 旷祥军
 * @Email : franjxjkuang@gmail.com
 * @Description : --
 * @Revise : --
 */

package binary_tree

import (
	"testing"
)

var tree = &TreeNode{}

func init()  {
	tree.Data = "A"
	tree.Left = &TreeNode{Data: "B"}
	tree.Right = &TreeNode{Data: "C"}
	tree.Left.Left = &TreeNode{Data: "D"}
	tree.Left.Right = &TreeNode{Data: "E"}
	tree.Right.Left = &TreeNode{Data: "F"}
}

func TestTreeNode_PreOrder(t *testing.T) {
	result := tree.PreOrder()
	t.Logf("PreOrder: %v", result)
	// out：[A B D E C F]
}

func TestTreeNode_MidOrder(t *testing.T) {
	result := tree.MidOrder()
	t.Logf("MidOrder: %v", result)
	// out：[D B E A F C]
}

func TestTreeNode_PostOrder(t *testing.T) {
	result := tree.PostOrder()
	t.Logf("PostOrder: %v", result)
	// out：[D E B F C A]
}

func TestTreeNode_LayerOrder(t *testing.T) {
	result := tree.LayerOrder()
	t.Logf("LayerOrder: %v", result)
	// out：[A B C D E F]
}