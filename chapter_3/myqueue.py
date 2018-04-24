
class MyQueue(object):
    def __init__(self):
        self.stack_in = []
        self.stack_out = []

    def enqueue(self, item);
        self.stack_in.push(item)

    def dequeue(self):
        if len(self.stack_out) == 0:
            if len(self.stack_in) == 0:
                raise Exception("Queue Empty")

            while len(self.stack_in) > 0:
                val = self.stack_in.pop()
                self.stack_out.push(val)

        return self.stack_out.pop()
