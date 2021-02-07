package main

import (
"bufio"
_ "database/sql/driver"
"fmt"
"os"
"strconv"
"strings"
)

/* 	go 没有“对象”，没有继承多态，没有泛型，没有try/catch
变量定义要点回顾：变量类型写在变量名之后；编译器可推测变量类型；没有char，只有rune；原生支持复数类型
常量的定义:用小写，必须赋值;const 数值可作为各种类型使用;
函数式编程
命令式编程是面向计算机硬件的抽象，有变量，赋值语句，表达式和控制语句，一句话，命令式程序就是一个冯诺依曼机的指令序列。
函数式编程是面向数学的抽象，将计算机描述为一种表达式求值，一句话，函数式程序就是一个表达式。
*/

//const file = "abc.txt"
//const x, y = 3, 4

//var a int = int(math.Sqrt(x*x + y * y))

// 用常量定义枚举类型。iota可实现自增
//const (
//	b = 1 << (10 * iota)
//	kb
//	mb
//	gb
//	tb
//	pb
//)

/* if 条件语句
if的条件里不需要括号
if条件里也可以定义变量
*/
//func ifGrammar() {
//	const filename = "abc.txt"
//	if contents, err := ioutil.ReadFile(filename); err == nil {
//		fmt.Println(string(contents))
//	} else {
//		fmt.Println("cannot print file contents:", err)
//	}
//}

// switch会自动break，除非使用fallthrough

/* for循环
for的条件里不需要括号
for的条件里可以省略初始条件，结束条件，递增表达式
无while，省略初始条件，就是while；省略所有条件，就是无限循环
*/
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err) // 报错打印
	}

	scanner := bufio.NewScanner(file)

	// 只有结束条件的for
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

/* 函数
返回值类型写在最后面
函数可以返回多个值；
函数返回多个值时可以起名字，仅用于非常简单的函数,对于调用者而言没有区别
func 函数名(变量1， 变量2 变量类型) （返回值1， 返回值2 返回值类型）{}
没有函数重载，没有默认参数、可选参数
只有可变参数列表
*/

/* 指针
go的指针不能运算
参数传递：
- c、c++既可以值传递，有可以引用传递
- java、python：引用传递，基本系统自定义类型都是引用传递
- go 只有值传递一种方式
*/

/* 数组
数量写在类型前
定义，有3种方式：
- var arr1 [5]int
- arr2 := [3]int{1, 3, 5}
- arr3 := [...]int{2, 4, 6, 8, 10}
var grid [4][5]bool //二维数组，4行5列
points := [][]int{
	{3,1},
	{9,0},
	{1,0},
	{1,4},
	{5,3},
	{8,8},
}

数组遍历，下标、值
for i,v := range arr {
	fmt.Println(i, v)
}
为什么要用range：
- 意义明确，美观
- c++没有类似能力
- java/python：只能for each value，不能同时获取i，v

重要概念：数组是值类型
eg：[10]int 和[20]int 是不同类型
调用func f(arr [10]int) 会拷贝数组
在go语言中一般不直接使用数组

应用：
字符串转数组：s_arr := strings.Split(s,"")
数组转字符串：// 参数：要拼接的数组，拼接的内容
    str := strings.Join([]string{`你好`, `世界`}, `,`)
*/

/* 切片 Slice
定义： []int
arr := [...]int{0, 2, 4, 6, 8, 10, 12}
arr[2:6]
slice是数组的视图
Slice本身没有数据，是对底层array的一个view
Slice可以向后扩展，不可以向前扩展
s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)
向Slice添加元素：
- 添加元素时如果超越cap，系统会重新分配更大的底层数组
- 由于值传递的关系，必须接受append的返回值
- s = append(s,val)
*/

func createSlice() {
	// 3种方式
	var s1 []int // 默认初始值为空nil
	s2 := []int{0, 1, 2, 3}
	s3 := make([]int, 16, 32) // 不给初始值，len=16, cap=32
	fmt.Println("s1, s2, s3 =", s1, s2, s3)
}

func slicePractice() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6] // 2,3,4,5
	s2 := s1[3:5]  // 5,6
	// 增加元素
	s3 := append(s2, 10) // [5,6,10]
	s4 := append(s3, 11) // [5,6,10,11]
	s5 := append(s4, 12) // [5,6,10,11,12]

	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Println("s3,s4,s5=", s3, s4, s5)
	fmt.Println("arr=", arr) // [0 1 2 3 4 5 6 10]

	// 删除第三个元素
	s6 := append(s2[:3], s2[4:]...)
	fmt.Println(s6)

	// 拷贝
	copy(s6, s5)
	fmt.Println(s5)
}

/* map
定义：
- map[K]V,map[K]map[K]V
- make(map[string]int)
遍历：range
删除：delete(m, "keyName")
使用len获得map个数
map的key：
- map使用哈希表，必须可以比较相等
- 除了slice，map，function的内建类型都可以作为key
- Struct类型不包含上述字段，也可作为key
*/

func createMap() {
	m1 := map[string]string{
		"name":   "Iris",
		"age":    "27",
		"course": "golang",
	}
	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]int      // m3 == nil

	// empty和nil都可以进行运算
	fmt.Println("m1, m2, m3 = ", m1, m2, m3)

	// 遍历
	for k, v := range m1 {
		fmt.Println("k, v = ", k, v)
	}

	// 判断key是否存在，如果不存在也不会报错,会获得value的初始值
	if username, ok := m1["username"]; ok {
		fmt.Println(username)
	} else {
		fmt.Println("Key does not exist")
	}

	// delete
	name, ok := m1["name"]
	fmt.Println(name, ok)
	delete(m1, "name")
	name, ok = m1["name"]
	fmt.Println(name, ok)
}

func findNoRepeatLetter(letters string) string {
	arrayLetters := strings.Split(letters, "")
	//fmt.Println("type:", reflect.TypeOf(arrayLetters))

	length := len(letters)
	fmt.Println("letters length=", length)
	location := 0
	maxLetters := arrayLetters[0]
	fmt.Println("maxLetters=", maxLetters)
	maxLength := 1

	for i := 0; i <= length-maxLength; i++ {
		var tmpMap map[int]string
		//tmpLetters := arrayLetters[i:i+maxLength]

		for j := i; j < i+maxLength; j++ {
			tmpMap[0] = "test"
			fmt.Println("tmpMap key, value = ", 0, tmpMap[0])
		}

	}

	fmt.Println(length, location, maxLetters)
	return "ok"
}

func main() {
	/*fmt.Println(
		convertToBin(5),		//101
		convertToBin(13),	//1101
	)

	*/
	//printFile("abc.txt")
	//slicePractice()
	//createMap()
	letters := "hello"
	findNoRepeatLetter(letters)
}

