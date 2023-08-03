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
      
   
   

### 案例3-并发安全Lock

这个案例主要学习使用 sync.Mutex 锁让并发变得安全

1. 建新文件夹`2-1-3mutex`

2. 创建文件`mutex.go`

3. 在`mutex.go`文件中写入：

   1. 基本框架

      ```go
      package main
      
      func main() {
      	
      }
      ```

   2. 定义全局变量x和全局锁lock

      ```go
      var x int64
      var lock sync.Mutex
      ```

   3. 定义一个带锁操作x 和 一个不带锁操作x 的函数`addWithLock`和`addWithoutLock`

      ```go
      func addWithLock() {
      	for i := 0; i < 2000; i++ {
      		lock.Lock()
      		x += 1
      		lock.Unlock()
      	}
      }
      
      func addWithoutLock() {
      	for i := 0; i < 2000; i++ {
      		x += 1
      	}
      }
      ```

   4. 在main函数中并发五个goroutine去调用无锁的函数，等待1秒输出x的值。然后把x清零，再并发五个goroutine去调用有锁的函数，等待1秒输出x的值

      ```go
      func main() {
      	x = 0
      	for i := 0; i < 5; i++ {
      		go addWithoutLock()
      	}
      	time.Sleep(time.Second)
      	println("withoutLock: ", x)
      
      	x = 0
      	for i := 0; i < 5; i++ {
      		go addWithLock()
      	}
      	time.Sleep(time.Second)
      	println("withLock: ", x)
      }
      ```

      

4. 保存文件，自动导入`sync`和`time`包。最终代码[src/2-1-3mutex/mutex.go](../src/2-1-3mutex/mutex.go)

5. 多次运行`go run mutex.go`，查看效果。（有锁的能准确输出10000，无锁的就会有差错）

   ```powershell
   (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-1-3mutex> go run .\mutex.go
   withoutLock:  8475
   withLock:  10000
   (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-1-3mutex> go run .\mutex.go
   withoutLock:  8519
   withLock:  10000
   (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-1-3mutex> go run .\mutex.go
   withoutLock:  8359
   withLock:  10000
   ```



### 案例4 - WaitGroup

在[案例1](###案例1 - 快速打印)中，我们并发打印数字，是通过main中的time.Sleep()函数给予充足时间等待并发完全执行，这是一种不太理想的做法，我们在实际情况中通常无法评估需要多长时间进行等待合适。通过等待管道可以解决这个问题。在此案例，我们这里通过`sync`包下的`WaitGroup`结构更优雅地处理。

- sync.WaitGroup
  - Add(delta int) - 计数器 +delta
  - Done() - 计数器 -1
  - Wait() - 阻塞直到计数器为 0

1. 新文件夹 `2-1-4waitGroup`，新文件`waitGroup.go`

2. 在`waitGroup`中：

   1. 基本框架

   2. 直接在main中写：

      ```go
      func main() {
      	// 新建一个WaitGroup结构
      	wg := sync.WaitGroup{}
      
      	// 计数器 + 5
      	wg.Add(5)
      
      	// 并发快速打印 0-5
      	for i := 0; i < 5; i++ {
      		go func(num int) {
      			// 延迟计数器 -1
      			defer wg.Done()
      			println(num)
      		}(i)
      	}
      
      	// 阻塞等待计数器清零
      	wg.Wait()
      }
      ```

3. 保存自动导包`sync`，最终代码见[src/2-1-4waitGroup/waitGroup.go](../src/2-1-4waitGroup/waitGroup.go)

4. 直接运行

   ```powershell
   (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-1-3mutex> go run .\mutex.go
   withoutLock:  8519
   4
   2
   3
   0
   1
   (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-1-4waitGroup> go run .\waitGroup.go
   4
   2
   3
   1
   0
   ```

   

## 02 依赖管理

- 背景
- Go依赖管理演进
- Go Module实践

GoPath -> GoVendor -> GoModule

Go Module

- 通过go.mod文件管理依赖包版本
- 通过go get / go mod 指令工具管理依赖包

终极目标：定义版本规则和管理项目的依赖关系

依赖管理三要素

- 配置文件，描述依赖 - go.mod
- 中心仓库管理依赖库 - proxy
- 本地工具 - go get / mod

依赖配置-version

- 语义化版本
  - `${Major}.${Minor}.${Patch}`
  - Major版本不同会被认为是不同的模块
  - Minor版本通常是新增函数或功能，向后兼容
  - Patch版本一般是修复bug
- 基于commit伪版本
  - `vX.0.0-时间戳yyyymmddhhmmss-哈希前缀12位abcdefgh1234`

### 案例5 - 后面补充



## 03 测试

- 单元测试

  - 规则

    - 所有测试文件以 `_test.go` 结尾

    - 测试函数的签名：`func TestXxx(*testing.T)`

    - 初始化逻辑放到TestMain中 

      ```go
      func TestMain(m *testing.M) {
      	// 测试前：装载数据、配置初始化等前置工作
          code := m.Run()
          // 测试后：释放资源等收尾工作
      }
      ```

- Mock测试

- 基准测试

### 案例6 - 单元测试

1. 新建文件夹`2-1-6unitTest`，在内新建文件`helloTom.go`和`helloTom_test.go`

2. 在`helloTom.go`中书写以下代码

   ```go
   package __1_6unitTest
   
   func HelloTom() string {
   	return "Jerry"
   	//return "Tom"
   }
   ```

3. 在`helloTom_test.go`中书写以下代码

   ```go
   package __1_6unitTest
   
   import (
   	"testing"
   )
   
   func TestHelloTom(t *testing.T) {
   	// T is a type passed to Test functions to manage test state and support formatted test logs.
   	output := HelloTom()
   	expectOutput := "Tom"
   
   	if output != expectOutput {
   		t.Errorf("Expected %s do not match actual %s", expectOutput, output)
   	}
   }
   
   //func TestMain(m *testing.M) {
   // M is a type passed to a TestMain function to run the actual tests.
   //	code := m.Run()
   //	os.Exit(code)
   //}
   ```

4. 最终代码内容见 [src/2-1-6unitTest/](../src/2-1-6unitTest/)

5. 在终端中执行 `go test helloTom.go helloTom_test.go`，查看效果

   ```shell
   (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-1-6unitTest> go test helloTom_test.go helloTom.go    
   --- FAIL: TestHelloTom (0.00s)
       helloTom_test.go:13: Expected Tom do not match actual Jerry
   FAIL
   FAIL    command-line-arguments  0.184s
   FAIL
   
   # 把HelloTom()的返回值改为Tom后测试
   (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-1-6unitTest> go test helloTom_test.go helloTom.go
   ok      command-line-arguments  0.180s
   ```

   


### 案例7 - 单元测试assert

   此次案例采用外部库`github.com/stretchr/testify/assert`来测试。

   1. 复制[上一个案例](###案例6 - 单元测试)的`2-1-6unitTest`文件夹，更名为`2-1-7assertTest`
   
   2. 可以把两个文件的package都指定为main（就是代码的第一行）
   
   3. 在`2-1-7assertTest目录下`，执行命令`go mod int assertTest` 和 `go get "github.com/stretchr/testify/assert"`
   
      ```shell
      (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-1-7assertTest> go mod init assertTest
      go: creating new go.mod: module assertTest
      go: to add module requirements and sums:
              go mod tidy
      
      (base) 【...】\2-1-7assertTest> go get "github.com/stretchr/testify/assert"
      go: added github.com/davecgh/go-spew v1.1.1
      go: added github.com/pmezard/go-difflib v1.0.0
      go: added github.com/stretchr/testify v1.8.4
      go: added gopkg.in/yaml.v3 v3.0.1 
      ```
   
   4. 然后你可以看到新建的 `go.mod`和`go.sum`文件记录了模块和依赖信息
   
   5. 修改`helloTom_test.go`的代码，先在import导入新的包
   
      ```go
      import (
      	"testing"
      	"github.com/stretchr/testify/assert"
      )
      ```
   
   6. 把下面手动比较的代码换成assert的Equal函数
   
      ```go
      // 原来
      if output != expectOutput {
          t.Errorf("Expected %s do not match actual %s", expectOutput, output)
      }
      
      // 替换为
      assert.Equal(t, expectOutput, output)
      ```
   
   7. 最终代码见[src/2-1-7assertTest/](../src/2-1-7assertTest/)

   8. 和[案例6](###案例6 - 单元测试)一样执行测试命令，查看效果如下
   
      ```shell
      # HelloTom 返回 Tom 时
      (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-1-7assertTest> go test .\helloTom.go .\helloTom_test.go
      ok      command-line-arguments  0.199s
      
      # HelloTom 返回 Jerry时
      (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-1-7assertTest> go test .\helloTom.go .\helloTom_test.go
      --- FAIL: TestHelloTom (0.00s)
          helloTom_test.go:13:
                      Error Trace:    D:/code/MoFishXiaodui/ExecutableManual/src/2-1-7assertTest/helloTom_test.go:13
                      Error:          Not equal:
                                      expected: "Tom"
                                      actual  : "Jerry"
      
                                      Diff:
                                      --- Expected
                                      +++ Actual
                                      @@ -1 +1 @@
                                      -Tom
                                      +Jerry
                      Test:           TestHelloTom
      FAIL
      FAIL    command-line-arguments  0.204s
      FAIL
      ```
      



### 案例8 - 单元测试 覆盖率





