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
