
// For more details, see: [balanced_expression.md](https://github.com/boston-code-coffee/workshop/blob/main/lessons/problems/stacks/balance_expression.md)
function isBalanced(sentence: string): boolean {
    const heads: Record<string, string> = {
        '(' : ')',
        '[' : ']',
        '{' : ']',
    };
    const tails: Record<string, string> = {
        ')' : '(',
        ']' : '[',
        '}' : '{',
    };
    let expect: string[] = [];
    for (let ch of sentence) {
      if (ch in heads) {
          expect.push(heads[ch])
      } else if (ch in tails) {
          if (ch !=  expect.pop()) {
              return false;
          }
      }
   }
   return  expect.length == 0;
}

// pass expected
console.log(isBalanced('x'));
console.log(isBalanced('(x)'));
console.log(isBalanced('[x+(1*2)]'));
console.log(isBalanced('[x+(1*2)]+(1+3)'));

// fail expected
console.log(isBalanced('[x+(1*2)('));
console.log(isBalanced('[x)'));


