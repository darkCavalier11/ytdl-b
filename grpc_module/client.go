package grpc_module

import (
	"context"
	"fmt"
	"github.com/darkCavalier11/downloader_backend/grpc_module/gen"
	"google.golang.org/grpc"
	"io"
	"os"
)

func serverStreamingHandler(cc *grpc.ClientConn) error {
	maxSizeOption := grpc.MaxCallRecvMsgSize(4 * 1024 * 1024)

	c := gen.NewFileStreamingServiceClient(cc)
	req := &gen.FileRequest{
		FormatId: "140",
		Url:      "https://www.youtube.com/watch?v=xUwePVuH1PM",
	}
	res, err := c.GetFileBytesStream(context.Background(), req, maxSizeOption)
	if err != nil {
		return fmt.Errorf("unable to make request to the server. %v", err)
	}
	f, err := os.Create("file.webm")
	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error occured during server streaming > %v", err)
		}
		_, err = f.Write(msg.GetFileBytes())
		if err != nil {
			return fmt.Errorf("error occured writing file > %v", err)
		}
	}
	return nil
}

func getFileMetaHandler(cc *grpc.ClientConn) error {
	c := gen.NewGetFileMetaServiceClient(cc)
	req := gen.RequestUrl{Url: "https://www.youtube.com/watch?v=geYV5R7Nv2g"}
	res, err := c.GetFileMeta(context.Background(), &req)
	if err != nil {
		return fmt.Errorf("unable to get file meta %v", err)
	}
	fmt.Println(res.Title)
	return nil
}
