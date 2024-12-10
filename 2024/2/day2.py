input = open("./inputs/input.txt", "r").read()

sequences = [list(map(int, line.split(" "))) for line in input.split("\n")]


def is_valid_sequence(numbers):
    if len(numbers) < 2:
        return False

    diffs = [numbers[i] - numbers[i + 1] for i in range(len(numbers) - 1)]
    
    valid_diffs = all(1 <= abs(diff) <= 3 for diff in diffs)
    all_positive = all(diff > 0 for diff in diffs)
    all_negative = all(diff < 0 for diff in diffs)

    return valid_diffs and (all_positive or all_negative)

def is_valid_sequence_with_one_error(numbers):
    for i in range(len(numbers)):
        new_numbers = numbers[:i] + numbers[i+1:]
        if is_valid_sequence(new_numbers):
            return True
    return False

valid_sequences = sum(1 for sequence in sequences if is_valid_sequence(sequence))
valid_sequences_with_one_error = sum(1 for sequence in sequences if is_valid_sequence_with_one_error(sequence))

print(valid_sequences)
print(valid_sequences_with_one_error)
