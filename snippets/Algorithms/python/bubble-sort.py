# Function to perform Bubble Sort on a list
def bubble_sort(arr):
    n = len(arr)  # Get the length of the list
    # Outer loop for each element in the list
    for i in range(n):
        # Inner loop to compare adjacent elements in the unsorted part of the list
        for j in range(0, n - i - 1):
            # If the current element is greater than the next, swap it
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]  # Swap
    return arr  # Return the sorted list

# Example usage of the Bubble Sort function
myList = [64, 34, 25, 12, 22, 11, 90]  # Sample unsorted list
sortedList = bubble_sort(myList)        # Call the bubble_sort function
print("Sorted array:", sortedList)      # Print the sorted result
