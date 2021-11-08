// noempty はスライス内アルゴリズムの例
package main

import (
	"fmt"
)

// nonempty は空文字列
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// append を使ったバージョン
func nonempty2(strings []string) []string {
	out := strings[:0] // 元の長さ0のスライス
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// 順序を保証しなければ最後の要素を空き要素へ
func remove(slice []int, i int) []int{
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}


func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // ["one" "three"]
	fmt.Printf("%q\n", data) // ["one" "three" "three"]
	// 大抵は data = nonempty(data) と書く

	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s,2))
}

