// https://www.acmicpc.net/problem/11382

// Q: 꼬마 정민이는 이제 A + B 정도는 쉽게 계산할 수 있다. 이제 A + B + C를 계산할 차례이다!
// input : 첫 번째 줄에 A, B, C (1 ≤ A, B, C ≤ 10^12)이 공백을 사이에 두고 주어진다.
// output : A+B+C의 값을 출력한다.

package main

import "fmt"

func main() {
	var a, b, c int64
	fmt.Scan(&a, &b, &c)
	fmt.Println(a + b + c)
}
