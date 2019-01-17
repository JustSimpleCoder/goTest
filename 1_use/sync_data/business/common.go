package business

import (
    "fmt"
    "goPush/K12User/api"
    "goPush/K12User/config"
    "goPush/K12User/db"
    "time"
)

func init() {

    db.NewSentinelPool(config.GetRedisDb(3))

}

func ClassList(c api.ClassAPI) (res []map[string]string) {

    //json -> 标准数组
    _ = c

    // classList := [["id":"1","classCode":"bj"]]
    // var tmp  = make( []map[string][string],10)
    var tmp = map[string]string{"id": "1"}

    res = append(res, tmp)
    return
}

func RedisToZAdd(redisKey, msgId, stime string) {

    if redisKey != "" && msgId != "" {

        conn := db.GetConn()

        defer conn.Close()

        unix_time := stimeToUnix(stime)

        fmt.Println(redisKey, msgId, unix_time)
        _, err := conn.Do("ZADD", redisKey, unix_time, msgId)

        if err != nil {
            fmt.Println(err)
        }

    }
}

// func RedisToClass(schoolid)

func GetMsgId(ishomework, jdxk, id int) string {
    return fmt.Sprintf("VPS_%d_%d_%d", ishomework, jdxk, id)
}

func stimeToUnix(stime string) int64 {

    var unix_time int64
    unix_time = 0

    if stime != "0" {
        //返回现在时间

        loc, _ := time.LoadLocation("Asia/Chongqing")
        t, _ := time.ParseInLocation(config.BASETIME, stime, loc)
        unix_time = t.Unix()
    }

    fmt.Println(stime, unix_time)
    return unix_time
}

func TeacherReadKey(teacherid string) string {

    readKey := fmt.Sprintf("ppss_z_message_read_%s", teacherid)
    return readKey
}

func TeacherMsgKey(teacherid, schoolid, classcode string) string {

    key := ""
    if schoolid != "" {

        key = fmt.Sprintf("ppss_z_message_list_%s_%s_%s", teacherid, schoolid, classcode)
    }

    return key
}

func TeacherReadVal(schoolid, classcode string) string {

    readVal := ""
    if schoolid != "" {
        readVal = fmt.Sprintf("VPS_%s_%s", schoolid, classcode)
    }

    return readVal
}

func keyForStudent(s string, t int) string {

    return fmt.Sprintf("z_student_message_%s_%d", s, t)

}

func keyForClass(s, c string) string {

    return fmt.Sprintf("z_class_message_%s_%s", s, c)
}

func keyForClassOld(s, c string) string {

    return fmt.Sprintf("message_class_%s_%s", s, c)
}
