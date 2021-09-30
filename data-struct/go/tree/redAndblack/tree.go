package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	RED   bool = true
	BLACK bool = false
)

type Node struct {
	key   int
	value interface{}

	left  *Node
	right *Node
	//parent *Node

	color bool
}

type RedBlackTree struct {
	size int
	root *Node
}

func NewNode(key, val int) *Node {
	// 默认添加红节点
	return &Node{
		key:   key,
		value: val,
		left:  nil,
		right: nil,
		//parent: nil,
		color: RED,
	}
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

func (n *Node) IsRed() bool {
	if n == nil {
		return BLACK
	}
	return n.color
}

func (tree *RedBlackTree) GetTreeSize() int {
	return tree.size
}

//   node                     x
//  /   \     左旋转         /  \
// T1   x   --------->   node   T3
//     / \              /   \
//    T2 T3            T1   T2
func (n *Node) leftRotate() *Node {
	// 左旋转
	retNode := n.right
	n.right = retNode.left

	retNode.left = n
	retNode.color = n.color
	n.color = RED

	return retNode
}

//     node                    x
//    /   \     右旋转       /  \
//   x    T2   ------->   y   node
//  / \                       /  \
// y  T1                     T1  T2
func (n *Node) rightRotate() *Node {
	//右旋转
	retNode := n.left
	n.left = retNode.right

	retNode.right = n
	retNode.color = n.color
	n.color = RED

	return retNode
}

// 颜色变换
func (n *Node) flipColors() {
	n.color = RED
	n.left.color = BLACK
	n.right.color = BLACK
}

// 维护红黑树
func (n *Node) updateRedBlackTree(isAdd int) *Node {
	// isAdd=0 说明没有新节点，无需维护
	if isAdd == 0 {
		return n
	}

	// 需要维护
	if n.right.IsRed() == RED && n.left.IsRed() != RED {
		n = n.leftRotate()
	}

	// 判断是否为情形3，是需要右旋转
	if n.left.IsRed() == RED && n.left.left.IsRed() == RED {
		n = n.rightRotate()
	}

	// 判断是否为情形4，是需要颜色翻转
	if n.left.IsRed() == RED && n.right.IsRed() == RED {
		n.flipColors()
	}

	return n
}

// 递归写法:向树的root节点中插入key,val
// 返回1, 代表加了节点
// 返回0, 代表没有添加新节点, 只更新key对应的value值
func (n *Node) add(key, val int) (int, *Node) {
	if n == nil { // 默认插入红色节点
		return 1, NewNode(key, val)
	}

	isAdd := 0
	if key < n.key {
		isAdd, n.left = n.left.add(key, val)
	} else if key > n.key {
		isAdd, n.right = n.right.add(key, val)
	} else {
		// 对value值更新,节点数量不增加,isAdd = 0
		n.value = val
	}

	// 维护红黑树
	n = n.updateRedBlackTree(isAdd)

	return isAdd, n
}

func (tree *RedBlackTree) Add(key, val int) {
	isAdd, nd := tree.root.add(key, val)
	tree.size += isAdd
	tree.root = nd
	tree.root.color = BLACK //根节点为黑色节点
}

// 前序遍历打印出key,val,color
func (tree *RedBlackTree) PrintPreOrder() {
	resp := make([][]interface{}, 0)
	tree.root.printPreOrder(&resp)
	fmt.Println(resp)
}

func (n *Node) printPreOrder(resp *[][]interface{}) {
	if n == nil {
		return
	}
	*resp = append(*resp, []interface{}{n.key, n.value, n.color})
	n.left.printPreOrder(resp)
	n.right.printPreOrder(resp)
}

// 测试红黑树代码
func main() {
	count := 10
	redBlackTree := NewRedBlackTree()
	nums := make([]int, 0)
	for i := 0; i < count; i++ {
		nums = append(nums, rand.Intn(count))
	}

	fmt.Println("source data: ", nums)
	now := time.Now()
	for _, v := range nums {
		redBlackTree.Add(v, v)
	}

	fmt.Println("redBlackTree:", now.Sub(time.Now()))
	redBlackTree.PrintPreOrder()
	fmt.Println("节点数量:", redBlackTree.GetTreeSize())
}
