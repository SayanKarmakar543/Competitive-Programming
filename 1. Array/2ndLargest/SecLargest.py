lst = [10, 40, 50, 30, 20]

def largestElement(lst):
    largest = lst[0]
    for element in lst:
        if element > largest:
            largest = element
    return largest

def secondLargestElement(lst, largestElement):
    secondLargest = lst[0]
    for element in lst:
        if element > secondLargest and element < largestElement:
            secondLargest = element
    return secondLargest
    
print(secondLargestElement(lst, largestElement(lst)))