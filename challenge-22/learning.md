# Learning Materials for Greedy Coin Change

## Understanding Greedy Algorithms

Greedy algorithms are a simple but powerful approach to solving optimization problems. They make locally optimal choices at each step, hoping that these choices will lead to a globally optimal solution. While greedy algorithms do not always yield the optimal solution for every problem, they are efficient and often provide good approximations.

### How Greedy Algorithms Work

The general process for a greedy algorithm is:

1. At each step, choose the locally optimal solution
2. Never reconsider previous choices
3. Proceed until a complete solution is reached

### Characteristics of Greedy Algorithms

- **Greedy Choice Property**: A globally optimal solution can be reached by making locally optimal choices.
- **Optimal Substructure**: The optimal solution to the problem contains optimal solutions to its subproblems.
- **Simplicity**: Greedy algorithms are typically easier to implement and understand than other approaches.
- **Efficiency**: Greedy algorithms are usually more efficient than dynamic programming or brute force approaches.

## Coin Change Problem

The coin change problem asks: given a set of coin denominations and a target amount, what is the minimum number of coins needed to make that amount?

### Greedy Approach to Coin Change

A greedy approach to the coin change problem follows these steps:

1. Sort the denominations in descending order
2. Start with the largest denomination
3. Take as many coins of the current denomination as possible without exceeding the target amount
4. Move to the next largest denomination
5. Repeat until the target amount is reached or all denominations are used

Here's a simple implementation in Go:

```go
func minCoins(amount int, denominations []int) int {
    // Sort denominations in descending order
    sort.Sort(sort.Reverse(sort.IntSlice(denominations)))
    
    coinCount := 0
    remainingAmount := amount
    
    for _, coin := range denominations {
        // Take as many coins of this denomination as possible
        count := remainingAmount / coin
        coinCount += count
        remainingAmount -= count * coin
        
        // If we've reached the target amount, we're done
        if remainingAmount == 0 {
            return coinCount
        }
    }
    
    // If we couldn't make the exact amount
    if remainingAmount > 0 {
        return -1
    }
    
    return coinCount
}
```

### Finding the Coin Combination

To find the specific combination of coins, we modify the algorithm to keep track of how many coins of each denomination are used:

```go
func coinCombination(amount int, denominations []int) map[int]int {
    // Sort denominations in descending order
    sort.Sort(sort.Reverse(sort.IntSlice(denominations)))
    
    combination := make(map[int]int)
    remainingAmount := amount
    
    for _, coin := range denominations {
        // Take as many coins of this denomination as possible
        count := remainingAmount / coin
        if count > 0 {
            combination[coin] = count
        }
        remainingAmount -= count * coin
        
        // If we've reached the target amount, we're done
        if remainingAmount == 0 {
            return combination
        }
    }
    
    // If we couldn't make the exact amount, return an empty map
    if remainingAmount > 0 {
        return map[int]int{}
    }
    
    return combination
}
```

## When Greedy Works (and When It Doesn't)

The greedy approach for coin change works optimally for the standard U.S. coin denominations (1, 5, 10, 25, 50, 100), but it doesn't always produce the optimal result for all sets of denominations.

### Example where Greedy Works

With U.S. denominations [1, 5, 10, 25, 50]:
- To make 42 cents, the greedy approach gives:
  - 1 quarter (25) + 1 dime (10) + 1 nickel (5) + 2 pennies (1) = 5 coins
  - This is indeed the optimal solution

### Example where Greedy Fails

With denominations [1, 3, 4]:
- To make 6 cents, the greedy approach gives:
  - 1 coin of value 4 + 2 coins of value 1 = 3 coins
  - But the optimal solution is 2 coins: 2 coins of value 3 = 6 cents

## When Does the Greedy Algorithm Work for Coin Change?

The greedy algorithm provides an optimal solution for the coin change problem if the coin system has the **canonical property**. A coin system has this property if each coin can be represented optimally using coins of smaller denomination.

The U.S. coin system (1, 5, 10, 25, 50, 100) has this property, which is why the greedy approach works for it.

## Dynamic Programming Alternative

For coin systems where the greedy approach doesn't always yield the optimal solution, dynamic programming is a more reliable approach:

```go
func minCoinsDP(amount int, denominations []int) int {
    // Initialize dp array with amount+1 (which is greater than max possible coins)
    dp := make([]int, amount+1)
    for i := range dp {
        dp[i] = amount + 1
    }
    dp[0] = 0 // Base case: 0 coins needed to make amount 0
    
    // Build up the dp array
    for _, coin := range denominations {
        for i := coin; i <= amount; i++ {
            if dp[i-coin]+1 < dp[i] {
                dp[i] = dp[i-coin] + 1
            }
        }
    }
    
    // If dp[amount] is still amount+1, then it's not possible to make the amount
    if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}
```

## Handling Edge Cases

When implementing the coin change algorithm, it's important to handle these edge cases:

1. **Zero amount**: Return 0 (no coins needed)
2. **Negative amount**: Return -1 or throw an error
3. **Cannot make exact change**: Return -1 or some indicator that it's not possible
4. **Empty denominations array**: Return -1 or throw an error

## Performance Considerations

- **Time Complexity**:
  - Greedy approach: O(n log n) for sorting + O(n) for finding coins = O(n log n)
  - Dynamic programming approach: O(n × amount) where n is the number of denominations
  
- **Space Complexity**:
  - Greedy approach: O(n) for storing the result
  - Dynamic programming approach: O(amount) for the dp array

## Further Reading

1. [Greedy Algorithms (GeeksforGeeks)](https://www.geeksforgeeks.org/greedy-algorithms/)
2. [Coin Change Problem (GeeksforGeeks)](https://www.geeksforgeeks.org/coin-change-dp-7/)
3. [When does the Greedy Algorithm Work for Coin Change?](https://graal.ift.ulaval.ca/~dadub100/ChapIV/node14.html)
4. [Dynamic Programming vs Greedy Approach](https://www.geeksforgeeks.org/dynamic-programming-vs-greedy-approach/) 