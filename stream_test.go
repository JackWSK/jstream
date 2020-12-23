package jstream

import (
    "fmt"
    "testing"
)

type User struct {
    name string
}

func Test_ToArray(t *testing.T) {
    var output []string
    FromArray([]*User{{name: "123"}, {name: "2222"}, {name: "123"}, {name: "123"}}).
        Filter(func(e interface{}) bool {
            return e.(*User).name != "2222"
        }).
        Map(func(e interface{}) interface{} {
            return e.(*User).name
        }).
        Collect(ToArray(&output))
    fmt.Println(output)
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
        Filter(func(e interface{}) bool {
            return e.(*User).name != "2222"
        }).
        Collect(Group(&users, func(i interface{}) interface{} {
            return (i.(*User)).name
        }))
    fmt.Println(users)
}
