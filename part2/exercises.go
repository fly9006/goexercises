package part2

import (
	"fmt"
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
		stack.node ++
	}
	fmt.Println(stack.data)
}

func (stack *stack) pop() int {
	if stack.node > -1 {
		stack.node --
	}
	fmt.Println(stack.data[stack.node])
	return stack.data[stack.node]
}

// Exercise011 创建一个固定大小保存整数的栈。它无须超出限制的增长。定义 push 函数—— 将数据放入栈，和 pop 函数——从栈中取得内容。栈应当是后进先出（LIFO） 的。

func Exercise011() {
	stack_instance := stack{0, [10]int{}}
	stack_instance.push(1)
	stack_instance.push(4)
	stack_instance.push(5)
	stack_instance.pop()
}