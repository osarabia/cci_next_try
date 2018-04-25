import random

from linkedlist import DoublyLinkedList

class AnimalShelter(object):
    def __init__(self):
        self.cats_ll = DoublyLinkedList()
        self.dogs_ll = DoublyLinkedList()

    def enqueue(self, animal):
        if animal.endswith("_c"):
            self.cats_ll.add(animal)
        else:
            self.dogs_ll.add(animal)

    def dequeueAny(self):
        num = random.randint(0,1)
        if num == 0:
            return self.dequeueCat()

        return self.dequeueDog()

    def dequeueCat(self):
        cat = self.cats_ll.head
        self.cats_ll.head = cat.next

        return cat.value


    def dequeueDog(self):
        dog = self.dogs_ll.head
        self.dogs_ll.head = dog.next

        return dog.value


if __name__ == "__main__":
    a = AnimalShelter()
    a.enqueue("black_c")
    a.enqueue("minty_c")
    a.enqueue("bony")
    a.enqueue("dollar")
    assert a.dequeueCat() == "black_c"
    assert a.dequeueCat() == "minty_c"
    assert a.dequeueDog() == "bony"
    assert a.dequeueDog() == "dollar"
