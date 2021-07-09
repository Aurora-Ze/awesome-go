package day5

import (
	"strings"
)

// Node 定义了前缀路由树的节点
type Node struct {
	pattern      string  // 路径，除叶子节点外其他节点的pattern均为空
	curPart      string  // 路径的一部分，用来表示当前匹配的值
	children     []*Node // 子节点
	isFuzzyMatch bool    // 当前匹配的部分是否采用模糊匹配(即动态路由)的方式
}

// insert 从根节点到叶子，构造出一条路径
func (n *Node) insert(pattern string, parts []string, height int) {
	// 设置叶子节点的pattern，用于查询时确认是否匹配
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil { // 如果不存在，则插入一个子节点
		child = &Node{
			curPart:      part,
			isFuzzyMatch: part[0] == '*' || part[0] == ':',
		}
		n.children = append(n.children, child)
	}
	// 向下递归
	child.insert(pattern, parts, height+1)
}

// matchChild 遍历节点n的所有子节点，返回匹配该部分路径的第一个节点
func (n *Node) matchChild(curPart string) *Node {
	for _, node := range n.children {
		if node.curPart == curPart || node.isFuzzyMatch {
			return node
		}
	}
	return nil
}

// search DFS遍历路由树，查找匹配到的最终节点。如果没有节点匹配，则返回nil
func (n *Node) search(parts []string, height int) *Node {
	// 递归终止条件：来到叶子节点或遇到*号
	if len(parts) == height || strings.HasPrefix(n.curPart, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil

}

// matchChildren 遍历节点n的所有子节点，返回匹配的节点数组
func (n *Node) matchChildren(curPart string) []*Node {
	children := make([]*Node, 0)
	for _, node := range n.children {
		if node.curPart == curPart || node.isFuzzyMatch {
			children = append(children, node)
		}
	}
	return children
}
