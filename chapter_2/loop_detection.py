from linkedlist import LinkedList

def loop_detection(ll):
    start = forward = ll.head

    while forward and forward.next:
        forward = forward.next.next
        start = start.next
        if start is forward:
            break

    #check by no loop
    if forward is None or forward.next is None:
        return

    start = ll.head
    while start is not forward:
        start = start.next
        forward = forward.next

    return forward
