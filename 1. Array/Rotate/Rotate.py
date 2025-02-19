lst = [1, 2, 3, 4, 5]

def reverse(lst, low, high):
    while low < high :
        lst[low], lst[high] = lst[high], lst[low]
        low += 1
        high -= 1

    return lst

def rotate(lst, k):
    reverse(lst, 0, k-1)
    reverse(lst, k, len(lst)-1)
    reverse(lst, 0, len(lst)-1)

    return lst

print(rotate(lst, 2))