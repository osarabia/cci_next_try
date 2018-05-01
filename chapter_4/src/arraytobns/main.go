package main

import (
    "fmt"
    "node"
)

func main(){
    array := []int{1,2,3,4,5,6,7,8,9,10,20,22}
    arraytobns(array, 0, len(array) - 1)
}

func arraytobns(array []int, start int, end int) *node.Node{
    if start > end {
        return &node.Node{}
    }
    mid := (start + end) / 2
    root := &node.Node{Value: array[mid]}
    root.Left = arraytobns(array, start, mid - 1)
    root.Right = arraytobns(array, mid + 1, end)
    fmt.Println(root)
    return root
}
