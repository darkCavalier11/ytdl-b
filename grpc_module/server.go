package grpc_module

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/darkCavalier11/downloader_backend/grpc_module/gen"
	"github.com/darkCavalier11/downloader_backend/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os/exec"
)

type Server struct {
	gen.UnimplementedFileStreamingServiceServer
	gen.UnimplementedGetFileMetaServiceServer
}

func (s *Server) GetFileMeta(ctx context.Context, req *gen.RequestUrl) (*gen.FileMeta, error) {
	url := req.GetUrl()
	cmd := exec.Command("yt-dlp", "--dump-json", "--skip-download", url)
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error occured getting json meta for url %v > %v", url, err)
	}
	var fileMeta models.FileMeta
	err = json.Unmarshal(stdout, &fileMeta)
	if err != nil {
		return nil, fmt.Errorf("error parsing json %v", err)
	}
	return fileMeta.ConvertToGRPCFileMeta(), nil
}

func (*Server) FileStreaming(req *gen.FileRequest, stream gen.FileStreamingService_FileStreamingServer) error {
	cmd := exec.Command("yt-dlp", "-o", "-", "-f", req.GetFormatId(), req.GetUrl())
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error at stdout pipe > %v", err)
	}
	if err = cmd.Start(); err != nil {
		return fmt.Errorf("error starting command > %v", err)
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
			return fmt.Errorf("error completing file streaming %v", err)
		}
	}
	return nil
}

func InitClient() {
	cc, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to the server %v", err)
	}
	defer cc.Close()
	getFileMetaHandler(cc)
}
