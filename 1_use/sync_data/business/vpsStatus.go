package business

import (
    "fmt"
    "goPush/K12User/api"
    "goPush/K12User/config"
    "goPush/K12User/db"
    "net/url"
    // "os"
)

func SyncStatus(u db.WaitingData, c api.ClassAPI) {

    userStatus := db.GetStatusByUser(u.StudentCode)

    for _, us := range userStatus {

        sInfo := db.GetStatusInfo(us.StatusId)

        if sInfo.StatusType == "1" {

            syncReport(sInfo, us)

        } else {

            syncMaterial(sInfo, us)

        }

    }

}

func getReportMsgId(si db.Status, us db.StatusUser) (key string) {

    // VPS_0_消息类型（成绩报告29，知识点分析30）_教师id_时间戳_学校id_班级编码_subjectCode?_测试类型?

    sc := getSubjectCode(us)

    tt := getTestType(si.PostId)

    t := stimeToUnix(si.CreateTime)
    key = fmt.Sprintf("VPS_0_29_%s_%d_%s_%s_%s_%s", si.TeacherUid, t, us.SchoolIdStr, us.ClassCode, sc, tt)

    return
}

func getSubjectCode(us db.StatusUser) string {

    c := db.GetClassInfo(us.SchoolIdStr, us.ClassCode)

    s := db.GetObject(c.Subject)

    return s.Code
}

func getTestType(postid string) string {

    ri := db.GetReportInfo(postid)

    ti := db.GetExaminate(ri.InOutType)
    u := url.Values{}
    u.Set("a", ti.Name)
    s := u.Encode()
    return string([]byte(s)[2:])
}

func syncReport(si db.Status, us db.StatusUser) {

    msgId := getReportMsgId(si, us)
    RedisToZAdd(keyForStudent(us.StudentCode, config.ZstuMsgType["r"]), msgId, si.CreateTime)

}

func syncMaterial(s db.Status, us db.StatusUser) {

    if s.StatusId > 0 {

        msgId := GetMsgId(0, 28, s.StatusId)

        RedisToZAdd(keyForClass(us.SchoolIdStr, us.ClassCode), msgId, s.CreateTime)
        RedisToZAdd(keyForClassOld(us.SchoolIdStr, us.ClassCode), msgId, s.CreateTime)

        RedisToZAdd(keyForStudent(us.StudentCode, config.ZstuMsgType["m"]), msgId, s.CreateTime)
    }

}

func TeacherStatus(u db.WaitingData) {

    tss := db.GetTeacherStatus(u.UserId)

    for _, ts := range tss {

        us := db.GetStatusOneStudent(ts.StatusId)

        if us.StatusId > 0 {

            RedisToZAdd(TeacherReadKey(ts.TeacherUid), TeacherReadVal(us.SchoolIdStr, us.ClassCode), "0")

            msgId := GetMsgId(0, 28, ts.StatusId)

            RedisToZAdd(TeacherMsgKey(ts.TeacherUid, us.SchoolIdStr, us.ClassCode), msgId, ts.CreateTime)
        }

    }

}
