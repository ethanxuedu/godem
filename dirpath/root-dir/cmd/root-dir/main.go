package main

import "godem/dirpath/root-dir/cmd/configs"

func main() {
    // configs.GetWorkPath() // C:\workplace\go_workplace\src\godem
    // configs.GetWorkPathByArg()
    configs.GetWorkPathByExec()
    // configs.GetWorkPathByCaller()
    // configs.GetWorkPathByEnv()
}
