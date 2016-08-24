package launch
import (
    "net/http"
    "log"
    "encoding/json"
    "fmt"
    "net/url"
)

type InputType url.Values
type OutputType map[string]interface{}

type Action interface {
    Execute(InputType, OutputType) error
}

type WrapAction struct {
    action Action
}

//this function should do some basic control
//such as:
//1. cookie
//2. template explain
//3. session
//4. parse input and output 
func (wa WrapAction)Execute(w http.ResponseWriter, r *http.Request) {
    if wa.action == nil {
        log.Panicln("no action")
    }
    r.ParseForm()
    log.Printf("receive input:[%+v]", r.Form)
    output := make(OutputType)
    err := wa.action.Execute(InputType(r.Form), output)
    if err != nil {
        log.Println(err)
        return
    }
    outputByte, err := json.Marshal(output)
    outputStr := string(outputByte)
    log.Printf("output:[%v]", outputStr)
    fmt.Fprint(w, outputStr)
    return
}
