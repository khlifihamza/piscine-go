package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root != nil {
		leftLevel := BTreeLevelCount(root.Left)
		rightLevel := BTreeLevelCount(root.Right)
		if leftLevel > rightLevel {
			return leftLevel + 1
		} else {
			return rightLevel + 1
		}
	}
	return 0
}
