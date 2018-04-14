from linkedlist import LinkedList


def intersection(ll1, ll2):
    if ll1.tail is not ll2.tail:
        return

    long_ll = ll1 if len(ll1) > len(ll2) else ll2
    short_ll = ll2 if len(ll1) > len(ll2) else ll1


    diff = long_ll - short_ll

    long_node, short_node = long_ll.head, short_ll.head

    for _ in range(diff):
        long_node = long_node.next

    while long_node is not short_node:
        long_node = long_node.next
        short_node = short_node.next

    return short_node
