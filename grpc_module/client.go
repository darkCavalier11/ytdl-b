package grpc_module

import (
	"context"
	"fmt"
	"github.com/darkCavalier11/downloader_backend/grpc_module/gen"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"
)

func serverStreamingHandler(cc *grpc.ClientConn) {
	maxSizeOption := grpc.MaxCallRecvMsgSize(4 * 1024 * 1024)

	c := gen.NewFileStreamingServiceClient(cc)
	req := &gen.FileRequest{
		FormatId: "140",
		Url:      "https://www.youtube.com/watch?v=xUwePVuH1PM",
	}
	res, err := c.ServerStreaming(context.Background(), req, maxSizeOption)
	if err != nil {
		log.Fatalf("Unable to make request to the server. %v", err)
	}
	f, err := os.Create("file.webm")
	n := 0
	t := time.Now()
	for {
		msg, err := res.Recv()
		if err == io.EOF {
			log.Println("Finished file receiving")
			break
		}
		if err != nil {
			log.Fatalf("Error occured during server streaming %v", err)
		}
		rc, err := f.Write(msg.GetFileBytes())
		if err != nil {
			log.Fatalln(err)
		}
		n += rc
	}
	fmt.Println(time.Since(t).Seconds())
}
