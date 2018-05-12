package main

import (
    "fmt"
    "node"
)

func main(){
    fmt.Println("Starting")
    tree := createBT()

    n :=  commonAncestor(tree, tree.Left, tree.Right)
    fmt.Println(n.GetInt()) //5
    n = commonAncestor(tree, tree.Right.Right, tree.Right.Left)
    fmt.Println(n.GetInt()) //7
}

func commonAncestor(root *node.Node, node1 *node.Node, node2 *node.Node) *node.Node{
    if root == nil{
        return nil
    }

    if root.GetInt() == node1.GetInt() || root.GetInt() == node2.GetInt(){
        return root
    }

    left := commonAncestor(root.Left, node1, node2)
    right := commonAncestor(root.Right, node1, node2)

    if left != nil && right != nil {
        return root
    }

    if left == nil && right == nil{
        return nil
    }

    if left != nil{
        return left
    }

    if right != nil {
        return right
    }

    return nil
}

func createBT() *node.Node{
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
