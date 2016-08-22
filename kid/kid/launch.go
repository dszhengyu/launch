package main
import (
    "github.com/dszhengyu/launch"
    _ "controllers"
)

func main() {
    objLaunch := launch.Launch{}
    objLaunch.Run()
}
