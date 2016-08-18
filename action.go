package launch
import (
    "net/http"
)
type Action interface {
    Execute(http.ResponseWriter, *http.Request)
}
