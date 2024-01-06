# Stack problem #2 - Verify an expression is balanced

principle: stacks

Write a function called `isBalanced() which takes a string and returns a boolean.
The input string is a math expression. An expression  is string contains whitespace,
character symbols, integer numbers, the operators `+` `-` `*` `/` and the
nesting operators `()` `[]` `{}`.

The nesting operators group subexpressions.
Nesting operators may nest any number of subexpression, which also may nest.
An expression is called "balanced" if head head nesting
operators `(` `[` `{` has a matching tail operator `)` `]` `}`.

## Assumptions

- inputs are strings of any length that may fit in memory.
- inputs are ASCII whitespace, characters, digits, and the operators `+-/*{}[]()`
- input may or may not have nested subexpression and they may be of any depth
- the isBalanced() function only is verifying the nesting of the expression.

|Expression|IsBalanced|
|--|--|
||true|
|`123`|true|
|`123+x`|true|
|`()`|true|
|`(x)`|true|
|`x * (y-3)`|true|
|`{(x[1] + x[2]) * (y[3]-y[4]) }`|true|
|`(x]`|false|
|`)x+1(`|false|
|`((x+1)+y)`|false|
