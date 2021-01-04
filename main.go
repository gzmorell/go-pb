package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	pb "github.com/gzmorell/go-pb/tutorialpb/addressbookpb"
)

func main() {
	// Create Addressbook Object with few persons
	home := pb.Person_HOME
	book := &pb.AddressBook{
		People: []*pb.Person{
			{
				Id:    proto.Int32(1234),
				Name:  proto.String("John Doe"),
				Email: proto.String("jdoe@example.com"),
				Phones: []*pb.Person_PhoneNumber{
					{Number: proto.String("555-4321"), Type: &home},
				},
			},
			{
				Id:    proto.Int32(1235),
				Name:  proto.String("Alex"),
				Email: proto.String("alex@example.com"),
				Phones: []*pb.Person_PhoneNumber{
					{Number: proto.String("1234-5678"), Type: &home},
				},
			},
		},
	}

	// let's unmarshal our Address object into byte array so
	// that we can use it transfer the data across services
	data, err := proto.Marshal(book)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// printing out our raw protobuf object
	fmt.Println("Raw data", data)

	// let's go the other way and unmarshal
	// our byte array into an object we can modify
	// and use
	addresssBook := pb.AddressBook{}
	err = proto.Unmarshal(data, &addresssBook)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	for _, person := range addresssBook.People {
		// print out our `person` object
		// for good measure
		fmt.Println("===========================")
		fmt.Println("ID: ", person.GetId())
		fmt.Println("Name: ", person.GetName())
		fmt.Println("Email: ", person.GetEmail())
		fmt.Println("Phones: ", person.GetPhones())
	}
}
