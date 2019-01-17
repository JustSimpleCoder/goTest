package service

import (
    "fmt"
    "goPush/K12User/api"
    "goPush/K12User/business"
    "goPush/K12User/db"
    "time"
)

var chanUser = make(chan db.WaitingData, 10)

func init() {

    go func() {
        for {
            select {

            case u := <-chanUser:
                time.Sleep(time.Millisecond * 100)
                go syncOneUser(u)
            default:

            }
        }
    }()
}

func ProcessUser() {

    fmt.Println("processUser")

    ticker := time.NewTicker(time.Second * 5)
    for {

        select {

        case <-ticker.C:

            waiting := db.GetToProcessUser(10)

            // fmt.Println(waiting)

            for _, stu := range waiting {

                if stu.Code.Valid {
                    stu.StudentCode = stu.Code.String
                } else {
                    stu.StudentCode = ""
                }

                chanUser <- stu

            }

        }

    }
}

func syncOneUser(u db.WaitingData) {

    fmt.Println("Begin Sync User %s %s", u.UserId, u.StudentCode)

    if u.StudentCode == "" {

        business.TeacherHomework(u)
        business.TeacherStatus(u)

    } else {

        c := api.GetStuClass(u)
        business.SyncHomework(u, c)
        business.SyncStatus(u, c)

    }

    fmt.Println("End Sync User %s %s", u.UserId, u.StudentCode)
}
