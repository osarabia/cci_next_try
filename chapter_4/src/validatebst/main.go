package main

import (
    "fmt"
    "node"
)

func main(){
    fmt.Println("Starting")
    bt := createBT()
    leftSub := isLeftBalanced(bt)
    rightSub := isRightBalanced(bt)
    fmt.Println(leftSub)
    fmt.Println(rightSub)
}

func isRightBalanced(rootNode *node.Node) bool{
    var max, min, leftVal, rightVal,Val int
    min = rootNode.GetInt()

    Val = rootNode.GetInt()
    if rootNode.Right != nil {
       rightVal = rootNode.Right.GetInt()
       if rightVal > Val {
           max = rightVal
       } else {
           return false
       }
    }

    array := []*node.Node{rootNode.Right}

    for len(array) > 0 {
        array_children := make([]*node.Node, 0, 5)

        for idx, rN := range array{
             Val = rN.GetInt()
             // evaluating left
             if rN.Left != nil {
                 leftVal = rN.Left.GetInt()
                 if leftVal > Val || leftVal < min {
                     return false
                 }
                 array_children = append(array_children, rN.Left)
             }

             // evaluating Right
             if rN.Right != nil{
                 rightVal = rN.Right.GetInt()
                 if rightVal <= Val{
                     return false
                 }
                 if idx < len(array) - 1{
                     array_children = append(array_children, rN.Right)
                     continue
                 }
                 if idx == len(array) - 1 && rightVal > max{
                     max = rightVal
                     array_children = append(array_children, rN.Right)
                 } else {
                     return false
                 }
             }
        }

        array = array_children
    }

    return true

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

    array := []*node.Node{rootNode.Left}

    for len(array) > 0 {
        array_children := make([]*node.Node, 0, 5)

        for idx, rN := range array{
             Val = rN.GetInt()
             // evaluating left
             if rN.Left != nil {
                 leftVal = rN.Left.GetInt()
                 if leftVal > Val {
                     return false
                 }
                 if idx > 0 {
                     array_children = append(array_children, rN.Left)
                     continue
                 }
                 if idx == 0 && leftVal <  min {
                     min = leftVal
                     array_children = append(array_children, rN.Left)
                 } else {
                     return false
                 }
             }

             // evaluating Right
             if rN.Right != nil{
                 rightVal = rN.Right.GetInt()
                 if rightVal < Val || rightVal > max {
                     fmt.Println("right:", rightVal)
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
    root.Right = &node.Node{Value:7}
    left := root.Left
    left.Left = &node.Node{Value:1}
    left.Right = &node.Node{Value:4}

    right := root.Right
    right.Left = &node.Node{Value:6}
    right.Right = &node.Node{Value:21}

    return root
}
