package main

import (
    "fmt"
    "graph"
    "queue"
)


func main(){
    fmt.Println("Starting")
    g := createGraph()

    nodes := g.GetNodes()

    start := nodes[0]
    end := nodes[5]


    bf := breadthFirst(g, start, end)
    fmt.Println(bf)
}

func breadthFirst(g *graph.Graph,
                    start *graph.Node,
                    end *graph.Node) bool {
    if start.Vertex == end.Vertex {
        return true
    }

    start.Visited = true
    aqueue := queue.New()
    aqueue.Enqueue(start)

    for !aqueue.Empty() {
        n := aqueue.Dequeue()
        adjacents := g.GetAdjacents(n)
        for _, a := range adjacents{
            if a.Visited == false {
                if a.Vertex == end.Vertex {
                    return true
                } else {
                    aqueue.Enqueue(a)
                }
                a.Visited = true
            }
        }
    }

    return false
}

func createGraph() *graph.Graph {
    g := graph.New()

    n1 := graph.Node{Vertex: "a"}
    n2 := graph.Node{Vertex: "b"}
    n3 := graph.Node{Vertex: "c"}
    n4 := graph.Node{Vertex: "d"}
    n5 := graph.Node{Vertex: "e"}
    n6 := graph.Node{Vertex: "f"}

    g.AddNode(n1)
    g.AddNode(n2)
    g.AddNode(n3)
    g.AddNode(n4)
    g.AddNode(n5)
    g.AddNode(n6)

    g.AddAdjacent(n1, n2)
    g.AddAdjacent(n1, n3)
    g.AddAdjacent(n1, n4)
    g.AddAdjacent(n4, n5)
    g.AddAdjacent(n5, n6)

    return g
}
