#读取最多N个字节

> `os.File`提供了文件操作的基本功能， 而`io`、`ioutil`、`bufio`提供了额外的辅助函数。

```go
package main

import (
        "log"
        "os"
)

func main() {
        //打开文件，只读
        file, err := os.Open("test.txt")
        if err != nil {
                log.Fatal(err)
        }
        defer file.Close()

        // 从文件中控读取len(b)字节的文件；
        // 返回0字节意味着读取到文件尾了；
        // 读取到文件尾后会返回的io.EOF的error；
        byteSlice := make([]byte, 16)
        bytesRead, err := file.Read(byteSlice)
        if err != nil {
                log.Fatal(err)
        }
        log.Printf("Number of bytes read: %d\n", bytesRead)
        log.Printf("Data read: %s\n", byteSlice)
}
```

