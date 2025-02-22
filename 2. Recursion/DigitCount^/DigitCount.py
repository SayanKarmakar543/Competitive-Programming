def DigitCount(n):
    if n==0:
        return 0
    return (n%10)+DigitCount(int(n/10))

print(DigitCount(1234))