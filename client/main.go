package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/grpcdemo/lib"
	"github.com/grpcdemo/person"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	addr *string = flag.StringP("port", "p", "50051", "[<host>:]<port>")
	cmd  *string = flag.StringP("cmd", "c", "commands.json", "list of commands")
)

func main() {
	flag.Parse()

	config, err := parseConfig()
	if err != nil {
		log.Fatalf("unable to get port due to %s", err)
	}

	dialAddress := fmt.Sprintf("%s:%d", config.ServerHost, config.ServerPort)

	dialOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(dialAddress, dialOption)
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
				Id: lib.Hash(cmd.Arg, 10),
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
	ServerHost string `yaml:"serverHost"`
	ServerPort int    `yaml:"serverPort"`
	Port       int    `yaml:"port"`
}

func parseConfig() (*configClient, error) {
	var host string
	var port int
	if matches, _ := regexp.MatchString(`\w*:\d`, *addr); matches {
		addrParts := strings.Split(*addr, ":")
		host = addrParts[0]
		port, _ = strconv.Atoi(addrParts[1])
	} else if matches, _ := regexp.MatchString(`\d`, *addr); matches {
		port, _ = strconv.Atoi(*addr)
	}

	return &configClient{
		ServerHost: host,
		ServerPort: port,
	}, nil
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
