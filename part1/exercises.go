package part1

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

 
//创建一个基于 for 的简单的循环。使其循环 10 次，并且使用 fmt 包打印出计数 器的值。
func Exercise001() {
	for i := 1; i < 11; i++ {
		fmt.Println("the current value is:", i)
	}
}

//用 goto 改写 1 的循环。关键字 for 不可使用
func Exercise002() {
	x := 1
	BEGIN:
		if x <= 10{
			fmt.Println("the current value is:", x)
			x++
			goto BEGIN
		}
}

//再次改写这个循环，使其遍历一个 array，并将这个 array 打印到屏幕上
func Exercise003() {
	arr := [...]int{1, 2, 3, 4, 5}
	//arr := [5]int{1, 2, 3, 4, 5}
	for _, v := range arr{
		fmt.Println("the current value is:", v)
	}
}

//解决这个叫做 Fizz-Buzz 的问题： 编写一个程序，打印从 1 到 100 的数字。当是3个倍数数就打印 “Fizz” 代替数字，当是5的倍数就打印 “Buzz” 。当数字同时是3和5的倍数 时，打印 “FizzBuzz” 。
func Exercise004() {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 != 0{
			fmt.Println("Fizz")
		}else if i % 3  != 0 && i % 5 == 0{
			fmt.Println("Buzz")
		}else if i % 3 == 0 && i % 5 == 0{
			fmt.Println("FizzBuzz")
		}else{
			fmt.Println("the number is:", i)
		}
	}

}

/*
建立一个 Go 程序打印下面的内容（到 100 个字符）：

      A
      
      AA
      
      AAA
      
      AAAA
      
      AAAAA
      
      AAAAAA
      
      AAAAAAA
*/
func Exercise005() {
	a := "A"
	for i := 1; i <= 100; i++ {
		fmt.Println(strings.Repeat(a, i), "\n")
	}
}

// 建立一个程序统计字符串里的字符数量： asSASA ddd dsjkdsjs dk 同时输出这个字符串的字节数。 提示： 看看 unicode/utf8 包。
func Exercise006() {
	str := "asSASA ddd dsjkdsjs dk你好"
	strCount := utf8.RuneCountInString(str)
	strRune := []rune(str)
	fmt.Println(strCount)
	fmt.Println(len(strRune))
}

// 扩展/修改上一个问题的程序，替换位置 4 开始的三个字符为 “abc”
func Exercise007() {
	str := "asSASA ddd dsjkdsjs dk你好"
	strRune := []rune(str)
	copy(strRune[4: 4+3], []rune("abc"))
	cstr := string(strRune)
	fmt.Println(cstr)
}

// 编写一个 Go 程序可以逆转字符串，例如 “foobar” 被打印成 “raboof”
func Exercise008() {
	str := "foobar"
	strRune := []rune(str)
	for from, to := 0, len(strRune)-1; from < to; from, to = from+1, to-1 {
		strRune[from], strRune[to] = strRune[to], strRune[from]
	}
	fmt.Println(string(strRune))
}

// 编写计算一个类型是 float64 的 slice 的平均值的代码
func Exercise009() {
	
	slice := []float64{12.1, 2314.123, 354.657}
	var total float64
	for _, v := range slice{
		total += v
	}
	fmt.Println(total / float64(len(slice)))
}
