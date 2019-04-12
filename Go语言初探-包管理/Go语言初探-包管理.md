golang语言初探-包管理

Goroot path：`C:\Go\bin`

Gopath :`C:\Users\zyf\go`

`go get `命令的包全部都保存在`Gopath`中

```go
package main
import (
    "hello2"
)
func main(){
    //注意这里变量名是hello
    hello.Say()
}

```
实际上导入的是目录
`C:\Users\zyf\go\src`下的`hello2`文件夹

导入的包的默认变量名是什么呢？

`C:\Users\zyf\go\src\hello2`里只用一个文件`main.go`
```go
package hello
import "fmt"
func Say(){
    fmt.Println("Hello, World!")
}
```
这里的package为hello，表示这个文件为hello包的内容。
Say方法大写，暴露在外面。所以推荐把package名和文件夹名命名一致，这样方便引用。即导入的文件夹名就是package的名字

golang包管理如果是`go get`的话也是先下载在gopath下，然后在本地读取比如


````go

package main
import (
    "hello2"
    FirstGoPackage "github.com/yukinotech/FirstGoPackage"
)
func main(){
    hello.Say()
    FirstGoPackage.Say("good bye")
}

```
