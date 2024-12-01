from collections import Counter

first_list = []
second_list = []

with open('input.txt') as f:
    for content in f:
        a, b = content.split("   ")
        a, b = int(a), int(b)
        first_list.append(a)
        second_list.append(b)

# First Part
first_list.sort()
second_list.sort()

diff = 0
for a, b in zip(first_list, second_list):
    diff += abs(a - b)

print(diff)

# Second Part
f_c = Counter(first_list)
s_c = Counter(second_list)

score = 0

for num in f_c.keys():
    if num in s_c:
        score += num * s_c[num]

print(score)



