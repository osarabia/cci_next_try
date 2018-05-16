package main

import (
    "fmt"
)

func main(){
    fmt.Println("Starting")

    root := createBST()
    results := findSequences(root)
}

func createBST() *node.Node{
    root := &node.Node{Value: 5}
    root.Left = &node.Node{Value:3}
    root.Right = &node.Node{Value:7}
    left := root.Left
    left.Left = &node.Node{Value:1}
    left.Right = &node.Node{Value:4}

    right := root.Right
    right.Left = &node.Node{Value:6}
    right.Right = &node.Node{Value:21}

    return root
}

func findSequences(n *node.Node) []int{
    if n == nil {
        var empty []int

        return  empty
    }

    //Leaf Nodes
    if n.Left && n.Right {
        return []int{n.GetInt()}
    }

    left := findSequences(n.Left)
    right := findSequences(n.Right)

    //merging everything
    return []int{0}
}
