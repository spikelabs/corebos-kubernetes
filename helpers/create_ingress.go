package helpers

import (
	pb "corebos-kubernetes/kubernetes"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1beta "k8s.io/api/extensions/v1beta1"
)
func create_ingress(ingressData *pb.Ingress, clientSet *kubernetes.Clientset) error {
	ingressClient := clientSet.ExtensionsV1beta1().Ingresses(apiv1.NamespaceDefault)

	ingress := &v1beta.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: ingressData.Name,
		},
		Spec: v1beta.IngressSpec{
			Rules: []v1beta.IngressRule{
				{
					Host: ingressData.SubDomain,
					//TODO continue ingress implementation
				},
			},
		},
	}

	_, err := ingressClient.Create(ingress)
	if err != nil {
		return err
	}
	return nil
}
