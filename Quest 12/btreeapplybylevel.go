package piscine

func printLevel(root *TreeNode, f func(...interface{}) (int, error), level int) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else {
		printLevel(root.Left, f, level-1)
		printLevel(root.Right, f, level-1)
	}
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	h := BTreeLevelCount(root)
	for i := 1; i <= h; i++ {
		printLevel(root, f, i)
	}
}
