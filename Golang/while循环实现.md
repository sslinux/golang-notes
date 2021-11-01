#while循环实现

```go
循环变量初始化
for {
	if 循环条件表达式 {
		break // 跳出表达式
	}
	循环操作(语句)
	循环变量迭代
}
```

```go
package main

func main() {
	var i int = 1
	for {
		if i > 10 {
			break
		}
		fmt.Println(i, " 小于10")
		i++
	}
	
	fmt.Println("i=", i)
}
```