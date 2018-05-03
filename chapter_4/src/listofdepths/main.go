package main

import (
    "fmt"
    "node"
    "linkedlist"
)

func main(){
    bt := createBT()
    depths := listOfDepths(bt)
    for key, ll := range depths{
        fmt.Println(key)
        node := ll.GetHead()
        for node != nil{
            fmt.Println(node)
            node = node.Next()
        }
    }
}

func listOfDepths(bt *node.Node) map[int]*linkedlist.List{
    depths := make(map[int]*linkedlist.List)
    i := 0

    array :=  []*node.Node{bt}

    for len(array) > 0 {
        array_children := make([]*node.Node, 0, 5)

        depths[i] = &linkedlist.List{} 
        ll := depths[i]
        for _, n := range array{
            ll.Insert(n)
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

    return depths
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
