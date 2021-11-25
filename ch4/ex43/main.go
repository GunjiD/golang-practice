package main

func reverse(array *[5]int) {
	for i, j := 0, 4; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}
