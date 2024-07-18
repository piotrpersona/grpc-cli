package main

import (
	"context"
	"fmt"
	"os"

	api "github.com/itsksaurabh/go-grpc-examples/unary/sum/sumpb"
	"google.golang.org/grpc"
)

func exit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	conn, err := grpc.NewClient("localhost:9999")
	exit(err)

	cli := api.NewSumClient(conn)

	request := &api.SumRequest{
		Numbers: &api.Numbers{
			A: 4,
			B: 5,
		},
	}
	response, err := cli.Add(context.Background(), request)
	exit(err)

	fmt.Printf("Response %+v\n", response)
}
