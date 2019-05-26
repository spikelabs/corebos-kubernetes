package helpers

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
)

func DeleteDatabase(databaseName string, clientSet *kubernetes.Clientset) error {
	deploymentsClient := clientSet.AppsV1().Deployments(os.Getenv("NAMESPACE"))

	deletePolicy := metav1.DeletePropagationForeground
	return deploymentsClient.Delete(databaseName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}
