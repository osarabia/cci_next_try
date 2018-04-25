import unittest

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
        self.indexes = IndexTracker(0, stacksize-1, -1)

    def isFull(self):
        return self.indexes.end == self.indexes.current

    def isEmpty(self):
        return self.indexes.current == -1

    def push(self, item):
        if self.isFull():
            raise Exception("Stack is full")

        self.indexes.current += 1
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

class SortedStack(object):
    def __init__(self, stacksize):
        self.stacksize = stacksize
        self.array = [0] * stacksize
        self.indexes = IndexTracker(0, stacksize-1, -1)

    def isFull(self):
        return self.indexes.end == self.indexes.current

    def isEmpty(self):
        return self.indexes.current == -1

    def push(self, item):
        if self.isFull():
            raise Exception("Stack is full")

        if self.isEmpty():
            self.indexes.current += 1
            self.array[self.indexes.current] = item
        else:
            tmpStack = Stack(self.stacksize)
            while not self.isEmpty() and  self.peek() < item:
                tmpStack.push(self.pop())

            self.indexes.current += 1
            self.array[self.indexes.current] = item

            while not tmpStack.isEmpty():
                value = tmpStack.pop()
                self.indexes.current += 1
                self.array[self.indexes.current] = value

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

class Tests(unittest.TestCase):
    def test_sorted_stack(self):
        ss = SortedStack(10)
        ss.push(8)
        ss.push(7)
        ss.push(6)
        ss.push(10)
        ss.push(4)
        ss.push(5)
        ss.push(20)
        ss.push(12)

        out = []
        for _ in range(8):
            out.append(ss.pop())
        self.assertEqual(out, [4, 5, 6, 7, 8, 10, 12, 20])


if __name__ == '__main__':
    unittest.main()
