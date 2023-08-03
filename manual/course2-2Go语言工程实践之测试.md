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













