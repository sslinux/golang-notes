* 引入：随机生成1-100之间的一个数，生成的数为99，则退出；

* 代码：

```go
// 在go中，需要生成一个随机的种子，否则返回的值总是固定的；

// time.Now().Unix() : 返回一个从1970-1-1 0:0:0到现在的秒数；

rand.Seed(time.Now().Unix())
fmt.Println("n", rand.Intn(100)+1)
```

```go
func main() {
	var count_int = 0
	for {
		rand.Seed(time.Now().Unix())
		n := rand.Intn(100) + 1
		fmt.Println("n=",n)
		count_int++
		if n == 99 {
			break
		}
	}
	fmt.Printf("生成99一共使用了%d次\n", count_int)
}
```

^b18103

## break基本介绍

`break语句用于终止某个语句块的执行，用于中断当前for循环或跳出switch语句`

## break基本语法

```go
{
	......
	break
	......
}
```

## break注意事项和使用细节

* 1、break语句出现在多层嵌套的语句块中时，可以通过标签指明要终止的是哪一层语句块；
```go
// 演示，通过指定标签的形式来使用break
label2:
for i := 0; i < 4; i++ {
	label1:
	for j := 0; j < 10; j++ {
		if j == 2 {
			// break  // break 默认会跳出最近的for循环；
		   // break lable1
		   
		   break lable2    // j=0   j=1
		}
	}
}
```


## break应用示例

* 实现登录验证，有三次机会，如果用户名为“熊贵银”，密码为"888"则提示登录成功，否则提示还有几次机会；

```go
package main

func main() {
	var name string
	var pwd string
	var loginChance = 3
	
	for i := 1; i <= 3; i++ {
		fmt.Println("Please Type in Username:")
		fmt.Scanln(&name)
		fmt.Println("Please Type in Password:")
		fmt.Scanln(&pwd)
		
		if name == "熊贵银" && pwd == "888" {
			fmt.Println("恭喜你登录成功！")
			break
		} else {
			loginChance--
			fmt.Printf("你还有%d次登录机会，请珍惜\n", loginChance)
		}
	}
	if loginChance == 0 {
		fmt.Println("机会用完，没有登录成功！")
	}
}
```