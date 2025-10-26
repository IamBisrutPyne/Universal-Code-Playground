import math

def is_prime(n):
    # Check if number is less than 2
    if n < 2:
        return False
    
    # Check if number is 2
    if n == 2:
        return True
    
    # Check if number is even
    if n % 2 == 0:
        return False
    
    # Check odd numbers up to square root of n
    for i in range(3, int(math.sqrt(n)) + 1, 2):
        if n % i == 0:
            return False
    
    return True

# Get input from user
try:
    num = int(input("Enter a number to check if it's prime: "))
    if is_prime(num):
        print(f"{num} is a prime number")
    else:
        print(f"{num} is not a prime number")
except ValueError:
    print("Please enter a valid integer")