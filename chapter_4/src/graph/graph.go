package graph

type Graph struct{
    adjacents map[string][]*Node
}

type Node struct {
    Vertex string
    Visited bool
}

func (g *Graph) AddAdjacent(n Node, a Node){
    _, ok := g.adjacents[n.Vertex]
    if !ok{
        g.adjacents[n.Vertex] = []*Node{&a}
        return
    }

    g.adjacents[n.Vertex] = append(g.adjacents[n.Vertex], &a)
}

func (g *Graph) GetAdjacents(node Node) []*Node{
    adjacents := g.adjacents[node.Vertex]

    return  adjacents
}

func New() *Graph{
    return &Graph{adjacents: make(map[string][]*Node)}
}
