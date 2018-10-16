package RBTree

const (
	RED = true
	BLACK = false
)

type RBTree struct {
	root *RBNode
}

func (t *RBTree) Insert(value int) {
	newNode := NewRBNode(RED, value, nil, nil, nil)
	t.insert(newNode)
}

func (t *RBTree) insert(n *RBNode) bool {
	var current, x *RBNode

	for x = t.root; x != nil ; { // 找到插入的位置
		current = x
		if cmp := n.CompareWith(x); cmp == 1 {
			x = x.right
		} else if cmp == -1 {
			x = x.left
		} else {
			return false
		}
	}
	n.parent = current // 确定了n的父节点

	if current != nil { // 确定新节点是左/右节点
		if cmp := n.CompareWith(current); cmp == 1 {
			current.right = n
		} else {
			current.left = n
		}
	} else {
		t.root = n
	}

	t.insertFixUp(n)
	return true
}

func (t *RBTree) insertFixUp(node *RBNode) {
	current := *node
	var parent *RBNode
	var gParent *RBNode

	for parent = current.parent; parent != nil && parent.IsRed() ; {
		// 修整的条件：父节点存在，且父节点的颜色是红色
		gParent = parent.parent
		if parent == gParent.left {
			uncle := gParent.right
			// case i
			if uncle != nil && uncle.IsRed() {
				parent.SetBlack()
				uncle.SetBlack()
				gParent.SetBlack()
				current = *gParent
				continue
			}

			//case ii
		} else {
			uncle := gParent.left

		}
	}
}