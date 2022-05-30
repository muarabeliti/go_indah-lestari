package main

import "fmt"

func moneyCoins(money int) []int {
       var money, numCoins int
       fmt.Scanf("%d%d",&money, &numCoins )
       k := make([]int,numCoins)
       for i:=0;i<numCoins;i++ {
              fmt.Scanf("%d",&k[i] )
       }
       // make a DP array
       dp := make([]int,money+1)
       dp[0] = 1
       for i:=0;i<numCoins;i++ {
              start := k[i]
              for j:=start;j<=money;j++ {
                     // use the saved solution
                     dp[j] += dp[j-start]
              }
       }
       fmt.Println(dp[money])
}

func main() {
	fmt.Println(moneyCoins(123))
	fmt.Println(moneyCoins(432))
	fmt.Println(moneyCoins(543))
	fmt.Println(moneyCoins(7752))
	fmt.Println(moneyCoins(15321))
}