/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        if subRoot == nil {
            return true
        }else{
            return false
        }
    }

    if subRoot == nil {
        if root == nil {
            return true
        }else{
            return false
        }
    }

    if root.Val == subRoot.Val && isSameTree(root, subRoot) {
        return true
    }

    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

 func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil {
        if q == nil {
            return true
        }else {
            return false
        }
    }

    if q == nil {
        if p == nil {
            return true
        }else {
            return false
        }
    }

    if p.Val != q.Val {
        return false
    }

    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}