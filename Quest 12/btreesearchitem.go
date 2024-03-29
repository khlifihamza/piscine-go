package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root != nil {
		if root.Data == elem {
			return root
		}
		if root.Data > elem && root.Left != nil {
			root.Left.Parent = root
			return BTreeSearchItem(root.Left, elem)
		}
		if root.Data < elem && root.Right != nil {
			root.Right.Parent = root
			return BTreeSearchItem(root.Right, elem)
		}
	}
	return nil
}
