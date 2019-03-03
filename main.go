package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	kb "corebos-kubernetes/kubernetes"
	"corebos-kubernetes/server"
)

func main()  {
	lis, err := net.Listen("tcp", ":8080")
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
