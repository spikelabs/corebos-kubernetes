package helpers

import (
	pb "corebos-kubernetes/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
	"os"
)

func UpdateIngress(ingressData *pb.Ingress, clientSet *kubernetes.Clientset) error {
	ingressClient := clientSet.ExtensionsV1beta1().Ingresses(os.Getenv("NAMESPACE"))

	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := ingressClient.Get(ingressData.Name, metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}

		result.Spec.Rules[0].Host = ingressData.SubDomain

		_, updateErr := ingressClient.Update(result)
		return updateErr
	})
}
