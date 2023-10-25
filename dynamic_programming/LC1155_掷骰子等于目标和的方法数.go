package dynamic_programming

// 题目:
//这里有 n 个一样的骰子，每个骰子上都有 k 个面，分别标号为 1 到 k 。
//给定三个整数 n ,  k 和 target ，返回可能的方式(从总共 kn 种方式中)滚动骰子的数量，使正面朝上的数字之和等于 target 。
//答案可能很大，你需要对 109 + 7 取模 。

// 思路: 动态规划
// 二维数组f[i][j]表示,骰子个数为i,目标和为j的方法数
// f[i][j] = f[i-1][j-1] + f[i-1][j-2] + ... + f[i-1][j-k]

func numRollsToTarget(n int, k int, target int) int {
	var mod int = 1e9 + 7
	// initialize
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1
	// dynamic programmig
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(target, i*k); j++ {
			for h := 1; h <= min(j, k); h++ {
				dp[i][j] = (dp[i][j] + dp[i-1][j-h]) % mod
			}
		}
	}

	return dp[n][target]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
