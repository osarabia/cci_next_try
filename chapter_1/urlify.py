import unittest

def urlify(s, length):
    index = length - 1
    index1 = len(s) -1

    l = list(s)

    while index != index1:
        if l[index] != ' ':
            l[index1] = l[index]
            index -= 1
            index1 -= 1
        else:
            l[index1] = '0'
            l[index1-1] = '2'
            l[index1-2] = '%'
            index1 -= 3
            index -= 1

    return "".join(l)


class Test(unittest.TestCase):
    '''Test Cases'''
    data = [
        ('much ado about nothing      ', 22, 'much%20ado%20about%20nothing'),
        ('Mr John Smith    ', 13, 'Mr%20John%20Smith')]

    def test_urlify(self):
        for [test_string, length, expected] in self.data:
            actual = urlify(test_string, length)
            self.assertEqual(actual, expected)

if __name__ == "__main__":
    unittest.main()
