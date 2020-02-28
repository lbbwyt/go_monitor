package dynamic

/**
一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法
*/
func ForgUpStair(n int) int {
	if n <= 1 {
		return 0
	}
	var dp = make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”,有多少种方法？
*/

func RobotMN(m, n int) int {
	if m < 0 || n < 0 {
		return 0
	}
	// 定义一个二维切片
	row, column := m+1, n+1
	var dp [][]int
	for i := 0; i < row; i++ {
		inline := make([]int, column)
		dp = append(dp, inline)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
