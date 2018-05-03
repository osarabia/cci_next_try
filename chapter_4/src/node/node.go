package node

import "fmt"

type Node struct{
    Value interface{}
    Left *Node
    Right *Node
}

func (n *Node) ListChildren() (*Node,*Node){
    return n.Left, n.Right
}

func (n *Node) String() string{
    return fmt.Sprintf("Left:%d, Value:%d, Right:%v", n.Left.Value, n.Value, n.Right.Value)
}
