package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v2"

	"github.com/grpcdemo/lib"
	"github.com/grpcdemo/person"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	file *string = flag.StringP("file", "f", "config.yaml", "client configuration")
	cmd  *string = flag.StringP("cmd", "c", "commands.json", "list of commands")
)

func main() {
	flag.Parse()

	config, err := parseConfig()
	if err != nil {
		log.Fatalf("unable to get port due to %s", err)
	}

	dialOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(fmt.Sprintf(":%d", config.Port), dialOption)
	if err != nil {
		log.Fatalf("fail to dial: %s", err)
	}
	defer conn.Close()
	personQueryClient := person.NewPersonQueryClient(conn)

	ctx := context.Background()

	commands, err := parseCommands()
	if err != nil {
		log.Fatalf("unable to get commands: %s", err)
	}

	for _, cmd := range commands {
		switch cmd.Op {
		case "set":
			var value person.Person
			if err := protojson.Unmarshal([]byte(cmd.Arg), &value); err != nil {
				log.Fatalf("unable to cast object to person: %s", err)
			}
			if _, err := personQueryClient.SetPerson(ctx, &value); err != nil {
				log.Fatalf("unable to set person: %s", err)
			}
		case "get":
			key := person.PersonKey{
				IdHash: lib.Hash(cmd.Arg),
			}
			person, err := personQueryClient.GetPerson(ctx, &key)
			if err != nil {
				log.Fatalf("unable to get person: %s", err)
			} else {
				fmt.Println(person)
			}
		default:
			log.Fatalf("undefined mode: %s", cmd.Op)
		}
	}
}

type configClient struct {
	Port int `yaml:"port"`
}

func parseConfig() (*configClient, error) {
	bytes, err := lib.ReadFile(*file)
	if err != nil {
		return nil, fmt.Errorf("unable to read config file: %w", err)
	}

	var config configClient
	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, fmt.Errorf("unable to cast object to config struct: %w", err)
	}

	return &config, nil
}

type kwArg struct {
	Op  string `json:"op"`
	Arg string `json:"arg"`
}

func parseCommands() ([]kwArg, error) {
	bytes, err := lib.ReadFile(*cmd)
	if err != nil {
		return nil, fmt.Errorf("unable to read config file: %w", err)
	}

	var commands []kwArg
	if err := json.Unmarshal(bytes, &commands); err != nil {
		return nil, fmt.Errorf("unable to cast object to config struct: %w", err)
	}

	return commands, nil
}
