package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/grpcdemo/person"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(
		fmt.Sprintf("localhost:%d", 50051),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	personQueryClient := person.NewPersonQueryClient(conn)

	ctx := context.Background()
	person, err := personQueryClient.GetPerson(ctx, &person.PersonKey{
		IdHash: hash("Yamac"),
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(person)
	}
}

func hash(v string) string {
	h := sha256.New()
	h.Write([]byte(v))
	return fmt.Sprintf("%x", h.Sum(nil))
}
