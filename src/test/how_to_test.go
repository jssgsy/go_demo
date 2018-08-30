package test

//1. 必须引入testing包
import (
	"fmt"
	"testing"
)

/*
 * 2. 测试方法必须以Test开头，后面部分必须以大写字母开头
 * 3. 测试方法的参数类型为*testing.T
 */
func TestDemo(t *testing.T) {
	if 3 == Add(1, 2) {
		t.Log("测试通过了")
	} else {
		t.Error("测试失败了")
	}
}

func TestDemo2(t *testing.T) {
	if true {
		t.Log("测试通过了")
	} else {
		t.Error("测试失败了")
	}
}

/**
测试条件表示式
1. Go 完全省略了 if、switch 和 for 结构中条件语句两侧的括号
*/
func TestCondition(t *testing.T) {
	if i := 1; i >= 0 {
		fmt.Println("go中的条件表示式可以有初始化条件")
	}
	// 在if外读取不到初始化条件中的变量i
	//fmt.Println(i)
}

//----------------------------------------------------------------------
/*
函数相关测试
1. Go 默认使用按值传递来传递参数，也就是传递参数的副本。函数接收参数副本之后，在使用变量的过程中可能对副本的值进行更改，但不会影响到原来的变量；
2. 如果传递给函数的是一个指针，指针的值（一个地址）会被复制，但指针的值所指向的地址上的值不会被复制；我们可以通过这个指针的值来修改这个值所指向的地址上的值；
3. 在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）；
4. 指针也是变量类型，有自己的地址和值，通常指针的值指向一个变量的地址。所以，按引用传递也是按值传递;
*/
func TestNameReturn(t *testing.T) {
	if 3 == add(1, 2) {
		t.Log("命名返回值")
	} else {
		t.Error("命名返回值测试失败")
	}
}

/**
命名返回值，已经命名了返回值为sum，所以此时只需要使用return，而不用return sum
*/
func add(a, b int) (sum int) {
	sum = a + b
	return
}

//----------------------------------------------------------------------
func TestVariableParmas(t *testing.T) {
	variable_params(1, 2, 3, 4, 5)
}

/*
变长参数
1. 如果一个变长参数的类型没有被指定，则可以使用默认的空接口 interface{}，这样就可以接受任何类型的参数；
*/
func variable_params(params ...int) {
	fmt.Println("-------------func variable_params begin----------------")
	for i := 0; i < len(params); i++ {
		fmt.Println(params[i])
	}
	fmt.Println("-------------func variable_params end----------------")
}

// 此时可以接收任何类型
func variable_params2(param2 ...interface{}) {

}

//----------------------------------------------------------------------
func TestDefer(t *testing.T) {
	deferFunc()
}

/*
函数中的defer
类似于java中的finally块，只是defer会在return前一刻进行执行
*/
func deferFunc() {
	fmt.Println("entering deferFunc")
	// 注意语法
	defer func() {
		fmt.Println("这是在return前一刻执行的")
	}()
	fmt.Println("leaving deferFunc")
}

//----------------------------------------------------------------------
/*
测试将函数作为参数传递
*/
func TestCallback(t *testing.T) {
	callbackFn(1, callback)
}

/*
将参数作为参数传递
*/
func callbackFn(i int, fn func(x int) int) {
	fmt.Println(fn(i))
}
func callback(x int) int {
	return x + 10
}

//----------------------------------------------------------------------
