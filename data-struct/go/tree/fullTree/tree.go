package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

//二叉树深度
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	lhight := maxDepth(root.left)
	rhight := maxDepth(root.right)
	return max(lhight, rhight)
}
func max(a, b int) int {
	if a > b {
		return a + 1
	}
	return b + 1
}

//判断
func isFull(root *Node) bool {
	if root == nil {
		return true
	}
	lheight := maxDepth(root.left)
	rheight := maxDepth(root.right)
	fmt.Println("l:", lheight, "r:", rheight)
	return isFull(root.left) && isFull(root.right) && (lheight == rheight)

	//完全二叉树
	//if math.Abs(float64(Height(root.left)-Height(root.right))) ==1  {
	//	flag = true
	//} else {
	//	flag = false
	//}
	//return flag && isFull(root.left) && isFull(root.right)
}

//2.最小深度
func minDepth(root *Node) int {
	if root == nil {
		return 0
	}
	lhight := minDepth(root.left)
	rhight := minDepth(root.right)
	if root.left == nil {
		return lhight + 1
	} else if root.right == nil {
		return rhight + 1
	} else {
		return min(lhight, rhight) + 1
	}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 中序遍历 左 -> 中 -> 右。
func MidOrderTraversal(tree *Node) {
	if tree == nil {
		return
	}

	MidOrderTraversal(tree.left)
	fmt.Printf(" %d -> ", tree.val)
	MidOrderTraversal(tree.right)
}

// 后序遍历   左 -> 右 -> 中。
func PostOrderTraversal(tree *Node) {
	if tree == nil {
		return
	}

	PostOrderTraversal(tree.left)
	PostOrderTraversal(tree.right)
	fmt.Printf(" %d -> ", tree.val)
}

// 前序遍历 遍历顺序
//  中 -> 左 -> 右。
func PreOrderTraversal(tree *Node) {
	if tree == nil {
		return
	}
	fmt.Printf(" %d -> ", tree.val)
	PreOrderTraversal(tree.left)
	PreOrderTraversal(tree.right)
}

// 按层遍历
func LevelOrderTraversal(tree *Node) {
	if tree == nil {
		return
	}

	// 采用队列实现
	queue := make([]*Node, 0)
	queue = append(queue, tree) // queue push
	for len(queue) > 0 {
		tree = queue[0]
		fmt.Printf(" %d -> ", tree.val)

		queue = queue[1:] // queue pop

		if tree.left != nil {
			queue = append(queue, tree.left)
		}
		if tree.right != nil {
			queue = append(queue, tree.right)
		}
	}
}

func main() {
	root := &Node{1, nil, nil}
	n1 := &Node{2, nil, nil}
	n2 := &Node{3, nil, nil}
	n3 := &Node{4, nil, nil}
	n4 := &Node{5, nil, nil}
	//n5 := &Node{6, nil, nil}
	//n6 := &Node{7, nil, nil}

	root.left = n1
	root.right = n2
	n1.left = n3
	n1.right = n4
	//n2.left = n5
	//n2.right = n6
	fmt.Println("前------------")
	PreOrderTraversal(root)
	fmt.Println()
	fmt.Println("中------------")
	MidOrderTraversal(root)
	fmt.Println()
	fmt.Println("后------------")
	PostOrderTraversal(root)
	fmt.Println()
	fmt.Println("层------------")
	LevelOrderTraversal(root)
	fmt.Println("------------")
	fmt.Println(isFull(root))
}
