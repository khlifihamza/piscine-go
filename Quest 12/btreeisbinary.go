package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root != nil {
		if root.Left != nil {
			return BTreeIsBinary(root.Left) && root.Data > root.Left.Data
		}
		if root.Left != nil {
			return BTreeIsBinary(root.Right) && root.Data < root.Right.Data
		}
	}
	return true
}
