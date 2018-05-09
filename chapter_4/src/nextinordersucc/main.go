package main

import (
    "fmt"
    "node"
)

func main(){
    fmt.Println("Starting")

    m := insert(nil, 20)
    m = insert(m, 8)
    m = insert(m, 22)
    m = insert(m, 4)
    m = insert(m, 12)
    m = insert(m, 10)
    m = insert(m, 14)

    successor := nextInOrder(m.Left)
    fmt.Println(successor.GetInt()) //should Print 10

    successor = nextInOrder(m.Left.Right.Left)
    fmt.Println(successor.GetInt())//should Print 12

    successor = nextInOrder(m.Left.Right.Right)
    fmt.Println(successor.GetInt())//should Print 20
}

func nextInOrder(n *node.Node) *node.Node {
    //If right subtree of node is not NULL, then succ lies in right subtree
    if n.Right != nil {
        return leftMost(n.Right)
    }

    //Travel up using the parent pointer
    p := n.Parent
    for p != nil {

        //until you see a node which is left child of it’s parent. The parent of such a node is the succ.
        if n.GetInt() != p.Right.GetInt(){
            break
        }
        n = p
        p = p.Parent
    }

    return p
}

//looking for the left most :)
func leftMost(n *node.Node) *node.Node {
    current := n

    for current != nil {
        if current.Left == nil {
            break
        }

        current = current.Left
    }

    return current
}

func insert(n *node.Node, value int) *node.Node {
    if n == nil{
        return &node.Node{Value: value}
    }

    if value <= n.GetInt() {
        tmp := insert(n.Left, value)
        n.Left = tmp
        tmp.Parent = n
    } else {
        tmp := insert(n.Right, value)
        n.Right = tmp
        tmp.Parent = n
    }

    return n
}
