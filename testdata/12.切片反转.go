package main

import "fmt"

func main() {
	i := []int{6, 2, 3, 4}
	Reverse(i)
	fmt.Println(i) //[4 3 2 6]

	s := []string{"哈哈哈", "嘻嘻嘻", "略略略", "lbwnb", "321"}
	Reverse(s)
	fmt.Println(s) //[321 lbwnb 略略略 嘻嘻嘻 哈哈哈]
}

// Reverse 切片反转
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
