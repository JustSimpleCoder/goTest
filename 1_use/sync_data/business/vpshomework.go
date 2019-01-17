package business

import (
    "fmt"
    "goPush/K12User/api"
    "goPush/K12User/config"
    "goPush/K12User/db"
    // "os"
)

func SyncHomework(u db.WaitingData, c api.ClassAPI) {

    // just student

    homeworkusers := db.GetHomeworkByUser(u.StudentCode)

    fmt.Println(homeworkusers)

    for _, uh := range homeworkusers {

        homework := db.GetHomworkInfo(uh.HomeworkId)

        // fmt.Println(homework)

        // os.Exit(1)

        if homework.Id > 0 {

            var jdxk = 28

            if homework.JDXKID.Valid {
                jdxk = 27
            }

            msgId := GetMsgId(1, jdxk, homework.Id)

            if homework.IsUser == 0 {

                RedisToZAdd(keyForClass(homework.SchoolIdStr, homework.ClassCode), msgId, homework.CreateTime)
                RedisToZAdd(keyForClassOld(homework.SchoolIdStr, homework.ClassCode), msgId, homework.CreateTime)

            }

            RedisToZAdd(keyForStudent(u.StudentCode, config.ZstuMsgType["h"]), msgId, homework.CreateTime)

        }
    }

}

func TeacherHomework(u db.WaitingData) {

    ths := db.GetTeacherHomework(u.UserId)

    for _, th := range ths {

        RedisToZAdd(TeacherReadKey(th.TeacherUid), TeacherReadVal(th.SchoolIdStr, th.ClassCode), "0")

        var jdxk = 28

        if th.JDXKID.Valid {
            jdxk = 27
        }

        msgId := GetMsgId(1, jdxk, th.Id)

        // fmt.Println("==========================", msgId, "======================")

        RedisToZAdd(TeacherMsgKey(th.TeacherUid, th.SchoolIdStr, th.ClassCode), msgId, th.CreateTime)

    }

}
