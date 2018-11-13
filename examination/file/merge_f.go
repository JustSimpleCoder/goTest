package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
)

var osPath = "C:/Users/yaobin/Desktop/xdf_woxue_struct/"

func merge_file(p string) {

    var targetPath = "C:/Users/yaobin/Desktop/xdf_woxue_merge.sql"

    filepath.Walk(p, func(path string, f os.FileInfo, err error) error {

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

        line, err := ioutil.ReadFile(path)
        if err != nil {
            fmt.Println(err)
        }

        nf, err := os.OpenFile(targetPath, os.O_CREATE|os.O_APPEND, 0777)
        if err != nil {
            fmt.Println(err)
        }

        nf.WriteString(string(line))
        nf.Close()

        return nil
    })

}

func main() {

    merge_file(osPath)

}
