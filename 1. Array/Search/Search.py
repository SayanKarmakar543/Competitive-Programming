lst = [1, 2, 3, 4, 5]

def isExist(lst, n):
    if n in lst:
        return True
    return False

print(isExist(lst, 8))