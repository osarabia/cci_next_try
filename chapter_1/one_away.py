import unittest

def one_away(s1, s2):
    if len(s1) > len(s2) and len(s1) - len(s2) > 1:
        return False

    if len(s2) > len(s1) and len(s2) - len(s1) > 1:
        return False

    i = 0
    j = 0

    while i < len(s1) and j < len(s2):
        if s1[i] != s2[j]:
            if len(s1) == len(s2):
                return s1[i+1:] == s2[i+1:]
            else:
                if len(s1) > len(s2):
                    return s1[i+1:] == s2[j:]

                return s2[j+1:] == s1[i:]
        i += 1
        j += 1

    return True


class Test(unittest.TestCase):
    '''Test Cases'''
    data = [
        ('pale', 'ple', True),
        ('pales', 'pale', True),
        ('pale', 'bale', True),
        ('paleabc', 'pleabc', True),
        ('pale', 'ble', False),
        ('a', 'b', True),
        ('', 'd', True),
        ('d', 'de', True),
        ('pale', 'pale', True),
        ('pale', 'ple', True),
        ('ple', 'pale', True),
        ('pale', 'bale', True),
        ('pale', 'bake', False),
        ('pale', 'pse', False),
        ('ples', 'pales', True),
        ('pale', 'pas', False),
        ('pas', 'pale', False),
        ('pale', 'pkle', True),
        ('pkle', 'pable', False),
        ('pal', 'palks', False),
        ('palks', 'pal', False)
    ]

    def test_one_away(self):
        for [test_s1, test_s2, expected] in self.data:
            actual = one_away(test_s1, test_s2)
            self.assertEqual(actual, expected)

if __name__ == "__main__":
    unittest.main()
