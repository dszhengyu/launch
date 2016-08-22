package launch
import (
    "encoding/json"
    "os"
    "log"
    "io/ioutil"
)

type Configuration struct {
    Name string
    Port string
}

func ResolveConfiguration() (conf Configuration) {
    goPath := os.Getenv("GOPATH")
    fileName := goPath + "/src/conf.json"
    confFile, err := os.Open(fileName)
    defer confFile.Close()
    if err != nil {
        log.Panicln(err)
    }
    origin, err := ioutil.ReadAll(confFile)
    if err != nil {
        log.Panicln(err)
    }
    err = json.Unmarshal(origin, &conf)
    if err != nil {
        log.Panicln(err)
    }
    return conf
}
