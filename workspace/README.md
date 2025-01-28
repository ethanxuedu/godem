Go 工作区是 Go 1.18 引入的一个新特性，允许你在一个目录中管理多个模块。

## 1. 创建 Go 工作区

首先，创建一个目录作为你的工作区：

```
mkdir my-workspace
cd my-workspace
```

## 2. 初始化工作区

在工作区目录中，初始化 Go 工作区：

```api
go work init
```

这会在当前目录下生成一个 go.work 文件。

## 3. 创建模块

在工作区中创建两个模块：

```api
mkdir module1
cd module1
go mod init example.com/module1
cd ..

mkdir module2
cd module2
go mod init example.com/module2
cd ..
```

## 4. 添加模块到工作区

将这两个模块添加到工作区中：

```api
go work use ./module1
go work use ./module2
```

这会在 go.work 文件中添加这两个模块的路径。

## 5. 编写代码

在 module1 中创建一个简单的 Go 文件：

```api
cd module1
echo 'package module1

import "fmt"

func Hello() {
fmt.Println("Hello from module1")
}' > module1.go
```

在 module2 中创建一个 Go 文件，并调用 module1 中的函数：

```api
cd ../module2
echo 'package main

import (
"example.com/module1"
)

func main() {
module1.Hello()
}' > main.go
```

## 6. 运行代码

在 module2 中运行代码：

```api
go build

module2.exe
```

你应该会看到输出：

```api
Hello from module1
```

## 7. go.work 文件内容

go.work 文件的内容应该类似于：

```api
go 1.20

use (
./module1
./module2
)
```


## 8. 总结

通过 Go 工作区，你可以在一个目录中管理多个模块，并且这些模块可以相互引用。这对于开发大型项目或多个相关项目时非常有用。

## 9. 清理

如果你想清理工作区，只需删除 go.work 文件即可：

```api
rm go.work
```

这样，Go 将不再将这些模块视为工作区的一部分。