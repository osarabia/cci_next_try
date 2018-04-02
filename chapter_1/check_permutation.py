
def check_permutation(string1, string2):

    if len(string1) != len(string2):
        return False

    counter_s1, counter_s2 = {}, {}

    for c1 in string1:
        if c1 in counter_s1:
            counter_s1[c1] += 1
        else:
            counter_s1[c1] = 1

    print(counter_s1)
    for c2 in string2:
        if c2 in counter_s2:
            counter_s2[c2] += 1
        else:
            counter_s2[c2] = 1
    print(counter_s2)

    for c in counter_s1:
        if counter_s1[c] != counter_s2[c]:
            return False

    return True

assert check_permutation("omar", "ramo") == True
assert check_permutation("oomar", "omara") == False
