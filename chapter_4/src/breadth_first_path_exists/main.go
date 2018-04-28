package main

import (
    "fmt"
    "graph"
    "queue"
)


func main(){
    g := graph.New()

    n1 := graph.Node{Vertex: "a"}
    n2 := graph.Node{Vertex: "b"}
    n3 := graph.Node{Vertex: "c"}
    n4 := graph.Node{Vertex: "d"}
    n5 := graph.Node{Vertex: "e"}
    n6 := graph.Node{Vertex: "f"}

    g.AddAdjacent(n1, n2)
    g.AddAdjacent(n1, n3)
    g.AddAdjacent(n1, n4)
    g.AddAdjacent(n4, n5)
    g.AddAdjacent(n5, n6)

    q := queue.New()

    fmt.Println(q)
}
