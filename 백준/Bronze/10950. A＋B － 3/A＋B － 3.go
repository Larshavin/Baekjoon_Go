// https://www.acmicpc.net/problem/10950
/*
Q) 두 정수 A와 B를 입력받은 다음, A+B를 출력하는 프로그램을 작성하시오.
input: 첫째 줄에 테스트 케이스의 개수 T가 주어진다.
각 테스트 케이스는 한 줄로 이루어져 있으며, 각 줄에 A와 B가 주어진다. (0 < A, B < 10)
output: 각 테스트 케이스마다 A+B를 출력한다.
*/

package main

import "fmt"

func main() {
	// Your solution here

	var t, a, b int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
