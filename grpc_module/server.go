package grpc_module

import (
	"bufio"
	"github.com/darkCavalier11/downloader_backend/grpc_module/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os/exec"
)

type Server struct {
	gen.UnimplementedFileStreamingServiceServer
}

func (*Server) ServerStreaming(req *gen.FileRequest, stream gen.FileStreamingService_ServerStreamingServer) error {
	cmd := exec.Command("yt-dlp", "-f", req.GetFormatId(), req.GetUrl())
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		log.Fatalf("error occured %v", err)
	}
	r := bufio.NewReader(stdout)
	buf := make([]byte, 0, 4*1024*1024)
	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				break
			}
			if err == io.EOF {
				break
			}
		}
		err = stream.Send(&gen.FileResponse{
			FileBytes: buf,
		})
		if err != nil {
			log.Fatalf("error, %v", err)
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func InitClient() {
	cc, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to the server %v", err)
	}
	defer cc.Close()
	serverStreamingHandler(cc)
}
