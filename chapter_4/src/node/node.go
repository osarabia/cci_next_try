package node

import "fmt"

type Node struct{
    Value interface{}
    Parent *Node
    Left *Node
    Right *Node
}

func (n *Node) ListChildren() (*Node,*Node){
    return n.Left, n.Right
}

func (n *Node) String() string{
    return fmt.Sprintf("Left:%v, Value:%v, Right:%v", n.Left.Value, n.Value, n.Right.Value)
}

func (n *Node) GetInt() int{
    val, _ := n.Value.(int)

    return val
}
