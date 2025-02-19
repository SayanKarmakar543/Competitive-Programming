lst = [1, 4, 2, 10, 23, 3, 1, 0, 20]
k = 4

def sliding_window(lst, k):
    sum_lst = []
    index = 0
    for i in range(len(lst)-k+1):
        sum = 0
        for j in range(i, i+k):
            sum += lst[j]
        sum_lst.append(sum)

    largest = sum_lst[0]
    for item in sum_lst:
        if item > largest:
            largest = item
    
    return largest

print(sliding_window(lst, k))    