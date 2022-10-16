package main

import (
	"fmt"
	"log"
	"protobuf-demo/protobuf-demo/example"

	"google.golang.org/protobuf/proto"
)

func main() {
	person := example.Person{
		Id:    112,
		Name:  "mamaz",
		Email: "mamaz@mamazo.com",
		Phones: []*example.PhoneNumber{
			{
				Number: "0821126007805",
				Type:   example.PhoneType_MOBILE,
			},
		},
	}
	bytes, err := proto.Marshal(&person)
	if err != nil {
		log.Fatalf("fail to marshal %v", err)
	}

	fmt.Println(string(bytes))

	mamaz := example.Person{}
	proto.Unmarshal(bytes, &mamaz)

	fmt.Printf("%v", mamaz)
}
