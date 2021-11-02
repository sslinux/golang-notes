#goto

# goto基本介绍
* 1、Go语言的goto语句可以无条件地转移到程序中指定的行；
* 2、goto语句通常与条件语句配合使用。可用来实现条件转移，跳出循环体等功能；
* 3、在Go程序设计中一般不主张使用goto语句，以免造成程序流程的混乱，使理解和调试程序都产生困难；

# goto基本语法

```go
goto label:
....
label: statement
```

# goto代码示例

```go
func main() {
	var n int = 30
	
	if n > 20 {
		goto lable1
	}
	
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	label1:
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
}
```