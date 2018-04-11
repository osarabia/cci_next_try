from linkedlist import LinkedList


def delete_mid_node(node):
    node.value = node.next.value
    node.next = node.next.next


if __name__ == "__main__":
    ll = LinkedList()
    ll.add_multiple([10,0,9])
    m_node = ll.add(3)
    ll.add_multiple([4,5,6])
    print(ll)
    delete_mid_node(m_node)
    print(ll)


