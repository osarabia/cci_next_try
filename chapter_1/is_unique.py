from sets import Set

def is_unique(astring):
    chars = Set([])

    for c in astring:
        if c in chars:
            return False
        chars.add(c)

    return True

assert is_unique("omar") == True
assert is_unique("omaro") == False
assert is_unique("anita") == False
