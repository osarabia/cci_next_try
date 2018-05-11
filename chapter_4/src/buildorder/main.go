package main


import (
    "fmt"
    "graph"
)

func main(){
    fmt.Println("Just Starting")
    //Building Graph
    gph := graph.NewGraphHashLike(6)
    gph.AddEdge("a", "b")
    gph.AddEdge("f", "b")
    gph.AddEdge("b", "d")
    gph.AddEdge("f", "a")
    gph.AddEdge("d", "c")

    //util variables
    vertices := []string{"a", "b", "c", "d", "e", "f"}
    visited := map[string]bool{"a": false,"b": false, "c": false, "d": false, "e": false, "f": false}

    topologySort(vertices, visited, gph)
}

func topologySortHelper(stack *[]string, v string, visited map[string]bool, gph *graph.GraphHashLike){
    visited[v] = true

    for _, vv := range gph.Vertices[v]{
        if visited[vv] == false{
            topologySortHelper(stack, vv, visited, gph)
        }
    }

    *stack = append(*stack, v)
}

func topologySort(vertices []string, visited map[string]bool, gph *graph.GraphHashLike) {
    stack := make([]string, 0, gph.TotalVertex)

    for _,v := range vertices {
        if visited[v] == false{
            topologySortHelper(&stack, v, visited, gph)
        }
    }

    for i := len(stack)-1; i >= 0; i--{
        fmt.Print(stack[i])
    }
    fmt.Println("")
}
