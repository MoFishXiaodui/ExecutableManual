# 2-1 go语言进阶与依赖管理

相关资料

- PPT链接 [‌⁤‌⁡⁡⁢⁢⁢⁢‌﻿‍⁤﻿‍⁡‍⁣⁢‌⁢⁢﻿‬⁢⁤﻿‬‌‌‌‬‍⁣‍﻿⁣⁢‌‬‍⁤‌Go 语言入门 - 工程实践 - 赵征.pptx - 飞书云文档 (feishu.cn)](https://bytedance.feishu.cn/file/boxcnp0EHowxuLif7of3cBYxAmb)
- 官方预习文档 [【后端专场 学习资料一】字节跳动青训营 - 掘金 (juejin.cn)](https://juejin.cn/post/7188225875211452476#heading-8)

# Manual



此次可操作手册以PPT顺序为主，建议大家对照着PPT一起阅读，并且跟着操作

## 01 语言进阶

`从并发编程的视角带大家了解Go性能的本质`

### 案例1 - 快速打印

1. 建立新文件夹，命名为`2-1-1`

2. 创建文件，命名为 `main.go` (可自定义命名)

3. 在文件逐步中写入一下代码

   1. 先写个基本框架

      ```go
      package main
      
      func main() {
      	
      }
      ```

   2. 写个hello函数，实现打印整数 i 的功能

      ```go
      func hello(i int) {
      	println("hello goroutine :", fmt.Sprint(i))
      }
      ```

   3. 写一个函数HelloGoRoutine来调用hello函数，实现并发调用5个hello函数。

      ```go
      func HelloGoRoutine() {
      	for i := 0; i < 5; i++ {     
      		go func(j int) {
      			hello(j)
      		}(i) // 声明一个匿名函数并立即调用
      	}
          // 设置延时等待
          time.Sleep(time.Second)
      }
      ```

   4. 把HelloGoRoutine函数放到main里进行调用

      ```go
      func main() {
      	HelloGoRoutine()
      }
      ```

   5. 当你保存时，可以看到代码上方已经自动imports了fmt和time包。

   6. 最终代码见 [src/2-2-1/main](../src/2-1-1/main.go)
   
   7. 然后可以执行命令 `go run main.go` 看到一下效果
   
      ```bash
      (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-1-1> go run .\main.go
      hello goroutine : 4
      hello goroutine : 2
      hello goroutine : 3
      hello goroutine : 0
      hello goroutine : 1
      ```
   
      为了查看并发的效果，可以多跑几次。可以观察到每次输出的顺序都是不一样的

### 案例2 -  Channel

这个案例主要学习Channel的作用以及尝试使用有缓冲的Channel

1. 建立新的文件夹 `2-1-2channel` (文件夹命名不要用空格)

2. 创建新的文件 `channelWithBuf.go`

3. 然后往`channelWithBuf.go`文件写入代码

   1. 先写个基本框架

      ```go
      package main
      
      func main() {
      	CalSquare()
      }
      ```
      
   2. 定义函数`CalSquare`，阅读代码顺序参见注释的编号

      ```go
      func CalSquare() {
          // 1-创建一个无缓冲的channel: src，创建一个有缓冲的channel: dest
          src := make(chan int)
          dest := make(chan int, 3)
          
          // 2-创建一个立即执行的匿名routine，依次往src管道里面放入数字 0~9，并在最后关闭src管道
          go func() {
              defer close(src)
              for i := 0; i < 10; i++ {
                  src <- i
              }
          }()
          
          // 3-创建一个立即执行的匿名routine，接收src管道的数字，计算其平方放入dest管道中，并在最后关闭dest管道
          go func() {
              defer close(dest)
              for i := range src {
                  dest <- i * i
              }
          }()
          
          // 4-接收dest管道的数字并打印出来
          for i := range dest {
              fmt.Println(i)
          }
      }
      ```

   3. 保存文件时自动导入相关的包，最终的代码见 [src/2-1-2channel/channelWithBuf.go](../src/2-1-2channel/channelWithBuf.go)

   4. 执行`go run channelWithBuf.go`查看效果，大致如下：

      ```powershell
      (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-1-2channel> go run .\channelWithBuf.go
      0
      1
      4
      9
      16
      25
      36
      49
      64
      81
      ```
      
   
   

