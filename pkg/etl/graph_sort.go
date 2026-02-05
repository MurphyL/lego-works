package etl

import "errors"

/**
 * 拓扑排序
 */

type Graph[T comparable] struct {
	nodes map[T]bool
	edges map[T][]T
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{
		nodes: make(map[T]bool),
		edges: make(map[T][]T),
	}
}

func (g *Graph[T]) AddNode(node T) {
	g.nodes[node] = true
}

func (g *Graph[T]) AddEdge(from, to T) {
	g.edges[from] = append(g.edges[from], to)
}

func (g *Graph[T]) TopologicalSort() ([]T, error) {
	inDegree := make(map[T]int)

	// 初始化所有节点的入度
	for node := range g.nodes {
		inDegree[node] = 0
	}

	// 计算入度
	for _, neighbors := range g.edges {
		for _, neighbor := range neighbors {
			inDegree[neighbor]++
		}
	}

	// 收集入度为0的节点
	queue := make([]T, 0)
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	result := make([]T, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, neighbor := range g.edges[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(result) != len(g.nodes) {
		return nil, errors.New("graph has at least one cycle")
	}

	return result, nil
}

/**
1. 拓扑排序的结果不唯一，同一个 DAG 可能有多个有效的拓扑顺序；
2. 如果图中存在环，拓扑排序会失败（因为 DAG 必须是无环的）；
3. 对于大型图，基于 BFS 的 Kahn 算法通常比 DFS 实现更高效；
4. 实际应用中，可能需要根据具体需求调整实现，比如处理并发或添加权重；

func main() {
    g := NewGraph[string]()
    g.AddNode("A")
    g.AddNode("B")
    g.AddNode("C")
    g.AddNode("D")

    g.AddEdge("A", "B")
    g.AddEdge("A", "C")
    g.AddEdge("B", "D")
    g.AddEdge("C", "D")

    sorted, err := g.TopologicalSort()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Topological order:", sorted)
    // 可能的输出: [A B C D] 或 [A C B D]
}
**/
