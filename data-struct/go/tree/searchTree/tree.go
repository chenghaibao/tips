package main

import (
	"fmt"
	"math/rand"
	"time"
)

const ( //常量
	min, max = 0, 100
)

var (
	A    []int
	root *BTNode = &BTNode{}
)

type BTNode struct {
	m_value int
	m_left  *BTNode
	m_right *BTNode
	parent  *BTNode
}

func Init() bool { //初始化切片
	rand := rand.New(rand.NewSource(time.Now().UnixNano())) //时间种子
	for i := 0; i < 5; i++ {
		A = append(A, (rand.Intn(max-min) + min)) //生成随机数
	}
	return true
}

//插入一个节点到二叉树中
func TREE_INSERT(v int) { //递归
	y := &BTNode{}
	z := &BTNode{}
	z.m_value = v
	x := root

	if x != nil {
		for {
			y = x //y始终是x的父节点
			if v < x.m_value {
				x = x.m_left
			} else {
				x = x.m_right
			}
			if x == nil {
				break
			}
		}
	}

	z.parent = y

	if y.m_left == nil && y.m_right == nil && y.parent == nil { //树为空
		root = z
		//fmt.Println(root.m_value)
	} else {
		if v < y.m_value {
			y.m_left = z
		} else {
			y.m_right = z
		}
	}

}

func Create_Tree() { //插入一棵树
	length := len(A)
	for i := 0; i < length; i++ {
		TREE_INSERT(A[i])
	}
}

func TREE_SEARCH(root *BTNode, x int) *BTNode { //查询
	if root == nil || x == root.m_value {
		return root
	} else {
		if x < root.m_value {
			return TREE_SEARCH(root.m_left, x)
		} else {
			return TREE_SEARCH(root.m_right, x)
		}
	}
}

func INORDER_TREE_WALK(root *BTNode) { //中序打印一棵二叉树
	if root != nil {
		INORDER_TREE_WALK(root.m_left)
		fmt.Print(root.m_value, " ")
		INORDER_TREE_WALK(root.m_right)
	}
}

func TRANSPLANT(root, u, v *BTNode) {
	if u.parent == nil {
		root = v
	} else {
		if u == u.parent.m_left {
			u.parent.m_left = v
		} else {
			u.parent.m_right = v
		}
	}

	if v != nil {
		v.parent = u.parent
	}
}

func TREE_MINIMUM(root *BTNode) *BTNode {
	if root.m_left != nil {

		for { //由于Go中没有while语句
			root = root.m_left
			if root.m_left == nil {
				break
			}
		}
	}
	return root
}

func TREE_DELETE(z *BTNode) {
	y := &BTNode{}
	if z.m_left == nil {
		TRANSPLANT(root, z, z.m_right)
	} else {
		if z.m_right == nil {
			TRANSPLANT(root, z, z.m_left)
		} else { //z有左右两个子树
			y = TREE_MINIMUM(root.m_right)
			if y.parent != z {
				TRANSPLANT(root, y, y.m_right)
				y.m_right = z.m_right
				y.m_right.parent = y
			}
			TRANSPLANT(root, z, y)
			y.m_left = z.m_left
			y.m_left.parent = y
		}
	}
}

func main() {
	Init()
	fmt.Println(A)
	t1 := time.Now()
	Create_Tree()
	INORDER_TREE_WALK(root)
	fmt.Println()
	//删除操作
	s := TREE_SEARCH(root, A[2])
	TREE_DELETE(s)
	INORDER_TREE_WALK(root)
	elapsed := time.Since(t1)
	fmt.Println("运行时间: ", elapsed)

}
