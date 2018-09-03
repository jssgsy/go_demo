package test

//1. 必须引入testing包
import (
	"fmt"
	"reflect"
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

//----------------------------------------------------------------------
/*
for-range结构
可以应用于数组和切片
key是数组或者切片的索引，value该索引位置的值；他们都是仅在 for 循环内部可见的局部变量。
注意：value 只是某个索引位置的值的一个拷贝，不能用来修改 slice1 该索引位置的值。
*/
func TestForRange(t *testing.T) {
	arr := []int{8, 43, 23, 8, 19}
	for key, value := range arr {
		value += 10
		fmt.Println("key: ", key, " value: ", value)
	}
	// value只是一个拷贝，不会影响原数组、切片
	fmt.Println(arr) // [8 43 23 8 19]
}

//----------------------------------------------------------------------

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
	// 注意语法,这是一个闭包(匿名函数)，最后的一对括号表示对该匿名函数的调用
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
	callbackFn(1, func(x int) int {
		return x + 10
	})
}

/*
将函数作为参数传递
*/
func callbackFn(i int, fn func(int) int) {
	fmt.Println(fn(i))
}

//----------------------------------------------------------------------

/*
闭包(匿名)函数的测试，闭包经常和defer一起使用
*/
func TestAnonyouFunc(t *testing.T) {
	k := 30
	// 1.1 匿名函数直接调用
	// 1.2 匿名函数可以引用其外部的变量
	func(i, j int) {
		fmt.Println(i + j + k)
	}(10, 20)

	// 2. 匿名函数赋值给变量，通过变量调用
	fn := func(i, j int) int { return i + j }
	fmt.Println(fn(10, 20))
}

//----------------------------------------------------------------------

/*
测试指针
Go 语言和 C、C++ 以及 D 语言这些低级（系统）语言一样，都有指针的概念。
但是对于经常导致 C 语言内存泄漏继而程序崩溃的指针运算（所谓的指针算法，如：pointer+2，移动指针指向字符串的字节数或数组的某个位置）是不被允许的。
Go 语言中的指针保证了内存安全，更像是 Java、C# 和 VB.NET 中的引用
指针的一个高级应用是你可以传递一个变量的引用（如函数的参数），这样不会传递变量的拷贝。指针传递是很廉价的
*/
func TestPoint(t *testing.T) {
	i := 1
	j := &i
	fmt.Println("指针的值为：", j, "指针的类型为：", reflect.TypeOf(j)) // 0xc4200961b0, *int
	fmt.Println("指针指向的内容为：", *j)                           // 1

	var p *int
	// 两者类型不一致，不能赋值
	//p = i
	// 两者类型一致，可赋值
	p = j
	// 此时修改会反映到j指向的内容上上
	*p = 3
	fmt.Println(*j)

	// 不能得到一个文字或常量的地址
	const n = 7
	//fmt.Println(&n)	// Cannot take the address of 'n'
	//fmt.Println(&10)	// Cannot take the address of 'n'
}

//----------------------------------------------------------------------

/*
测试slice结构
*/
func TestSlice(t *testing.T) {
	// 定义一个数组
	var arr = []int{1, 3, 4, 6, 7, 9}
	fmt.Println("最原始的数组：", arr) //[1 3 4 6 7 9]
	// 数组当然也是一个slice
	fmt.Println("数组容量：", cap(arr)) //6

	// 从数组中获取一个slice，左闭右开
	var slice1 = arr[2:5]
	fmt.Println("slice1大小：", len(slice1)) // 3

	// 容量是向右扩展到极限时的大小
	fmt.Println("slice1容量：", cap(slice1)) // 4

	// 注意，slice只是原数组的一段区间的引用，所以的这里的变化会反应到原数组上
	slice1[0] = 100
	fmt.Println("slice1: ", slice1) // [100 6 7]
	fmt.Println("arr: ", arr)       // [1 3 100 6 7 9]

	// slice就是一个小数组，所以可以使用for与for-range等
	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}

	for key, value := range slice1 {
		fmt.Println("slice1 key: ", key, " slice1 value: ", value)
	}

	// 数组定义,此时索引仍然会从0开始补齐，没赋值索引处的值会被置为相应类型的空值
	var arr2 = []int{2: 222, 3: 333, 7: 777}
	for key, value := range arr2 {
		fmt.Println("arr2 key: ", key, " arr2 value: ", value)
	}

	// 注意，长度也是数组类型的一部分
	var arr3 = [3]int{1, 2, 3}
	var arr4 = []int{1, 2, 3}
	var arr5 = new([4]int)
	fmt.Println(reflect.TypeOf(arr3)) // [3]int
	fmt.Println(reflect.TypeOf(arr4)) // []int
	fmt.Println(reflect.TypeOf(arr5)) // *[4]int
	fmt.Println(arr5)                 // &[0 0 0 0]
}

/**
测试切片1
*/
func TestSlice1(t *testing.T) {
	// 使用make来创建切片，这种方法可以用在还没有创建相关数组时，
	// make方法签名：func make([]T, len, cap)，第二个参数是切片的len，第三个参数是切片的cap，即底层数组的len
	var slice1 = make([]int, 3, 10)
	fmt.Println("slice1", slice1)
	fmt.Println("len(slice1): ", len(slice1))
	fmt.Println("cap(slice1): ", cap(slice1))

	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] += 1

	}
	fmt.Println("重组后的slice1: ", slice1)

	// 切片复制与追加
	slice2 := []int{45, 23, 6, 1, 98, 3}
	slice3 := []int{10, 20, 30}
	// slice3是dest，slice2是src，相应处的元素会被覆盖,copy返回被复制的个数
	n := copy(slice3, slice2)
	fmt.Println("复制的元素个数为：", n)
	fmt.Println(slice2)
	fmt.Println(slice3)

	// append返回扩容后的新切片
	result := append(slice3, 9, 8, 7)
	fmt.Println(result)
}

//----------------------------------------------------------------------
/**
new与make的区别
new(T)：为新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：适用于值类型如数组和结构体;
func make([]T, len, cap)：返回一个类型为T的初始值，它只适用于3种内建的引用类型：切片、map 和 channel
*/
func TestNewMake(t *testing.T) {
	n := new([]int)
	m := make([]int, 3)
	fmt.Println(n)
	fmt.Println(reflect.TypeOf(n)) // *[]int
	fmt.Println(m)
	fmt.Println(reflect.TypeOf(m)) // []int
}

//----------------------------------------------------------------------
