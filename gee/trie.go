package gee

import "strings"

//实现路由模糊匹配的前缀树

type node struct {
	pattern  string  //待匹配的路由/user/login这种
	part     string  //当前节点的路由地址login这种
	children []*node //子节点
	isWild   bool    //是否是模糊匹配,part 含有 : 或 * 时为true
}

//第一个匹配的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

//所有匹配的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern != "" {
			return n
		} else {
			return nil
		}
	}

	part := parts[height]
	nodes := n.matchChildren(part)
	for _, nodes := range nodes {
		res := nodes.search(parts, height+1)
		if res != nil {
			return res
		}
	}
	return nil
}
