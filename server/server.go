package server

import (
	"context"
	"corebos-kubernetes/helpers"
	pb "corebos-kubernetes/kubernetes"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Server struct {
}

func (s *Server) CreateClientDatabase(ctx context.Context, request *pb.CreateClientDatabaseRequest) (*pb.ClientResponse, error) {

	database := request.Database
	databaseService := request.DatabaseService
	databasePvc := request.DatabasePvc
	databaseNodePort := request.DatabaseNodePort
	configData := request.ConfigData

	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(configData))
	if err != nil {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.CreateDatabasePvc(databasePvc, clientSet)
	if err != nil && !errors.IsAlreadyExists(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.CreateDatabase(database, databasePvc.Name, clientSet)
	if err != nil && !errors.IsAlreadyExists(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.CreateDatabaseService(databaseService, clientSet)
	if err != nil && !errors.IsAlreadyExists(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.CreateDatabaseNodePort(databaseNodePort, clientSet)
	if err != nil && !errors.IsAlreadyExists(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	return &pb.ClientResponse{
		Success: 1,
	}, nil
}

func (s *Server) CreateClientDeployment(ctx context.Context, request *pb.CreateClientDeploymentRequest) (*pb.ClientResponse, error) {

	deployment := request.Deployment
	service := request.Service
	deploymentPvc := request.DeploymentPvc
	ingress := request.Ingress
	nodePortName := request.DatabaseNodePortName
	configData := request.ConfigData

	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(configData))
	if err != nil {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.DeleteDatabaseNodePort(nodePortName, clientSet)
	if err != nil && !errors.IsNotFound(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.CreateDeploymentPvc(deploymentPvc, clientSet)
	if err != nil && !errors.IsAlreadyExists(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.CreateDeployment(deployment, deploymentPvc.Name, clientSet)
	if err != nil && !errors.IsAlreadyExists(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.CreateService(service, clientSet)
	if err != nil && !errors.IsAlreadyExists(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.CreateIngress(ingress, clientSet)
	if err != nil && !errors.IsAlreadyExists(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	return &pb.ClientResponse{
		Success: 1,
	}, nil
}

func (s *Server) UpdateClientIngress(ctx context.Context, request *pb.UpdateClientIngressRequest) (*pb.ClientResponse, error) {

	ingress := request.Ingress
	configData := request.ConfigData

	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(configData))
	if err != nil {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.UpdateIngress(ingress, clientSet)
	if err != nil && !errors.IsNotFound(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	return &pb.ClientResponse{
		Success: 1,
	}, nil
}

func (s *Server) UpdateClientDeployment(ctx context.Context, request *pb.UpdateClientDeploymentRequest) (*pb.ClientResponse, error) {

	deployment := request.Deployment
	configData := request.ConfigData

	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(configData))
	if err != nil {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.UpdateDeployment(deployment, clientSet)
	if err != nil && !errors.IsNotFound(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	return &pb.ClientResponse{
		Success: 1,
	}, nil
}

func (s *Server) DeleteClient(ctx context.Context, request *pb.DeleteClientRequest) (*pb.ClientResponse, error) {

	deploymentName := request.DeploymentName
	deploymentPvcName := request.DeploymentPvcName
	serviceName := request.ServiceName
	ingressName := request.IngressName
	databaseName := request.DatabaseName
	databasePvcName := request.DatabasePvcName
	databaseServiceName := request.DatabaseServiceName
	configData := request.ConfigData

	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(configData))
	if err != nil {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.DeleteDeployment(deploymentName, clientSet)
	if err != nil && !errors.IsNotFound(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.DeleteDeploymentPvc(deploymentPvcName, clientSet)
	if err != nil && !errors.IsNotFound(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.DeleteService(serviceName, clientSet)
	if err != nil && !errors.IsNotFound(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.DeleteIngress(ingressName, clientSet)
	if err != nil && !errors.IsNotFound(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.DeleteDatabase(databaseName, clientSet)
	if err != nil && !errors.IsNotFound(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.DeleteDatabasePvc(databasePvcName, clientSet)
	if err != nil && !errors.IsNotFound(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	err = helpers.DeleteDatabaseService(databaseServiceName, clientSet)
	if err != nil && !errors.IsNotFound(err) {
		return &pb.ClientResponse{
			Success: 0,
			Error:   err.Error(),
		}, nil
	}

	return &pb.ClientResponse{
		Success: 1,
	}, nil
}
