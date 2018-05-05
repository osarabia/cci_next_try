package main

import (
    "fmt"
    "node"
)

func main(){
    fmt.Println("Starting")
    bt := createBT()
    leftSub := isLeftBalanced(bt)
    fmt.Println(leftSub)
}

func isLeftBalanced(rootNode *node.Node) bool{
    var max, min, leftVal, rightVal,Val int
    max = rootNode.GetInt()

    Val = rootNode.GetInt()
    if rootNode.Left != nil {
       leftVal = rootNode.Left.GetInt()
       if leftVal <= Val {
           min = leftVal
       } else {
           return false
       }
    }

    array_children := make([]*node.Node, 0, 5)
    array := []*node.Node{rootNode.Left}

    for len(array) > 0 {
        for idx, rN := range array{
             Val = rN.GetInt()
             // evaluating left
             if rN.Left != nil {
                 leftVal = rN.Left.GetInt()
                 if idx == 0 && leftVal <= Val && leftVal <  min {
                     min = leftVal
                 } else if idx > 0 && leftVal <= Val {
                     fmt.Println("Si")
                 }else{
                     fmt.Println(leftVal)
                     return false
                 }
                 array_children = append(array_children, rN.Left)
             }

             // evaluating Right
             if rN.Right != nil{
                 rightVal = rN.Right.GetInt()
                 if rightVal < Val || rightVal > max {
                     fmt.Println(rightVal)
                     return false
                 }
                 array_children = append(array_children, rN.Right)
             }
        }

        array = array_children
    }

    return true
}

func createBT() *node.Node{
    root := &node.Node{Value: 5}
    root.Left = &node.Node{Value:3}
    root.Right = &node.Node{Value:6}
    left := root.Left
    left.Left = &node.Node{Value:1}
    left.Right = &node.Node{Value:2}

    right := root.Right
    right.Left = &node.Node{Value:20}
    right.Right = &node.Node{Value:21}

    return root
}
