package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	api "github.com/itsksaurabh/go-grpc-examples/unary/sum/sumpb"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func exit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	pctx := context.Background()

	addr := flag.String("addr", "localhost:9999", "gRPC server address")
	concurrency := flag.Int("concurrency", 1, "how many concurrent requests to send")
	single := flag.Bool("single", false, "send only 1 request and exit")
	flag.Parse()

	conn, err := grpc.NewClient(*addr)
	exit(err)

	cli := api.NewSumClient(conn)

	if *concurrency > 1 {
		fmt.Printf("running with concurrency: %d", *concurrency)
	}

	for {
		wg, ctx := errgroup.WithContext(pctx)
		wg.SetLimit(*concurrency)
		for range *concurrency {
			wg.Go(func() error {
				request := &api.SumRequest{
					Numbers: &api.Numbers{
						A: 4,
						B: 5,
					},
				}
				response, err := cli.Add(ctx, request)
				if err != nil {
					return err
				}

				fmt.Printf("Response %+v\n", response)
				return nil
			})
		}
		exit(wg.Wait())

		if *single {
			break
		}
	}
}
