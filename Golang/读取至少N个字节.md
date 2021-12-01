#读取至少N个字节

```go
package main

import (
        "io"
        "log"
        "os"
)

func main() {
        // 打开文件，只读
        file, err := os.Open("test.txt")
        if err != nil {
                log.Fatal(err)
        }
        defer file.Close()

        byteSlice := make([]byte, 512)
        minBytes := 8
        // io.ReadAtLeast() 读取文件的最小字节数, 在不能得到最小的字节的时候会返回错误，但会把已读的文件保留；
        numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)
        if err != nil {
                log.Fatal(err)
        }
        log.Printf("Number of bytes read: %d\n", numBytesRead)
        log.Printf("Data read: %s\n", byteSlice)
}
```