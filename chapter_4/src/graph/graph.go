package graph

type Graph struct{
    vertices  []*Node
    adjacents map[string][]*Node
}

type Node struct {
    Vertex string
    Visited bool
}

func (n *Node) GetVertex() string {
    return n.Vertex
}

func (g *Graph) AddAdjacent(n Node, a Node){
    _, ok := g.adjacents[n.Vertex]
    if !ok{
        g.adjacents[n.Vertex] = []*Node{&a}
        return
    }

    g.adjacents[n.Vertex] = append(g.adjacents[n.Vertex], &a)
}

func (g *Graph) GetAdjacents(node *Node) []*Node{
    vertex := node.GetVertex()
    return g.adjacents[vertex]
}

func (g *Graph) AddNode(node Node) {
    g.vertices = append(g.vertices, &node)
}

func (g *Graph) GetNodes() []*Node{
    return g.vertices
}

func New() *Graph{
    return &Graph{adjacents: make(map[string][]*Node)}
}
