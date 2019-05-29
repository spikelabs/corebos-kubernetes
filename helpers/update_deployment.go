package helpers

import (
	pb "corebos-kubernetes/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
	"os"
)

func UpdateDeployment(deploymentData *pb.Deployment, clientSet *kubernetes.Clientset) error {

	deploymentsClient := clientSet.AppsV1().Deployments(os.Getenv("NAMESPACE"))

	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := deploymentsClient.Get(deploymentData.Name, metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}

		result.Spec.Template.Spec.Containers[0].Image = deploymentData.Image
		_, updateErr := deploymentsClient.Update(result)
		return updateErr
	})
}
