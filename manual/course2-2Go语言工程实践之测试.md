# 2-2 Go工程实践之测试

相关资料

- PPT链接 [‌⁤‌⁡⁡⁢⁢⁢⁢‌﻿‍⁤﻿‍⁡‍⁣⁢‌⁢⁢﻿‬⁢⁤﻿‬‌‌‌‬‍⁣‍﻿⁣⁢‌‬‍⁤‌Go 语言入门 - 工程实践 - 赵征.pptx - 飞书云文档 (feishu.cn)](https://bytedance.feishu.cn/file/boxcnp0EHowxuLif7of3cBYxAmb)
- 官方预习文档 [【后端专场 学习资料一】字节跳动青训营 - 掘金 (juejin.cn)](https://juejin.cn/post/7188225875211452476#heading-8)

# Manual 

此次可操作手册以PPT顺序为主，建议大家对照着PPT一起阅读，并且跟着操作

## 03 测试

**由于忘记了测试是新一节课的内容，导致部分 03测试章节的内容留存在2-1中，未学习的读者需要自行往回翻阅**

### 案例10 - 依赖文件的单元测试

1. 新文件夹 `2-2-10file`，在内新建`file.go` 、`file_test.go`和`log.txt` 文件

2. log.txt

   ```
   line11
   line22
   line33
   ```

3. file.go 

   ```go
   package main
   
   import (
   	"bufio"
   	"fmt"
   	"os"
   	"strings"
   )
   
   func ReadFirstLine() string {
   	file, err := os.Open("log.txt")
   	defer file.Close()
   	if err != nil {
   		return ""
   	}
   	scanner := bufio.NewScanner(file)
   	for scanner.Scan() {
   		return scanner.Text()
   	}
   	return ""
   }
   
   func ProcessFirstLine() string {
   	line := ReadFirstLine()
   	destLine := strings.ReplaceAll(line, "11", "00")
   	return destLine
   }
   
   func main() {
   	line := ProcessFirstLine()
   	fmt.Println(line)
   }
   ```

4. file_test.go

   ```go
   package main
   
   import (
   	"github.com/stretchr/testify/assert"
   	"testing"
   )
   
   func TestProcessFirstLine(t *testing.T) {
   	firstLine := ProcessFirstLine()
   	assert.Equal(t, "line00", firstLine)
   }
   ```

5. 先运行 `go run file.go`查看效果

   ```shell
   src\2-2-10file> go run .\file.go
   line00
   ```

6. 再运行 `go test file.go file_test.go` 查看效果

   ```shell
   (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-2-10file> go test .\file.go .\file_test.go
   ok      command-line-arguments  0.213s
   
   (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-2-10file> go test .\file.go .\file_test.go --cover
   ok      command-line-arguments  0.227s  coverage: 69.2% of statements
   ```

7. 总结

   1. `ReadFirstLine` 函数可以读取 `log.txt` 文件并返回第一行
   2. `ProcessFirstLine` 函数可以调用 `ReadFirstLine` 函数获取第一行文本副本，并对该副本做加工，把"11"替换成"00"
   3. `TestProcessFirstLine` 函数用来测试 `ProcessFirstLine` 函数是否正确
   4. `ReadFirstLine` 函数的内容返回依赖于文件`log.txt`，如果`log.txt`被删除，就不能测试`ProcessFirstLine`的正确性。接下来的案例我们将设法替换 `ReadFirstLine`函数，使测试不依赖于 `log.txt`文件

### 案例11 - mock

mock 

- 为一个函数打桩
- 为一个方法打桩

此次mock需要用到"bou.ke/monkey"包

1. 复制[案例10](###案例10 - 依赖文件的单元测试)的整个文件夹，命名为:`2-2-11mock`，删掉`log.txt`文件。(把`ReadFirstLine`的依赖删除)

2. 添加monkey包，在终端运行`go get "bou.ke/monkey"`

   ```shell
   (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-2-11mock> go get "bou.ke/monkey"
   go: added bou.ke/monkey v1.0.2
   ```

3. 修改`file_test.go`文件的测试函数代码。（就是在中间多了一个函数替换操作，也就是打桩）

   ```go
   // 原来
   func TestProcessFirstLine(t *testing.T) {
   	firstLine := ProcessFirstLine()
   	assert.Equal(t, "line00", firstLine)
   }
   
   // 现在
   func TestProcessFirstLine(t *testing.T) {
   	monkey.Patch(ReadFirstLine, func() string {
   		return "line11"
   	})
   	defer monkey.Unpatch(ReadFirstLine)
   
   	firstLine := ProcessFirstLine()
   	assert.Equal(t, "line00", firstLine)
   }
   ```

4. 最终代码见[src/2-2-11mock/](../src/2-2-11mock/)

5. 同[案例10](###案例10 - 依赖文件的单元测试)一样测试

   ```shell
   (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-2-11mock> go test .\file.go .\file_test.go --cover 
   ok      command-line-arguments  0.218s  coverage: 23.1% of statements
   ```

6. 总结：mock可以通过替换函数的方式，从而可以删去部分依赖。



### 案例12 - 基准测试

1. 新建文件夹 `2-2-12bench` ，在内新建文件 `server.go` 和  `server_test.go`

2. 初始化，在 `2-2-12bench` 目录下执行 `go mod init server`

3. 在 `server.go`中写入代码

   ```go
   package bench
   
   import "math/rand"
   
   var ServerIndex [10]int
   
   // 初始化服务函数，（其实就是往数组里面各个项写入特定数字，i+100本身应该没什么含义）
   func InitServerIndex() {
   	for i := 0; i < 10; i++ {
   		ServerIndex[i] = i + 100
   	}
   }
   
   func Select() int {
   	// 随意选择一个服务器（随机选择一个数组元素，模拟随机一个服务器返回数字）
   	return ServerIndex[rand.Intn(10)]
   }
   ```

4. 在  `server_test.go` 中写入代码

   ```go
   package bench
   
   import "testing"
   
   func BenchmarkSelect(b *testing.B) {
   	InitServerIndex()
   	b.ResetTimer()
   
   	for i := 0; i < b.N; i++ {
   		Select()
   	}
   }
   
   // parallel 平行的，同时发生的
   func BenchmarkSelectParallel(b *testing.B) {
   	InitServerIndex()
   	b.ResetTimer()
   	b.RunParallel(func(pb *testing.PB) {
   		for pb.Next() {
   			Select()
   		}
   	})
   }
   ```

5. 最终代码见[src/2-2-12bench/](../src/2-2-12bench/)

6. 运行命令 `go test -bench=Ben`，查看效果。指令中`Ben`是指以`Ben`开头的测试函数都执行

   ```shell
   (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-2-12bench> go test -bench=B
   goos: windows
   goarch: amd64
   pkg: server
   cpu: 12th Gen Intel(R) Core(TM) i5-12490F
   BenchmarkSelect-12              71638796                15.41 ns/op
   BenchmarkSelectParallel-12      28514331                42.61 ns/op
   PASS
   ok      server  2.562s
   
   # 只测并行
   (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-2-12bench> go test -bench=BenchmarkSelectP
   goos: windows
   goarch: amd64
   pkg: server
   cpu: 12th Gen Intel(R) Core(TM) i5-12490F
   BenchmarkSelectParallel-12      27168068                43.41 ns/op
   PASS
   ok      server  1.399s
   ```




### 案例13 - 基准测试优化

1. 复制[案例12](###案例12 - 基准测试) 的 `2-2-12bench` 文件夹，重命名为  `2-2-13benchOptimize`

2. 在 `2-2-13bencOptimize`文件夹中执行命令， `go get "github.com/bytedance/gopkg@main"` 添加字节的gopkg

   ```shell
   (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-2-13benchOptimize> go get "github.com/bytedance/gopkg@main"
   go: downloading github.com/bytedance/gopkg v0.0.0-20220118071334-3db87571198b
   go: added github.com/bytedance/gopkg v0.0.0-20220118071334-3db87571198b
   ```

3. 在文件 `server.go`中导入字节的fastrand包

   ```go
   import "github.com/bytedance/gopkg/lang/fastrand"
   ```

   这里说一下是怎么知道在路径 `gopkg/lang/fastrand`

   方法一：

   - 当你执行go get 命令后，你可以去 `$GOPATH/pkg/mod`找到你下载的包，然后在里面查找就能找到`fastrand`包的位置
   - 如果你不知道你的 `$GOPATH`在哪，你可以在终端输入`go env` 确定

   方法二（其实我也是用方法二的时候才确定用方法一再找一遍的）：

   - 我直接看课件的实例代码 [go-project-example/benchmark/load_balance_selector.go at V0 · Moonlight-Zhao/go-project-example (github.com)](https://github.com/Moonlight-Zhao/go-project-example/blob/V0/benchmark/load_balance_selector.go)

4. 修改下方的随机选择服务器代码中的随机函数为快速随机函数

   ```go
   // 原来
   return ServerIndex[rand.Intn(10)]
   
   // 现在
   return ServerIndex[fastrand.Intn(10)]
   ```

5. 与[案例12](###案例12 - 基准测试)一样进行测试

   ```shell
   (base) PS D:\code\MoFishXiaodui\ExecutableManual\src\2-2-13benchOptimize> go test -bench=B
   goos: windows
   goarch: amd64
   pkg: server
   cpu: 12th Gen Intel(R) Core(TM) i5-12490F
   BenchmarkSelect-12              523615494                2.292 ns/op
   BenchmarkSelectParallel-12      1000000000               0.3863 ns/op
   PASS
   ok      server  2.039s
   ```

6. 可以发现这比[案例12](###案例12 - 基准测试)要快很多



## 04 实践

### 案例14 - 克隆老师的社区话题页面案例进行测试

因为这是一个完整的仓库，且此次案例不需要我们写代码，所以这里只放链接，就不把仓库代码放到这里了。

[Moonlight-Zhao/go-project-example at V0 (github.com)](https://github.com/Moonlight-Zhao/go-project-example/tree/V0)

注意，我们**只用分支 V0** 上面的代码，主分支的代码需要用到数据库，不方便初学者测试。大家主要克隆代码时只克隆V0上的代码就行。

你可以克隆主分支后再切到V0分支，然后拉取V0分支。

也可以直接只克隆V0分支

#### 附上GPT的解答：

问：git clone 如何只克隆特定分支？

ChatGPT：

要克隆一个特定分支而不是整个仓库，可以使用Git的`-b`选项来指定要克隆的分支。以下是在`git clone`命令中使用`-b`选项的示例：

```bash
git clone -b branch_name repository_url
```

其中：

- `branch_name`是你要克隆的特定分支的名称。
- `repository_url`是你要克隆的远程仓库的URL。

例如，如果你想克隆一个名为`development`的特定分支，仓库的URL为`https://github.com/example/repo.git`，你可以执行以下命令：

```bash
git clone -b development https://github.com/example/repo.git
```

这将只克隆`development`分支的内容到你的本地目录中，而不会包含其他分支和提交历史。

#### 开始实操

1. 找一个合适的位置，执行 `git clone -b V0 git@github.com:Moonlight-Zhao/go-project-example.git `

2. 然后在项目根目录上执行`go run server.go`

3. 当看到终端输出`[GIN-debug] Listening and serving HTTP on :8080`时，即为启动成功。

4. 先去浏览器查看`localhost:8080`，可以看到404，这是因为代码没有配置根路由的处理函数，默认404了。

5. 为了找到有用的路由，我们可以到`server.go`中查看gin的路由处理函数

   ```go
   r.GET("/community/page/get/:id", func(c *gin.Context) {
       topicId := c.Param("id")
       data := cotroller.QueryPageInfo(topicId)
       c.JSON(200, data)
   })
   ```

6. 从上面的代码中，可以看到gin引擎r可以处理路径为 `/community/page/get/:id`的GET请求。于是尝试在浏览器访问：[localhost:8080/community/page/get/1](http://localhost:8080/community/page/get/1)，既可以看到响应数据，然后请自行把id换成2和3等进行测试。

7. 这些数据来源于根目录data文件夹下的topic文件。请尝试修改以及增加数据。

8. 当你修改了文件夹的数据后，刷新浏览器并不能直接看到新数据，需要你关掉服务再执行命令`go run server.go`开启服务













