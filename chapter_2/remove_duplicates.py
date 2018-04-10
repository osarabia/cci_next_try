from linkedlist import LinkedList


def remove_duplicates(ll):

    if ll.head is None:
        return

    current = ll.head
    values = set([current.value])

    while current.next is not None:
        if current.next.value in values:
            current.next = current.next.next
        else:
            values.add(current.next.value)
            current = current.next

    return ll

if __name__ == "__main__":
    ll = LinkedList()
    ll.generate(100, 0, 9)
    print(ll)
    remove_duplicates(ll)
    print("------")
    print(ll)
