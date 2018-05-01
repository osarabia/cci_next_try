package node

import "fmt"

type Node struct{
    Value int
    Left *Node
    Right *Node
}

func (n *Node) String() string{
    return fmt.Sprintf("Left:%d, Value:%d, Right:%v", n.Left.Value, n.Value, n.Right.Value)
}
