# HTTP协议

# Manual

## 01再谈HTTP协议

### 案例4-1创建一个gin服务

这部分ppt没有讲，这里先做这个服务，为下一个案例做准备

1. 新建文件夹 `4-1gin`，作为项目根目录

2. 根目录上运行 `go mod init ginServer`

3. 运行 `go get -u github.com/gin-gonic/gin`加载包

4. 在根目录创建 `main.go`

5. 在main.go中写入

   ```go
   package main
   
   import "github.com/gin-gonic/gin"
   
   func main() {
   	r := gin.Default()
   	r.GET("/ping", func(c *gin.Context) {
   		c.JSON(200, gin.H{
   			"message": "pong",
   		})
   	})
   	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
   }
   ```

6. 直接运行

7. 到浏览器查看 `localhost:8080/ping` 的结果，若能看到 `{"message":"pong"}`则说明服务启动成功

### 案例4-2发送报文

本次案例采用telnet发送报文作为测试

1. 启动案例4-1的gin服务

2. cmd运行 `telnet localhost 8080`，会显示出![image-20230811111750235](course4http协议.assets/image-20230811111750235.png)

3. 如果telnet没有安装，需要自行安装，参考下文附录[Windows安装telnet](##Windows安装telnet)

4. 如果如前面所示显示会话框等待输入，那么可以进行一下操作：

5. 首先按下快捷键 `ctrl+]` 打开回显，然后再按回车回到会话框。如果你没有开回显的话，你在会话框输入的内容会直接传到服务器上，但没有显示出来，你不能确定自己输入了什么

6. 接下来输入以下语句，注意，不要敲错，不支持回车删除，你输入的每一个字符都是会马上传到服务器

   ```http
   GET /ping HTTP/1.1
   host: localhost
   
   ```

   输入完指令敲两个回车（第一次的回车是为了标识分开headers信息和body信息，第二次回车就结束报文了）

7. 然后可以看到gin服务传回来的json数据

   ```cmd
   GET /ping HTTP/1.1
   host: localhost
   
   HTTP/1.1 200 OK
   Content-Type: application/json; charset=utf-8
   Date: Fri, 11 Aug 2023 03:29:20 GMT
   Content-Length: 18
   
   {"message":"pong"}
   ```

### 案例4-3发送post报文

1. 复制案例4-1的文件夹，改名为4-3postGin

2. 在main.go中新增一个post请求的处理方法

   ```go
   r.POST("/sis", func(c *gin.Context) {
       c.Data(200, "text/plain; charset=utf-8", []byte("ok"))
   })
   ```

3. 直接运行

4. cmd执行`telnet localhost 8080`进入会话框，关闭回显，发送报文如下

   ```http
   POST /sis HTTP/1.1
   host: 127.0.0.1
   ```

5. 测试结果

   ```shell
   POST /sis HTTP/1.1	# 请求行
   host: 127.0.0.1		# 请求头部 host
   
   HTTP/1.1 200 OK		# 响应行
   Content-Type: text/plain; charset=utf-8	# 响应头部 内容/文件类型
   Date: Fri, 11 Aug 2023 04:02:53 GMT		# 响应头部 响应时间
   Content-Length: 2						# 响应头部 主体长度
   
   ok		# 响应主体
   ```

   





# 附录

## Windows安装telnet

1. 在设置查看telnet![image-20230811110634238](course4http协议.assets/image-20230811110634238.png)

2. 选择`启用或关闭Windows功能`

3. 在列表勾选`Telnet客户端`，点击确定，等待安装完毕![image-20230811110730459](course4http协议.assets/image-20230811110730459.png)

4. 此时在cmd运行`telnet`可以看到 如下内容，说明telnet启动成功

   > 欢迎使用 Microsoft Telnet Client
   >
   > Escape 字符为 'CTRL+]'
   >
   > Microsoft Telnet>



## HTTP 请求头

- 标准文档 [RFC 9110: HTTP Semantics (rfc-editor.org)](https://www.rfc-editor.org/rfc/rfc9110)
- MDN文档 [HTTP 标头（header） - HTTP | MDN (mozilla.org)](https://developer.mozilla.org/zh-CN/docs/web/http/headers)





