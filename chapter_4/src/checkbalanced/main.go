package main

import (
    "fmt"
    "node"
)

func main(){
    bt := createBT()
    fmt.Println(bt)
    leftdepth := levelsDepth(bt.Left)
    rightdepth := levelsDepth(bt.Right)
    diff := leftdepth - rightdepth
    if diff < 0 {
        diff  = diff * -1
    }

    if diff <= 1 {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}

func levelsDepth(rootNode *node.Node) int{
    array := make([]*node.Node, 0, 5)
    i := 0

    if rootNode != nil{
        array = append(array, rootNode)
        i += 1
    } else {
        return 0
    }

    for len(array) > 0 {
        array_children := make([]*node.Node, 0, 5)

        for _, n := range array{
            left, right := n.ListChildren()
            if left != nil{
              array_children = append(array_children, left)
            }

            if right != nil {
              array_children = append(array_children, right)
            }
        }

        i += 1
        array = array_children

    }

    return i
}

func createBT() *node.Node{
    root := &node.Node{Value: 5}
    root.Left = &node.Node{Value:3}
    root.Right = &node.Node{Value:6}
    left := root.Left
    left.Left = &node.Node{Value:2}
    left.Right = &node.Node{Value:1}

    right := root.Right
    right.Left = &node.Node{Value:20}
    right.Right = &node.Node{Value:21}

    return root
}
