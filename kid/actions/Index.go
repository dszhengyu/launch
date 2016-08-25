package actions
import (
    . "github.com/dszhengyu/launch"
)

type Index struct {
}

func (action Index)Execute(input InputType, output OutputType)(err error) {
    for key, value := range input {
        output[key] = value[0]
    }
    return nil
}
