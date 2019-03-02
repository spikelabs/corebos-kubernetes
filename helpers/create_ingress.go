package helpers

import (
	pb "corebos-kubernetes/kubernetes"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1beta "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func CreateIngress (ingressData *pb.Ingress, clientSet *kubernetes.Clientset) error {
	ingressClient := clientSet.ExtensionsV1beta1().Ingresses(apiv1.NamespaceDefault)

	ingress := &v1beta.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: ingressData.Name,
		},
		Spec: v1beta.IngressSpec{
			Rules: []v1beta.IngressRule{
				{
					Host: ingressData.SubDomain,
					IngressRuleValue: v1beta.IngressRuleValue{
						HTTP: &v1beta.HTTPIngressRuleValue{
							Paths: []v1beta.HTTPIngressPath{
								{
									Path: "/",
									Backend: v1beta.IngressBackend{
										ServiceName: ingressData.Resource,
										ServicePort: intstr.IntOrString{
											IntVal: 80,
											Type: intstr.Int,
										},
									},
								},
							},
						},
					},
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
