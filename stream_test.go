package jstream

import (
    "fmt"
    "reflect"
    "testing"
)

type User struct {
    name string
}

func Test_Array(t *testing.T) {
    var users []*User
    FromArray([]*User{{name: "123"}, {name: "2222"}, {name: "123"}, {name: "123"}}).
        Filter(func(e interface{}) bool {
            return e.(*User).name == "2222"
        }).
        //Map(func(e interface{}) interface{} {
        //   return fmt.Sprintf("%d", e)
        //}).
        Collect(ToArray(&users))
    fmt.Println(users)
}

func Test_Map(t *testing.T) {
    var users map[string]*User
    FromArray([]*User{{name: "1234"}, {name: "2222"}, {name: "12"}, {name: "1235"}}).
        Filter(func(e interface{}) bool {
           return e.(*User).name != "2222"
        }).
        //Map(func(e interface{}) interface{} {
        //   return fmt.Sprintf("%d", e)
        //}).
        Collect(ToMap(&users, func(i interface{}) interface{} {
            return (i.(*User)).name
        }))
    fmt.Println(users)
}

func Test_Group(t *testing.T) {
    var users map[string][]*User
    FromArray([]*User{{name: "2222"}, {name: "2222"}, {name: "12"}, {name: "1235"}}).
        //Filter(func(e interface{}) bool {
        //    return e.(*User).name != "2222"
        //}).
        //Map(func(e interface{}) interface{} {
        //   return fmt.Sprintf("%d", e)
        //}).
        Collect(Group(&users, func(i interface{}) interface{} {
            return (i.(*User)).name
        }))
    fmt.Println(users)
}

func Test_MapKey(t *testing.T) {
   var users map[string][]*User

    value := reflect.ValueOf(users)
    listValue := value.MapIndex(reflect.ValueOf("abc"))
    if listValue.Kind() == reflect.Invalid {
        listValue = reflect.MakeSlice(value.Type().Elem(), 0, 0)
    }
}
