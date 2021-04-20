package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/mohammadinasab-dev/grpc/cmd"
	"github.com/mohammadinasab-dev/grpc/configuration"
	"github.com/mohammadinasab-dev/grpc/grpcserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	op := flag.String("op", "s", "s for Server and c for Client.")
	flag.Parse()

	switch strings.ToLower(*op) {
	case "s":
		runGrpcServer()
	case "c":
		runGrpcClient()
	}

}

func runGrpcServer() {
	grpclog.Println("starting server...")
	ls, err := net.Listen("tcp", ":8282")
	if err != nil {
		log.Fatalln("failed to listen")
	}
	grpclog.Println("listenning established...")
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	config, err := configuration.LoadConfig(".")
	if err != nil {
		log.Fatalln("faild to load config")
	}
	userserver, err := grpcserver.NewGrpcServer(config)
	if err != nil {
		log.Fatalln(err)
	}
	cmd.RegisterUserServiceServer(server, userserver)
	err = server.Serve(ls)
	if err != nil {
		log.Fatalln(err)
	}
}

func runGrpcClient() {
	conn, err := grpc.Dial("127.0.0.1:8282", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := cmd.NewUserServiceClient(conn)
	input := ""
	fmt.Println("All users? (y/n)")
	fmt.Scanln(&input)
	if strings.EqualFold(input, "y") {
		users, err := client.GetUsers(context.Background(), &cmd.Request{})
		if err != nil {
			log.Fatalln(err)
		}
		for {
			user, err := users.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(user)
		}
		return
	}
	fmt.Println("name?")
	fmt.Scanln(&input)

	user, err := client.GetUser(context.Background(), &cmd.Request{Name: input})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user)

}
