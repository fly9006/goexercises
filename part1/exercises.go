package part1

import (
	"fmt"
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

}








