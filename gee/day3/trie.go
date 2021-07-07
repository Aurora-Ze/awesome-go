package day3

// Node 定义了前缀路由树的节点
type Node struct {
	pattern        string  // 路径，除叶子节点外其他节点的pattern均为空
	curPart        string  // 路径的一部分，用来表示当前匹配的值
	children       []*Node // 子节点
	isFuzzyMatch   bool    // 当前匹配的部分是否采用模糊匹配(即动态路由)的方式
}

// search DFS遍历路由树，查找匹配到的节点
func (n *Node) search(parts []string, height int)  {

}
