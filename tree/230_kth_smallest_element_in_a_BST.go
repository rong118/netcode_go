/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 var ans []int

 func kthSmallest(root *TreeNode, k int) int {
	 dfs(root)
	 return ans[k - 1]
 }
  
 func dfs(root *TreeNode){
	 if root == nil {
		 return
	 }
	 dfs(root.Left)
	 ans = append(ans, root.Val)
	 dfs(root.Right)
 }

// stack dfs
 func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0, k)

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		k--
		if k == 0 {
			return root.Val
		}

		root = root.Right
	}

	return -1
}
