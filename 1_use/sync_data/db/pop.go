package db

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "goPush/K12User/config"
    // "os"
    "strconv"
    "strings"
)

var PopDB *sqlx.DB

func init() {

    var link_str string = config.POPdb.User + ":" + config.POPdb.Pwd + "@tcp(" + config.POPdb.IP + ":" + config.POPdb.Port + ")/" + config.POPdb.Name + "?charset=" + config.POPdb.Charset

    xdf_woxue, err := sqlx.Open("mysql", link_str)
    if err != nil {
        fmt.Println("Open XDF_WOXUE Database Failed,", err)
        return
    }

    PopDB = xdf_woxue

    PopDB.SetMaxIdleConns(10)

}

type WaitingData struct {
    Id          int64          `db:"id"`
    UserId      string         `db:"userId"`
    Code        sql.NullString `db:"studentCode"`
    StudentCode string
    Status      int `db:"status"`
}

func GetToProcessUser(l int) (waiting []WaitingData) {

    err := PopDB.Select(&waiting, "SELECT id,userId,studentCode,status from k12_user_sync where status = 0  limit ? ", l)
    if err != nil {
        fmt.Println("exec failed, ", err)
        return
    }

    // strings.Join(s1 []string, sep string) //：用sep把s1中的所有元素链接起来
    var strIds = ""
    for _, u := range waiting {
        strIds += strconv.FormatInt(u.Id, 10) + ","
    }

    strIds = strings.Trim(strIds, ",")

    if strIds != "" {
        PopDB.Exec("update k12_user_sync set status = 1 where id in (" + strIds + ")")
    }

    return
}

func GetWaitingData() int {

    var cnt int
    err := PopDB.Get(&cnt, "select count(1) as cnt from k12_user_sync where status = 0 ")
    if err != nil {
        fmt.Println(" count k12_user_sync ", err)
    }
    return cnt

}
