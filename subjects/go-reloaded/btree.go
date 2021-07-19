package student

type TreeNode struct {
	Parent *TreeNode
	Left   *TreeNode
	Right  *TreeNode
	Data   string
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == node {
		return rplc
	} else if root == nil || node == nil {
		return root
	}
	if root.Data < node.Data {
		root.Right = BTreeTransplant(root.Right, node, rplc)
		root.Right.Parent = root
	} else if root.Data > node.Data {
		root.Left = BTreeTransplant(root.Left, node, rplc)
		root.Left.Parent = root
	}
	return root
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return root
	}
	if root.Data < data {
		if root.Right == nil {
			root.Right = &TreeNode{Parent: root, Data: data}
		} else {
			BTreeInsertData(root.Right, data)
		}
	} else {
		if root.Left == nil {
			root.Left = &TreeNode{Parent: root, Data: data}
		} else {
			BTreeInsertData(root.Left, data)
		}
	}
	return root
}

func BTreeSearchItem(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return nil
	} else if root.Data > data {
		return BTreeSearchItem(root.Left, data)
	} else if root.Data < data {
		return BTreeSearchItem(root.Right, data)
	} else { // root.Data == data
		return root
	}
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (n int, err error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	nodes, size := []*TreeNode{root}, 1
	for size > 0 {
		node := nodes[0]
		if node.Left != nil {
			nodes = append(nodes, node.Left)
			size++
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
			size++
		}
		f(node.Data)
		nodes = nodes[1:]
		size--
	}
}

func BTreeApplyForLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	} else if level == 1 {
		f(root.Data)
	} else if level > 1 {
		BTreeApplyForLevel(root.Left, level-1, f)
		BTreeApplyForLevel(root.Right, level-1, f)
	}
}

func GetMaxLevelOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lLvl := GetMaxLevelOfTree(root.Left) + 1
	rLvl := GetMaxLevelOfTree(root.Right) + 1
	return getMax(lLvl, rLvl)
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else if root.Data == node.Data {
		parent := root.Parent
		if root.Right == nil && root.Left == nil {
			return nil
		} else if root.Right == nil {
			root = root.Left
		} else if root.Left == nil {
			root = root.Right
		} else { // They Have Left and Right
			inserter := root.Right
			root = root.Left
			mover := root
			for mover.Right != nil {
				mover = mover.Right
			}
			mover.Right = inserter
			inserter.Parent = mover
		}
		root.Parent = parent
	} else if root.Data < node.Data { // root.Data < node.Data
		res := BTreeDeleteNode(root.Right, node)
		if root.Right != res {
			root.Right = res
		}
	} else { // root.Data > node.Data
		res := BTreeDeleteNode(root.Left, node)
		if root.Left != res {
			root.Left = res
		}
	}
	return root
}

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil || root.Right == nil || root.Left == nil {
		return true
	}
	checker := false
	if root.Left.Data < root.Data && root.Right.Data >= root.Data {
		checker = BTreeIsBinary(root.Right)
		if checker {
			checker = BTreeIsBinary(root.Left)
		}
	}
	return checker
}
