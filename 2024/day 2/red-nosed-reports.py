# part 1
def process_line_1(line):
    arr = [int(a) for a in line.split(" ")]

    is_increasing = arr[0] < arr[-1]
    
    for i in range(1, len(arr)):
        if is_increasing:
            if 1 <= arr[i] - arr[i - 1] <= 3:
                continue
            else:
                return 0 
        else:
            if 1 <= arr[i - 1] - arr[i] <= 3:
                continue
            else:
                return 0 
    return 1

# part 2
def process_line_3(line):
    arr = [int(a) for a in line.split(" ")]
    
    if process_line_1(line) == 1:
        return 1

    for i in range(len(arr)):
        l = ' '.join([str(a) for a in arr[:i]] + [str(a) for a in arr[i + 1:]])
        if process_line_1(l) == 1:
            return 1
    return 0

count = 0
with open("input.txt") as f:
    for line in f:
        count += process_line_3(line)
print(count)


        
        