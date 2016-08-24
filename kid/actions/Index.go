package actions
import (
    . "github.com/dszhengyu/launch"
)

type Index struct {
}

func (action Index)Execute(input InputType, output OutputType)(err error) {
    output["kaka"] = "hehe"
    output["list"] = []string{"i1","i2"}
    mapp := make(map[string]string)
    mapp["mapp"] = "mapppp"
    output["map"] = mapp
    return nil
}
