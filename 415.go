// 415
package leecode

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World!")
	res := addStrings("541", "9")
	fmt.Println(res)
	input := []byte{}
	input = append(input, []byte{'h', 'e', 'l', 'l', 'o'}...)
	fmt.Printf("%v\n", input)
	fmt.Printf("%#v\n", input)
	fmt.Printf("%+v\n", input)
	fmt.Printf("%T\n", input)
	fmt.Printf("%U\n", input)
	fmt.Printf("%c\n", input)
	reverseString(input)
	fmt.Println(input[:])

	world := "world"
	fmt.Println(world)
	fmt.Println(world[:2])

	text1 := "abcdefg"
	for index, char := range text1 {
		fmt.Printf("%d %U %c \n", index, char, char)
	}

	fmt.Println(3 / 2)
}

func addStrings(num1 string, num2 string) string {
	// max := max(len(num1), len(num2))
	flag := 0
	res := ""
	i, j := len(num1)-1, len(num2)-1
	for i >= 0 || j >= 0 || flag != 0 {
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(num1[i] - '0')
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
		}
		sum := n1 + n2 + flag
		flag = sum / 10
		mod := sum % 10
		res = strconv.Itoa(mod) + res
		i--
		j--
	}
	return res
}

// func max(n1 int, n2 int) int {
// 	if n1 > n2 {
// 		return n1
// 	}
// 	return n2
// }

func reverseString(s []byte) {
	i := 0
	l := len(s) / 2
	for ; i < l; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
