package helpers

import (
	pb "corebos-kubernetes/kubernetes"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/apimachinery/pkg/api/resource"
	"os"
)

func CreateDeploymentPvc(deploymentPvcData *pb.DeploymentPvc, clientSet *kubernetes.Clientset) (error) {

	pvcClient := clientSet.CoreV1().PersistentVolumeClaims(os.Getenv("NAMESPACE"))

	deploymentPvc := &apiv1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentPvcData.Name,
		},
		Spec: apiv1.PersistentVolumeClaimSpec{
			AccessModes: []apiv1.PersistentVolumeAccessMode{
				apiv1.ReadWriteOnce,
			},
			Resources: apiv1.ResourceRequirements{
				Requests: map[apiv1.ResourceName]resource.Quantity{
					apiv1.ResourceName(apiv1.ResourceStorage): resource.MustParse(deploymentPvcData.Storage),
				},
			},
		},
	}


	_, err := pvcClient.Create(deploymentPvc)
	if err != nil {
		return err
	}
	return nil
}