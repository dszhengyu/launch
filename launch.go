package launch
import (
    "net/http"
    "log"
)

type Launch struct {
}

func (launch *Launch) Run () {
    conf := ResolveConfiguration() 
    log.Printf("%+v", conf)
    err := http.ListenAndServe(":" + conf.Port, nil)
    if (err != nil) {
        log.Println(err)
        panic("ouch, listen fail")
    }
    return;
}
