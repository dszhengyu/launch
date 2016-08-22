package actions
import (
    "net/http"
    "fmt"
)

type Index struct {
}

func (action Index)Execute(w http.ResponseWriter, request *http.Request) {
    request.ParseForm()
    fmt.Fprintf(w, "hello world\n")
}
