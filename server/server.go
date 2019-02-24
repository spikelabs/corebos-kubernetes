package server

import (
	"context"
	pb "corebos-kubernetes/kubernetes"
	)

type server struct {

}

func (s *server)  CreateClient(ctx context.Context, request *pb.CreateClientRequest) (*pb.CreateClientResponse, error) {

	deployment := request.Deployment
	service := request.Service
	deploymentPvc := request.DeploymentPvc
	ingress := request.Ingress
	database := request.Database
	databasePvc := request.DatabasePvc



	return &pb.CreateClientResponse{
		Success: 1,
	}, nil
}