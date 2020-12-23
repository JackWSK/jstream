# jstream
StreamApi inspired by java8 stream

just finish little future, keep updating

### Use jstream
```shell
go get -u github.com/JackWSK/jstream
```

### Example

#### Collect into array
```go

import "github.com/JackWSK/jstream"

type User struct {
    name string
}
func main() {
    var output []string
    jstream.FromArray([]*User{{name: "123"}, {name: "2222"}, {name: "123"}, {name: "123"}}).
        Filter(func(e interface{}) bool {
            return e.(*User).name != "2222"
        }).
        Map(func(e interface{}) interface{} {
            return e.(*User).name
        }).
        Collect(jstream.ToArray(&output))
    fmt.Println(output)
}

```

#### Collect into map
```go
func func main() {
    var users map[string]*User
    jstream.FromArray([]*User{{name: "1234"}, {name: "2222"}, {name: "12"}, {name: "1235"}}).
        Filter(func(e interface{}) bool {
            return e.(*User).name != "2222"
        }).
        //Map(func(e interface{}) interface{} {
        //   return fmt.Sprintf("%d", e)
        //}).
        Collect(jstream.ToMap(&users, func(i interface{}) interface{} {
            return (i.(*User)).name
        }))
    fmt.Println(users)
}
```

#### Group into map
```go
func main() {
    var users map[string][]*User
    jstream.FromArray([]*User{{name: "2222"}, {name: "2222"}, {name: "12"}, {name: "1235"}}).
        Filter(func(e interface{}) bool {
            return e.(*User).name != "2222"
        }).
        Collect(jstream.Group(&users, func(i interface{}) interface{} {
            return (i.(*User)).name
        }))
    fmt.Println(users)
}
```