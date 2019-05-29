package main

import (
	kb "corebos-kubernetes/kubernetes"
	"corebos-kubernetes/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.ServerOption(grpc.MaxRecvMsgSize(2000000000)),
		grpc.ServerOption(grpc.MaxSendMsgSize(2000000000)),
	)

	grpcServer := &server.Server{}

	kb.RegisterClusterManagerServer(s, grpcServer)

	s.Serve(lis)
}
