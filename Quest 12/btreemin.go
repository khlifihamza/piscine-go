package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root.Left != nil {
		root.Left.Parent = root
		return BTreeMin(root.Left)
	}
	return root
}
