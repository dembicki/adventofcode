const fs = require("fs");

const input = fs.readFileSync("./inputs/input.txt", "utf8");
// python regex: regex = r"mul\((\d{1,3}),(\d{1,3})\)"
const regex = /mul[\(\[]\d+,\d+[\)\]>](?!\(\d+,\d+\))/g;
const matches = input.matchAll(regex);

let sum = 0;
for (const match of matches) {
  sum += parseInt(match[1]) * parseInt(match[2]);
}

console.log(sum);
