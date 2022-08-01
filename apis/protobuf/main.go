package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/mohamedallam1991/gopl/apis/protobuf/apis/randomuser"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

func getUser() *randomuser.User {
	theDob := randomuser.Date{
		Year:  1991,
		Month: 11,
		Day:   6,
	}
	user := randomuser.User{
		Uuid:     "7a0eed16-9430-4d68-901f-c0d4c1c3bf00",
		Email:    "jennie.nichols@example.com",
		Active:   true,
		Postcode: 123,
		Dob:      &theDob,
		Todos:    []string{"TDD", "gRPC", "Stuff"},
	}

	// for i, v := range user.GetTodos() {
	// 	fmt.Printf("the value: %v %T \n", v, v)
	// 	fmt.Printf("the index: %v %T \n", i, i)
	// }

	// fmt.Printf("%v %T \n", user.GetUuid(), user.GetUuid())
	// fmt.Printf("%v %T \n", user.GetActive(), user.GetActive())
	// fmt.Printf("%v %T \n", user.GetEmail(), user.GetEmail())
	// fmt.Printf("%v %T \n", user.GetPostcode(), user.GetPostcode())
	// fmt.Printf("%v %T \n", user.GetTodos(), user.GetTodos())
	// fmt.Print(user.GetDob())

	return &user
}
func main() {
	a := getUser()
	js := toJSON(a)
	fmt.Printf(js)
}

func all() {

	fmt.Println("gettign the user")

	user := getUser()
	fmt.Println("finished gettting the user")
	fmt.Println()

	readandWriteDemo(user)

	js := toJSON(user)
	fmt.Printf(js)

}

func toJSON(m protoiface.MessageV1) string {
	marshalling := jsonpb.Marshaler{}
	out, err := marshalling.MarshalToString(m)
	if err != nil {
		log.Fatal("cant ReadFile", err)
	}
	// fmt.Println(out)
	return out
	// return nil
}

func readandWriteDemo(pb protoreflect.ProtoMessage) {
	fmt.Println("writeTofile the user")
	writeTofile("user.bin", pb)
	fmt.Println("finish writeTofile the user")
	fmt.Println()

	fmt.Println("RadFile the user")
	user2 := &randomuser.User{}
	readFromFile("user.bin", user2)
	fmt.Println("finish RadFile the user")
	fmt.Println(user2)
	fmt.Println()
}
func readFromFile(fname string, pb protoreflect.ProtoMessage) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal("cant ReadFile", err)
		return err
	}
	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatal("cant Unmarshal", err)
		return err
	}
	return nil
}

func writeTofile(fname string, pb protoreflect.ProtoMessage) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal("cant marshall", err)
		return err
	}
	err = ioutil.WriteFile(fname, out, 0644)
	if err != nil {
		log.Fatal("cant WriteFile", err)
		return err
	}
	fmt.Printf("file has been created, check %v\n", fname)
	return nil
}
