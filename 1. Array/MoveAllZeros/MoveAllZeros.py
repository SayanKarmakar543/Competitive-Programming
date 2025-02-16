lst = [10, 20, 0, 0, 0, 30, 40, 50]

def MoveAllZeros(lst):
    for i in range(len(lst)):
        if lst[i]==0:
            j = i
            for j in range(i+1, len(lst)):
                if lst[j]!=0:
                    lst[i], lst[j] = lst[j], lst[i] # swap of two elements
                    break
    return lst

print(MoveAllZeros(lst))