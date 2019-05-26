package helpers

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
)

func DeleteDeployment(deploymentName string, clientSet *kubernetes.Clientset) error {
	deploymentsClient := clientSet.AppsV1().Deployments(os.Getenv("NAMESPACE"))

	deletePolicy := metav1.DeletePropagationForeground
	return deploymentsClient.Delete(deploymentName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}
