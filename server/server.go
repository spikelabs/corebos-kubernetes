package server

import (
	"context"
	pb "corebos-kubernetes/kubernetes"
	helpers "corebos-kubernetes/helpers"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes"
)

type Server struct {

}

func (s *Server)  CreateClient(ctx context.Context, request *pb.CreateClientRequest) (*pb.CreateClientResponse, error) {

	deployment := request.Deployment
	service := request.Service
	deploymentPvc := request.DeploymentPvc
	ingress := request.Ingress
	database := request.Database
	databaseService := request.DatabaseService
	databasePvc := request.DatabasePvc

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// call helpers to create resources in kubernetes cluster

	err = helpers.CreateDeploymentPvc(deploymentPvc, clientset)
	if err != nil{
		return &pb.CreateClientResponse{
			Success: 0,
			Error: err.Error(),
		}, nil
	}

	err = helpers.CreateDeployment(deployment, clientset)
	if err != nil{
		return &pb.CreateClientResponse{
			Success: 0,
			Error: err.Error(),
		}, nil
	}

	err = helpers.CreateService(service, clientset)
	if err != nil{
		return &pb.CreateClientResponse{
			Success: 0,
			Error: err.Error(),
		}, nil
	}

	err = helpers.CreateDatabasePvc(databasePvc, clientset)
	if err != nil{
		return &pb.CreateClientResponse{
			Success: 0,
			Error: err.Error(),
		}, nil
	}

	err = helpers.CreateDatabase(database, clientset)
	if err != nil{
		return &pb.CreateClientResponse{
			Success: 0,
			Error: err.Error(),
		}, nil
	}

	err = helpers.CreateDatabaseService(databaseService, clientset)
	if err != nil{
		return &pb.CreateClientResponse{
			Success: 0,
			Error: err.Error(),
		}, nil
	}

	err = helpers.CreateIngress(ingress, clientset)
	if err != nil{
		return &pb.CreateClientResponse{
			Success: 0,
			Error: err.Error(),
		}, nil
	}


	return &pb.CreateClientResponse{
		Success: 1,
	}, nil
}