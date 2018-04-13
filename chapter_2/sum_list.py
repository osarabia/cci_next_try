from linkedlist import LinkedList

def sum_list(ll1, ll2):
    remaining = 0
    current1 = ll1.head
    current2 = ll2.head
    ll = LinkedList()

    while current1 and current2:
        number = current1.value + current2.value + remaining
        remaining = number // 10
        ll.add(number % 10)

        current1 = current1.next
        current2 = current2.next

    while current1:
        number = current1.value + remaining
        remaining = number // 10
        ll.add(number % 10)
        current1 = current1.next

    while current2:
        number = current2.value + remaining
        remaining = number // 10
        ll.add(number % 10)
        current2 = current2.next

    if remaining > 0:
        ll.add(remaining)

    return ll

if __name__ == "__main__":
    ll1 = LinkedList()
    ll1.add_multiple([7,1,6])
    ll2 = LinkedList()
    ll2.add_multiple([5,9,2])
    new_ll = sum_list(ll1, ll2)
    print(new_ll)
