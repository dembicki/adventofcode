const fs = require("fs");

const input = fs.readFileSync("./inputs/input.txt", "utf8");

const sequences = input.split("\n").map((line) => line.split(" ").map(Number));

const isValidSequence = (numbers) => {
  if (numbers.length < 2) return false;

  const diffs = numbers.slice(1).map((n, i) => n - numbers[i]);

  const validDiffs = diffs.every(
    (diff) => Math.abs(diff) >= 1 && Math.abs(diff) <= 3
  );

  const allPositive = diffs.every((diff) => diff > 0);
  const allNegative = diffs.every((diff) => diff < 0);

  return validDiffs && (allPositive || allNegative);
};

const isValidSequenceWithOneError = (numbers) => {
  for (let i = 0; i < numbers.length; i++) {
    const newNumbers = [...numbers];
    newNumbers.splice(i, 1);
    if (isValidSequence(newNumbers)) return true;
  }
  return false;
};

const validSequences = sequences.filter(isValidSequence).length;
const validSequencesWithOneError = sequences.filter(
  isValidSequenceWithOneError
).length;

console.log(validSequences);
console.log(validSequencesWithOneError);
