package main

import "fmt"

// golang中字符类型的使用
func main() {
	var c1 byte = 'a'
	var c2 byte = '0' //字符的0

	// 当直接输出byte值时，输出的是对应字符的码值；
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)

	// var c3 byte = '北' // overflow溢出
	var c3 int = '北'
	fmt.Printf("c3=%c, c3对应码值=%d\n", c3, c3)

	// for i := 32; i < 50000; i++ {
	// 	fmt.Printf("%c", i)
	// 	if i%10 == 0 {
	// 		fmt.Println()
	// 	}
	// }
	var str = "hello"
	fmt.Printf("%T\n", str[1])
	fmt.Printf("%T\n", str)
}
