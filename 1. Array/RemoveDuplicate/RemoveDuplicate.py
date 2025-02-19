lst =[10, 10, 10, 30, 20, 20, 20, 50, 60, 60, 60, 60]

def MoveAllZeros(lst)->list:
    unique_list = []
    for item in lst:
        if item not in unique_list:
            unique_list.append(item)
    return unique_list

print(MoveAllZeros(lst))
