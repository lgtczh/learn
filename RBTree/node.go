package RBTree

type RBNode struct {
	color  bool   //颜色
	value  int    //关键字(键值)
	parent *RBNode //父节点
	left   *RBNode //左子节点
	right  *RBNode //右子节点
}

func NewRBNode(color bool,value int, parent *RBNode, left *RBNode, right *RBNode) *RBNode{
	return &RBNode{color, value, parent, left, right}
}

/*************对红黑树节点n进行左旋操作 ******************/
/*
 * 左旋示意图：对节点n进行左旋
 *     p                       p
 *    /                       /
 *   n                       y
 *  / \                     / \
 * ln  y      ----->       n  ry
 *    / \                 / \
 *   ly ry               ln ly
 * 左旋做了三件事：
 * 1. 将y的左子节点赋值给n的右子节点，同时将y的左子节点的父节点更新为n节点(y的左节点不空时)
 * 2. 将n的父节点p(非空时)变为y的父节点，同时更新p的子节点为y(左或右)
 * 3. 将y的左子节点设为n，将n的父节点设为y
*/
func (n *RBNode) leftRotate(root *RBNode) {
	right := n.right
	n.right = right.left
	if n.right != nil {
		n.right.parent = n
	}

	right.parent = n.parent
	if n.parent != nil {
		if n == n.parent.left {
			n.parent.left = right
		} else {
			n.parent.right = right
		}
	} else {
		root = right
	}

	right.left = n
	n.parent = right
}

/*************对红黑树节点n进行右旋操作 ******************/
/*
 * 右旋示意图：对节点n进行右旋
 *        p                   p
 *       /                   /
 *      n                   x
 *     / \                 / \
 *    x  rn   ----->      lx  n
 *   / \                     / \
 * lx  rx                   rx rn
*/
func (n *RBNode) rightRotate(root *RBNode) {
	left := n.left
	n.left = left.right
	if left.right != nil {
		left.right.parent = n
	}

	left.parent = n.parent
	if n.parent != nil {
		if n == n.parent.left {
			left.parent.left = left
		} else {
			left.parent.right = left
		}
	} else {
		root = left
	}

	left.right = n
	n.parent = left
}

func (n *RBNode) CompareWith(m *RBNode) int {
	if n.value > m.value {
		return 1
	} else if n.value < m.value {
		return -1
	}
	return 0
}

func (n *RBNode) IsRed() bool {
	return n.color
}

func (n *RBNode) SetBlack() {
	n.color = BLACK
}