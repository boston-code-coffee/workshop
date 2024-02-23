# Map problem #2 Finding Sums in Arrays

principle: maps

Imagine you're given a list of numbers (we'll call this list `nums`) and a special number called `target`. Your task is to find two numbers in `nums` that add up to `target` and then, as a bonus challenge, find three numbers in `nums` that add up to `target`.

### Part 1: Two Sum

**Your Task**: Write a function named `findTwoSum` that:
- Takes in `nums` (a list of numbers) and `target` (a single number).
- Finds two different numbers in `nums` that add up to `target`.
- Returns the positions (indexes) of those two numbers.

**Example**:
- If `nums` is `[2, 7, 11, 15]` and `target` is `9`, your function should return `[0, 1]` because `nums[0] + nums[1] = 2 + 7 = 9`.

### Part 2: Three Sum (Bonus Challenge)

**Your Task**: Enhance your solution to create a new function named `findThreeSum` that:
- Still takes in `nums` and `target`.
- Finds any three different numbers in `nums` that add up to `target`.
- Returns a list of these numbers.

**Example**:
- If `nums` is `[-1, 0, 1, 2, -1, -4]` and `target` is `0`, your function could return `[[-1, -1, 2], [-1, 0, 1]]`.

### Before You Start Coding

1. **Think About Assumptions**: What assumptions are you making about the input? How large can the list of numbers be? Can numbers repeat? Are there negative numbers?


2. **Ask Clarifying Questions**: If you're unsure about anything or if something seems ambiguous, what questions would you ask to get a clearer picture?


3. **Consider Different Approaches**: There's more than one way to solve these challenges. Can you think of a simple way to start? What are the limitations of your initial approach, and how might you overcome them?
