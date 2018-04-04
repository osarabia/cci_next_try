import unittest

def string_compression(s):
    prev_char = s[0]
    count = 1
    new_string = []

    for i in xrange(1, len(s)):
        if prev_char == s[i]:
            count += 1
            prev_char = s[i]
        else:
            new_string.append(prev_char)
            new_string.append(str(count))
            prev_char = s[i]
            count = 1

    new_string.append(prev_char)
    new_string.append(str(count))

    if len(new_string) < len(s):
        return "".join(new_string)

    return s

class Test(unittest.TestCase):
    '''Test Cases'''
    data = [
        ('aabcccccaaa', 'a2b1c5a3'),
        ('abcdef', 'abcdef'),
        ('a', 'a'),
        ('bb', 'bb')
    ]

    def test_string_compression(self):
        for [test_string, expected] in self.data:
            actual = string_compression(test_string)
            self.assertEqual(actual, expected)

if __name__ == "__main__":
    unittest.main()
