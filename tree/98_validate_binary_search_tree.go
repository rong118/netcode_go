/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 var ans []int

 func isValidBST(root *TreeNode) bool {
	 dfs(root)
 
	 for i := 0; i < len(ans) - 1; i++ {
		 if ans[i] > ans[i + 1]{
			 return false
		 }
	 }
 
	 return true
 }
 
 func dfs(root *TreeNode){
	 if root == nil {
		 return
	 }
	 dfs(root.Left)
	 ans = append(ans, root.Val)
	 dfs(root.Right)
 }

// recursive way
func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= min.Val {
		return false
	}

	if max != nil && root.Val >= max.Val {
		return false
	}

	return isValid(root.Left, min, root) && isValid(root.Right, root, max)
}