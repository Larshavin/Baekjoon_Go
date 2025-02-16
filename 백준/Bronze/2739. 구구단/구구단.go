// https://www.acmicpc.net/problem/2739
/*
Q) N을 입력받은 뒤, 구구단 N단을 출력하는 프로그램을 작성하시오. 출력 형식에 맞춰서 출력하면 된다.
input: 첫째 줄에 N이 주어진다. N은 1보다 크거나 같고, 9보다 작거나 같다.
output: 출력형식과 같게 N*1부터 N*9까지 출력한다.
*/

package main

import (
	"fmt"
)

func main() {
	// Your solution here

	var n int
	fmt.Scan(&n)

	for i := 1; i <= 9; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, n*i)
	}
}
