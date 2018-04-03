
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

if __name__ == "__main__":
    urlify("Mr John Smith    ", 13)
