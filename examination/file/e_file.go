// 读取压缩的文件 gzip  if .json 结尾的文件 打印 统计key
// 一个文件统计文件中  字符 数字 空格的个数
// 实现一个自定义的错误和panic
// 使用map 和 struct

package main

import (
    // "compress/gzip"
    "bufio"
    "encoding/json"
    "fmt"
    "github.com/mholt/archiver"
    "io"
    "io/ioutil"
    "math/rand"
    "os"
    "path/filepath"
    "strings"
    "time"
)

type Student struct {
    Name  string
    Age   int
    Socre string
}

var osPath = "C:/Users/yaobin/go/"

type MyFileError struct {
    Name string
    Msg  string
}

func (me *MyFileError) Error() string {

    return fmt.Sprintf("FileName:%s \r\nErrorMsg: %s \r\n", me.Name, me.Msg)
}

func makeFile() (path string) {

    str := "a.log"

    path = osPath + str

    file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, 0777)

    // 自定义错误使用还是不对, 需要研究 Day 7 错误定义
    if err != nil {

        me := MyFileError{Name: path, Msg: err.Error()}
        return me.Error()
    }

    file.WriteString(RandString() + "\r\n")

    file.Close()
    return
}

func RandString() (str string) {

    rand.Seed(time.Now().UnixNano())
    var i = rand.Intn(100)
    var randStr = "0123456789 ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$^&*()"
    bytes := []byte(randStr)
    result := []byte{}
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for j := 0; j < i; j++ {
        result = append(result, bytes[r.Intn(len(bytes))])

    }

    str = string(result)

    return
}

func makeJsonFile() (path string) {

    filename := "b.json"

    path = osPath + filename

    file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)

    if err != nil {

        fmt.Println(err)
        return err.Error()
    }

    defer file.Close()

    file.WriteString(getJson())

    return

}

func getJson() string {

    stu1 := Student{
        Name: "Katherine",
        Age:  rand.Intn(18),

        Socre: fmt.Sprintf("%2.2f", rand.Float32()*100),
    }

    data, err := json.Marshal(stu1)

    if err != nil {

        fmt.Println(err)
    }

    return string(data)

}

type ChatType struct {
    Letter int
    Number int
    Space  int
    Other  int
}

func countChat(dir string) ChatType {

    var ct = ChatType{
        Letter: 0,
        Number: 0,
        Space:  0,
        Other:  0,
    }

    filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {

        if f == nil {
            return err
        }

        if f.IsDir() {
            return nil
        }

        file, err := os.Open(path)
        if err != nil {
            fmt.Println(err)
        }

        defer file.Close()

        fmt.Println(f.Name())

        if strings.HasSuffix(f.Name(), ".json") {

            line, err := ioutil.ReadFile(path)
            if err != nil {
                fmt.Println(err)
            }

            var nowKatherine = Student{}
            json.Unmarshal(line, &nowKatherine)
            fmt.Println(nowKatherine)

        } else {

            reader := bufio.NewReader(file)
            for {

                arrStr, _, err := reader.ReadLine()
                if err == io.EOF {
                    break
                }

                if err != nil {
                    fmt.Println(err)
                    break
                }

                // arrStr := []rune(str)
                for _, v := range arrStr {

                    switch {
                    case v >= 'a' && v <= 'z':
                        fallthrough
                    case v >= 'A' && v <= 'Z':

                        ct.Letter++

                    case v == ' ' || v == '\t':
                        ct.Space++

                    case v >= '0' || v <= '9':
                        ct.Number++

                    default:
                        ct.Other++
                    }

                }
            }

        }
        return err

    })

    fmt.Println(ct)

    return ct

}

func main() {

    f1 := makeFile()
    f2 := makeJsonFile()

    zipname := osPath + "c.gz"
    archiver.Zip.Make(zipname, []string{f1, f2})

    targetPath := osPath + "zipRes/"

    archiver.Zip.Open(zipname, targetPath)

    countChat(targetPath)

    fmt.Println(zipname)
}
