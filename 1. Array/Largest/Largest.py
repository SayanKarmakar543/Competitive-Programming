def largestElement(lst)->int:
    largest = lst[0]
    for element in lst:
        if element > largest:
            largest = element
    return largest

lst = [20, 30, 50, 10, 40]

print(largestElement(lst))