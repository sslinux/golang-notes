#return
#return语句

# return介绍

return使用在[[方法]]或[[函数]]中，表示跳出所在的方法或[[函数]]。

* 示例

```go
func main() {
	for i := 1; i <= 10; i++ {
		if i==3 {
			return
		}
		fmt.Println("haha", i)
	}
	fmt.Println("hehe")
}
```

* 1、如果return是在普通的函数，则表示跳出该函数，即不在执行函数中return后面的代码，也可以理解成终止函数；
* 2、如果return是在main函数，表示终止main函数，也就是终止程序；

