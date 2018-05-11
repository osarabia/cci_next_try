package graph

type Graph struct{
    vertices  []*Node
    adjacents map[string][]*Node
}

type GraphHashLike struct{
    Vertices map[string][]string
    TotalVertex int
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

func (g *GraphHashLike) AddEdge(k string, v string){
    s, ok := g.Vertices[k]
    if ok {
        g.Vertices[k] = append(s, v)
        return
    }

    g.Vertices[k] = make([]string, 0, g.TotalVertex)
    g.Vertices[k] = append(g.Vertices[k], v)
}

func New() *Graph{
    return &Graph{adjacents: make(map[string][]*Node)}
}

func NewGraphHashLike(vertex int) *GraphHashLike{
    return &GraphHashLike{Vertices: make(map[string][]string), TotalVertex: vertex}
}
