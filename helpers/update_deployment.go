package helpers

import (
	pb "corebos-kubernetes/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
	"k8s.io/client-go/util/retry"
)

func UpdateDeployment (deploymentData *pb.Deployment, clientSet *kubernetes.Clientset) (error) {

	deploymentsClient := clientSet.AppsV1().Deployments(os.Getenv("NAMESPACE"))

	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, getErr := deploymentsClient.Get(deploymentData.Name, metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}

		result.Spec.Template.Spec.Containers[0].Image = deploymentData.Image // change nginx version
		_, updateErr := deploymentsClient.Update(result)
		return updateErr
	})
}
