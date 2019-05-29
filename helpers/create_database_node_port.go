package helpers

import (
	pb "corebos-kubernetes/kubernetes"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"os"
)

func CreateDatabaseNodePort(nodePortData *pb.DatabaseNodePort, clientSet *kubernetes.Clientset) error {

	serviceClient := clientSet.CoreV1().Services(os.Getenv("NAMESPACE"))

	nodePort := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: nodePortData.Name,
		},
		Spec: apiv1.ServiceSpec{
			Type: apiv1.ServiceTypeNodePort,
			Selector: map[string]string{
				"component": nodePortData.Label,
			},
			Ports: []apiv1.ServicePort{
				{
					Port: nodePortData.Port,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 3306,
					},
					NodePort: nodePortData.Port,
				},
			},
		},
	}

	_, err := serviceClient.Create(nodePort)

	return err
}
