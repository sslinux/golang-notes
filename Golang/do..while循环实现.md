* 语法格式：
```go
循环变量初始化
for {
	循环操作(语句)
	循环变量迭代
	if 循环条件表达式 {
		break  // 跳出for循环；
	}
}
```

* 示例

```go
var j int = 1
for {
	fmt.Println("hello, ok", j)
	j++
	
	if j > 10 {
		break
	}
}
```