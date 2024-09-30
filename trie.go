package gwh

import "strings"

type node struct {
    pattern  string           // 完整路由，只有叶子节点才有
    part     string           // 当前节点代表的路由部分
    children map[string]*node // 子节点
    isWild   bool             // 是否是通配符节点（: 或 * 开头）
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
    // 优先精准匹配
    if child, exists := n.children[part]; exists {
        return child
    }
    for _, child := range n.children {
        if child.isWild {
            return child
        }
    }
    return nil
}

// 所有匹配成功的节点，用于查找
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
        child = &node{
            part:   part,
            isWild: part[0] == ':' || part[0] == '*',
        }
        if n.children == nil {
            n.children = make(map[string]*node)
        }
        n.children[part] = child
    }
    child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
    if len(parts) == height || strings.HasPrefix(n.part, "*") {
        if n.pattern == "" {
            return nil
        }
        return n
    }
    children := n.matchChildren(parts[height])
    for _, child := range children {
        result := child.search(parts, height+1)
        if result != nil {
            return result
        }
    }
    return nil
}

// 遍历整棵 Trie，将所有有 pattern 的节点（即完整路由节点）添加到 list 中
func (n *node) travel(list *[]*node) {
    if n.pattern != "" {
        *list = append(*list, n)
    }
    for _, child := range n.children {
        child.travel(list)
    }
}
