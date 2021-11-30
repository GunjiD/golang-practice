package main

// 参照はここ
// https://github.com/vinceyuan/gopl-solutions/blob/master/ch04/ex4.4/ex4.4.go
// append で slice ゴニョゴニョするやり方しか浮かばなかった

func rotate(s []int, n int) []int {
	tmp := make([]int, len(s))

	for i, j := 0, n; i < len(s); i, j = i+1, j+1 {
		if j == len(s) {
			j = 0
		}
		tmp[i] = s[j]
	}
	return tmp
}

//!-rev
