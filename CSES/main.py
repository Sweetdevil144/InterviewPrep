from typing import List


def main():
    increasing_array()

def increasing_array():
    n = int(input())
    arr = list(map(int, input().split()))
    base, count = arr[0], 0
    for i in range(1, n):
        if arr[i] < base:
            count += base - arr[i]
            arr[i] = base
        base = arr[i]
    print(count)

def repetitions(str: str):
    m, i, j = 0, 0, 0
    while j < len(str) and i <= j:
        if str[i] == str[j]:
            j += 1
        else:
            m = max(m, j - i)
            i = j
            j += 1
    m = max(m, j - i)
    print(m)
    
def find_missing_number():
    n = int(input())
    nums = list(map(int, input().split()))
    total = n * (n + 1) // 2
    actual = sum(nums)
    print(total - actual)

if __name__ == "__main__":
    main()
