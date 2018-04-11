from linkedlist import LinkedList

def partition(ll, partitionVal):
    current = ll.tail = ll.head

    while current is not None:
        nextNode = current.next
        if current.value < partitionVal:
            current.next = ll.head
            ll.head = current
        else:
            ll.tail.next = current
            ll.tail = current
        current = nextNode

    if ll.tail.next is not None:
        ll.tail.next = None

if __name__ == "__main__":
    ll = LinkedList()
    ll.generate(10, 0, 9)
    print(ll)
    partition(ll, 5)
    print(ll)
