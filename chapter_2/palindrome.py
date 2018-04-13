from linkedlist import LinkedList

def palindrome(ll):
    new_ll = LinkedList()

    node = ll.head
    while node:
        new_ll.add_to_beginning(node.value)
        node = node.next

    node = ll.head
    node1 = new_ll.head
    while node and node1:
        if node.value != node1.value:
            return False
        node = node.next
        node1 = node1.next

    return True

if __name__ == "__main__":
    ll = LinkedList()
    ll.add_multiple([7,1,6,1,7])

    ll1 = LinkedList()
    ll1.add_multiple([7,1,6])
    assert palindrome(ll) == True
    assert palindrome(ll1) == False
