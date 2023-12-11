package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"buf.build/gen/go/ahmad-ibrahim/count-api/connectrpc/go/count/v1/countv1connect"
	countv1 "buf.build/gen/go/ahmad-ibrahim/count-api/protocolbuffers/go/count/v1"
	connect "connectrpc.com/connect"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	fmt.Println("attempting to connect to host: ", host, " on port: ", port)

	client := countv1connect.NewCountServiceClient(
		http.DefaultClient,
		"http://"+host+":"+port,
	)

	for {
		res, err := client.Count(
			context.Background(),
			connect.NewRequest(&countv1.CountRequest{}),
		)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(res.Msg)
		time.Sleep(1 * time.Second)
	}
}
