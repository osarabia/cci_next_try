package linkedlist

import "fmt"

type Node struct{
    prev *Node
    next *Node
    value interface{}
}

type List struct{
    head *Node
    tail *Node
}

func (N *Node) String() string{
    return fmt.Sprintf("%d", N.value)
}
func (N *Node) Next() *Node{
    return N.next
}

func (L *List) Insert(value interface{}){
    node := &Node{value: value}

    if L.head != nil{
        L.tail.next = node
        L.tail = L.tail.next
    } else {
        L.head = node
        L.tail = node
    }
}

func (L *List) GetHead() *Node{
    return L.head
}
