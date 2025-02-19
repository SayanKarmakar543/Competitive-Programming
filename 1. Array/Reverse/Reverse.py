lst = [1, 2, 3, 4, 5, 6]

def reverse_list(lst):
    j = len(lst)-1
    for i in range(0, int(len(lst)/2)):
        lst[i], lst[j] = lst[j], lst[i]
        j -= 1
    return lst

print(reverse_list(lst))