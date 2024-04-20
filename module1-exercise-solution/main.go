package main

import (
	"github.com/cshep4/grpc-course/module1/proto"
	"log"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	person := proto.Person{
		Name: argsWithoutProg[0],
	}

	log.Printf("Hello %s!", person.GetName())
}
