# 在 Go 项目中，获取项目的根目录可以通过以下几种方式实现：

## 1. 使用 os.Getwd 获取当前工作目录

os.Getwd 可以获取当前工作目录，通常这是项目的根目录（前提是你从根目录运行程序）。

```api
package main

import (
    "fmt"
    "os"
)

func main() {
    dir, err := os.Getwd()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Current directory:", dir)
}
```

## 2. 使用 os.Executable 获取可执行文件路径

os.Executable 返回当前可执行文件的路径，你可以通过解析路径来获取项目的根目录。

```api
package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    exePath, err := os.Executable()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    rootDir := filepath.Dir(filepath.Dir(exePath)) // 假设项目根目录是可执行文件的上一级目录
    fmt.Println("Root directory:", rootDir)
}
```

## 3. 使用 go build 的 -ldflags 注入根目录

在构建时，可以通过 -ldflags 注入根目录路径。

```api
package main

import (
    "fmt"
)

var rootDir string

func main() {
    fmt.Println("Root directory:", rootDir)
}

构建时注入路径：

bash
复制
go build -ldflags="-X 'main.rootDir=$(pwd)'" -o myapp
```

## 4. 使用环境变量

通过环境变量传递项目根目录。

```api
package main

import (
    "fmt"
    "os"
)

func main() {
    rootDir := os.Getenv("PROJECT_ROOT")
    if rootDir == "" {
        fmt.Println("PROJECT_ROOT environment variable not set")
        return
    }
    fmt.Println("Root directory:", rootDir)
}
```

运行程序时设置环境变量：
```api
PROJECT_ROOT=$(pwd) ./myapp
```

## 5. 使用 go.mod 文件定位

如果你的项目使用了 Go Modules，可以通过解析 go.mod 文件来定位项目根目录。

```api
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func findModuleRoot(dir string) (string, error) {
    for {
        if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
            return dir, nil
        }
        parent := filepath.Dir(dir)
        if parent == dir {
            return "", fmt.Errorf("go.mod not found")
        }
        dir = parent
    }
}

func main() {
    rootDir, err := findModuleRoot(".")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Root directory:", rootDir)
}
```

## 总结
- 如果你从项目根目录运行程序，os.Getwd 是最简单的方式。
- 如果需要更精确的定位，可以使用 go.mod 文件或环境变量。
- 对于复杂的项目结构，可能需要结合多种方法来确定根目录。