package helpers

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
)

func DeleteIngress(ingressName string, clientSet *kubernetes.Clientset) error {
	ingressClient := clientSet.ExtensionsV1beta1().Ingresses(os.Getenv("NAMESPACE"))

	deletePolicy := metav1.DeletePropagationForeground
	return ingressClient.Delete(ingressName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}
