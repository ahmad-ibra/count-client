package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"buf.build/gen/go/ahmad-ibrahim/count-api/connectrpc/go/count/v1/countv1connect"
	countv1 "buf.build/gen/go/ahmad-ibrahim/count-api/protocolbuffers/go/count/v1"
	connect "connectrpc.com/connect"
)

func main() {
	client := countv1connect.NewCountServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
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
