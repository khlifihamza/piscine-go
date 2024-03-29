package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root.Right != nil {
		root.Right.Parent = root
		return BTreeMax(root.Right)
	}
	return root
}
