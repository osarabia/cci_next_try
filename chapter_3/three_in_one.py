class Stack(object):
    def __init__(self, start, end, current):
        self.start_index = start
        self.end_index = end
        self.current_index = current

    def __str__(self):
        return "{}-{}-{}".format(self.start_index,
                                 self.end_index,
                                 self.current_index)

class ThreeInOne(object):
    def __init__(self, stacksize):
        self.stacksize = stacksize
        self.size = stacksize * 3
        self.array = [0] * (self.size)
        self.stacks = {}

        i = 1
        for x in xrange(0, self.size, self.stacksize):
            self.stacks[i] = Stack(x, x + self.stacksize -1, -1)
            #print(self.stacks[i])
            i += 1

    def isFull(self, stacknum):
        stack = self.stacks[stacknum]

        return stack.end_index == stack.current_index

    def isEmpty(self, stacknum):
        stack = self.stacks[stacknum]

        return stack.current_index == -1

    def push(self, item, stacknum):
        stack = self.stacks[stacknum]
        if self.isFull(stacknum):
            raise Exception("Stack {} is full".format(stacknum))

        stack.current_index += 1
        self.array[stack.current_index] = item

    def pop(self, stacknum):
        stack = self.stacks[stacknum]
        if self.isEmpty(stacknum):
            raise Exception("Stack {} is empty".format(stacknum))

        value = self.array[stack.current_index]
        stack.current_index -= 1

        return value

    def peek(self, stacknum):
        stack = self.stacks[stacknum]
        if self.isEmpty(stacknum):
            raise Exception("Stack {} is empty".format(stacknum))

        return self.array[stack.current_index]


if __name__ == "__main__":
    newstack = ThreeInOne(2)
    print(newstack.isEmpty(1))
    newstack.push(3, 1)
    print(newstack.peek(1))
    print(newstack.isEmpty(1))
    newstack.push(2, 1)
    print(newstack.peek(1))
    print(newstack.pop(1))
    print(newstack.peek(1))
    newstack.push(4, 1)
    print(newstack.peek(1))
