package linkedlist

type Node struct{
    prev *Node
    next *Node
    value interface{}
}

type List struct{
    head *Node
    tail *Node
}

func (L *List) Insert(value interface{}){
    node := &Node{
        value: value
    }

    if L.head != nil{
        L.tail.next = node
        L.tail = L.tail.next
    } else {
        L.head = L.tail = node
    }
}
