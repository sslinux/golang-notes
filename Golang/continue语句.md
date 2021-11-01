#continue

## continue基本介绍

* continue语句用于结束本次循环，继续下一次循环；
* continue语句出现在多层嵌套的循环语句体中时，可以通过标签指明要跳过的是哪一层循环，和break一样； ^f65a81

## continue基本语法

```go
{
	......
	continue
	......
}
```

## continue示例

```go
lable2:
for i := 0; i < 4; i++ {
	label1:
	for j := 0; j < 10; j++ {
		if j == 2 {
			continue lable1
		}
		fmt.Println("j=", j)
	}
}
```