package univ

import (
	"fmt"
	"os"
)

//这种写法主要用于声明包级别的全局变量，当你在函数体内声明局部变量时，应使用简短声明语法 :=
var (
	HOME = os.Getenv("HOME")
	USER = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

/*
 左大括号 { 必须与方法的声明放在同一行，这是编译器的强制规定，否则报错
 */
func fn() {
	fmt.Println("首字母为小写相当于private，不能被外部引用")
}

func Fn2() {
	fmt.Println("首字母为大写相当于public，可以被外部引用")
}
//------------------------------------------------------------------------

// 全局变量允许声明但不使用,局部变量必须要被使用到
var g int

/**
变量定义
 */
func Var_definition()  {
	// 变量,类型在后
	// 变量的命名规则遵循骆驼命名法，即首个单词小写，每个新单词的首字母大写，例如：numShips 和 startDate
	var i int = 1
	// 变量至少被使用到，不然编译出错，但全局变量是允许声明但不使用
	fmt.Println(i)

	// 如果可以直接推导出变量类型，则可以省略var关键字，此时用 :=表示,这种形式一般用在局部变量的定义中
	j := i
	fmt.Println(j)

	// 可省略类型，只要此时可以推断出变量的值推断出变量的类型
	var k = 3
	fmt.Println(k)

	/*
	常量定义
	常量的值必须是能够在编译时就能够确定的；你可以在其赋值表达式中涉及计算过程，但是所有用于计算的值必须在编译期间就能获得(不能是需要动态计算才能获取的)
	  */
	const pi int = 12
	const pi2 = 34

	// 当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil。
	// 记住，所有的内存在 Go 中都是经过初始化的。
	var p int
	var q string
	fmt.Println(p)
	fmt.Println(q)

	// 只写变量"_"，即其不可读，
	// 常用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃
	_ = 4
	//fmt.Println(_)

	// 并行赋值，并行赋值也被用于当一个函数返回多个返回值时
	a,b := 1,2
	fmt.Println(a)
	fmt.Println(b)
}
//------------------------------------------------------------------------

/**
三种格式的输出
 */
func basic_print()  {
	// 换行输出
	fmt.Println("the basic print expression two ")

	fmt.Print("the basic print expression one ")
	fmt.Printf("the basic print expression %s","three")
	fmt.Println()
}
//------------------------------------------------------------------------


