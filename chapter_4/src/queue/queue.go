package queue

import (
    "graph"
)

type queue struct{
    array []*graph.Node
    start int
    end int
}

func (q *queue) Enqueue(n *graph.Node){
    q.array = append(q.array, n)
    q.end += 1
}

func (q *queue) Dequeue() *graph.Node{
    n := q.array[q.start]
    q.start += 1

    if q.start > 0 && q.start > q.end {
        q.array = make([]*graph.Node, 0, 5)
        q.start = 0
        q.end = -1
    }

    return n
}

func (q *queue) Empty() bool{
    return q.end == -1
}

func New() *queue{
    return &queue{ make([]*graph.Node, 0, 5), 0, -1}
}
