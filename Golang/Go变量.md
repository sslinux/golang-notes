#变量

## 变量可见性

* 声明在函数内部，是函数的本地值，类似private；
* 声明在函数外部，是对当前包可见(包内所有.go文件都可见)的全局值，类似protect；
* 声明在函数外部且首字母大写是所有包可见的全局值，类似public；


> 不论是使用哪种高级程序语言编写程序，变量都是其程序的基本组成单位；

```go
package main

func getVal(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2
	return sum, sub
}

func main() {
	sum, sub := getVal(30, 30)
	fmt.Println("sum=", sum, "sub=", sub)
	sum2, _ := getVal(10, 30)  // 忽略第二个返回值
	fmt.Println("sum=", sum2)
}
```

* 变量的概念：
> 变量相当于内存中一个数据存储空间的表示，通过变量名可以访问到变量(值)

* 变量的使用步骤：
> 1、声明变量，也叫定义变量；
> 2、变量赋值；
> 3、使用变量；

* 变量使用示例：

```go
package main
import "fmt"

func main() {
	var i int		// 定义(声明)
	i = 10         // 赋值
	fmt.Println("i=", i)    // 使用
}
```


* 变量的三要素：
> 1、变量类型；
> 2、变量名；
> 3、变量值；

## Golang变量使用的三种方式

### 第一种：指定变量类型，声明后若不赋值，使用[[默认值]]

```go
package main

func main() {
	var i int
	fmt.Println("i=", i)
}
````

### 第二种： 根据值自行判定变量类型，即类型推导；

```go
var num = 10.11
fmt.Println("num=", num)
```
### 第三种： 省略var关键字，注意 := 左侧的变量不应该是已经声明过的，否则会导致编译错误；

```go
name := "tom"     // 这种方式只能在函数中声明变量，不能用来定义全局变量；
fmt.Println("name=", name)
```

* 多变量声明：

```go
package main
import "fmt"

func main() {
// 方式一：
var n1, n2, n3 int
fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

// 方式二：
var n1, name, n3 = 100, "tom", 888
fmt.Println("n1=", n1, "name=", name, "n3=", n3)

// 方式三：使用类型推导
n1, name, n3 := 100, "tom~", 888
fmt.Println("n1=", n1, "name=", name, "n3=", n3)
}
```

* 声明多个全局变量

```go
// 定义全局变量
var n1 = 100
var n2 = 200
var name = "jack"

// 一次性声明：
var (
	n3 = 300
	n4 = 900
	name2 = "mary"
)
```

* 变量在同一个作用域(一个函数或者在代码块)内不能重名；
* 变量=变量名+值+数据类型， 变量三要素；
* 变量如果没有赋初值，编译器会使用[[默认值]]


## [[变量声明]]
## [[变量初始化]]
## [[变量赋值]]
## [[加号]]
## [[Golang基本数据类型]]
## [[Golang扩展数据类型]]
# [[特殊标识符]]
