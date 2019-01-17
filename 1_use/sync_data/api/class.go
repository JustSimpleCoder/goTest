package api

import (
    "github.com/json-iterator/go"

    "fmt"
    "goPush/K12User/config"
    "goPush/K12User/db"
    "io/ioutil"
    "net/http"
    // "os"
)

type ClassAPI []struct {
    ClassState int `json:"classState"`
    Classes    []struct {
        SchoolID       int    `json:"schoolId"`
        ClassName      string `json:"className"`
        ClassCode      string `json:"classCode"`
        DeptCode       string `json:"deptCode"`
        SubjectCode    string `json:"subjectCode"`
        SubjectName    string `json:"subjectName"`
        ClassStartTime string `json:"classStartTime"`
        ClassEndTime   string `json:"classEndTime"`
        StudentCount   int    `json:"studentCount"`
        LessonCount    int    `json:"lessonCount"`
        Teachers       []struct {
            UserID      string `json:"userId"`
            TeacherCode string `json:"teacherCode"`
            TeacherName string `json:"teacherName"`
        } `json:"teachers"`
        PrintAddress string `json:"printAddress,omitempty"`
    } `json:"classes"`
}

func GetStuClass(u db.WaitingData) (result ClassAPI) {

    // 暂不需要获取班级信息
    result = ClassAPI{}

    return

    url := config.ApiConf["url"] + "/1/student/classesInfo_byCode?classState=0&studentCode=" + u.StudentCode + "&accessToken=" + config.ApiConf["token"] + "&appId=" + config.ApiConf["appid"]
    res, err := http.Get(url)
    defer res.Body.Close()

    data, err := ioutil.ReadAll(res.Body)

    if err != nil {
        fmt.Println(err)
        return
    }
    result = ClassAPI{}

    var json = jsoniter.ConfigCompatibleWithStandardLibrary
    json.Unmarshal(data, &result)

    // fmt.Println(result)

    return
    // os.Exit(1)
}
