package datalayer

import (
    "log"
    "database/sql"
    _"github.com/go-sql-driver/mysql"
    "reflect"
    "../module"
)

func GetStudent() []interface{}{
    db, err := sql.Open("mysql", "root:Changeme_123@tcp(10.22.98.175:3306)/orange?charset=utf8")
    if err != nil{
        log.Fatal(err)
	    return nil
    }
    defer db.Close()

    stmt, err := db.Prepare("SELECT ID as Id, Name, Grade FROM Student where ID = ?")
    if err != nil {
        log.Fatal(err)
        return nil
    }
    defer stmt.Close()
    // id := 002
    // s := new(Student)
    // s.ID = "002"
    rows, err := stmt.Query("002")
    // rows, err := db.Query("SELECT ID as id, Name as name, Grade as grade FROM Student")
    if err != nil{
        log.Fatalln(err)
	    return nil
    }

    columns, err := rows.Columns()
    if err != nil {
        log.Fatalln(err)
        return nil
    }
    log.Println(columns);
    
    newFunc := reflect.ValueOf(module.NewStudent)
    
    var student reflect.Value
    if newFunc.Kind() == reflect.Func {
        results := newFunc.Call(nil)
        if len(results) == 1 {
            student = results[0]
        }
    }
    
    log.Println("numField->", student.Elem().NumField())
    
    len := len(columns)
    v := make([]interface{}, len)
    for i:=0; i<len; i++ {
        log.Println("field=", columns[i])
        field := student.Elem().FieldByName(columns[i])
        v[i] = field.Addr().Interface()
    }
    
    results := make([]interface{}, 0)
    
    for rows.Next(){
        // var t []interface{}
        /*var id string
        var name string
        var grade int8*/
        // err = rows.Scan(&id, &name, &grade)
        err = rows.Scan(v...)
        if err !=nil{
            log.Fatalln(err)
        }
        /*log.Println(id)
        log.Println(name)
        log.Println(grade)*/
        
        // log.Printf("found row containing %q", s)
        results = append(results, student.Elem().Interface())
    }
    return results
}
