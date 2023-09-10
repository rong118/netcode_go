/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 import "math"

 func levelOrder(root *TreeNode) [][]int {
	 var ans [][]int
	 if root == nil { return ans }
	 
	 var q []*TreeNode
	 q = append(q, root)
	 q = append(q, &TreeNode{ Val : math.MaxInt })
	 var l []int
	 for true {
		 f := q[0]
		 q = q[1:]
		 fmt.Println("f", f.Val)
		 if f.Val == math.MaxInt {
			 ans = append(ans, l)
			 l = make([]int, 0)
			 if len(q) == 0 { break }
			 q = append(q, f)
		 }else{
			 l = append(l, f.Val)
			 if f.Left != nil {
				 q = append(q, f.Left)
			 }
 
			 if f.Right != nil {
				 q = append(q, f.Right)
			 }
		 }
	 }
 
	 return ans
 }