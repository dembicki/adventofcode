const fs = require("fs");
const input = fs.readFileSync("./inputs/input.txt", "utf8");

const list1 = [];
const list2 = [];

input.split("\n").forEach((line) => {
  const a = parseInt(line.split("  ")[0]);
  const b = parseInt(line.split("  ")[1]);
  list1.push(a);
  list2.push(b);
});

// Part 1
const sum1 = list1.sort().reduce((acc, item, index) => {
  return acc + Math.abs(item - list2.sort()[index]);
}, 0);

console.log(sum1);

// Part 2
const sum2 = list1.reduce((acc, item) => {
  let count = 0;
  list2.forEach((item2) => {
    return item === item2 ? count++ : count;
  });
  return acc + item * count;
}, 0);

console.log(sum2);
