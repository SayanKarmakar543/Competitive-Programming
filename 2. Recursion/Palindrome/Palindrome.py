str = "aba"

def is_palindrome(str, left, right):
    if left >= right:
        return True
    return (str[left]==str[right]) and is_palindrome(str, left+1, right-1)

print(is_palindrome(str, 0, len(str)-1))