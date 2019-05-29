package helpers

import (
	pb "corebos-kubernetes/kubernetes"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"os"
)

func CreateDatabaseService(databaseServiceData *pb.DatabaseService, clientSet *kubernetes.Clientset) error {

	serviceClient := clientSet.CoreV1().Services(os.Getenv("NAMESPACE"))

	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: databaseServiceData.Name,
		},
		Spec: apiv1.ServiceSpec{
			Type: apiv1.ServiceTypeClusterIP,
			Selector: map[string]string{
				"component": databaseServiceData.Label,
			},
			Ports: []apiv1.ServicePort{
				{
					Port: 3306,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 3306,
					},
				},
			},
		},
	}

	_, err := serviceClient.Create(service)
	return err
}
