import unittest

class MyQueue(object):
    def __init__(self):
        self.stack_in = []
        self.stack_out = []

    def enqueue(self, item):
        self.stack_in.append(item)

    def dequeue(self):
        if len(self.stack_out) == 0:
            if len(self.stack_in) == 0:
                raise Exception("Queue Empty")

            while len(self.stack_in) > 0:
                val = self.stack_in.pop()
                self.stack_out.append(val)

        return self.stack_out.pop()

class Tests(unittest.TestCase):
    def test_stacks(self):
        queue = MyQueue()
        in_sto = []
        for i in range(35):
            queue.enqueue(i)
            in_sto.append(i)
        lst = []
        for _ in range(35):
            lst.append(queue.dequeue())
        self.assertEqual(in_sto, lst)


if __name__ == '__main__':
    unittest.main()
