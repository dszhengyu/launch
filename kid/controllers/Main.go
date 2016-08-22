package controllers
import (
    "github.com/dszhengyu/launch"
    "actions"
)

func init() {
    var Main = make (launch.Controller)
    Main["/"] = actions.Index{}
    launch.RegistController(Main)
}
