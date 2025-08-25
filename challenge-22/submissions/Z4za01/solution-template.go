package main

import (
	"fmt"
	"sort"
)

func main() {
	// Standard U.S. coin denominations in cents
	denominations := []int{1, 5, 10, 25, 50}

	// Test amounts
	amounts := []int{87, 42, 99, 33, 7}

	for _, amount := range amounts {
		// Find minimum number of coins
		minCoins := MinCoins(amount, denominations)

		// Find coin combination
		coinCombo := CoinCombination(amount, denominations)

		// Print results
		fmt.Printf("Amount: %d cents\n", amount)
		fmt.Printf("Minimum coins needed: %d\n", minCoins)
		fmt.Printf("Coin combination: %v\n", coinCombo)
		fmt.Println("---------------------------")
	}
}

// MinCoins returns the minimum number of coins needed to make the given amount.
// If the amount cannot be made with the given denominations, return -1.
func MinCoins(amount int, denominations []int) int {
	// TODO: Implement this function
	dp := make([]int, amount + 1)
	
	for i := range dp{
	    dp[i] = amount + 1
	}
	dp[0] = 0
	
	for _, coin := range denominations {
	    for i := coin;i <= amount;i++{
	        if dp[i - coin] + 1 < dp[i]{
	            dp[i] = dp[i - coin] + 1
	        }
	    }
	}
	
	if dp[amount] > amount {
	    return -1
	}
	
	return dp[amount]
}

// CoinCombination returns a map with the specific combination of coins that gives
// the minimum number. The keys are coin denominations and values are the number of
// coins used for each denomination.
// If the amount cannot be made with the given denominations, return an empty map.
func CoinCombination(amount int, denominations []int) map[int]int {
	// TODO: Implement this function
	sort.Sort(sort.Reverse(sort.IntSlice(denominations)))
	m := map[int]int{
	    50:0,
	    25:0,
	    10:0,
	    5:0,
	    1:0,
	}

	remainingAmount := amount
	
	for _, coin := range denominations {
        // Take as many coins of this denomination as possible
        count := remainingAmount / coin
        if count > 0 {
            m[coin] = count
        }
        remainingAmount -= count * coin
        
        // If we've reached the target amount, we're done
        if remainingAmount == 0 {
            for coin := range m{
                if m[coin] == 0 {
                    delete(m, coin)
                }
            }
            return m
        }
    }
    
    if remainingAmount > 0 {
        for coin := range m{
            if m[coin] == 0 {
                delete(m, coin)
            }
        }
        return m
    }

	return m
}
