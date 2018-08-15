package main

import (
    "encoding/xml"
    "fmt"
    "os"
    "io/ioutil"
)

type Mapper struct {
    XMLName xml.Name `xml:"mapper"`
    Namespace string `xml:"namespace,attr"`
    Select []Selecter `xml:"select"`
    Delete []Deleter `xml:"delete"`
    Insert []Inserter `xml:"insert"`
    Update []Updater `xml:"update"`
}

type Selecter struct {
    Id string `xml:"id,attr"`
    Sql string `xml:",innerxml"`
}

type Deleter struct {
    Id string `xml:"id,attr"`
    Sql string `xml:",innerxml"`
}

type Inserter struct {
    Id string `xml:"id,attr"`
    Sql string `xml:",innerxml"`
}

type Updater struct {
    Id string `xml:"id,attr"`
    Sql string `xml:",innerxml"`
}


func main() {
    file, err := os.Open("test.xml")
    defer file.Close()
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
    data, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
    
    d := Mapper{}
    err = xml.Unmarshal(data, &d)
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
    
    fmt.Println(d)
    fmt.Println(d.Namespace)
    for index, value := range d.Select {
        fmt.Printf("(%d)id=%s, sql=%s\n", index, value.Id, value.Sql)
    }
}
