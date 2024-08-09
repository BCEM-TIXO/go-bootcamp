package Tree

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func CreateNode(value bool) *TreeNode {
	return &TreeNode{
		HasToy: value,
		Left:   nil,
		Right:  nil,
	}
}

func areToysBalanced(head *TreeNode) bool {
	if head == nil {
		return true
	}
	return getToysCount(head.Left) == getToysCount(head.Right)
}

func getToysCount(head *TreeNode) int {
	if head == nil {
		return 0
	}

	if head.HasToy {
		return 1 + getToysCount(head.Left) + getToysCount(head.Right)
	}

	return getToysCount(head.Left) + getToysCount(head.Right)
}
