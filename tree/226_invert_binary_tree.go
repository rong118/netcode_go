/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func invertTree(root *TreeNode) *TreeNode {
    if root == nil { return nil }
    left := root.Left
    right := root.Right

    root.Right = invertTree(left)
    root.Left = invertTree(right)

    return root
}