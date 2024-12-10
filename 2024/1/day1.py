input = open("./inputs/input.txt", "r").read()
list1 = []
list2 = []

for line in input.split("\n"):
    list1.append(int(line.split("  ")[0]))
    list2.append(int(line.split("  ")[1]))

list1.sort()
list2.sort()

sum1 = 0
for i in range(len(list1)):
    sum1 += abs(list1[i] - list2[i])

print(sum1)

# Part 2
sum2 = 0
for i in range(len(list1)):
    sum2 += list1[i] * list2.count(list1[i])

print(sum2)
