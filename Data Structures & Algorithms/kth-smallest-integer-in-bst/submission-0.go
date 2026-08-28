/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var count int
var result int

func kthSmallest(root *TreeNode, k int) int {
   	count = k
	inorder(root)
	return result
}

func inorder(root *TreeNode){
	if root == nil{
		return
	}
	inorder(root.Left)
	count--
	if count==0{
		result = root.Val
		return
	}
	inorder(root.Right)
}