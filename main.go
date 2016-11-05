
package main

import (
    "io"
    "log"
    "net/http"
    "strings"
)

func logger( fn http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    log.Println("Before")
    fn(w, r)
    log.Println("After")
  }
}

func MustParams(fn http.HandlerFunc, params ...string) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    var missingParams []string
    for _, param := range params {
      if len(r.URL.Query().Get(param)) == 0 {
        missingParams = append(missingParams, param)
      }
    }

    if len(missingParams) > 0 {
      http.Error(w, "missing " + strings.Join(missingParams, ","), http.StatusBadRequest)
      return
    }

    fn(w, r)
  }
}
func pingHandler(w http.ResponseWriter, r *http.Request){
    io.WriteString(w, "PONG\n")
}

func main() {
    // http.HandleFunc("/ping", MustParams(logger(pingHandler), "auth", "key"))
    // log.Fatal(http.ListenAndServe(":8081", nil))
}

