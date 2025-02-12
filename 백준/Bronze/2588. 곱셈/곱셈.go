// https://www.acmicpc.net/problem/2588
// Q: (세 자리 수) × (세 자리 수)는 다음과 같은 과정을 통하여 이루어진다.
//     - 생략 -
//    (1)과 (2)위치에 들어갈 세 자리 자연수가 주어질 때 (3), (4), (5), (6)위치에 들어갈 값을 구하는 프로그램을 작성하시오.

// Input: 첫째 줄에 (1)의 위치에 들어갈 세 자리 자연수가, 둘째 줄에 (2)의 위치에 들어갈 세 자리 자연수가 주어진다.
// Output: 첫째 줄부터 넷째 줄까지 차례대로 (3), (4), (5), (6)에 들어갈 값을 출력한다.

package main

import (
	"fmt"
	"math"
)

func main() {

	var x, y int
	var b [3]int

	fmt.Scan(&x, &y)

	// 문자열을 숫자 배열로 변환
	// EX 123 -> [1, 2, 3]
	for i := 0; i < 3; i++ {
		b[i] = (y / int(math.Pow10(2-i))) % 10
	}

	sum2 := 0
	for i := 0; i <= 2; i++ {
		sum := x * b[2-i]
		fmt.Println(sum)
		sum2 += sum * int(math.Pow10(i))
	}
	fmt.Println(sum2)
}