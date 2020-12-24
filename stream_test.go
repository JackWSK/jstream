package jstream

import (
    "fmt"
    "github.com/jucardi/go-streams/streams"
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
            return e.(*User).name == "2222"
        }).
        Collect(Group(&users, func(i interface{}) interface{} {
            return (i.(*User)).name
        }))
    fmt.Println(users)
}

func BenchmarkRawToArray(b *testing.B) {
    arr := []*User{}
    for i := 0; i < 1000; i++ {
        arr = append(arr, &User{name: fmt.Sprintf("%d", i)})
    }

    b.Run("RawToArray", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            var output []string
            for _, user := range arr {
                if user.name != "2222" {
                    output = append(output, user.name)
                }
            }

        }
    })
}

func BenchmarkToArray(b *testing.B) {
    arr := []*User{}
    for i := 0; i < 1000; i++ {
        arr = append(arr, &User{name: fmt.Sprintf("%d", i)})
    }

    b.Run("RawToArray", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            var output []string
            for _, user := range arr {
                if user.name != "2222" {
                    output = append(output, user.name)
                }
            }

        }
    })

    b.Run("go-streams", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            _ = streams.
                // Creates a stream from the given array
                FromArray(arr).
                Filter(func(e interface{}) bool {
                return e.(*User).name != "2222"
            }).Map(func(e interface{}) interface{} {
                return e.(*User).name
            }).ToArray().([]string)

                //// Adds a filter for strings that start with 'p'
                //Filter(func(e interface{}) bool {
                //    return e.(*User).name != "2222"
                //}).
                //// Converts back to an array
                //ToArray().([]string)
        }
    })

    b.Run("jstream", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            var output []string
            FromArray(arr).
                Filter(func(e interface{}) bool {
                return e.(*User).name != "2222"
                }).
                Map(func(e interface{}) interface{} {
                return e.(*User).name
                }).
                Collect(ToArray(&output))
        }
    })
}

func BenchmarkGroup(b *testing.B) {
    arr := []*User{{name: "2222"}, {name: "2222"}, {name: "12"}, {name: "1235"}}
    for i := 0; i < b.N; i++ {
        var users map[string][]*User
        FromArray(arr).
            Filter(func(e interface{}) bool {
                return e.(*User).name != "2222"
            }).
            Collect(Group(&users, func(i interface{}) interface{} {
                return (i.(*User)).name
            }))
        //fmt.Println(users)
    }
}
