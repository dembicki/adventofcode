import re

input = open("./inputs/input.txt", "r").read()
regex = r"mul\((\d{1,3}),(\d{1,3})\)"
matches = re.findall(regex, input)

# Part 1
sum = 0
for match in matches:
    sum += int(match[0]) * int(match[1])

print(sum)

# Part 2

# We should stop treating as valid if we see dont() and start treating as valid if we see do() again

doRegex = r"do\((\d{1,3})\)"
dontRegex = r"dont\((\d{1,3})\)"

doMatches = re.findall(doRegex, input)
dontMatches = re.findall(dontRegex, input)

print(doMatches)
print(dontMatches)

print(len(doMatches) - len(dontMatches))