package unroll

import tree "day5/ex00"

func UnrollGarland(head *tree.TreeNode) []bool {
	if head == nil {
		return []bool{}
	}

	queue := []*tree.TreeNode{head}
	values := []bool{head.HasToy}
	isEven := true
	for len(queue) > 0 {
		var curNode *tree.TreeNode
		levelSize := len(queue)
		for i, j := 0, levelSize; i < levelSize; i, j = i+1, j-1 {
			if isEven {
				curNode = queue[i]
				if curNode.Left != nil {
					queue = append(queue, curNode.Left)
					values = append(values, curNode.Left.HasToy)
				}
				if curNode.Right != nil {
					queue = append(queue, curNode.Right)
					values = append(values, curNode.Right.HasToy)

				}
			} else {
				curNode = queue[j-1]
				if curNode.Right != nil {
					queue = append(queue, curNode.Right)
					values = append(values, curNode.Right.HasToy)

				}
				if curNode.Left != nil {
					queue = append(queue, curNode.Left)
					values = append(values, curNode.Left.HasToy)
				}
			}

		}
		queue = queue[levelSize:]
		isEven = !isEven
	}
	return values
}
