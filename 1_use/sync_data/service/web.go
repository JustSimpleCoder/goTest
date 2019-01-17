package service

import (
    "fmt"
    "goPush/K12User/db"
    "net/http"
)

func WebBegin() {

    fmt.Println("webBegin")

    http.HandleFunc("/countWaiting", countWaiting)

    err := http.ListenAndServe(":21918", nil)
    if err != nil {

        fmt.Println(" http failed")
    }
}

func countWaiting(w http.ResponseWriter, r *http.Request) {

    cnt := db.GetWaitingData()
    fmt.Fprintf(w, "countWaiting = %d", cnt)

}
