package main

import (
	"context"
	"fmt"
	"log"
	"net"

	lib "github.com/grpcdemo/lib"
	"github.com/grpcdemo/person"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/yaml.v2"
)

var (
	file *string = flag.StringP("file", "f", "config.yaml", "server configuration")
)

type PersonQueryServerStruct struct {
	db map[string]string
	person.UnimplementedPersonQueryServer
}

func (x *PersonQueryServerStruct) SetPerson(ctx context.Context, in *person.Person) (*emptypb.Empty, error) {
	empty := new(emptypb.Empty)
	idHash := lib.Hash(in.GetName(), 10)
	personMarshalled, err := protojson.Marshal(in)
	if err != nil {
		return empty, fmt.Errorf("unable to marshal person: %w", err)
	}

	fmt.Printf("Saved: %s\n", string(personMarshalled))

	x.db[idHash] = string(personMarshalled)
	return empty, nil
}

func (x *PersonQueryServerStruct) GetPerson(ctx context.Context, in *person.PersonKey) (*person.Person, error) {
	idHash := in.GetId()
	personMarshalled := x.db[idHash]

	fmt.Printf("Requested: %s\n", personMarshalled)

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
	flag.Parse()
	config, err := parseConfig()
	if err != nil {
		log.Fatalf("unable to read server params: %s", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	person.RegisterPersonQueryServer(grpcServer, &PersonQueryServerStruct{
		db: make(map[string]string),
	})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type configServer struct {
	Port int `yaml:"port"`
}

func parseConfig() (*configServer, error) {
	bytes, err := lib.ReadFile(*file)
	if err != nil {
		return nil, fmt.Errorf("unable to read config file: %w", err)
	}

	var config configServer
	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, fmt.Errorf("unable to cast object to config struct: %w", err)
	}

	return &config, nil
}
