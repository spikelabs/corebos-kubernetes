package helpers

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
)

func DeleteDeploymentPvc(deploymentPvcName string, clientSet *kubernetes.Clientset) error {
	pvcClient := clientSet.CoreV1().PersistentVolumeClaims(os.Getenv("NAMESPACE"))

	deletePolicy := metav1.DeletePropagationForeground
	return pvcClient.Delete(deploymentPvcName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}
