package launch
import (
    "net/http"
)

type Controller map[string] Action

func RegistController(controller Controller) error {
    for path, action := range controller {
        http.HandleFunc(path, WrapAction{action}.Execute)
    }
    return nil
}
