package leetcode200

import (
	"fmt"
	"testing"
)
/**
	给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

	岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

	此外，你可以假设该网格的四条边均被水包围。
 */

// 文字讲解
// 有思路但是写的还是有问题  必须看讲解才能写出来   太菜了太菜了 要努力要努力
// 可以当作 图来想 就是无向图  1 1 有边相连  由此得出 1 1只是一块陆地  第二个为水
// 不用想，这种题就是dfs大法
// dfs 模版  一直走，直到找到解或者走不下去后

var grid [][]byte

func init() {
	grid = [][]byte{[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011")}
}
func Test(t *testing.T) {
	fmt.Println(numIslands(grid))
}
func numIslands(grid [][]byte) int {
	// 行
	row := len(grid)
	// 列
	col := len(grid[0])
	// 岛屿数量
	count := 0
	// 遍历行
	for i := 0; i < row; i++ {
		//  遍历列
		for j := 0; j < col; j++ {

			fmt.Println(grid[i][j])
			// 不是水
			if  grid[i][j]!= '0' {
				// 那么必存在一个陆地
				count++
				// dfs求是否还有其他陆地
				DFS(i, j, grid)
			}
		}
	}
	return count
}

func DFS(i, j int, grid [][]byte) {
	// 很简单 等于0 水，大于行，大于列 都是水   找不下去了
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	// 每一个我都去找  找过的1   相邻陆地可以当作水来处理
	grid[i][j] = '0'
	// 左面
	DFS(i-1, j, grid)
	// 右面
	DFS(i+1, j, grid)
	// 上面
	DFS(i, j-1, grid)
	// 下面
	DFS(i, j+1, grid)
}
