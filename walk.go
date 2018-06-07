package main

import (
    "fmt"
    "path/filepath"
    "os"
)

func walk(path string) error {
    err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
        if f == nil {
            fmt.Println("nil")
            return err
        }
        if f.IsDir() {
            //fmt.Println("dir")
            return nil
        }
        fmt.Println(path)
        return nil
    })
    return err
}

func main() {
    fmt.Println("hello")
    walk("server")
}
