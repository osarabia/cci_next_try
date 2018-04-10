from linkedlist import LinkedList

def kth_to_last(kth, ll):
    if ll.head is None:
        return

    back_node = ll.head
    front_node = ll.head

    i = 0

    while i < kth:
        if front_node is None:
            return

        front_node = front_node.next
        i += 1

    if front_node is None:
        return

    while front_node.next is not None:
        back_node = back_node.next
        front_node = front_node.next

    return back_node

if __name__ == "__main__":
    ll = LinkedList()
    ll.generate(10, 0, 9)
    print(ll)
    node = kth_to_last(9, ll)
    print(node)
