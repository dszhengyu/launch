package launch
import (
    "net/http"
    "net/url"
    "log"
    "encoding/json"
    "fmt"
    "html/template"
    "os"
)

type InputType url.Values
type OutputType map[string]interface{}

type Action interface {
    Execute(InputType, OutputType) error
}

type WrapAction struct {
    action Action
}

/*this function should do some basic control
*such as:
*1. cookie
*2. template explain
*3. session
*4. parse input and output 
*/
func (wa WrapAction)Execute(w http.ResponseWriter, r *http.Request) {
    if wa.action == nil {
        log.Panicln("no action")
        return
    }
    r.ParseForm()
    log.Printf("receive input:[%+v]", r.Form)
    output := make(OutputType)
    err := wa.action.Execute(InputType(r.Form), output)
    if err != nil {
        log.Println(err)
        return
    }
    log.Printf("output:[%+v]", output)
    //to print template or just output json data
    tplName, exist := output["template"]
    if exist == true {
        tplNameStr, ok := tplName.(string)
        if ok != true {
            log.Panicln("error template name")
            return
        }
        goPath := os.Getenv("GOPATH")
        tpl, err := template.ParseFiles(goPath + "/src/templates/" + tplNameStr)
        if err != nil {
            log.Panicln(err)
            return
        }
        tpl.Execute(w, output)
    } else {
        log.Printf("output json")
        outputByte, err := json.Marshal(output)
        if err != nil {
            log.Panicln(err)
            return
        }
        outputStr := string(outputByte)
        fmt.Fprint(w, outputStr)
    }
    return
}
