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
	err = helpers.CreateDeployment(deployment, clientset)
	err = helpers.CreateService(service, clientset)
	err = helpers.CreateDatabasePvc(databasePvc, clientset)
	err = helpers.CreateDatabase(database, clientset)
	err = helpers.CreateIngress(ingress, clientset)


	return &pb.CreateClientResponse{
		Success: 1,
	}, nil
}