from collections import Counter
import string

letters = set(string.letters)

def even_palin(letters_counter):
    for k in letters_counter.keys():
        if letters_counter[k] % 2 > 0:
            return False

    return True

def odd_palin(letters_counter):
    impar_seen = False
    for k in letters_counter.keys():
        if letters_counter[k] % 2 == 1:
            if impar_seen is False:
                impar_seen = True
            else:
                return False

    return True

def is_palindrome_permu(s):

    if len(s) == 1:
        return True

    total_letters = 0
    letters_counter = Counter()

    for c in s:
        if c in letters:
            letters_counter[c.lower()] += 1
            total_letters += 1

    if total_letters % 2 == 0:
        return even_palin(letters_counter)

    return odd_palin(letters_counter)

assert(is_palindrome_permu('Tact Coa')) == True
assert(is_palindrome_permu('jhsabckuj ahjsbckj')) == True
assert(is_palindrome_permu('Able was I ere I saw Elba')) == True
assert(is_palindrome_permu('So patient a nurse to nurse a patient so')) == False
assert(is_palindrome_permu('Random Words')) == False
assert(is_palindrome_permu('Not a Palindrome')) == False
assert(is_palindrome_permu('no x in nixon')) == True
assert(is_palindrome_permu('azAZ')) == True
