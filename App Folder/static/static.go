package static

import (
    "net/http"
)

func init() {
    http.HandleFunc("/", static)
}

func static(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "files/"+r.URL.Path)
}
