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

func Test_Distinct(t *testing.T) {
	var output []string
	FromArray([]*User{{name: "2222"}, {name: "2222"}, {name: "12"}, {name: "1235"}}).
		Filter(func(e interface{}) bool {
			return e.(*User).name == "2222"
		}).
		Map(func(e interface{}) interface{} {
			return e.(*User).name
		}).
	    Distinct().
		Collect(ToArray(&output))
	fmt.Println(output)
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

	//b.Run("stream", func(b *testing.B) {
	//    for i := 0; i < b.N; i++ {
	//
	//         _ = stream.OfSlice(arr).
	//            Filter(func(e types.T) bool {
	//                return e.(*User).name != "2222"
	//            }).
	//            Map(func(e types.T) types.R {
	//                return e.(*User).name
	//            }).
	//            ToSliceOf(reflect.TypeOf(""))
	//
	//            //// Adds a filter for strings that start with 'p'
	//            //Filter(func(e interface{}) bool {
	//            //    return e.(*User).name != "2222"
	//            //}).
	//            //// Converts back to an array
	//            //ToArray().([]string)
	//    }
	//})
	//
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
