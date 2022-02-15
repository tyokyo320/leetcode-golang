/*
 * @lc app=leetcode id=589 lang=golang
 *
 * [589] N-ary Tree Preorder Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {

	traversePath := []int{}

	// 数组是值类型 Go中的数组是值类型，而不是引用类型
	dfs(root, &traversePath)

	return traversePath
}

func dfs(root *Node, traversePath *[]int) {

	if root == nil {
		return
	}

	*traversePath = append(*traversePath, root.Val)

	// root.Children是根节点的（相同深度的）子节点
	for _, child := range root.Children {
		dfs(child, traversePath)
	}
}

// @lc code=end

