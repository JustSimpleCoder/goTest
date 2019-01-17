package db

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "goPush/K12User/config"
)

var Vpsdb *sqlx.DB

func init() {

    var link_str string = config.Vpsdb.User + ":" + config.Vpsdb.Pwd + "@tcp(" + config.Vpsdb.IP + ":" + config.Vpsdb.Port + ")/" + config.Vpsdb.Name + "?charset=" + config.Vpsdb.Charset

    xdf_woxue, err := sqlx.Open("mysql", link_str)
    if err != nil {
        fmt.Println("Open POP Database Failed,", err)
        return
    }

    Vpsdb = xdf_woxue

    Vpsdb.SetMaxIdleConns(10)
}

func GetHomeworkByUser(s string) (hus []HomeworkUser) {

    err := Vpsdb.Select(&hus, "SELECT homework_id,school_id,class_code,stu_num,create_time FROM t_homework_user where stu_num = ?", s)
    if err != nil {
        fmt.Println("exec failed GetHomeworkByUser, ", err, "SELECT homework_id,school_id,class_code,stu_num,create_time FROM t_homework_user where stu_num = ?", s)
        return
    }

    return
}

func GetHomworkInfo(id int) (h Homework) {

    err := Vpsdb.Get(&h, "SELECT id,class_code,school_id,jdxk_id,is_homework_user,create_time,creator_id FROM t_homework where id = ?", id)
    if err != nil {
        fmt.Println("exec failed GetHomworkInfo, ", err, "SELECT id,class_code,school_id,jdxk_id,is_homework_user,create_time,creator_id FROM t_homework where id = ?", id)
        return
    }

    if h.SchoolId.Valid {
        h.SchoolIdStr = h.SchoolId.String
    } else {
        h.SchoolIdStr = ""
    }

    return
}

func GetStatusByUser(s string) (sus []StatusUser) {

    err := Vpsdb.Select(&sus, "SELECT status_id,stu_num,class_code,school_id FROM t_status_student where stu_num = ?", s)
    if err != nil {
        fmt.Println("exec failed GetStatusByUser, ", err, "SELECT status_id,stu_num,class_code,school_id FROM t_status_student where stu_num = ?", s)
        return
    }

    for k, th := range sus {

        if th.SchoolId.Valid {
            sus[k].SchoolIdStr = th.SchoolId.String
        } else {
            sus[k].SchoolIdStr = ""
        }
    }

    return
}

func GetStatusInfo(id int) (s Status) {

    err := Vpsdb.Get(&s, "SELECT id,post_id,status_type,create_time,creator_id FROM t_status where id = ?", id)
    if err != nil {
        fmt.Println("exec failed GetStatusInfo, ", err, "SELECT post_id,status_type,create_time,creator_id FROM t_status where id = ?", id)
        return
    }

    return
}

func GetReportInfo(pid string) (r T_app_quiz_result) {

    err := Vpsdb.Get(&r, "SELECT stu_number,class_code,school_id,creater_id,indoor_outdoor FROM t_app_quiz_result where id = ?", pid)
    if err != nil {
        fmt.Println("exec failed GetReportInfo, ", err, "SELECT stu_number,class_code,school_id,creater_id,indoor_outdoor FROM t_app_quiz_result where id = ?", pid)
        return
    }

    return
}

func GetClassInfo(s, c string) (i ClassInfo) {

    err := Vpsdb.Get(&i, "SELECT class_code,school_id,subject FROM t_class_info where school_id = ? and class_code = ?  ", s, c)
    if err != nil {
        fmt.Println("exec failed GetClassInfo , ", err, "SELECT class_code,school_id,subject FROM t_class_info where school_id = ? and class_code = ?  ", s, c)
        return
    }

    return
}

func GetObject(id int) (i ClassSubject) {

    err := Vpsdb.Get(&i, "SELECT id,subject_code FROM t_subject where id = ?", id)
    if err != nil {
        fmt.Println("exec failed GetObject ,", err, "SELECT id,subject_code FROM t_subject where id = ?", id)
        return
    }

    return

}

func GetExaminate(id string) (e T_app_examination) {

    err := Vpsdb.Get(&e, "SELECT id,exam_name FROM t_app_examination where id = ?", id)
    if err != nil {
        fmt.Println("exec failed GetExaminate, ", err, "SELECT id,exam_name FROM t_app_examination where id = ?", id)
        return
    }

    return

}

func GetTeacherHomework(uid string) (ths []Homework) {

    err := Vpsdb.Select(&ths, "SELECT id,class_code,school_id,jdxk_id,is_homework_user,create_time,creator_id FROM t_homework where creator_id = ? ", uid)
    if err != nil {
        fmt.Println("exec failed GetTeacherHomework, ", err, "SELECT id,class_code,school_id,jdxk_id,is_homework_user,create_time,creator_id FROM t_homework where creator_id = ? ", uid)
        return
    }

    for k, th := range ths {

        if th.SchoolId.Valid {
            ths[k].SchoolIdStr = th.SchoolId.String
        } else {
            ths[k].SchoolIdStr = ""
        }
    }

    return
}

func GetTeacherStatus(uid string) (tss []Status) {

    err := Vpsdb.Select(&tss, "SELECT id,post_id,status_type,create_time,creator_id FROM t_status where creator_id = ? ", uid)
    if err != nil {
        fmt.Println("exec failed GetTeacherStatus, ", err, "SELECT id,post_id,status_type,create_time,creator_id FROM t_status where creator_id = ? ", uid)
        return
    }

    return
}

func GetStatusOneStudent(id int) (s StatusUser) {

    err := Vpsdb.Get(&s, "SELECT status_id,stu_num,class_code,school_id FROM t_status_student where status_id = ?", id)
    if err != nil {
        fmt.Println("exec failed GetStatusOneStudent, ", err, "SELECT status_id,stu_num,class_code,school_id FROM t_status_student where status_id = ?", id)
        return
    }

    if s.SchoolId.Valid {
        s.SchoolIdStr = s.SchoolId.String
    } else {
        s.SchoolIdStr = ""
    }

    return
}

type Homework struct {
    Id          int            `db:"id"`
    ClassCode   string         `db:"class_code"`
    SchoolId    sql.NullString `db:"school_id"`
    SchoolIdStr string
    JDXKID      sql.NullInt64 `db:"jdxk_id"`
    IsUser      int           `db:"is_homework_user"`
    CreateTime  string        `db:"create_time"`
    TeacherUid  string        `db:"creator_id"`
}

type Status struct {
    StatusId   int    `db:"id"`
    PostId     string `db:"post_id"`
    StatusType string `db:"status_type"`
    CreateTime string `db:"create_time"`
    TeacherUid string `db:"creator_id"`
}

type HomeworkUser struct {
    HomeworkId  int            `db:"homework_id"`
    SchoolId    sql.NullString `db:"school_id"`
    SchoolIdStr string
    ClassCode   string `db:"class_code"`
    StudentCode string `db:"stu_num"`
    CreateTime  string `db:"create_time"`
}

type StatusUser struct {
    StatusId    int            `db:"status_id"`
    StudentCode string         `db:"stu_num"`
    ClassCode   string         `db:"class_code"`
    SchoolId    sql.NullString `db:"school_id"`
    SchoolIdStr string
}

type T_app_quiz_result struct {
    Id          string `db:"id"`
    StudentCode string `db:"stu_number"`
    ClassCode   string `db:"class_code"`
    SchoolId    string `db:"school_id"`
    InOutType   string `db:"indoor_outdoor"`
    TeacherUid  string `db:"creater_id"`
}

type ClassInfo struct {
    ClassCode string `db:"class_code"`
    SchoolId  string `db:"school_id"`
    Subject   int    `db:"subject"`
}

type ClassSubject struct {
    Id   int    `db:"id"`
    Code string `db:"subject_code"`
}

type T_app_examination struct {
    Id   int    `db:"id"`
    Name string `db:"exam_name"`
}

func TT() {
    fmt.Println(Vpsdb)
}
