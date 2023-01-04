package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"net"

	"github.com/go-redis/redis"
	"github.com/grpcdemo/person"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

type PersonQueryServerStruct struct {
	redisClient *redis.Client
	person.UnimplementedPersonQueryServer
}

func (x *PersonQueryServerStruct) GetPerson(ctx context.Context, in *person.PersonKey) (*person.Person, error) {
	idHash := in.GetIdHash()
	personMarshalled, err := x.redisClient.Get(idHash).Result()
	if err != nil {
		return nil, fmt.Errorf("unable to find person with hash %s due to %w", idHash, err)
	}
	var person person.Person
	if err := protojson.Unmarshal([]byte(personMarshalled), &person); err != nil {
		return nil, fmt.Errorf("unable to extract person from %s due to %w", personMarshalled, err)
	}
	return &person, nil
}

func (x *PersonQueryServerStruct) GetPersonBlockedInput(person.PersonQuery_GetPersonBlockedInputServer) error {
	return nil
}

func (x *PersonQueryServerStruct) GetPersonBlockedOutput(*person.PersonKey, person.PersonQuery_GetPersonBlockedOutputServer) error {
	return nil
}

func (x *PersonQueryServerStruct) GetPersonBlockedInputOutput(person.PersonQuery_GetPersonBlockedInputOutputServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	db := map[string]*person.Person{
		"Yamac": {
			IsMale: true,
			Age:    21,
			Weight: 82.5,
			Height: 1.86,
			Name:   "Yamac",
			Occupation: &person.Person_JobTitle{
				JobTitle: "Software Engineer",
			},
		},
	}

	for name, person := range db {
		personMarshalled, err := protojson.Marshal(person)
		if err != nil {
			log.Fatalf("unable to marshal %v", person)
		}

		if err := redisClient.Set(hash(name), string(personMarshalled), 0).Err(); err != nil {
			log.Fatalf("unable to set person struct due to %s", err)
		}
	}

	grpcServer := grpc.NewServer()
	person.RegisterPersonQueryServer(grpcServer, &PersonQueryServerStruct{
		redisClient: redisClient,
	})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func hash(v string) string {
	h := sha256.New()
	h.Write([]byte(v))
	return fmt.Sprintf("%x", h.Sum(nil))
}
