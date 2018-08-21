package main

// 如果你导入了一个包却没有使用它，则会在构建程序时引发错误，如 imported and not used: os，这正是遵循了 Go 的格言：“没有不必要的代码！“
import (
	"fmt"
	"com/univ"
)

/**
不能够被人为调用,可以在init函数中进行变量初始化
每个源文件都只能包含一个 init 函数
 */
func init()  {
	fmt.Println("init方法最先被执行，甚至在main方法之前")
}

/*
1. main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）
2. main 函数既没有参数，也没有返回类型（与 C 家族中的其它语言恰好相反）
*/
func main()  {
	univ.Var_definition()
	univ.Fn2()
}




