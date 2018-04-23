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


class SetOfStacks(object):

    def __init__(self, stacksize):
        self.stacksize = stacksize
        self.stacks = {}
        self.stackId = 1
        self.stacks[self.stackId] = Stack(stacksize)

    def createNewStack(self):
        self.stackId += 1
        self.stacks[self.stackId] = Stack(self.stacksize)

    def push(self, item):
        stack = self.stacks[self.stackId]
        if stack.isFull():
            self.createNewStack()

        stack = self.stacks[self.stackId]
        stack.push(item)

    def pop(self):
        stack = self.stacks[self.stackId]
        if stack.isEmpty():
            if self.stackId == 1:
                raise Exception("Empty")
            self.stackId -= 1
        stack = self.stacks[self.stackId]

        return stack.pop()

    def popAt(self, stackIndex):
        stack = self.stacks[stackIndex]
        if stack.isEmpty():
            raise Exception("Stack is Empty")

        return stack.pop()

class Tests(unittest.TestCase):
    def test_stacks(self):
        stacks = SetOfStacks(5)
        for i in range(35):
            stacks.push(i)
        lst = []
        for _ in range(35):
            lst.append(stacks.pop())
        self.assertEqual(lst, list(reversed(range(35))))

    def test_pop_at(self):
        stacks = SetOfStacks(5)
        for i in range(35):
            stacks.push(i)
        value1 = stacks.popAt(1)
        value2 = stacks.popAt(7)
        self.assertEqual(4, value1)
        self.assertEqual(34, value2)

if __name__ == '__main__':
    unittest.main()
