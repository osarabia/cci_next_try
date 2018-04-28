package queue

import "graph"

type queue struct{
    array []*graph.Node
    index int
}

func (q *queue) Enqueue(n graph.Node){
    q.array = append(q.array, &n)
    q.index += 1
}

func (q *queue) Dequeue() *graph.Node{
    item := q.array[q.index]
    q.index -= 1

    return item
}

func (q *queue) Empty() bool{
    return q.index == -1
}

func New() *queue{
    return &queue{ make([]*graph.Node, 5 ,5), -1}
}
