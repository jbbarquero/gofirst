package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"

	pb "github.com/jbbarquero/gofirst/protobuf"
)

func main() {

	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	book := &pb.AddressBook{}
	book.People = append(book.People, &p)

	fmt.Println("Book: ", book)

	//Writing a Message(marshall)
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	fmt.Println("Marshalled book: ", out)

	//Reading a Message (unmarshall)
	sameBook := &pb.AddressBook{}
	if err := proto.Unmarshal(out, sameBook); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println("Same book: ", book)

}
