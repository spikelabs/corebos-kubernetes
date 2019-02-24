package helpers

import (
	pb "corebos-kubernetes/kubernetes"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func create_deployment_vpc(deploymentPvcData *pb.DeploymentPvc, clientSet *kubernetes.Clientset) (error) {

	pvcClient := clientSet.CoreV1().PersistentVolumeClaims(apiv1.NamespaceDefault)

	deploymentPvc := &apiv1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentPvcData.Name,
		},
		Spec: apiv1.PersistentVolumeClaimSpec{

		},
	}


	_, err := pvcClient.Create(deploymentPvc)
	if err != nil {
		return err
	}
	return nil

	return nil
}