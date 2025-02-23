import (
	"math"
)

func solution(info [][]int, n int, m int) int {
	const INF = math.MaxInt32
	itemCount := len(info)

	// DP 테이블 선언
	dp := make([][]int, itemCount+1)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	
	// 초기 상태
	dp[0][0] = 0

	// DP 진행
	for i := 0; i < itemCount; i++ {
		aTrace := info[i][0] // A 도둑의 흔적
		bTrace := info[i][1] // B 도둑의 흔적

		for j := 0; j < m; j++ {
			if dp[i][j] == INF {
				continue
			}

			// A가 훔친 경우
			newBTrace := j
			newATrace := dp[i][j] + aTrace
			if newATrace < n {
				dp[i+1][newBTrace] = min(dp[i+1][newBTrace], newATrace)
			}

			// B가 훔친 경우
			newBTrace = j + bTrace
			newATrace = dp[i][j]
			if newBTrace < m {
				dp[i+1][newBTrace] = min(dp[i+1][newBTrace], newATrace)
			}
		}
	}

	// 최적해 찾기
	result := INF
	for j := 0; j < m; j++ {
		result = min(result, dp[itemCount][j])
	}

	if result == INF {
		return -1
	}
	return result
}

// 최소값 함수
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}