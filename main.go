package main

import (
	"github.com/darkCavalier11/downloader_backend/grpc_module"
	"github.com/darkCavalier11/downloader_backend/grpc_module/gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("unable to start server %v", err)
	}
	s := grpc.NewServer()
	gen.RegisterFileStreamingServiceServer(s, &grpc_module.Server{})
	go grpc_module.InitClient()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Unable to bind with grpc_module %v", err)
	}
	log.Println("server started")
}
