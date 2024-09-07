package main

import (
	"github.com/cshep4/grpc-course/module1/proto"
	"log"
)

func main() {
	person := proto.Person{
		Name: "Chris",
	}

	log.Println(person.GetName())
}
