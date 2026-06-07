# Count Brackets / Parentheses Matcher

You are expected to implement a parenthesis matcher algorithm that can be used to review programs or strings for balanced structure based on opening and closing brackets.

## Problem Statement

Given a string expression containing text and braces, write a program to examine whether the pairs and the order of the following are correct in the expression:

- `{` and `}`
- `(` and `)`
- `[` and `]`

## Assumption

- You can assume that there are only three types of parenthesis combinations: `{}`, `()`, and `[]`.

## Expected Behaviour

Determine whether the expression is balanced, meaning every opening bracket has a matching closing bracket in the correct order.

## Test Cases

### Success Cases

1. Input: `expression = "{([])}"
   Output: `Balanced`
   Explanation: all the brackets are well-formed.

2. Input: `exp = "[()]{}{[()()]()}"
   Output: `Balanced`
   Explanation: all the brackets are well-formed.

3. Input: `exp = "T{his(is [a ] well)struc tured} expression."
   Output: `Balanced`
   Explanation: all the brackets are well-formed.

### Failure Case

- Input: `exp = "[(])"
  Output: `Not Balanced`
  Explanation: the brackets are not balanced because there is a closing `]` before the matching closing `)`.
