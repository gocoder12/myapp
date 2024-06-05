package main

import (
    "log"
    "myapp/internal/router"
    "myapp/pkg/util"
    "net/http"
)

func main() {
    config := util.LoadConfig()
    r := router.NewRouter()

    log.Printf("Server starting on port %s", config.Server.Port)
    log.Fatal(http.ListenAndServe(config.Server.Port, r))
}
