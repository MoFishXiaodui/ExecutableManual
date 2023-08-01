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
   
   

