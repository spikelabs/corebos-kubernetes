package helpers

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
)

func DeleteDatabaseNodePort(databaseNodePortName string, clientSet *kubernetes.Clientset) error {
	serviceClient := clientSet.CoreV1().Services(os.Getenv("NAMESPACE"))

	deletePolicy := metav1.DeletePropagationForeground
	return serviceClient.Delete(databaseNodePortName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}
