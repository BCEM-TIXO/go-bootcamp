package unroll

import (
	tree "day5/ex00"
	"reflect"
	"testing"
)

func TestUnroll(t *testing.T) {
	root := &tree.TreeNode{HasToy: true}
	root.Left = tree.CreateNode(true)
	root.Right = tree.CreateNode(false)
	root.Left.Left = tree.CreateNode(true)
	root.Left.Right = tree.CreateNode(false)
	root.Right.Left = tree.CreateNode(true)
	root.Right.Right = tree.CreateNode(true)
	expectedResult := []bool{true, true, false, true, true, false, true}
	myResult := UnrollGarland(root)
	t.Run(
		"TestCase1", func(t *testing.T) {
			if !reflect.DeepEqual(expectedResult, myResult) {
				t.Errorf("[Expected] %v != %v [Real]\\n", expectedResult, myResult)
			}
		},
	)
}
