iterator4go
===========

A iterator for Go. Based on http://ewencp.org/blog/golang-iterators/

A sample code:
```
package main

import (
    "fmt"
    it "github.com/virtao/iterator4go"
    "math/rand"
)

func main() {
    useIterator()
    useIntIterator()
    useTxtFileIterator("D:\\test.txt")
}

const dataSize int = 1000000

func initData() (ret []interface{}) {
    data := make([]interface{}, dataSize)
    for i := 0; i < dataSize; i++ {
        data[i] = rand.Int()
    }

    return data
}

func initIntData() (ret []int) {
    data := make([]int, dataSize)
    for i := 0; i < dataSize; i++ {
        data[i] = rand.Int()
    }

    return data
}

func useIterator() {
    data := initData()
    var sum int = 0
    itor := it.NewIterator(data)
    for itor.Next() {
        val, _ := itor.Value().(int)
        sum += val
    }
}

func useIntIterator() {
    data := initIntData()

    var sum int = 0
    itor := it.NewIntIterator(data)
    for itor.Next() {
        sum += itor.Value()
    }
}

func useTxtFileIterator(filePath string) {
    if itor := it.NewTxtFileIterator(filePath); itor == nil {
        fmt.Println("文件打开失败。")
    } else {
        defer itor.Close() //使用文本迭代器一定别忘了使用Close()函数。
        defer func() {
            if r := recover(); r != nil {
                fmt.Println("Panic : ", r)
            }
        }()
        for itor.Next() {
            fmt.Println(itor.Value())
        }
    }
}
```

Note:
interface{} is slow. If the type is known, then please try to use explicit type.

After calling the TxtFileIterator is complete, you must not forget to call Close() function.