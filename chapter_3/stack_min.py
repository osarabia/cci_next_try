import sys

class IndexTracker(object):
    def __init__(self, start, end, current):
        self.start = start
        self.end = end
        self.current = current

    def __str__(self):
        return "{}-{}-{}".format(self.start,
                                 self.end,
                                 self.current)

class Stack(object):
    def __init__(self, stacksize):
        self.array = [0] * stacksize
        self.mins = [sys.maxint] * stacksize
        self.indexes = IndexTracker(0, stacksize-1, -1)

    def isFull(self):
        return self.indexes.end == self.indexes.current

    def isEmpty(self):
        return self.indexes.current == -1

    def min(self):
        if self.isEmpty():
            raise Exception("Stack is empty")

        return self.mins[self.indexes.current]

    def push(self, item):
        if self.isFull():
            raise Exception("Stack is full")

        self.indexes.current += 1
        if self.isEmpty():
            self.mins[self.indexes.current] = item
        else:
            self.mins[self.indexes.current] = min(item, self.mins[self.indexes.current -1])
        self.array[self.indexes.current] = item

    def pop(self):
        if self.isEmpty():
            raise Exception("Stack is empty")

        value = self.array[self.indexes.current]
        self.indexes.current -= 1

        return value

    def peek(self):
        if self.isEmpty():
            raise Exception("Stack is empty")

        return self.array[self.indexes.current]


if __name__ == "__main__":
    newstack = Stack(10)
    newstack.push(3)
    newstack.push(2)
    newstack.push(4)
    print("Min:", newstack.min())
    newstack.pop()
    newstack.push(1)
    print("Min:", newstack.min())
