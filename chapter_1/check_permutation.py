from collections import Counter


def check_permutation(string1, string2):

    if len(string1) != len(string2):
        return False

    counter_s1 = Counter(string1)
    counter_s2 = Counter(string2)

    for c in counter_s1:
        if counter_s1[c] != counter_s2[c]:
            return False

    return True

assert check_permutation("omar", "ramo") == True
assert check_permutation("oomar", "omara") == False
