package part2

import (
	"fmt"
	"strconv"
)

// Exerciese010 编写函数，返回其（两个）参数正确的（自然）数字顺序
func Exercise010(x, y int) (int, int) {
	if x <= y {
		fmt.Println(x, y)
		return x, y
	} else {
		fmt.Println(y, x)
		return y, x
	}
}

type stack struct {
	node int
	data [10]int
}

func (stack *stack) push(number int) {
	if stack.node < len(stack.data) {
		stack.data[stack.node] = number
		stack.node++
	}
	fmt.Println(stack.data)
}

func (stack *stack) pop() int {
	if stack.node > -1 {
		stack.node--
	}
	fmt.Println(stack.data[stack.node])
	return stack.data[stack.node]
}

func (stack *stack) string() string {
	var s string
	for i := 0; i < stack.node; i++ {
		k := strconv.Itoa(i)
		v := strconv.Itoa(stack.data[i])
		s += "[" + k + ":" + v + "]"
	}
	return s
}

// Exercise011 创建一个固定大小保存整数的栈。它无须超出限制的增长。定义 push 函数—— 将数据放入栈，和 pop 函数——从栈中取得内容。栈应当是后进先出（LIFO） 的。
func Exercise011() {
	stack_instance := stack{0, [10]int{}}
	stack_instance.push(1)
	stack_instance.push(4)
	stack_instance.push(5)
	stack_instance.push(12)
	stack_instance.push(31)
	stack_instance.push(42)
	stack_instance.pop()
	s := stack_instance.string()
	fmt.Println(s)
}

// Exercise012 编写函数接受整数类型变参，并且每行打印一个数字。
func Exercise012(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}

// Exercise013 斐波那契数列以：1,1,2,3,5,8,13,... 开始。或者用数学形式表达：x 1 = 1;x 2 = 1;x n = x n−1 + x n−2 ∀n > 2。 编写一个接受 int 值的函数，并给出这个值得到的斐波那契数列。
func Exercise013(number int) []int {
	s := make([]int, number)
	for i := 1; i <= number; i++ {
		if i < 3 {
			s[i-1] = 1
		} else {
			s[i-1] = s[i-2] + s[i-3]
		}
	}
	fmt.Println(s)
	return s
}

// Exercise014 编写一个函数，找到 int slice ( []int ) 中的最大值。
func Exercise014(numbers []int) int {
	var max int
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
	return max
}

// Exercise015 编写一个函数，找到 int slice ( []int ) 中的最小值。
func Exercise015(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	var min = numbers[0]
	for _, v := range numbers {
		if v < min {
			min = v
		}
	}
	fmt.Println(min)
	return min
}

// Exercise016 编写一个函数返回另一个函数，返回的函数的作用是对一个整数 +2。函数的名称叫做 plusTwo
func PlusTwo() func(int) int {
	f := func(number int) int {
		return number + 2
	}
	return f
}
func Exercise016(number int) {
	f := PlusTwo()
	fmt.Println(f(number))
}

// Exercise017 使 Exercise016 中的函数更加通用化，创建一个 plusX(x) 函数，返回一个函数用于对整 数加上 x
func PlusX(x int) func(int) int {
	f := func(y int) int {
		return x + y
	}
	return f
}
func Exercise017(x, y int) {
	f := PlusX(x)
	fmt.Println(f(y))
}

// Exercise018 冒泡排序
func Exercise018(numbers []int) {
	length := len(numbers)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	fmt.Println(numbers)
}
