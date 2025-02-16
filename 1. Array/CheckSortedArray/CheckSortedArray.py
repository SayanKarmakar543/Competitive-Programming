lst = [10, 20, 30, 40, 50, 5]

def isSortedArray(lst):
    for element in lst:
        if (lst.index(element)!=len(lst)-1) and (element > lst[lst.index(element)+1]):
                return False
    return True

print(isSortedArray(lst))