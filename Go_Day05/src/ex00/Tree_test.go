package Tree

import (
	"testing"
)

func TestTreeFalse(t *testing.T) {
	t.Run("Three nodes", func(t *testing.T) {
		/*
			  1
			 /

			1   0
		*/
		root := &TreeNode{HasToy: true}
		root.Left = CreateNode(true)
		root.Right = CreateNode(false)
		areBalanced := areToysBalanced(root)

		if false != areBalanced {
			t.Errorf("[Expected] %v != %v [Real]\\n", true, areBalanced)
		}
	})

	t.Run("Five nodes", func(t *testing.T) {
		/*
			  0
			 /

			1   0



			  1   1
		*/
		root := &TreeNode{HasToy: false}
		root.Left = CreateNode(true)
		root.Right = CreateNode(false)
		root.Right.Left = CreateNode(true)
		root.Right.Right = CreateNode(true)
		areBalanced := areToysBalanced(root)

		if false != areBalanced {
			t.Errorf("[Expected] %v != %v [Real]\\n", true, areBalanced)
		}
	})

	t.Run("Seven nodes", func(t *testing.T) {
		/*
			    1
			   /

			  1     0
			 /
			       /

			1   0 1   0
		*/
		root := &TreeNode{HasToy: true}
		root.Left = CreateNode(true)
		root.Left.Left = CreateNode(true)
		root.Left.Right = CreateNode(false)
		root.Right = CreateNode(false)
		root.Right.Left = CreateNode(true)
		root.Right.Right = CreateNode(false)
		areBalanced := areToysBalanced(root)

		if false != areBalanced {
			t.Errorf("[Expected] %v != %v [Real]\\n", true, areBalanced)
		}
	})
}

func TestTreeTrue(t *testing.T) {
	t.Run("Three nodes", func(t *testing.T) {
		/*
		    1
		   /

		  1   1
		*/
		root := &TreeNode{HasToy: true}
		root.Left = CreateNode(true)
		root.Right = CreateNode(true)
		areBalanced := areToysBalanced(root)

		if true != areBalanced {
			t.Errorf("[Expected] %v != %v [Real]\\n", true, areBalanced)
		}
	})

	t.Run("Five nodes", func(t *testing.T) {
		/*
			    0
			   /

			  0   1
			 /

			0   1
		*/
		root := &TreeNode{HasToy: false}
		root.Left = CreateNode(false)
		root.Left.Left = CreateNode(false)
		root.Left.Right = CreateNode(true)
		root.Right = CreateNode(true)
		areBalanced := areToysBalanced(root)

		if true != areBalanced {
			t.Errorf("[Expected] %v != %v [Real]\\n", true, areBalanced)
		}
	})

	t.Run("Seven nodes", func(t *testing.T) {
		/*
			    1
			   /

			  1     0
			 /
			       /

			1   0 1   1
		*/
		root := &TreeNode{HasToy: true}
		root.Left = CreateNode(true)
		root.Left.Left = CreateNode(true)
		root.Left.Right = CreateNode(false)
		root.Right = CreateNode(false)
		root.Right.Left = CreateNode(true)
		root.Right.Right = CreateNode(true)
		areBalanced := areToysBalanced(root)

		if true != areBalanced {
			t.Errorf("[Expected] %v != %v [Real]\\n", true, areBalanced)
		}
	})
}
