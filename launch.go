package launch
import (
    "net/http"
    "fmt"
    "os"
)

type Launch struct {
}

func (launch *Launch) Run () {
    err := http.ListenAndServe(":8080", nil)
    if (err != nil) {
        fmt.Fprintf(os.Stderr, err.Error())
        panic("ouch, listen fail")
    }
    fmt.Println("launch...");
    return;
}
