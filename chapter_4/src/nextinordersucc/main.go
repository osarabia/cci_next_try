package main

import (
    "fmt"
    "node"
)

func main(){
    fmt.Println("Starting")

    m := insert(nil, 20)
    m = insert(m.Node, 8)
    m = insert(m.Node, 22)
    m = insert(m.Node, 4)
    m = insert(m.Node, 12)
    m = insert(m.Node, 10)
    m = insert(m.Node, 10)
    fmt.Println(m.Value)
    fmt.Println(m.Left.Value)
    fmt.Println(m.Right.Value)
}

func insert(n *node.Node, value int) *node.NodeInOrderSucc {
    if n == nil{
        tmp := &node.Node{Value: value}

        return &node.NodeInOrderSucc{Node: tmp}
    }

    if value <= n.GetInt() {
        tmp := insert(n.Left, value)
        n.Left = tmp.Node
        tmp.Parent = n
    } else {
        tmp := insert(n.Right, value)
        n.Right = tmp.Node
        tmp.Parent = n
    }

    return &node.NodeInOrderSucc{Node: n}
}
