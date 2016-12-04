package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Returnable struct {
    Url, Method, Proto, Body, Host string
    Headers http.Header
}

func main() {
    http.HandleFunc("/", api)
    http.ListenAndServe(":8000", nil)
}

func api(w http.ResponseWriter, r *http.Request){
    var returner string

    buf := new(bytes.Buffer)
    buf.ReadFrom(r.Body)

    var returnable Returnable
    returnable.Body = buf.String()
    returnable.Headers = r.Header
    returnable.Host = r.Host
    returnable.Method = r.Method
    returnable.Proto = r.Proto
    returnable.Url = r.URL.String()

    if rData,err := json.Marshal(returnable); err != nil {
        returner = err.Error()
    } else {
        returner = string(rData)
    }

    log.Println(returner)

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, returner)
}
